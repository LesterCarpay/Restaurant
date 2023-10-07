package terminalInterface

import (
	"Restaurant/db"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type dialogueState int

const (
	quit                      dialogueState = 0
	home                      dialogueState = 1
	manageIngredients         dialogueState = 2
	addIngredient             dialogueState = 3
	deleteIngredient          dialogueState = 4
	manageMenuItems           dialogueState = 5
	addMenuItem               dialogueState = 6
	deleteMenuItem            dialogueState = 7
	changeMenuItemDescription dialogueState = 8
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
	dialogueManager.CurrentState = home
	dialogueManager.loopDialogue()
}

func (dialogueManager *dialogueManager) checkTables() bool {
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

func (dialogueManager *dialogueManager) loopDialogue() {
	for dialogueManager.CurrentState != quit {
		dialogueManager.showNextWindow()
	}
}

func (dialogueManager *dialogueManager) showNextWindow() {
	switch dialogueManager.CurrentState {
	case home:
		dialogueManager.showHome()
	case manageIngredients:
		dialogueManager.showManageIngredients()
	case addIngredient:
		dialogueManager.showIngredientAdd()
	case deleteIngredient:
		dialogueManager.showIngredientDelete()
	case manageMenuItems:
		dialogueManager.showManageMenuItems()
	case addMenuItem:
		dialogueManager.showMenuItemAdd()
	case deleteMenuItem:
		dialogueManager.showMenuItemDelete()
	case changeMenuItemDescription:
		dialogueManager.showMenuItemChangeDescription()
	}
}
