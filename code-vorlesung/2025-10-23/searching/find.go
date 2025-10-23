package searching

// / Liefert die Position von x in der Liste l,
// / oder liefert -1, falls x nicht in l vorkommt.
func Find(l []int, x int) int {
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			return i
		}
	}
	return -1
}
