package domain

import "context"

type UserRepository interface{
	Create(ctx context.Context, user* User) error
	GetByUsername(ctx context.Context, username string) (*User, error)
}

type ItemRepository interface {
	Create(ctx context.Context, item * Item) error
	GetById(ctx context.Context, id int) (*Item, error)
	GetByName(ctx context.Context, name string) (*Item, error)
	GetAll(ctx context.Context) ([]*Item, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, item *Item) error
}

type StoreRepository interface {
	Create(ctx context.Context, store *Store) error
	GetById(ctx context.Context, id int) (*Store, error)
	GetAll(ctx context.Context) ([]*Store, error)
	GetByNum(ctx context.Context, number int) (*Store, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, store *Store) error
}

type BookingRepository interface {
	CreateRoomBooking(ctx context.Context, booking *RoomBooking) (*RoomBooking, error)
	UpdateRoomBooking(ctx context.Context, booking *RoomBooking) error
	DeleteRoomBooking(ctx context.Context, id int) error
	GetRoomBookingById(ctx context.Context, id int) (*RoomBooking, error)
	GetAllRoomBookings(ctx context.Context) ([]*RoomBooking, error)

	CreateItemBooking(ctx context.Context, booking *ItemBooking) (*ItemBooking, error)
	UpdateItemBooking(ctx context.Context, booking *ItemBooking) error
	DeleteItemBooking(ctx context.Context, id int) error
	GetItemBookingById(ctx context.Context, id int) (*ItemBooking, error)
	GetAllItemBookings(ctx context.Context) ([]*ItemBooking, error)
}
