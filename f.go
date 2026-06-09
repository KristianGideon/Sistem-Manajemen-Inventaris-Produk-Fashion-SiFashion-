package main

import "fmt"

type akun struct {
	username string
	password string
}

var idMul int = 10

var ak = [5]akun{{username: "admin", password: "admin"}, {username: "user", password: "user"}, {username: "1", password: "1"}}

type Product struct {
	Id       int
	Nama     string
	kategori string
	Ukuran   string
	Warna    string
	Harga    float64
	Stok     int
}

var produk = [100]Product{{1, "Kaos Polos", "Atasan", "L", "Hitam", 75000, 50}, {2, "Celana Chino", "Bawah", "32", "Khaki", 150000, 30}, {3, "Topi Snapback", "Aksesoris", "One Size", "Merah", 90000, 25}, {4, "Boxer Brief", "Pakaian Dalam", "M", "Abu-abu", 45000, 100}, {5, "Celana Training", "Pakaian Olahraga", "XL", "Biru Navy", 120000, 40}, {6, "Jaket Hoodie", "Outerwear", "L", "Hitam", 200000, 20}, {7, "Kemeja Putih", "Pakaian Formal", "42", "Putih", 180000, 15}, {8, "Gelang Tali", "Aksesoris", "One Size", "Coklat", 25000, 75}, {9, "Rok Lipit", "Bawah", "S", "Hitam", 110000, 35}, {10, "Rompi Wanita", "Atasan", "M", "Krem", 85000, 28}}

func menu() {
	clear()
	var pilihan int
	fmt.Println(bor(15), "Menu Utama", bor(15))
	fmt.Println("1. Manajemen Produk")
	fmt.Println("2. Manajemen Stok  ")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih Menu Dipilih : ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		manajemenProduk()
	case 2:
		manajemenStock()
	case 3:
		fmt.Println("Terimakasih sudah menggunakan sistem manajemen fashion")
		return
	default:
		fmt.Println("Menu tidak tersedia")
	}
}
func manajemenProduk() {
	clear()
	var pilihan int
	fmt.Println(bor(15), "Manajemen Produk", bor(15))
	fmt.Println("1. Tambah Produk")
	fmt.Println("2. Edit Produk")
	fmt.Println("3. Hapus Produk")
	fmt.Println("4. Kembali")
	fmt.Print("Pilih Menu Dipilih : ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahProduk()
	case 2:
		editProduk()
	case 3:
		hapusProduk()
	case 4:
		menu()
	default:
		fmt.Println("Menu tidak tersedia")
	}
}

func tambahProduk() {
	clear()
	var a int
	var pil string
	if idMul >= len(produk)-1 {
		fmt.Println("Penyimpanan produk penuh!")
		return
	}
	idMul++
	produk[idMul].Id = idMul
	fmt.Print("Masukan Nama Produk : ")
	fmt.Scan(&produk[idMul].Nama)
	fmt.Printf("Pilihan Kategori : \n1. Atasan\n2. Bawahan\n3. Aksesoris\n4. Pakaian Dalam\n5. Pakaian Olahraga\n6. Outerwear\n7. Pakaian Formal\n8. Aksesoris\n")
	fmt.Print("Masukan Kategori Produk : ")
	fmt.Scan(&a)
	switch a {
	case 1:
		produk[idMul].kategori = "Atasan"
	case 2:
		produk[idMul].kategori = "Bawahan"
	case 3:
		produk[idMul].kategori = "Aksesoris"
	case 4:
		produk[idMul].kategori = "Pakaian Dalam"
	case 5:
		produk[idMul].kategori = "Pakaian Olahraga"
	case 6:
		produk[idMul].kategori = "Outerwear"
	case 7:
		produk[idMul].kategori = "Pakaian Formal"
	case 8:
		produk[idMul].kategori = "Aksesoris"
	default:
		fmt.Println("Kategori tidak tersedia")
		return
	}
	fmt.Print("Masukan Ukuran Produk : ")
	fmt.Scan(&produk[idMul].Ukuran)
	fmt.Print("Masukan Warna Produk : ")
	fmt.Scan(&produk[idMul].Warna)
	fmt.Print("Masukan Harga Produk : ")
	fmt.Scan(&produk[idMul].Harga)
	fmt.Print("Masukan Stok Produk : ")
	fmt.Scan(&produk[idMul].Stok)
	fmt.Println("Produk berhasil ditambahkan dengan ID:", idMul)
	fmt.Print("Apakah Anda ingin menambahkan produk lagi? (y/n): ")
	fmt.Scan(&pil)
	if pil == "y" || pil == "Y" {
		tambahProduk()
	} else {
		manajemenProduk()
	}
}
func editProduk() {
	clear()
	fmt.Println(bor(15), "Edit Produk", bor(15))
	fmt.Println("Cari produk yang akan diedit:")
	tampil()

	var id int
	fmt.Print("Masukkan ID produk yang akan diedit: ")
	fmt.Scan(&id)

	idx := -1
	for i := 0; i < idMul; i++ {
		if produk[i].Id == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println("ID tidak ditemukan.")
	} else {
		fmt.Printf("Data lama: %s | %s | Rp%.2f\n", produk[idx].Nama, produk[idx].kategori, produk[idx].Harga)
		fmt.Print("Nama baru: ")
		fmt.Scan(&produk[idx].Nama)
		fmt.Print("Kategori Baru: ")
		fmt.Scan(&produk[idx].kategori)
		fmt.Print("Harga Baru: ")
		fmt.Scan(&produk[idx].Harga)
		fmt.Println("Produk berhasil diperbarui.")
	}

	fmt.Println("Tekan Enter untuk kembali...")
	fmt.Scanln()
	manajemenProduk()
}

func hapusProduk() {
	clear()
	var pilihan string
	fmt.Println(bor(15), "Hapus Produk", bor(15))
	fmt.Print("Cari Berdasarkan Nama Produk Yang Mau Dihapus: ")
	tampil()
	fmt.Println("Pencarian Lagi (y/n) ?")
	fmt.Scan(&pilihan)
	if pilihan == "y" || pilihan == "Y" {
		hapusProduk()
	} else {
		fmt.Println("Kembali ke Menu Manajemen Produk / Hapus Produk")
		fmt.Println("1. Kembali ke Menu Manajemen Produk")
		fmt.Println("2. Hapus Produk Berdasarkan ID")
		fmt.Print("Pilih Menu Dipilih : ")
		var pil int
		fmt.Scan(&pil)
		switch pil {
		case 1:
			manajemenProduk()
		case 2:
			var id int
			fmt.Print("Masukkan ID produk yang akan dihapus: ")
			fmt.Scan(&id)
			idx := -1
			for i := 0; i < idMul; i++ {
				if produk[i].Id == id {
					idx = i
					break
				}
			}
			if idx != -1 {
				for i := idx; i < idMul; i++ {
					produk[i] = produk[i+1]
				}
				idMul--
				fmt.Println("Produk berhasil dihapus.")
			} else {
				fmt.Println("Produk tidak ditemukan.")
			}
			manajemenProduk()
		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}
func manajemenStock() {
	clear()
	fmt.Println(bor(15), "Manajemen Stok", bor(15))
	var pil int
	fmt.Println("1. Cek Stok")
	fmt.Println("2. Tambah Stok")
	fmt.Println("3. Hapus Stok")
	fmt.Println("4. Kembali")
	fmt.Print("Pilih Menu Dipilih : ")
	fmt.Scan(&pil)
	switch pil {
	case 1:
		cekStok()
	case 2:
		tambahStock()
	case 3:
		hapusStock()
	case 4:
		menu()
	default:
		fmt.Println("Menu tidak tersedia")
	}
}
func tambahStock() {
	clear()
	fmt.Println(bor(15), "Tambah Stock", bor(15))
	var id int
	fmt.Print("Masukkan ID produk yang akan ditambahkan stok: ")
	fmt.Scan(&id)
	idx := -1
	for i := 0; i < idMul; i++ {
		if produk[i].Id == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("ID tidak ditemukan.")
	} else {
		fmt.Printf("Stok lama: %d\n", produk[idx].Stok)
		var jumlah int
		fmt.Print("Masukkan jumlah stok yang akan ditambahkan: ")
		fmt.Scan(&jumlah)
		produk[idx].Stok += jumlah
		fmt.Println("Stok berhasil ditambahkan.")
	}
	fmt.Println("Tekan Enter untuk kembali...")
	fmt.Scanln()
	manajemenStock()
}
func hapusStock() {
	clear()
	fmt.Println(bor(15), "Hapus Stok", bor(15))
	var id int
	fmt.Print("Masukkan ID produk yang akan dihapus stoknya: ")
	fmt.Scan(&id)
	idx := -1
	for i := 0; i < idMul; i++ {
		if produk[i].Id == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("ID tidak ditemukan.")
	} else {
		fmt.Printf("Stok lama: %d\n", produk[idx].Stok)
		var jumlah int
		fmt.Print("Masukkan jumlah stok yang akan dihapus: ")
		fmt.Scan(&jumlah)
		produk[idx].Stok -= jumlah
		fmt.Println("Stok berhasil dihapus.")
	}
	fmt.Println("Tekan Enter untuk kembali...")
	fmt.Scanln()
	manajemenStock()
}
func cekStok() {
	clear()
	fmt.Println(bor(15), "Cek Stok", bor(15))
	for i := 0; i < idMul; i++ {
		fmt.Printf("id: %d, Nama: %s, Stok: %d\n", produk[i].Id, produk[i].Nama, produk[i].Stok)
	}
	fmt.Println("Tekan Enter untuk kembali...")
	fmt.Scanln()
	manajemenStock()
}
func main() {
	fmt.Println(bor(6), "Selamat datang di sistem manajemen fashion", bor(6))
	fmt.Println(bor(12), "Login Terlebih Dahulu", bor(12))
	if login() {
		menu()
	}
}

// FUNCTION BANTUAN atau INCLUDER

func bor(n int) string {
	var hasil string
	for i := 0; i < n; i++ {
		hasil += "="
	}
	return hasil
}
func login() bool {
	var username string
	var password string
	fmt.Print("Username ")
	fmt.Scanln(&username)
	fmt.Print("Password ")
	fmt.Scanln(&password)
	for i := 0; i < len(ak); i++ {
		if ak[i].username != "" && username == ak[i].username && password == ak[i].password {
			fmt.Println("Login Berhasil")
			return true
		}
	}
	fmt.Println("Login Gagal")
	return false
}

func tampil() {
	var keyword string
	fmt.Scan(&keyword)
	ada := false
	for i := 0; i < idMul; i++ {
		if containsI(produk[i].Nama, keyword) {
			fmt.Printf("id: %d, Nama: %s, Kategori: %s, Ukuran: %s, Warna: %s, Harga: %.2f, Stok: %d\n",
				produk[i].Id, produk[i].Nama, produk[i].kategori,
				produk[i].Ukuran, produk[i].Warna, produk[i].Harga, produk[i].Stok)
			ada = true
		}
	}
	if !ada {
		fmt.Println("Tidak ada produk yang cocok.")
	}
}

func containsI(s, sub string) bool {
	sLow := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 32
		}
		sLow[i] = c
	}
	subLow := make([]byte, len(sub))
	for i := 0; i < len(sub); i++ {
		c := sub[i]
		if c >= 'A' && c <= 'Z' {
			c += 32
		}
		subLow[i] = c
	}
	for i := 0; i <= len(sLow)-len(subLow); i++ {
		sama := true
		for j := 0; j < len(subLow); j++ {
			if sLow[i+j] != subLow[j] {
				sama = false
				break
			}
		}
		if sama {
			return true
		}
	}
	return false
}

func clear() {
	fmt.Print("\033[2J\033[H")
	fmt.Println("")
	fmt.Println("")
}
