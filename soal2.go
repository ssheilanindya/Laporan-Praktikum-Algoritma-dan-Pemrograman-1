package main

import "fmt"

func main() {
	var nama, nim, kelas string

	fmt.Println("Masukkan data mahasiswa :")
	fmt.Print("Nama :")
	fmt.Scanln(&nama)
	fmt.Print("Nim :")
	fmt.Scanln(&nim)
	fmt.Print("Kelas :")
	fmt.Scanln(&kelas)

	fmt.Println("\nBerikut resume singkat mahasiswa Tel-U Purwokerto")
	fmt.Printf("Perkenalkan nama saya %s, saya dari Prodi S1-IF, dengan NIM %s, saya dari kelas %s", nama, nim, kelas)

}
