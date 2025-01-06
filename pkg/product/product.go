package product

import "time"

// Model of product
type Model struct {
	ID          uint
	Name        string
	Observation string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Models []*Model

type Storage interface {
	Migrate() error
	//Create(model *Model) error
	//Update(model *Model) error
	//GetAll() (Models, error)
	//GetByID(id uint) (*Model, error)
	//Delete(id uint) error
}

// Service product service
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
