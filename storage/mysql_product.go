package storage

import (
	"database/sql"
	"fmt"
)

const (
	mySqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(25),
    observations VARCHAR(100),
    price INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP
)`
)

// mySqlProduct used to work with postgres - product
type mySqlProduct struct {
	db *sql.DB
}

// newMySqlProduct return a new pointer to sql product
func newMySqlProduct(db *sql.DB) *mySqlProduct {
	return &mySqlProduct{db: db}
}

// Migrate implements product.storage interface
func (m *mySqlProduct) Migrate() error {
	stmt, err := m.db.Prepare(mySqlMigrateProduct)
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
