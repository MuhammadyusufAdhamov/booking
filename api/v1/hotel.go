package v1

import (
	"database/sql"
	"errors"
	"github.com/MuhammadyusufAdhamov/booking/api/models"
	"github.com/MuhammadyusufAdhamov/booking/storage/repo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Router /hotels [post]
// @Summary Create a hotel
// @Description Create a hotel
// @Tags hotel
// @Accept json
// @Produce json
// @Param hotel body models.CreateHotelRequest true "Hotel"
// @Success 201 {object} models.Hotel
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateHotel(c *gin.Context) {
	var (
		req models.CreateHotelRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.Hotel().Create(&repo.Hotel{
		UserID:        req.UserID,
		HotelName:     req.HotelName,
		HotelLocation: req.HotelLocation,
		HotelImageUrl: req.HotelImageUrl,
		NumberOfRooms: req.NumberOfRooms,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, parseHotelModel(resp))
}

// @Router /hotels/{id} [get]
// @Summary Get hotel by id
// @Description Get hotel by id
// @Tags hotel
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Hotel
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetHotel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.Hotel().Get(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, parseHotelModel(resp))
}

// @Router /hotels [get]
// @Summary Get all hotels
// @Description Get all hotels
// @Tags hotel
// @Accept json
// @Produce json
// @Param filter query models.GetAllParams false "Filter"
// @Success 200 {object} models.GetAllHotelsResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllHotels(c *gin.Context) {
	req, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.Hotel().GetAll(&repo.GetAllHotelsParams{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getHotelsResponse(result))
}

func getHotelsResponse(data *repo.GetAllHotelsResult) *models.GetAllHotelsResponse {
	response := models.GetAllHotelsResponse{
		Hotels: make([]*models.Hotel, 0),
		Count:  data.Count,
	}

	for _, hotel := range data.Hotels {
		u := parseHotelModel(hotel)
		response.Hotels = append(response.Hotels, &u)
	}

	return &response
}

func parseHotelModel(hotel *repo.Hotel) models.Hotel {
	return models.Hotel{
		ID:            hotel.ID,
		UserID:        hotel.UserID,
		HotelName:     hotel.HotelName,
		HotelLocation: hotel.HotelLocation,
		HotelImageUrl: hotel.HotelImageUrl,
		NumberOfRooms: hotel.NumberOfRooms,
		CreatedAt:     hotel.CreatedAt,
	}
}

// @Router /hotels/{id} [put]
// @Summary Update a hotel
// @Description Update a hotel
// @Tags hotel
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param hotel body models.CreateHotelRequest true "Hotel"
// @Success 200 {object} models.Hotel
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateHotel(c *gin.Context) {
	var (
		req models.CreateHotelRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.Hotel().Update(&repo.Hotel{
		ID:            int64(id),
		UserID:        req.UserID,
		HotelName:     req.HotelName,
		HotelLocation: req.HotelLocation,
		HotelImageUrl: req.HotelImageUrl,
		NumberOfRooms: req.NumberOfRooms,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.Hotel{
		ID:            resp.ID,
		UserID:        resp.UserID,
		HotelName:     resp.HotelName,
		HotelLocation: resp.HotelLocation,
		HotelImageUrl: resp.HotelImageUrl,
		NumberOfRooms: resp.NumberOfRooms,
	})
}

// @Router /hotel/{id} [delete]
// @Summary Delete a hotel
// @Description Delete a hotel
// @Tags hotel
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.ResponseOK
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeleteHotel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = h.storage.Hotel().Delete(int64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully deleted",
	})
}
