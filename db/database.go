/*
Package db implements simple, common operations with a PostgreSQL database.
It implements methods for the specific use case of a restaurant database.
*/
package db

import (
	"database/sql"
	"fmt"
	"slices"
)

// Table corresponds to and implements functionality of a relational database table.
type Table struct {
	Name         string
	IDColumnName string
	OtherColumns map[string]string
}

// Database connects to an SQL database and its methods implement common database operations.
type Database struct {
	sqlDB *sql.DB
}

// ConnectToDatabase accepts a connection string and establishes an active connection
// between a database struct and the SQL database.
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

// DeleteTable accepts a table struct and deletes the equivalent table from the SQL database, if it exists.
func (db *Database) DeleteTable(table Table) error {
	_, err := db.sqlDB.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %v", table.Name))
	if err != nil {
		return err
	}
	return nil
}

// CreateTable accepts a table struct and creates an equivalent table in the SQL database.
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

// TableExists accepts a table struct and returns true if a table with the same name exists
// in the SQL database, false if not.
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

// ColumnsExist accepts a table struct and returns true if the table of the same name
// in the SQL database has the columns that the table struct has, false if not.
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

// AddRowToTable accepts a table and a column-value mapping and adds a row
// with the specified values to the database.
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

// DeleteItem accepts a table and row id and deletes the row with that id from the database.
func (db *Database) DeleteItem(table Table, id int) error {
	_, err := db.sqlDB.Exec(fmt.Sprintf("DELETE from %v WHERE %v=%v",
		table.Name, table.IDColumnName, id))
	return err
}

// GetColumnValues accepts a table and column name and returns the row values corresponding to
// that column from the database as a slice of strings.
func (db *Database) GetColumnValues(table Table, columnName string) (map[int]string, error) {
	var (
		item   string
		itemID int
	)
	items := make(map[int]string)

	rows, err := db.sqlDB.Query(fmt.Sprintf("SELECT %v, %v FROM %v",
		table.IDColumnName, columnName, table.Name))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&itemID, &item)
		if err != nil {
			return nil, err
		}
		items[itemID] = item
	}
	return items, nil
}

// ChangeRowValue accepts a table, column name, id and new value, and changes the
// corresponding value in the database to the new value.
func (db *Database) ChangeRowValue(table Table, col string, id int, newValue string) error {
	_, err := db.sqlDB.Exec(fmt.Sprintf("UPDATE %v SET %v = '%v' WHERE %v = %v",
		table.Name, col, newValue, table.IDColumnName, id))
	if err != nil {
		return err
	}
	return nil
}

// getCompositionElements accepts the composing table in a composition relationship, the element table in this
// relationship and the table that encodes the composition relation, as well as the id of the
// row of the composer table for which elements should be returned. The function returns the ids of the elements that
// the row of the composer table contains as a dictionary of the relationship id and the element ids.
func (db *Database) getCompositionElements(composerTable Table, elementTable Table,
	compositionTable Table, rowID int) (map[int]int, error) {
	query := fmt.Sprintf("SELECT %v, %v "+
		"FROM %v "+
		"INNER JOIN %v "+
		"ON %v = %v "+
		"WHERE %v = %v",
		compositionTable.Name+"."+compositionTable.IDColumnName,
		elementTable.Name+"."+elementTable.IDColumnName,
		compositionTable.Name,
		elementTable.Name,
		compositionTable.Name+"."+elementTable.IDColumnName,
		elementTable.Name+"."+elementTable.IDColumnName,
		composerTable.IDColumnName,
		rowID)

	var (
		compositionRelationshipID int
		itemID                    int
	)
	items := make(map[int]int)

	rows, err := db.sqlDB.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&compositionRelationshipID, &itemID)
		if err != nil {
			return nil, err
		}
		items[compositionRelationshipID] = itemID
	}
	return items, nil
}
