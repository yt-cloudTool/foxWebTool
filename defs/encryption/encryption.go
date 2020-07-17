package encryption

import (
	"golang.org/x/crypto/bcrypt"
)

// 加密
func Generate(pwd_str string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd_str), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hash)
}

// 解密
func Compare(hash_str string, pwd_str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash_str), []byte(pwd_str))
	if err != nil {
		return false
	} else {
		return true
	}
}
