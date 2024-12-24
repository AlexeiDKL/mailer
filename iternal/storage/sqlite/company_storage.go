package sqlites

import (
	"database/sql"
	"fmt"

	dklserrors "dkl.dklsa.mailer/iternal/dkls_errors"
	_ "github.com/mattn/go-sqlite3"
)

type Pair struct {
	Type  string
	Value string
}

type CompanyStorages struct {
	Tables
	db *Storage
}

func CreateCompanyStorages(db *Storage) *CompanyStorages {
	return &CompanyStorages{db: db}
}

func NewTablesCompany(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS company (
			id integer primary key NOT NULL UNIQUE,
			name TEXT NOT NULL UNIQUE
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

func (s CompanyStorages) New() error {

	return dklserrors.NotRelizedError()
}

func (s CompanyStorages) Get(args Pair) (*sql.Rows, error) {
	switch args.Type {
	case "id":
		return s.getWithId(args.Value)
	case "names":
		return s.getWithNames(args.Value)
	default:
		return nil, fmt.Errorf("unsupported type: %s", args.Type)
	}
}

func (s CompanyStorages) getWithNames(companyNames string) (*sql.Rows, error) {
	const op = "storage.sqlite.getWithNames"

	return s.db.db.Query("SELECT id, name FROM company WHERE name =?", companyNames)
}

func (s CompanyStorages) getWithId(id string) (*sql.Rows, error) {
	const op = "storage.sqlite.getWithId"

	return s.db.db.Query("SELECT id, name FROM company WHERE id =?", id)
}

func (s CompanyStorages) Save(companyName string) (int64, error) {
	const op = "storage.sqlite.SaveCompany"

	stmt, err := s.db.db.Prepare("INSERT INTO company(name) VALUES(?)")
	if err != nil {
		return -1, dklserrors.Wrap(op, err)
	}

	res, err := stmt.Exec(companyName)
	if err != nil {
		return -2, dklserrors.Wrap(op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -3, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}

	return id, nil
}

func (s CompanyStorages) Delete(st Storage) error {
	return dklserrors.NotRelizedError()
}

func (s CompanyStorages) Update(st Storage) error {
	return dklserrors.NotRelizedError()
}

func (s CompanyStorages) DropTable() error {
	const op = "storage.sqlite.DropTableCompany"

	stmt, err := s.db.db.Prepare("DROP TABLE if EXISTS company;")

	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	return nil
}
