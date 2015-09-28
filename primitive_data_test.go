/* Examples with primitive data types */

package nanoexamples

import (
	"fmt"
	"math"
)

func ExampleIntDefaultType() {
	i := 1000
	fmt.Printf("%#v %T\n", i, i)
	// Output: 1000 int
}

func ExampleUintOverflow() {
	maxUint := ^uint(0)
	log2 := math.Log2(float64(maxUint))
	fmt.Printf("   maximum uint: %d (2↑%.0f)\n", maxUint, log2)
	maxUint++
	fmt.Printf("maximum uint ++: %d\n", maxUint)
	// Output:
	//    maximum uint: 18446744073709551615 (2↑64)
	// maximum uint ++: 0
}

func ExampleIntOverflow() {
	maxInt := int(^uint(0) >> 1)
	log2 := math.Log2(float64(maxInt))
	fmt.Printf("   maximum int: %d (2↑%.0f)\n", maxInt, log2)
	maxInt++
	fmt.Printf("maximum int ++: %d\n", maxInt)
	// Output:
	//    maximum int: 9223372036854775807 (2↑63)
	// maximum int ++: -9223372036854775808
}
