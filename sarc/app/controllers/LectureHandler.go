package controllers

import (
	"net/http"
	"strconv"

	"sarc/core/domain"
	"sarc/core/services"

	"github.com/gin-gonic/gin"
)

type LectureHandler struct {
	Service services.LectureService
}

func NewLectureHandler(service services.LectureService) *LectureHandler {
	return &LectureHandler{Service: service}
}

// Create Lecture
// @Summary      Create a new lecture
// @Description  Creates a new lecture in the system
// @Tags         lectures
// @Accept       json
// @Produce      json
// @Param        lecture  body      domain.Lecture   true  "Lecture data"
// @Success      201   {object}  domain.Lecture
// @Failure      400   {object}  domain.ErrorResponse "Invalid request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /lectures [post]
func (h *LectureHandler) CreateLecture(c *gin.Context) {
	var lecture domain.Lecture
	if err := c.ShouldBindJSON(&lecture); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := h.Service.CreateLecture(&lecture)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// Get All Lectures
// @Summary      Get all lectures
// @Description  Retrieves all lectures
// @Tags         lectures
// @Produce      json
// @Success      200   {array}   domain.Lecture
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /lectures [get]
func (h *LectureHandler) GetLectures(c *gin.Context) {
	lectures, err := h.Service.GetLectures()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lectures)
}

// Get Lecture by ID
// @Summary      Get lecture by ID
// @Description  Retrieves a lecture by its ID
// @Tags         lectures
// @Produce      json
// @Param        id   path      int  true  "Lecture ID"
// @Success      200  {object}  domain.Lecture
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      404  {object}  domain.ErrorResponse "Lecture not found"
// @Router       /lectures/{id} [get]
func (h *LectureHandler) GetLectureByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	lecture, err := h.Service.GetLectureByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lecture)
}

// Update Lecture
// @Summary      Update an existing lecture
// @Description  Updates the lecture information for the given lecture ID
// @Tags         lectures
// @Accept       json
// @Produce      json
// @Param        id      path      int             true  "Lecture ID"
// @Param        lecture body      domain.Lecture  true  "Lecture data"
// @Success      200   {object}  domain.Lecture
// @Failure      400   {object}  domain.ErrorResponse "Invalid ID or bad request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /lectures/{id} [put]
func (h *LectureHandler) UpdateLecture(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var lecture domain.Lecture
	if err := c.ShouldBindJSON(&lecture); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.Service.UpdateLecture(uint(id), &lecture)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Delete Lecture
// @Summary      Delete a lecture
// @Description  Deletes a lecture by its ID
// @Tags         lectures
// @Param        id   path      int  true  "Lecture ID"
// @Success      204  {string}  string "No Content"
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      500  {object}  domain.ErrorResponse "Internal server error"
// @Router       /lectures/{id} [delete]
func (h *LectureHandler) DeleteLecture(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.Service.DeleteLecture(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
