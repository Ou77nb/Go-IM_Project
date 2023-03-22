package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// MakePassword 加密
func MakePassword(pwd, salt string) string {
	data := pwd + salt
	return MD5Encode(data)
}
