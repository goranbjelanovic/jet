package postgres

import "github.com/goranbjelanovic/jet/v2/internal/jet"

// Statement is common interface for all statements(SELECT, INSERT, UPDATE, DELETE, LOCK)
type Statement = jet.Statement

// Projection is interface for all projection types. Types that can be part of, for instance SELECT clause.
type Projection = jet.Projection

// ProjectionList can be used to create conditional constructed projection list.
type ProjectionList = jet.ProjectionList

// ColumnAssigment is interface wrapper around column assigment
type ColumnAssigment = jet.ColumnAssigment

// PrintableStatement is a statement which sql query can be logged
type PrintableStatement = jet.PrintableStatement

// SetLogger sets automatic statement logging
var SetLogger = jet.SetLoggerFunc
