package hanoi

import "fmt"

func Hanoi0(start, mitte, ziel string) {

}

func Hanoi1(start, mitte, ziel string) {
	Hanoi0(start, ziel, mitte)
	Move(start, ziel)
	Hanoi0(mitte, start, ziel)
}

func Hanoi2(start, mitte, ziel string) {
	Hanoi1(start, ziel, mitte)
	Move(start, ziel)
	Hanoi1(mitte, start, ziel)
}

func Hanoi3(start, mitte, ziel string) {
	Hanoi2(start, ziel, mitte)
	Move(start, ziel)
	Hanoi2(mitte, start, ziel)
}

func Move(start, ziel string) {
	fmt.Printf("%s -> %s\n", start, ziel)
}

func Hanoi(height int, start, mitte, ziel string) {
	if height == 0 {
		return
	}
	Hanoi(height-1, start, ziel, mitte)
	Move(start, ziel)
	Hanoi(height-1, mitte, start, ziel)
}
