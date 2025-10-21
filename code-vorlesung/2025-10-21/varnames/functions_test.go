package varnames

import "fmt"

func Example_functions() {
	fmt.Println(Foo(15))

	// Output:
	// 17
	// 0, 4
	// 17
	// 0
}

func Foo(x int) int {
	x += 2
	y := 0
	// x == 17 && y == 0
	fmt.Println(Bar(x, y))

	fmt.Println(x)

	return y
}

func Bar(z, k int) string { // z == 17 && k == 0
	fmt.Println(z)
	y := 4
	return fmt.Sprintf("%v, %v", k, y)
}
