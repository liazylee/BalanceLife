package models

import (
	"time"
)

// Gender represents the gender of a user
type Gender string

// ActivityLevel represents the user's physical activity level
type ActivityLevel string

// GoalType represents the user's fitness goal type
type GoalType string

// Constants for Gender, ActivityLevel, and GoalType
const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
	GenderOther  Gender = "OTHER"

	ActivityLow      ActivityLevel = "LOW"
	ActivityModerate ActivityLevel = "MODERATE"
	ActivityHigh     ActivityLevel = "HIGH"

	GoalTypeLose GoalType = "LOSE"
	GoalTypeGain GoalType = "GAIN"
	GoalTypeAll  GoalType = "" // Used for filtering all goal types
)

// GoalInfo represents a user's fitness goal
type GoalInfo struct {
	Type           GoalType  `json:"type" bson:"type"`
	TargetCalories int       `json:"targetCalories" bson:"targetCalories"`
	TargetProtein  int       `json:"targetProtein" bson:"targetProtein"`
	TargetCarbs    int       `json:"targetCarbs" bson:"targetCarbs"`
	TargetFat      int       `json:"targetFat" bson:"targetFat"`
	StartDate      time.Time `json:"startDate" bson:"startDate"`
	StartWeight    float64   `json:"startWeight,omitempty" bson:"startWeight,omitempty"`
	TargetWeight   float64   `json:"targetWeight,omitempty" bson:"targetWeight,omitempty"`
}

// User represents a user in the system
type User struct {
	ID            string        `json:"userId" bson:"_id"`
	Name          string        `json:"name" bson:"name"`
	Email         string        `json:"email" bson:"email"`
	Password      string        `json:"-" bson:"password"` // Never expose password
	Gender        Gender        `json:"gender" bson:"gender"`
	BirthDate     time.Time     `json:"birthDate" bson:"birthDate"`
	CreatedAt     time.Time     `json:"createdAt" bson:"createdAt"`
	LastLoginAt   time.Time     `json:"lastLoginAt" bson:"lastLoginAt"`
	Height        float64       `json:"height" bson:"height"`
	Weight        float64       `json:"weight" bson:"weight"`
	ActivityLevel ActivityLevel `json:"activityLevel" bson:"activityLevel"`
	Goal          GoalInfo      `json:"goal" bson:"goal"`
}
type UserInfo struct {
	ID            string        `json:"_id" bson:"_id"`
	UserId        string        `json:"userId" bson:"userId"`
	Height        float64       `json:"height" bson:"height"`
	Weight        float64       `json:"weight" bson:"weight"`
	ActivityLevel ActivityLevel `json:"activityLevel" bson:"activityLevel"`
	Goal          GoalInfo      `json:"goal" bson:"goal"`
	BMIRate       float64       `json:"bmiRate" bson:"bmiRate"`
	MetabolicRate float64       `json:"metabolicRate" bson:"metabolicRate"`
	CreatedAt     time.Time     `json:"createdAt" bson:"createdAt"`
}
