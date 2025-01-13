package sqlites

import (
	"database/sql"
	"fmt"

	dklserrors "dkl.dklsa.mailer/iternal/dkls_errors"
	"dkl.dklsa.mailer/iternal/storage"
)

type Url struct {
	ID       int    `json:"id"`
	Company  string `json:"company"`
	Url      string `json:"url"`
	Url_type string `json:"url_type"`
}

type UrlStorages struct {
	Tables
	db *Storage
}

func CreateUrlStorages(db *Storage) *UrlStorages {
	return &UrlStorages{db: db}
}

func CreateUrlTable(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.CreateUrlTable"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS url_table (
        id integer primary key NOT NULL UNIQUE,
        company TEXT NOT NULL,
        url TEXT NOT NULL,
        url_type TEXT NOT NULL,
        FOREIGN KEY(company) REFERENCES company(id),
        FOREIGN KEY(url_type) REFERENCES url_type(id)
    );`)
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	return &Storage{db: db}, nil
}

func (s UrlStorages) Select(args storage.Pair) (*Url, error) {
	op := "storage.sqlite.select"
	var result *sql.Rows
	var err error
	switch args.Type {
	case "id":
		result, err = s.selectWithId(args.Value.(int))
	case "url":
		result, err = s.selectWithUrl(args.Value.(string))
	default:
		return nil, dklserrors.UnsupportedType(args.Type)
	}
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	defer result.Close()

	url := &Url{}
	for result.Next() {
		err := result.Scan(&url.ID, &url.Company, &url.Url, &url.Url_type)
		if err != nil {
			return nil, dklserrors.Wrap(op, err)
		}
	}
	if url.ID == 0 {
		return nil, fmt.Errorf("%s: no url found with %s: %s", op, args.Type, args.Value)
	}
	return url, nil
}

func (s UrlStorages) selectWithId(id int) (*sql.Rows, error) {
	const op = "storage.sqlite.selectWithId"

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

func (s UrlStorages) selectWithUrl(url string) (*sql.Rows, error) {
	const op = "storage.sqlite.selectWithUrl"

	stmt, err := s.db.db.Prepare("SELECT * FROM url_table WHERE url =?")
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	rows, err := stmt.Query(url)
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	return rows, nil
}

func (s UrlStorages) SelectAll() ([]Url, error) {
	const op = "storage.sqlite.SelectAll"

	rows, err := s.db.db.Query("SELECT * FROM url_table")
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	defer rows.Close()

	urls := []Url{}
	for rows.Next() {
		url := Url{}
		err := rows.Scan(&url.ID, &url.Company, &url.Url, &url.Url_type)
		if err != nil {
			return nil, dklserrors.Wrap(op, err)
		}
		urls = append(urls, url)
	}
	return urls, nil
}

func (s UrlStorages) Insert(url *Url) (int64, error) {
	const op = "storage.sqlite.InsertUrl"

	stmt, err := s.db.db.Prepare("INSERT INTO url_table (company, url, url_type) VALUES (?,?,?)")
	if err != nil {
		return 0, dklserrors.Wrap(op, err)
	}
	result, err := stmt.Exec(url.Company, url.Url, url.Url_type)
	if err != nil {
		return 0, dklserrors.Wrap(op, err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, dklserrors.Wrap(op, err)
	}
	return id, nil
}

func (s UrlStorages) Update(url *Url) error {
	const op = "storage.sqlite.UpdateUrl"

	stmt, err := s.db.db.Prepare("UPDATE url_table SET company =?, url =?, url_type =? WHERE id =?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	_, err = stmt.Exec(url.Company, url.Url, url.Url_type, url.ID)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}

func (s UrlStorages) Delete(args storage.Pair) error {
	const op = "storage.sqlite.DeleteUrl"

	switch args.Type {
	case "id":
		return s.deleteWithId(args.Value.(int))
	case "url":
		return s.deleteWithUrl(args.Value.(string))
	default:
		return dklserrors.UnsupportedType(args.Type)
	}
}

func (s UrlStorages) deleteWithId(id int) error {
	const op = "storage.sqlite.deleteWithId"

	stmt, err := s.db.db.Prepare("DELETE FROM url_table WHERE id =?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}

func (s UrlStorages) deleteWithUrl(url string) error {
	const op = "storage.sqlite.deleteWithUrl"

	stmt, err := s.db.db.Prepare("DELETE FROM url_table WHERE url =?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	_, err = stmt.Exec(url)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}

func (s UrlStorages) Drop() error {
	const op = "storage.sqlite.DropUrlTable"

	_, err := s.db.db.Exec("DROP TABLE IF EXISTS url_table")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}
