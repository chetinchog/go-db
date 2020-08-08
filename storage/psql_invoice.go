package storage

import (
	"database/sql"

	"github.com/chetinchog/go-db/pkg/invoice"
	"github.com/chetinchog/go-db/pkg/invoiceheader"
	"github.com/chetinchog/go-db/pkg/invoiceitem"
)

// PsqlInvoice used to work with postgres - invoice
type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

// NewPsqlInvoice returns new pointer to PsqlInvoice
func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

// Create implements interface invoice.Storage
func (p *PsqlInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	// if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	// fmt.Printf("Invoice created with ID: %d\n", m.Header.ID)

	// if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	// fmt.Printf("Items created: %d\n", len(m.Items))

	return tx.Commit()
}
