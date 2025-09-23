package views

import "fmt"

// Print menampilkan satu baris pesan ke layar
func Print(msg string) {
	fmt.Println(msg)
}

// menampilkan beberapa baris pesan ke layar
func PrintLines(lines []string) {
	for _, l := range lines {
		fmt.Println(l)
	}
}