package postgres

import (
	"context"

	"github.com/g0shi4ek/store/internal/store/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StoresRepository struct {
	pg *pgxpool.Pool
}

func NewStoresRepository(pgx *pgxpool.Pool) domain.IStoreRepository {
	return &StoresRepository{pg: pgx}
}

func (r *StoresRepository) Create(ctx context.Context, store *domain.Store) error {
	query := "INSERT INTO stores (number, room_total, room_booked, room_occupied) VALUES ($1, $2, $3, $4) RETURNING id"
	return r.pg.QueryRow(ctx, query, store.Num, store.RoomTotal, store.RoomBooked, store.RoomOccupied).Scan(&store.Id)
}

func (r *StoresRepository) Update(ctx context.Context, store *domain.Store) error {
	query := "UPDATE stores SET number = $1, room_total = $2, room_booked = $3, room_occupied = $4 WHERE id = $5"
	_, err := r.pg.Exec(ctx, query, store.Num, store.RoomTotal, store.RoomBooked, store.RoomOccupied, store.Id)
	return err
}

func (r *StoresRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM stores WHERE id = $1"
	_, err := r.pg.Exec(ctx, query, id)
	return err
}

func (r *StoresRepository) GetById(ctx context.Context, id int) (*domain.Store, error) {
	var newStore domain.Store
	query := "SELECT id, number, room_total, room_booked, room_occupied FROM stores WHERE id = $1"
	err := r.pg.QueryRow(ctx, query, id).Scan(&newStore.Id, &newStore.Num, &newStore.RoomTotal, &newStore.RoomBooked, &newStore.RoomOccupied)
	if err != nil {
		return nil, err
	}
	return &newStore, nil
}

func (r *StoresRepository) GetByNum(ctx context.Context, num int) (*domain.Store, error) {
	var newStore domain.Store
	query := "SELECT id, number, room_total, room_booked, room_occupied FROM stores WHERE number = $1"
	err := r.pg.QueryRow(ctx, query, num).Scan(&newStore.Id, &newStore.Num, &newStore.RoomTotal, &newStore.RoomBooked, &newStore.RoomOccupied)
	if err != nil {
		return nil, err
	}
	return &newStore, nil
}

func (r *StoresRepository) GetAll(ctx context.Context) ([]*domain.Store, error) {
	query := "SELECT id, number, room_total, room_booked, room_occupied FROM stores"
	rows, err := r.pg.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stores []*domain.Store
	for rows.Next() {
		var store domain.Store
		err := rows.Scan(&store.Id, &store.Num, &store.RoomTotal, &store.RoomBooked, &store.RoomOccupied)
		if err != nil {
			return nil, err
		}
		stores = append(stores, &store)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stores, nil
}
