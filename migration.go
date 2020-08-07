package main

import (
	"log"

	"github.com/chetinchog/go-db/pkg/invoiceheader"
	"github.com/chetinchog/go-db/pkg/invoiceitem"
	"github.com/chetinchog/go-db/pkg/product"
	"github.com/chetinchog/go-db/storage"
)

func migrateProduct() {
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
}
func migrateInvoiceHeader() {
	storageInvoiceheader := storage.NewPsqlInvoiceHeader(storage.Pool())
	serviceInvoiceheader := invoiceheader.NewService(storageInvoiceheader)
	if err := serviceInvoiceheader.Migrate(); err != nil {
		log.Fatalf("invoiceheader.Migrate: %v", err)
	}
}
func migrateInvoiceItem() {
	storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceitem.Migrate: %v", err)
	}
}

func migrate() {
	migrateProduct()
	migrateInvoiceHeader()
	migrateInvoiceItem()
}
