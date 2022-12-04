package models

import "time"

type Hotel struct {
	ID            int64     `json:"id"`
	OwnerID       int64     `json:"owner_id"`
	HotelName     string    `json:"hotel_name"`
	HotelRating   string    `json:"hotel_rating"`
	HotelLocation string    `json:"hotel_location"`
	HotelImageUrl *string   `json:"hotel_image_url"`
	NumberOfRooms int32     `json:"number_of_rooms"`
	CreatedAt     time.Time `json:"created_at"`
}

type CreateHotelRequest struct {
	OwnerID       int64   `json:"owner_id" binding:"required"`
	HotelName     string  `json:"hotel_name" binding:"required"`
	HotelRating   string  `json:"hotel_rating" binding:"required"`
	HotelLocation string  `json:"hotel_location" binding:"required"`
	HotelImageUrl *string `json:"hotel_image_url" binding:"required"`
	NumberOfRooms int32   `json:"number_of_rooms" binding:"required"`
}

type GetAllHotelsResponse struct {
	Hotels []*Hotel `json:"hotels"`
	Count  int32    `json:"count"`
}
