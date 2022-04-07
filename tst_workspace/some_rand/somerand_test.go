package somerand

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	got := select_rand(3)

	if got < 3 {
		fmt.Println("got:", got)
		t.Logf("select_rand(3) = %d", got)
	}

}
