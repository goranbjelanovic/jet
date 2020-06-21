//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package view

import (
	"github.com/goranbjelanovic/jet/v2/postgres"
)

var ActorInfo = newActorInfoTable()

type actorInfoTable struct {
	postgres.Table

	//Columns
	ActorID   postgres.ColumnInteger
	FirstName postgres.ColumnString
	LastName  postgres.ColumnString
	FilmInfo  postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ActorInfoTable struct {
	actorInfoTable

	EXCLUDED actorInfoTable
}

// AS creates new ActorInfoTable with assigned alias
func (a *ActorInfoTable) AS(alias string) *ActorInfoTable {
	aliasTable := newActorInfoTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newActorInfoTable() *ActorInfoTable {
	return &ActorInfoTable{
		actorInfoTable: newActorInfoTableImpl("dvds", "actor_info"),
		EXCLUDED:       newActorInfoTableImpl("", "excluded"),
	}
}

func newActorInfoTableImpl(schemaName, tableName string) actorInfoTable {
	var (
		ActorIDColumn   = postgres.IntegerColumn("actor_id")
		FirstNameColumn = postgres.StringColumn("first_name")
		LastNameColumn  = postgres.StringColumn("last_name")
		FilmInfoColumn  = postgres.StringColumn("film_info")
		allColumns      = postgres.ColumnList{ActorIDColumn, FirstNameColumn, LastNameColumn, FilmInfoColumn}
		mutableColumns  = postgres.ColumnList{ActorIDColumn, FirstNameColumn, LastNameColumn, FilmInfoColumn}
	)

	return actorInfoTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		ActorID:   ActorIDColumn,
		FirstName: FirstNameColumn,
		LastName:  LastNameColumn,
		FilmInfo:  FilmInfoColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
