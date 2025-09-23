package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/EkyGalih/parking-lot/controllers"
)


// entry point aplikasi
// Program berjalan dengan mode interaktif (CLI) untuk menerima perintah dari user
func main() {
	controller := &controllers.ParkingController{}

	fmt.Println("Parking Lot (Enter 'exit' to quit):")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Tampilkan prompt dan baca input user
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())

		// keluar jika user mengetik "exit"
		if strings.ToLower(line) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		// kirim perintah ke controller
		// dan eksekusi kode berdasarkan kategori perintah
		controller.HandleCommand(line)
	}

	// jika terjadi error saat membaca input
	if err := scanner.Err(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}