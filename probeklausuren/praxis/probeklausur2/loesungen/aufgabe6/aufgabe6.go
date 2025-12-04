package aufgabe6

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// DuplicateSinglets erwartet eine int-Liste list.
// Die Funktion liefert eine int-Liste, bei der alle Elemente,
// die in list nur einmal vorkommen, verdoppelt sind,
// also zwei Mal hintereinander stehen.
// Elemente, die schon in list mehrfach vorkommen, sollen wie sie sind
// ins Ergebnis übertragen werden.
func DuplicateSinglets(list []int) []int {
	result := []int{}
	for _, el := range list {
		result = append(result, el)
		if Count(list, el) == 1 {
			result = append(result, el)
		}
	}
	return result
}

func Count(list []int, x int) int {
	result := 0
	for _, el := range list {
		if el == x {
			result++
		}
	}
	return result
}
