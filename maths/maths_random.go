package maths

import (
	"math/rand"
	"time"
)

// Random 返回 [start,end) 之间的随机整数
func Random(start int, end int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(end-start) + start
}
