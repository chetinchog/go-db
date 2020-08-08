package invoiceheader

import (
	"database/sql"
	"time"
)

// Model of invoiceheader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Models slice of Model
type Models []*Model

// Storage interface that must implement a db storage
type Storage interface {
	Migrate() error
	CreateTx(*sql.Tx, *Model) error
}

// Service of invoiceheader
type Service struct {
	storage Storage
}

// NewService returns a pointer to Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used to migrate invoiceheader
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
