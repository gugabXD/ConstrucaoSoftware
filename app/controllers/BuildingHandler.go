package controllers

import (
	"net/http"
	"strconv"

	"sarc/core/domain"
	serviceinterfaces "sarc/core/services/interfaces"

	"github.com/gin-gonic/gin"
)

type BuildingHandler struct {
	Service serviceinterfaces.BuildingService
}

func NewBuildingHandler(service serviceinterfaces.BuildingService) *BuildingHandler {
	return &BuildingHandler{Service: service}
}

// Create Building
// @Summary      Create a new building
// @Description  Creates a new building in the system
// @Tags         buildings
// @Accept       json
// @Produce      json
// @Param        building  body      domain.Building   true  "Building data"
// @Success      201   {object}  domain.Building
// @Failure      400   {object}  domain.ErrorResponse "Invalid request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /buildings [post]
func (h *BuildingHandler) CreateBuilding(c *gin.Context) {
	var building domain.Building
	if err := c.ShouldBindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := h.Service.CreateBuilding(&building)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// Get All Buildings
// @Summary      Get all buildings
// @Description  Retrieves all buildings
// @Tags         buildings
// @Produce      json
// @Success      200   {array}   domain.Building
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /buildings [get]
func (h *BuildingHandler) GetBuildings(c *gin.Context) {
	buildings, err := h.Service.GetBuildings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, buildings)
}

// Get Building by ID
// @Summary      Get building by ID
// @Description  Retrieves a building by its ID
// @Tags         buildings
// @Produce      json
// @Param        id   path      int  true  "Building ID"
// @Success      200  {object}  domain.Building
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      404  {object}  domain.ErrorResponse "Building not found"
// @Router       /buildings/{id} [get]
func (h *BuildingHandler) GetBuildingByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	building, err := h.Service.GetBuildingByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, building)
}

// Update Building
// @Summary      Update an existing building
// @Description  Updates the building information for the given building ID
// @Tags         buildings
// @Accept       json
// @Produce      json
// @Param        id       path      int             true  "Building ID"
// @Param        building body      domain.Building true  "Building data"
// @Success      200   {object}  domain.Building
// @Failure      400   {object}  domain.ErrorResponse "Invalid ID or bad request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /buildings/{id} [put]
func (h *BuildingHandler) UpdateBuilding(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var building domain.Building
	if err := c.ShouldBindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.Service.UpdateBuilding(uint(id), &building)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Delete Building
// @Summary      Delete a building
// @Description  Deletes a building by its ID
// @Tags         buildings
// @Param        id   path      int  true  "Building ID"
// @Success      204  {string}  string "No Content"
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      500  {object}  domain.ErrorResponse "Internal server error"
// @Router       /buildings/{id} [delete]
func (h *BuildingHandler) DeleteBuilding(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.Service.DeleteBuilding(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
