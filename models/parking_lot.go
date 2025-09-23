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
	SlotToReg map[int]string
	RegToSlot map[string]int
}

// NewParkingLot membuat instance baru ParkingLot
func NewParkingLot(cap int) *ParkingLot {
	return &ParkingLot{
		Capacity: cap,
		FreeSlots: NewIntHeap(cap),
		SlotToReg: make(map[int]string),
		RegToSlot: make(map[string]int),
	}
}

// Park menambahkan kendaraan ke slot kosong
// Mengembalikan nomor slot atau error jika penuh
func (p *ParkingLot) Park(reg string) (int, error) {
	if p == nil {
		return 0, errors.New("parking lot not created")
	}

	if p.FreeSlots.Len() == 0 {
		return 0, errors.New("full")
	}
	slot := heap.Pop((p.FreeSlots)).(int)
	p.SlotToReg[slot] = reg
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
	delete(p.SlotToReg, slot)
	heap.Push(p.FreeSlots, slot)

	charge := 10
	if hours > 2 {
		charge += (hours-2) * 10
	}
	return slot, charge, nil
}

// Status menampilkan daftar slot terisi
func (p *ParkingLot) Status() []string {
	if p == nil {
return []string{"Parking ot not created"}
	}
	if len(p.SlotToReg) == 0 {
return []string{"Slot No. Registration No."}
	}

lines := []string{"Slot No. Registration No."}
	slots := make([]int, 0, len(p.SlotToReg))
	for s := range p.SlotToReg {
		slots = append(slots, s)
	}
	sort.Ints(slots)
	for _, s := range slots {
		lines = append(lines,
			stringifSlot(s, p.SlotToReg[s]),
		)
	}
	return lines
}

// fungsi ini dibuat untuk memudahkan debugging dan format output status
func stringifSlot(slot int, reg string) string {
	return  fmt.Sprintf("%d %s", slot, reg)
}