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
	var id int
	row := r.db.QueryRow("INSERT INTO invoices(title, description, company_name, date, total_cost) VALUES( $1, $2, $3, $4, $5 ) RETURNING id",
		invoice.Title, invoice.Description, invoice.CompanyName, invoice.Date, invoice.TotalCost)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *InvoiceRepository) GetAll() ([]invoices.Invoice, error) {
	rows, err := r.db.Query("SELECT id, title, description, company_name, date, total_cost FROM invoices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []invoices.Invoice
	for rows.Next() {
		var invoice invoices.Invoice
		err := rows.Scan(&invoice.Id, &invoice.Title, &invoice.Description, &invoice.CompanyName, &invoice.Date, &invoice.TotalCost)
		if err != nil {
			return nil, err
		}

		results = append(results, invoice)
	}

	return results, nil
}

func (r *InvoiceRepository) GetById(id int) (invoices.Invoice, error) {
	var invoice invoices.Invoice

	row := r.db.QueryRow("SELECT id, title, description, company_name, date, total_cost FROM invoices WHERE id=$1", id)
	err := row.Scan(&invoice.Id, &invoice.Title, &invoice.Description, &invoice.CompanyName, &invoice.Date, &invoice.TotalCost)
	if err != nil {
		if err == sql.ErrNoRows {
			return invoices.Invoice{}, invoices.NewErrorInvoiceNotFound(id)
		}

		return invoices.Invoice{}, err
	}

	return invoice, nil
}

func (r *InvoiceRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM invoices WHERE id=$1", id)
	return err
}

func (r *InvoiceRepository) Update(id int, invoice invoices.Invoice) error {
	_, err := r.db.Exec("UPDATE invoices SET title=$1, description=$2, company_name=$3, date=$4, total_cost=$5 WHERE id=$6",
		invoice.Title, invoice.Description, invoice.CompanyName, invoice.Date, invoice.TotalCost, id)
	return err
}
