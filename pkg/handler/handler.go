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

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"user1": "qwerty",
		"user2": "asd123",
		"user3": "asd321",
	}))

	authorized.POST("/invoice", h.createInvoice)
	authorized.GET("/invoice/:id", h.getInvoiceById)
	authorized.GET("/invoices", h.getAllInvoices)
	authorized.PUT("/invoice/:id", h.updateInvoice)
	authorized.DELETE("/invoice/:id", h.deleteInvoice)

	return r.Run(":8080")
}
