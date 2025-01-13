package sqlites

import (
	"database/sql"
	"fmt"

	dklserrors "dkl.dklsa.mailer/iternal/dkls_errors"
	"dkl.dklsa.mailer/iternal/storage"
	_ "github.com/mattn/go-sqlite3"
)

type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CompanyStorages struct {
	Tables
	db *Storage
}

func CreateCompanyStorages(db *Storage) *CompanyStorages {
	return &CompanyStorages{db: db}
}

func CreateCompanyTable(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.CreateCompanyTable"

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

func (s CompanyStorages) Select(args storage.Pair) (*Company, error) {
	op := "storage.sqlite.select"
	var result *sql.Rows
	var err error
	switch args.Type {
	case "id":
		result, err = s.selectWithId(args.Value.(string))
	case "names":
		result, err = s.selectWithNames(args.Value.(string))
	default:
		return nil, dklserrors.UnsupportedType(args.Type)
	}
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}

	c := &Company{}

	result.Next()

	err = result.Scan(&c.ID, &c.Name)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("%s: no company found with id: %s", op, id)
	}
	return c, err
}

func (s CompanyStorages) SelectAll() ([]Company, error) {
	return nil, dklserrors.NotRelizedError()
}

func (s CompanyStorages) selectWithNames(companyNames string) (*sql.Rows, error) {
	const op = "storage.sqlite.SelectWithNames"

	return s.db.db.Query("SELECT id, name FROM company WHERE name =?", companyNames)
}

func (s CompanyStorages) selectWithId(id string) (*sql.Rows, error) {
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

func (s CompanyStorages) Delete(args storage.Pair) error {
	const os = "storage.sqlite.Delete"
	switch args.Type {
	case id:
		return s.deleteById(args.Value.(string))
	case name:
		return s.deleteByNames(args.Value.(string))
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

func (s CompanyStorages) Update(arg storage.Pair) error {
	const op = "storage.sqlite.UpdateCompany"
	switch arg.Type {
	case "id":
		return s.updateById(arg.Type, arg.Value.(string))
	case "name":
		return s.updateByName(arg.Type, arg.Value.(string))
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
