package length

type Length int

// FromMillimeters konstruiert eine Länge aus einer Millimeter-Angabe.
func FromMillimeters(m int) Length {
	return Length(m)
}

// FromCentimeters konstruiert eine Länge aus einer Zentimeter-Angabe.
func FromCentimeters(c int) Length {
	return Length(c * 10)
}

// FromMeters konstruiert eine Länge aus einer Meter-Angabe.
func FromMeters(m int) Length {
	return Length(m * 10 * 100)
}

// FromKilometers konstruiert eine Länge aus einer Kilometer-Angabe.
func FromKilometers(k int) Length {
	return Length(k * 10 * 100 * 1000)
}

// Millimeters liefert die Länge in Millimetern.
func (l Length) Millimeters() int {
	return int(l)
}

// Centimeters liefert die Länge in Zentimetern.
func (l Length) Centimeters() int {
	return l.Millimeters() / 10
}

// Meters liefert die Länge in Metern
func (l Length) Meters() int {
	return l.Centimeters() / 100
}

// Kilometers liefert die Länge in Metern
func (l Length) Kilometers() int {
	return l.Meters() / 1000
}
