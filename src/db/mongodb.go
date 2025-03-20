package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/zhenyili/BalanceLife/src/config"
	"github.com/zhenyili/BalanceLife/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection names
const (
	usersCollection           = "users"
	mealPackagesCollection    = "meal_packages"
	workoutPackagesCollection = "workout_packages"
	mealEntriesCollection     = "meal_entries"
	workoutEntriesCollection  = "workout_entries"
)

// MongoStore implements the Store interface using MongoDB
type MongoStore struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
}

// NewMongoStore creates a new MongoDB-backed store
func NewMongoStore(cfg *config.AppConfig) (*MongoStore, error) {
	ctx := context.Background()

	// Configure MongoDB connection options
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(cfg.MongoDB.URI).
		SetServerAPIOptions(serverAPIOptions)

	// Add TLS certificate if provided
	if cfg.MongoDB.CertificatePath != "" {
		clientOptions.SetTLSConfig(nil) // Setting to nil lets the driver use system certs
		clientOptions = clientOptions.ApplyURI(fmt.Sprintf(
			"%s&tlsCertificateKeyFile=%s",
			cfg.MongoDB.URI,
			cfg.MongoDB.CertificatePath,
		))
	}

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		// Clean up if connection fails
		if closeErr := client.Disconnect(ctx); closeErr != nil {
			log.Printf("Error disconnecting from MongoDB: %v", closeErr)
		}
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	// Create MongoStore instance
	database := client.Database(cfg.MongoDB.Database)
	store := &MongoStore{
		client: client,
		db:     database,
		ctx:    ctx,
	}

	// Create indexes
	if err := store.createIndexes(); err != nil {
		log.Printf("Warning: Failed to create indexes: %v", err)
	}

	return store, nil
}

// Close closes the MongoDB connection
func (s *MongoStore) Close() error {
	return s.client.Disconnect(s.ctx)
}

// createIndexes creates indexes for the MongoDB collections
func (s *MongoStore) createIndexes() error {
	// Create indexes for users collection
	_, err := s.db.Collection(usersCollection).Indexes().CreateOne(
		s.ctx,
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		return err
	}

	// More indexes could be added here as needed

	return nil
}

// GetUsers returns all users
func (s *MongoStore) GetUsers() []models.User {
	var users []models.User
	cursor, err := s.db.Collection(usersCollection).Find(s.ctx, bson.D{})
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return users
	}
	defer cursor.Close(s.ctx)

	if err := cursor.All(s.ctx, &users); err != nil {
		log.Printf("Error decoding users: %v", err)
	}

	return users
}

// GetUser returns a specific user by ID
func (s *MongoStore) GetUser(id string) (models.User, error) {
	var user models.User

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// If not a valid ObjectID, try finding by ID string directly
		err = s.db.Collection(usersCollection).FindOne(s.ctx, bson.M{"id": id}).Decode(&user)
	} else {
		// Try finding by ObjectID
		err = s.db.Collection(usersCollection).FindOne(s.ctx, bson.M{"_id": objID}).Decode(&user)
	}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}

	return user, nil
}

// CreateUser creates a new user
func (s *MongoStore) CreateUser(user models.User) (models.User, error) {
	// Ensure the user has an ID
	if user.ID == "" {
		user.ID = primitive.NewObjectID().Hex()
	}

	// Convert the model to BSON
	_, err := s.db.Collection(usersCollection).InsertOne(s.ctx, user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// GetMealPackages returns meal packages, optionally filtered by goal type
func (s *MongoStore) GetMealPackages(goalType models.GoalType) []models.MealPackage {
	var packages []models.MealPackage

	filter := bson.D{}
	if goalType != models.GoalTypeAll {
		filter = bson.D{{Key: "goalType", Value: goalType}}
	}

	cursor, err := s.db.Collection(mealPackagesCollection).Find(s.ctx, filter)
	if err != nil {
		log.Printf("Error fetching meal packages: %v", err)
		return packages
	}
	defer cursor.Close(s.ctx)

	if err := cursor.All(s.ctx, &packages); err != nil {
		log.Printf("Error decoding meal packages: %v", err)
	}

	return packages
}

// GetMealPackage returns a specific meal package by ID
func (s *MongoStore) GetMealPackage(id string) (models.MealPackage, error) {
	var pkg models.MealPackage

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// If not a valid ObjectID, try finding by ID string directly
		err = s.db.Collection(mealPackagesCollection).FindOne(s.ctx, bson.M{"id": id}).Decode(&pkg)
	} else {
		// Try finding by ObjectID
		err = s.db.Collection(mealPackagesCollection).FindOne(s.ctx, bson.M{"_id": objID}).Decode(&pkg)
	}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.MealPackage{}, errors.New("meal package not found")
		}
		return models.MealPackage{}, err
	}

	return pkg, nil
}

// GetWorkoutPackages returns workout packages, optionally filtered by goal type
func (s *MongoStore) GetWorkoutPackages(goalType models.GoalType) []models.WorkoutPackage {
	var packages []models.WorkoutPackage

	filter := bson.D{}
	if goalType != models.GoalTypeAll {
		filter = bson.D{{Key: "goalType", Value: goalType}}
	}

	cursor, err := s.db.Collection(workoutPackagesCollection).Find(s.ctx, filter)
	if err != nil {
		log.Printf("Error fetching workout packages: %v", err)
		return packages
	}
	defer cursor.Close(s.ctx)

	if err := cursor.All(s.ctx, &packages); err != nil {
		log.Printf("Error decoding workout packages: %v", err)
	}

	return packages
}

// GetWorkoutPackage returns a specific workout package by ID
func (s *MongoStore) GetWorkoutPackage(id string) (models.WorkoutPackage, error) {
	var pkg models.WorkoutPackage

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// If not a valid ObjectID, try finding by ID string directly
		err = s.db.Collection(workoutPackagesCollection).FindOne(s.ctx, bson.M{"id": id}).Decode(&pkg)
	} else {
		// Try finding by ObjectID
		err = s.db.Collection(workoutPackagesCollection).FindOne(s.ctx, bson.M{"_id": objID}).Decode(&pkg)
	}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.WorkoutPackage{}, errors.New("workout package not found")
		}
		return models.WorkoutPackage{}, err
	}

	return pkg, nil
}

// CreateMealEntry creates a new meal entry
func (s *MongoStore) CreateMealEntry(entry models.MealEntry) (models.MealEntry, error) {
	// Ensure the entry has an ID
	if entry.ID == "" {
		entry.ID = primitive.NewObjectID().Hex()
	}
	// Ensure the timestamp is set
	if entry.Timestamp.IsZero() {
		entry.Timestamp = time.Now()
	}

	// Insert the entry
	_, err := s.db.Collection(mealEntriesCollection).InsertOne(s.ctx, entry)
	if err != nil {
		return models.MealEntry{}, err
	}

	return entry, nil
}

// GetMealEntriesByUserAndDateRange returns meal entries for a user within a date range
func (s *MongoStore) GetMealEntriesByUserAndDateRange(userID string, startDate, endDate time.Time) []models.MealEntry {
	var entries []models.MealEntry

	// Create a date range filter
	filter := bson.M{
		"userID": userID,
		"timestamp": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	}

	cursor, err := s.db.Collection(mealEntriesCollection).Find(s.ctx, filter)
	if err != nil {
		log.Printf("Error fetching meal entries: %v", err)
		return entries
	}
	defer cursor.Close(s.ctx)

	if err := cursor.All(s.ctx, &entries); err != nil {
		log.Printf("Error decoding meal entries: %v", err)
	}

	return entries
}

// CreateWorkoutEntry creates a new workout entry
func (s *MongoStore) CreateWorkoutEntry(entry models.WorkoutEntry) (models.WorkoutEntry, error) {
	// Ensure the entry has an ID
	if entry.ID == "" {
		entry.ID = primitive.NewObjectID().Hex()
	}
	// Ensure the timestamp is set
	if entry.Timestamp.IsZero() {
		entry.Timestamp = time.Now()
	}

	// Insert the entry
	_, err := s.db.Collection(workoutEntriesCollection).InsertOne(s.ctx, entry)
	if err != nil {
		return models.WorkoutEntry{}, err
	}

	return entry, nil
}

// GetWorkoutEntriesByUserAndDateRange returns workout entries for a user within a date range
func (s *MongoStore) GetWorkoutEntriesByUserAndDateRange(userID string, startDate, endDate time.Time) []models.WorkoutEntry {
	var entries []models.WorkoutEntry

	// Create a date range filter
	filter := bson.M{
		"userID": userID,
		"timestamp": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	}

	cursor, err := s.db.Collection(workoutEntriesCollection).Find(s.ctx, filter)
	if err != nil {
		log.Printf("Error fetching workout entries: %v", err)
		return entries
	}
	defer cursor.Close(s.ctx)

	if err := cursor.All(s.ctx, &entries); err != nil {
		log.Printf("Error decoding workout entries: %v", err)
	}

	return entries
}

// DeleteUser deletes a user by ID and returns the deleted user
func (s *MongoStore) DeleteUser(id string) (models.User, error) {
	// Find the user first to return it
	userCollection := s.db.Collection(usersCollection)
	filter := bson.M{"id": id}

	// First find the user to return it
	var user models.User
	err := userCollection.FindOne(s.ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, fmt.Errorf("user not found: %s", id)
		}
		return models.User{}, err
	}

	// Then delete it
	_, err = userCollection.DeleteOne(s.ctx, filter)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to delete user: %w", err)
	}

	return user, nil
}
