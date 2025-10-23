package power

// Pow2 berechnet "2 hoch n".
func Pow2(n int) int {
	// 2^0 == 1
	if n == 0 {
		return 1
	}
	// 2^n = 2^(n-1) * 2
	result := Pow2(n-1) * 2
	return result
}
