package util

import (
	"errors"
	"os"
	"time"
)

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func LastModify(path string) time.Time {
	s, err := os.Stat(path)
	if err != nil {
		return time.Time{}
	}
	return s.ModTime()
}
