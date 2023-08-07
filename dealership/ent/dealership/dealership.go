// Copyright 2022-present Vlabs Development Kft
//
// All rights reserved under a proprietary license.
//
// Code generated by entc, DO NOT EDIT.

package dealership

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the dealership type in the database.
	Label = "dealership"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeCars holds the string denoting the cars edge name in mutations.
	EdgeCars = "cars"
	// Table holds the table name of the dealership in the database.
	Table = "dealerships"
	// CarsTable is the table that holds the cars relation/edge. The primary key declared below.
	CarsTable = "cars_dealership_cars"
	// CarsInverseTable is the table name for the Cars entity.
	// It exists in this package in order to avoid circular dependency with the "cars" package.
	CarsInverseTable = "cars"
)

// Columns holds all SQL columns for dealership fields.
var Columns = []string{
	FieldID,
	FieldCity,
	FieldName,
}

var (
	// CarsPrimaryKey and CarsColumn2 are the table columns denoting the
	// primary key for the cars relation (M2M).
	CarsPrimaryKey = []string{"cars_id", "dealership_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Dealership queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCity orders the results by the city field.
func ByCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCity, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCarsCount orders the results by cars count.
func ByCarsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCarsStep(), opts...)
	}
}

// ByCars orders the results by cars terms.
func ByCars(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCarsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCarsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CarsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, CarsTable, CarsPrimaryKey...),
	)
}