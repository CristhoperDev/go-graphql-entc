// Code generated by entc, DO NOT EDIT.

package entc

import (
	"fmt"
	"strings"
	"time"

	"github.com/CristhoperDev/go-graphql-entc/internal/provider/entc/author"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/google/uuid"
)

// Author is the model entity for the Author schema.
type Author struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// FullName holds the value of the "fullName" field.
	FullName string `json:"fullName,omitempty"`
	// Age holds the value of the "age" field.
	Age int8 `json:"age,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AuthorQuery when eager-loading is set.
	Edges AuthorEdges `json:"edges"`
}

// AuthorEdges holds the relations/edges for other nodes in the graph.
type AuthorEdges struct {
	// Films holds the value of the films edge.
	Films []*Film
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FilmsOrErr returns the Films value or an error if the edge
// was not loaded in eager-loading.
func (e AuthorEdges) FilmsOrErr() ([]*Film, error) {
	if e.loadedTypes[0] {
		return e.Films, nil
	}
	return nil, &NotLoadedError{edge: "films"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Author) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullInt64{},  // status
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullString{}, // fullName
		&sql.NullInt64{},  // age
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Author fields.
func (a *Author) assignValues(values ...interface{}) error {
	if m, n := len(values), len(author.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		a.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[0])
	} else if value.Valid {
		a.Status = int8(value.Int64)
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[1])
	} else if value.Valid {
		a.CreatedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[2])
	} else if value.Valid {
		a.UpdatedAt = value.Time
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field fullName", values[3])
	} else if value.Valid {
		a.FullName = value.String
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[4])
	} else if value.Valid {
		a.Age = int8(value.Int64)
	}
	return nil
}

// QueryFilms queries the films edge of the Author.
func (a *Author) QueryFilms() *FilmQuery {
	return (&AuthorClient{config: a.config}).QueryFilms(a)
}

// Update returns a builder for updating this Author.
// Note that, you need to call Author.Unwrap() before calling this method, if this Author
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Author) Update() *AuthorUpdateOne {
	return (&AuthorClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Author) Unwrap() *Author {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("entc: Author is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Author) String() string {
	var builder strings.Builder
	builder.WriteString("Author(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", a.Status))
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", fullName=")
	builder.WriteString(a.FullName)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", a.Age))
	builder.WriteByte(')')
	return builder.String()
}

// Authors is a parsable slice of Author.
type Authors []*Author

func (a Authors) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
