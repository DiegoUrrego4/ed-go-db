package product

import (
	"fmt"
	"strings"
	"time"
)

// Model of product
type Model struct {
	ID          uint
	Name        string
	Observation string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (ms Models) String() string {
	var result strings.Builder
	for _, m := range ms {
		result.WriteString(fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s\n",
			m.ID, m.Name, m.Observation, m.Price,
			m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02")))
	}
	return result.String()
}

type Models []*Model

type Storage interface {
	Migrate() error
	Create(model *Model) error
	//Update(model *Model) error
	GetAll() (Models, error)
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

// Create is used to create a new product
func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}

// GetAll is used to get all the products
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}
