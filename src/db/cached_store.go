package db

import (
	"log"
	"time"

	"github.com/zhenyili/BalanceLife/src/config"
	"github.com/zhenyili/BalanceLife/src/models"
)

// MongodbStore implements the Store interface using MongoDB for persistence
// It provides direct access to the MongoDB database
type MongodbStore struct {
	db *MongoStore
}

// NewMongodbStore creates a new MongoDB-based store
// It requires a valid MongoDB connection
func NewMongodbStore(cfg *config.AppConfig) (*MongodbStore, error) {
	// Initialize MongoDB connection
	db, err := NewMongoStore(cfg)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB successfully")
	return &MongodbStore{
		db: db,
	}, nil
}

// Close closes all connections
func (s *MongodbStore) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

// HasMongoDB returns true if MongoDB is available
func (s *MongodbStore) HasMongoDB() bool {
	return s.db != nil
}

// User-related methods

// GetUsers returns all users
func (s *MongodbStore) GetUsers() []models.User {
	return s.db.GetUsers()
}

// GetUser returns a user by ID
func (s *MongodbStore) GetUser(id string) (models.User, error) {
	// Get from database
	user, err := s.db.GetUser(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// CreateUser creates a new user
func (s *MongodbStore) CreateUser(user models.User) (models.User, error) {
	// Create in database
	return s.db.CreateUser(user)
}

// DeleteUser deletes a user by ID
func (s *MongodbStore) DeleteUser(id string) (models.User, error) {
	// Delete from database
	return s.db.DeleteUser(id)
}

// MealPackage-related methods

// GetMealPackages returns all meal packages, optionally filtered by goal type
func (s *MongodbStore) GetMealPackages(goalType models.GoalType) []models.MealPackage {
	return s.db.GetMealPackages(goalType)
}

// GetMealPackage returns a meal package by ID
func (s *MongodbStore) GetMealPackage(id string) (models.MealPackage, error) {
	return s.db.GetMealPackage(id)
}

// WorkoutPackage-related methods

// GetWorkoutPackages returns all workout packages, optionally filtered by goal type
func (s *MongodbStore) GetWorkoutPackages(goalType models.GoalType) []models.WorkoutPackage {
	return s.db.GetWorkoutPackages(goalType)
}

// GetWorkoutPackage returns a workout package by ID
func (s *MongodbStore) GetWorkoutPackage(id string) (models.WorkoutPackage, error) {
	return s.db.GetWorkoutPackage(id)
}

// MealEntry-related methods

// CreateMealEntry adds a new meal entry
func (s *MongodbStore) CreateMealEntry(entry models.MealEntry) (models.MealEntry, error) {
	return s.db.CreateMealEntry(entry)
}

// GetMealEntriesByUserAndDateRange returns meal entries for a user within a date range
func (s *MongodbStore) GetMealEntriesByUserAndDateRange(userID string, startDate, endDate time.Time) []models.MealEntry {
	return s.db.GetMealEntriesByUserAndDateRange(userID, startDate, endDate)
}

// WorkoutEntry-related methods

// CreateWorkoutEntry adds a new workout entry
func (s *MongodbStore) CreateWorkoutEntry(entry models.WorkoutEntry) (models.WorkoutEntry, error) {
	return s.db.CreateWorkoutEntry(entry)
}

// GetWorkoutEntriesByUserAndDateRange returns workout entries for a user within a date range
func (s *MongodbStore) GetWorkoutEntriesByUserAndDateRange(userID string, startDate, endDate time.Time) []models.WorkoutEntry {
	return s.db.GetWorkoutEntriesByUserAndDateRange(userID, startDate, endDate)
}
