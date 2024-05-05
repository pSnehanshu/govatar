// Code generated by ent, DO NOT EDIT.

package avatar

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the avatar type in the database.
	Label = "avatar"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldLoc holds the string denoting the loc field in the database.
	FieldLoc = "loc"
	// FieldRating holds the string denoting the rating field in the database.
	FieldRating = "rating"
	// FieldEmailID holds the string denoting the email_id field in the database.
	FieldEmailID = "email_id"
	// EdgeEmail holds the string denoting the email edge name in mutations.
	EdgeEmail = "email"
	// Table holds the table name of the avatar in the database.
	Table = "avatars"
	// EmailTable is the table that holds the email relation/edge.
	EmailTable = "avatars"
	// EmailInverseTable is the table name for the Email entity.
	// It exists in this package in order to avoid circular dependency with the "email" package.
	EmailInverseTable = "emails"
	// EmailColumn is the table column denoting the email relation/edge.
	EmailColumn = "email_id"
)

// Columns holds all SQL columns for avatar fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldLoc,
	FieldRating,
	FieldEmailID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Rating defines the type for the "rating" enum field.
type Rating string

// RatingG is the default value of the Rating enum.
const DefaultRating = RatingG

// Rating values.
const (
	RatingG  Rating = "G"
	RatingPG Rating = "PG"
	RatingR  Rating = "R"
	RatingX  Rating = "X"
)

func (r Rating) String() string {
	return string(r)
}

// RatingValidator is a validator for the "rating" field enum values. It is called by the builders before save.
func RatingValidator(r Rating) error {
	switch r {
	case RatingG, RatingPG, RatingR, RatingX:
		return nil
	default:
		return fmt.Errorf("avatar: invalid enum value for rating field: %q", r)
	}
}

// OrderOption defines the ordering options for the Avatar queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByLoc orders the results by the loc field.
func ByLoc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLoc, opts...).ToFunc()
}

// ByRating orders the results by the rating field.
func ByRating(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRating, opts...).ToFunc()
}

// ByEmailID orders the results by the email_id field.
func ByEmailID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailID, opts...).ToFunc()
}

// ByEmailField orders the results by email field.
func ByEmailField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEmailStep(), sql.OrderByField(field, opts...))
	}
}
func newEmailStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EmailInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, EmailTable, EmailColumn),
	)
}