package invoiceheader

import "time"

// Model of invoice header
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
