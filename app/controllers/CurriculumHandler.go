package controllers

import (
	"net/http"
	"strconv"

	"sarc/core/domain"
	"sarc/core/services"

	"github.com/gin-gonic/gin"
)

type CurriculumHandler struct {
	Service services.CurriculumService
}

func NewCurriculumHandler(service services.CurriculumService) *CurriculumHandler {
	return &CurriculumHandler{Service: service}
}

// Create Curriculum
// @Summary      Create a new curriculum
// @Description  Creates a new curriculum in the system
// @Tags         curriculums
// @Accept       json
// @Produce      json
// @Param        curriculum  body      domain.Curriculum   true  "Curriculum data"
// @Success      201   {object}  domain.Curriculum
// @Failure      400   {object}  domain.ErrorResponse "Invalid request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /curriculums [post]
func (h *CurriculumHandler) CreateCurriculum(c *gin.Context) {
	var curriculum domain.Curriculum
	if err := c.ShouldBindJSON(&curriculum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := h.Service.CreateCurriculum(&curriculum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// Get All Curriculums
// @Summary      Get all curriculums
// @Description  Retrieves all curriculums
// @Tags         curriculums
// @Produce      json
// @Success      200   {array}   domain.Curriculum
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /curriculums [get]
func (h *CurriculumHandler) GetCurriculums(c *gin.Context) {
	curriculums, err := h.Service.GetCurriculums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, curriculums)
}

// Get Curriculum by ID
// @Summary      Get curriculum by ID
// @Description  Retrieves a curriculum by its ID
// @Tags         curriculums
// @Produce      json
// @Param        id   path      int  true  "Curriculum ID"
// @Success      200  {object}  domain.Curriculum
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      404  {object}  domain.ErrorResponse "Curriculum not found"
// @Router       /curriculums/{id} [get]
func (h *CurriculumHandler) GetCurriculumByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	curriculum, err := h.Service.GetCurriculumByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, curriculum)
}

// Update Curriculum
// @Summary      Update an existing curriculum
// @Description  Updates the curriculum information for the given curriculum ID
// @Tags         curriculums
// @Accept       json
// @Produce      json
// @Param        id         path      int                true  "Curriculum ID"
// @Param        curriculum body      domain.Curriculum  true  "Curriculum data"
// @Success      200   {object}  domain.Curriculum
// @Failure      400   {object}  domain.ErrorResponse "Invalid ID or bad request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /curriculums/{id} [put]
func (h *CurriculumHandler) UpdateCurriculum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var curriculum domain.Curriculum
	if err := c.ShouldBindJSON(&curriculum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.Service.UpdateCurriculum(uint(id), &curriculum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Delete Curriculum
// @Summary      Delete a curriculum
// @Description  Deletes a curriculum by its ID
// @Tags         curriculums
// @Param        id   path      int  true  "Curriculum ID"
// @Success      204  {string}  string "No Content"
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      500  {object}  domain.ErrorResponse "Internal server error"
// @Router       /curriculums/{id} [delete]
func (h *CurriculumHandler) DeleteCurriculum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.Service.DeleteCurriculum(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// Add Discipline to Curriculum
// @Summary      Add a discipline to a curriculum
// @Description  Associates a discipline with a curriculum (many-to-many relation)
// @Tags         curriculums
// @Accept       json
// @Produce      json
// @Param        id           path      int     true  "Curriculum ID"
// @Param        discipline   body      object  true  "Discipline ID to add"  Schema({"disciplineId":1})
// @Success      204  {string}  string "No Content"
// @Failure      400  {object}  domain.ErrorResponse "Invalid curriculum ID or bad request"
// @Failure      500  {object}  domain.ErrorResponse "Internal server error"
// @Router       /curriculums/{id}/disciplines [post]
func (h *CurriculumHandler) AddDisciplineToCurriculum(c *gin.Context) {
	curriculumID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid curriculum ID"})
		return
	}
	var req struct {
		DisciplineID uint `json:"disciplineId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = h.Service.AddDisciplineToCurriculum(uint(curriculumID), req.DisciplineID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}
