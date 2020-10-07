package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		logrus.WithField("handler", "createInvoice").Errorf("error: %s", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}

	id, err := h.repo.Create(input)
	h.cache.CreateCache(&input)

	if err != nil {
		logrus.WithField("handler", "createInvoice").Errorf("error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getInvoiceById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var invoice invoices.Invoice
	if err != nil {
		logrus.WithField("handler", "getInvoiceById").Errorf("error: %s", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}
	invoice, err = h.cache.GetCache(id)
	if err != nil {
		invoice, err = h.repo.GetById(id)
		if err != nil {
			logrus.WithField("handler", "getInvoiceById").Errorf("error: %s", err.Error())
			c.JSON(http.StatusInternalServerError, errorResponse{
				Error: err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, invoice)
}

func (h *Handler) getAllInvoices(c *gin.Context) {
	var invoice []invoices.Invoice

	invoice, err := h.cache.GetAllCache()
	if err != nil {
		invoice, err = h.repo.GetAll()
		if err != nil {
			logrus.WithField("handler", "getAllInvoices").Errorf("error: %s", err.Error())
			c.JSON(http.StatusBadRequest, errorResponse{
				Error: err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, invoice)
}

func (h *Handler) updateInvoice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithField("handler", "updateInvoice").Errorf("error: %s", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}

	var invoice invoices.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		logrus.WithField("handler", "updateInvoice").Errorf("error: %s", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}

	if err := h.repo.Update(id, invoice); err != nil {
		logrus.WithField("handler", "updateInvoice").Errorf("error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse{
			Error: err.Error(),
		})
		return
	}

	h.cache.UpdateCache(id, invoice)

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deleteInvoice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithField("handler", "deleteInvoice").Errorf("error: %s", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}

	if err := h.repo.Delete(id); err != nil {
		logrus.WithField("handler", "deleteInvoice").Errorf("error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse{
			Error: err.Error(),
		})
		return
	}
	h.cache.DeleteCache(id)

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
