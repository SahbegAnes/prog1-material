package datatypes

import "fmt"

func Example_integers() {
	var i int
	k := 55

	fmt.Println(i)
	fmt.Println(k)

	// Output:
	// 0
	// 55
}

func Example_various_data_types() {
	var i1 int
	i2 := 55

	fmt.Println(i1)
	fmt.Println(i2)

	var f1 float64
	f2 := 3.0 / 4.0

	fmt.Println(f1)
	fmt.Println(f2)

	var b1 bool
	b2 := true

	fmt.Println(b1)
	fmt.Println(b2)

	var s1 string // s1 := ""
	s2 := "Hallo Welt"
	s3 := s2[4]
	s4 := s2[2:7]
	s5 := s4 + s2

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Printf("%c\n", s3)
	fmt.Println(s4)
	fmt.Println(s5)

	for i := 0; i < len(s2); i++ {
		fmt.Printf("%c\n", s2[i])
	}

	for i, c := range s2 {
		fmt.Printf("%v: %c\n", i, c)
	}

	// Output:
}
