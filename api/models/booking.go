package models

import "time"

type Booking struct {
	ID            int64     `json:"id"`
	RoomId        int       `json:"room_id"`
	UserId        int       `json:"user_id"`
	Stay          string    `json:"stay"`
	NumberOfUsers int       `json:"number_of_users"`
	StayDate      time.Time `json:"stay_date"`
	CreatedAt     time.Time `json:"created_at"`
}

type CreateBookingRequest struct {
	RoomId        int       `json:"room_id"`
	UserId        int       `json:"user_id"`
	Stay          string    `json:"stay"`
	NumberOfUsers int       `json:"number_of_users"`
	StayDate      time.Time `json:"stay_date"`
}

type GetAllBookingsResponse struct {
	Bookings []*Booking `json:"bookings"`
	Count    int32      `json:"count"`
}
