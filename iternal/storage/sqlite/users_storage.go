package sqlites

import (
	"database/sql"

	dklserrors "dkl.dklsa.mailer/iternal/dkls_errors"
	"dkl.dklsa.mailer/iternal/storage"
)

type Users struct {
	ID    int    `json:"id"`
	Mail  string `json:"mail"`
	Domen int    `json:"domen"`
}

type UsersStorages struct {
	Tables
	db *Storage
}

func CreateUsersStorages(db *Storage) *UsersStorages {
	return &UsersStorages{db: db}
}

func CreateUsersTable(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.CreateUsersTable"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Users (
			id integer primary key NOT NULL UNIQUE,
			mail TEXT NOT NULL,
			domen INTEGER NOT NULL,
			FOREIGN KEY(domen) REFERENCES domens(id)
		);
	`)
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}

	return &Storage{db: db}, nil
}

func (s UsersStorages) Select(args storage.Pair) (*Users, error) {
	const op = "storage.sqlite.selectUsers"
	var result *sql.Rows
	var err error
	switch args.Type {
	case "id":
		result, err = s.db.db.Query("SELECT * FROM Users WHERE id = ?", args.Value)
	case "mail":
		result, err = s.db.db.Query("SELECT * FROM Users WHERE mail = ?", args.Value)
	default:
		return nil, dklserrors.UnsupportedType(args.Type)
	}
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	defer result.Close()

	var user Users
	if result.Next() {
		err = result.Scan(&user.ID, &user.Mail, &user.Domen)
		if err != nil {
			return nil, dklserrors.Wrap(op, err)
		}
	}
	return &user, nil
}

func (s *UsersStorages) SelectAll() ([]Users, error) {
	const op = "storage.sqlite.selectAllUsers"

	rows, err := s.db.db.Query("SELECT * FROM Users")
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	defer rows.Close()

	users := []Users{}
	for rows.Next() {
		var user Users
		err = rows.Scan(&user.ID, &user.Mail, &user.Domen)
		if err != nil {
			return nil, dklserrors.Wrap(op, err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *UsersStorages) Insert(user *Users) error {
	const op = "storage.sqlite.insertUsers"

	stmt, err := s.db.db.Prepare("INSERT INTO Users (id, mail, domen) VALUES (?,?,?)")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Mail, user.Domen)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}

func (s *UsersStorages) Update(user *Users) error {
	const op = "storage.sqlite.updateUsers"

	stmt, err := s.db.db.Prepare("UPDATE Users SET mail=?, domen=? WHERE id=?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Mail, user.Domen, user.ID)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}

func (s *UsersStorages) Delete(args storage.Pair) error {
	const op = "storage.sqlite.deleteUsers"

	switch args.Type {
	case "id":
		return s.deleteById(args.Value.(int))
	case "mail":
		return s.deleteByMail(args.Value.(string))
	}

	return nil
}

func (s *UsersStorages) deleteById(id int) error {
	const op = "storage.sqlite.deleteByIdUsers"

	stmt, err := s.db.db.Prepare("DELETE FROM Users WHERE id=?")
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

func (s *UsersStorages) deleteByMail(mail string) error {
	const op = "storage.sqlite.deleteByMailUsers"

	stmt, err := s.db.db.Prepare("DELETE FROM Users WHERE mail=?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(mail)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}

func (s UsersStorages) Drop() error {
	const op = "storage.sqlite.Drop"

	_, err := s.db.db.Exec("DROP TABLE IF EXISTS Users")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}
