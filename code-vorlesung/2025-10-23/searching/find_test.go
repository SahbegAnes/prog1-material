package searching

import "fmt"

func ExampleFind() {
	l1 := []int{1, 6, 8, 10, 25, 42, 125, 277}

	fmt.Println(Find(l1, 3))
	fmt.Println(Find(l1, 8))

	// Output:
	// -1
	// 2
}
