package models

import "time"

// WorkoutPackage represents a predefined workout package in the system
type WorkoutPackage struct {
	ID                  string   `json:"packageId" bson:"_id"`
	Name                string   `json:"name" bson:"name"`
	Description         string   `json:"description" bson:"description"`
	GoalType            GoalType `json:"goalType" bson:"goalType"` // LOSE, GAIN, or BOTH
	WorkoutType         string   `json:"workoutType" bson:"workoutType"`
	BaseDurationMinutes int      `json:"baseDurationMinutes" bson:"baseDurationMinutes"`
	BaseCaloriesBurn    int      `json:"baseCaloriesBurn" bson:"baseCaloriesBurn"`
	CaloriesBurnFormula string   `json:"caloriesBurnFormula" bson:"caloriesBurnFormula"`
	ImageURL            string   `json:"imageUrl" bson:"imageUrl"`
	Instructions        []string `json:"instructions,omitempty" bson:"instructions,omitempty"`
}

// WorkoutEntry represents a logged workout by a user
type WorkoutEntry struct {
	ID                  string    `json:"entryId" bson:"_id"`
	UserID              string    `json:"userId" bson:"userId"`
	PackageID           string    `json:"packageId" bson:"packageId"`
	IntensityMultiplier float64   `json:"intensityMultiplier" bson:"intensityMultiplier"`
	DurationMinutes     int       `json:"durationMinutes" bson:"durationMinutes"`
	CaloriesBurned      int       `json:"caloriesBurned" bson:"caloriesBurned"`
	Date                time.Time `json:"date" bson:"date"`
	Timestamp           time.Time `json:"timestamp" bson:"timestamp"` // Used for querying by time range
	CreatedAt           time.Time `json:"createdAt" bson:"createdAt"`
}
