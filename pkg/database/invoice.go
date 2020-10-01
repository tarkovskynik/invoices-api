package database

import (
	"database/sql"
	"invoices"
)

type InvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) Create(invoice invoices.Invoice) (int, error) {
	//row := r.db.QueryRow()
	return 0, nil
}