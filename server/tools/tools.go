package tools

import (
	"math/rand"
)

const CHARS = "abcdefghijklmnopqrstuvwxyz"

func RandomString(len int) string {
	rtn := ""
	for i := 0; i < len; i++ {
		index := rand.Intn(26)
		rtn = rtn + CHARS[index:index+1]
	}
	return rtn
}
