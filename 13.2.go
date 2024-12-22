package main

import "fmt"

const nMax = 7919

type Buku struct {
	id        string
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}
type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	fmt.Println("\n PENDAFTARAN BUKU ")
	fmt.Println("Silahkan masukkan data", n, "buku:")
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("\nBuku ke-%d:\n", i+1)
		fmt.Print("Masukkan ID Buku (contoh: B001): ")
		fmt.Scan(&buku.id)
		fmt.Print("Masukkan Judul Buku (gunakan' _'untuk spasi): ")
		fmt.Scan(&buku.judul)
		fmt.Print("Masukkan Nama Penulis (gunakan '_' untuk spasi): ")
		fmt.Scan(&buku.penulis)
		fmt.Print("Masukkan Nama Penerbit (gunakan '_' untuk spasi): ")
		fmt.Scan(&buku.penerbit)
		fmt.Print("Masukkan Jumlah Eksemplar: ")
		fmt.Scan(&buku.eksemplar)
		fmt.Print("Masukkan Tahun Terbit: ")
		fmt.Scan(&buku.tahun)
		fmt.Print("Masukkan Rating (0-100): ")
		fmt.Scan(&buku.rating)
		pustaka[i] = buku
	}
	fmt.Println("\nPendaftaran buku selesai")
}
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	fmt.Println("\n BUKU TERFAVORIT ")
	maxIdx := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[maxIdx].rating {
			maxIdx = i
		}
	}
	buku := pustaka[maxIdx]
	fmt.Printf("Buku terfavorit : %s %s %s %d\n",
		buku.judul, buku.penulis, buku.penerbit, buku.tahun)
}
func UrutBuku(pustaka *DaftarBuku, n int) {
	fmt.Println("\n PENGURUTAN BUKU ")
	fmt.Println("Mengurutkan buku berdasarkan rating...")
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
	fmt.Println("Pengurutan selesai!")
}
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("\n 5 BUKU RATING TERTINGGI ")
	count := 5
	if n < 5 {
		count = n
		fmt.Printf("Menampilkan %d buku (karena total buku kurang dari 5):\n", n)
	} else {
		fmt.Println("Menampilkan 5 buku rating tertinggi:")
	}
	for i := 0; i < count; i++ {
		fmt.Printf("%d. %s\n", i+1, pustaka[i].judul)
	}
}
func CariBuku(pustaka DaftarBuku, n int, r int) {
	fmt.Printf("\n PENCARIAN BUKU DENGAN RATING %d \n", r)
	left := 0
	right := n - 1
	found := false
	for left <= right && !found {
		mid := (left + right) / 2
		if pustaka[mid].rating == r {
			buku := pustaka[mid]
			fmt.Println("Buku ditemukan:")
			fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nEksemplar: %d\nRating: %d\n",
				buku.judul, buku.penulis, buku.penerbit,
				buku.tahun, buku.eksemplar, buku.rating)
			found = true
		} else if pustaka[mid].rating < r {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating itu")
	}
}
func main() {
	var n, ratingCari int
	var pustaka DaftarBuku
	fmt.Println("PROGRAM MANAJEMEN PERPUSTAKAAN")
	fmt.Print("Masukkan jumlah buku yang ingin didaftarkan: ")
	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)
	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	fmt.Print("\nMasukkan rating buku yang dicari: ")
	fmt.Scan(&ratingCari)
	CariBuku(pustaka, n, ratingCari)
	fmt.Println("\n PROGRAM KELAR")
}
