package rectangles

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Der Rand des Rechtecks soll aus `#`-Zeichen bestehen, der Innenraum soll leer sein.
func DrawEmptyRectangle(height, width int) {
	DrawRectangle(height, width, " ", "#")
}

// REMARKS
// - Diese Aufgabe ist eine etwas komplexere Variante der Aufgabe "Gefülltes Rechteck zeichnen".
//   Man geht dabei ähnlich vor, muss allerdings zusätzlich prüfen, ob man sich am Rand des Rechtecks befindet.
