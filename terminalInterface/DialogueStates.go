package terminalInterface

import (
	"fmt"
	"slices"
	"strconv"
)

func (dialogueManager *dialogueManager) showHome() {
	fmt.Println("Welcome to your restaurant management environment.")
	fmt.Println("What would you like to do?")
	var chosenOption int
	chosenOption = showChoiceMenu([]string{"Manage ingredients", "Manage menu items", "quit"})
	switch chosenOption {
	case 1:
		dialogueManager.CurrentState = manageIngredients
	case 2:
		dialogueManager.CurrentState = manageMenuItems
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
	println("Added the ingredient \"" + newIngredient + "\" to your list, do you want to add another ingredient?")
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
	ids, err := dialogueManager.RestaurantDataBase.GetIngredientIDs()
	handleError(err)
	fmt.Println("Which ingredient would you like to delete?")
	chosenOption := showChoiceMenu(items) - 1
	id, _ := strconv.Atoi(ids[chosenOption])
	err = dialogueManager.RestaurantDataBase.DeleteIngredient(id)
	handleError(err)
	fmt.Println("Successfully removed ingredient.")
	fmt.Println("Would you like to delete another ingredient?")
	chosenOption = showChoiceMenu([]string{"Yes", "No"})
	if chosenOption != 1 {
		dialogueManager.CurrentState = home
	}
}

func (dialogueManager *dialogueManager) showManageMenuItems() {
	menuItems, err := dialogueManager.RestaurantDataBase.GetMenuItems()
	if err != nil {
		fmt.Println("Could not get menu items, error message:")
		fmt.Println(err)
		fmt.Println("Returning home.")
		dialogueManager.CurrentState = home
		return
	}
	if len(menuItems) < 1 {
		fmt.Println("Currently no menu items exist.")
	} else {
		fmt.Println("Currently the following menu items exist:")
		for _, menuItem := range menuItems {
			fmt.Printf("-%v\n", menuItem)
		}
	}
	fmt.Println("What would you like to do?")
	chosenOption := showChoiceMenu([]string{"Add menu item",
		"Delete menu item", "Change menu item", "Return home"})
	switch chosenOption {
	case 1:
		dialogueManager.CurrentState = addMenuItem
	case 2:
		dialogueManager.CurrentState = deleteMenuItem
	case 3:
		dialogueManager.CurrentState = changeMenuItem
	default:
		dialogueManager.CurrentState = home
	}
}

func (dialogueManager *dialogueManager) showMenuItemAdd() {
	fmt.Println("Adding new menu item.")
	fmt.Print("Name: ")
	newIngredientName := getUserInput()
	fmt.Print("Description: ")
	newIngredientDescription := getUserInput()
	err := dialogueManager.RestaurantDataBase.AddMenuItem(newIngredientName, newIngredientDescription)
	if err != nil {
		fmt.Printf("Failed to add new menu item \"%v\". Error message:\n", newIngredientName)
		fmt.Println(err)
		fmt.Println("Try again.")
		return
	}
	println("Added the item \"" + newIngredientName + "\" to your list, do you want to add another item?")
	var chosenOption int
	chosenOption = showChoiceMenu([]string{"Yes", "No"})
	if chosenOption != 1 {
		dialogueManager.CurrentState = home
	}
}

func (dialogueManager *dialogueManager) showMenuItemDelete() {
	items, err := dialogueManager.RestaurantDataBase.GetMenuItems()
	handleError(err)
	if len(items) < 1 {
		fmt.Println("No menu items to delete. Returning home.")
		dialogueManager.CurrentState = home
		return
	}
	ids, err := dialogueManager.RestaurantDataBase.GetMenuItemIDs()
	handleError(err)
	fmt.Println("Which menu item would you like to delete?")
	chosenOption := showChoiceMenu(items) - 1
	id, _ := strconv.Atoi(ids[chosenOption])
	err = dialogueManager.RestaurantDataBase.DeleteMenuItem(id)
	handleError(err)
	fmt.Println("Successfully removed item.")
	fmt.Println("Would you like to delete another item?")
	chosenOption = showChoiceMenu([]string{"Yes", "No"})
	if chosenOption != 1 {
		dialogueManager.CurrentState = home
	}
}

func (dialogueManager *dialogueManager) showChangeMenuItem() {
	items, err := dialogueManager.RestaurantDataBase.GetMenuItems()
	handleError(err)
	if len(items) < 1 {
		fmt.Println("No menu items to change. Returning home.")
		dialogueManager.CurrentState = home
		return
	}
	ids, err := dialogueManager.RestaurantDataBase.GetMenuItemIDs()

	handleError(err)
	fmt.Println("Which menu item would you like to change?")
	chosenOption := showChoiceMenu(items) - 1
	menuItemID, err := strconv.Atoi(ids[chosenOption])
	handleError(err)

	fmt.Println("Menu item:", items[chosenOption])
	description, err := dialogueManager.RestaurantDataBase.GetMenuItemDescription(menuItemID)
	handleError(err)
	fmt.Println("Description:", description)
	ingredients, err := dialogueManager.RestaurantDataBase.GetIngredientsOfMenuItem(menuItemID)
	handleError(err)
	fmt.Printf("Ingredients: %v\n", ingredients)
	fmt.Println("What would you like to change?")
	chosenOption = showChoiceMenu([]string{"Modify description", "Add ingredient"})
	if chosenOption == 1 {
		dialogueManager.showChangeMenuItemDescription(menuItemID)
	} else {
		dialogueManager.showAddIngredientToMenuItem(menuItemID)
	}
	fmt.Println("Would you like to make another change?")
	chosenOption = showChoiceMenu([]string{"Yes", "No, return home"})
	if chosenOption == 1 {
		dialogueManager.CurrentState = changeMenuItem
	} else {
		dialogueManager.CurrentState = home
	}
}

func (dialogueManager *dialogueManager) showChangeMenuItemDescription(menuItemID int) {
	fmt.Println("New menu item description:")
	newIngredientDescription := getUserInput()
	err := dialogueManager.RestaurantDataBase.ChangeMenuItemDescription(menuItemID, newIngredientDescription)
	handleError(err)
	fmt.Println("Successfully changed menu item description.")
}

func (dialogueManager *dialogueManager) showAddIngredientToMenuItem(menuItemID int) {
	ingredients, err := dialogueManager.RestaurantDataBase.GetIngredients()
	handleError(err)
	ingredientIDs, err := dialogueManager.RestaurantDataBase.GetIngredientIDs()
	handleError(err)
	currentIngredients, err := dialogueManager.RestaurantDataBase.GetIngredientsOfMenuItem(menuItemID)
	handleError(err)
	var ingredientOptions []string
	var ingredientOptionIDs []string
	for i, ingredient := range ingredients {
		if !slices.Contains(currentIngredients, ingredient) {
			ingredientOptions = append(ingredientOptions, ingredient)
			ingredientOptionIDs = append(ingredientOptionIDs, ingredientIDs[i])
		}
	}
	if len(ingredientOptions) < 1 {
		fmt.Println("There are no available ingredients to add.")
		return
	}
	fmt.Println("Which ingredient would you like to add?")
	ingredientID, err := strconv.Atoi(ingredientOptionIDs[showChoiceMenu(ingredientOptions)-1])
	handleError(err)
	err = dialogueManager.RestaurantDataBase.AddIngredientToMenuItem(menuItemID, ingredientID)
	handleError(err)
}
