package repo

import "time"

type Owner struct {
	ID          int64
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber *string
	Username    *string
	Password    string
	CreatedAt   time.Time
}

type GetAllOwnersParams struct {
	Limit  int32
	Page   int32
	Search string
}

type GetAllOwnersResult struct {
	Count  int32
	Owners []*Owner
}

type OwnerStorageI interface {
	Create(u *Owner) (*Owner, error)
	Get(id int64) (*Owner, error)
	GetAll(params *GetAllOwnersParams) (*GetAllOwnersResult, error)
	Update(u *Owner) (*Owner, error)
	Delete(id int64) error
}
