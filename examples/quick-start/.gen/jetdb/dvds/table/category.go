//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/goranbjelanovic/jet/v2/postgres"
)

var Category = newCategoryTable()

type categoryTable struct {
	postgres.Table

	//Columns
	CategoryID postgres.ColumnInteger
	Name       postgres.ColumnString
	LastUpdate postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CategoryTable struct {
	categoryTable

	EXCLUDED categoryTable
}

// AS creates new CategoryTable with assigned alias
func (a *CategoryTable) AS(alias string) *CategoryTable {
	aliasTable := newCategoryTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newCategoryTable() *CategoryTable {
	return &CategoryTable{
		categoryTable: newCategoryTableImpl("dvds", "category"),
		EXCLUDED:      newCategoryTableImpl("", "excluded"),
	}
}

func newCategoryTableImpl(schemaName, tableName string) categoryTable {
	var (
		CategoryIDColumn = postgres.IntegerColumn("category_id")
		NameColumn       = postgres.StringColumn("name")
		LastUpdateColumn = postgres.TimestampColumn("last_update")
		allColumns       = postgres.ColumnList{CategoryIDColumn, NameColumn, LastUpdateColumn}
		mutableColumns   = postgres.ColumnList{NameColumn, LastUpdateColumn}
	)

	return categoryTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		CategoryID: CategoryIDColumn,
		Name:       NameColumn,
		LastUpdate: LastUpdateColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
