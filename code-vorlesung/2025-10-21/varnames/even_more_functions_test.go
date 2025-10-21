package varnames

import "fmt"

func Example_even_more_functions() {
	fmt.Println(Baz(4))

	// Output:
}

func Baz(x int) int {
	if x == 0 {
		return 1
	}
	return x * Baz(x-1)
}
