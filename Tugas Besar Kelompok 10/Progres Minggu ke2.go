package main

import "fmt"
import "time"

type Barang struct {
	ID       int    
	Nama     string 
	Stok     int    
	Dipinjam int    
}

const maxItems = 100 

var inventori [maxItems]Barang 
var jumlahBarang int = 0        
var idBerikutnya int = 1        

type Transaksi struct {
	Tipe      string 
	IDBarang  int
	Jumlah    int
	Waktu     string 
}

var riwayatTransaksi [maxItems]Transaksi 
var jumlahTransaksi int = 0                

func catatTransaksi(tipe string, idBarang int, jumlah int) {
	if jumlahTransaksi < maxItems {
		waktu := time.Now().Format("2006-01-02 15:04:05") 
		transaksi := Transaksi{Tipe: tipe, IDBarang: idBarang, Jumlah: jumlah, Waktu: waktu}
		riwayatTransaksi[jumlahTransaksi] = transaksi
		jumlahTransaksi++
	}
}

func tambahBarang(nama string, stok int) {
	if jumlahBarang < maxItems {
		barang := Barang{ID: idBerikutnya, Nama: nama, Stok: stok, Dipinjam: 0}
		inventori[jumlahBarang] = barang
		jumlahBarang++
		idBerikutnya++
		fmt.Printf("Barang '%s' berhasil ditambahkan dengan ID %03d.\n", nama, barang.ID)
	} else {
		fmt.Println("Inventori penuh, tidak dapat menambah barang baru.")
	}
}

func ubahBarang(id int, nama string, stok int) {
	index := binarySearch(id)
	if index != -1 {
		inventori[index].Nama = nama
		inventori[index].Stok = stok
		fmt.Printf("Barang ID %03d berhasil diperbarui.\n", id)
	} else {
		fmt.Printf("Barang dengan ID %03d tidak ditemukan.\n", id)
	}
}

func hapusBarang(id int) {
	index := binarySearch(id)
	if index != -1 {
		for j := index; j < jumlahBarang-1; j++ {
			inventori[j] = inventori[j+1]
		}
		inventori[jumlahBarang-1] = Barang{} 
		jumlahBarang--
		fmt.Printf("Barang ID %03d berhasil dihapus.\n", id)
	} else {
		fmt.Printf("Barang dengan ID %03d tidak ditemukan.\n", id)
	}
}

func binarySearch(id int) int {
	low, high := 0, jumlahBarang-1
	for low <= high {
		mid := (low + high) / 2
		if inventori[mid].ID == id {
			return mid
		} else if inventori[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func sequentialSearch(nama string) int {
	for i := 0; i < jumlahBarang; i++ {
		if inventori[i].Nama == nama {
			return i
		}
	}
	return -1
}

func tampilkanInventori() {
	if jumlahBarang == 0 {
		fmt.Println("Inventori kosong.")
		return
	}
	fmt.Println("Inventori Barang:")
	for i := 0; i < jumlahBarang; i++ {
		fmt.Printf("ID: %03d, Nama: %s, Stok: %d, Dipinjam: %d\n", inventori[i].ID, inventori[i].Nama, inventori[i].Stok, inventori[i].Dipinjam)
	}
}

func urutkanBerdasarkanStok(ascending bool) {
	if jumlahBarang == 0 {
		fmt.Println("Inventori kosong.")
		return
	}
	for i := 0; i < jumlahBarang-1; i++ {
		for j := i + 1; j < jumlahBarang; j++ {
			if (ascending && inventori[i].Stok > inventori[j].Stok) || (!ascending && inventori[i].Stok < inventori[j].Stok) {
				inventori[i], inventori[j] = inventori[j], inventori[i]
			}
		}
	}
	fmt.Println("Barang berhasil diurutkan berdasarkan stok.")
}

func urutkanBerdasarkanNama(ascending bool) {
	if jumlahBarang == 0 {
		fmt.Println("Inventori kosong.")
		return
	}
	for i := 1; i < jumlahBarang; i++ {
		key := inventori[i]
		j := i - 1
		for j >= 0 && ((ascending && inventori[j].Nama > key.Nama) || (!ascending && inventori[j].Nama < key.Nama)) {
			inventori[j+1] = inventori[j]
			j--
		}
		inventori[j+1] = key
	}
	fmt.Println("Barang berhasil diurutkan berdasarkan nama.")
}

func cariBarang(kataKunci string, kategori string) {
	fmt.Println("Hasil Pencarian:")
	if jumlahBarang == 0 {
		fmt.Println("Inventori kosong.")
		return
	}
	for i := 0; i < jumlahBarang; i++ {
		if (kategori == "nama" && inventori[i].Nama == kataKunci) || 
		   (kategori == "stok" && fmt.Sprintf("%d", inventori[i].Stok) == kataKunci) {
			fmt.Printf("ID: %03d, Nama: %s, Stok: %d, Dipinjam: %d\n", inventori[i].ID, inventori[i].Nama, inventori[i].Stok, inventori[i].Dipinjam)
		}
	}
}

func tampilkanRiwayatTransaksi() {
	fmt.Println("Riwayat Transaksi:")
	if jumlahTransaksi == 0 {
		fmt.Println("Tidak ada transaksi.")
		return
	}
	for i := 0; i < jumlahTransaksi; i++ {
		fmt.Printf("Tipe: %s, ID Barang: %03d, Jumlah: %d, Waktu: %s\n", riwayatTransaksi[i].Tipe, riwayatTransaksi[i].IDBarang, riwayatTransaksi[i].Jumlah, riwayatTransaksi[i].Waktu)
	}
}

func pinjamBarang(id int, jumlah int) {
	index := binarySearch(id)
	if index != -1 && inventori[index].Stok >= jumlah {
		inventori[index].Dipinjam += jumlah
		inventori[index].Stok -= jumlah
		catatTransaksi("keluar", id, jumlah) 
		fmt.Printf("Barang ID %03d berhasil dipinjam sebanyak %d.\n", id, jumlah)
	} else {
		fmt.Println("Barang tidak cukup atau tidak ditemukan.")
	}
}

func kembalikanBarang(id int, jumlah int) {
	index := binarySearch(id)
	if index != -1 && inventori[index].Dipinjam >= jumlah {
		inventori[index].Dipinjam -= jumlah
		inventori[index].Stok += jumlah
		catatTransaksi("masuk", id, jumlah) 
		fmt.Printf("Barang ID %03d berhasil dikembalikan sebanyak %d.\n", id, jumlah)
	} else {
		fmt.Println("Jumlah barang yang dikembalikan melebihi yang dipinjam atau barang tidak ditemukan.")
	}
}

func main() {
	for {
		fmt.Println("\n=== Menu Aplikasi Inventori ===")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Ubah Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Tampilkan Semua Barang")
		fmt.Println("5. Urutkan Barang Berdasarkan Stok")
		fmt.Println("6. Urutkan Barang Berdasarkan Nama")
		fmt.Println("7. Cari Barang")
		fmt.Println("8. Tampilkan Catatan Transaksi")
		fmt.Println("9. Pinjam Barang") 
		fmt.Println("10. Kembalikan Barang") 
		fmt.Println("11. Keluar")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var nama string
			var stok int
			fmt.Print("Masukkan nama barang: ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan jumlah stok: ")
			fmt.Scan(&stok)
			tambahBarang(nama, stok)

		case 2:
			var id, stok int
			var nama string
			fmt.Print("Masukkan ID barang yang ingin diubah: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan nama baru: ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan stok baru: ")
			fmt.Scan(&stok)
			ubahBarang(id, nama, stok)

		case 3:
			var id int
			fmt.Print("Masukkan ID barang yang ingin dihapus: ")
			fmt.Scan(&id)
			hapusBarang(id)

		case 4:
			tampilkanInventori()

		case 5:
			var ascending int
			fmt.Print("Urutkan stok barang (1 untuk ascending, 0 untuk descending): ")
			fmt.Scan(&ascending)
			urutkanBerdasarkanStok(ascending == 1)

		case 6:
			var ascending int
			fmt.Print("Urutkan nama barang (1 untuk ascending, 0 untuk descending): ")
			fmt.Scan(&ascending)
			urutkanBerdasarkanNama(ascending == 1)

		case 7:
			var kataKunci, kategori string
			fmt.Print("Masukkan kata kunci: ")
			fmt.Scan(&kataKunci)
			fmt.Print("Masukkan kategori (nama/stok): ")
			fmt.Scan(&kategori)
			cariBarang(kataKunci, kategori)

		case 8:
			tampilkanRiwayatTransaksi()

		case 9:
			var id, jumlah int
			fmt.Print("Masukkan ID barang yang ingin dipinjam: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan jumlah barang yang ingin dipinjam: ")
			fmt.Scan(&jumlah)
			pinjamBarang(id, jumlah)

		case 10:
			var id, jumlah int
			fmt.Print("Masukkan ID barang yang ingin dikembalikan: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan jumlah barang yang ingin dikembalikan: ")
			fmt.Scan(&jumlah)
			kembalikanBarang(id, jumlah)

		case 11:
			fmt.Println("Keluar dari aplikasi. Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}