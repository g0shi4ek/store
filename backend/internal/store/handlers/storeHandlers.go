package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *StoreHandler) GetAllStores(c *gin.Context) {
	items, err := h.storeService.ViewAllStores(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *StoreHandler) GetStore(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.storeService.ViewStore(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}