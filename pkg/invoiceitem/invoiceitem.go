package invoiceitem

import (
	"database/sql"
	"time"
)

// Model of invoiceItem
type Model struct {
	ID              uint
	InvoiceHeaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Models []*Model

type Storage interface {
	Migrate() error
	CreateTx(tx *sql.Tx, headerID uint, invoiceItemModel Models) error
}

// Service invoice item service
type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used to migrate product
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
