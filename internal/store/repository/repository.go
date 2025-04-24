package repository

import (
	"github.com/g0shi4ek/store/internal/store/repository/postgres"

	"github.com/g0shi4ek/store/internal/store/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)


type Repository interface {
    Store() 	domain.StoreRepository
    Item() 		domain.ItemRepository
    Booking() 	domain.BookingRepository
}

type postgresRepository struct {
    pg * pgxpool.Pool
}

func NewRepository(pgx *pgxpool.Pool) Repository {
    return &postgresRepository{pg: pgx}
}

func (r *postgresRepository) Store() domain.StoreRepository {
    return postgres.NewStoresRepository(r.pg)
}

func (r *postgresRepository) Item() domain.ItemRepository {
    return postgres.NewItemsRepository(r.pg)
}

func (r *postgresRepository) Booking() domain.BookingRepository {
    return postgres.NewBookingRepository(r.pg)
}