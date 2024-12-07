package util

import (
	"crypto/sha256"
	"fmt"
	"regexp"
)

func MakePassHash(password string) string {
	to_hash := []byte("salty"+password)
	idx := sha256.Sum256(to_hash)

	hashString := fmt.Sprintf("%x", idx)
	return hashString
}

func ContainUnacceptableChar(s string) bool {
	specialCharPattern := `[^a-zA-Z0-9_-]`
	re := regexp.MustCompile(specialCharPattern)
	return re.MatchString(s)
}
