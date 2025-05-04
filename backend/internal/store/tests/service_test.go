package tests

import (
	"context"
	"testing"

	"github.com/g0shi4ek/store/config"
	"github.com/g0shi4ek/store/internal/store/domain"
	"github.com/g0shi4ek/store/internal/store/mocks"
	"github.com/g0shi4ek/store/internal/store/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestStoreService_RegisterUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// моки репы
	mockRepo := mocks.NewMockIRepository(ctrl)
	mockUserRepo := mocks.NewMockIUserRepository(ctrl)
	
	cfg := &config.Config{}

	// ожидания
	testUser := &domain.User{
		Username: "testuser",
		PasswordHash: "password123",
	}

	// у мока главной репы вызовется метод User() и вернется mockUserRepo
	mockRepo.EXPECT().
		User().
		Return(mockUserRepo).
		Times(1)

	// у мока репы юзера вызов Create с любым контекстом и тестовым пользователем
	mockUserRepo.EXPECT().
		Create(
			gomock.Any(),
			gomock.AssignableToTypeOf(&domain.User{}),
		).
		DoAndReturn(func(ctx context.Context, u *domain.User) error {
			assert.Equal(t, "testuser", u.Username)
			assert.True(t, len(u.PasswordHash) > 0)
			return nil
		}).
		Times(1)

	service := services.NewStoreService(mockRepo, cfg)
	err := service.RegisterUser(context.Background(), testUser)

	assert.NoError(t, err)
}
