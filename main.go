package main

import (
	"fmt"

	"github.com/chetinchog/go-db/storage"
)

func main() {
	fmt.Println("Go DB")
	storage.NewPostgresDB()

	// migrate()
	// createProduct()
	// getList()
	// getProduct()
	// updateProduct()
	deleteProduct()
}
