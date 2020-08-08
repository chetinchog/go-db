package storage

import (
	"database/sql"
	"fmt"

	"github.com/chetinchog/go-db/pkg/invoiceitem"
)

// PsqlInvoiceItem used to work with postgres - invoiceitem
type PsqlInvoiceItem struct {
	db *sql.DB
}

// NewPsqlInvoiceItem returns new pointer to PsqlInvoiceItem
func NewPsqlInvoiceItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db}
}

// Migrate implements interface invoiceitem.Storage
func (p *PsqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)
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
func (p *PsqlInvoiceItem) CreateTx(tx *sql.Tx, headerID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range ms {
		if err = stmt.QueryRow(headerID, item.ProductID).Scan(
			&item.ID,
			&item.CreatedAt,
		); err != nil {
			return err
		}
	}

	return nil
}
