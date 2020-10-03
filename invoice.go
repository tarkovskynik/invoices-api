package invoices

import "fmt"

func NewErrorInvoiceNotFound(id int) error {
	return fmt.Errorf("invoice with id #%d not found", id)
}

type Invoice struct {
	Id          int     `json:"id"`
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	CompanyName string  `json:"company_name" binding:"required"`
	Date        string  `json:"date"`
	TotalCost   float32 `json:"total_cost" binding:"required"`
}
