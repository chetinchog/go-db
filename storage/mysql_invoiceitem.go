package storage

import (
	"database/sql"
	"fmt"
)

// MySQLInvoiceItem used to work with MySQL - invoiceitem
type MySQLInvoiceItem struct {
	db *sql.DB
}

// NewMySQLInvoiceItem returns new pointer to MySQLInvoiceItem
func NewMySQLInvoiceItem(db *sql.DB) *MySQLInvoiceItem {
	return &MySQLInvoiceItem{db}
}

// Migrate implements interface invoiceitem.Storage
func (p *MySQLInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("InvoiceItem migration Succeeded!")
	return nil
}
