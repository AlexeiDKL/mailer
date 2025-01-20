package sqlites

import (
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

const (
	Id          = "id"
	Name        = "name"
	description = "description"
	Email       = "email"
	Phone       = "phone"
	Address     = "address"
	Created_at  = "created_at"
)

type Tables interface {
	CreateTable() error
	Select(id string) (Storage, error)
	Insert(Storage) error
	Delete(Storage) error
	Drop(Storage) error
	Update(Storage) error
}
