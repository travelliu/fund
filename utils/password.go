// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package utils

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"unicode"
)

// Password validates plain password against the rules defined below.
//
// upp: at least one upper case letter.
// low: at least one lower case letter.
// num: at least one digit.
// sym: at least one special character.
// tot: at least eight characters long.
// No empty string or whitespace.
func Password(pass string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}

// PWEncrypt smtpserver password encrypt
func PWEncrypt(pw string) string {
	return base64.RawStdEncoding.EncodeToString([]byte(pw))
}

// PWDecrypt smtpserver password decrypt
func PWDecrypt(p string) (string, error) {
	r, err := base64.RawStdEncoding.DecodeString(p)
	return string(r), err
}

// HashPassword hash password
func HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash), err
}

// ComparePassword compare password
func ComparePassword(hashPassword, target string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(target)) == nil
}
