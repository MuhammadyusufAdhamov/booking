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

// @Router /rooms [post]
// @Summary Create a room
// @Description Create a room
// @Tags room
// @Accept json
// @Produce json
// @Param room body models.CreateRoomRequest true "Room"
// @Success 201 {object} models.Room
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateRoom(c *gin.Context) {
	var (
		req models.CreateRoomRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.Room().Create(&repo.Room{
		Type:         req.Type,
		NumberOfRoom: req.NumberOfRoom,
		Sleeps:       req.Sleeps,
		RoomImageUrl: req.RoomImageUrl,
		Price:        req.Price,
		Status:       req.Status,
		HotelId:      req.HotelId,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, parseRoomModel(resp))
}

// @Router /rooms/{id} [get]
// @Summary Get room by id
// @Description Get room by id
// @Tags room
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Room
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetRoom(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.Room().Get(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, parseRoomModel(resp))
}

// @Router /rooms [get]
// @Summary Get all rooms
// @Description Get all rooms
// @Tags room
// @Accept json
// @Produce json
// @Param filter query models.GetAllParams false "Filter"
// @Success 200 {object} models.GetAllRoomsResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllRooms(c *gin.Context) {
	req, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.Room().GetAll(&repo.GetAllRoomsParams{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getRoomsResponse(result))
}

func getRoomsResponse(data *repo.GetAllRoomsResult) *models.GetAllRoomsResponse {
	response := models.GetAllRoomsResponse{
		Rooms: make([]*models.Room, 0),
		Count: data.Count,
	}

	for _, room := range data.Rooms {
		u := parseRoomModel(room)
		response.Rooms = append(response.Rooms, &u)
	}

	return &response
}

func parseRoomModel(room *repo.Room) models.Room {
	return models.Room{
		ID:           room.ID,
		Type:         room.Type,
		NumberOfRoom: room.NumberOfRoom,
		Sleeps:       room.Sleeps,
		RoomImageUrl: room.RoomImageUrl,
		Price:        room.Price,
		Status:       room.Status,
		HotelId:      room.HotelId,
		CreatedAt:    room.CreatedAt,
	}
}

// @Router /rooms/{id} [put]
// @Summary Update a room
// @Description Update a room
// @Tags room
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param room body models.CreateRoomRequest true "Room"
// @Success 200 {object} models.Room
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateRoom(c *gin.Context) {
	var (
		req models.CreateRoomRequest
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

	resp, err := h.storage.Room().Update(&repo.Room{
		ID:           int64(id),
		Type:         req.Type,
		NumberOfRoom: req.NumberOfRoom,
		Sleeps:       req.Sleeps,
		RoomImageUrl: req.RoomImageUrl,
		Price:        req.Price,
		Status:       req.Status,
		HotelId:      req.HotelId,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.Room{
		ID:           resp.ID,
		Type:         resp.Type,
		NumberOfRoom: resp.NumberOfRoom,
		Sleeps:       resp.Sleeps,
		RoomImageUrl: resp.RoomImageUrl,
		Price:        resp.Price,
		Status:       resp.Status,
		HotelId:      resp.HotelId,
	})
}

// @Router /room/{id} [delete]
// @Summary Delete a room
// @Description Delete a room
// @Tags room
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.ResponseOK
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeleteRoom(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = h.storage.Room().Delete(int64(id))
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
