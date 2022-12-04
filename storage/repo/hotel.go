package repo

import "time"

type Hotel struct {
	ID            int64
	OwnerID       int64
	HotelName     string
	HotelRating   string
	HotelLocation string
	HotelImageUrl *string
	NumberOfRooms int32
	CreatedAt     time.Time
}

type GetAllHotelsParams struct {
	Limit  int32
	Page   int32
	Search string
}

type GetAllHotelsResult struct {
	Hotels []*Hotel
	Count  int32
}

type HotelStorageI interface {
	Create(u *Hotel) (*Hotel, error)
	Get(id int64) (*Hotel, error)
	GetAll(params *GetAllHotelsParams) (*GetAllHotelsResult, error)
	Update(u *Hotel) (*Hotel, error)
	Delete(id int64) error
}
