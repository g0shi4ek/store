package handlers

import (
	"strings"
	"time"

	"github.com/g0shi4ek/store/config"
	"github.com/g0shi4ek/store/internal/store/domain"
	"github.com/g0shi4ek/store/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type StoreHandler struct {
	storeService domain.IStoreService
	cfg          *config.Config
}

func NewStoreHandler(serv domain.IStoreService, cfg *config.Config) *StoreHandler {
	return &StoreHandler{
		storeService: serv,
		cfg:          cfg,
	}
}

func (h *StoreHandler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Cache-Control", "no-store, no-cache, must-revalidate")

		// Для API отключаем кэширование
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.Header("Pragma", "no-cache")
			c.Header("Expires", "0")
		}
		c.Next()
	})

	// Улучшенный CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://store-front-4ulo.onrender.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Явная обработка OPTIONS запросов
	router.OPTIONS("/*any", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "https://store-front-4ulo.onrender.com")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,X-Requested-With")
		c.Status(200)
	})

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
