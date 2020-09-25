package utils

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
)

// 密码加密算法，scrypt 加密
func ScryptPassword(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 64, 33, 221, 213, 11}
	hashpwd, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		logger.Println("scrypt error", err)
		return ""
	}

	fpwd := base64.StdEncoding.EncodeToString(hashpwd)
	return fpwd
}

// scrypt 比较密码，直接和加密后的密码比较即可
