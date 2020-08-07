package invoiceheader

import "time"

// Model of invoiceheader
type Model struct {
	ID        uint
	Client    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
