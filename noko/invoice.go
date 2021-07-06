package noko

// Invoice represents a Noko invoice
type Invoice struct {
	ID          int     `json:"id"`
	Reference   string  `json:"reference"`
	InvoiceDate string  `json:"invoice_date"`
	State       string  `json:"state"`
	TotalAmount float64 `json:"total_amount"`
	URL         string  `json:"url"`
}
