package greet

import "fmt"

func Greet(name string) string {
	// return fmt.Sprintf("Hallo %s!", name)
	return fmt.Sprint("Hallo ", name)
}
