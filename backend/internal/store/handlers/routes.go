package handlers

import (
	"github.com/g0shi4ek/store/config"
	"github.com/g0shi4ek/store/internal/store/services"
	"github.com/g0shi4ek/store/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type StoreHandler struct {
	storeService services.StoreService
	cfg          *config.Config
}

func NewStoreHandler(serv *services.StoreService, cfg *config.Config) *StoreHandler {
	return &StoreHandler{
		storeService: *serv,
		cfg:          cfg,
	}
}

func (h *StoreHandler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	api.POST("/register", h.RegisterUser)
	api.POST("/login", h.LoginUser)

	{
		manager := api.Group("/manager")
		manager.Use(middleware.AuthMiddleware("manager", h.cfg))
		{
			bookings := manager.Group("/bookings")
			{
				bookings.POST("/room", h.BookRooms)
				bookings.DELETE("/room/:id", h.CancelRoomBookings)
				bookings.GET("/all_rooms", h.GetAllRoomBookings)
				bookings.POST("/buy_room/:id", h.BuyRoomBookings)
			}
			items := manager.Group("/items")
			{
				items.POST("/", h.CreateItem)
				items.POST("/:id", h.UpdateItem)
				items.GET("/all", h.GetAllItems)
				items.GET("/:id", h.GetItem)
			}
			stores := manager.Group("/stores")
			{
				stores.GET("/:id", h.GetStore)
				stores.GET("/all", h.GetAllStores)
			}
		}

		seller := api.Group("/seller")
		seller.Use(middleware.AuthMiddleware("seller", h.cfg))
		{
			bookings := seller.Group("/bookings")
			{
				bookings.POST("/item", h.BookItems)
				bookings.GET("/all_items", h.GetAllItemBookings)
				bookings.DELETE("/item/:id", h.CancelItemBookings)
				bookings.POST("/buy_item/:id", h.BuyItemBookings)
			}
			items := api.Group("/items")
			{
				items.GET("/all", h.GetAllItems)
				items.GET("/:id", h.GetItem)
			}
			stores := api.Group("/stores")
			{
				stores.GET("/:id", h.GetStore)
				stores.GET("/all", h.GetAllStores)
			}

		}

	}

	return router
}
