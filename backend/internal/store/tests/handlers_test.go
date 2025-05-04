package tests

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/g0shi4ek/store/config"
	"github.com/g0shi4ek/store/internal/store/handlers"
	"github.com/g0shi4ek/store/internal/store/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_Register(t *testing.T) {

	cfg := &config.Config{}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мок сервиса
	mockService := mocks.NewMockIStoreService(ctrl)
	handler := handlers.NewStoreHandler(mockService, cfg)

	userJSON := `{"username":"test","password":"pass"}`

	mockService.EXPECT().
		RegisterUser(gomock.Any(), gomock.Any()).
		Return(nil)

	req := httptest.NewRequest("POST", "/register", strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := gin.Default()
	router.POST("/register", handler.RegisterUser)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}
