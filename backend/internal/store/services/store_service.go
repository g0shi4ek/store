package services

import (
	"context"
	"fmt"
	"log"

	"github.com/g0shi4ek/store/config"
	"github.com/g0shi4ek/store/internal/store/domain"
	"github.com/g0shi4ek/store/internal/store/repository"
	"github.com/g0shi4ek/store/pkg/jwt"
)

type StoreService struct {
	repo repository.Repository
	cfg  *config.Config
}

func NewStoreService(repo repository.Repository, cfg *config.Config) *StoreService {
	return &StoreService{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *StoreService) BookItems(ctx context.Context, booking *domain.ItemBooking) (int, error) {
	booking, err := s.repo.Booking().CreateItemBooking(ctx, booking)
	if err != nil {
		return -1, err
	}
	log.Println("add items booking")
	return booking.Id, nil
}

func (s *StoreService) BookRooms(ctx context.Context, booking *domain.RoomBooking) (int, error) {
	booking, err := s.repo.Booking().CreateRoomBooking(ctx, booking)
	if err != nil {
		return -1, err
	}
	log.Println("add rooms booking")
	return booking.Id, nil
}

func (s *StoreService) CancelItemsBooking(ctx context.Context, bookingId int) error {
	err := s.repo.Booking().DeleteItemBooking(ctx, bookingId)
	log.Println("delete items booking")
	return err
}

func (s *StoreService) CancelRoomsBooking(ctx context.Context, bookingId int) error {
	err := s.repo.Booking().DeleteRoomBooking(ctx, bookingId)
	log.Println("delete rooms booking")
	return err
}

func (s *StoreService) BuyItemsBooking(ctx context.Context, bookingId int) error {
	booking, err := s.repo.Booking().GetItemBookingById(ctx, bookingId)
	if err != nil {
		return err
	}
	booking.IsActive = true
	err = s.repo.Booking().UpdateItemBooking(ctx, booking)
	log.Println("buy items")
	return err
}

func (s *StoreService) BuyRoomsBooking(ctx context.Context, bookingId int) error {
	booking, err := s.repo.Booking().GetRoomBookingById(ctx, bookingId)
	if err != nil {
		return err
	}
	booking.IsActive = true
	err = s.repo.Booking().UpdateRoomBooking(ctx, booking)
	log.Println("buy rooms")
	return err
}

func (s *StoreService) ViewAllItemBookings(ctx context.Context) ([]*domain.ItemBooking, error) {
	itemBookings, err := s.repo.Booking().GetAllItemBookings(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("get all item bookings")
	return itemBookings, nil
}

func (s *StoreService) ViewAllRoomBookings(ctx context.Context) ([]*domain.RoomBooking, error) {
	roomBookings, err := s.repo.Booking().GetAllRoomBookings(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("get allroom bookings")
	return roomBookings, nil
}

func (s *StoreService) ViewItem(ctx context.Context, itemId int) (*domain.Item, error) {
	item, err := s.repo.Item().GetById(ctx, itemId)
	if err != nil {
		return nil, err
	}
	log.Println("get item")
	return item, nil
}

func (s *StoreService) ViewStore(ctx context.Context, storeId int) (*domain.Store, error) {
	store, err := s.repo.Store().GetById(ctx, storeId)
	if err != nil {
		return nil, err
	}
	log.Println("get store")
	return store, nil
}

func (s *StoreService) AddItem(ctx context.Context, item *domain.Item) error {
	err := s.repo.Item().Create(ctx, item)
	log.Println("add item")
	return err
}

func (s *StoreService) UpdateItem(ctx context.Context, item *domain.Item) error {
	err := s.repo.Item().Update(ctx, item)
	log.Println("update item")
	return err
}

func (s *StoreService) ViewAllItems(ctx context.Context) ([]*domain.Item, error) {
	items, err := s.repo.Item().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("get all items")
	return items, nil
}

func (s *StoreService) ViewAllStores(ctx context.Context) ([]*domain.Store, error) {
	rooms, err := s.repo.Store().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("get all stores")
	return rooms, nil
}


func (s * StoreService) RegisterUser(ctx context.Context, user * domain.User) error{
	hash, err := HashPassword(user.PasswordHash)
	if err != nil{
		return err
	}
	user.PasswordHash = hash
	err = s.repo.User().Create(ctx, user)
	return err
}

func (s * StoreService) LoginUser(ctx context.Context, username, password string) (string, error){
	user, err := s.repo.User().GetByUsername(ctx, username)
	if err != nil{
		return "", err
	}
	validPassword := user.PasswordHash
	if !CheckPassword(validPassword, password) {
		return "", fmt.Errorf("wrong password")
	}
	token, err := jwt.CreateNewToken(user, s.cfg)
	if err != nil{
		return "", err
	}
	return token, nil
} 
