# Material zur Vorlesung Programmierung 1

Dieses Repository enthält Materialien zur Vorlesung _Programmierung I_
im Kurs WWI25AMA an der DHBW Mannheim.

## Lern-Ressourcen

Die folgenden Links bieten gute Einstiegs- und Nachschlage-Ressourcen für die Programmiersprache Go:

- [Go by Example](https://gobyexample.com) - Eine Sammlung von kleinen, gut erklärten Code-Beispielen zu vielen Themen.
- [A Tour of Go](https://tour.golang.org) - Ein interaktives Tutorial zu den wichtigsten Go-Konzepten.
- [Exercism](https://exercism.org) - Eine Sammlung von Programmieraufgaben mit automatischer Überprüfung.
- [CodinGame](https://www.codingame.com) - Eine Sammlung von Programmieraufgaben in Form von Spielen.
- [Project Euler](https://projecteuler.net) - Eine Sammlung von mathematisch orientierten Programmieraufgaben.

## Software

In der Vorlesung wird die Programmiersprache [Go](https://go.dev) verwendet.
Dies ist eine relativ junge, freie Programmiersprache, die dank einer einfachen
Entwicklungs-Toolchain und einer guten Standardbibliothek einen leichten Einstieg
und kurze Entwicklungszyklen ermöglicht.
Gleichzeitig ist Go hochgradig praxisrelevant:
Die Sprache wird u.A. im Web- bzw. Cloud-Umfeld oder z.B. für
service-orientierte Softwaresysteme verwendet.

Zusätzlich zur Programmiersprache werden wir in der Vorlesung auch
[Visual Studio Code](https://code.visualstudio.com) (VsCode) als Editor
und [Git](https://git-scm.com) als Versionsverwaltung verwenden.

### Lokale Verwendung auf dem eigenen Rechner

Um Go-Programme auf dem eigenen Rechner zu schreiben und auszuführen, sollten
Sie die folgenden Software-Pakete installieren:

- [Go](https://go.dev) (Compiler und Toolchain)
- [Visual Studio Code](https://code.visualstudio.com) (Editor)
- [Git](https://git-scm.com) (Versionsverwaltung)

Dies sind alles freie Software-Pakete, die auf Windows, Linux und MacOS
verfügbar sind. Alternativ kann auch [GoLand](https://www.jetbrains.com/de-de/go) als
Editor verwendet werden (s.u.).

### Software-Übersicht

| Software | Download-Link | Erläuterung
| --- | --- | ---
| Go-[Compiler](https://de.wikipedia.org/wiki/Compiler) | <https://go.dev>                     | Das wichtigste Werkzeug beim Programmieren: Übersetzt den [Quellcode](https://de.wikipedia.org/wiki/Quelltext) in eine [ausführbare Datei](https://de.wikipedia.org/wiki/Ausf%C3%BChrbare_Datei).
| Visual Studio Code                                    | <https://code.visualstudio.com>      | [Code-Editor](https://de.wikipedia.org/wiki/Editor_(Software)) bzw. [Entwicklungsumgebung](https://de.wikipedia.org/wiki/Integrierte_Entwicklungsumgebung) zum Schreiben von Quelltext.
| GoLand                                                | <https://www.jetbrains.com/de-de/go> | Eine Alternative zu VsCode. Gut integriert und einfach einzusetzen, "batteries included". Für Studierende kostenlos.
| Git                                                   | <https://git-scm.com>                | Eine [Versionsverwaltung](https://de.wikipedia.org/wiki/Versionsverwaltung) für Quellcode.
| Go-Playground                                         | <https://go.dev/play>                | Ein einfacher Online-Editor und -Compiler für Go-Code. Gut zum schnellen Ausprobieren.
| GitHub Codespaces                                                | Jeweils bei den Material-Repos                 | Eine fertige Online-Entwicklungsumgebung für jedes Repo zur Vorlesung.

Dies alles sollte auf jedem Windows- oder Linux-PC und auch auf jedem Mac
ohne Admin-Rechte installierbar sein.
Falls es dabei dennoch Probleme gibt, ist auch die Verwendung eines webbasierten
Programmiersystems denkbar (z.B. [Go Playground](https://go.dev/play) oder [Codespaces](https://github.com/features/codespaces)).

## Bemerkungen zur Software-Verwendung und -Installation

### VsCode

In der Vorlesung und Prüfung wird VsCode als Editor verwendet. Wer möglichst ähnliche Programmierumgebung bei sich einrichten will,
sollte sich also den Go-Compiler von <https://go.dev> sowie VsCode installieren.
_Wichtig dabei_: Sobald zum ersten Mal eine Go-Quelldatei mit VsCode geöffnet wird, wird zuerst die Installation einer _Go-Erweiterung_
und anschließend noch die Installation diverser Tools angeboten. Dies sollte alles angenommen werden.
VsCode installiert aber nicht selbständig den Go-Compiler, dieser muss vorab selbständig installiert werden.

### GoLand

Gegen die Verwendung von GoLand ist grundsätzlich nichts einzuwenden, dies ist ebenfalls eine sehr gute Entwicklungsumgebung.
Die Installation ist etwas einfacher, die Umgebung ist stärker automatisiert. Die Bedienung ist vom Konzept her ähnlich wie bei VsCode,
unterscheidet sich aber in Details.
Als Vorbereitung für die Prüfung ist es daher ratsam, VsCode wenigstens einmal ausprobiert zu haben.
Generell ist GoLand aber eine sehr gute Wahl.

### GitHub Codespaces

Mit [GitHub Codespaces](https://github.com/features/codespaces) kann auch vollwertig gearbeitet werden,
andere webbasierte IDEs wie z.B. [GitPod](https://www.gitpod.io) bieten ähnliche Funktionalitäten.
Um Codespaces zu verwenden, ist ein GitHub-Account notwendig und es empfiehlt sich,
diesen mit einer DHBW-E-Mail-Adresse zu erstellen, um die kostenlose Nutzung
über das Education-Programm von GitHub zu ermöglichen.

### Go-Playground

Der Go-Playground bietet eine sehr einfache und leicht zugängliche Möglichkeit,
um Go-Code auszuprobieren. Er benötigt keine Registrierung und kann direkt im Browser
verwendet werden, allerdings kann man dort keinen Code speichern (bzw. nur sehr eingeschränkt).

### Git

[Git](https://git-scm.com) ist eine Versionsverwaltung für Quellcode, deren Hauptzweck es ist, gemeinsam  an Quellcode zu arbeiten und sich dabei
möglichst nicht gegenseitig mit Änderungen zu stören.
Services wie [GitHub](https://github.com) setzen darauf auf und bieten einen zentralen Speicherort für Quellcode.
Dadurch ist Git mittlerweile zum weltweiten Quasi-Standard für die Verwaltung von Software-Quellcode geworden.

Sämtlicher Code für diese Vorlesung wird über GitHub bereitgestellt.
Die Verwendung von Git ist dabei nicht zwingend notwendig und nicht Kernbestandteil der Vorlesung.
Sich mit den einfachsten Git-Kommandos und Arbeitsweisen vertraut zu machen, vereinfacht aber auch die Arbeit mit dem Code aus der Vorlesung.
Beispiele dafür werden in der Vorlesung besprochen.
Insbesondere bei der Arbeit mit GitHub-Codespaces ist die Verwendung von Git unerlässlich.

Neben dieser allgemeinen Anwendung wird Git auch von Go selbst als Werkzeug für die Verwaltung und Bereitstellung von
Paketen genutzt. Ob dies in der Vorlesung von Relevanz sein wird, ist jetzt am Anfang noch nicht klar.
Zu einer funktionierenden und vollständigen Go-Installation gehört Git dadurch aber auf jeden Fall dazu und sollte bei Ihnen installiert sein.
