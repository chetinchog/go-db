package main

import (
	"github.com/chetinchog/go-db/godbgorm/model"
	"github.com/chetinchog/go-db/godbgorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	storage.DB().AutoMigrate(
		&model.Product{},
		&model.InvoiceHeader{},
		&model.InvoiceItem{},
	)
}
