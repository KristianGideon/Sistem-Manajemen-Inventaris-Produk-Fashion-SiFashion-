package main

import "fmt"

type Product struct {
	ID       int
	Name     string
	Category string
	Size     string
	Color    string
	Price    float64
	Stock    int
}

var daftarProduk []Product
var idBerikutnya int = 1

func bacaBaris() string {
	var hasil string
	for {
		var huruf byte
		fmt.Scanf("%c", &huruf)
		if huruf == '\n' {
			return hasil
		}
		if huruf != '\r' {
			hasil += string(huruf)
		}
	}
}

func adaSubstring(teks, cari string) bool {
	if len(cari) == 0 {
		return true
	}
	if len(teks) < len(cari) {
		return false
	}
	for i := 0; i <= len(teks)-len(cari); i++ {
		cocok := true
		for j := 0; j < len(cari); j++ {
			a := teks[i+j]
			b := cari[j]
			if a >= 'A' && a <= 'Z' {
				a += 32
			}
			if b >= 'A' && b <= 'Z' {
				b += 32
			}
			if a != b {
				cocok = false
				break
			}
		}
		if cocok {
			return true
		}
	}
	return false
}

func tambahProduk() {
	var p Product
	p.ID = idBerikutnya
	idBerikutnya++
	fmt.Print("Nama: ")
	p.Name = bacaBaris()
	fmt.Print("Kategori: ")
	p.Category = bacaBaris()
	fmt.Print("Ukuran: ")
	p.Size = bacaBaris()
	fmt.Print("Warna: ")
	p.Color = bacaBaris()
	fmt.Print("Harga: ")
	hargaStr := bacaBaris()
	fmt.Sscanf(hargaStr, "%f", &p.Price)
	fmt.Print("Stok: ")
	stokStr := bacaBaris()
	fmt.Sscanf(stokStr, "%d", &p.Stock)
	daftarProduk = append(daftarProduk, p)
	fmt.Println("Produk berhasil ditambahkan.")
}

func lihatProduk() {
	if len(daftarProduk) == 0 {
		fmt.Println("Tidak ada produk.")
		return
	}
	for i := 0; i < len(daftarProduk); i++ {
		p := daftarProduk[i]
		fmt.Printf("ID: %d, Nama: %s, Kategori: %s, Ukuran: %s, Warna: %s, Harga: %.2f, Stok: %d\n",
			p.ID, p.Name, p.Category, p.Size, p.Color, p.Price, p.Stock)
	}
}

func ubahProduk() {
	fmt.Print("Masukkan ID produk yang akan diubah: ")
	idStr := bacaBaris()
	var id int
	fmt.Sscanf(idStr, "%d", &id)
	posisi := -1
	for i := 0; i < len(daftarProduk); i++ {
		if daftarProduk[i].ID == id {
			posisi = i
			break
		}
	}
	if posisi == -1 {
		fmt.Println("Produk tidak ditemukan.")
		return
	}
	fmt.Print("Nama baru (kosongkan untuk tidak mengubah): ")
	nama := bacaBaris()
	if nama != "" {
		daftarProduk[posisi].Name = nama
	}
	fmt.Print("Kategori baru (kosongkan untuk tidak mengubah): ")
	kategori := bacaBaris()
	if kategori != "" {
		daftarProduk[posisi].Category = kategori
	}
	fmt.Print("Ukuran baru (kosongkan untuk tidak mengubah): ")
	ukuran := bacaBaris()
	if ukuran != "" {
		daftarProduk[posisi].Size = ukuran
	}
	fmt.Print("Warna baru (kosongkan untuk tidak mengubah): ")
	warna := bacaBaris()
	if warna != "" {
		daftarProduk[posisi].Color = warna
	}
	fmt.Print("Harga baru (kosongkan untuk tidak mengubah): ")
	hargaStr := bacaBaris()
	if hargaStr != "" {
		var harga float64
		fmt.Sscanf(hargaStr, "%f", &harga)
		daftarProduk[posisi].Price = harga
	}
	fmt.Print("Stok baru (kosongkan untuk tidak mengubah): ")
	stokStr := bacaBaris()
	if stokStr != "" {
		var stok int
		fmt.Sscanf(stokStr, "%d", &stok)
		daftarProduk[posisi].Stock = stok
	}
	fmt.Println("Produk berhasil diubah.")
}

func hapusProduk() {
	fmt.Print("Masukkan ID produk yang akan dihapus: ")
	idStr := bacaBaris()
	var id int
	fmt.Sscanf(idStr, "%d", &id)
	posisi := -1
	for i := 0; i < len(daftarProduk); i++ {
		if daftarProduk[i].ID == id {
			posisi = i
			break
		}
	}
	if posisi == -1 {
		fmt.Println("Produk tidak ditemukan.")
		return
	}
	var sisaProduk []Product
	for i := 0; i < len(daftarProduk); i++ {
		if i != posisi {
			sisaProduk = append(sisaProduk, daftarProduk[i])
		}
	}
	daftarProduk = sisaProduk
	fmt.Println("Produk berhasil dihapus.")
}

func cariProduk() {
	fmt.Print("Masukkan kata kunci pencarian (nama/kategori): ")
	katakunci := bacaBaris()
	ketemu := false
	for i := 0; i < len(daftarProduk); i++ {
		if adaSubstring(daftarProduk[i].Name, katakunci) || adaSubstring(daftarProduk[i].Category, katakunci) {
			p := daftarProduk[i]
			fmt.Printf("ID: %d, Nama: %s, Kategori: %s, Ukuran: %s, Warna: %s, Harga: %.2f, Stok: %d\n",
				p.ID, p.Name, p.Category, p.Size, p.Color, p.Price, p.Stock)
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Produk tidak ditemukan.")
	}
}

func main() {
	for {
		fmt.Println("=== SiFashion - Manajemen Produk ===")
		fmt.Println("1. Tambah Produk")
		fmt.Println("2. Lihat Semua Produk")
		fmt.Println("3. Ubah Produk")
		fmt.Println("4. Hapus Produk")
		fmt.Println("5. Cari Produk")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		pilihan := bacaBaris()
		if pilihan == "0" {
			fmt.Println("")
			break
		} else if pilihan == "1" {
			tambahProduk()
		} else if pilihan == "2" {
			lihatProduk()
		} else if pilihan == "3" {
			ubahProduk()
		} else if pilihan == "4" {
			hapusProduk()
		} else if pilihan == "5" {
			cariProduk()
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
