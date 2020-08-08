package storage

import (
	"database/sql"
	"fmt"

	"github.com/chetinchog/go-db/pkg/invoice"
	"github.com/chetinchog/go-db/pkg/invoiceheader"
	"github.com/chetinchog/go-db/pkg/invoiceitem"
)

// MySQLInvoice used to work with MySQL - invoice
type MySQLInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

// NewMySQLInvoice returns new pointer to MySQLInvoice
func NewMySQLInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *MySQLInvoice {
	return &MySQLInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

// Create implements interface invoice.Storage
func (p *MySQLInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return err
	}
	fmt.Printf("Invoice created with ID: %d\n", m.Header.ID)

	if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return err
	}
	fmt.Printf("Items created: %d\n", len(m.Items))

	return tx.Commit()
}
