package v1

import (
	"errors"
	"github.com/MuhammadyusufAdhamov/booking/api/models"
	"github.com/MuhammadyusufAdhamov/booking/config"
	"github.com/MuhammadyusufAdhamov/booking/storage"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	ErrWrongEmailOrPass = errors.New("wrong email or password")
	ErrEmailExists      = errors.New("email already exists")
	ErrUserNotVerified  = errors.New("user not verified")
	ErrIncorrectCode    = errors.New("incorrect verification code")
	ErrCodeExpired      = errors.New("verification code has been expired")
	ErrForbidden        = errors.New("forbidden")
)

type handlerV1 struct {
	cfg      *config.Config
	storage  storage.StorageI
	inMemory storage.InMemoryStorageI
}

type HandlerV1Options struct {
	Cfg      *config.Config
	Storage  storage.StorageI
	inMemory storage.InMemoryStorageI
}

//goland:noinspection GoExportedFuncWithUnexportedType
func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg:      options.Cfg,
		storage:  options.Storage,
		inMemory: options.inMemory,
	}
}

func errorResponse(err error) *models.ErrorResponse {
	return &models.ErrorResponse{
		Error: err.Error(),
	}
}

func validateGetAllParams(c *gin.Context) (*models.GetAllParams, error) {
	var (
		limit int = 10
		page  int = 1
		err   error
	)

	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("page") != "" {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			return nil, err
		}
	}

	return &models.GetAllParams{
		Limit:  int32(limit),
		Page:   int32(page),
		Search: c.Query("search"),
	}, nil
}
