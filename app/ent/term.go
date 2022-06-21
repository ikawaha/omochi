// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/YadaYuki/omochi/app/ent/term"
	"github.com/google/uuid"
)

// Term is the model entity for the Term schema.
type Term struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Word holds the value of the "word" field.
	Word string `json:"word,omitempty"`
	// PostingListCompressed holds the value of the "posting_list_compressed" field.
	PostingListCompressed []byte `json:"posting_list_compressed,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Term) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case term.FieldPostingListCompressed:
			values[i] = new([]byte)
		case term.FieldWord:
			values[i] = new(sql.NullString)
		case term.FieldCreatedAt, term.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case term.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Term", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Term fields.
func (t *Term) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case term.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case term.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case term.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case term.FieldWord:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field word", values[i])
			} else if value.Valid {
				t.Word = value.String
			}
		case term.FieldPostingListCompressed:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field posting_list_compressed", values[i])
			} else if value != nil {
				t.PostingListCompressed = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Term.
// Note that you need to call Term.Unwrap() before calling this method if this Term
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Term) Update() *TermUpdateOne {
	return (&TermClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Term entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Term) Unwrap() *Term {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Term is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Term) String() string {
	var builder strings.Builder
	builder.WriteString("Term(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", word=")
	builder.WriteString(t.Word)
	builder.WriteString(", posting_list_compressed=")
	builder.WriteString(fmt.Sprintf("%v", t.PostingListCompressed))
	builder.WriteByte(')')
	return builder.String()
}

// Terms is a parsable slice of Term.
type Terms []*Term

func (t Terms) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
