// Code generated by entc, DO NOT EDIT.

package invertindexcompressed

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/YadaYuki/omochi/app/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// PostingListCompressed applies equality check predicate on the "posting_list_compressed" field. It's identical to PostingListCompressedEQ.
func PostingListCompressed(v []byte) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPostingListCompressed), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.InvertIndexCompressed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.InvertIndexCompressed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.InvertIndexCompressed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.InvertIndexCompressed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// PostingListCompressedEQ applies the EQ predicate on the "posting_list_compressed" field.
func PostingListCompressedEQ(v []byte) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPostingListCompressed), v))
	})
}

// PostingListCompressedNEQ applies the NEQ predicate on the "posting_list_compressed" field.
func PostingListCompressedNEQ(v []byte) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPostingListCompressed), v))
	})
}

// PostingListCompressedIn applies the In predicate on the "posting_list_compressed" field.
func PostingListCompressedIn(vs ...[]byte) predicate.InvertIndexCompressed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPostingListCompressed), v...))
	})
}

// PostingListCompressedNotIn applies the NotIn predicate on the "posting_list_compressed" field.
func PostingListCompressedNotIn(vs ...[]byte) predicate.InvertIndexCompressed {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPostingListCompressed), v...))
	})
}

// PostingListCompressedGT applies the GT predicate on the "posting_list_compressed" field.
func PostingListCompressedGT(v []byte) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPostingListCompressed), v))
	})
}

// PostingListCompressedGTE applies the GTE predicate on the "posting_list_compressed" field.
func PostingListCompressedGTE(v []byte) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPostingListCompressed), v))
	})
}

// PostingListCompressedLT applies the LT predicate on the "posting_list_compressed" field.
func PostingListCompressedLT(v []byte) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPostingListCompressed), v))
	})
}

// PostingListCompressedLTE applies the LTE predicate on the "posting_list_compressed" field.
func PostingListCompressedLTE(v []byte) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPostingListCompressed), v))
	})
}

// HasTermRelated applies the HasEdge predicate on the "term_related" edge.
func HasTermRelated() predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TermRelatedTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TermRelatedTable, TermRelatedColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTermRelatedWith applies the HasEdge predicate on the "term_related" edge with a given conditions (other predicates).
func HasTermRelatedWith(preds ...predicate.Term) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TermRelatedInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TermRelatedTable, TermRelatedColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.InvertIndexCompressed) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.InvertIndexCompressed) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.InvertIndexCompressed) predicate.InvertIndexCompressed {
	return predicate.InvertIndexCompressed(func(s *sql.Selector) {
		p(s.Not())
	})
}
