package functions

import (
	"fmt"
	"slices"
)

func Plus2(x int) int {
	return x + 2
}

func Plus5(x int) int {
	return x + 5
}

func Times5(x int) int {
	return x * 5
}

func ForEach(list []int, f func(int) int) {
	for i, el := range list {
		list[i] = f(el)
	}
}

func Example_function_params() {
	l1 := []int{1, 2, 3, 4, 5}

	fmt.Println(l1)

	ForEach(l1, Plus2)
	ForEach(l1, Plus5)
	ForEach(l1, Times5)

	fmt.Println(l1)

	// Output:
}

func Example_closures() {
	l1 := []int{1, 2, 3, 4, 5}

	double_element := func(i int) {
		l1[i] *= 2
	}

	for i := 0; i < len(l1); i++ {
		double_element(i)
	}

	fmt.Println(l1)

	// Output:
}

func Example_sort() {
	l1 := []int{17, 2, 3, 15, 25, 42, 38}

	comp := func(i, j int) int {
		if i == j {
			return 0
		}
		if i > j {
			return -1
		}
		return 1
	}

	slices.SortFunc(l1, comp)

	fmt.Println(l1)

	// Output:
}
