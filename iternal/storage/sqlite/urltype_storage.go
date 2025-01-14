package sqlites

import (
	"database/sql"

	dklserrors "dkl.dklsa.mailer/iternal/dkls_errors"
	"dkl.dklsa.mailer/iternal/storage"
	_ "github.com/mattn/go-sqlite3"
)

type UrlType struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

type UrlTypeStorage struct {
	Tables
	db *Storage
}

func CreateUrlTypeStorages(db *Storage) *UrlTypeStorage {
	return &UrlTypeStorage{db: db}
}

func CreateUrlTypeTable(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.CreateUrlTypeTable"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS url_type (
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

func (s *UrlTypeStorage) Select(args storage.Pair) (*UrlType, error) {
	op := "storage.sqlite.selectUrlType"
	var result *sql.Rows
	var err error
	switch args.Type {
	case "id":
		result, err = s.selectWithId(args.Value.(int))
	case "type":
		result, err = s.selectWithType(args.Value.(string))
	default:
		return nil, dklserrors.Wrap(op, err)
	}

	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}

	defer result.Close()

	urlType := &UrlType{}
	if result.Next() {
		err = result.Scan(&urlType.ID, &urlType.Type)
		if err != nil {
			return nil, dklserrors.Wrap(op, err)
		}
	}
	return urlType, nil
}

func (s UrlTypeStorage) selectWithId(id int) (*sql.Rows, error) {
	const op = "storage.sqlite.selectWithId"

	return s.db.db.Query("SELECT * FROM url_type WHERE id = ?", id)
}

func (s UrlTypeStorage) selectWithType(url_type string) (*sql.Rows, error) {
	const op = "storage.sqlite.selectWithType"

	return s.db.db.Query("SELECT * FROM url_type WHERE type = ?", url_type)
}

func (s UrlTypeStorage) SelectAll() ([]UrlType, error) {
	const op = "storage.sqlite.selectAll"

	rows, err := s.db.db.Query("SELECT * FROM url_type")
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}

	defer rows.Close()

	var urls []UrlType
	for rows.Next() {
		var url UrlType
		err := rows.Scan(&url.ID, &url.Type)
		if err != nil {
			return nil, dklserrors.Wrap(op, err)
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func (s UrlTypeStorage) Insert(urls_type *UrlType) (int64, error) {
	const op = "storage.sqlite.insertUrlType"

	stmt, err := s.db.db.Prepare("INSERT INTO url_type (id, type) VALUES (?,?)")
	if err != nil {
		return 0, dklserrors.Wrap(op, err)
	}

	res, err := stmt.Exec(urls_type.ID, urls_type.Type)
	if err != nil {
		return 0, dklserrors.Wrap(op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, dklserrors.Wrap(op, err)
	}

	return id, nil
}

func (s UrlTypeStorage) Update(url_type *UrlType) error {
	const op = "storage.sqlite.updateUrlType"

	stmt, err := s.db.db.Prepare("UPDATE url_type SET type =? WHERE id =?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	_, err = stmt.Exec(url_type.Type, url_type.ID)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	return nil
}

func (s UrlTypeStorage) Delete(args storage.Pair) error {
	const op = "storage.sqlite.DeleteUrlType"

	switch args.Type {
	case "id":
		return s.deleteById(args.Value.(int))
	case "type":
		return s.deleteByType(args.Value.(string))
	default:
		return dklserrors.UnsupportedType(args.Type)
	}
}

func (s UrlTypeStorage) deleteByType(url_type string) error {
	const op = "storage.sqlite.deleteByType"

	stmt, err := s.db.db.Prepare("DELETE FROM url_type WHERE type =?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	_, err = stmt.Exec(url_type)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	return nil
}

func (s UrlTypeStorage) deleteById(id int) error {
	const op = "storage.sqlite.deleteById"

	stmt, err := s.db.db.Prepare("DELETE FROM url_type WHERE id =?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	return nil
}

func (s UrlTypeStorage) Drop() error {
	const op = "storage.sqlite.dropUrlType"

	_, err := s.db.db.Exec("DROP TABLE IF EXISTS url_type")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}
