package rectangles

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Das Rechteck soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidRectangle(height, width int) {
	DrawRectangle(height, width, "#", "#")
}
