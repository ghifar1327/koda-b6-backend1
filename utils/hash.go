package utils

import "github.com/matthewhartstonge/argon2"

var argon = argon2.DefaultConfig()

func HashPassword(password string) (string, error) {
	hash, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(password string, encodedHash string) (bool, error) {
	return argon2.VerifyEncoded([]byte(password), []byte(encodedHash))
}