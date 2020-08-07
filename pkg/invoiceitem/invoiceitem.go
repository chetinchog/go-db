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

// Models slice of Model
type Models []*Model

// Storage interface that must implement a db storage
type Storage interface {
	Migrate() error
}

// Service of invoiceitem
type Service struct {
	storage Storage
}

// NewService returns a pointer to Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used to migrate invoiceitem
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
