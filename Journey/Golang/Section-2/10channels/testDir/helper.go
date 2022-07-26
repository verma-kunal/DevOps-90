package testDir

import (
	"math/rand"
)

func RandomNumber(n int) int {
	value := rand.Intn(n)

	return value
}
