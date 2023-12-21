package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Smbrer1/go-short/internal/storage"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: open db connection: %w", op, err)
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS url(
			id INTEGER PRIMARY KEY,
			url TEXT NON NULL);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(urlToSave string) (int64, error) {
	const op = "storage.sqlite.SaveURL"

	stmt, err := s.db.Prepare("INSERT INTO url(url) VALUES (?)")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	res, err := stmt.Exec(urlToSave)
	if err != nil {
		return 0, fmt.Errorf("%s: execute statement: %w", op, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last inserted id: %w", op, err)
	}
	return id, nil
}

func (s *Storage) GetURL(id int) (string, error) {
	const op = "storage.sqlite.GetURL"

	stmt, err := s.db.Prepare(`SELECT url FROM url WHERE id = ?`)
	if err != nil {
		return "", fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	var res string

	err = stmt.QueryRow(id).Scan(&res)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", storage.ErrURLNotFound
		}
		return "", fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return res, nil
}

// TODO Make Delete Method
