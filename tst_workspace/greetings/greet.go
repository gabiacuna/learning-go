package greetings

import (
	"fmt"

	"github.com/gabiacuna/learning-go/tst_workspace/hello"
)

func main() {
	// Get a greeting message and print it.
	message := hello.Hello("Gladys")
	fmt.Println(message)
}
