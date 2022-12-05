package models

import "time"

type Booking struct {
	ID        int64     `json:"id"`
	RoomId    int       `json:"room_id"`
	UserId    int       `json:"user_id"`
	HotelId   int       `json:"hotel_id"`
	FromDate  string    `json:"from_date"`
	ToDate    string    `json:"to_date"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateBookingRequest struct {
	RoomId   int     `json:"room_id"`
	UserId   int     `json:"user_id"`
	HotelId  int     `json:"hotel_id"`
	FromDate string  `json:"from_date"`
	ToDate   string  `json:"to_date"`
	Price    float64 `json:"price"`
}

type GetAllBookingsResponse struct {
	Bookings []*Booking `json:"bookings"`
	Count    int32      `json:"count"`
}
