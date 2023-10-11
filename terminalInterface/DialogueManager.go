/*
Package terminalInterface implements a dialogue loop in the terminal for a restaurant database.
An instance of DialogueManager needs to be created and StartDialogue needs to be called on it to initiate the dialogue.
*/
package terminalInterface

import (
	"Restaurant/db"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// dialogueState implements constants for the different state the dialogue can be in.
type dialogueState int

const (
	quit              dialogueState = iota
	home
	manageIngredients
	addIngredient
	deleteIngredient
	manageMenuItems
	addMenuItem
	deleteMenuItem
	changeMenuItem
)

// dialogueManager features fields and encodes behavior for managing the dialogue loop.
type dialogueManager struct {
	RestaurantDataBase db.RestaurantDataBase
	CurrentState       dialogueState
}

// handleError stops the program and prints an error log in case of an error.
func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// StartDialogue initiates the dialogue loop.
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

// checkTables checks whether all tables in the dialogue manager's RestaurantDatabase exist.
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

// loopDialogue continuously calls for the next window to be shown until the quit state is reached.
func (dialogueManager *dialogueManager) loopDialogue() {
	for dialogueManager.CurrentState != quit {
		dialogueManager.showNextWindow()
	}
}

// showNextWindow calls for the next window of the appropriate state to be showns.
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
	case changeMenuItem:
		dialogueManager.showChangeMenuItem()
	}
}
