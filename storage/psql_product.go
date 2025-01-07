package storage

import (
	"database/sql"
	"fmt"
	"github.com/DiegoUrrego4/go-db/pkg/product"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
    id SERIAL NOT NULL,
    name VARCHAR(25),
    observations VARCHAR(100),
    price INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    CONSTRAINT products_id_pk PRIMARY KEY (id)
)`
	psqlCreateProduct = `INSERT INTO products(name, observations, price, created_at) 
VALUES ($1, $2, $3, $4) RETURNING id
`
	psqlGetAllProducts = `SELECT id, name, observations, price, created_at, updated_at FROM products`
)

// PsqlProduct used to work with postgres - product
type PsqlProduct struct {
	db *sql.DB
}

// NewPsqlProduct return a new pointer to sql product
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db: db}
}

// Migrate implements product.storage interface
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Product migration executed successfully!")
	return nil
}

func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observation),
		m.Price,
		m.CreatedAt).Scan(&m.ID)
	if err != nil {
		return err
	}

	fmt.Println("Product created successfully!")
	return nil
}

// GetAll implements the interface product.storage
func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProducts)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(product.Models, 0)
	for rows.Next() {
		m := &product.Model{}
		observationNull := sql.NullString{}
		updatedAtNull := sql.NullTime{}

		err := rows.Scan(
			&m.ID,
			&m.Name,
			&observationNull,
			&m.Price,
			&m.CreatedAt,
			&updatedAtNull,
		)
		if err != nil {
			return nil, err
		}

		m.Observation = observationNull.String
		m.UpdatedAt = updatedAtNull.Time
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}
