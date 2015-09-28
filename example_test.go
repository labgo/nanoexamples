package nanoexamples

import (
	"fmt"

	"github.com/golang/example/stringutil"
)

func ExampleReverse() {
	fmt.Println(stringutil.Reverse("hello"))
	// Output: olleh
}
