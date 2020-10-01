package invoices

type Invoice struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	CompanyName string  `json:"company_name"`
	Date        string  `json:"date"`
	TotalCost   float32 `json:"total_cost"`
}