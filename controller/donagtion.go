package controller

import (
	"donation-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DonationController struct {
	donations []models.Donation
}

func NewController() *DonationController {
	return &DonationController{
		donations: make([]models.Donation, 0),
	}
}

// GetDonations returns all donations
func (dc *DonationController) GetDonations(c *gin.Context) {
	c.JSON(http.StatusOK, dc.donations)
}

// CreateDonation creates a new donation
func (dc *DonationController) CreateDonation(c *gin.Context) {
	var newDonation models.Donation
	if err := c.ShouldBindJSON(&newDonation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	newDonation.ID = uuid.New().String()
	newDonation.CreatedAt = time.Now().Format(time.RFC3339)
	newDonation.Status = "available"

	dc.donations = append(dc.donations, newDonation)
	c.JSON(http.StatusCreated, newDonation)
}

// GetDonation gets a specific donation by ID
func (dc *DonationController) GetDonation(c *gin.Context) {
	id := c.Param("id")

	for _, donation := range dc.donations {
		if donation.ID == id {
			c.JSON(http.StatusOK, donation)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Donation not found"})
}

// UpdateDonation updates an existing donation
func (dc *DonationController) UpdateDonation(c *gin.Context) {
	id := c.Param("id")
	var updatedDonation models.Donation

	if err := c.ShouldBindJSON(&updatedDonation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	for i, donation := range dc.donations {
		if donation.ID == id {
			updatedDonation.ID = id
			updatedDonation.CreatedAt = donation.CreatedAt
			dc.donations[i] = updatedDonation
			c.JSON(http.StatusOK, updatedDonation)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Donation not found"})
}

// DeleteDonation removes a donation
func (dc *DonationController) DeleteDonation(c *gin.Context) {
	id := c.Param("id")

	for i, donation := range dc.donations {
		if donation.ID == id {
			dc.donations = append(dc.donations[:i], dc.donations[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Donation successfully removed"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Donation not found"})
}
