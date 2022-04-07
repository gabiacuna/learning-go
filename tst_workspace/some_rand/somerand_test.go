package somerand

import (
	"fmt"
	"testing"
)

func test_sel_rand(t *testing.T) {
	// want := int(0)

	got := select_rand(3)

	if got > 0 {
		panic("select_rand(3) != 0")
		fmt.Sprint((got))
	}
	// println("got:", got)

}
