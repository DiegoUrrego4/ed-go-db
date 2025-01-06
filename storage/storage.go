package storage

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once // Esta estructura es para crear el singleton
)

func NewPostgresDB() {
	once.Do(func() { // esto se ejecuta una sola vez
		var err error
		db, err = sql.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}
		defer db.Close()

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping: %v", err)
		}
	})
}
