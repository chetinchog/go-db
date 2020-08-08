package main

import (
	"fmt"
	"log"

	"github.com/chetinchog/go-db/pkg/product"

	"github.com/chetinchog/go-db/storage"
)

func main() {
	fmt.Println("-----------------------")
	fmt.Println("> Go DB")
	fmt.Println("-----------------------")
	fmt.Println("")

	// // Postgres
	// storage.NewPostgresDB()
	// migratePSQL()
	// createProduct()
	// getList()
	// getProduct()
	// updateProduct()
	// deleteProduct()
	// createInvoice()

	// MySQL
	// storage.NewMySQLDB()
	// migrateMySQL()
	// createMySQLProduct()
	// getMySQLList()
	// getMySQLProduct()
	// updateMySQLProduct()
	// deleteMySQLProduct()
	// createMySQLInvoice()

	// driver := storage.Postgres
	driver := storage.MySQL
	storage.New(driver)
	fmt.Println("")

	myStorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}
	myService := product.NewService(myStorage)

	results, err := myService.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll : %v", err)
	}
	fmt.Println(results)

	result, err := myService.GetByID(2)
	if err != nil {
		log.Fatalf("product.GetByID : %v", err)
	}
	fmt.Println(result)
	fmt.Println("")
}
