package listenausgabe

import "fmt"

// Output gibt die Elemente der Liste einzeln auf die Konsole aus.
func Output(l []int) {
	if len(l) == 0 {
		return
	}

	head, tail := l[0], l[1:]

	fmt.Println(head)
	Output(tail)
}

func Example() {
	l1 := []int{5, 2, 17, 23, 42, 25, 38}

	Output(l1)

	// Output:
}
