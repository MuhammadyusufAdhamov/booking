package models

import "time"

type Room struct {
	ID           int64     `json:"id"`
	Type         string    `json:"type"`
	NumberOfRoom int       `json:"number_of_room"`
	RoomImageUrl *string   `json:"room_image_url"`
	Status       string    `json:"status"`
	HotelId      int       `json:"hotel_id"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateRoomRequest struct {
	Type         string  `json:"type"`
	NumberOfRoom int     `json:"number_of_room"`
	RoomImageUrl *string `json:"room_image_url"`
	Status       string  `json:"status"`
	HotelId      int     `json:"hotel_id"`
}

type GetAllRoomsResponse struct {
	Rooms []*Room `json:"rooms"`
	Count int32   `json:"count"`
}
