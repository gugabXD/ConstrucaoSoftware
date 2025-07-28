package controllers

import (
	"net/http"
	"strconv"

	"sarc/core/domain"
	serviceinterfaces "sarc/core/services/interfaces"

	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
	Service serviceinterfaces.RoomService
}

func NewRoomHandler(service serviceinterfaces.RoomService) *RoomHandler {
	return &RoomHandler{Service: service}
}

// Create Room
// @Summary      Create a new room
// @Description  Creates a new room in the system
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param        room  body      domain.Room   true  "Room data"
// @Success      201   {object}  domain.Room
// @Failure      400   {object}  domain.ErrorResponse "Invalid request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /rooms [post]
func (h *RoomHandler) CreateRoom(c *gin.Context) {
	var room domain.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := h.Service.CreateRoom(&room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// Get All Rooms
// @Summary      Get all rooms
// @Description  Retrieves all rooms
// @Tags         rooms
// @Produce      json
// @Success      200   {array}   domain.Room
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /rooms [get]
func (h *RoomHandler) GetRooms(c *gin.Context) {
	rooms, err := h.Service.GetRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rooms)
}

// Get Room by ID
// @Summary      Get room by ID
// @Description  Retrieves a room by its ID
// @Tags         rooms
// @Produce      json
// @Param        id   path      int  true  "Room ID"
// @Success      200  {object}  domain.Room
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      404  {object}  domain.ErrorResponse "Room not found"
// @Router       /rooms/{id} [get]
func (h *RoomHandler) GetRoomByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	room, err := h.Service.GetRoomByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, room)
}

// Update Room
// @Summary      Update an existing room
// @Description  Updates the room information for the given room ID
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param        id    path      int           true  "Room ID"
// @Param        room  body      domain.Room   true  "Room data"
// @Success      200   {object}  domain.Room
// @Failure      400   {object}  domain.ErrorResponse "Invalid ID or bad request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /rooms/{id} [put]
func (h *RoomHandler) UpdateRoom(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var room domain.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.Service.UpdateRoom(uint(id), &room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Delete Room
// @Summary      Delete a room
// @Description  Deletes a room by its ID
// @Tags         rooms
// @Param        id   path      int  true  "Room ID"
// @Success      204  {string}  string "No Content"
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      500  {object}  domain.ErrorResponse "Internal server error"
// @Router       /rooms/{id} [delete]
func (h *RoomHandler) DeleteRoom(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.Service.DeleteRoom(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
