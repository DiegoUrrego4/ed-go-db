package storage

import (
	"database/sql"
	"fmt"
)

const (
	mySqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    invoice_header_id INT NOT NULL,
    product_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
    CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
)`
)

// MySqlInvoiceItem used to work with postgres - invoice items
type MySqlInvoiceItem struct {
	db *sql.DB
}

// NewMySqlInvoiceItem return a new pointer to sql invoice items
func NewMySqlInvoiceItem(db *sql.DB) *MySqlInvoiceItem {
	return &MySqlInvoiceItem{db}
}

// Migrate implements invoice items.storage interface
func (m *MySqlInvoiceItem) Migrate() error {
	stmt, err := m.db.Prepare(mySqlMigrateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Invoice items migration executed successfully!")
	return nil
}
