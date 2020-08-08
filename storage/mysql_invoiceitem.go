package storage

import (
	"database/sql"
	"fmt"

	"github.com/chetinchog/go-db/pkg/invoiceitem"
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

// CreateTx implements interface invoiceitem.Storage
func (p *MySQLInvoiceItem) CreateTx(tx *sql.Tx, headerID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(mySQLCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range ms {
		result, err := stmt.Exec(headerID, item.ProductID)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		item.ID = uint(id)
	}

	return nil
}
