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
	CreateTable() error
	Select(id string) (Storage, error)
	Insert(Storage) error
	Delete(Storage) error
	Drop(Storage) error
	Update(Storage) error
}
