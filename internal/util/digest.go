package util

import (
	"crypto/md5"
	"fmt"
	"os"
)

func GetDigest(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	buffer := make([]byte, 64*1024) // 64 KiB
	_, _ = f.Read(buffer)
	return fmt.Sprintf("%x", md5.Sum(buffer)), nil
}
