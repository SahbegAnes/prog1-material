package add

import "fmt"

func Add(x, y int) int {

	if y == 0 {
		return x
	}

	if y < 0 {
		return Add(x-1, y+1)
	}

	return Add(x, y-1) + 1
}

func ExampleAdd() {
	fmt.Println(Add(2, 3))
	fmt.Println(Add(2, -3))

	// Output:
	// 5
	// -1
}
