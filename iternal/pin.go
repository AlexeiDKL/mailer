package iternal

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandPin(n int) int {
	rand.Seed(time.Now().UnixNano())

	min := 1
	max := 9
	for i := 0; i < n-1; i++ {
		min = min*10 + 0
		max = max*10 + 9
	}
	min = min - 1
	max = max + 1
	result := rand.Intn(max-min) + min

	fmt.Println(min, max, result)
	return result
}
