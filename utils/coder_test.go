package utils

import (
	"strings"
	"testing"
)

func TestByteMaskGen(t *testing.T) {
	lenRes := ByteMaskGen()
	if len(lenRes) != 10 {
		t.Fatalf("Want %v, but got %v", 10, lenRes)
	}
	conRes := ByteMaskGen()
	if strings.Contains(conRes, "+") {
		t.Fatalf("Contains blocked letter %v", "+")
	}
}
