package main

import (
	"fmt"
	"github.com/DiegoUrrego4/go-db/pkg/product"
	"github.com/DiegoUrrego4/go-db/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()
	defer storage.Pool().Close()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	//m := &product.Model{
	//	Name:        "Curso de BBDD con Go",
	//	Price:       70,
	//	Observation: "On fire",
	//}

	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(ms)

	//storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	//serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)
	//if err := serviceInvoiceHeader.Migrate(); err != nil {
	//	log.Fatalf("invoiceHeader.Migrate: %v", err)
	//}

	//storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	//serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	//if err := serviceInvoiceItem.Migrate(); err != nil {
	//	log.Fatalf("invoiceItem.Migrate: %v", err)
	//}
}
