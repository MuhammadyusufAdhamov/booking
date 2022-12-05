package models

import "time"

type Hotel struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user_id"`
	HotelName     string    `json:"hotel_name"`
	HotelLocation string    `json:"hotel_location"`
	HotelImageUrl *string   `json:"hotel_image_url"`
	NumberOfRooms int32     `json:"number_of_rooms"`
	CreatedAt     time.Time `json:"created_at"`
}

type CreateHotelRequest struct {
	UserID        int64   `json:"user_id" binding:"required"`
	HotelName     string  `json:"hotel_name" binding:"required"`
	HotelLocation string  `json:"hotel_location" binding:"required"`
	HotelImageUrl *string `json:"hotel_image_url" binding:"required"`
	NumberOfRooms int32   `json:"number_of_rooms" binding:"required"`
}

type GetAllHotelsResponse struct {
	Hotels []*Hotel `json:"hotels"`
	Count  int32    `json:"count"`
}
