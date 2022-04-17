package storage

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "123321"
	dbname = "postgres"
)

type Response struct {
	ID       int    `json:"id"`
	Short    string `json:"short"`
	original string `json:"original"`
}

var (
	ConnectFail = errors.New("connection failed")
	InertFail = errors.New("insertion failed")
	LookupFail = errors.New("lookup failed")
	NotFound = errors.New("URL does not found")
	AlreadyExists = errors.New("URL already exists")
)

func Connect() (*sql.DB, error) {
	cfgcon := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable",
		host,
		port,
		user,
		password,
		dbname,
	)
	db, err := sql.Open("postgres", cfgcon)
	if err != nil {
		return db, ConnectFail
	}

	return db, nil
}

func Insert(short string, original string, db *sql.DB) error {
	query := "SELECT * FROM shortener WHERE short = ($1) OR original = ($2)"
	value := db.QueryRow(query, short, original)
	url := &Response{}
	err := value.Scan(
		&url.ID,
		&url.Short,
		&url.original,
	)
	if err == sql.ErrNoRows {

		query = "INSERT INTO shortener(short, original) VALUES($1, $2)"
		_, err = db.Exec(query, short, original)
		if err != nil {
			return InertFail
		}
	} else {
		return AlreadyExists
	}
	return nil
}

func Lookup(short string, db *sql.DB) (string, error) {
	query := "SELECT * FROM shortener WHERE short = ($1)"
	value := db.QueryRow(query, short)
	url := &Response{}
	err := value.Scan(
		&url.ID,
		&url.Short,
		&url.original,
	)
	if err != nil {
		return "", LookupFail
	}
	return url.original, nil
}
