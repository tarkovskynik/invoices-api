package handler

import (
	"github.com/gin-gonic/gin"
	"invoices/pkg/database"
)

type Handler struct {
	repo *database.InvoiceRepository
}

func NewHandler(repo *database.InvoiceRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Init() error {
	r := gin.New()

	r.POST("/invoice", h.createInvoice)
	r.GET("/invoice/:id", h.getInvoiceById)
	r.GET("/invoices", h.getAllInvoices)
	r.PUT("/invoice/:id", h.getInvoiceById)
	r.DELETE("/invoice/:id", h.deleteInvoice)

	return r.Run(":8080")
}
