package storage

import (
	"database/sql"
	"fmt"
)

const (
	mySqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    client VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP
)`
)

// MySqlInvoiceHeader used to work with postgres - invoice header
type MySqlInvoiceHeader struct {
	db *sql.DB
}

// NewMySqlInvoiceHeader return a new pointer to sql invoice header
func NewMySqlInvoiceHeader(db *sql.DB) *MySqlInvoiceHeader {
	return &MySqlInvoiceHeader{db}
}

// Migrate implements invoice header.storage interface
func (m *MySqlInvoiceHeader) Migrate() error {
	stmt, err := m.db.Prepare(mySqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Invoice header migration executed successfully!")
	return nil
}

//func (p *PsqlInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
//	stmt, err := tx.Prepare(psqlCreateInvoiceHeader)
//	if err != nil {
//		return err
//	}
//	defer stmt.Close()
//
//	return stmt.QueryRow(m.Client).Scan(&m.ID, &m.CreatedAt)
//
//}
