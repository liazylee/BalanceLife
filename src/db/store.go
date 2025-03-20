package db

import (
	"time"

	"github.com/zhenyili/BalanceLife/src/models"
)

// Store defines the interface for data storage
type Store interface {
	// User operations
	GetUsers() []models.User
	GetUser(id string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	DeleteUser(id string) (models.User, error)

	// MealPackage operations
	GetMealPackages(goalType models.GoalType) []models.MealPackage
	GetMealPackage(id string) (models.MealPackage, error)

	// WorkoutPackage operations
	GetWorkoutPackages(goalType models.GoalType) []models.WorkoutPackage
	GetWorkoutPackage(id string) (models.WorkoutPackage, error)
	// MealEntry operations
	CreateMealEntry(entry models.MealEntry) (models.MealEntry, error)
	GetMealEntriesByUserAndDateRange(userID string, startDate, endDate time.Time) []models.MealEntry

	// WorkoutEntry operations
	CreateWorkoutEntry(entry models.WorkoutEntry) (models.WorkoutEntry, error)
	GetWorkoutEntriesByUserAndDateRange(userID string, startDate, endDate time.Time) []models.WorkoutEntry
}
