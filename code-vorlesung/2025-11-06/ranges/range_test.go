package ranges

import "fmt"

func Example_range() {
	l1 := []int{10, 20, 30, 40, 50}

	for i := range l1 {
		fmt.Println(i, l1[i])
	}

	fmt.Println()

	for i, el := range l1 {
		fmt.Println(i, el)
	}

	fmt.Println()

	for _, el := range l1 {
		fmt.Println(el)
	}

	fmt.Println()

	for i, el := range l1[1:] {
		fmt.Println(i, el)
	}

	// Output:
}
