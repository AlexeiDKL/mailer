package sqlites

import (
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

const (
	id          = "id"
	name        = "name"
	description = "description"
	email       = "email"
	phone       = "phone"
	address     = "address"
	created_at  = "created_at"
)

type Tables interface {
	New() error
	Get(id string) (Storage, error)
	Save(Storage) error
	Delete(Storage) error
	Update(Storage) error
}
