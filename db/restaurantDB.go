package db

type RestaurantDataBase struct {
	Database    Database
	Ingredients Table
	MenuItems   Table
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
	err := RestaurantDataBase.Database.ConnectToDatabase(connectionString)
	return RestaurantDataBase, err
}

func (RestaurantDataBase RestaurantDataBase) GetTables() []Table {
	return []Table{RestaurantDataBase.Ingredients, RestaurantDataBase.MenuItems}
}

func (RestaurantDataBase RestaurantDataBase) GetIngredients() ([]string, error) {
	return RestaurantDataBase.Database.GetColumnValues(RestaurantDataBase.Ingredients, "ingredient_name")
}

func (RestaurantDataBase RestaurantDataBase) GetIngredientIndices() ([]string, error) {
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
