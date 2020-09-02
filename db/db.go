package db

import (
	"context"
	"log"

	"github.com/TylerGrey/ent-shop/db/ent"
)

var client *ent.Client

// Init ...
func Init() {
	var err error
	client, err = connectDb()
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	createSchema()
}

func connectDb() (*ent.Client, error) {
	op := ent.Option(ent.Log(func(a ...interface{}) {
		log.Println("entgo", a)
	}))
	op2 := ent.Option(ent.Debug())

	return ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", op, op2)
}

func createSchema() {
	ctx := context.Background()
	// run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

// GetClient ...
func GetClient() *ent.Client {
	return client
}
