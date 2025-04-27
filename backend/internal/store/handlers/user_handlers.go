package handlers

import (
	"net/http"

	"github.com/g0shi4ek/store/internal/store/domain"
	"github.com/gin-gonic/gin"
)

func (h *StoreHandler) RegisterUser(c *gin.Context) {
	var input struct{ Username, Password, Role string }
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser := domain.User{
		Username:     input.Username,
		PasswordHash: input.Password,
		Role:         input.Role,
	}
	err := h.storeService.RegisterUser(c.Request.Context(), &newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *StoreHandler) LoginUser(c *gin.Context) {
	var input struct{ Username, Password string }
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	role, token, err := h.storeService.LoginUser(c.Request.Context(), input.Username, input.Password)
	if token == "" || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Header("Authorization", token)
	c.JSON(http.StatusOK, gin.H{"token": token, "role": role})
}
