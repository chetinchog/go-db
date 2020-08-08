package storage

import (
	"database/sql"
	"fmt"

	"github.com/chetinchog/go-db/pkg/invoiceheader"
)

// PsqlInvoiceHeader used to work with postgres - invoiceheader
type PsqlInvoiceHeader struct {
	db *sql.DB
}

// NewPsqlInvoiceHeader returns new pointer to PsqlInvoiceHeader
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

// Migrate implements interface invoiceheader.Storage
func (p *PsqlInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
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
func (p *PsqlInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRow(m.Client).Scan(&m.ID, &m.CreatedAt)
}
