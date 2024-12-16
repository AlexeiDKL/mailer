package sqlites

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func Hi() {
	q := "CREATE TABLE IF NOT EXISTS `company` ("
	q += "`id` INTEGER NOT NULL UNIQUE,"
	q += "`name` TEXT NOT NULL UNIQUE);"
	fmt.Println(q)
}
