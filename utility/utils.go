package utility

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"math/rand"
	"time"
)

// RandInt 随机数
func RandInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

// EncryptPassword 密码加密
func EncryptPassword(password, salt string) string {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + gmd5.MustEncryptString(salt))
}
