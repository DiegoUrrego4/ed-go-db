package main

import (
	"github.com/DiegoUrrego4/go-db/storage"
)

func main() {
	//storage.NewPostgresDB()
	//defer storage.Pool().Close()

	storage.NewMySqlDB()
	storage.NewMySqlDB()
	storage.NewMySqlDB()
}
