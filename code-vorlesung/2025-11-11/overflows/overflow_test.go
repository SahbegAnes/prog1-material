package overflows

import "fmt"

func Example_overflow() {
	var x1 int8
	for x1 != -1 {
		fmt.Println(x1)
		x1++
	}

	// Output:
}
