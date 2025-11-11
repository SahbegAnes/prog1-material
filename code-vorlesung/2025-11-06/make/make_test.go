package make

import "fmt"

func Example_make() {
	l1 := []int{}
	l2 := make([]int, 5)
	l3 := make([]string, 3)

	fmt.Println(l1)
	fmt.Println(len(l1))

	fmt.Println(l2)
	fmt.Println(len(l2))

	fmt.Println(l3)
	fmt.Println(len(l3))

	// Output:
}
