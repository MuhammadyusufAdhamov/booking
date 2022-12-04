package storage

import (
	"github.com/MuhammadyusufAdhamov/booking/storage/postgres"
	"github.com/MuhammadyusufAdhamov/booking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserStorageI
	Owner() repo.OwnerStorageI
	Hotel() repo.HotelStorageI
	Room() repo.RoomsStorageI
	Booking() repo.BookingsStorageI
}

type storagePg struct {
	userRepo    repo.UserStorageI
	ownerRepo   repo.OwnerStorageI
	hotelRepo   repo.HotelStorageI
	roomRepo    repo.RoomsStorageI
	bookingRepo repo.BookingsStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		userRepo:    postgres.NewUser(db),
		ownerRepo:   postgres.NewOwner(db),
		hotelRepo:   postgres.NewHotel(db),
		roomRepo:    postgres.NewRoom(db),
		bookingRepo: postgres.NewBooking(db),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s *storagePg) Owner() repo.OwnerStorageI {
	return s.ownerRepo
}

func (s *storagePg) Hotel() repo.HotelStorageI {
	return s.hotelRepo
}

func (s *storagePg) Room() repo.RoomsStorageI {
	return s.roomRepo
}

func (s *storagePg) Booking() repo.BookingsStorageI {
	return s.bookingRepo
}
