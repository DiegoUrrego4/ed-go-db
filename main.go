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

	err := serviceProduct.Delete(1)
	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}
}
