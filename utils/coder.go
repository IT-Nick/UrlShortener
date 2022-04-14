package utils

import (
	"math/rand"
)

const PATTERN = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789_"

const (
	bits = 6
	mask = 1<<bits - 1
	max  = 63/bits
)

func ByteMaskGen() string {
	result := make([]byte, 10)
	for i, cache, remain := 9, rand.Int63(), max; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), max
		}
		if idx := int(cache & mask); idx < len(PATTERN) {
			result[i] = PATTERN[idx]
			i--
		}
		cache >>= bits
		remain--
	}

	return string(result)
}