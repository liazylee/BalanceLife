package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// GenerateID generates a random ID for entities
func GenerateID() string {
	// Simple ID generation for demo purposes
	// In production, use UUID or similar library
	return fmt.Sprintf("%d%d", time.Now().UnixNano(), rnd.Intn(1000))
}
