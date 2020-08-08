package storage

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	// MySQL Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// Postgres Driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Driver of sotrage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New creates connection with DB
func New(d Driver) {
	switch d {
	case Postgres:
		newPostgresDB()
	case MySQL:
		newMySQLDB()
	}
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		// docker run --name postgres -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres
		db, err = gorm.Open("postgres", "postgres://postgres:root@localhost:5432/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}
		fmt.Println("Connected to Postgres")
	})
}

func newMySQLDB() {
	once.Do(func() {
		var err error
		// docker run --name mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=root -e MYSQL_PASSWORD=root -e MYSQL_DATABASE=godb -p 3306:3306 -d mysql
		// docker exec -ti mysql mysql -uroot -proot
		db, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/godb?parseTime=true")
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}
		fmt.Println("Connected to MySQL")
	})
}

// DB returns a unique instance of db
func DB() *gorm.DB {
	return db
}
