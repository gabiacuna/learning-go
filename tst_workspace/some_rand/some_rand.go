package somerand

import "math/rand"

func select_rand(n int) int {
	return rand.Intn(n)
}
