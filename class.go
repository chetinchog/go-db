package main

import (
	"database/sql"
	"errors"
	"fmt"
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

func createProduct() {
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		Name:         "Curso de DB con Go",
		Price:        70,
		Observations: "On fire!",
	}
	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Create: %v", err)
	}
	fmt.Printf("%+v\n", m)
}

func getList() {
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	listProduct, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}
	fmt.Println("Get All Products")
	fmt.Println(listProduct)
}

func getProduct() {
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	prod, err := serviceProduct.GetByID(1)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("Product not found")
	case err != nil:
		log.Fatalf("product.GetAll: %v", err)
	default:
		fmt.Println("Get Product")
		fmt.Println(prod)
	}
}

func updateProduct() {
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		ID:   90,
		Name: "Curso Testing",
		// Observations: "This is the course",
		Price: 50,
	}
	if err := serviceProduct.Update(m); err != nil {
		log.Fatalf("product.Update: %v", err)
	}
}

func deleteProduct() {
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Delete(1); err != nil {
		log.Fatalf("product.Delete: %v", err)
	}
}
