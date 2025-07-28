package controllers

import (
	"net/http"
	"strconv"

	"sarc/core/domain"
	"sarc/core/services"

	"github.com/gin-gonic/gin"
)

type ClassHandler struct {
	Service services.ClassService
}

func NewClassHandler(service services.ClassService) *ClassHandler {
	return &ClassHandler{Service: service}
}

// Create Class
// @Summary      Create a new class
// @Description  Creates a new class in the system
// @Tags         classes
// @Accept       json
// @Produce      json
// @Param        class  body      domain.Class   true  "Class data"
// @Success      201   {object}  domain.Class
// @Failure      400   {object}  domain.ErrorResponse "Invalid request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /classes [post]
func (h *ClassHandler) CreateClass(c *gin.Context) {
	var class domain.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := h.Service.CreateClass(&class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// Get All Classes
// @Summary      Get all classes
// @Description  Retrieves all classes
// @Tags         classes
// @Produce      json
// @Success      200   {array}   domain.Class
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /classes [get]
func (h *ClassHandler) GetClasses(c *gin.Context) {
	classes, err := h.Service.GetClasses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, classes)
}

// Get Class by ID
// @Summary      Get class by ID
// @Description  Retrieves a class by its ID
// @Tags         classes
// @Produce      json
// @Param        id   path      int  true  "Class ID"
// @Success      200  {object}  domain.Class
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      404  {object} domain.ErrorResponse "Class not found"
// @Router       /classes/{id} [get]
func (h *ClassHandler) GetClassByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	class, err := h.Service.GetClassByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, class)
}

// Update Class
// @Summary      Update an existing class
// @Description  Updates the class information for the given class ID
// @Tags         classes
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Class ID"
// @Param        class body      domain.Class true "Class data"
// @Success      200   {object}  domain.Class
// @Failure      400   {object}  domain.ErrorResponse "Invalid ID or bad request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /classes/{id} [put]
func (h *ClassHandler) UpdateClass(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var class domain.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.Service.UpdateClass(uint(id), &class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Delete Class
// @Summary      Delete a class
// @Description  Deletes a class by its ID
// @Tags         classes
// @Param        id   path      int  true  "Class ID"
// @Success      204  {string}  string "No Content"
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      500  {object}  domain.ErrorResponse "Internal server error"
// @Router       /classes/{id} [delete]
func (h *ClassHandler) DeleteClass(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.Service.DeleteClass(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
