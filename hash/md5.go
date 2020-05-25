package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func StringMd5(s string) string  {
	md55 := md5.New()
	md55.Write([]byte(s))
	return hex.EncodeToString(md55.Sum(nil))
}