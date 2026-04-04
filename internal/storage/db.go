package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	DB *sql.DB
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages (
					  id INTEGER PRIMARY KEY AUTOINCREMENT,
					  content TEXT
	)`)

	if err != nil {
		return nil, err
	}
	return &Storage{DB: db}, nil
}

func (s *Storage) SaveMessage(content string) error {
	_, err := s.DB.Exec("INSERT INTO messages(content) VALUES(?)", content)
	return err
}

func (s *Storage) GetMessages() ([]string, error) {
	rows, err := s.DB.Query(`SELECT content FROM messages ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []string
	for rows.Next() {
		var content string
		if err := rows.Scan(&content); err != nil {
			return nil, err
		}
		msgs = append(msgs, content)
	}
	return msgs, nil
}
