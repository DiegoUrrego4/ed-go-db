package storage

import (
	"database/sql"
	"fmt"
	"github.com/DiegoUrrego4/go-db/pkg/product"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once // Esta estructura es para crear el singleton
)

type Driver string

const (
	MySql    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New create the connection with DB
func New(d Driver) {
	switch d {
	case MySql:
		newMySqlDB()
	case Postgres:
		newPostgresDB()
	}
}

func newPostgresDB() {
	once.Do(func() { // esto se ejecuta una sola vez
		var err error

		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Can't load environment variables: %v", err)
		}

		dbUser := os.Getenv("POSTGRES_USER")
		dbPass := os.Getenv("POSTGRES_PASS")
		dbName := os.Getenv("POSTGRES_DB")

		connStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", dbUser, dbPass, dbName)

		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping: %v", err)
		}

		fmt.Println("Conectado a postgres")
	})
}

func newMySqlDB() {
	once.Do(func() { // esto se ejecuta una sola vez
		var err error

		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Can't load environment variables: %v", err)
		}

		dbPass := os.Getenv("MYSQL_PASS")
		dbName := os.Getenv("MYSQL_DB")

		connStr := fmt.Sprintf("root:%s@tcp(localhost:3306)/%s", dbPass, dbName)

		db, err = sql.Open("mysql", connStr)
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping: %v", err)
		}

		fmt.Println("Conectado a Mysql")
	})
}

// Pool returns an unique instance of db
func Pool() *sql.DB {
	return db
}

// stringToNull handles null strings values
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

// timeToNull handles null times values
func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{
		Time: t,
	}

	if !null.Time.IsZero() {
		null.Valid = true
	}

	return null
}

// DAOProduct factory of product.Storage
func DAOProduct(driver Driver) (product.Storage, error) {
	switch driver {
	case Postgres:
		return newPsqlProduct(db), nil
	case MySql:
		return newMySqlProduct(db), nil
	default:
		return nil, fmt.Errorf("driver not implemented")
	}
}
