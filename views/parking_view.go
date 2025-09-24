package views

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

// untuk membaca input user
func Input(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}