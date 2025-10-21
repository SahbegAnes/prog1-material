package varnames

import "fmt"

func Example_scopes() {

	x := 1
	{
		x := 2
		{
			x := 3
			{
				fmt.Println(x)
			}
			fmt.Println(x)
		}
		fmt.Println(x)
	}
	fmt.Println(x)

	// Output:
}

func Example_scopes_loops() {
	x := 42
	y := 23

	z := 10
	for x := 0; x < 5; x++ {
		z = 4
		fmt.Println(z)
		fmt.Println(x)
		fmt.Println(y)
	}

	fmt.Println(z)
	fmt.Println(x)

	// Output:
}
