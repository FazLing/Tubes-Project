package main

import (
	"fmt"
	"sort"
	"strings"
)

// Struct untuk menyimpan data pabrikan
type Pabrikan struct {
	Nama  string
	Motor []Motor
}

// Struct untuk menyimpan data motor
type Motor struct {
	Nama      string
	Penjualan int
}

var daftarPabrikan []Pabrikan

// Fungsi untuk menambahkan pabrikan
func tambahPabrikan(nama string) {
	daftarPabrikan = append(daftarPabrikan, Pabrikan{Nama: nama})
}

// Fungsi untuk menambahkan motor ke pabrikan
func tambahMotor(pabrikanIndex int, namaMotor string, penjualan int) {
	if pabrikanIndex >= 0 && pabrikanIndex < len(daftarPabrikan) {
		motor := Motor{Nama: namaMotor, Penjualan: penjualan}
		daftarPabrikan[pabrikanIndex].Motor = append(daftarPabrikan[pabrikanIndex].Motor, motor)
	} else {
		fmt.Println("Indeks pabrikan tidak valid.")
	}
}

// Fungsi untuk mengubah data motor
func ubahMotor(pabrikanIndex int, motorIndex int, namaMotor string, penjualan int) {
	if pabrikanIndex >= 0 && pabrikanIndex < len(daftarPabrikan) && motorIndex >= 0 && motorIndex < len(daftarPabrikan[pabrikanIndex].Motor) {
		daftarPabrikan[pabrikanIndex].Motor[motorIndex] = Motor{Nama: namaMotor, Penjualan: penjualan}
	} else {
		fmt.Println("Indeks pabrikan atau motor tidak valid.")
	}
}

// Fungsi untuk menghapus motor
func hapusMotor(pabrikanIndex int, motorIndex int) {
	if pabrikanIndex >= 0 && pabrikanIndex < len(daftarPabrikan) && motorIndex >= 0 && motorIndex < len(daftarPabrikan[pabrikanIndex].Motor) {
		daftarPabrikan[pabrikanIndex].Motor = append(daftarPabrikan[pabrikanIndex].Motor[:motorIndex], daftarPabrikan[pabrikanIndex].Motor[motorIndex+1:]...)
	} else {
		fmt.Println("Indeks pabrikan atau motor tidak valid.")
	}
}

// Fungsi untuk mencari motor
func cariMotor(namaMotor string) (*Pabrikan, *Motor) {
	for _, pabrikan := range daftarPabrikan {
		for _, motor := range pabrikan.Motor {
			if strings.EqualFold(motor.Nama, namaMotor) {
				return &pabrikan, &motor
			}
		}
	}
	return nil, nil
}

// Fungsi untuk menampilkan daftar motor berdasarkan pabrikan
func tampilkanMotorBerdasarkanPabrikan(namaPabrikan string) {
	for _, pabrikan := range daftarPabrikan {
		if strings.EqualFold(pabrikan.Nama, namaPabrikan) {
			fmt.Printf("Motor dari pabrikan %s:\n", pabrikan.Nama)
			for i, motor := range pabrikan.Motor {
				fmt.Printf("%d. Nama: %s, Penjualan: %d\n", i, motor.Nama, motor.Penjualan)
			}
			return
		}
	}
	fmt.Println("Pabrikan tidak ditemukan.")
}

// Fungsi untuk menampilkan daftar pabrikan dan jumlah motor
func tampilkanDaftarPabrikan() {
	fmt.Println("Daftar Pabrikan:")
	for i, pabrikan := range daftarPabrikan {
		fmt.Printf("%d. Nama: %s, Jumlah Motor: %d\n", i, pabrikan.Nama, len(pabrikan.Motor))
	}
}

// Fungsi untuk menampilkan 5 motor dengan penjualan tertinggi
func tampilkanMotorTerlaris() {
	var allMotors []Motor
	for _, pabrikan := range daftarPabrikan {
		allMotors = append(allMotors, pabrikan.Motor...)
	}
	sort.Slice(allMotors, func(i, j int) bool {
		return allMotors[i].Penjualan > allMotors[j].Penjualan
	})
	fmt.Println("5 Motor dengan Penjualan Tertinggi:")
	for i, motor := range allMotors {
		if i >= 5 {
			break
		}
		fmt.Printf("Nama: %s, Penjualan: %d\n", motor.Nama, motor.Penjualan)
	}
}

func main() {
	var choice int
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Pabrikan")
		fmt.Println("2. Tambah Motor")
		fmt.Println("3. Ubah Motor")
		fmt.Println("4. Hapus Motor")
		fmt.Println("5. Tampilkan Daftar Pabrikan")
		fmt.Println("6. Tampilkan Motor Berdasarkan Pabrikan")
		fmt.Println("7. Cari Motor")
		fmt.Println("8. Tampilkan 5 Motor Terlaris")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var namaPabrikan string
			fmt.Print("Masukkan nama pabrikan: ")
			fmt.Scanln(&namaPabrikan)
			tambahPabrikan(namaPabrikan)

		case 2:
			var pabrikanIndex int
			var namaMotor string
			var penjualan int
			tampilkanDaftarPabrikan()
			fmt.Print("Pilih pabrikan (index): ")
			fmt.Scanln(&pabrikanIndex)
			fmt.Print("Masukkan nama motor: ")
			fmt.Scanln(&namaMotor)
			fmt.Print("Masukkan jumlah penjualan: ")
			fmt.Scanln(&penjualan)
			tambahMotor(pabrikanIndex, namaMotor, penjualan)

		case 3:
			var pabrikanIndex, motorIndex, penjualan int
			var namaMotor string
			tampilkanDaftarPabrikan()
			fmt.Print("Pilih pabrikan (index): ")
			fmt.Scanln(&pabrikanIndex)
			tampilkanMotorBerdasarkanPabrikan(daftarPabrikan[pabrikanIndex].Nama)
			fmt.Print("Pilih motor (index): ")
			fmt.Scanln(&motorIndex)
			fmt.Print("Masukkan nama motor baru: ")
			fmt.Scanln(&namaMotor)
			fmt.Print("Masukkan jumlah penjualan baru: ")
			fmt.Scanln(&penjualan)
			ubahMotor(pabrikanIndex, motorIndex, namaMotor, penjualan)

		case 4:
			var pabrikanIndex, motorIndex int
			tampilkanDaftarPabrikan()
			fmt.Print("Pilih pabrikan (index): ")
			fmt.Scanln(&pabrikanIndex)
			tampilkanMotorBerdasarkanPabrikan(daftarPabrikan[pabrikanIndex].Nama)
			fmt.Print("Pilih motor (index): ")
			fmt.Scanln(&motorIndex)
			hapusMotor(pabrikanIndex, motorIndex)

		case 5:
			tampilkanDaftarPabrikan()

		case 6:
			var namaPabrikan string
			fmt.Print("Masukkan nama pabrikan: ")
			fmt.Scanln(&namaPabrikan)
			tampilkanMotorBerdasarkanPabrikan(namaPabrikan)

		case 7:
			var namaMotor string
			fmt.Print("Masukkan nama motor: ")
			fmt.Scanln(&namaMotor)
			pabrikan, motor := cariMotor(namaMotor)
			if motor != nil {
				fmt.Printf("Motor %s ditemukan dari pabrikan %s dengan penjualan %d\n", motor.Nama, pabrikan.Nama, motor.Penjualan)
			} else {
				fmt.Println("Motor tidak ditemukan.")
			}

		case 8:
			tampilkanMotorTerlaris()

		case 9:
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
