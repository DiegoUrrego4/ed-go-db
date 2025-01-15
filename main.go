package main

import (
	"github.com/DiegoUrrego4/go-db/pkg/product"
	"github.com/DiegoUrrego4/go-db/storage"
	"log"
)

func main() {
	driver := storage.Postgres

	storage.New(driver)
	myStorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}

	serviceProduct := product.NewService(myStorage)
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
}
