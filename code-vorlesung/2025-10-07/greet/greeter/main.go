package main

import (
	"fmt"
	"prog1-material/code-vorlesung/2025-10-07/greet"
)

func main() {
	fmt.Println(greet.Greet("Kurs"))
	fmt.Println("Alle " + greet.Greet("Kursteilnehmer"))
}
