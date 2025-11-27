package ackermann

import "fmt"

func Ack(m, n int) int {
	if m == 0 {
		return n + 1
	}
	if n == 0 {
		return Ack(m-1, 1)
	}
	return Ack(m-1, Ack(m, n-1))
}

func ExampleAck() {
	fmt.Println(Ack(4, 4))

	// Output:
}
