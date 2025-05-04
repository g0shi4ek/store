package domain

import "context"

//go:generate mockgen -destination=../mocks/repositoryMock.go -package=mocks . IRepository
//go:generate mockgen -destination=../mocks/userRepositoryMock.go -package=mocks . IUserRepository
//go:generate mockgen -destination=../mocks/itemRepositoryMock.go -package=mocks . IItemRepository
//go:generate mockgen -destination=../mocks/storeRepositoryMock.go -package=mocks . IStoreRepository
//go:generate mockgen -destination=../mocks/bookingRepositoryMock.go -package=mocks . IBookingRepository
//go:generate mockgen -destination=../mocks/serviceMock.go -package=mocks . IStoreService


type IStoreService interface{
	BookItems(ctx context.Context, booking *ItemBooking) (int, error)
	BookRooms(ctx context.Context, booking *RoomBooking) (int, error)
	CancelItemsBooking(ctx context.Context, bookingId int) error
	CancelRoomsBooking(ctx context.Context, bookingId int) error
	BuyItemsBooking(ctx context.Context, bookingId int) error
	BuyRoomsBooking(ctx context.Context, bookingId int) error
	ViewAllItemBookings(ctx context.Context) ([]*ItemBooking, error)
	ViewAllRoomBookings(ctx context.Context) ([]*RoomBooking, error)
	ViewItem(ctx context.Context, itemId int) (*Item, error)
	ViewStore(ctx context.Context, storeId int) (*Store, error)
	AddItem(ctx context.Context, item *Item) error
	UpdateItem(ctx context.Context, item *Item) error
	ViewAllItems(ctx context.Context) ([]*Item, error)
	ViewAllStores(ctx context.Context) ([]*Store, error)
	RegisterUser(ctx context.Context, user *User) error
	LoginUser(ctx context.Context, username, password string) (string, string, error)
}

type IUserRepository interface{
	Create(ctx context.Context, user* User) error
	GetByUsername(ctx context.Context, username string) (*User, error)
}

type IItemRepository interface {
	Create(ctx context.Context, item * Item) error
	GetById(ctx context.Context, id int) (*Item, error)
	GetByName(ctx context.Context, name string) (*Item, error)
	GetAll(ctx context.Context) ([]*Item, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, item *Item) error
}

type IStoreRepository interface {
	Create(ctx context.Context, store *Store) error
	GetById(ctx context.Context, id int) (*Store, error)
	GetAll(ctx context.Context) ([]*Store, error)
	GetByNum(ctx context.Context, number int) (*Store, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, store *Store) error
}

type IBookingRepository interface {
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

type IRepository interface {
    Store() 	IStoreRepository
    Item() 		IItemRepository
    Booking() 	IBookingRepository
    User()      IUserRepository
}