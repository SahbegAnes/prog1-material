package duration

import "fmt"

func ExampleDuration_conversions() {
	fmt.Println(FromMinutes(3).Seconds())

	// Output:
	// 180
}
