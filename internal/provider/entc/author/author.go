// Code generated by entc, DO NOT EDIT.

package author

import (
	"time"
)

const (
	// Label holds the string label denoting the author type in the database.
	Label = "author"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldFullName holds the string denoting the fullname field in the database.
	FieldFullName = "full_name"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"

	// EdgeFilms holds the string denoting the films edge name in mutations.
	EdgeFilms = "films"

	// Table holds the table name of the author in the database.
	Table = "authors"
	// FilmsTable is the table the holds the films relation/edge. The primary key declared below.
	FilmsTable = "film_authors"
	// FilmsInverseTable is the table name for the Film entity.
	// It exists in this package in order to avoid circular dependency with the "film" package.
	FilmsInverseTable = "films"
)

// Columns holds all SQL columns for author fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldFullName,
	FieldAge,
}

var (
	// FilmsPrimaryKey and FilmsColumn2 are the table columns denoting the
	// primary key for the films relation (M2M).
	FilmsPrimaryKey = []string{"film_id", "author_id"}
)

var (
	// DefaultStatus holds the default value on creation for the status field.
	DefaultStatus int8
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
)