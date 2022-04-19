package somerand

import (
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	got := some_json()

	if got == -1 {
		fmt.Println("got:", got)
		t.Logf("some_json() = %d", got)
	}
	if got == 0 {
		fmt.Println("got:", got)
		t.Logf("some_json() = %d", got)
	}
}
