package main

import (
	"fmt"

	"github.com/chetinchog/go-db/godbgorm/model"
	"github.com/chetinchog/go-db/godbgorm/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	// // Migration
	// storage.DB().AutoMigrate(
	// 	&model.Product{},
	// 	&model.InvoiceHeader{},
	// 	&model.InvoiceItem{},
	// )

	// // Create
	// product1 := model.Product{
	// 	Name:  "Curso de GO",
	// 	Price: 120,
	// }
	// obs := "Testing con GO"
	// product2 := model.Product{
	// 	Name:         "Curso de Testing",
	// 	Observations: &obs,
	// 	Price:        150,
	// }
	// product3 := model.Product{
	// 	Name:  "Curso de Python",
	// 	Price: 200,
	// }
	// storage.DB().Create(&product1)
	// storage.DB().Create(&product2)
	// storage.DB().Create(&product3)

	// // Find
	// listProduct := make([]model.Product, 0)
	// storage.DB().Find(&listProduct)
	// for _, product := range listProduct {
	// 	fmt.Printf("%d - %s\n", product.ID, product.Name)
	// }

	// Find One
	product := model.Product{}
	storage.DB().First(&product, 3)
	fmt.Printf("%d - %s\n", product.ID, product.Name)
}
