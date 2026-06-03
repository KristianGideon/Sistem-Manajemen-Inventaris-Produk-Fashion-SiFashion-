package main

import "fmt"

type akun struct {
	username string
	password string
}

var idMul int = 10

var ak = [5]akun{{username: "admin", password: "admin"}, {username: "user", password: "user"}}

type Product struct {
	Id       int
	Nama     string
	kategori string
	Ukuran   string
	Warna    string
	Harga    float64
	Stok     int
}

var produk [100]Product

func clear() {
	fmt.Print("")
}
func menu() {
	clear()
	var pilihan int
	fmt.Println(bor(15), "Menu Utama", bor(15))
	fmt.Println("1. Manajemen Produk")
	fmt.Println("2. Manajemen Stok  ")
	fmt.Println("3. Cek Stok Gudang")
	fmt.Println("4. Keluar")
	fmt.Print("Pilih Menu Dipilih : ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		manajemenProduk()
	case 2:
		manajemenStock()
	case 3:
		cekStok()
	case 4:
		fmt.Println("Terimakasih sudah menggunakan sistem manajemen fashion")
		return
	default:
		fmt.Println("Menu tidak tersedia")
	}
}
func manajemenProduk() {
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

}
func hapusProduk() {
	fmt.Println(bor(15), "Hapus Produk", bor(15))
	var pilihPro string
	fmt.Print("Cari Berdasarkan Nama Produk Yang Mau Dihapus: ")
	fmt.Scan(&pilihPro)
	for i := 0; i <= idMul-1; i++ {
		if produk[i].Nama == pilihPro {
			produk[i] = produk[idMul]
			idMul--
			fmt.Println("Produk berhasil dihapus.")
			return
		}
	}
	fmt.Println("Produk tidak ditemukan.")
}
func manajemenStock() {

}
func cekStok() {

}
func main() {
	fmt.Println(bor(6), "Selamat datang di sistem manajemen fashion", bor(6))
	fmt.Println(bor(12), "Login Terlebih Dahulu", bor(12))
	if login() {
		menu()
	}
}

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
