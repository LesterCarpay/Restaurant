package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"slices"
)

type ToDoListDB struct {
	db *sql.DB
}

func scanWithDefault(name string, defaultValue string) string {
	fmt.Print(name + " [" + defaultValue + "]:")
	var result string
	_, err := fmt.Scanln(&result)
	if err != nil || result == "" {
		return defaultValue
	}
	return result

}

func getConnectionString() string {
	var password string

	host := scanWithDefault("Host", "localhost")
	database := scanWithDefault("Database", "todos")
	username := scanWithDefault("Username", "postgres")
	fmt.Print("Password:")
	_, err := fmt.Scanln(&password)
	if err != nil {
		fmt.Println("Incorrect password.")
		log.Fatalln(err)
	}
	return "postgresql://" +
		username + ":" +
		password + "@" +
		host + "/" +
		database + "?sslmode=disable"
}

func (tdl *ToDoListDB) ConnectToDatabase() {
	fmt.Println("Loading database...")

	if tdl.db != nil {
		return
	}
	connectionString := getConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Failed to initialize database")
		log.Fatalln(err)
	}
	tdl.db = db
	//tdl.GetAllItems()
	fmt.Println("Connected to database")
}

func (tdl *ToDoListDB) CreateToDosTable() {
	_, err := tdl.db.Exec("DROP TABLE IF EXISTS todos;" +
		"CREATE TABLE todos (todo_id INT GENERATED ALWAYS AS IDENTITY, item text)")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to create table.")
	}
}

func (tdl *ToDoListDB) TableExists() bool {
	result, err := tdl.db.Query("SELECT COUNT(*) " +
		"FROM INFORMATION_SCHEMA.TABLES " +
		"WHERE TABLE_NAME = 'todos'")
	if err != nil {
		log.Fatalln(err)
	}
	var rowCount int
	for result.Next() {
		err := result.Scan(&rowCount)
		if rowCount < 1 || err != nil {
			return false
		}
	}
	return true
}

func (tdl *ToDoListDB) ColumnsExist() bool {
	result, err := tdl.db.Query("SELECT column_name " +
		"FROM information_schema.columns " +
		"WHERE TABLE_NAME = 'todos'")
	if err != nil {
		log.Fatalln(err)
	}
	var column string
	var columns []string
	for result.Next() {
		err := result.Scan(&column)
		if err != nil {
			return false
		}
		columns = append(columns, column)
	}
	return slices.Contains(columns, "todo_id") &&
		slices.Contains(columns, "item")
}

func (tdl *ToDoListDB) AddItem(newItem string) {
	_, err := tdl.db.Exec("INSERT into todos (item) VALUES ($1)", newItem)
	if err != nil {
		fmt.Println("Failed to add item", newItem)
		log.Fatalln(err)
	}
}

func (tdl *ToDoListDB) DeleteItem(item_id int) {
	_, err := tdl.db.Exec("DELETE from todos WHERE todo_id=$1", item_id)
	if err != nil {
		log.Fatalln("Failed to delete item.")
	}
}

func (tdl *ToDoListDB) getColumnValues(column string) []string {
	var item string
	var items []string

	rows, err := tdl.db.Query(fmt.Sprintf("SELECT %v FROM todos", column))
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		err := rows.Scan(&item)
		if err != nil {
			print(err)
		}
		items = append(items, item)
	}
	return items
}

func (tdl *ToDoListDB) GetAllItems() []string {
	return tdl.getColumnValues("item")
}

func (tdl *ToDoListDB) GetIDs() []string {
	return tdl.getColumnValues("todo_id")
}

func (tdl *ToDoListDB) GetItem(i int) string {
	return tdl.GetAllItems()[i]
}
