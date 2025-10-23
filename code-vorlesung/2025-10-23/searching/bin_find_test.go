package searching

import "fmt"

func ExampleBinFind1() {
	l1 := []int{1, 6, 8, 10, 25, 42, 125, 277}

	//fmt.Println(BinFind1(l1, 3))
	fmt.Println(BinFind1(l1, 1))
	fmt.Println(BinFind1(l1, 277))
	fmt.Println(BinFind1(l1, 25))
	fmt.Println(BinFind1(l1, 10))
	fmt.Println(BinFind1(l1, 125))
	fmt.Println(BinFind1(l1, 3))

	// Output:
	// 0
	// 7
	// 4
	// 3
	// 6
	// -1
}

func ExampleBinFind2() {
	l1 := []int{1, 6, 8, 10, 25, 42, 125, 277}

	fmt.Println(BinFind2(l1, 1))
	fmt.Println(BinFind2(l1, 277))
	fmt.Println(BinFind2(l1, 25))
	fmt.Println(BinFind2(l1, 10))
	fmt.Println(BinFind2(l1, 125))
	fmt.Println(BinFind2(l1, 3))

	// Output:
	// 0
	// 7
	// 4
	// 3
	// 6
	// -1
}

func ExampleBinFindRek() {
	l1 := []int{1, 6, 8, 10, 25, 42, 125, 277}
	fmt.Println(BinFindRek(l1, 1))
	fmt.Println(BinFindRek(l1, 277))
	fmt.Println(BinFindRek(l1, 25))
	fmt.Println(BinFindRek(l1, 10))
	fmt.Println(BinFindRek(l1, 125))
	fmt.Println(BinFindRek(l1, 3))

	// Output:
	// 0
	// 7
	// 4
	// 3
	// 6
	// -1
}
