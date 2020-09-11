package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Encode(str string) string {
	encode := md5.New()
	encode.Write([]byte(str))
	return hex.EncodeToString(encode.Sum(nil))
}
