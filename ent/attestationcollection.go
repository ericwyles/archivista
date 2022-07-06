// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/testifysec/archivist/ent/attestationcollection"
	"github.com/testifysec/archivist/ent/statement"
)

// AttestationCollection is the model entity for the AttestationCollection schema.
type AttestationCollection struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AttestationCollectionQuery when eager-loading is set.
	Edges                             AttestationCollectionEdges `json:"edges"`
	statement_attestation_collections *int
}

// AttestationCollectionEdges holds the relations/edges for other nodes in the graph.
type AttestationCollectionEdges struct {
	// Attestations holds the value of the attestations edge.
	Attestations []*Attestation `json:"attestations,omitempty"`
	// Statement holds the value of the statement edge.
	Statement *Statement `json:"statement,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AttestationsOrErr returns the Attestations value or an error if the edge
// was not loaded in eager-loading.
func (e AttestationCollectionEdges) AttestationsOrErr() ([]*Attestation, error) {
	if e.loadedTypes[0] {
		return e.Attestations, nil
	}
	return nil, &NotLoadedError{edge: "attestations"}
}

// StatementOrErr returns the Statement value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttestationCollectionEdges) StatementOrErr() (*Statement, error) {
	if e.loadedTypes[1] {
		if e.Statement == nil {
			// The edge statement was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: statement.Label}
		}
		return e.Statement, nil
	}
	return nil, &NotLoadedError{edge: "statement"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AttestationCollection) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case attestationcollection.FieldID:
			values[i] = new(sql.NullInt64)
		case attestationcollection.FieldName:
			values[i] = new(sql.NullString)
		case attestationcollection.ForeignKeys[0]: // statement_attestation_collections
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AttestationCollection", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AttestationCollection fields.
func (ac *AttestationCollection) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attestationcollection.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ac.ID = int(value.Int64)
		case attestationcollection.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ac.Name = value.String
			}
		case attestationcollection.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field statement_attestation_collections", value)
			} else if value.Valid {
				ac.statement_attestation_collections = new(int)
				*ac.statement_attestation_collections = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAttestations queries the "attestations" edge of the AttestationCollection entity.
func (ac *AttestationCollection) QueryAttestations() *AttestationQuery {
	return (&AttestationCollectionClient{config: ac.config}).QueryAttestations(ac)
}

// QueryStatement queries the "statement" edge of the AttestationCollection entity.
func (ac *AttestationCollection) QueryStatement() *StatementQuery {
	return (&AttestationCollectionClient{config: ac.config}).QueryStatement(ac)
}

// Update returns a builder for updating this AttestationCollection.
// Note that you need to call AttestationCollection.Unwrap() before calling this method if this AttestationCollection
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AttestationCollection) Update() *AttestationCollectionUpdateOne {
	return (&AttestationCollectionClient{config: ac.config}).UpdateOne(ac)
}

// Unwrap unwraps the AttestationCollection entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *AttestationCollection) Unwrap() *AttestationCollection {
	tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: AttestationCollection is not a transactional entity")
	}
	ac.config.driver = tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AttestationCollection) String() string {
	var builder strings.Builder
	builder.WriteString("AttestationCollection(")
	builder.WriteString(fmt.Sprintf("id=%v", ac.ID))
	builder.WriteString(", name=")
	builder.WriteString(ac.Name)
	builder.WriteByte(')')
	return builder.String()
}

// AttestationCollections is a parsable slice of AttestationCollection.
type AttestationCollections []*AttestationCollection

func (ac AttestationCollections) config(cfg config) {
	for _i := range ac {
		ac[_i].config = cfg
	}
}