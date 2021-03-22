// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/contrib/entproto/internal/entprototest/ent/blogpost"
	"entgo.io/contrib/entproto/internal/entprototest/ent/user"
	"entgo.io/ent/dialect/sql"
)

// BlogPost is the model entity for the BlogPost schema.
type BlogPost struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// ExternalID holds the value of the "external_id" field.
	ExternalID int `json:"external_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BlogPostQuery when eager-loading is set.
	Edges            BlogPostEdges `json:"edges"`
	blog_post_author *int
}

// BlogPostEdges holds the relations/edges for other nodes in the graph.
type BlogPostEdges struct {
	// Author holds the value of the author edge.
	Author *User `json:"author,omitempty"`
	// Categories holds the value of the categories edge.
	Categories []*Category `json:"categories,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BlogPostEdges) AuthorOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Author == nil {
			// The edge author was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Author, nil
	}
	return nil, &NotLoadedError{edge: "author"}
}

// CategoriesOrErr returns the Categories value or an error if the edge
// was not loaded in eager-loading.
func (e BlogPostEdges) CategoriesOrErr() ([]*Category, error) {
	if e.loadedTypes[1] {
		return e.Categories, nil
	}
	return nil, &NotLoadedError{edge: "categories"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BlogPost) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case blogpost.FieldID, blogpost.FieldExternalID:
			values[i] = &sql.NullInt64{}
		case blogpost.FieldTitle, blogpost.FieldBody:
			values[i] = &sql.NullString{}
		case blogpost.ForeignKeys[0]: // blog_post_author
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type BlogPost", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BlogPost fields.
func (bp *BlogPost) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blogpost.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bp.ID = int(value.Int64)
		case blogpost.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				bp.Title = value.String
			}
		case blogpost.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				bp.Body = value.String
			}
		case blogpost.FieldExternalID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field external_id", values[i])
			} else if value.Valid {
				bp.ExternalID = int(value.Int64)
			}
		case blogpost.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field blog_post_author", value)
			} else if value.Valid {
				bp.blog_post_author = new(int)
				*bp.blog_post_author = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAuthor queries the "author" edge of the BlogPost entity.
func (bp *BlogPost) QueryAuthor() *UserQuery {
	return (&BlogPostClient{config: bp.config}).QueryAuthor(bp)
}

// QueryCategories queries the "categories" edge of the BlogPost entity.
func (bp *BlogPost) QueryCategories() *CategoryQuery {
	return (&BlogPostClient{config: bp.config}).QueryCategories(bp)
}

// Update returns a builder for updating this BlogPost.
// Note that you need to call BlogPost.Unwrap() before calling this method if this BlogPost
// was returned from a transaction, and the transaction was committed or rolled back.
func (bp *BlogPost) Update() *BlogPostUpdateOne {
	return (&BlogPostClient{config: bp.config}).UpdateOne(bp)
}

// Unwrap unwraps the BlogPost entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bp *BlogPost) Unwrap() *BlogPost {
	tx, ok := bp.config.driver.(*txDriver)
	if !ok {
		panic("ent: BlogPost is not a transactional entity")
	}
	bp.config.driver = tx.drv
	return bp
}

// String implements the fmt.Stringer.
func (bp *BlogPost) String() string {
	var builder strings.Builder
	builder.WriteString("BlogPost(")
	builder.WriteString(fmt.Sprintf("id=%v", bp.ID))
	builder.WriteString(", title=")
	builder.WriteString(bp.Title)
	builder.WriteString(", body=")
	builder.WriteString(bp.Body)
	builder.WriteString(", external_id=")
	builder.WriteString(fmt.Sprintf("%v", bp.ExternalID))
	builder.WriteByte(')')
	return builder.String()
}

// BlogPosts is a parsable slice of BlogPost.
type BlogPosts []*BlogPost

func (bp BlogPosts) config(cfg config) {
	for _i := range bp {
		bp[_i].config = cfg
	}
}
