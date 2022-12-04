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

// @Router /bookings [post]
// @Summary Create a booking
// @Description Create a booking
// @Tags booking
// @Accept json
// @Produce json
// @Param booking body models.CreateBookingRequest true "Booking"
// @Success 201 {object} models.Booking
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateBooking(c *gin.Context) {
	var (
		req models.CreateBookingRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.Booking().Create(&repo.Booking{
		RoomId:        req.RoomId,
		UserId:        req.UserId,
		Stay:          req.Stay,
		NumberOfUsers: req.NumberOfUsers,
		StayDate:      req.StayDate,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, parseBookingModel(resp))
}

// @Router /bookings/{id} [get]
// @Summary Get booking by id
// @Description Get booking by id
// @Tags booking
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Booking
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetBooking(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.Booking().Get(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, parseBookingModel(resp))
}

// @Router /bookings [get]
// @Summary Get all bookings
// @Description Get all bookings
// @Tags booking
// @Accept json
// @Produce json
// @Param filter query models.GetAllParams false "Filter"
// @Success 200 {object} models.GetAllBookingsResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllBookings(c *gin.Context) {
	req, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.Booking().GetAll(&repo.GetAllBookingsParams{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getBookingsResponse(result))
}

func getBookingsResponse(data *repo.GetAllBookingResult) *models.GetAllBookingsResponse {
	response := models.GetAllBookingsResponse{
		Bookings: make([]*models.Booking, 0),
		Count:    data.Count,
	}

	for _, booking := range data.Bookings {
		u := parseBookingModel(booking)
		response.Bookings = append(response.Bookings, &u)
	}

	return &response
}

func parseBookingModel(booking *repo.Booking) models.Booking {
	return models.Booking{
		ID:            booking.ID,
		RoomId:        booking.RoomId,
		UserId:        booking.UserId,
		Stay:          booking.Stay,
		NumberOfUsers: booking.NumberOfUsers,
		StayDate:      booking.StayDate,
		CreatedAt:     booking.CreatedAt,
	}
}

// @Router /bookings/{id} [put]
// @Summary Update a booking
// @Description Update a booking
// @Tags booking
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param booking body models.CreateBookingRequest true "Booking"
// @Success 200 {object} models.Booking
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateBooking(c *gin.Context) {
	var (
		req models.CreateBookingRequest
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

	resp, err := h.storage.Booking().Update(&repo.Booking{
		ID:            int64(id),
		RoomId:        req.RoomId,
		UserId:        req.UserId,
		Stay:          req.Stay,
		NumberOfUsers: req.NumberOfUsers,
		StayDate:      req.StayDate,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.Booking{
		ID:            resp.ID,
		RoomId:        resp.RoomId,
		UserId:        resp.UserId,
		Stay:          resp.Stay,
		NumberOfUsers: resp.NumberOfUsers,
		StayDate:      resp.StayDate,
	})
}

// @Router /booking/{id} [delete]
// @Summary Delete a booking
// @Description Delete a booking
// @Tags booking
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.ResponseOK
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeleteBooking(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = h.storage.Booking().Delete(int64(id))
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
