package listen

import "fmt"

func ExampleDuplicateElements() {
	l1 := []int{1, 2, 3, 4, 5}
	fmt.Println(DuplicateElements(l1))
	l2 := []int{10, 222, 35, 47, 53}
	fmt.Println(DuplicateElements(l2))
	l3 := []int{}
	fmt.Println(DuplicateElements(l3))

	// Output:
	// [1 1 2 2 3 3 4 4 5 5]
	// [10 10 222 222 35 35 47 47 53 53]
	// []
}

func ExampleReverse() {
	l1 := []int{1, 2, 3, 4, 5}
	fmt.Println(Reverse(l1))

	// Output:
	// [5 4 3 2 1]
}
