package common

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomSequence(n int) string {
	b := make([]rune, n)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range b {
		b[i] = letters[r1.Intn(999999)%len(letters)]

	}
	return string(b)
}

func GenSalt(lenght int) string {
	if lenght < 0 {
		lenght = 50
	}
	return randomSequence(lenght)
}
