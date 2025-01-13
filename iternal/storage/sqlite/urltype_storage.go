package sqlites

import (
	"database/sql"

	dklserrors "dkl.dklsa.mailer/iternal/dkls_errors"
	"dkl.dklsa.mailer/iternal/storage"
)

type UrlTypeStorage struct {
	Tables
	db *Storage
}

func NewTablesUrlType(db *Storage) *UrlTypeStorage {
	return &UrlTypeStorage{db: db}
}

func CreateUrlTypeTable(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.createUrlTypeTable"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS url_table (
	    id integer primary key NOT NULL UNIQUE,
        type TEXT NOT NULL UNIQUE
		);
	`)
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	return &Storage{db: db}, nil
}

func (s UrlTypeStorage) Select(args storage.Pair) (*sql.Rows, error) {
	switch args.Type {
	case "id":
		return s.selectWithId(args.Value.(string))
	case "type":
		return s.selectWithType(args.Value.(string))
	default:
		return nil, dklserrors.UnsupportedType(args.Type)
	}
}

func (s UrlTypeStorage) selectWithId(id string) (*sql.Rows, error) {
	const op = "storage.sqlite.SelectWithId"

	stmt, err := s.db.db.Prepare("SELECT * FROM url_table WHERE id =?")
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	return rows, nil
}

func (s UrlTypeStorage) selectWithType(urlType string) (*sql.Rows, error) {
	const op = "storage.sqlite.SelectWithType"

	stmt, err := s.db.db.Prepare("SELECT * FROM url_table WHERE type =?")
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	rows, err := stmt.Query(urlType)
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	return rows, nil
}
