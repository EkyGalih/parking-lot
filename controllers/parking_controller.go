package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/EkyGalih/parking-lot/models"
	"github.com/EkyGalih/parking-lot/views"
)

// ParkingController bertugas menerima perintah (command) dari user
// dan menghubungkannya dengan model (ParkingLot) serta view
type ParkingController struct {
	Lot *models.ParkingLot
}


// HanndleCommand memproses input dari user
func (c *ParkingController) HandleCommand(line string) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return
	}
	cmd := strings.ToLower(parts[0])

	switch cmd {
	case "create_parking_lot":
		// Membuat tempat parkir baru dengan kapasitas tertentu
		if len(parts) < 2 {
			return
		}
		n, _ := strconv.Atoi(parts[1])
		c.Lot = models.NewParkingLot(n)
		views.Print(fmt.Sprintf("Created a parking lot with %d slots", n))
	case "park":
		// Menambahkan kendaraan ke slot kosong
		if len(parts) < 2 || c.Lot == nil {
			views.Print("Parking lot not created")
			return
		}
		slot, err := c.Lot.Park((parts[1]))
		if err != nil {
			if err.Error() == "full" {
				views.Print("Sorry, parking lot is full")
			}
			return
		}
		views.Print(fmt.Sprintf("Allocated slot number: %d", slot))
	case "leave":
		// Menghapus kendaraan dari slot berdasarkan nomor registrasi
		// dan menghitung biaya parkir sesuai lama waktu
		if len(parts) < 3 || c.Lot == nil {
			views.Print("Parking lot not created")
			return
		}
		hours, _ := strconv.Atoi(parts[2])
		slot, charge, err := c.Lot.Leave(parts[1], hours)
		if err != nil {
			views.Print(fmt.Sprintf("Registration number %s not found", parts[1]))
			return
		}
		views.Print(fmt.Sprintf(
			"Registration number %s with Slot number %d is free with Charge $%d",
			parts[1], slot, charge,
		))
	case "status":
		// Menampilkan daftar slot yang terisi
		if c.Lot == nil {
			views.Print("Parking lot not created")
			return
		}
		views.PrintLines(c.Lot.Status())
	}
}
