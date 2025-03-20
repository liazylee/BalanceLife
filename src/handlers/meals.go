package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhenyili/BalanceLife/src/db"
	"github.com/zhenyili/BalanceLife/src/models"
	"github.com/zhenyili/BalanceLife/src/utils"
)

// MealHandler handles meal-related requests
type MealHandler struct {
	store db.Store
}

// NewMealHandler creates a new meal handler
func NewMealHandler(store db.Store) *MealHandler {
	return &MealHandler{
		store: store,
	}
}

// RegisterRoutes registers meal routes to the router
func (h *MealHandler) RegisterRoutes(router *gin.RouterGroup) {
	meals := router.Group("/meals")
	{
		// Meal packages
		meals.GET("/packages", h.GetMealPackages)
		meals.GET("/packages/:id", h.GetMealPackage)

		// Meal entries
		meals.POST("/entries", h.CreateMealEntry)
		meals.GET("/entries", h.GetMealEntries)
	}
}

// GetMealPackages godoc
// @Summary      Get all meal packages
// @Description  Returns a list of all meal packages, optionally filtered by goal type
// @Tags         meals
// @Produce      json
// @Param        goalType  query     string  false  "Goal type filter (LOSE, GAIN, ALL)"
// @Success      200       {array}   models.MealPackage
// @Router       /meals/packages [get]
func (h *MealHandler) GetMealPackages(c *gin.Context) {
	goalType := models.GoalType(c.Query("goalType"))
	packages := h.store.GetMealPackages(goalType)
	c.JSON(http.StatusOK, packages)
}

// GetMealPackage godoc
// @Summary      Get a meal package by ID
// @Description  Returns details of a specific meal package
// @Tags         meals
// @Produce      json
// @Param        id   path      string  true  "Meal Package ID"
// @Success      200  {object}  models.MealPackage
// @Failure      404  {object}  map[string]string
// @Router       /meals/packages/{id} [get]
func (h *MealHandler) GetMealPackage(c *gin.Context) {
	id := c.Param("id")
	pkg, err := h.store.GetMealPackage(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkg)
}

// mealEntryRequest defines the structure for meal entry creation
type mealEntryRequest struct {
	UserID            string  `json:"userId" binding:"required" example:"usr1"`
	PackageID         string  `json:"packageId" binding:"required" example:"meal1"`
	PortionMultiplier float64 `json:"portionMultiplier" binding:"required,min=0.1,max=3" example:"1.0"`
	Date              string  `json:"date" binding:"required" example:"2023-03-18"`
}

// CreateMealEntry godoc
// @Summary      Create a new meal entry
// @Description  Logs a meal for a user with specified portion size
// @Tags         meals
// @Accept       json
// @Produce      json
// @Param        entry  body      mealEntryRequest  true  "Meal entry details"
// @Success      201    {object}  models.MealEntry
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /meals/entries [post]
func (h *MealHandler) CreateMealEntry(c *gin.Context) {
	var req mealEntryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse date
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use YYYY-MM-DD"})
		return
	}

	// Get meal package to calculate nutritional values
	pkg, err := h.store.GetMealPackage(req.PackageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meal package: " + err.Error()})
		return
	}

	// Calculate nutritional values based on portion size
	newEntry := models.MealEntry{
		ID:                utils.GenerateID(),
		UserID:            req.UserID,
		PackageID:         req.PackageID,
		PortionMultiplier: req.PortionMultiplier,
		Calories:          int(float64(pkg.BaseCalories) * req.PortionMultiplier),
		Protein:           int(float64(pkg.BaseProtein) * req.PortionMultiplier),
		Carbs:             int(float64(pkg.BaseCarbs) * req.PortionMultiplier),
		Fat:               int(float64(pkg.BaseFat) * req.PortionMultiplier),
		MealType:          pkg.MealType,
		Date:              date,
		Timestamp:         time.Now(),
		CreatedAt:         time.Now(),
	}

	// Save the entry
	createdEntry, err := h.store.CreateMealEntry(newEntry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save meal entry: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdEntry)
}

// GetMealEntries godoc
// @Summary      Get meal entries for a user
// @Description  Returns meal entries for a user within a date range
// @Tags         meals
// @Produce      json
// @Param        userId     query     string  true   "User ID"
// @Param        startDate  query     string  false  "Start date (YYYY-MM-DD)"
// @Param        endDate    query     string  false  "End date (YYYY-MM-DD)"
// @Success      200        {array}   models.MealEntry
// @Failure      400        {object}  map[string]string
// @Router       /meals/entries [get]
func (h *MealHandler) GetMealEntries(c *gin.Context) {
	userID := c.Query("userId")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}

	// Parse start and end dates
	startDateStr := c.DefaultQuery("startDate", time.Now().Format("2006-01-02"))
	endDateStr := c.DefaultQuery("endDate", time.Now().Format("2006-01-02"))

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid startDate format, use YYYY-MM-DD"})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid endDate format, use YYYY-MM-DD"})
		return
	}

	// Make sure the end date is inclusive by setting it to the end of the day
	endDate = endDate.Add(24*time.Hour - time.Second)

	entries := h.store.GetMealEntriesByUserAndDateRange(userID, startDate, endDate)
	c.JSON(http.StatusOK, entries)
}
