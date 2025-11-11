package length

import "fmt"

func ExampleLength_conversions() {
	a := FromCentimeters(5)
	b := FromKilometers(5)
	c := FromMillimeters(50)

	fmt.Println(a.Centimeters())
	fmt.Println(b.Centimeters())
	fmt.Println(c.Centimeters())
	fmt.Println(a.Millimeters())

	// Output:
	// 5
	// 500000
	// 5
	// 50
}
