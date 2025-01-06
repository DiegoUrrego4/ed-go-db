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
	Create(model *Model) error
	Update(model *Model) error
	GetAll() (Models, error)
	GetByID(id uint) (*Model, error)
	Delete(id uint) error
}
