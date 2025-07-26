package controllers

import (
	"net/http"
	"strconv"

	"sarc/core/domain"
	"sarc/core/services"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	Service services.ProfileService
}

func NewProfileHandler(service services.ProfileService) *ProfileHandler {
	return &ProfileHandler{Service: service}
}

// Create Profile
// @Summary      Create a new profile
// @Description  Creates a new profile in the system
// @Tags         profiles
// @Accept       json
// @Produce      json
// @Param        profile  body      domain.Profile   true  "Profile data"
// @Success      201   {object}  domain.Profile
// @Failure      400   {object}  domain.ErrorResponse "Invalid request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /profiles [post]
func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	var profile domain.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := h.Service.CreateProfile(&profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// Get All Profiles
// @Summary      Get all profiles
// @Description  Retrieves all profiles
// @Tags         profiles
// @Produce      json
// @Success      200   {array}   domain.Profile
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /profiles [get]
func (h *ProfileHandler) GetProfiles(c *gin.Context) {
	profiles, err := h.Service.GetProfiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, profiles)
}

// Get Profile by ID
// @Summary      Get profile by ID
// @Description  Retrieves a profile by its ID
// @Tags         profiles
// @Produce      json
// @Param        id   path      int  true  "Profile ID"
// @Success      200  {object}  domain.Profile
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      404  {object}  domain.ErrorResponse "Profile not found"
// @Router       /profiles/{id} [get]
func (h *ProfileHandler) GetProfileByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	profile, err := h.Service.GetProfileByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, profile)
}

// Update Profile
// @Summary      Update an existing profile
// @Description  Updates the profile information for the given profile ID
// @Tags         profiles
// @Accept       json
// @Produce      json
// @Param        id      path      int             true  "Profile ID"
// @Param        profile body      domain.Profile  true  "Profile data"
// @Success      200   {object}  domain.Profile
// @Failure      400   {object}  domain.ErrorResponse "Invalid ID or bad request"
// @Failure      500   {object}  domain.ErrorResponse "Internal server error"
// @Router       /profiles/{id} [put]
func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var profile domain.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.Service.UpdateProfile(uint(id), &profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Delete Profile
// @Summary      Delete a profile
// @Description  Deletes a profile by its ID
// @Tags         profiles
// @Param        id   path      int  true  "Profile ID"
// @Success      204  {string}  string "No Content"
// @Failure      400  {object}  domain.ErrorResponse "Invalid ID"
// @Failure      500  {object}  domain.ErrorResponse "Internal server error"
// @Router       /profiles/{id} [delete]
func (h *ProfileHandler) DeleteProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.Service.DeleteProfile(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
