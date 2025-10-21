package varnames

import "fmt"

func Example_more_functions() {
	PrintSomething(3)

	// Output:
}

func PrintSomething(x int) {
	for i := 0; i < 5; i++ {
		PrintMore(x + i)
	}
}

func PrintMore(x int) {
	for i := 0; i < x; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}
