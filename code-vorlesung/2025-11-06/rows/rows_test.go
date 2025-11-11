package rows

import (
	"fmt"
	"strings"
)

type Row []string

// New liefert eine neue Zeile der Länge `length` für
// ein Spielfeld und füllt sie mit `fill`.
func New(length int, fill string) Row {
	// Erzeugt eine Row (Liste von Strings) mit der angegebenen Länge.
	result := make(Row, length)

	//for i := 0; i < length; i++ {
	for i := range result {
		result[i] = fill
	}

	return result
}

// ContainsOnly erwartet eine Row und ein Zeichen.
// Liefert true, falls die Row nur das Zeichen enthält.
func (r Row) ContainsOnly(z string) bool {
	for i := 0; i < len(r); i++ {
		if r[i] != z {
			return false
		}
	}
	return true
}

func (r Row) String() string {
	return fmt.Sprintf("| %v |", strings.Join(r, " | "))
}

func ExampleRow_ContainsOnly() {
	r1 := New(3, "*")

	fmt.Println(r1)
	fmt.Println(r1.ContainsOnly("*"))

	r1[0] = "+"

	fmt.Println(r1)
	fmt.Println(r1.ContainsOnly("*"))

	// Output:
}
