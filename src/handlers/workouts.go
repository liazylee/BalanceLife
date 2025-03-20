package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhenyili/BalanceLife/src/db"
	"github.com/zhenyili/BalanceLife/src/models"
	"github.com/zhenyili/BalanceLife/src/utils"
)

// WorkoutHandler handles workout-related requests
type WorkoutHandler struct {
	store db.Store
}

// NewWorkoutHandler creates a new workout handler
func NewWorkoutHandler(store db.Store) *WorkoutHandler {
	return &WorkoutHandler{
		store: store,
	}
}

// RegisterRoutes registers workout routes to the router
func (h *WorkoutHandler) RegisterRoutes(router *gin.RouterGroup) {
	workouts := router.Group("/workouts")
	{
		// Workout packages
		workouts.GET("/packages", h.GetWorkoutPackages)
		workouts.GET("/packages/:id", h.GetWorkoutPackage)

		// Workout entries
		workouts.POST("/entries", h.CreateWorkoutEntry)
		workouts.GET("/entries", h.GetWorkoutEntries)
	}
}

// GetWorkoutPackages godoc
// @Summary      Get all workout packages
// @Description  Returns a list of all workout packages, optionally filtered by goal type
// @Tags         workouts
// @Produce      json
// @Param        goalType  query     string  false  "Goal type filter (LOSE, GAIN, ALL)"
// @Success      200       {array}   models.WorkoutPackage
// @Router       /workouts/packages [get]
func (h *WorkoutHandler) GetWorkoutPackages(c *gin.Context) {
	goalType := models.GoalType(c.Query("goalType"))
	packages := h.store.GetWorkoutPackages(goalType)
	c.JSON(http.StatusOK, packages)
}

// GetWorkoutPackage godoc
// @Summary      Get a workout package by ID
// @Description  Returns details of a specific workout package
// @Tags         workouts
// @Produce      json
// @Param        id   path      string  true  "Workout Package ID"
// @Success      200  {object}  models.WorkoutPackage
// @Failure      404  {object}  map[string]string
// @Router       /workouts/packages/{id} [get]
func (h *WorkoutHandler) GetWorkoutPackage(c *gin.Context) {
	id := c.Param("id")
	pkg, err := h.store.GetWorkoutPackage(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pkg)
}

// workoutEntryRequest defines the structure for workout entry creation
type workoutEntryRequest struct {
	UserID              string  `json:"userId" binding:"required" example:"usr1"`
	PackageID           string  `json:"packageId" binding:"required" example:"workout1"`
	IntensityMultiplier float64 `json:"intensityMultiplier" binding:"required,min=0.5,max=2" example:"1.0"`
	DurationMinutes     int     `json:"durationMinutes" binding:"required,min=5,max=180" example:"30"`
	Date                string  `json:"date" binding:"required" example:"2023-03-18"`
}

// CreateWorkoutEntry godoc
// @Summary      Create a new workout entry
// @Description  Logs a workout for a user with specified intensity and duration
// @Tags         workouts
// @Accept       json
// @Produce      json
// @Param        entry  body      workoutEntryRequest  true  "Workout entry details"
// @Success      201    {object}  models.WorkoutEntry
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /workouts/entries [post]
func (h *WorkoutHandler) CreateWorkoutEntry(c *gin.Context) {
	var req workoutEntryRequest
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

	// Get workout package
	pkg, err := h.store.GetWorkoutPackage(req.PackageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid workout package: " + err.Error()})
		return
	}

	// Get user information for calorie calculation
	user, err := h.store.GetUser(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user: " + err.Error()})
		return
	}

	// Calculate calories burned based on workout package, duration, intensity, and user weight
	// Note: In a real app, you would use the actual formula from the package
	// For MVP, we'll use a simple formula
	caloriesBurned := int(float64(pkg.BaseCaloriesBurn) *
		(float64(req.DurationMinutes) / float64(pkg.BaseDurationMinutes)) *
		req.IntensityMultiplier *
		(user.Weight / 70.0)) // Adjust for user weight relative to 70kg reference

	newEntry := models.WorkoutEntry{
		ID:                  utils.GenerateID(),
		UserID:              req.UserID,
		PackageID:           req.PackageID,
		IntensityMultiplier: req.IntensityMultiplier,
		DurationMinutes:     req.DurationMinutes,
		CaloriesBurned:      caloriesBurned,
		Date:                date,
		Timestamp:           time.Now(),
		CreatedAt:           time.Now(),
	}

	// Save the entry
	createdEntry, err := h.store.CreateWorkoutEntry(newEntry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save workout entry: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdEntry)
}

// GetWorkoutEntries godoc
// @Summary      Get workout entries for a user
// @Description  Returns workout entries for a user within a date range
// @Tags         workouts
// @Produce      json
// @Param        userId     query     string  true   "User ID"
// @Param        startDate  query     string  false  "Start date (YYYY-MM-DD)"
// @Param        endDate    query     string  false  "End date (YYYY-MM-DD)"
// @Success      200        {array}   models.WorkoutEntry
// @Failure      400        {object}  map[string]string
// @Router       /workouts/entries [get]
func (h *WorkoutHandler) GetWorkoutEntries(c *gin.Context) {
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

	entries := h.store.GetWorkoutEntriesByUserAndDateRange(userID, startDate, endDate)
	c.JSON(http.StatusOK, entries)
}
