package util

import (
	"fmt"
	"crypto/sha256"
	"regexp"
)

func MakeUserPassHash(username string, password string) string {
	to_hash := []byte(username+"_"+password)
	idx:= sha256.Sum256(to_hash)

	hashString := fmt.Sprintf("%x",idx)
	return hashString
}

func ContainUnacceptableChar(s string) bool {
	specialCharPattern := `[^a-zA-Z0-9_-]`
	re := regexp.MustCompile(specialCharPattern)
	return re.MatchString(s)
}