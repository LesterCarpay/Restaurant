package db

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

func (RestaurantDataBase RestaurantDataBase) GetTables() []Table {
	return []Table{RestaurantDataBase.Ingredients, RestaurantDataBase.MenuItems}
}

//Ingredient methods

func (RestaurantDataBase RestaurantDataBase) GetIngredients() ([]string, error) {
	return RestaurantDataBase.Database.GetColumnValues(RestaurantDataBase.Ingredients, "ingredient_name")
}

func (RestaurantDataBase RestaurantDataBase) GetIngredientIDs() ([]string, error) {
	return RestaurantDataBase.Database.GetTableIndices(RestaurantDataBase.Ingredients)
}

func (RestaurantDataBase RestaurantDataBase) AddIngredient(newIngredient string) error {
	err := RestaurantDataBase.Database.AddRowToTable(RestaurantDataBase.Ingredients,
		map[string]string{"ingredient_name": newIngredient})
	return err
}

func (RestaurantDataBase RestaurantDataBase) DeleteIngredient(id int) error {
	return RestaurantDataBase.Database.DeleteItem(RestaurantDataBase.Ingredients, id)
}

//Menu item methods

func (RestaurantDataBase RestaurantDataBase) GetMenuItems() ([]string, error) {
	return RestaurantDataBase.Database.GetColumnValues(RestaurantDataBase.MenuItems, "menu_item_name")
}

func (RestaurantDataBase RestaurantDataBase) GetMenuItemDescriptions() ([]string, error) {
	return RestaurantDataBase.Database.GetColumnValues(RestaurantDataBase.MenuItems, "menu_item_description")
}

func (RestaurantDataBase RestaurantDataBase) GetMenuItemIDs() ([]string, error) {
	return RestaurantDataBase.Database.GetColumnValues(RestaurantDataBase.MenuItems, "menu_item_id")
}

func (RestaurantDataBase RestaurantDataBase) AddMenuItem(newItemName string, newItemDescription string) error {
	err := RestaurantDataBase.Database.AddRowToTable(RestaurantDataBase.MenuItems,
		map[string]string{"menu_item_name": newItemName, "menu_item_description": newItemDescription})
	return err
}

func (RestaurantDataBase RestaurantDataBase) DeleteMenuItem(id int) error {
	return RestaurantDataBase.Database.DeleteItem(RestaurantDataBase.MenuItems, id)
}

func (RestaurantDataBase RestaurantDataBase) ChangeMenuItemDescription(id int, newValue string) error {
	return RestaurantDataBase.Database.ChangeRowValue(RestaurantDataBase.MenuItems,
		"menu_item_description", id, newValue)
}
