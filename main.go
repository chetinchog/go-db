package main

import (
	"fmt"

	"github.com/chetinchog/go-db/storage"
)

func main() {
	fmt.Println("Go DB")
	// storage.NewPostgresDB()

	// // Postgres
	// migratePSQL()
	// createProduct()
	// getList()
	// getProduct()
	// updateProduct()
	// deleteProduct()
	// createInvoice()

	// MySQL
	storage.NewMySQLDB()
	// migrateMySQL()
	// createMySQLProduct()
	// getMySQLList()
	// getMySQLProduct()
	// updateMySQLProduct()
	// deleteMySQLProduct()
	createMySQLInvoice()
}
