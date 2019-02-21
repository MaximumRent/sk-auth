package crypto

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// encryption constants
const (
	_DEFAULT_COST = 10
)

// working constants
const (
	_INVALID_RESULT = ""
)

// Encrypt password using BCRYPT algorithm
// Returns array definition of result password
func EncryptPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	encryptedBytes, err := bcrypt.GenerateFromPassword(passwordBytes, _DEFAULT_COST)
	if err != nil {
		return _INVALID_RESULT, errors.New("Can't encrypt password! Cause: " + err.Error())
	}
	return string(encryptedBytes[:]), nil
}

func ComparePasswords(password string, encryptedPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPass), []byte(password))
}
