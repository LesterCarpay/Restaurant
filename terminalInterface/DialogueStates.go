package terminalInterface

import (
	"fmt"
	"strconv"
)

func (dialogueManager *dialogueManager) showHome() {
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

func (dialogueManager *dialogueManager) showManageIngredients() {
	ingredients, err := dialogueManager.RestaurantDataBase.GetIngredients()
	if err != nil {
		fmt.Println("Could not get ingredients, error message:")
		fmt.Println(err)
		fmt.Println("Returning home.")
		dialogueManager.CurrentState = home
		return
	}
	if len(ingredients) < 1 {
		fmt.Println("Currently no ingredients exist.")
	} else {
		fmt.Println("Currently the following ingredients exist:")
		for _, ingredient := range ingredients {
			fmt.Printf("-%v\n", ingredient)
		}
	}
	fmt.Println("What would you like to do?")
	chosenOption := showChoiceMenu([]string{"Add ingredients",
		"Delete ingredients", "Return home"})
	switch chosenOption {
	case 1:
		dialogueManager.CurrentState = addIngredient
	case 2:
		dialogueManager.CurrentState = deleteIngredient
	default:
		dialogueManager.CurrentState = home
	}
}

func (dialogueManager *dialogueManager) showIngredientAdd() {
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

func (dialogueManager *dialogueManager) showIngredientDelete() {
	items, err := dialogueManager.RestaurantDataBase.GetIngredients()
	handleError(err)
	if len(items) < 1 {
		fmt.Println("No ingredients to delete. Returning home.")
		dialogueManager.CurrentState = home
		return
	}
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
