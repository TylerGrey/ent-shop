package models

import (
	"context"

	"github.com/TylerGrey/ent-shop/db"
	"github.com/TylerGrey/ent-shop/db/ent"
	"github.com/TylerGrey/ent-shop/forms"
)

// MemberModel ...
type MemberModel struct{}

// Join ...
func (mm MemberModel) Join(form forms.JoinForm) (*int, error) {
	m, err := db.GetClient().Member.
		Create().
		SetName(form.Name).
		SetCity(form.City).
		SetStreet(form.Street).
		SetZipcode(form.Zipcode).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return &m.ID, nil
}

// FindMembers ...
func (mm MemberModel) FindMembers() ([]*ent.Member, error) {
	return db.GetClient().Member.
		Query().
		All(context.Background())
}
