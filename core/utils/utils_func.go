package utils

import (
	"os"
)

// this one needs more checks to be robust.... but later
//return bool + error is better than just bool
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
	return false
	}
	return !info.IsDir()
}
