package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encode 小写的
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// MD5Encode 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// MakePasswd 密码加密
func MakePasswd(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

// ValidatePasswd 密码校验
func ValidatePasswd(plainpwd, salt, passwd string) bool {
	return Md5Encode(plainpwd+salt) == passwd
}
