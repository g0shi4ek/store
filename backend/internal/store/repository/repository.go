package repository

import (
	"github.com/g0shi4ek/store/internal/store/repository/postgres"

	"github.com/g0shi4ek/store/internal/store/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)


type postgresRepository struct {
    pg * pgxpool.Pool
}

func NewRepository(pgx *pgxpool.Pool) domain.IRepository {
    return &postgresRepository{pg: pgx}
}

func (r *postgresRepository) Store() domain.IStoreRepository {
    return postgres.NewStoresRepository(r.pg)
}

func (r *postgresRepository) Item() domain.IItemRepository {
    return postgres.NewItemsRepository(r.pg)
}

func (r *postgresRepository) Booking() domain.IBookingRepository {
    return postgres.NewBookingRepository(r.pg)
}

func (r *postgresRepository) User() domain.IUserRepository {
    return postgres.NewUserRepository(r.pg)
}