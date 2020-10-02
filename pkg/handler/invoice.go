package handler

import (
	"github.com/gin-gonic/gin"
	"invoices"
	"net/http"
	"strconv"
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
	h.repo.Create(&input)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": input.Id,
	})

	return
}

func (h *Handler) getInvoiceById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	invoice,err := h.repo.Get(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func (h *Handler) getAllInvoices(c *gin.Context) {
	invoice, err := h.repo.GetAll()
	if err != nil{
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func (h *Handler) updateInvoice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil{
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}
	var invoice invoices.Invoice

	h.repo.Update(id,invoice)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": invoice.Id,
	})
}

func (h *Handler) deleteInvoice(c *gin.Context) {
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}
	h.repo.Delete(id)
	c.JSON(http.StatusOK, "invoice deleted")
}
