package controllers

import (
	"net/http"
	"strconv"

	"sarc/core/domain"
	"sarc/core/services"

	"github.com/gin-gonic/gin"
)

type ReservationsHandler struct {
	Service services.ReservationsService
}

func NewReservationsHandler(service services.ReservationsService) *ReservationsHandler {
	return &ReservationsHandler{Service: service}
}

// Create Reservation
// @Summary      Create a new reservation
// @Description  Creates a new reservation in the system
// @Tags         reservations
// @Accept       json
// @Produce      json
// @Param        reservation  body      domain.Reservation   true  "Reservation data"
// @Success      201   {object}  domain.Reservation
// @Failure      400   {object}  domain.ErrorResponse "Invalid request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /reservations [post]
func (h *ReservationsHandler) CreateReservation(c *gin.Context) {
	var reservation domain.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := h.Service.CreateReservation(&reservation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// Get All Reservations
// @Summary      Get all reservations
// @Description  Retrieves all reservations
// @Tags         reservations
// @Produce      json
// @Success      200   {array}   domain.Reservation
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /reservations [get]
func (h *ReservationsHandler) GetReservations(c *gin.Context) {
	reservations, err := h.Service.GetReservations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reservations)
}

// Get Reservation by ID
// @Summary      Get reservation by ID
// @Description  Retrieves a reservation by its ID
// @Tags         reservations
// @Produce      json
// @Param        id   path      int  true  "Reservation ID"
// @Success      200  {object}  domain.Reservation
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      404  {object}  domain.ErrorResponse "Reservation not found"
// @Router       /reservations/{id} [get]
func (h *ReservationsHandler) GetReservationByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	reservation, err := h.Service.GetReservationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reservation)
}

// Update Reservation
// @Summary      Update an existing reservation
// @Description  Updates the reservation information for the given reservation ID
// @Tags         reservations
// @Accept       json
// @Produce      json
// @Param        id           path      int                true  "Reservation ID"
// @Param        reservation  body      domain.Reservation true  "Reservation data"
// @Success      200   {object}  domain.Reservation
// @Failure      400   {object}  domain.ErrorResponse "Invalid ID or bad request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /reservations/{id} [put]
func (h *ReservationsHandler) UpdateReservation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var reservation domain.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.Service.UpdateReservation(uint(id), &reservation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Delete Reservation
// @Summary      Delete a reservation
// @Description  Deletes a reservation by its ID
// @Tags         reservations
// @Param        id   path      int  true  "Reservation ID"
// @Success      204  {string}  string "No Content"
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      500  {object}  domain.ErrorResponse "Internal server error"
// @Router       /reservations/{id} [delete]
func (h *ReservationsHandler) DeleteReservation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.Service.DeleteReservation(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// Add Resource to Reservation
// @Summary      Add a resource to a reservation
// @Description  Associates a resource with a reservation (many-to-many relation)
// @Tags         reservations
// @Accept       json
// @Produce      json
// @Param        id         path      int  true  "Reservation ID"
// @Param        resource   body      object  true  "Resource ID to add"  Schema({"resourceId":1})
// @Success      204  {string}  string "No Content"
// @Failure      400  {object}  domain.ErrorResponse "Invalid reservation ID or bad request"
// @Failure      500  {object}  domain.ErrorResponse "Internal server error"
// @Router       /reservations/{id}/resources [post]
func (h *ReservationsHandler) AddResourceToReservation(c *gin.Context) {
	reservationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid reservation ID"})
		return
	}
	var req struct {
		ResourceID uint `json:"resourceId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = h.Service.AddResourceToReservation(uint(reservationID), req.ResourceID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}
