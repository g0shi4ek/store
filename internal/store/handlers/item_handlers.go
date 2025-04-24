package handlers

import (
	"net/http"
	"strconv"

	"github.com/g0shi4ek/store/internal/store/domain"
	"github.com/gin-gonic/gin"
)

func (h * StoreHandler) CreateItem(c * gin.Context){
	var input domain.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.storeService.AddItem(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h * StoreHandler) UpdateItem(c * gin.Context){
	var input domain.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.storeService.UpdateItem(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *StoreHandler) GetAllItems(c *gin.Context) {
    items, err := h.storeService.ViewAllItems(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
    }
    c.JSON(http.StatusOK, items)
}

func (h *StoreHandler) GetItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    item, err := h.storeService.ViewItem(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
    }
    c.JSON(http.StatusOK, item)
}
