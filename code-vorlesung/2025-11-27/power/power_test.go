package power

import "fmt"

// Power2 berechnet 2^x
func Power2(x int) int {
	// 2^0 == 1
	if x == 0 {
		return 1
	}
	// 2^x == 2 * 2^(x-1)
	return 2 * Power2(x-1)
}

func ExamplePower2() {
	fmt.Println(Power2(3))
	fmt.Println(Power2(2))
	fmt.Println(Power2(5))
	fmt.Println(Power2(0))

	// Output:
	// 8
	// 4
	// 32
	// 1
}
