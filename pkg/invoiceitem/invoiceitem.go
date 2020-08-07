package invoiceitem

import "time"

// Model of invoiceitem
type Model struct {
	ID              uint
	InvoiceHeaderID uint
	PorudctID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
