package domain

type Item struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	Provider string `json:"provider"`
	ItemBooked int //добавить забронированные 
	Price    int    `json:"price"`
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
