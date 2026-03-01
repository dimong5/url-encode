package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Создаём таблицу если нет
	_, errDbCreate := db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			short VARCHAR(10) PRIMARY KEY,
			original TEXT NOT NULL
		)
	`)
	if errDbCreate != nil {
		return nil, errDbCreate
	}

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Save(short, original string) error {
	_, err := s.db.Exec(
		"INSERT INTO urls (short, original) VALUES ($1, $2) ON CONFLICT (short) DO UPDATE SET original = $2",
		short, original,
	)
	return err
}

func (s *PostgresStorage) Get(short string) (string, error) {
	var original string
	err := s.db.QueryRow("SELECT original FROM urls WHERE short = $1", short).Scan(&original)
	if err == sql.ErrNoRows {
		return "", ErrNotFound
	}
	if err != nil {
		return "", err
	}
	return original, nil
}

func (s *PostgresStorage) FindByOriginal(original string) string {
	var short string
	err := s.db.QueryRow("SELECT short FROM urls WHERE original = $1", original).Scan(&short)
	if err != nil {
		return ""
	}
	return short
}