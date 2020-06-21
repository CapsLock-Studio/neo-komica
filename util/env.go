package util

import (
	"os"
)

// Getenv extends os.Getenv with failback string when the target env not exits.
func Getenv(key, failback string) string {
	value := os.Getenv(key)
	if value == "" {
		return failback
	}

	return value
}
