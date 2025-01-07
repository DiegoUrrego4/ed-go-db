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

	//m := &product.Model{
	//	ID:   19,
	//	Name: "Curso de testing",
	//	//Observation: "Este curso tiene otro nivel.",
	//	Price: 150,
	//}

	err := serviceProduct.Delete(1)
	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}

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
