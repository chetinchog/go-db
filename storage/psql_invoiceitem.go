package storage

import (
	"database/sql"
	"fmt"
)

// PsqlInvoiceitem used to work with postgres - invoiceitem
type PsqlInvoiceitem struct {
	db *sql.DB
}

// NewPsqlInvoiceItem returns new pointer to PsqlInvoiceitem
func NewPsqlInvoiceItem(db *sql.DB) *PsqlInvoiceitem {
	return &PsqlInvoiceitem{db}
}

// Migrate implements interface invoiceitem.Storage
func (p *PsqlInvoiceitem) Migrate() error {
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
