package matrices

import "fmt"

func Example_matrix() {
	m1 := [][]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(m1)

	// Zeile 0 ausgeben:
	fmt.Println(m1[0])

	// Element (0,0) ausgeben:
	fmt.Println(m1[0][0])

	// Output:
}

// PrintMatrix gibt die Matrix m Zeile für Zeile aus.
func PrintMatrix(m [][]int) {
	for row := 0; row < len(m); row++ {
		for col := 0; col < len(m[row]); col++ {
			//fmt.Printf("%v ", m[row][col])
			fmt.Print(m[row][col])
			if col < len(m[row])-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func ExamplePrintMatrix() {
	m1 := [][]int{{1, 2}, {3, 4, 5}}

	PrintMatrix(m1)

	// Output:
	// 1 2
	// 3 4 5
}
