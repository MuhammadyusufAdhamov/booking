package repo

import "time"

type Booking struct {
	ID            int64
	RoomId        int
	UserId        int
	Stay          string
	NumberOfUsers int
	StayDate      time.Time
	CreatedAt     time.Time
}

type GetAllBookingsParams struct {
	Limit  int32
	Page   int32
	Search string
}

type GetAllBookingResult struct {
	Bookings []*Booking
	Count    int32
}

type BookingsStorageI interface {
	Create(u *Booking) (*Booking, error)
	Get(id int64) (*Booking, error)
	GetAll(params *GetAllBookingsParams) (*GetAllBookingResult, error)
	Update(u *Booking) (*Booking, error)
	Delete(id int64) error
}
