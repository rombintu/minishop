package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rombintu/minishop/config"
)

type Store struct {
	Database *sql.DB
	Config   *config.Postgres
}

// Struct for test
type Ping struct {
	Message string `json:"message"`
}

func (s *Store) Open() error {
	if !s.Config.Dev {
		connStr := fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=%s",
			s.Config.User,
			s.Config.Password,
			s.Config.Dbname,
			s.Config.SSLMode,
		)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			return err
		}
		s.Database = db
		return nil
	}
	db, err := sql.Open("sqlite3", "dev.db")
	if err != nil {
		return err
	}
	s.Database = db
	return nil
}

func (s *Store) Close() {
	s.Database.Close()
}
