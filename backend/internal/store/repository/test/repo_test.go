package test

import (
	"context"
	"testing"

	"github.com/g0shi4ek/store/config"
	"github.com/g0shi4ek/store/internal/store/domain"
	"github.com/g0shi4ek/store/internal/store/repository/postgres"
	"github.com/g0shi4ek/store/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestStoreRepository(t *testing.T) {

	cfg := &config.Config{
		DbConf: config.DbConfig{
			DbHost:     "localhost",
			DbPort:     "5432",
			DbUser:     "postgres",
			DbPassword: "1111",
			DbName:     "store_db",
		},
		StoreConf: config.StoreConfig{
			Port: "8000",
		},
	}
	ctx := context.Background()

	pool, err := db.NewPool(ctx, cfg)
	assert.NoError(t, err)
	defer pool.Close()

	repoStore := postgres.NewStoresRepository(pool)

	t.Run("Create,Get,Delete Store", func(t *testing.T) {
		store := &domain.Store{
			Num:          1,
			RoomTotal:    10,
			RoomBooked:   0,
			RoomOccupied: 0,
		}

		err := repoStore.Create(ctx, store)
		assert.NoError(t, err)
		assert.NotZero(t, store.Id)

		found, err := repoStore.GetById(ctx, store.Id)
		assert.NoError(t, err)
		assert.Equal(t, store.Num, found.Num)

		err = repoStore.Delete(ctx, store.Id)
		assert.NoError(t, err)
	})

	repoItem := postgres.NewItemsRepository(pool)

	t.Run("Create,Get,Delete Item", func(t *testing.T) {
		item := &domain.Item{
			Name:     "book_6",
			ItemTotal:   10,
			Provider: "mmm",
			Price:    10,
		}

		err := repoItem.Create(ctx, item)
		assert.NoError(t, err)
		assert.NotZero(t, item.Id)

		found, err := repoItem.GetById(ctx, item.Id)
		assert.NoError(t, err)
		assert.Equal(t, item.Name, found.Name)

		err = repoStore.Delete(ctx, item.Id)
		assert.NoError(t, err)
	})

	repoBooking := postgres.NewBookingRepository(pool)

	t.Run("Create,Get,Delete Booking", func(t *testing.T) {

		b := &domain.ItemBooking{
			ItemId:   3,
			Amount:   3,
			IsActive: false,
		}

		_, err = repoBooking.CreateItemBooking(ctx, b)
		assert.NoError(t, err)
		assert.NotZero(t, b.Id)

		b.IsActive = true

		err = repoBooking.UpdateItemBooking(ctx, b)
		assert.NoError(t, err)
		assert.Equal(t, b.IsActive, true)

		err = repoBooking.DeleteItemBooking(ctx, b.Id)
		assert.NoError(t, err)
	})
}
