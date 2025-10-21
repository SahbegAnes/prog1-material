package listprops

import "fmt"

func ExampleContains() {
	l := []string{"Hallo", "Welt", "gute", "Nacht"}
	fmt.Println(Contains(l, "Welt"))
	fmt.Println(Contains(l, "Haus"))

	// Output:
	// true
	// false
}
func ExampleContainsOnly() {
	l1 := []string{"Hallo", "Welt", "gute", "Nacht"}
	fmt.Println(ContainsOnly(l1, "Welt"))
	l2 := []string{"Welt", "Welt", "Welt", "Welt"}
	fmt.Println(ContainsOnly(l2, "Welt"))
	fmt.Println(ContainsOnly(l2, "Haus"))
	l3 := []string{}
	fmt.Println(ContainsOnly(l3, "Welt"))

	// Output:
	// false
	// true
	// false
	// true
}
func ExampleContainsN() {
	l1 := []string{"Welt", "Haus", "Welt", "Welt"}
	fmt.Println(ContainsN(l1, "Welt", 2))
	fmt.Println(ContainsN(l1, "Haus", 2))

	// Output:
	// true
	// false
}
func ExampleContainsNRow() {
	l1 := []string{"Welt", "Haus", "Welt", "Welt", "Welt"}
	fmt.Println(ContainsNRow(l1, "Welt", 2))
	fmt.Println(ContainsNRow(l1, "Welt", 3))
	fmt.Println(ContainsNRow(l1, "Welt", 4))

	// Output:
	// true
	// true
	// false
}
