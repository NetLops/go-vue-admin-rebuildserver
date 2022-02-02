package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5V
//
// Description: MD5加密
//
//
// param: str []byte
// param: b ...byte
//
// return: string
func MD5V(str []byte, b ...byte) string {
	hash := md5.New()
	hash.Write(str)
	return hex.EncodeToString(hash.Sum(b))
}
