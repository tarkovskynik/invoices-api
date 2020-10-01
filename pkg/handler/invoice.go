package handler

import (
	"github.com/gin-gonic/gin"
	"invoices"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func (h *Handler) createInvoice(c *gin.Context) {
	var input invoices.Invoice

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}

	_, err := h.repo.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{
			Error: err.Error(),
		})
		return
	}

	return
}

func (h *Handler) getInvoiceById(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllInvoices(c *gin.Context) {

}

func (h *Handler) updateInvoice(c *gin.Context) {

}

func (h *Handler) deleteInvoice(c *gin.Context) {

}
