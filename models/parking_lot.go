package models

import (
	"container/heap"
	"errors"
	"fmt"
	"sort"
)

// ParkingLot menyimpan state dari tempat parkir
type ParkingLot struct {
	Capacity  int
	FreeSlots *IntHeap
	SlotToCar map[int]Car
	RegToSlot map[string]int
}

// NewParkingLot membuat instance baru ParkingLot
func NewParkingLot(cap int) *ParkingLot {
	return &ParkingLot{
		Capacity:  cap,
		FreeSlots: NewIntHeap(cap),
		SlotToCar: make(map[int]Car),
		RegToSlot: make(map[string]int),
	}
}

// Park menambahkan kendaraan ke slot kosong
// Mengembalikan nomor slot atau error jika penuh
func (p *ParkingLot) Park(reg, color string) (int, error) {
	if p == nil {
		return 0, errors.New("parking lot not created")
	}

	if p.FreeSlots.Len() == 0 {
		return 0, errors.New("full")
	}

	slot := heap.Pop(p.FreeSlots).(int)
	car := Car{RegNo: reg, Color: color}
	p.SlotToCar[slot] = car
	p.RegToSlot[reg] = slot
	return slot, nil
}

// Leave menghapus kendaraan berdasarkan nomor registrasi
// Menghitung biaya parkir sesuai lama waktu (hours) kendaraan terparkir
func (p *ParkingLot) Leave(reg string, hours int) (int, int, error) {
	if p == nil {
		return 0, 0, errors.New("parking lot not created")
	}
	slot, ok := p.RegToSlot[reg]
	if !ok {
		return 0, 0, errors.New(("Not Found"))
	}
	delete(p.RegToSlot, reg)
	delete(p.SlotToCar, slot)
	heap.Push(p.FreeSlots, slot)

	charge := 10
	if hours > 2 {
		charge += (hours - 2) * 10
	}
	return slot, charge, nil
}

// Status menampilkan daftar slot terisi
func (p *ParkingLot) Status() []string {
	if p == nil {
		return []string{"Parking ot not created"}
	}
	if len(p.SlotToCar) == 0 {
		return []string{"Slot No. Registration No."}
	}

	lines := []string{"Slot No. Registration No."}
	slots := make([]int, 0, len(p.SlotToCar))
	for s := range p.SlotToCar {
		slots = append(slots, s)
	}
	sort.Ints(slots)
	for _, s := range slots {
		car := p.SlotToCar[s]
		regColored := colorize(car.Color, car.RegNo)
		lines = append(lines, fmt.Sprintf("%d %s", s, regColored))
	}
	return lines
}
