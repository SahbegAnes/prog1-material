package fizzbuzz

import "fmt"

// Fizz gibt die Zahlen von 1 bis 30 aus und ersetzt dabei
// jede durch 3 teilbare Zahl durch "fizz" und jede durch
// 5 teilbare Zahl durch "buzz".
// Bei Zahlen, die sowohl durch 3 als auch durch 5 teilbar sind,
// wird "fizzbuzz" ausgegeben.
func Fizz() {
	for i := 1; i <= 30; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
