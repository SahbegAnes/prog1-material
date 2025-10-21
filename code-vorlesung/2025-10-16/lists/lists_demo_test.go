package lists

import "fmt"

func Example_lists() {
	l1 := []int{1, 3, 17, 25, 2, 4, 7}
	fmt.Println(l1[0])
	fmt.Println(l1[3])
	l1[3] = 42
	fmt.Println(l1[3])

	for i := 0; i < len(l1); i++ {
		fmt.Println(l1[i])
	}

	// Output:
}
