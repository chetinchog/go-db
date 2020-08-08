package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	// Load Postgres driver
	_ "github.com/lib/pq"
	// Load MySQL driver
	_ "github.com/go-sql-driver/mysql"
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

// NewMySQLDB sets singleton db
func NewMySQLDB() {
	once.Do(func() {
		var err error
		// docker run --name mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=root -e MYSQL_PASSWORD=root -e MYSQL_DATABASE=godb -p 3306:3306 -d mysql
		// docker exec -ti mysql mysql -uroot -proot
		db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/godb?parseTime=true")
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}
		if err := db.Ping(); err != nil {
			log.Fatalf("Can't ping db: %v", err)
		}
		fmt.Println("Connected to MySQLDB")
	})
}

// Pool returns a unique instance of db
func Pool() *sql.DB {
	return db
}
