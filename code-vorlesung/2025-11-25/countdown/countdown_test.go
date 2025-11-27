package countdown

import "fmt"

func CountDown(n int) {
	if n <= 0 {
		return
	}
	fmt.Println(n)
	CountDown(n - 1)
}

func ExampleCountDown() {
	CountDown(3)

	// Output:
}

// CountUp gibt die Zahlen von 1 bis n aufsteigend auf der Konsole aus.
func CountUp(n int) {
	if n <= 0 {
		return
	}
	CountUp(n - 1)
	fmt.Println(n)
}

func ExampleCountUp() {
	CountUp(3)

	// Output:
}
