package storage

import (
	"database/sql"
	"fmt"

	"github.com/chetinchog/go-db/pkg/invoiceheader"
)

// MySQLInvoiceHeader used to work with MySQL - invoiceheader
type MySQLInvoiceHeader struct {
	db *sql.DB
}

// NewMySQLInvoiceHeader returns new pointer to MySQLInvoiceHeader
func NewMySQLInvoiceHeader(db *sql.DB) *MySQLInvoiceHeader {
	return &MySQLInvoiceHeader{db}
}

// Migrate implements interface invoiceheader.Storage
func (p *MySQLInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("InvoiceHeader migration Succeeded!")
	return nil
}

// CreateTx implements interface invoiceheader.Storage
func (p *MySQLInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(mySQLCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(m.Client)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = uint(id)
	return nil
}
