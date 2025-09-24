Aplikasi **Parking Lot** berbasis **Command Line Interface (CLI)** menggunakan bahasa pemrograman **Golang**.
Project ini merupakan implementasi sederhana untuk mengelola parkiran, termasuk fitur membuat slot parkir, parkir kendaraan, keluar kendaraan dengan biaya, dan melihat status slot parkir.


---

## 📂 Struktur Proyek

```
.
├── pengontrol
│ └── parking_controller.go # Menangani perintah dan koneksi ke model & tampilan
│
├── model
│ ├── int_heap.go # Implementasi min-heap untuk slot yang tersedia
│ └── parking_lot.go # Tempat logika bisnis sistem
│
├── dilihat
│ └── parking_view.go # melibatkan tampilan CLI
│
├── main.go # Entri Poin dari aplikasi
├── go.mod # Definisi Go Module
└── go.sum # Kunci File Go
```

---

## ✨ Fitur

- Membuat tempat parkir dengan jumlah slot tertentu.
- Parkir kendaraan berdasarkan nomor registrasi.
- Mengeluarkan kendaraan berdasarkan nomor registrasi, sekaligus menghitung biaya parkir.
- Menampilkan status slot parkir (nomor slot dan nomor registrasi kendaraan).
- Mendukung interaksi langsung melalui CLI.

---

## ⚙️ Persyaratan

* [Go](https://golang.org/dl/) versi **1.20+**

---

## 📦 Instalasi & Menjalankan Aplikasi

1. Klon repositori ini:

```bash
git clone https://github.com/your-username/parking-lot.git
cd parking-lot
```

2. Jalankan Aplikasi:

```bash
go run main.go
```

---

## 💻 Penggunaan

Saat program dimulai, Anda akan melihat:

```text
Parking Lot (Enter 'exit' to quit):
>
```

Anda dapat memasukkan perintah berikut:

* **Buat tempat parkir dengan N slot**

```text
| [command] [params]
Contoh: create_parking_lot 6
```

* **Parkir mobil**

```text
| [command]
> park
Reg No: KA-01-HH-1234 Putih
Color: blue
```

* **Tinggalkan slot**

```teks
| [command] [parms] [params]
contoh: leave KA-01-HH-1234 4
```

* **Tampilkan status**

```teks
status
```

* **Keluar program**

```teks
exit
```

---

## 📖 Contoh Sesi

```teks
Parking Lot (Enter 'exit' to quit):
> create_parking_lot 6
Created a parking lot with 6 slots

> park
Reg No: KA-01-HH-1234
Color: White
Allocated slot number: 1

> park KA-01-HH-9999 Black
Allocated slot number: 2

> leave 1
Slot number 1 is free

> status
Slot No.    Registration No    Colour
2           KA-01-HH-9999      Black

> exit
Goodbye!
```

---

## 🛠 Catatan Pengembangan

* **IntHeap** (`int_heap.go`) digunakan untuk selalu mengalokasikan slot parkir terdekat yang tersedia.
* **ParkingController** mengelola perintah pengguna dan merutekannya ke model dan tampilan.
* Model **ParkingLot** berisi semua logika inti untuk mengelola mobil dan slot.
* **View** hanya bertanggung jawab untuk mencetak pesan.

---

## 📜 Lisensi

Proyek ini dilisensikan di bawah Lisensi MIT.