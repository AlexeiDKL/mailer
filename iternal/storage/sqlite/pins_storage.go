package sqlites

import (
	"database/sql"

	dklserrors "dkl.dklsa.mailer/iternal/dkls_errors"
	"dkl.dklsa.mailer/iternal/storage"
)

type Pins struct {
	ID   int    `json:"id"`
	Pin  string `json:"pin"`
	User int    `json:"user"`
}

type PinsStorage struct {
	Tables
	db *Storage
}

func CreatePinsStorages(db *Storage) *PinsStorage {
	return &PinsStorage{db: db}
}

func CreatePinsTable(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.CreatePinsTable"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`
    CREATE TABLE IF NOT EXISTS pins (
        id integer primary key NOT NULL UNIQUE,
        pin TEXT NOT NULL UNIQUE,
        user INTEGER NOT NULL
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

func (s *PinsStorage) Select(args storage.Pair) (*Pins, error) {
	op := "storage.sqlite.selectPins"
	var result *sql.Rows
	var err error
	switch args.Type {
	case "id":
		result, err = s.db.db.Query("SELECT * FROM pins WHERE id=?", args.Value)
	case "user":
		result, err = s.db.db.Query("SELECT * FROM pins WHERE user=?", args.Value)
	default:
		return nil, dklserrors.UnsupportedType(args.Type)
	}

	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	defer result.Close()

	pin := &Pins{}
	for result.Next() {
		err = result.Scan(&pin.ID, &pin.Pin, &pin.User)
		if err != nil {
			return nil, dklserrors.Wrap(op, err)
		}
	}
	return pin, nil
}

func (s *PinsStorage) SelectAll() ([]Pins, error) {
	op := "storage.sqlite.selectAllPins"
	rows, err := s.db.db.Query("SELECT * FROM pins")
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	defer rows.Close()

	pins := []Pins{}
	for rows.Next() {
		pin := Pins{}
		err := rows.Scan(&pin.ID, &pin.Pin, &pin.User)
		if err != nil {
			return nil, dklserrors.Wrap(op, err)
		}
		pins = append(pins, pin)
	}
	return pins, nil
}

func (s *PinsStorage) Insert(pin *Pins) (int64, error) {
	op := "storage.sqlite.insertPins"
	stmt, err := s.db.db.Prepare("INSERT INTO pins (id, pin, user) VALUES (?,?,?)")
	if err != nil {
		return 0, dklserrors.Wrap(op, err)
	}
	res, err := stmt.Exec(pin.ID, pin.Pin, pin.User)
	if err != nil {
		return 0, dklserrors.Wrap(op, err)
	}
	id, err := res.LastInsertId()
	return id, nil
}

func (s *PinsStorage) Update(pin *Pins) error {
	op := "storage.sqlite.updatePins"
	stmt, err := s.db.db.Prepare("UPDATE pins SET pin=?, user=? WHERE id=?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	_, err = stmt.Exec(pin.Pin, pin.User, pin.ID)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}

func (s *PinsStorage) Delete(args storage.Pair) error {
	op := "storage.sqlite.deletePins"
	switch args.Type {
	case "id":
		stmt, err := s.db.db.Prepare("DELETE FROM pins WHERE id=?")
		if err != nil {
			return dklserrors.Wrap(op, err)
		}
		_, err = stmt.Exec(args.Value)
		if err != nil {
			return dklserrors.Wrap(op, err)
		}
	default:
		return dklserrors.UnsupportedType(args.Type)
	}
	return nil
}

func (s *PinsStorage) Drop() error {
	op := "storage.sqlite.dropPins"
	_, err := s.db.db.Exec("DROP TABLE IF EXISTS pins")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}
