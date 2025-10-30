package arrays

import "fmt"

func Example_slice_reference() {
	// Vorab-Überlegung: Variablen werden als Werte in Funktionen kopiert.
	// Der Wert von `x` wird durch `Foo` nicht verändert, obwohl `Foo` innerhalb eine
	// Änderung macht.
	x := 42
	y := Foo(x)
	fmt.Println(x, y)

	// Bei Arrays und Slices ist das etwas anders.
	// Wir definieren ein Array (`arr`) und einen Ausschnitt daraus (`s1`):
	arr := [5]int{10, 20, 30, 40, 50}
	s1 := arr[1:4]

	fmt.Println(arr)
	fmt.Println(s1)

	// Die Funktion `Bar` verändert sowohl die Werte von `s1` als auch die von `arr`
	Bar(s1)

	// Beim Anhängen von Werten an `s1` wird eine Kopie von `arr` erzeugt,
	// weil der Platz in `arr` nicht ausreicht, um zwei neue Werte anzuhängen.
	// `s2` verwendet diese Kopie, `s1` aber nicht.
	s2 := append(s1, 555, 999)

	fmt.Println(arr)
	fmt.Println(s1)
	fmt.Println(s2)

	// Output:
}

func Foo(z int) int {
	z = z + 2
	return z
}

func Bar(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = s[i] * 2
	}
}
