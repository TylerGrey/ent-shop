// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/TylerGrey/ent-shop/db/ent/category"
	"github.com/TylerGrey/ent-shop/db/ent/item"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ItemCreate is the builder for creating a Item entity.
type ItemCreate struct {
	config
	mutation *ItemMutation
	hooks    []Hook
}

// SetCreatedAt sets the created_at field.
func (ic *ItemCreate) SetCreatedAt(t time.Time) *ItemCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (ic *ItemCreate) SetNillableCreatedAt(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the updated_at field.
func (ic *ItemCreate) SetUpdatedAt(t time.Time) *ItemCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (ic *ItemCreate) SetNillableUpdatedAt(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetName sets the name field.
func (ic *ItemCreate) SetName(s string) *ItemCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetPrice sets the price field.
func (ic *ItemCreate) SetPrice(i int32) *ItemCreate {
	ic.mutation.SetPrice(i)
	return ic
}

// SetStockQuantity sets the stockQuantity field.
func (ic *ItemCreate) SetStockQuantity(i int32) *ItemCreate {
	ic.mutation.SetStockQuantity(i)
	return ic
}

// SetDtype sets the dtype field.
func (ic *ItemCreate) SetDtype(i item.Dtype) *ItemCreate {
	ic.mutation.SetDtype(i)
	return ic
}

// AddCategoryIDs adds the categories edge to Category by ids.
func (ic *ItemCreate) AddCategoryIDs(ids ...int) *ItemCreate {
	ic.mutation.AddCategoryIDs(ids...)
	return ic
}

// AddCategories adds the categories edges to Category.
func (ic *ItemCreate) AddCategories(c ...*Category) *ItemCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ic.AddCategoryIDs(ids...)
}

// Mutation returns the ItemMutation object of the builder.
func (ic *ItemCreate) Mutation() *ItemMutation {
	return ic.mutation
}

// Save creates the Item in the database.
func (ic *ItemCreate) Save(ctx context.Context) (*Item, error) {
	if err := ic.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Item
	)
	if len(ic.hooks) == 0 {
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ic.mutation = mutation
			node, err = ic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ItemCreate) SaveX(ctx context.Context) *Item {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ic *ItemCreate) preSave() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := item.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := item.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := ic.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	if _, ok := ic.mutation.StockQuantity(); !ok {
		return &ValidationError{Name: "stockQuantity", err: errors.New("ent: missing required field \"stockQuantity\"")}
	}
	if _, ok := ic.mutation.Dtype(); !ok {
		return &ValidationError{Name: "dtype", err: errors.New("ent: missing required field \"dtype\"")}
	}
	if v, ok := ic.mutation.Dtype(); ok {
		if err := item.DtypeValidator(v); err != nil {
			return &ValidationError{Name: "dtype", err: fmt.Errorf("ent: validator failed for field \"dtype\": %w", err)}
		}
	}
	return nil
}

func (ic *ItemCreate) sqlSave(ctx context.Context) (*Item, error) {
	i, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	i.ID = int(id)
	return i, nil
}

func (ic *ItemCreate) createSpec() (*Item, *sqlgraph.CreateSpec) {
	var (
		i     = &Item{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: item.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: item.FieldID,
			},
		}
	)
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldCreatedAt,
		})
		i.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldUpdatedAt,
		})
		i.UpdatedAt = value
	}
	if value, ok := ic.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldName,
		})
		i.Name = value
	}
	if value, ok := ic.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: item.FieldPrice,
		})
		i.Price = value
	}
	if value, ok := ic.mutation.StockQuantity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: item.FieldStockQuantity,
		})
		i.StockQuantity = value
	}
	if value, ok := ic.mutation.Dtype(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: item.FieldDtype,
		})
		i.Dtype = value
	}
	if nodes := ic.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.CategoriesTable,
			Columns: item.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return i, _spec
}

// ItemCreateBulk is the builder for creating a bulk of Item entities.
type ItemCreateBulk struct {
	config
	builders []*ItemCreate
}

// Save creates the Item entities in the database.
func (icb *ItemCreateBulk) Save(ctx context.Context) ([]*Item, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Item, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*ItemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (icb *ItemCreateBulk) SaveX(ctx context.Context) []*Item {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}