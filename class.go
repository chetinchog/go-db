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

// Postgres
func migratePSQLProduct() {
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
}
func migratePSQLInvoiceHeader() {
	storageInvoiceheader := storage.NewPsqlInvoiceHeader(storage.Pool())
	serviceInvoiceheader := invoiceheader.NewService(storageInvoiceheader)
	if err := serviceInvoiceheader.Migrate(); err != nil {
		log.Fatalf("invoiceheader.Migrate: %v", err)
	}
}
func migratePSQLInvoiceItem() {
	storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceitem.Migrate: %v", err)
	}
}

func migratePSQL() {
	migratePSQLProduct()
	migratePSQLInvoiceHeader()
	migratePSQLInvoiceItem()
}

// func createProduct() {
// 	storageProduct := storage.NewPsqlProduct(storage.Pool())
// 	serviceProduct := product.NewService(storageProduct)

// 	m := &product.Model{
// 		Name:         "Curso de DB con Go",
// 		Price:        70,
// 		Observations: "On fire!",
// 	}
// 	if err := serviceProduct.Create(m); err != nil {
// 		log.Fatalf("product.Create: %v", err)
// 	}
// 	fmt.Printf("%+v\n", m)
// }

// func getList() {
// 	storageProduct := storage.NewPsqlProduct(storage.Pool())
// 	serviceProduct := product.NewService(storageProduct)

// 	listProduct, err := serviceProduct.GetAll()
// 	if err != nil {
// 		log.Fatalf("product.GetAll: %v", err)
// 	}
// 	fmt.Println("Get All Products")
// 	fmt.Println(listProduct)
// }

// func getProduct() {
// 	storageProduct := storage.NewPsqlProduct(storage.Pool())
// 	serviceProduct := product.NewService(storageProduct)

// 	prod, err := serviceProduct.GetByID(1)
// 	switch {
// 	case errors.Is(err, sql.ErrNoRows):
// 		fmt.Println("Product not found")
// 	case err != nil:
// 		log.Fatalf("product.GetAll: %v", err)
// 	default:
// 		fmt.Println("Get Product")
// 		fmt.Println(prod)
// 	}
// }

// func updateProduct() {
// 	storageProduct := storage.NewPsqlProduct(storage.Pool())
// 	serviceProduct := product.NewService(storageProduct)

// 	m := &product.Model{
// 		ID:   90,
// 		Name: "Curso Testing",
// 		// Observations: "This is the course",
// 		Price: 50,
// 	}
// 	if err := serviceProduct.Update(m); err != nil {
// 		log.Fatalf("product.Update: %v", err)
// 	}
// }

// func deleteProduct() {
// 	storageProduct := storage.NewPsqlProduct(storage.Pool())
// 	serviceProduct := product.NewService(storageProduct)

// 	if err := serviceProduct.Delete(1); err != nil {
// 		log.Fatalf("product.Delete: %v", err)
// 	}
// }

// func createInvoice() {
// 	storage.NewPostgresDB()
// 	storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
// 	storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
// 	storageInvoice := storage.NewPsqlInvoice(storage.Pool(), storageHeader, storageItems)

// 	m := &invoice.Model{
// 		Header: &invoiceheader.Model{Client: "Alexys"},
// 		Items: invoiceitem.Models{
// 			&invoiceitem.Model{ProductID: 2},
// 		},
// 	}
// 	serviceInvoice := invoice.NewService(storageInvoice)
// 	if err := serviceInvoice.Create(m); err != nil {
// 		log.Fatalf("invoice.Create: %v", err)
// 	}
// }

// MySQL
func migrateMySQLProduct() {
	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
}
func migrateMySQLInvoiceHeader() {
	storageInvoiceHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)
	if err := serviceInvoiceHeader.Migrate(); err != nil {
		log.Fatalf("invoiceheader.Migrate: %v", err)
	}
}
func migrateMySQLInvoiceItem() {
	storageInvoiceItem := storage.NewMySQLInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceitem.Migrate: %v", err)
	}
}

func migrateMySQL() {
	migrateMySQLProduct()
	migrateMySQLInvoiceHeader()
	migrateMySQLInvoiceItem()
}

func createMySQLProduct() {
	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		Name:  "Curso de testing con Go",
		Price: 120,
		// Observations: "On fire!",
	}
	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Create: %v\n", err)
	}
	fmt.Printf("%+v\n", m)
}

func getMySQLList() {
	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	listProduct, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}
	fmt.Println("Get All Products")
	fmt.Println(listProduct)
}

func getMySQLProduct() {
	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	prod, err := serviceProduct.GetByID(2)
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

func updateMySQLProduct() {
	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		ID:           1,
		Name:         "Curso CSS",
		Observations: "CSS!!",
		Price:        200,
	}
	if err := serviceProduct.Update(m); err != nil {
		log.Fatalf("product.Update: %v", err)
	}
}

func deleteMySQLProduct() {
	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Delete(3); err != nil {
		log.Fatalf("product.Delete: %v", err)
	}
}
