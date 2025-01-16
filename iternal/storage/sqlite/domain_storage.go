package sqlites

import (
	"database/sql"

	dklserrors "dkl.dklsa.mailer/iternal/dkls_errors"
	"dkl.dklsa.mailer/iternal/storage"
	_ "github.com/mattn/go-sqlite3"
)

type Domens struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CompanyID int    `json:"company_id"`
}

type DomensStorage struct {
	Tables
	db *Storage
}

func CreateDomensStorages(db *Storage) *DomensStorage {
	return &DomensStorage{db: db}
}

func CreateDomensTable(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.CreatedomensTable"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	defer db.Close()
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS domens (
			id integer primary key NOT NULL UNIQUE,
			name TEXT UNIQUE,
			company_id INTEGER NOT NULL,
		FOREIGN KEY(company_id) REFERENCES company(id)
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

func (s DomensStorage) Select(args storage.Pair) (*Domens, error) {
	const op = "storage.sqlite.select"
	var result *sql.Rows
	var err error
	switch args.Type {
	case "id":
		result, err = s.selectWithId(args.Value.(int))
	case "name":
		result, err = s.selectWithName(args.Value.(string))
	case "company_id":
		result, err = s.selectWithCompanyID(args.Value.(int))
	default:
		return nil, dklserrors.UnsupportedType(args.Type)
	}

	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}

	defer result.Close()

	domens := &Domens{}
	if result.Next() {
		err = result.Scan(&domens.ID, &domens.Name, &domens.CompanyID)
		if err != nil {
			return nil, dklserrors.Wrap(op, err)
		}
	}
	return domens, nil
}

func (s DomensStorage) selectWithId(id int) (*sql.Rows, error) {
	const op = "storage.sqlite.selectWithId"

	return s.db.db.Query("SELECT * FROM domens WHERE id =?", id)
}

func (s DomensStorage) selectWithName(name string) (*sql.Rows, error) {
	const op = "storage.sqlite.selectWithName"

	return s.db.db.Query("SELECT * FROM domens WHERE name = ?", name)
}

func (s DomensStorage) selectWithCompanyID(companyID int) (*sql.Rows, error) {
	const op = "storage.sqlite.selectWithCompanyID"

	return s.db.db.Query("SELECT * FROM domens WHERE id = ?", companyID)
}

func (s DomensStorage) SelectAll() ([]Domens, error) {
	const op = "storage.sqlite.selectAll"

	rows, err := s.db.db.Query("SELECT * FROM company")
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	defer rows.Close()

	domenss := []Domens{}
	for rows.Next() {
		domens := Domens{}
		err = rows.Scan(&domens.ID, &domens.Name, &domens.CompanyID)
		if err != nil {
			return nil, dklserrors.Wrap(op, err)
		}
		domenss = append(domenss, domens)
	}
	return domenss, nil
}

func (s DomensStorage) Insert(domens *Domens) (int64, error) {
	const op = "storage.sqlite.Insert"

	stmt, err := s.db.db.Prepare("INSERT INTO domens (id, name, company_id) VALUES (?,?,?)")
	if err != nil {
		return 0, dklserrors.Wrap(op, err)
	}
	res, err := stmt.Exec(domens.ID, domens.Name, domens.CompanyID)
	if err != nil {
		return 0, dklserrors.Wrap(op, err)
	}
	return res.LastInsertId()
}

func (s DomensStorage) Update(domens *Domens) error {
	const op = "storage.sqlite.Update"

	stmt, err := s.db.db.Prepare("UPDATE domens SET name=?, company_id=? WHERE id=?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	_, err = stmt.Exec(domens.Name, domens.CompanyID, domens.ID)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}

func (s DomensStorage) Delete(args storage.Pair) error {
	const op = "storage.sqlite.Delete"

	switch args.Type {
	case "id":
		return s.deleteById(args.Value.(int))
	case "name":
		return s.deleteByName(args.Value.(string))
	default:
		return dklserrors.UnsupportedType(args.Type)
	}
}

func (s DomensStorage) deleteById(id int) error {
	const op = "storage.sqlite.deleteById"

	stmt, err := s.db.db.Prepare("DELETE FROM domens WHERE id=?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}

func (s DomensStorage) deleteByName(name string) error {
	const op = "storage.sqlite.deleteByName"

	stmt, err := s.db.db.Prepare("DELETE FROM domens WHERE name=?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	_, err = stmt.Exec(name)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}

func (s DomensStorage) Drop() error {
	const op = "storage.sqlite.Drop"

	_, err := s.db.db.Exec("DROP TABLE IF EXISTS domens")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}
