package models

import "time"

// MealType represents the type of meal (breakfast, lunch, etc.)
type MealType string

// Constants for MealType
const (
	MealTypeBreakfast MealType = "BREAKFAST"
	MealTypeLunch     MealType = "LUNCH"
	MealTypeDinner    MealType = "DINNER"
	MealTypeSnack     MealType = "SNACK"
)

// MealPackage represents a predefined meal package in the system
type MealPackage struct {
	ID               string   `json:"packageId" bson:"_id"`
	Name             string   `json:"name" bson:"name"`
	Description      string   `json:"description" bson:"description"`
	GoalType         GoalType `json:"goalType" bson:"goalType"` // LOSE, GAIN, or BOTH
	MealType         MealType `json:"mealType" bson:"mealType"`
	BaseCalories     int      `json:"baseCalories" bson:"baseCalories"`
	BaseProtein      int      `json:"baseProtein" bson:"baseProtein"`
	BaseCarbs        int      `json:"baseCarbs" bson:"baseCarbs"`
	BaseFat          int      `json:"baseFat" bson:"baseFat"`
	ImageURL         string   `json:"imageUrl" bson:"imageUrl"`
	PreparationSteps []string `json:"preparationSteps,omitempty" bson:"preparationSteps,omitempty"`
	Ingredients      []string `json:"ingredients,omitempty" bson:"ingredients,omitempty"`
}

// MealEntry represents a logged meal by a user
type MealEntry struct {
	ID                string    `json:"entryId" bson:"_id"`
	UserID            string    `json:"userId" bson:"userId"`
	PackageID         string    `json:"packageId" bson:"packageId"`
	PortionMultiplier float64   `json:"portionMultiplier" bson:"portionMultiplier"`
	Calories          int       `json:"calories" bson:"calories"`
	Protein           int       `json:"protein" bson:"protein"`
	Carbs             int       `json:"carbs" bson:"carbs"`
	Fat               int       `json:"fat" bson:"fat"`
	MealType          MealType  `json:"mealType" bson:"mealType"`
	Date              time.Time `json:"date" bson:"date"`
	Timestamp         time.Time `json:"timestamp" bson:"timestamp"` // Used for querying by time range
	CreatedAt         time.Time `json:"createdAt" bson:"createdAt"`
}
