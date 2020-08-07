package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	// Load driver
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// NewPostgresDB sets singleton db
func NewPostgresDB() {
	once.Do(func() {
		var err error
		// docker run --name postgres -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres
		db, err = sql.Open("postgres", "postgres://postgres:root@localhost:5432/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}
		if err := db.Ping(); err != nil {
			log.Fatalf("Can't ping db: %v", err)
		}
		fmt.Println("Connected to PostgresDB")
	})
}

// Pool returns a unique instance of db
func Pool() *sql.DB {
	return db
}
