package aufgabe1

// LongestAbc erwartet eine Liste von Strings und liefert
// das längste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
func LongestAbc(list []string) string {
	return findLongestABC(list, "")
}

func findLongestABC(list []string, longest string) string {
	if len(list) == 0 {
		return longest
	}
	for i := range list {
		if len(list[i]) >= 3 && list[i][:3] == "abc" {
			if len(list[i]) > len(longest) {
				longest = list[i]
			}
		}
	}
	return longest
}
