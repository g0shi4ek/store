package handlers

import (
	"github.com/g0shi4ek/store/internal/store/services"
	"github.com/gin-gonic/gin"
)

type StoreHandler struct {
	storeService services.StoreService
}

func NewStoreHandler(serv *services.StoreService) *StoreHandler {
	return &StoreHandler{storeService: *serv}
}

func (h *StoreHandler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		bookings := api.Group("/bookings")
		{
			bookings.POST("/room", h.BookRooms)
			bookings.POST("/item", h.BookItems)
			bookings.GET("/all_items", h.GetAllItemBookings)
			bookings.GET("/all_rooms", h.GetAllRoomBookings)
			bookings.DELETE("/item/:id", h.CancelItemBookings)
			bookings.DELETE("/room/:id", h.CancelRoomBookings)
			bookings.POST("/buy_item/:id", h.BuyItemBookings)
			bookings.POST("/buy_room/:id", h.BuyRoomBookings)
		}
		items := api.Group("/items")
		{
			items.POST("/", h.CreateItem)
			items.POST("/:id", h.UpdateItem)
			items.GET("/all", h.GetAllItems)
			items.GET("/:id", h.GetItem)
		}
		stores := api.Group("/stores")
		{
			stores.GET("/:id", h.GetStore)
			stores.GET("/all", h.GetAllStores)
		}
	}

	return router
}
