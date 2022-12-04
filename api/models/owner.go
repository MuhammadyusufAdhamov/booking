package models

import "time"

type Owner struct {
	ID          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber *string   `json:"phone_number"`
	Username    *string   `json:"username"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateOwnerRequest struct {
	FirstName   string  `json:"first_name" binding:"required,min=2,max=30"`
	LastName    string  `json:"last_name" binding:"required,min=2,max=30"`
	Email       string  `json:"email" binding:"required,email"`
	PhoneNumber *string `json:"phone_number"`
	Username    *string `json:"username" binding:"required,min=2,max=30"`
	Password    string  `json:"password" binding:"required,min=6,max=16"`
}

type GetAllOwnersResponse struct {
	Owners []*Owner `json:"owners"`
	Count  int32    `json:"count"`
}
