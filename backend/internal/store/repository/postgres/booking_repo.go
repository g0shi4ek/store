package postgres

import (
	"context"

	"github.com/g0shi4ek/store/internal/store/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookingRepository struct {
	pg *pgxpool.Pool
}

func NewBookingRepository(pgx *pgxpool.Pool) domain.BookingRepository {
	return &BookingRepository{pg: pgx}
}

func (r *BookingRepository) CreateRoomBooking(ctx context.Context, booking *domain.RoomBooking) (*domain.RoomBooking, error) {
	query := "INSERT INTO room_bookings (store_id, item_id, amount) VALUES ($1, $2, $3) RETURNING id"
	err := r.pg.QueryRow(ctx, query, booking.StoreId, booking.ItemId, booking.Amount).Scan(&booking.Id)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (r *BookingRepository) UpdateRoomBooking(ctx context.Context, booking *domain.RoomBooking) error {
	query := "UPDATE room_bookings SET is_active = $1 WHERE id = $2"
	_, err := r.pg.Exec(ctx, query, booking.IsActive, booking.Id)
	return err
}

func (r *BookingRepository) DeleteRoomBooking(ctx context.Context, id int) error {
	query := "DELETE FROM room_bookings WHERE id = $1"
	_, err := r.pg.Exec(ctx, query, id)
	return err
}

func (r *BookingRepository) GetRoomBookingById(ctx context.Context, id int) (*domain.RoomBooking, error) {
	var newBooking domain.RoomBooking
	query := "SELECT id, store_id, item_id, amount, is_active FROM room_bookings WHERE id = $1"
	err := r.pg.QueryRow(ctx, query, id).Scan(&newBooking.Id, &newBooking.StoreId, &newBooking.ItemId, &newBooking.Amount, &newBooking.IsActive)
	if err != nil {
		return nil, err
	}
	return &newBooking, nil
}

func (r *BookingRepository) GetAllRoomBookings(ctx context.Context) ([]*domain.RoomBooking, error) {
	query := "SELECT id, store_id, item_id, amount, is_active FROM room_bookings"
	rows, err := r.pg.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*domain.RoomBooking
	for rows.Next() {
		var b domain.RoomBooking
		err := rows.Scan(&b.Id, &b.StoreId, &b.ItemId, &b.Amount, &b.IsActive)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, &b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *BookingRepository) CreateItemBooking(ctx context.Context, booking *domain.ItemBooking) (*domain.ItemBooking, error) {
	query := "INSERT INTO item_bookings (item_id, amount) VALUES ($1, $2) RETURNING id"
	err := r.pg.QueryRow(ctx, query, booking.ItemId, booking.Amount).Scan(&booking.Id)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (r *BookingRepository) UpdateItemBooking(ctx context.Context, booking *domain.ItemBooking) error {
	query := "UPDATE item_bookings SET is_active = $1 WHERE id = $2"
	_, err := r.pg.Exec(ctx, query, booking.IsActive, booking.Id)
	return err
}

func (r *BookingRepository) DeleteItemBooking(ctx context.Context, id int) error {
	query := "DELETE FROM item_bookings WHERE id = $1"
	_, err := r.pg.Exec(ctx, query, id)
	return err
}

func (r *BookingRepository) GetItemBookingById(ctx context.Context, id int) (*domain.ItemBooking, error) {
	var newBooking domain.ItemBooking
	query := "SELECT id, item_id, amount, is_active FROM item_bookings WHERE id = $1"
	err := r.pg.QueryRow(ctx, query, id).Scan(&newBooking.Id, &newBooking.ItemId, &newBooking.Amount, &newBooking.IsActive)
	if err != nil {
		return nil, err
	}
	return &newBooking, nil
}

func (r *BookingRepository) GetAllItemBookings(ctx context.Context) ([]*domain.ItemBooking, error) {
	query := "SELECT id, item_id, amount, is_active FROM item_bookings"
	rows, err := r.pg.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*domain.ItemBooking
	for rows.Next() {
		var b domain.ItemBooking
		err := rows.Scan(&b.Id, &b.ItemId, &b.Amount, &b.IsActive)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, &b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}
