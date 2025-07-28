package controllers

import (
	"net/http"
	"strconv"

	"sarc/core/domain"
	"sarc/core/services"

	"github.com/gin-gonic/gin"
)

type DisciplineHandler struct {
	Service services.DisciplineService
}

func NewDisciplineHandler(service services.DisciplineService) *DisciplineHandler {
	return &DisciplineHandler{Service: service}
}

// Create Discipline
// @Summary      Create a new discipline
// @Description  Creates a new discipline in the system
// @Tags         disciplines
// @Accept       json
// @Produce      json
// @Param        discipline  body      domain.Discipline   true  "Discipline data"
// @Success      201   {object}  domain.Discipline
// @Failure      400   {object}  domain.ErrorResponse "Invalid request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /disciplines [post]
func (h *DisciplineHandler) CreateDiscipline(c *gin.Context) {
	var discipline domain.Discipline
	if err := c.ShouldBindJSON(&discipline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := h.Service.CreateDiscipline(&discipline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// Get All Disciplines
// @Summary      Get all disciplines
// @Description  Retrieves all disciplines
// @Tags         disciplines
// @Produce      json
// @Success      200   {array}   domain.Discipline
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /disciplines [get]
func (h *DisciplineHandler) GetDisciplines(c *gin.Context) {
	disciplines, err := h.Service.GetDisciplines()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, disciplines)
}

// Get Discipline by ID
// @Summary      Get discipline by ID
// @Description  Retrieves a discipline by its ID
// @Tags         disciplines
// @Produce      json
// @Param        id   path      int  true  "Discipline ID"
// @Success      200  {object}  domain.Discipline
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      404  {object}  domain.ErrorResponse "Discipline not found"
// @Router       /disciplines/{id} [get]
func (h *DisciplineHandler) GetDisciplineByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	discipline, err := h.Service.GetDisciplineByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, discipline)
}

// Update Discipline
// @Summary      Update an existing discipline
// @Description  Updates the discipline information for the given discipline ID
// @Tags         disciplines
// @Accept       json
// @Produce      json
// @Param        id         path      int                true  "Discipline ID"
// @Param        discipline body      domain.Discipline  true  "Discipline data"
// @Success      200   {object}  domain.Discipline
// @Failure      400   {object}  domain.ErrorResponse "Invalid ID or bad request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /disciplines/{id} [put]
func (h *DisciplineHandler) UpdateDiscipline(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var discipline domain.Discipline
	if err := c.ShouldBindJSON(&discipline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.Service.UpdateDiscipline(uint(id), &discipline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Delete Discipline
// @Summary      Delete a discipline
// @Description  Deletes a discipline by its ID
// @Tags         disciplines
// @Param        id   path      int  true  "Discipline ID"
// @Success      204  {string}  string "No Content"
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      500  {object}  domain.ErrorResponse "Internal server error"
// @Router       /disciplines/{id} [delete]
func (h *DisciplineHandler) DeleteDiscipline(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.Service.DeleteDiscipline(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
