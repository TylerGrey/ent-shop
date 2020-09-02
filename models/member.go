package models

import (
	"context"
	"log"

	"github.com/TylerGrey/ent-shop/db"
	"github.com/TylerGrey/ent-shop/db/ent"
	"github.com/TylerGrey/ent-shop/forms"
)

// MemberModel ...
type MemberModel struct{}

// Join ...
func (mm MemberModel) Join(member forms.JoinForm) (*int, error) {
	log.Println("client: ", db.GetClient())
	m, err := db.GetClient().Member.
		Create().
		SetName(member.Name).
		SetCity(member.City).
		SetStreet(member.Street).
		SetZipcode(member.Zipcode).
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

// FindOne ...
func (mm MemberModel) FindOne() {

}
