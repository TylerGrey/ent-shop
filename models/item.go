package models

import (
	"context"

	"github.com/TylerGrey/ent-shop/db"
	"github.com/TylerGrey/ent-shop/db/ent"
	"github.com/TylerGrey/ent-shop/db/ent/item"
	"github.com/TylerGrey/ent-shop/forms"
)

// ItemModel ...
type ItemModel struct{}

// SaveItem ...
func (im ItemModel) SaveItem(form forms.BookForm) (*int, error) {
	m, err := db.GetClient().Item.
		Create().
		SetName(form.Name).
		SetPrice(form.Price).
		SetStockQuantity(form.StockQuantity).
		SetDtype(item.DtypeBOOK).
		SetAuthor(form.Author).
		SetIsbn(form.Isbn).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return &m.ID, nil
}

// UpdateItem ...
func (im ItemModel) UpdateItem(form forms.BookForm) (*int, error) {
	m, err := db.GetClient().Item.
		UpdateOneID(form.ID).
		SetName(form.Name).
		SetPrice(form.Price).
		SetStockQuantity(form.StockQuantity).
		SetDtype(item.DtypeBOOK).
		SetAuthor(form.Author).
		SetIsbn(form.Isbn).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return &m.ID, nil
}

// FindItems ...
func (im ItemModel) FindItems() ([]*ent.Item, error) {
	return db.GetClient().Item.
		Query().
		All(context.Background())
}

// FindOne ...
func (im ItemModel) FindOne(ID int) (*ent.Item, error) {
	return db.GetClient().Item.
		Get(context.Background(), ID)
}
