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

var products []Product
var nextID int = 1

func readLine() string {
	var result string
	for {
		var ch byte
		_, err := fmt.Scanf("%c", &ch)
		if err != nil || ch == '\n' {
			break
		}
		if ch != '\r' {
			result += string(ch)
		}
	}
	return result
}

func contains(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	if len(s) < len(substr) {
		return false
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
			c1 := s[i+j]
			c2 := substr[j]
			if c1 >= 'A' && c1 <= 'Z' {
				c1 += 32
			}
			if c2 >= 'A' && c2 <= 'Z' {
				c2 += 32
			}
			if c1 != c2 {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func addProduct() {
	var p Product
	p.ID = nextID
	nextID++
	fmt.Print("Nama: ")
	p.Name = readLine()
	fmt.Print("Kategori: ")
	p.Category = readLine()
	fmt.Print("Ukuran: ")
	p.Size = readLine()
	fmt.Print("Warna: ")
	p.Color = readLine()
	fmt.Print("Harga: ")
	var priceStr string
	priceStr = readLine()
	fmt.Sscanf(priceStr, "%f", &p.Price)
	fmt.Print("Stok: ")
	var stockStr string
	stockStr = readLine()
	fmt.Sscanf(stockStr, "%d", &p.Stock)
	products = append(products, p)
	fmt.Println("Produk berhasil ditambahkan.")
}

func viewProducts() {
	if len(products) == 0 {
		fmt.Println("Tidak ada produk.")
		return
	}
	for i := 0; i < len(products); i++ {
		p := products[i]
		fmt.Printf("ID: %d, Nama: %s, Kategori: %s, Ukuran: %s, Warna: %s, Harga: %.2f, Stok: %d\n", p.ID, p.Name, p.Category, p.Size, p.Color, p.Price, p.Stock)
	}
}

func updateProduct() {
	fmt.Print("Masukkan ID produk yang akan diubah: ")
	var idStr string
	idStr = readLine()
	var id int
	fmt.Sscanf(idStr, "%d", &id)
	idx := -1
	for i := 0; i < len(products); i++ {
		if products[i].ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("Produk tidak ditemukan.")
		return
	}
	fmt.Print("Nama baru (kosongkan untuk tidak mengubah): ")
	name := readLine()
	if name != "" {
		products[idx].Name = name
	}
	fmt.Print("Kategori baru (kosongkan untuk tidak mengubah): ")
	cat := readLine()
	if cat != "" {
		products[idx].Category = cat
	}
	fmt.Print("Ukuran baru (kosongkan untuk tidak mengubah): ")
	size := readLine()
	if size != "" {
		products[idx].Size = size
	}
	fmt.Print("Warna baru (kosongkan untuk tidak mengubah): ")
	color := readLine()
	if color != "" {
		products[idx].Color = color
	}
	fmt.Print("Harga baru (kosongkan untuk tidak mengubah): ")
	priceStr := readLine()
	if priceStr != "" {
		var price float64
		fmt.Sscanf(priceStr, "%f", &price)
		products[idx].Price = price
	}
	fmt.Print("Stok baru (kosongkan untuk tidak mengubah): ")
	stockStr := readLine()
	if stockStr != "" {
		var stock int
		fmt.Sscanf(stockStr, "%d", &stock)
		products[idx].Stock = stock
	}
	fmt.Println("Produk berhasil diubah.")
}

func deleteProduct() {
	fmt.Print("Masukkan ID produk yang akan dihapus: ")
	var idStr string
	idStr = readLine()
	var id int
	fmt.Sscanf(idStr, "%d", &id)
	idx := -1
	for i := 0; i < len(products); i++ {
		if products[i].ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("Produk tidak ditemukan.")
		return
	}
	var newProducts []Product
	for i := 0; i < len(products); i++ {
		if i != idx {
			newProducts = append(newProducts, products[i])
		}
	}
	products = newProducts
	fmt.Println("Produk berhasil dihapus.")
}

func searchProduct() {
	fmt.Print("Masukkan kata kunci pencarian (nama/kategori): ")
	keyword := readLine()
	found := false
	for i := 0; i < len(products); i++ {
		if contains(products[i].Name, keyword) || contains(products[i].Category, keyword) {
			p := products[i]
			fmt.Printf("ID: %d, Nama: %s, Kategori: %s, Ukuran: %s, Warna: %s, Harga: %.2f, Stok: %d\n", p.ID, p.Name, p.Category, p.Size, p.Color, p.Price, p.Stock)
			found = true
		}
	}
	if !found {
		fmt.Println("Produk tidak ditemukan.")
	}
}

func main() {
	for {
		fmt.Println("1. Tambah Produk")
		fmt.Println("2. Lihat Semua Produk")
		fmt.Println("3. Ubah Produk")
		fmt.Println("4. Hapus Produk")
		fmt.Println("5. Cari Produk")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		var choiceStr string
		choiceStr = readLine()
		if choiceStr == "0" {
			break
		} else if choiceStr == "1" {
			addProduct()
		} else if choiceStr == "2" {
			viewProducts()
		} else if choiceStr == "3" {
			updateProduct()
		} else if choiceStr == "4" {
			deleteProduct()
		} else if choiceStr == "5" {
			searchProduct()
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
