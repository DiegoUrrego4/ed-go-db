package storage

import (
	"database/sql"
	"fmt"
	"github.com/DiegoUrrego4/go-db/pkg/product"
)

type scanner interface {
	Scan(dest ...any) error
}

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
	psqlGetAllProducts = `SELECT * FROM products`
	psqlGetProductByID = psqlGetAllProducts + " WHERE id=$1"
	psqlUpdateProduct  = `UPDATE products SET name= $1, observations= $2, price= $3, updated_at = $4 WHERE id= $5`
	psqlDeleteProduct  = `DELETE from products WHERE id= $1`
)

// psqlProduct used to work with postgres - product
type psqlProduct struct {
	db *sql.DB
}

// newPsqlProduct return a new pointer to sql product
func newPsqlProduct(db *sql.DB) *psqlProduct {
	return &psqlProduct{db: db}
}

// Migrate implements product.storage interface
func (p *psqlProduct) Migrate() error {
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

func (p *psqlProduct) Create(m *product.Model) error {
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
func (p *psqlProduct) GetAll() (product.Models, error) {
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
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

// GetByID implements the interface product.storage
func (p *psqlProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductByID)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

// Update implements the interface product.storage
func (p *psqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(m.Name, stringToNull(m.Observation), m.Price, timeToNull(m.UpdatedAt), m.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no existe el producto con id: %d", m.ID)
	}

	fmt.Println("The register was updated successfully!")
	return nil
}

func (p *psqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("Product deleted successfully!")
	return nil
}

// scanRowProduct helper function to reuse the logic
func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}

	m.Observation = observationNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
