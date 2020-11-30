package database

import (
	"SeminarioGoLang/internal/config"
	"errors"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jmoiron/sqlx"
)

// NewDataBase ...
func NewDataBase(conf *config.Config) (*sqlx.DB, error) {
	switch conf.DB.Type {
	case "sqlite3":
		db, err := sqlx.Open(conf.DB.Driver, conf.DB.Conn)
		if err != nil {
			return nil, err
		}

		err = db.Ping()
		if err != nil {
			return nil, err
		}

		return db, nil
	default:
		return nil, errors.New("invalid db type")
	}
}

//CreateSchema ...
func CreateSchema(s *sqlx.DB) error {

	schema := "CREATE TABLE IF NOT EXISTS 'Doctor' ( " +
		"id integer PRIMARY KEY AUTOINCREMENT, " +
		"name varchar(50) NOT NULL, " +
		"enrollment varchar(100) NOT NULL, " +
		"age integer NOT NULL)"
	_, err := s.Exec(schema)

	if err != nil {
		return err
	}

	return nil

}
