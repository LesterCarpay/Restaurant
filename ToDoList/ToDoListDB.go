package ToDoList

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
)

type ToDoListDB struct {
	db *sql.DB
}

func (tdl *ToDoListDB) initialize() {
	if tdl.db != nil {
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Password:")
	password, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	connectionString := "postgresql://postgres:" +
		password[:len(password)-1] +
		"@localhost/todos?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Failed to initialize database")
		log.Fatalln(err)
	}
	tdl.db = db
	fmt.Println("Succesfully loaded database.")
}

func (tdl *ToDoListDB) AddItem(newItem string) {
	_, err := tdl.db.Exec("INSERT into todos VALUES ($1)", newItem)
	if err != nil {
		fmt.Println("Failed to add item", newItem)
	}
}

func (tdl *ToDoListDB) DeleteItem(i int) {
	item := tdl.GetItem(i)
	_, err := tdl.db.Exec("DELETE from todos WHERE item=$1", item)
	if err != nil {
		fmt.Println("Failed to delete item", item)
	}

}

func (tdl *ToDoListDB) GetAllItems() []string {
	tdl.initialize()

	var item string
	var items []string

	rows, err := tdl.db.Query("SELECT * FROM todos")
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(rows)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		rows.Scan(&item)
		items = append(items, item)
	}
	return items
}

func (tdl *ToDoListDB) GetItem(i int) string {
	return tdl.GetAllItems()[i]
}
