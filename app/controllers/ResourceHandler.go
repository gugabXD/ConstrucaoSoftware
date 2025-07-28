package controllers

import (
	"net/http"
	"strconv"

	"sarc/core/domain"
	serviceinterfaces "sarc/core/services/interfaces"

	"github.com/gin-gonic/gin"
)

type ResourceHandler struct {
	Service serviceinterfaces.ResourceService
}

func NewResourceHandler(service serviceinterfaces.ResourceService) *ResourceHandler {
	return &ResourceHandler{Service: service}
}

// Create Resource
// @Summary      Create a new resource
// @Description  Creates a new resource in the system
// @Tags         resources
// @Accept       json
// @Produce      json
// @Param        resource  body      domain.Resource   true  "Resource data"
// @Success      201   {object}  domain.Resource
// @Failure      400   {object}  domain.ErrorResponse "Invalid request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /resources [post]
func (h *ResourceHandler) CreateResource(c *gin.Context) {
	var resource domain.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := h.Service.CreateResource(&resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// Get All Resources
// @Summary      Get all resources
// @Description  Retrieves all resources
// @Tags         resources
// @Produce      json
// @Success      200   {array}   domain.Resource
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /resources [get]
func (h *ResourceHandler) GetResources(c *gin.Context) {
	resources, err := h.Service.GetResources()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resources)
}

// Get Resource by ID
// @Summary      Get resource by ID
// @Description  Retrieves a resource by its ID
// @Tags         resources
// @Produce      json
// @Param        id   path      int  true  "Resource ID"
// @Success      200  {object}  domain.Resource
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      404  {object}  domain.ErrorResponse "Resource not found"
// @Router       /resources/{id} [get]
func (h *ResourceHandler) GetResourceByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	resource, err := h.Service.GetResourceByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resource)
}

// Update Resource
// @Summary      Update an existing resource
// @Description  Updates the resource information for the given resource ID
// @Tags         resources
// @Accept       json
// @Produce      json
// @Param        id        path      int              true  "Resource ID"
// @Param        resource  body      domain.Resource  true  "Resource data"
// @Success      200   {object}  domain.Resource
// @Failure      400   {object}  domain.ErrorResponse "Invalid ID or bad request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /resources/{id} [put]
func (h *ResourceHandler) UpdateResource(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var resource domain.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.Service.UpdateResource(uint(id), &resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Delete Resource
// @Summary      Delete a resource
// @Description  Deletes a resource by its ID
// @Tags         resources
// @Param        id   path      int  true  "Resource ID"
// @Success      204  {string}  string "No Content"
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      500  {object}  domain.ErrorResponse "Internal server error"
// @Router       /resources/{id} [delete]
func (h *ResourceHandler) DeleteResource(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.Service.DeleteResource(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
