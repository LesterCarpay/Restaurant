package db

import (
	"fmt"
	"strconv"
)

type RestaurantDataBase struct {
	Database            Database
	Ingredients         Table
	MenuItems           Table
	MenuItemIngredients Table
}

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

func (restaurantDataBase RestaurantDataBase) GetTables() []Table {
	return []Table{restaurantDataBase.Ingredients,
		restaurantDataBase.MenuItems,
		restaurantDataBase.MenuItemIngredients}
}

//Ingredient methods

func (restaurantDataBase RestaurantDataBase) GetIngredients() ([]string, error) {
	return restaurantDataBase.Database.GetColumnValues(restaurantDataBase.Ingredients, "ingredient_name")
}

func (restaurantDataBase RestaurantDataBase) GetIngredientIDs() ([]string, error) {
	return restaurantDataBase.Database.GetTableIndices(restaurantDataBase.Ingredients)
}

func (restaurantDataBase RestaurantDataBase) AddIngredient(newIngredient string) error {
	err := restaurantDataBase.Database.AddRowToTable(restaurantDataBase.Ingredients,
		map[string]string{"ingredient_name": newIngredient})
	return err
}

func (restaurantDataBase RestaurantDataBase) DeleteIngredient(id int) error {
	return restaurantDataBase.Database.DeleteItem(restaurantDataBase.Ingredients, id)
}

//Menu item methods

func (restaurantDataBase RestaurantDataBase) GetMenuItems() ([]string, error) {
	return restaurantDataBase.Database.GetColumnValues(restaurantDataBase.MenuItems, "menu_item_name")
}

func (restaurantDataBase RestaurantDataBase) GetMenuItemDescription(id int) (string, error) {
	return restaurantDataBase.Database.GetColumnValue(restaurantDataBase.MenuItems, "menu_item_description", id)
}

func (restaurantDataBase RestaurantDataBase) GetMenuItemIDs() ([]string, error) {
	return restaurantDataBase.Database.GetColumnValues(restaurantDataBase.MenuItems, "menu_item_id")
}

func (restaurantDataBase RestaurantDataBase) AddMenuItem(newItemName string, newItemDescription string) error {
	err := restaurantDataBase.Database.AddRowToTable(restaurantDataBase.MenuItems,
		map[string]string{"menu_item_name": newItemName, "menu_item_description": newItemDescription})
	return err
}

func (restaurantDataBase RestaurantDataBase) DeleteMenuItem(id int) error {
	return restaurantDataBase.Database.DeleteItem(restaurantDataBase.MenuItems, id)
}

func (restaurantDataBase RestaurantDataBase) ChangeMenuItemDescription(id int, newValue string) error {
	return restaurantDataBase.Database.ChangeRowValue(restaurantDataBase.MenuItems,
		"menu_item_description", id, newValue)
}

//Menu item ingredient methods

func (restaurantDataBase RestaurantDataBase) AddIngredientToMenuItem(menuItemID int, ingredientID int) error {
	return restaurantDataBase.Database.AddRowToTable(restaurantDataBase.MenuItemIngredients,
		map[string]string{"menu_item_id": strconv.Itoa(menuItemID), "ingredient_id": strconv.Itoa(ingredientID)})
}

func (restaurantDataBase RestaurantDataBase) GetIngredientsOfMenuItem(menuItemID int) ([]string, error) {
	query := fmt.Sprintf("SELECT ingredient_name "+
		"FROM menu_item_ingredients "+
		"INNER JOIN ingredients "+
		"ON menu_item_ingredients.ingredient_id = ingredients.ingredient_id "+
		"WHERE menu_item_id = %v", menuItemID)
	return restaurantDataBase.Database.getQueryResults(query)
}

//func (RestaurantDataBase RestaurantDataBase) DeleteIngredientFromMenuItem(IngredientMenuItemID int) error {
//
//}
