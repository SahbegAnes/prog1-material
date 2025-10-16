package fizzbuzz

import "fmt"

// Fizz gibt die Zahlen von 1 bis n aus und ersetzt dabei
// jede durch x teilbare Zahl durch "fizz" und jede durch
// y teilbare Zahl durch "buzz".
// Bei Zahlen, die sowohl durch x als auch durch y teilbar sind,
// wird "fizzbuzz" ausgegeben.
func FizzImproved(x, y, n int) {
	for i := 1; i <= n; i++ {
		if i%y == 0 && i%x == 0 {
			fmt.Println("fizzbuzz")
		} else if i%x == 0 {
			fmt.Println("fizz")
		} else if i%y == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
