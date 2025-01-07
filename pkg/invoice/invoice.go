package invoice

import (
	"github.com/DiegoUrrego4/go-db/pkg/invoiceheader"
	"github.com/DiegoUrrego4/go-db/pkg/invoiceitem"
)

// Model of invoice
type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

// Storage interface that must implement a db storage
type Storage interface {
	Create(model *Model) error
}

// Service of an invoice
type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

// Create a new invoice
func (s Service) Create(m *Model) error {
	return s.storage.Create(m)
}
