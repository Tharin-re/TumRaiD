package util

import (
    "crypto/sha256"
    "fmt"
    "regexp"
    "github.com/Tharin-re/TumRaiD/src/config"
)

// MakePassHash generates a SHA-256 hash of the given password with added salt.
func MakePassHash(password string) string {
    to_hash := []byte("salty" + password) // Add salt to the password
    idx := sha256.Sum256(to_hash)         // Generate SHA-256 hash

    hashString := fmt.Sprintf("%x", idx)  // Convert hash to hexadecimal string
    return hashString
}

// ContainUnacceptableChar checks if the given string contains any unacceptable characters.
func ContainUnacceptableChar(s string) bool {
    specialCharPattern := config.Cfg.UserPassConstraints.IllegalChar // Get the pattern for illegal characters
    re := regexp.MustCompile(specialCharPattern)                     // Compile the regex pattern
    return re.MatchString(s)                                         // Check if the string matches the pattern
}

// ValidUserLength checks if the length of the username is within the acceptable range.
func ValidUserLength(username string) bool {
    validUserLengthMin := config.Cfg.UserPassConstraints.UserLengthMin // Minimum length for username
    validUserLengthMax := config.Cfg.UserPassConstraints.UserLengthMax // Maximum length for username
    return len(username) <= validUserLengthMax && len(username) >= validUserLengthMin
}

// ValidPasswordLength checks if the length of the password is within the acceptable range.
func ValidPasswordLength(password string) bool {
    validPasswordLengthMin := config.Cfg.UserPassConstraints.PasswordLengthMin // Minimum length for password
    validPasswordLengthMax := config.Cfg.UserPassConstraints.PasswordLengthMax // Maximum length for password
    return len(password) <= validPasswordLengthMax && len(password) >= validPasswordLengthMin
}
