package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//TODO: learn more about all this vv
type Datastore interface {
	GetShopper(id string) *Shopper
}

type DB struct {
	*sql.DB
}

//Env holds environment variables
type Env struct {
	db Datastore
}

var env *Env

func NewDB(dataSourceName string) error {
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}

	env = &Env{
		&DB{db},
	}

	return nil
}
