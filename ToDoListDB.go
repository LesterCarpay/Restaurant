package main

import (
	"Restaurant/db"
	_ "github.com/lib/pq"
	"log"
)

type ToDoListDB struct {
	db db.Database
}

var toDosTable = db.Table{Name: "todos",
	IDColumnName: "item_id",
	OtherColumns: map[string]string{"item": "text"}}

func (tdl *ToDoListDB) ConnectToDatabase() {
	err := tdl.db.ConnectToDatabase()
	if err != nil {
		log.Fatalln(err)
	}
}

func (tdl *ToDoListDB) CreateToDosTable() {
	err := tdl.db.DeleteTable(toDosTable)
	if err != nil {
		log.Fatalln(err)
	}
	err = tdl.db.CreateTable(toDosTable)
	if err != nil {
		log.Fatalln(err)
	}
}

func (tdl *ToDoListDB) TableExists() bool {
	result, err := tdl.db.TableExists(toDosTable)
	if err != nil {
		log.Fatalln(err)
	}
	return result
}

func (tdl *ToDoListDB) ColumnsExist() bool {
	result, err := tdl.db.ColumnsExist(toDosTable)
	if err != nil {
		log.Fatalln(err)
	}
	return result
}

func (tdl *ToDoListDB) AddItem(newItem string) {
	newRow := map[string]string{"item": newItem}
	err := tdl.db.AddRowToTable(toDosTable, newRow)
	if err != nil {
		log.Fatalln(err)
	}
}

func (tdl *ToDoListDB) DeleteItem(itemID int) {
	err := tdl.db.DeleteItem(toDosTable, itemID)
	if err != nil {
		log.Fatalln(err)
	}
}

func (tdl *ToDoListDB) getColumnValues(column string) []string {
	values, err := tdl.db.GetColumnValues(toDosTable, column)
	if err != nil {
		log.Fatalln(err)
	}
	return values
}

func (tdl *ToDoListDB) GetAllItems() []string {
	return tdl.getColumnValues("item")
}

func (tdl *ToDoListDB) GetIDs() []string {
	return tdl.getColumnValues(toDosTable.IDColumnName)
}

func (tdl *ToDoListDB) GetItem(i int) string {
	return tdl.GetAllItems()[i]
}
