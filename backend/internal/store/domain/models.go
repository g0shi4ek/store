package domain

type User struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
}

type Item struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	ItemTotal  int    `json:"item_total"`
	ItemBooked int    `json:"item_booked"`
	Provider   string `json:"provider"`
	Price      int    `json:"price"`
}

type Store struct {
	Id           int `json:"id"`
	Num          int `json:"number"`
	RoomTotal    int `json:"room_total"`
	RoomBooked   int `json:"room_booked"`
	RoomOccupied int `json:"room_occupied"`
}

type RoomBooking struct {
	Id       int  `json:"id"`
	StoreId  int  `json:"store_id"`
	ItemId   int  `json:"item_id"`
	Amount   int  `json:"amount"`
	IsActive bool `json:"is_active"`
}

type ItemBooking struct {
	Id       int  `json:"id"`
	ItemId   int  `json:"item_id"`
	Amount   int  `json:"amount"`
	IsActive bool `json:"is_active"`
}
