package fibonacci

import "fmt"

func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

func ExampleFib() {
	fmt.Println(Fib(1))
	fmt.Println(Fib(2))
	fmt.Println(Fib(3))
	fmt.Println(Fib(4))
	fmt.Println(Fib(5))
	fmt.Println(Fib(6))
	fmt.Println(Fib(7))

	// Output:
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// 13

}
