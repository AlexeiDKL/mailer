package sqlites

import (
	"database/sql"

	dklserrors "dkl.dklsa.mailer/iternal/dkls_errors"
	"dkl.dklsa.mailer/iternal/storage"
)

type Mails struct {
	ID      int    `json:"id"`
	User    int    `json:"user"`
	Body    string `json:"body"`
	Sending bool   `json:"sending"`
}

type MailsStorages struct {
	Tables
	db *Storage
}

func CreateMailsStorages(db *Storage) *MailsStorages {
	return &MailsStorages{db: db}
}

func CreateMailTable(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.CreateMailTable"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS mails (
			id integer primary key NOT NULL UNIQUE,
			user INTEGER NOT NULL,
			body TEXT NOT NULL,
	sending REAL NOT NULL DEFAULT '0',
		FOREIGN KEY(user) REFERENCES Users(id)
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

func (s *MailsStorages) Select(args storage.Pair) (*Mails, error) {
	const op = "storage.sqlite.select"
	var result *sql.Rows
	var err error
	switch args.Type {
	case "id":
		result, err = s.selectWithId(args.Value.(int))
	case "user":
		result, err = s.selectWithUser(args.Value.(int))
	}
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	defer result.Close()
	mails := &Mails{}
	if result.Next() {
		err = result.Scan(&mails.ID, &mails.User, &mails.Body, &mails.Sending)
		if err != nil {
			return nil, dklserrors.Wrap(op, err)
		}
	}
	return mails, nil
}

func (s *MailsStorages) selectWithId(id int) (*sql.Rows, error) {
	const op = "storage.sqlite.SelectWithId"
	return s.db.db.Query("SELECT * FROM mails WHERE id =?", id)
}

func (s *MailsStorages) selectWithUser(user_id int) (*sql.Rows, error) {
	const op = "storage.sqlite.SelectWithUser"
	return s.db.db.Query("SELECT * FROM mails WHERE user =?", user_id)
}

func (s *MailsStorages) SelectAll() ([]Mails, error) {
	const op = "storage.sqlite.SelectAll"
	rows, err := s.db.db.Query("SELECT * FROM mails")
	if err != nil {
		return nil, dklserrors.Wrap(op, err)
	}
	defer rows.Close()
	mails := []Mails{}
	for rows.Next() {
		mail := &Mails{}
		err := rows.Scan(&mail.ID, &mail.User, &mail.Body, &mail.Sending)
		if err != nil {
			return nil, dklserrors.Wrap(op, err)
		}
		mails = append(mails, *mail)
	}
	return mails, nil
}

func (s *MailsStorages) Insert(mail *Mails) error {
	const op = "storage.sqlite.Insert"
	stmt, err := s.db.db.Prepare("INSERT INTO mails (id, user, body, sending) VALUES (?,?,?,?)")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(mail.ID, mail.User, mail.Body, mail.Sending)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}

func (s *MailsStorages) Update(mail *Mails) error {
	const op = "storage.sqlite.Update"
	stmt, err := s.db.db.Prepare("UPDATE mails SET user=?, body=?, sending=? WHERE id=?")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(mail.User, mail.Body, mail.Sending, mail.ID)
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}

func (s *MailsStorages) Delete(id int) error {
	const op = "storage.sqlite.Delete"
	stmt, err := s.db.db.Prepare("DELETE FROM mails WHERE id=?")
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

func (s *MailsStorages) Drop() error {
	const op = "storage.sqlite.Drop"

	_, err := s.db.db.Exec("DROP TABLE IF EXISTS mails")
	if err != nil {
		return dklserrors.Wrap(op, err)
	}
	return nil
}
