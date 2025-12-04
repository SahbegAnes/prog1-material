package aufgabe4

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// ElementSums erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils die Summe der beiden Elemente enthält.
//
// Annahmen für die Berechnung:
// Falls eine Liste kürzer ist als die andere, soll für die Berechnung der
// hinteren Werte ihr letztes Element verwendet werden.
// Für leere Listen soll für die Berechnung ggf. 0 verwendet werden.
func ElementSums(l1, l2 []int) []int {
	//Beispiel l1 = 5 und l3 =3
	result := []int{}
	var diff int
	var max int
	if len(l1) == len(l2) {
		for i := 0; i < len(l1); i++ {
			x := l1[i] + l2[i]
			result = append(result, x)
		}
		return result
	}

	if len(l1) < len(l2) {
		diff = len(l2) - len(l1)
		max = len(l1)
		for j := 0; j < max; j++ {
			y := l1[j] + l2[j]
			result = append(result, y)
		}
		for k := 0; k < diff; k++ {
			z := l1[max-1] + l2[max+k]
			result = append(result, z)
		}
	}

	if len(l1) > len(l2) {
		diff = len(l1) - len(l2)
		max = len(l2)

		for j := 0; j < max; j++ {
			y := l2[j] + l1[j]
			result = append(result, y)
		}
		for k := 0; k < diff; k++ {
			z := l2[max-1] + l1[max+k]
			result = append(result, z)
		}
	}
	// diff = 2
	//max = 3

	return result
}
