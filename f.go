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
	Terjual  int
}

var produk = [100]Product{{1, "Kaos Polos", "Atasan", "L", "Hitam", 75000, 50, 0}, {2, "Celana Chino", "Bawahan", "32", "Khaki", 150000, 30, 0}, {3, "Topi Snapback", "Aksesoris", "One Size", "Merah", 90000, 25, 0}, {4, "Boxer Brief", "Pakaian Dalam", "M", "Abu-abu", 45000, 100, 0}, {5, "Celana Training", "Pakaian Olahraga", "XL", "Biru Navy", 120000, 40, 0}, {6, "Jaket Hoodie", "Outerwear", "L", "Hitam", 200000, 20, 0}, {7, "Kemeja Putih", "Pakaian Formal", "42", "Putih", 180000, 15, 0}, {8, "Gelang Tali", "Aksesoris", "One Size", "Coklat", 25000, 75, 0}, {9, "Rok Lipit", "Bawahan", "S", "Hitam", 110000, 35, 0}, {10, "Rompi Wanita", "Atasan", "M", "Krem", 85000, 28, 0}}

func main() {
	fmt.Println(bor(6), "Selamat datang di sistem manajemen fashion", bor(6))
	fmt.Println(bor(12), "Login Terlebih Dahulu", bor(12))
	if login() {
		menu()
	}
}

// FUNCTION BANTUAN atau INCLUDER

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
func menu() {
	clear()
	var pilihan int
	fmt.Println(bor(15), "Menu Utama", bor(15))
	fmt.Println("1. Manajemen Produk")
	fmt.Println("2. Manajemen Stok  ")
	fmt.Println("3. Statistik")
	fmt.Println("4. Keluar")
	fmt.Print("Pilih Menu Dipilih : ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		manajemenProduk()
	case 2:
		manajemenStock()
	case 3:
		statistik()
	case 4:
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
	fmt.Println("4. Cari Produk")
	fmt.Println("5. Urutkan Produk")
	fmt.Println("6. Kembali")
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
		menuCariProduk()
	case 5:
		menuUrutkanProduk()
	case 6:
		menu()
	default:
		fmt.Println("Menu tidak tersedia")
	}
}

func tambahProduk() {
	clear()
	var pil string
	if idMul >= len(produk)-1 {
		fmt.Println("Penyimpanan produk penuh!")
		return
	}
	idMul++
	produk[idMul].Id = idMul
	fmt.Print("Masukan Nama Produk : ")
	fmt.Scan(&produk[idMul].Nama)
	produk[idMul].kategori = pilihKategori()
	fmt.Println("Masukan Ukuran Produk :")
	produk[idMul].Ukuran = pilihUkuran()
	fmt.Println("Masukan Warna Produk :")
	produk[idMul].Warna = pilihWarna()
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
		fmt.Printf("Data lama: %s | Kategori: %s | Ukuran: %s | Warna: %s | Rp%.2f\n", produk[idx].Nama, produk[idx].kategori, produk[idx].Ukuran, produk[idx].Warna, produk[idx].Harga)
		fmt.Print("Nama baru: ")
		fmt.Scan(&produk[idx].Nama)
		fmt.Println("Kategori Baru:")
		produk[idx].kategori = pilihKategori()
		fmt.Println("Ukuran Baru:")
		produk[idx].Ukuran = pilihUkuran()
		fmt.Println("Warna Baru:")
		produk[idx].Warna = pilihWarna()
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

func menuCariProduk() {
	clear()
	var pilihan int
	fmt.Println(bor(15), "Cari Produk", bor(15))
	fmt.Println("1. Cari Berdasarkan Nama")
	fmt.Println("2. Cari Berdasarkan Ukuran (Sequential Search)")
	fmt.Println("3. Cari Berdasarkan Warna (Binary Search)")
	fmt.Println("4. Kembali")
	fmt.Print("Pilih Menu Dipilih : ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		cariByNama()
	case 2:
		cariByUkuran()
	case 3:
		cariByWarna()
	case 4:
		manajemenProduk()
	default:
		fmt.Println("Menu tidak tersedia")
	}
}

func cariByNama() {
	clear()
	var nama string
	fmt.Print("Masukkan nama produk yang dicari: ")
	fmt.Scan(&nama)

	var hasil []Product
	for i := 0; i < idMul; i++ {
		if produk[i].Id != 0 && containsI(produk[i].Nama, nama) {
			hasil = append(hasil, produk[i])
		}
	}

	if len(hasil) == 0 {
		fmt.Println("Produk dengan nama tersebut tidak ditemukan.")
		fmt.Println("Tekan Enter untuk kembali...")
		fmt.Scanln()
		fmt.Scanln()
		menuCariProduk()
	} else {
		cetakProdukPaginasi("Hasil Pencarian: "+nama, hasil)
		menuCariProduk()
	}
}

func cariByUkuran() {
	clear()
	fmt.Println("Pencarian Ukuran:")
	ukuran := pilihUkuran()

	var hasil []Product
	for i := 0; i < idMul; i++ {
		// Sequential search
		if produk[i].Id != 0 && (containsI(produk[i].Ukuran, ukuran) || toLower(produk[i].Ukuran) == toLower(ukuran)) {
			hasil = append(hasil, produk[i])
		}
	}

	if len(hasil) == 0 {
		fmt.Println("Produk dengan ukuran tersebut tidak ditemukan.")
		fmt.Println("Tekan Enter untuk kembali...")
		fmt.Scanln()
		fmt.Scanln()
		menuCariProduk()
	} else {
		cetakProdukPaginasi("Hasil Pencarian: "+ukuran, hasil)
		menuCariProduk()
	}
}

func cariByWarna() {
	clear()
	fmt.Println("Pencarian Warna:")
	warna := pilihWarna()

	urut := make([]Product, 0, idMul)
	for i := 0; i < idMul; i++ {
		if produk[i].Id != 0 {
			urut = append(urut, produk[i])
		}
	}

	// Sort by Warna (Selection Sort or similar) - case insensitive sort
	for i := 0; i < len(urut)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(urut); j++ {
			if toLower(urut[j].Warna) < toLower(urut[minIdx].Warna) {
				minIdx = j
			}
		}
		temp := urut[i]
		urut[i] = urut[minIdx]
		urut[minIdx] = temp
	}

	// Binary Search
	left := 0
	right := len(urut) - 1
	var hasil []Product
	wSearch := toLower(warna)

	for left <= right {
		mid := left + (right-left)/2
		wMid := toLower(urut[mid].Warna)

		if wMid == wSearch {
			hasil = append(hasil, urut[mid])

			// Expand left
			i := mid - 1
			for i >= 0 && toLower(urut[i].Warna) == wSearch {
				hasil = append(hasil, urut[i])
				i--
			}

			// Expand right
			i = mid + 1
			for i < len(urut) && toLower(urut[i].Warna) == wSearch {
				hasil = append(hasil, urut[i])
				i++
			}
			break
		} else if wMid < wSearch {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if len(hasil) == 0 {
		fmt.Println("Produk dengan warna tersebut tidak ditemukan.")
		fmt.Println("Tekan Enter untuk kembali...")
		fmt.Scanln()
		fmt.Scanln()
		menuCariProduk()
	} else {
		cetakProdukPaginasi("Hasil Pencarian: "+warna, hasil)
		menuCariProduk()
	}
}

func menuUrutkanProduk() {
	clear()
	var pilihan int
	fmt.Println(bor(15), "Urutkan Produk", bor(15))
	fmt.Println("1. Urutkan Berdasarkan Harga (Selection Sort)")
	fmt.Println("2. Urutkan Berdasarkan Stok (Insertion Sort)")
	fmt.Println("3. Kembali")
	fmt.Print("Pilih Menu Dipilih : ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		urutHarga()
	case 2:
		urutStok()
	case 3:
		manajemenProduk()
	default:
		fmt.Println("Menu tidak tersedia")
	}
}

func urutHarga() {
	urut := make([]Product, 0, idMul)
	for i := 0; i < idMul; i++ {
		if produk[i].Id != 0 {
			urut = append(urut, produk[i])
		}
	}

	// Selection Sort (Ascending)
	for i := 0; i < len(urut)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(urut); j++ {
			if urut[j].Harga < urut[minIdx].Harga {
				minIdx = j
			}
		}
		temp := urut[i]
		urut[i] = urut[minIdx]
		urut[minIdx] = temp
	}

	cetakProdukPaginasi("Produk Diurutkan by Harga", urut)
	menuUrutkanProduk()
}

func urutStok() {
	urut := make([]Product, 0, idMul)
	for i := 0; i < idMul; i++ {
		if produk[i].Id != 0 {
			urut = append(urut, produk[i])
		}
	}

	// Insertion Sort (Ascending)
	for i := 1; i < len(urut); i++ {
		key := urut[i]
		j := i - 1
		for j >= 0 && urut[j].Stok > key.Stok {
			urut[j+1] = urut[j]
			j = j - 1
		}
		urut[j+1] = key
	}

	cetakProdukPaginasi("Produk Diurutkan by Stok", urut)
	menuUrutkanProduk()
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

func cekStok() {
	for {
		clear()
		fmt.Println(bor(15), "Cek Stok Gudang", bor(15))
		fmt.Println("1. Cek Stok Dikit")
		fmt.Println("2. Cek Stok Habis")
		fmt.Println("3. Cek Stok Banyak")
		fmt.Println("4. Kembali")
		fmt.Print("Pilih Menu Dipilih : ")
		var pil int
		if _, err := fmt.Scan(&pil); err != nil {
			fmt.Println("Input tidak valid")
			return
		}
		switch pil {
		case 1:
			cekStokDikit()
		case 2:
			CekStokHabis()
		case 3:
			CekStokBanyak()
		case 4:
			menu()
			return
		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}
func tambahStock() {
	clear()
	fmt.Println(bor(15), "Tambah Stock", bor(15))
	fmt.Print("Cari produk (keyword nama): ")
	tampil()

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
	fmt.Print("Cari produk (keyword nama): ")
	tampil()

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
		produk[idx].Terjual += jumlah
		fmt.Println("Stok berhasil dihapus.")
	}
	fmt.Println("Tekan Enter untuk kembali...")
	fmt.Scanln()
	manajemenStock()
}

func cekStokDikit() {
	var stokDikit []Product
	for i := 0; i < idMul; i++ {
		stokDikit = append(stokDikit, produk[i])
	}

	if len(stokDikit) == 0 {
		fmt.Println("Tidak ada produk dengan stok dikit.")
		return
	}

	tampilkanProdukPaginasi("Cek Stok Dikit", stokDikit, true)
}

func CekStokHabis() {
	var habis []Product
	for i := 0; i < idMul; i++ {
		if produk[i].Stok == 0 {
			habis = append(habis, produk[i])
		}
	}

	if len(habis) == 0 {
		fmt.Println("Tidak ada produk yang habis stok.")
		cekStok()
		return
	}

	tampilkanProdukPaginasi("Cek Stok Habis", habis, false)
}

func CekStokBanyak() {
	clear()
	fmt.Println(bor(15), "Cek Stok Tertinggi", bor(15))

	urut := make([]Product, 0, idMul)
	for i := 0; i < idMul; i++ {
		if produk[i].Id != 0 {
			urut = append(urut, produk[i])
		}
	}

	if len(urut) == 0 {
		fmt.Println("Belum ada produk yang tersedia.")
		return
	}

	urutkanProdukByStok(urut, true)

	pageSize := 10
	page := 0
	for {
		clear()
		fmt.Println(bor(15), "Cek Stok Tertinggi", bor(15))
		fmt.Printf("Halaman %d\n", page+1)

		start := page * pageSize
		end := start + pageSize
		if end > len(urut) {
			end = len(urut)
		}

		if start >= len(urut) {
			fmt.Println("Tidak ada data lagi.")
			break
		}

		for i := start; i < end; i++ {
			p := urut[i]
			fmt.Printf("%d. %s | Stok: %d | Harga: %.2f\n", i+1, p.Nama, p.Stok, p.Harga)
		}

		fmt.Println("\n1. Next Page")
		fmt.Println("2. Kembali")
		fmt.Print("Pilih menu: ")
		var pil int
		fmt.Scan(&pil)
		switch pil {
		case 1:
			if end >= len(urut) {
				fmt.Println("Sudah di halaman terakhir.")
				return
			}
			page++
		case 2:
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
			return
		}
	}
}

func statistik() {
	clear()
	fmt.Println(bor(15), "Statistik", bor(15))

	// 1. Produk paling populer (Terjual terbanyak)
	urutTerjual := make([]Product, 0, idMul)
	for i := 0; i < idMul; i++ {
		if produk[i].Id != 0 {
			urutTerjual = append(urutTerjual, produk[i])
		}
	}

	for i := 0; i < len(urutTerjual)-1; i++ {
		maxIdx := i
		for j := i + 1; j < len(urutTerjual); j++ {
			if urutTerjual[j].Terjual > urutTerjual[maxIdx].Terjual {
				maxIdx = j
			}
		}
		temp := urutTerjual[i]
		urutTerjual[i] = urutTerjual[maxIdx]
		urutTerjual[maxIdx] = temp
	}

	fmt.Println("\n--- 3 Produk Paling Populer ---")
	maxPopuler := 3
	if len(urutTerjual) < 3 {
		maxPopuler = len(urutTerjual)
	}
	for i := 0; i < maxPopuler; i++ {
		fmt.Printf("%d. %s (Terjual: %d)\n", i+1, urutTerjual[i].Nama, urutTerjual[i].Terjual)
	}

	// 2. Sisa total stok per kategori
	type KategoriStok struct {
		Kategori string
		Stok     int
	}
	var statKategori []KategoriStok

	for i := 0; i < idMul; i++ {
		if produk[i].Id != 0 {
			kat := produk[i].kategori
			found := false
			for j := 0; j < len(statKategori); j++ {
				if statKategori[j].Kategori == kat {
					statKategori[j].Stok += produk[i].Stok
					found = true
					break
				}
			}
			if !found {
				statKategori = append(statKategori, KategoriStok{Kategori: kat, Stok: produk[i].Stok})
			}
		}
	}

	fmt.Println("\n--- Sisa Total Stok Per Kategori ---")
	for _, ks := range statKategori {
		fmt.Printf("- %s : %d stok\n", ks.Kategori, ks.Stok)
	}

	fmt.Println("\nTekan Enter untuk kembali...")
	fmt.Scanln()
	fmt.Scanln()
	menu()
}
func bor(n int) string {
	var hasil string
	for i := 0; i < n; i++ {
		hasil += "="
	}
	return hasil
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

func pilihKategori() string {
	var kategori string
	fmt.Println("Contoh Kategori: Atasan, Bawahan, Aksesoris, Pakaian Dalam, Pakaian Olahraga, Outerwear, Pakaian Formal")
	fmt.Print("Ketikkan Kategori: ")
	fmt.Scan(&kategori)
	return kategori
}

func pilihWarna() string {
	var warna string
	fmt.Println("Contoh Warna: Hitam, Putih, Merah, Biru Navy, Abu-abu")
	fmt.Print("Ketikkan Warna: ")
	fmt.Scan(&warna)
	return warna
}

func pilihUkuran() string {
	var ukuran string
	fmt.Println("Contoh Ukuran: S, M, L, XL, XXL, One Size")
	fmt.Print("Ketikkan Ukuran: ")
	fmt.Scan(&ukuran)
	return ukuran
}
func toLower(s string) string {
	res := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 32
		}
		res[i] = c
	}
	return string(res)
}
func tampilkanProdukPaginasi(title string, data []Product, ascending bool) {
	clear()
	fmt.Println(bor(15), title, bor(15))

	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			swap := false
			if ascending {
				if data[j].Stok > data[j+1].Stok {
					swap = true
				} else if data[j].Stok == data[j+1].Stok && data[j].Id > data[j+1].Id {
					swap = true
				}
			} else {
				if data[j].Stok < data[j+1].Stok {
					swap = true
				} else if data[j].Stok == data[j+1].Stok && data[j].Id > data[j+1].Id {
					swap = true
				}
			}
			if swap {
				temp := data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}
		}
	}

	pageSize := 10
	page := 0
	for {
		clear()
		fmt.Println(bor(15), title, bor(15))
		fmt.Printf("Halaman %d\n", page+1)

		start := page * pageSize
		end := start + pageSize
		if end > len(data) {
			end = len(data)
		}

		if start >= len(data) {
			fmt.Println("Tidak ada data lagi.")
			break
		}

		for i := start; i < end; i++ {
			p := data[i]
			fmt.Printf("%d. %s | ID: %d | Stok: %d | Harga: %.2f\n", i+1, p.Nama, p.Id, p.Stok, p.Harga)
		}

		fmt.Println("\n1. Next Page")
		fmt.Println("2. Kembali")
		fmt.Print("Pilih menu: ")
		var pil int
		fmt.Scan(&pil)
		switch pil {
		case 1:
			if end >= len(data) {
				fmt.Println("Sudah di halaman terakhir.")
				return
			}
			page++
		case 2:
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
			return
		}
	}
}
func urutkanProdukByStok(data []Product, descending bool) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			swap := false
			if descending {
				if data[j].Stok < data[j+1].Stok {
					swap = true
				} else if data[j].Stok == data[j+1].Stok && data[j].Id > data[j+1].Id {
					swap = true
				}
			} else {
				if data[j].Stok > data[j+1].Stok {
					swap = true
				} else if data[j].Stok == data[j+1].Stok && data[j].Id > data[j+1].Id {
					swap = true
				}
			}
			if swap {
				temp := data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}
		}
	}
}

func cetakProdukPaginasi(title string, data []Product) {
	pageSize := 10
	page := 0
	for {
		clear()
		fmt.Println(bor(15), title, bor(15))
		fmt.Printf("Halaman %d\n", page+1)

		start := page * pageSize
		end := start + pageSize
		if end > len(data) {
			end = len(data)
		}

		if start >= len(data) {
			fmt.Println("Tidak ada data lagi.")
			break
		}

		for i := start; i < end; i++ {
			p := data[i]
			fmt.Printf("%d. %s | ID: %d | Stok: %d | Harga: %.2f | Ukuran: %s | Warna: %s\n", i+1, p.Nama, p.Id, p.Stok, p.Harga, p.Ukuran, p.Warna)
		}

		fmt.Println("\n1. Next Page")
		fmt.Println("2. Kembali")
		fmt.Print("Pilih menu: ")
		var pil int
		fmt.Scan(&pil)
		switch pil {
		case 1:
			if end >= len(data) {
				fmt.Println("Sudah di halaman terakhir.")
				return
			}
			page++
		case 2:
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
			return
		}
	}
}
