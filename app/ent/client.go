// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/YadaYuki/omochi/app/ent/migrate"
	"github.com/google/uuid"

	"github.com/YadaYuki/omochi/app/ent/document"
	"github.com/YadaYuki/omochi/app/ent/term"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Document is the client for interacting with the Document builders.
	Document *DocumentClient
	// Term is the client for interacting with the Term builders.
	Term *TermClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Document = NewDocumentClient(c.config)
	c.Term = NewTermClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Document: NewDocumentClient(cfg),
		Term:     NewTermClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Document: NewDocumentClient(cfg),
		Term:     NewTermClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Document.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Document.Use(hooks...)
	c.Term.Use(hooks...)
}

// DocumentClient is a client for the Document schema.
type DocumentClient struct {
	config
}

// NewDocumentClient returns a client for the Document from the given config.
func NewDocumentClient(c config) *DocumentClient {
	return &DocumentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `document.Hooks(f(g(h())))`.
func (c *DocumentClient) Use(hooks ...Hook) {
	c.hooks.Document = append(c.hooks.Document, hooks...)
}

// Create returns a create builder for Document.
func (c *DocumentClient) Create() *DocumentCreate {
	mutation := newDocumentMutation(c.config, OpCreate)
	return &DocumentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Document entities.
func (c *DocumentClient) CreateBulk(builders ...*DocumentCreate) *DocumentCreateBulk {
	return &DocumentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Document.
func (c *DocumentClient) Update() *DocumentUpdate {
	mutation := newDocumentMutation(c.config, OpUpdate)
	return &DocumentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DocumentClient) UpdateOne(d *Document) *DocumentUpdateOne {
	mutation := newDocumentMutation(c.config, OpUpdateOne, withDocument(d))
	return &DocumentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DocumentClient) UpdateOneID(id int) *DocumentUpdateOne {
	mutation := newDocumentMutation(c.config, OpUpdateOne, withDocumentID(id))
	return &DocumentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Document.
func (c *DocumentClient) Delete() *DocumentDelete {
	mutation := newDocumentMutation(c.config, OpDelete)
	return &DocumentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DocumentClient) DeleteOne(d *Document) *DocumentDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DocumentClient) DeleteOneID(id int) *DocumentDeleteOne {
	builder := c.Delete().Where(document.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DocumentDeleteOne{builder}
}

// Query returns a query builder for Document.
func (c *DocumentClient) Query() *DocumentQuery {
	return &DocumentQuery{
		config: c.config,
	}
}

// Get returns a Document entity by its id.
func (c *DocumentClient) Get(ctx context.Context, id int) (*Document, error) {
	return c.Query().Where(document.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DocumentClient) GetX(ctx context.Context, id int) *Document {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DocumentClient) Hooks() []Hook {
	return c.hooks.Document
}

// TermClient is a client for the Term schema.
type TermClient struct {
	config
}

// NewTermClient returns a client for the Term from the given config.
func NewTermClient(c config) *TermClient {
	return &TermClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `term.Hooks(f(g(h())))`.
func (c *TermClient) Use(hooks ...Hook) {
	c.hooks.Term = append(c.hooks.Term, hooks...)
}

// Create returns a create builder for Term.
func (c *TermClient) Create() *TermCreate {
	mutation := newTermMutation(c.config, OpCreate)
	return &TermCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Term entities.
func (c *TermClient) CreateBulk(builders ...*TermCreate) *TermCreateBulk {
	return &TermCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Term.
func (c *TermClient) Update() *TermUpdate {
	mutation := newTermMutation(c.config, OpUpdate)
	return &TermUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TermClient) UpdateOne(t *Term) *TermUpdateOne {
	mutation := newTermMutation(c.config, OpUpdateOne, withTerm(t))
	return &TermUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TermClient) UpdateOneID(id uuid.UUID) *TermUpdateOne {
	mutation := newTermMutation(c.config, OpUpdateOne, withTermID(id))
	return &TermUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Term.
func (c *TermClient) Delete() *TermDelete {
	mutation := newTermMutation(c.config, OpDelete)
	return &TermDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TermClient) DeleteOne(t *Term) *TermDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TermClient) DeleteOneID(id uuid.UUID) *TermDeleteOne {
	builder := c.Delete().Where(term.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TermDeleteOne{builder}
}

// Query returns a query builder for Term.
func (c *TermClient) Query() *TermQuery {
	return &TermQuery{
		config: c.config,
	}
}

// Get returns a Term entity by its id.
func (c *TermClient) Get(ctx context.Context, id uuid.UUID) (*Term, error) {
	return c.Query().Where(term.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TermClient) GetX(ctx context.Context, id uuid.UUID) *Term {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TermClient) Hooks() []Hook {
	return c.hooks.Term
}
