// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/TylerGrey/ent-shop/db/ent/delivery"
	"github.com/TylerGrey/ent-shop/db/ent/order"
	"github.com/TylerGrey/ent-shop/db/ent/orderitem"
	"github.com/TylerGrey/ent-shop/db/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks      []Hook
	mutation   *OrderMutation
	predicates []predicate.Order
}

// Where adds a new predicate for the builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.predicates = append(ou.predicates, ps...)
	return ou
}

// SetUpdatedAt sets the updated_at field.
func (ou *OrderUpdate) SetUpdatedAt(t time.Time) *OrderUpdate {
	ou.mutation.SetUpdatedAt(t)
	return ou
}

// SetOrderDate sets the orderDate field.
func (ou *OrderUpdate) SetOrderDate(t time.Time) *OrderUpdate {
	ou.mutation.SetOrderDate(t)
	return ou
}

// SetStatus sets the status field.
func (ou *OrderUpdate) SetStatus(o order.Status) *OrderUpdate {
	ou.mutation.SetStatus(o)
	return ou
}

// AddOrderItemIDs adds the orderItems edge to OrderItem by ids.
func (ou *OrderUpdate) AddOrderItemIDs(ids ...int) *OrderUpdate {
	ou.mutation.AddOrderItemIDs(ids...)
	return ou
}

// AddOrderItems adds the orderItems edges to OrderItem.
func (ou *OrderUpdate) AddOrderItems(o ...*OrderItem) *OrderUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.AddOrderItemIDs(ids...)
}

// SetDeliveryID sets the delivery edge to Delivery by id.
func (ou *OrderUpdate) SetDeliveryID(id int) *OrderUpdate {
	ou.mutation.SetDeliveryID(id)
	return ou
}

// SetNillableDeliveryID sets the delivery edge to Delivery by id if the given value is not nil.
func (ou *OrderUpdate) SetNillableDeliveryID(id *int) *OrderUpdate {
	if id != nil {
		ou = ou.SetDeliveryID(*id)
	}
	return ou
}

// SetDelivery sets the delivery edge to Delivery.
func (ou *OrderUpdate) SetDelivery(d *Delivery) *OrderUpdate {
	return ou.SetDeliveryID(d.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// RemoveOrderItemIDs removes the orderItems edge to OrderItem by ids.
func (ou *OrderUpdate) RemoveOrderItemIDs(ids ...int) *OrderUpdate {
	ou.mutation.RemoveOrderItemIDs(ids...)
	return ou
}

// RemoveOrderItems removes orderItems edges to OrderItem.
func (ou *OrderUpdate) RemoveOrderItems(o ...*OrderItem) *OrderUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.RemoveOrderItemIDs(ids...)
}

// ClearDelivery clears the delivery edge to Delivery.
func (ou *OrderUpdate) ClearDelivery() *OrderUpdate {
	ou.mutation.ClearDelivery()
	return ou
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := ou.mutation.UpdatedAt(); !ok {
		v := order.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
	if v, ok := ou.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return 0, &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		},
	}
	if ps := ou.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ou.mutation.OrderDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldOrderDate,
		})
	}
	if value, ok := ou.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: order.FieldStatus,
		})
	}
	if nodes := ou.mutation.RemovedOrderItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderItemsTable,
			Columns: []string{order.OrderItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.OrderItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderItemsTable,
			Columns: []string{order.OrderItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.DeliveryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   order.DeliveryTable,
			Columns: []string{order.DeliveryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: delivery.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.DeliveryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   order.DeliveryTable,
			Columns: []string{order.DeliveryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: delivery.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// SetUpdatedAt sets the updated_at field.
func (ouo *OrderUpdateOne) SetUpdatedAt(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetUpdatedAt(t)
	return ouo
}

// SetOrderDate sets the orderDate field.
func (ouo *OrderUpdateOne) SetOrderDate(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetOrderDate(t)
	return ouo
}

// SetStatus sets the status field.
func (ouo *OrderUpdateOne) SetStatus(o order.Status) *OrderUpdateOne {
	ouo.mutation.SetStatus(o)
	return ouo
}

// AddOrderItemIDs adds the orderItems edge to OrderItem by ids.
func (ouo *OrderUpdateOne) AddOrderItemIDs(ids ...int) *OrderUpdateOne {
	ouo.mutation.AddOrderItemIDs(ids...)
	return ouo
}

// AddOrderItems adds the orderItems edges to OrderItem.
func (ouo *OrderUpdateOne) AddOrderItems(o ...*OrderItem) *OrderUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.AddOrderItemIDs(ids...)
}

// SetDeliveryID sets the delivery edge to Delivery by id.
func (ouo *OrderUpdateOne) SetDeliveryID(id int) *OrderUpdateOne {
	ouo.mutation.SetDeliveryID(id)
	return ouo
}

// SetNillableDeliveryID sets the delivery edge to Delivery by id if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableDeliveryID(id *int) *OrderUpdateOne {
	if id != nil {
		ouo = ouo.SetDeliveryID(*id)
	}
	return ouo
}

// SetDelivery sets the delivery edge to Delivery.
func (ouo *OrderUpdateOne) SetDelivery(d *Delivery) *OrderUpdateOne {
	return ouo.SetDeliveryID(d.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// RemoveOrderItemIDs removes the orderItems edge to OrderItem by ids.
func (ouo *OrderUpdateOne) RemoveOrderItemIDs(ids ...int) *OrderUpdateOne {
	ouo.mutation.RemoveOrderItemIDs(ids...)
	return ouo
}

// RemoveOrderItems removes orderItems edges to OrderItem.
func (ouo *OrderUpdateOne) RemoveOrderItems(o ...*OrderItem) *OrderUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.RemoveOrderItemIDs(ids...)
}

// ClearDelivery clears the delivery edge to Delivery.
func (ouo *OrderUpdateOne) ClearDelivery() *OrderUpdateOne {
	ouo.mutation.ClearDelivery()
	return ouo
}

// Save executes the query and returns the updated entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	if _, ok := ouo.mutation.UpdatedAt(); !ok {
		v := order.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
	if v, ok := ouo.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return nil, &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}

	var (
		err  error
		node *Order
	)
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	o, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return o
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (o *Order, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Order.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ouo.mutation.OrderDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldOrderDate,
		})
	}
	if value, ok := ouo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: order.FieldStatus,
		})
	}
	if nodes := ouo.mutation.RemovedOrderItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderItemsTable,
			Columns: []string{order.OrderItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.OrderItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderItemsTable,
			Columns: []string{order.OrderItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.DeliveryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   order.DeliveryTable,
			Columns: []string{order.DeliveryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: delivery.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.DeliveryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   order.DeliveryTable,
			Columns: []string{order.DeliveryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: delivery.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	o = &Order{config: ouo.config}
	_spec.Assign = o.assignValues
	_spec.ScanValues = o.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return o, nil
}
