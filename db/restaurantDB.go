package db

import (
	"strconv"
)

// RestaurantDataBase represents a postgreSQL database for a restaurant with relations for
// ingredients, menu items, and menu item ingredients, which encodes which menu items contain
// which ingredients.
type RestaurantDataBase struct {
	Database            Database
	Ingredients         Table
	MenuItems           Table
	MenuItemIngredients Table
}

// GetRestaurantDatabase accepts a connection string and returns a RestaurantDataBase instance
// with the necessary tables and a connection to the postgreSQL database.
func GetRestaurantDatabase(connectionString string) (RestaurantDataBase, error) {
	var RestaurantDataBase RestaurantDataBase
	RestaurantDataBase.Ingredients = Table{
		Name:         "ingredients",
		IDColumnName: "ingredient_id",
		OtherColumns: map[string]string{"ingredient_name": "text"}}
	RestaurantDataBase.MenuItems = Table{
		Name:         "menu_items",
		IDColumnName: "menu_item_id",
		OtherColumns: map[string]string{"menu_item_name": "text",
			"menu_item_description": "text"}}
	RestaurantDataBase.MenuItemIngredients = Table{
		Name:         "menu_item_ingredients",
		IDColumnName: "menu_item_ingredient_id",
		OtherColumns: map[string]string{"menu_item_id": "INT",
			"ingredient_id": "INT"}}
	err := RestaurantDataBase.Database.ConnectToDatabase(connectionString)
	return RestaurantDataBase, err
}

// GetTables returns the tables of the database.
func (restaurantDataBase RestaurantDataBase) GetTables() []Table {
	return []Table{restaurantDataBase.Ingredients,
		restaurantDataBase.MenuItems,
		restaurantDataBase.MenuItemIngredients}
}

// Ingredient methods

// GetIngredients returns all ingredients in the database.
func (restaurantDataBase RestaurantDataBase) GetIngredients() (map[int]string, error) {
	return restaurantDataBase.Database.GetColumnValues(restaurantDataBase.Ingredients, "ingredient_name")
}

// AddIngredient accepts an ingredient name and adds it to the database.
func (restaurantDataBase RestaurantDataBase) AddIngredient(newIngredient string) error {
	err := restaurantDataBase.Database.AddRowToTable(restaurantDataBase.Ingredients,
		map[string]string{"ingredient_name": newIngredient})
	return err
}

// DeleteIngredient accepts an ingredient id and deletes the corresponding ingredient from the database.
func (restaurantDataBase RestaurantDataBase) DeleteIngredient(id int) error {
	return restaurantDataBase.Database.DeleteItem(restaurantDataBase.Ingredients, id)
}

// Menu item methods

// GetMenuItems returns all menu items in the database.
func (restaurantDataBase RestaurantDataBase) GetMenuItems() (map[int]string, error) {
	return restaurantDataBase.Database.GetColumnValues(restaurantDataBase.MenuItems, "menu_item_name")
}

// GetMenuItemDescription accepts an id and returns the description of the menu item with that id.
func (restaurantDataBase RestaurantDataBase) GetMenuItemDescription(id int) (string, error) {
	descriptions, err := restaurantDataBase.Database.GetColumnValues(restaurantDataBase.MenuItems, "menu_item_description")
	if err != nil {
		return "", err
	}
	return descriptions[id], nil
}

// AddMenuItem accepts an item name and description and creates a corresponding menu item entry in the database.
func (restaurantDataBase RestaurantDataBase) AddMenuItem(newItemName string, newItemDescription string) error {
	err := restaurantDataBase.Database.AddRowToTable(restaurantDataBase.MenuItems,
		map[string]string{"menu_item_name": newItemName, "menu_item_description": newItemDescription})
	return err
}

// DeleteMenuItem accepts an id and deletes the menu item with that id from the database.
func (restaurantDataBase RestaurantDataBase) DeleteMenuItem(id int) error {
	return restaurantDataBase.Database.DeleteItem(restaurantDataBase.MenuItems, id)
}

// ChangeMenuItemDescription accepts an id and a new description and changes the description of the corresponding
// menu item to the new description.
func (restaurantDataBase RestaurantDataBase) ChangeMenuItemDescription(id int, newValue string) error {
	return restaurantDataBase.Database.ChangeRowValue(restaurantDataBase.MenuItems,
		"menu_item_description", id, newValue)
}

//Menu item ingredient methods

// AddIngredientToMenuItem accepts the ids of a menu item and an ingredient and creates a menu item ingredient
// connection between the corresponding entries.
func (restaurantDataBase RestaurantDataBase) AddIngredientToMenuItem(menuItemID int, ingredientID int) error {
	return restaurantDataBase.Database.AddRowToTable(restaurantDataBase.MenuItemIngredients,
		map[string]string{"menu_item_id": strconv.Itoa(menuItemID), "ingredient_id": strconv.Itoa(ingredientID)})
}

// GetIngredientsOfMenuItem accepts the id of a menu item and returns the names of the ingredients it contains.
func (restaurantDataBase RestaurantDataBase) GetIngredientsOfMenuItem(menuItemID int) (map[int]string, error) {
	return restaurantDataBase.Database.getCompositionElements("ingredient_name", restaurantDataBase.MenuItems, restaurantDataBase.Ingredients,
		restaurantDataBase.MenuItemIngredients, menuItemID)
}

// DeleteIngredientFromMenuItem accepts the id of a menu item ingredient row and deletes it from the table.
func (restaurantDataBase RestaurantDataBase) DeleteIngredientFromMenuItem(ingredientMenuItemID int) error {
	return restaurantDataBase.Database.DeleteItem(restaurantDataBase.MenuItemIngredients, ingredientMenuItemID)
}
