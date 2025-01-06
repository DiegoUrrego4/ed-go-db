package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlCreateProduct = `CREATE TABLE IF NOT EXISTS products(
    id SERIAL NOT NULL,
    name VARCHAR(25),
    observations VARCHAR(100),
    price INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    CONSTRAINT products_id_pk PRIMARY KEY (id)
)`
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
	stmt, err := p.db.Prepare(psqlCreateProduct)
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
