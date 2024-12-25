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

func CreateTable(storagePath string) (*Storage, error) {
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

func (s CompanyStorages) Select(args Pair) (*sql.Rows, error) {
	switch args.Type {
	case "id":
		return s.SelectWithId(args.Value)
	case "names":
		return s.SelectWithNames(args.Value)
	default:
		return nil, fmt.Errorf("unsupported type: %s", args.Type)
	}
}

func (s CompanyStorages) SelectWithNames(companyNames string) (*sql.Rows, error) {
	const op = "storage.sqlite.SelectWithNames"

	return s.db.db.Query("SELECT id, name FROM company WHERE name =?", companyNames)
}

func (s CompanyStorages) SelectWithId(id string) (*sql.Rows, error) {
	const op = "storage.sqlite.SelectWithId"

	return s.db.db.Query("SELECT id, name FROM company WHERE id =?", id)
}

func (s CompanyStorages) Insert(companyName string) (int64, error) {
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

func (s CompanyStorages) Delete(args Pair) error {
	const os = "storage.sqlite.Delete"
	switch args.Type {
	case id:
		return s.deleteById(args.Value)
	case name:
		return s.deleteByNames(args.Value)
	default:
		return fmt.Errorf("unsupported type: %s", args.Type)
	}
}

func (s CompanyStorages) deleteByNames(companyNames string) error {
	const op = "storage.sqlite.DeleteByNames"

	stmt, err := s.db.db.Prepare("DELETE FROM company WHERE name =?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	res, err := stmt.Exec(companyNames)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to get rows affected: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: no company found with name %s", op, companyNames)
	}

	return nil
}

func (s CompanyStorages) deleteById(id string) error {
	const op = "storage.sqlite.DeleteById"

	stmt, err := s.db.db.Prepare("DELETE FROM company WHERE id =?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to get rows affected: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: no company found with id %s", op, id)
	}

	return nil
}

func (s CompanyStorages) Drop() error {
	const op = "storage.sqlite.Drop"

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

func (s CompanyStorages) Update(arg Pair) error {
	const op = "storage.sqlite.UpdateCompany"
	switch arg.Type {
	case "id":
		return s.updateById(arg.Type, arg.Value)
	case "name":
		return s.updateByName(arg.Type, arg.Value)
	default:
		return fmt.Errorf("unsupported type: %s", arg.Type)
	}
}

func (s CompanyStorages) updateById(id, newName string) error {
	const op = "storage.sqlite.UpdateById"

	stmt, err := s.db.db.Prepare("UPDATE company SET name =? WHERE id =?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	res, err := stmt.Exec(newName, id)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to get rows affected: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: no company found with id %s", op, id)
	}

	return nil
}

func (s CompanyStorages) updateByName(name, newName string) error {
	const op = "storage.sqlite.UpdateByName"

	stmt, err := s.db.db.Prepare("UPDATE company SET name =? WHERE name =?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	res, err := stmt.Exec(newName, name)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to get rows affected: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: no company found with name %s", op, name)
	}

	return nil
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
