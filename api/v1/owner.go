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

// @Router /owners [post]
// @Summary Create an owner
// @Description Create an owner
// @Tags owner
// @Accept json
// @Produce json
// @Param owner body models.CreateOwnerRequest true "Owner"
// @Success 201 {object} models.Owner
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateOwner(c *gin.Context) {
	var (
		req models.CreateOwnerRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.Owner().Create(&repo.Owner{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Username:    req.Username,
		Password:    req.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, parseOwnerModel(resp))
}

// @Router /owners/{id} [get]
// @Summary Get owner by id
// @Description Get owner by id
// @Tags owner
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Owner
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetOwner(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.Owner().Get(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, parseOwnerModel(resp))
}

// @Router /owners [get]
// @Summary Get all owners
// @Description Get all owners
// @Tags owner
// @Accept json
// @Produce json
// @Param filter query models.GetAllParams false "Filter"
// @Success 200 {object} models.GetAllOwnersResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllOwners(c *gin.Context) {
	req, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.Owner().GetAll(&repo.GetAllOwnersParams{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getOwnersResponse(result))
}

func getOwnersResponse(data *repo.GetAllOwnersResult) *models.GetAllOwnersResponse {
	response := models.GetAllOwnersResponse{
		Owners: make([]*models.Owner, 0),
		Count:  data.Count,
	}

	for _, owner := range data.Owners {
		u := parseOwnerModel(owner)
		response.Owners = append(response.Owners, &u)
	}

	return &response
}

func parseOwnerModel(owner *repo.Owner) models.Owner {
	return models.Owner{
		ID:          owner.ID,
		FirstName:   owner.FirstName,
		LastName:    owner.LastName,
		PhoneNumber: owner.PhoneNumber,
		Email:       owner.Email,
		CreatedAt:   owner.CreatedAt,
	}
}

// @Router /owners/{id} [put]
// @Summary Update an owner
// @Description Update an owner
// @Tags owner
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param owner body models.CreateOwnerRequest true "Owner"
// @Success 200 {object} models.Owner
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateOwner(c *gin.Context) {
	var (
		req models.CreateOwnerRequest
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

	resp, err := h.storage.Owner().Update(&repo.Owner{
		ID:          int64(id),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Username:    req.Username,
		Password:    req.Password,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.Owner{
		ID:          resp.ID,
		FirstName:   resp.FirstName,
		LastName:    resp.LastName,
		Email:       resp.Email,
		PhoneNumber: resp.PhoneNumber,
		Username:    resp.Username,
		Password:    resp.Password,
	})
}

// @Router /owner/{id} [delete]
// @Summary Delete an owner
// @Description Delete an owner
// @Tags owner
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.ResponseOK
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeleteOwner(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = h.storage.Owner().Delete(int64(id))
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
