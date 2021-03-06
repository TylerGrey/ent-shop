// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/TylerGrey/ent-shop/db/ent/migrate"

	"github.com/TylerGrey/ent-shop/db/ent/category"
	"github.com/TylerGrey/ent-shop/db/ent/delivery"
	"github.com/TylerGrey/ent-shop/db/ent/item"
	"github.com/TylerGrey/ent-shop/db/ent/member"
	"github.com/TylerGrey/ent-shop/db/ent/order"
	"github.com/TylerGrey/ent-shop/db/ent/orderitem"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Category is the client for interacting with the Category builders.
	Category *CategoryClient
	// Delivery is the client for interacting with the Delivery builders.
	Delivery *DeliveryClient
	// Item is the client for interacting with the Item builders.
	Item *ItemClient
	// Member is the client for interacting with the Member builders.
	Member *MemberClient
	// Order is the client for interacting with the Order builders.
	Order *OrderClient
	// OrderItem is the client for interacting with the OrderItem builders.
	OrderItem *OrderItemClient
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
	c.Category = NewCategoryClient(c.config)
	c.Delivery = NewDeliveryClient(c.config)
	c.Item = NewItemClient(c.config)
	c.Member = NewMemberClient(c.config)
	c.Order = NewOrderClient(c.config)
	c.OrderItem = NewOrderItemClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Category:  NewCategoryClient(cfg),
		Delivery:  NewDeliveryClient(cfg),
		Item:      NewItemClient(cfg),
		Member:    NewMemberClient(cfg),
		Order:     NewOrderClient(cfg),
		OrderItem: NewOrderItemClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:    cfg,
		Category:  NewCategoryClient(cfg),
		Delivery:  NewDeliveryClient(cfg),
		Item:      NewItemClient(cfg),
		Member:    NewMemberClient(cfg),
		Order:     NewOrderClient(cfg),
		OrderItem: NewOrderItemClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Category.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.Category.Use(hooks...)
	c.Delivery.Use(hooks...)
	c.Item.Use(hooks...)
	c.Member.Use(hooks...)
	c.Order.Use(hooks...)
	c.OrderItem.Use(hooks...)
}

// CategoryClient is a client for the Category schema.
type CategoryClient struct {
	config
}

// NewCategoryClient returns a client for the Category from the given config.
func NewCategoryClient(c config) *CategoryClient {
	return &CategoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `category.Hooks(f(g(h())))`.
func (c *CategoryClient) Use(hooks ...Hook) {
	c.hooks.Category = append(c.hooks.Category, hooks...)
}

// Create returns a create builder for Category.
func (c *CategoryClient) Create() *CategoryCreate {
	mutation := newCategoryMutation(c.config, OpCreate)
	return &CategoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Category entities.
func (c *CategoryClient) CreateBulk(builders ...*CategoryCreate) *CategoryCreateBulk {
	return &CategoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Category.
func (c *CategoryClient) Update() *CategoryUpdate {
	mutation := newCategoryMutation(c.config, OpUpdate)
	return &CategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CategoryClient) UpdateOne(ca *Category) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategory(ca))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CategoryClient) UpdateOneID(id int) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategoryID(id))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Category.
func (c *CategoryClient) Delete() *CategoryDelete {
	mutation := newCategoryMutation(c.config, OpDelete)
	return &CategoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CategoryClient) DeleteOne(ca *Category) *CategoryDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CategoryClient) DeleteOneID(id int) *CategoryDeleteOne {
	builder := c.Delete().Where(category.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CategoryDeleteOne{builder}
}

// Query returns a query builder for Category.
func (c *CategoryClient) Query() *CategoryQuery {
	return &CategoryQuery{config: c.config}
}

// Get returns a Category entity by its id.
func (c *CategoryClient) Get(ctx context.Context, id int) (*Category, error) {
	return c.Query().Where(category.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CategoryClient) GetX(ctx context.Context, id int) *Category {
	ca, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ca
}

// QueryParent queries the parent edge of a Category.
func (c *CategoryClient) QueryParent(ca *Category) *CategoryQuery {
	query := &CategoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(category.Table, category.FieldID, id),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, category.ParentTable, category.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Category.
func (c *CategoryClient) QueryChildren(ca *Category) *CategoryQuery {
	query := &CategoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(category.Table, category.FieldID, id),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, category.ChildrenTable, category.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryItems queries the items edge of a Category.
func (c *CategoryClient) QueryItems(ca *Category) *ItemQuery {
	query := &ItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(category.Table, category.FieldID, id),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, category.ItemsTable, category.ItemsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CategoryClient) Hooks() []Hook {
	return c.hooks.Category
}

// DeliveryClient is a client for the Delivery schema.
type DeliveryClient struct {
	config
}

// NewDeliveryClient returns a client for the Delivery from the given config.
func NewDeliveryClient(c config) *DeliveryClient {
	return &DeliveryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `delivery.Hooks(f(g(h())))`.
func (c *DeliveryClient) Use(hooks ...Hook) {
	c.hooks.Delivery = append(c.hooks.Delivery, hooks...)
}

// Create returns a create builder for Delivery.
func (c *DeliveryClient) Create() *DeliveryCreate {
	mutation := newDeliveryMutation(c.config, OpCreate)
	return &DeliveryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Delivery entities.
func (c *DeliveryClient) CreateBulk(builders ...*DeliveryCreate) *DeliveryCreateBulk {
	return &DeliveryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Delivery.
func (c *DeliveryClient) Update() *DeliveryUpdate {
	mutation := newDeliveryMutation(c.config, OpUpdate)
	return &DeliveryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DeliveryClient) UpdateOne(d *Delivery) *DeliveryUpdateOne {
	mutation := newDeliveryMutation(c.config, OpUpdateOne, withDelivery(d))
	return &DeliveryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DeliveryClient) UpdateOneID(id int) *DeliveryUpdateOne {
	mutation := newDeliveryMutation(c.config, OpUpdateOne, withDeliveryID(id))
	return &DeliveryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Delivery.
func (c *DeliveryClient) Delete() *DeliveryDelete {
	mutation := newDeliveryMutation(c.config, OpDelete)
	return &DeliveryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DeliveryClient) DeleteOne(d *Delivery) *DeliveryDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DeliveryClient) DeleteOneID(id int) *DeliveryDeleteOne {
	builder := c.Delete().Where(delivery.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DeliveryDeleteOne{builder}
}

// Query returns a query builder for Delivery.
func (c *DeliveryClient) Query() *DeliveryQuery {
	return &DeliveryQuery{config: c.config}
}

// Get returns a Delivery entity by its id.
func (c *DeliveryClient) Get(ctx context.Context, id int) (*Delivery, error) {
	return c.Query().Where(delivery.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DeliveryClient) GetX(ctx context.Context, id int) *Delivery {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// QueryOrder queries the order edge of a Delivery.
func (c *DeliveryClient) QueryOrder(d *Delivery) *OrderQuery {
	query := &OrderQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(delivery.Table, delivery.FieldID, id),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, delivery.OrderTable, delivery.OrderColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DeliveryClient) Hooks() []Hook {
	return c.hooks.Delivery
}

// ItemClient is a client for the Item schema.
type ItemClient struct {
	config
}

// NewItemClient returns a client for the Item from the given config.
func NewItemClient(c config) *ItemClient {
	return &ItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `item.Hooks(f(g(h())))`.
func (c *ItemClient) Use(hooks ...Hook) {
	c.hooks.Item = append(c.hooks.Item, hooks...)
}

// Create returns a create builder for Item.
func (c *ItemClient) Create() *ItemCreate {
	mutation := newItemMutation(c.config, OpCreate)
	return &ItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Item entities.
func (c *ItemClient) CreateBulk(builders ...*ItemCreate) *ItemCreateBulk {
	return &ItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Item.
func (c *ItemClient) Update() *ItemUpdate {
	mutation := newItemMutation(c.config, OpUpdate)
	return &ItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ItemClient) UpdateOne(i *Item) *ItemUpdateOne {
	mutation := newItemMutation(c.config, OpUpdateOne, withItem(i))
	return &ItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ItemClient) UpdateOneID(id int) *ItemUpdateOne {
	mutation := newItemMutation(c.config, OpUpdateOne, withItemID(id))
	return &ItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Item.
func (c *ItemClient) Delete() *ItemDelete {
	mutation := newItemMutation(c.config, OpDelete)
	return &ItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ItemClient) DeleteOne(i *Item) *ItemDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ItemClient) DeleteOneID(id int) *ItemDeleteOne {
	builder := c.Delete().Where(item.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ItemDeleteOne{builder}
}

// Query returns a query builder for Item.
func (c *ItemClient) Query() *ItemQuery {
	return &ItemQuery{config: c.config}
}

// Get returns a Item entity by its id.
func (c *ItemClient) Get(ctx context.Context, id int) (*Item, error) {
	return c.Query().Where(item.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ItemClient) GetX(ctx context.Context, id int) *Item {
	i, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return i
}

// QueryCategories queries the categories edge of a Item.
func (c *ItemClient) QueryCategories(i *Item) *CategoryQuery {
	query := &CategoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, id),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, item.CategoriesTable, item.CategoriesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ItemClient) Hooks() []Hook {
	return c.hooks.Item
}

// MemberClient is a client for the Member schema.
type MemberClient struct {
	config
}

// NewMemberClient returns a client for the Member from the given config.
func NewMemberClient(c config) *MemberClient {
	return &MemberClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `member.Hooks(f(g(h())))`.
func (c *MemberClient) Use(hooks ...Hook) {
	c.hooks.Member = append(c.hooks.Member, hooks...)
}

// Create returns a create builder for Member.
func (c *MemberClient) Create() *MemberCreate {
	mutation := newMemberMutation(c.config, OpCreate)
	return &MemberCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Member entities.
func (c *MemberClient) CreateBulk(builders ...*MemberCreate) *MemberCreateBulk {
	return &MemberCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Member.
func (c *MemberClient) Update() *MemberUpdate {
	mutation := newMemberMutation(c.config, OpUpdate)
	return &MemberUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MemberClient) UpdateOne(m *Member) *MemberUpdateOne {
	mutation := newMemberMutation(c.config, OpUpdateOne, withMember(m))
	return &MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MemberClient) UpdateOneID(id int) *MemberUpdateOne {
	mutation := newMemberMutation(c.config, OpUpdateOne, withMemberID(id))
	return &MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Member.
func (c *MemberClient) Delete() *MemberDelete {
	mutation := newMemberMutation(c.config, OpDelete)
	return &MemberDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MemberClient) DeleteOne(m *Member) *MemberDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MemberClient) DeleteOneID(id int) *MemberDeleteOne {
	builder := c.Delete().Where(member.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MemberDeleteOne{builder}
}

// Query returns a query builder for Member.
func (c *MemberClient) Query() *MemberQuery {
	return &MemberQuery{config: c.config}
}

// Get returns a Member entity by its id.
func (c *MemberClient) Get(ctx context.Context, id int) (*Member, error) {
	return c.Query().Where(member.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MemberClient) GetX(ctx context.Context, id int) *Member {
	m, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return m
}

// QueryOrders queries the orders edge of a Member.
func (c *MemberClient) QueryOrders(m *Member) *OrderQuery {
	query := &OrderQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, id),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.OrdersTable, member.OrdersColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MemberClient) Hooks() []Hook {
	return c.hooks.Member
}

// OrderClient is a client for the Order schema.
type OrderClient struct {
	config
}

// NewOrderClient returns a client for the Order from the given config.
func NewOrderClient(c config) *OrderClient {
	return &OrderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `order.Hooks(f(g(h())))`.
func (c *OrderClient) Use(hooks ...Hook) {
	c.hooks.Order = append(c.hooks.Order, hooks...)
}

// Create returns a create builder for Order.
func (c *OrderClient) Create() *OrderCreate {
	mutation := newOrderMutation(c.config, OpCreate)
	return &OrderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Order entities.
func (c *OrderClient) CreateBulk(builders ...*OrderCreate) *OrderCreateBulk {
	return &OrderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Order.
func (c *OrderClient) Update() *OrderUpdate {
	mutation := newOrderMutation(c.config, OpUpdate)
	return &OrderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrderClient) UpdateOne(o *Order) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrder(o))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrderClient) UpdateOneID(id int) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrderID(id))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Order.
func (c *OrderClient) Delete() *OrderDelete {
	mutation := newOrderMutation(c.config, OpDelete)
	return &OrderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OrderClient) DeleteOne(o *Order) *OrderDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OrderClient) DeleteOneID(id int) *OrderDeleteOne {
	builder := c.Delete().Where(order.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrderDeleteOne{builder}
}

// Query returns a query builder for Order.
func (c *OrderClient) Query() *OrderQuery {
	return &OrderQuery{config: c.config}
}

// Get returns a Order entity by its id.
func (c *OrderClient) Get(ctx context.Context, id int) (*Order, error) {
	return c.Query().Where(order.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrderClient) GetX(ctx context.Context, id int) *Order {
	o, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return o
}

// QueryOrderItems queries the orderItems edge of a Order.
func (c *OrderClient) QueryOrderItems(o *Order) *OrderItemQuery {
	query := &OrderItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, id),
			sqlgraph.To(orderitem.Table, orderitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, order.OrderItemsTable, order.OrderItemsColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDelivery queries the delivery edge of a Order.
func (c *OrderClient) QueryDelivery(o *Order) *DeliveryQuery {
	query := &DeliveryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, id),
			sqlgraph.To(delivery.Table, delivery.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, order.DeliveryTable, order.DeliveryColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OrderClient) Hooks() []Hook {
	return c.hooks.Order
}

// OrderItemClient is a client for the OrderItem schema.
type OrderItemClient struct {
	config
}

// NewOrderItemClient returns a client for the OrderItem from the given config.
func NewOrderItemClient(c config) *OrderItemClient {
	return &OrderItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `orderitem.Hooks(f(g(h())))`.
func (c *OrderItemClient) Use(hooks ...Hook) {
	c.hooks.OrderItem = append(c.hooks.OrderItem, hooks...)
}

// Create returns a create builder for OrderItem.
func (c *OrderItemClient) Create() *OrderItemCreate {
	mutation := newOrderItemMutation(c.config, OpCreate)
	return &OrderItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of OrderItem entities.
func (c *OrderItemClient) CreateBulk(builders ...*OrderItemCreate) *OrderItemCreateBulk {
	return &OrderItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OrderItem.
func (c *OrderItemClient) Update() *OrderItemUpdate {
	mutation := newOrderItemMutation(c.config, OpUpdate)
	return &OrderItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrderItemClient) UpdateOne(oi *OrderItem) *OrderItemUpdateOne {
	mutation := newOrderItemMutation(c.config, OpUpdateOne, withOrderItem(oi))
	return &OrderItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrderItemClient) UpdateOneID(id int) *OrderItemUpdateOne {
	mutation := newOrderItemMutation(c.config, OpUpdateOne, withOrderItemID(id))
	return &OrderItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OrderItem.
func (c *OrderItemClient) Delete() *OrderItemDelete {
	mutation := newOrderItemMutation(c.config, OpDelete)
	return &OrderItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OrderItemClient) DeleteOne(oi *OrderItem) *OrderItemDeleteOne {
	return c.DeleteOneID(oi.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OrderItemClient) DeleteOneID(id int) *OrderItemDeleteOne {
	builder := c.Delete().Where(orderitem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrderItemDeleteOne{builder}
}

// Query returns a query builder for OrderItem.
func (c *OrderItemClient) Query() *OrderItemQuery {
	return &OrderItemQuery{config: c.config}
}

// Get returns a OrderItem entity by its id.
func (c *OrderItemClient) Get(ctx context.Context, id int) (*OrderItem, error) {
	return c.Query().Where(orderitem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrderItemClient) GetX(ctx context.Context, id int) *OrderItem {
	oi, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return oi
}

// QueryOrder queries the order edge of a OrderItem.
func (c *OrderItemClient) QueryOrder(oi *OrderItem) *OrderQuery {
	query := &OrderQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := oi.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(orderitem.Table, orderitem.FieldID, id),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderitem.OrderTable, orderitem.OrderColumn),
		)
		fromV = sqlgraph.Neighbors(oi.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OrderItemClient) Hooks() []Hook {
	return c.hooks.OrderItem
}
