package util

import (
	"strconv"
)

func Str2Uint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}
