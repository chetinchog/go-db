package main

import (
	"github.com/chetinchog/go-db/godbgorm/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)
}
