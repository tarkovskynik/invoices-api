package handler

import (
	"github.com/gin-gonic/gin"
	"invoices/pkg/cache"
	"invoices/pkg/database"
)

type Handler struct {
	repo  *database.InvoiceRepository
	cache *cache.Cache
}

func NewHandler(repo *database.InvoiceRepository, cache *cache.Cache) *Handler {
	return &Handler{
		repo:  repo,
		cache: cache,
	}

}

func (h *Handler) Init() error {
	r := gin.New()

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "qwerty",
	}))

	authorized.POST("/invoice", h.createInvoice)
	authorized.GET("/invoice/:id", h.getInvoiceById)
	authorized.GET("/invoices", h.getAllInvoices)
	authorized.PUT("/invoice/:id", h.updateInvoice)
	authorized.DELETE("/invoice/:id", h.deleteInvoice)

	return r.Run(":8080")
}
