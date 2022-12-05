package repo

import "time"

type Room struct {
	ID           int64
	Type         string
	NumberOfRoom int
	RoomImageUrl *string
	Status       string
	HotelId      int
	CreatedAt    time.Time
}

type GetAllRoomsParams struct {
	Limit  int32
	Page   int32
	Search string
}

type GetAllRoomsResult struct {
	Rooms []*Room
	Count int32
}

type RoomsStorageI interface {
	Create(u *Room) (*Room, error)
	Get(id int64) (*Room, error)
	GetAll(params *GetAllRoomsParams) (*GetAllRoomsResult, error)
	Update(u *Room) (*Room, error)
	Delete(id int64) error
}
