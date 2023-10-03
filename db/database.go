package db

import (
	"database/sql"
	"fmt"
	"slices"
)

type Table struct {
	Name         string
	IDColumnName string
	OtherColumns map[string]string
}

type Database struct {
	sqlDB *sql.DB
}

func (db *Database) ConnectToDatabase(connectionString string) error {
	fmt.Println("Loading Database...")

	sqlDB, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	db.sqlDB = sqlDB
	//tdl.GetAllItems()
	fmt.Println("Connected to Database")
	return nil
}

func (db *Database) DeleteTable(table Table) error {
	_, err := db.sqlDB.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %v", table.Name))
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) CreateTable(table Table) error {
	_, err := db.sqlDB.Exec(fmt.Sprintf("CREATE TABLE %v "+
		"(%v INT GENERATED ALWAYS AS IDENTITY)", table.Name, table.IDColumnName))
	if err != nil {
		return err
	}
	for colName, colType := range table.OtherColumns {
		_, err = db.sqlDB.Exec(fmt.Sprintf("ALTER TABLE %v ADD %v %v",
			table.Name, colName, colType))
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *Database) TableExists(table Table) (bool, error) {
	result, err := db.sqlDB.Query(fmt.Sprintf(
		"SELECT COUNT(*) "+
			"FROM INFORMATION_SCHEMA.TABLES "+
			"WHERE TABLE_NAME = '%v'", table.Name))
	if err != nil {
		return false, err
	}
	var rowCount int
	for result.Next() {
		err := result.Scan(&rowCount)
		if err != nil {
			return false, err
		}
		if rowCount < 1 {
			return false, nil
		}
	}
	return true, nil
}

func (db *Database) ColumnsExist(table Table) (bool, error) {
	result, err := db.sqlDB.Query(fmt.Sprintf("SELECT column_name "+
		"FROM information_schema.columns "+
		"WHERE TABLE_NAME = '%v'", table.Name))
	if err != nil {
		return false, err
	}
	var column string
	var columns []string
	for result.Next() {
		err := result.Scan(&column)
		if err != nil {
			return false, err
		}
		columns = append(columns, column)
	}
	for colName := range table.OtherColumns {
		if !slices.Contains(columns, colName) {
			return false, nil
		}
	}
	if !slices.Contains(columns, table.IDColumnName) {
		return false, nil
	}
	return true, nil
}

func (db *Database) AddRowToTable(table Table, newRow map[string]string) error {
	colString := "("
	valString := "("
	for colName, value := range newRow {
		colString += colName + ","
		valString += fmt.Sprintf("'%v',", value)
	}
	colString = colString[:len(colString)-1] + ")"
	valString = valString[:len(valString)-1] + ")"
	_, err := db.sqlDB.Exec(fmt.Sprintf("INSERT into %v %v VALUES %v",
		table.Name, colString, valString))
	return err
}

func (db *Database) DeleteItem(table Table, id int) error {
	_, err := db.sqlDB.Exec(fmt.Sprintf("DELETE from %v WHERE %v=%v",
		table.Name, table.IDColumnName, id))
	return err
}

func (db *Database) GetColumnValues(table Table, column string) ([]string, error) {
	var item string
	var items []string

	rows, err := db.sqlDB.Query(fmt.Sprintf("SELECT %v FROM %v", column, table.Name))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (db *Database) GetTableIndices(table Table) ([]string, error) {
	return db.GetColumnValues(table, table.IDColumnName)
}
