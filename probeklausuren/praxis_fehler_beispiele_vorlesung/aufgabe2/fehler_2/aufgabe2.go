package aufgabe2

// ExcludeBetween erwartet eine Liste und zwei Zahlen m und n.
// Die Funktion liefert eine Liste mit allen Elementen x, für die gilt: m < x < n.
func ExcludeBetween(list []int, m, n int) []int {
	result := []int{}                // result wird nie verändert!
	x := int{}                       // Sinn unklar
	for i := len(list); i < x; i++ { // Zählrichtung falsch?
		if m < x < n { // Unzulässige Bool-Formel
			return x // Early-Return macht keinen Sinn.
		}
	}
	return result
}
