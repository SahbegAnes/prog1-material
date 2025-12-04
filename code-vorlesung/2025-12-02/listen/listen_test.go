package listen

import "fmt"

func ExampleReverse() {
	l1 := []int{1, 2, 3, 4, 5}
	fmt.Println(Reverse(l1))

	// Output:
	// [ 5 4 3 2 1 ]
}

func ExampleDuplicateElements() {
	l1 := []int{1, 2, 3, 4, 5}
	fmt.Println(DuplicateElements(l1))

	// Output:
	// [ 1 1 2 2 3 3 4 4 5 5 ]
}
