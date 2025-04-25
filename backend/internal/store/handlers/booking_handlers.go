package handlers

import (
	"net/http"
	"strconv"

	"github.com/g0shi4ek/store/internal/store/domain"
	"github.com/gin-gonic/gin"
)

func (h * StoreHandler) BookItems(c * gin.Context){
	var input domain.ItemBooking
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.storeService.BookItems(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (h * StoreHandler) BookRooms(c * gin.Context){
	var input domain.RoomBooking
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.storeService.BookRooms(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (h *StoreHandler) GetAllItemBookings(c *gin.Context) {
    items, err := h.storeService.ViewAllItemBookings(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
    }
    c.JSON(http.StatusOK, items) 
}

func (h *StoreHandler) GetAllRoomBookings(c *gin.Context) {
    rooms, err := h.storeService.ViewAllRoomBookings(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
    }
    c.JSON(http.StatusOK, rooms) 
}

func (h *StoreHandler) CancelItemBookings(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    err = h.storeService.CancelItemsBooking(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
    }
    c.Status(http.StatusNoContent)
}


func (h *StoreHandler) CancelRoomBookings(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    err = h.storeService.CancelRoomsBooking(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
    }
    c.Status(http.StatusNoContent)
}


func (h *StoreHandler) BuyItemBookings(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    err = h.storeService.BuyItemsBooking(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
    }
    c.Status(http.StatusNoContent)
}


func (h *StoreHandler) BuyRoomBookings(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    err = h.storeService.BuyRoomsBooking(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
    }
    c.Status(http.StatusOK)
}



