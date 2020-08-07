package main

import (
	"fmt"
	"log"

	"github.com/chetinchog/go-db/pkg/product"

	"github.com/chetinchog/go-db/storage"
)

func main() {
	fmt.Println("Go DB")
	storage.NewPostgresDB()
	// migrate()

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
