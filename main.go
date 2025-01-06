package main

import (
	"github.com/DiegoUrrego4/go-db/pkg/product"
	"github.com/DiegoUrrego4/go-db/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()
	defer storage.Pool().Close()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
}
