package invoiceheader

import "time"

// Model of invoice header
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Storage interface {
	Migrate() error
}

// Service invoice header service
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
