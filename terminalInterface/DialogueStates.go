package terminalInterface

import (
	"Restaurant/db"
	"fmt"
	"log"
	"strconv"
)

type dialogueState int

const (
	quit              dialogueState = 0
	home              dialogueState = 1
	manageIngredients dialogueState = 2
	addIngredient     dialogueState = 3
	deleteIngredient  dialogueState = 4
)

type dialogueManager struct {
	RestaurantDataBase db.RestaurantDataBase
	CurrentState       dialogueState
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func StartDialogue() {
	var dialogueManager dialogueManager
	connectionString := GetConnectionString()
	restaurantDatabase, err := db.GetRestaurantDatabase(connectionString)
	handleError(err)
	dialogueManager.RestaurantDataBase = restaurantDatabase
	if !dialogueManager.checkTables() {
		return
	}
	dialogueManager.loopDialogue()
}

func (dialogueManager dialogueManager) checkTables() bool {
	for _, table := range dialogueManager.RestaurantDataBase.GetTables() {
		tableExists, err := dialogueManager.RestaurantDataBase.Database.TableExists(table)
		handleError(err)
		if !tableExists {
			fmt.Printf("Table %v does not exist, creating it.\n", table.Name)
			err = dialogueManager.RestaurantDataBase.Database.CreateTable(table)
			handleError(err)
		}
		columnsExist, err := dialogueManager.RestaurantDataBase.Database.ColumnsExist(table)
		handleError(err)
		if !columnsExist {
			fmt.Printf("One or more of the columns of table %v does not exist.\n", table.Name)
			fmt.Println("Do you wish to delete the existing table and create it?")
			if showChoiceMenu([]string{"Yes", "No, quit"}) == 1 {
				err = dialogueManager.RestaurantDataBase.Database.DeleteTable(table)
				err = dialogueManager.RestaurantDataBase.Database.CreateTable(table)
			} else {
				return false
			}
		}
	}
	_, err := dialogueManager.RestaurantDataBase.GetIngredients()
	handleError(err)
	return true
}

func (dialogueManager dialogueManager) loopDialogue() {
	for dialogueManager.CurrentState != quit {
		dialogueManager.showNextWindow()
	}
}

func (dialogueManager dialogueManager) showNextWindow() {
	switch dialogueManager.CurrentState {
	case home:
		dialogueManager.showHome()
	case manageIngredients:
		dialogueManager.showManageIngredients()
	case addIngredient:
		dialogueManager.showIngredientAdd()
	case deleteIngredient:
		dialogueManager.showIngredientDelete()
	}
}

func (dialogueManager dialogueManager) showHome() {
	fmt.Println("Welcome to your restaurant management environment.")
	fmt.Println("What would you like to do?")
	var chosenOption int
	chosenOption = showChoiceMenu([]string{"Manage ingredients", "quit"})
	switch chosenOption {
	case 1:
		dialogueManager.CurrentState = manageIngredients
	default:
		dialogueManager.CurrentState = quit
	}
}

func (dialogueManager dialogueManager) showManageIngredients() {
	fmt.Println("Currently the following ingredients exist:")
	ingredients, err := dialogueManager.RestaurantDataBase.GetIngredients()
	if err != nil {
		fmt.Println("Could not get ingredients, error message:")
		fmt.Println(err)
		fmt.Println("Returning home.")
		dialogueManager.CurrentState = home
		return
	}
	for _, ingredient := range ingredients {
		fmt.Printf("-%v\n", ingredient)
	}
	fmt.Println("What would you like to do?")
	chosenOption := showChoiceMenu([]string{"Add ingredients.",
		"Delete ingredients", "Return home."})
	switch chosenOption {
	case 1:
		dialogueManager.CurrentState = addIngredient
	case 2:
		dialogueManager.CurrentState = deleteIngredient
	default:
		dialogueManager.CurrentState = home
	}
}

func (dialogueManager dialogueManager) showIngredientAdd() {
	fmt.Println("Adding new ingredient.")
	fmt.Print("Name: ")
	newIngredient := getUserInput()
	err := dialogueManager.RestaurantDataBase.AddIngredient(newIngredient)
	if err != nil {
		fmt.Printf("Failed to add new ingredient \"%v\". Error message:\n", newIngredient)
		fmt.Println(err)
		fmt.Println("Try again.")
		return
	}
	println("Added the item \"" + newIngredient + "\" to your list, do you want to add another item?")
	var chosenOption int
	chosenOption = showChoiceMenu([]string{"Yes", "No"})
	if chosenOption != 1 {
		dialogueManager.CurrentState = home
	}
}

func (dialogueManager dialogueManager) showIngredientDelete() {
	items, err := dialogueManager.RestaurantDataBase.GetIngredients()
	handleError(err)
	ids, err := dialogueManager.RestaurantDataBase.GetIngredientIndices()
	handleError(err)
	fmt.Println("Which ingredient would you like to delete?")
	chosenOption := showChoiceMenu(items) - 1
	id, _ := strconv.Atoi(ids[chosenOption])
	err = dialogueManager.RestaurantDataBase.DeleteIngredient(id)
	handleError(err)
	fmt.Println("Successfully removed item.")
	fmt.Println("Would you like to delete another element?")
	chosenOption = showChoiceMenu([]string{"Yes", "No"})
	if chosenOption != 1 {
		dialogueManager.CurrentState = home
	}
}
