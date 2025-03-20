package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhenyili/BalanceLife/src/db"
	"github.com/zhenyili/BalanceLife/src/models"
	"github.com/zhenyili/BalanceLife/src/utils"
)

// UserHandler handles user-related requests
type UserHandler struct {
	store db.Store
}

// NewUserHandler creates a new user handler
func NewUserHandler(store db.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

// RegisterRoutes registers user routes to the router
func (h *UserHandler) RegisterRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.GET("", h.GetUsers)
		users.GET("/:id", h.GetUser)
		users.POST("", h.CreateUser)
		users.DELETE("/:id", h.DeleteUser)
	}
}

// GetUsers godoc
// @Summary      Get all users
// @Description  Returns a list of all users in the system
// @Tags         users
// @Produce      json
// @Success      200  {array}   models.User
// @Router       /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users := h.store.GetUsers()
	c.JSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary      Get a user by ID
// @Description  Returns details of a specific user
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  models.User
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.store.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// userRegistrationRequest defines the structure for user registration
type userRegistrationRequest struct {
	Name          string  `json:"name" binding:"required" example:"John Doe"`
	Email         string  `json:"email" binding:"required,email" example:"john@example.com"`
	Password      string  `json:"password" binding:"required,min=6" example:"SecurePassword123"`
	Gender        string  `json:"gender" binding:"required" example:"MALE" enums:"MALE,FEMALE,OTHER"`
	BirthDate     string  `json:"birthDate" binding:"required" example:"1990-01-01"`
	Height        float64 `json:"height" binding:"required" example:"180.0"`
	Weight        float64 `json:"weight" binding:"required" example:"80.0"`
	ActivityLevel string  `json:"activityLevel" binding:"required" example:"MODERATE" enums:"LOW,MODERATE,HIGH"`
	Goal          string  `json:"goal" binding:"required" example:"LOSE" enums:"LOSE,GAIN"`
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Creates a new user with the provided details and calculates their nutritional goals
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      userRegistrationRequest  true  "User details"
// @Success      201   {object}  models.User
// @Failure      400   {object}  map[string]string
// @Router       /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req userRegistrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse birth date
	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid birth date format, use YYYY-MM-DD"})
		return
	}

	// Validate gender
	gender := models.Gender(req.Gender)
	if gender != models.GenderMale && gender != models.GenderFemale && gender != models.GenderOther {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gender"})
		return
	}

	// Validate activity level
	activityLevel := models.ActivityLevel(req.ActivityLevel)
	if activityLevel != models.ActivityLow && activityLevel != models.ActivityModerate && activityLevel != models.ActivityHigh {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity level"})
		return
	}

	// Validate goal
	goalType := models.GoalType(req.Goal)
	if goalType != models.GoalTypeLose && goalType != models.GoalTypeGain {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid goal type"})
		return
	}

	// Calculate default target values based on user's data
	// Note: In a real application, this would be more sophisticated
	baseCalories := calculateBaseCalories(req.Weight, req.Height, birthDate, gender, activityLevel)

	var targetCalories int
	if goalType == models.GoalTypeLose {
		targetCalories = baseCalories - 500 // 500 calorie deficit for weight loss
	} else {
		targetCalories = baseCalories + 500 // 500 calorie surplus for muscle gain
	}

	// Create a new user
	newUser := models.User{
		ID:            utils.GenerateID(), // Simple ID generation
		Name:          req.Name,
		Email:         req.Email,
		Password:      req.Password, // In a real app, you would hash this
		Gender:        gender,
		BirthDate:     birthDate,
		Height:        req.Height,
		Weight:        req.Weight,
		ActivityLevel: activityLevel,
		Goal: models.GoalInfo{
			Type:           goalType,
			TargetCalories: targetCalories,
			TargetProtein:  int(float64(targetCalories) * 0.25 / 4), // 25% from protein (4 calories per gram)
			TargetCarbs:    int(float64(targetCalories) * 0.5 / 4),  // 50% from carbs (4 calories per gram)
			TargetFat:      int(float64(targetCalories) * 0.25 / 9), // 25% from fat (9 calories per gram)
			StartDate:      time.Now(),
			StartWeight:    req.Weight,
		},
		CreatedAt:   time.Now(),
		LastLoginAt: time.Now(),
	}

	createdUser, err := h.store.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Deletes a user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  models.User
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Call the store to delete the user
	deletedUser, err := h.store.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deletedUser)
}

// calculateBaseCalories calculates base calorie needs using a simplified formula
// In a real application, you would use more sophisticated formulas like Harris-Benedict
func calculateBaseCalories(weight float64, height float64, birthDate time.Time, gender models.Gender, activityLevel models.ActivityLevel) int {
	// Calculate age
	now := time.Now()
	age := now.Year() - birthDate.Year()
	if now.YearDay() < birthDate.YearDay() {
		age--
	}

	var bmr float64
	if gender == models.GenderMale {
		bmr = 88.362 + (13.397 * weight) + (4.799 * height) - (5.677 * float64(age))
	} else {
		bmr = 447.593 + (9.247 * weight) + (3.098 * height) - (4.330 * float64(age))
	}

	// Apply activity multiplier
	var activityMultiplier float64
	switch activityLevel {
	case models.ActivityLow:
		activityMultiplier = 1.2
	case models.ActivityModerate:
		activityMultiplier = 1.55
	case models.ActivityHigh:
		activityMultiplier = 1.9
	default:
		activityMultiplier = 1.4
	}

	return int(bmr * activityMultiplier)
}
