package utility

import (
	"math/rand"
	"time"
)

// RandInt 随机数
func RandInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
