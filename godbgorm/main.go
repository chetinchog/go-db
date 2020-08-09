package main

import (
	"github.com/chetinchog/go-db/godbgorm/model"
	"github.com/chetinchog/go-db/godbgorm/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	// Migration
	storage.DB().AutoMigrate(
		&model.Product{},
		&model.InvoiceHeader{},
		&model.InvoiceItem{},
	)
	storage.DB().Model(&model.InvoiceItem{}).AddForeignKey(
		"product_id",
		"products(id)",
		"RESTRICT",
		"RESTRICT",
	)
	storage.DB().Model(&model.InvoiceItem{}).AddForeignKey(
		"invoice_header_id",
		"invoice_headers(id)",
		"RESTRICT",
		"RESTRICT",
	)

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

	// // Find One
	// product := model.Product{}
	// storage.DB().First(&product, 3)
	// fmt.Printf("%d - %s\n", product.ID, product.Name)

	// // Update
	// product := model.Product{}
	// product.ID = 3
	// storage.DB().Model(&product).Updates(
	// 	model.Product{Name: "Curso de CSS", Price: 120},
	// )

	// // Delete Soft
	// product := model.Product{}
	// product.ID = 3
	// storage.DB().Delete(&product)

	// // Delete Hard
	// product := model.Product{}
	// product.ID = 3
	// storage.DB().Unscoped().Delete(&product)
}
