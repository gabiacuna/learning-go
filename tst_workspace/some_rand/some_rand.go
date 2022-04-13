package somerand

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func select_rand(n int) int {
	return rand.Intn(n)
}

// * https://go.dev/blog/json
func some_json() int {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		// panic(err)
		return -1
	}
	fmt.Printf("%s\n", b)
	return 0
}
