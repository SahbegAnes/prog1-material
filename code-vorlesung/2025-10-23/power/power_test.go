package power

import "fmt"

func ExamplePow2() {
	fmt.Println(Pow2(0))
	fmt.Println(Pow2(2))
	fmt.Println(Pow2(3))
	fmt.Println(Pow2(5))

	// Output:
	// 1
	// 4
	// 8
	// 32
}
