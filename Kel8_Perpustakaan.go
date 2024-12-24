package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Buku struct {
	ID            string
	Judul         string
	Penulis       string
	Penerbit      string
	TahunTerbit   int
	JumlahDipinjam int
}

type Peminjaman struct {
	IDBuku         string
	TanggalPinjam  time.Time
	TanggalKembali time.Time
	Denda          float64
}

const (
	maksBuku       = 100
	maksPeminjaman = 100
)

var daftarBuku [maksBuku]Buku
var jumlahBuku int
var daftarPeminjaman [maksPeminjaman]Peminjaman
var jumlahPeminjaman int


// Binary Search
func cariBukuBinary(kataKunci string) []Buku {
	var hasil []Buku
	left, right := 0, jumlahBuku-1
	for left <= right {
		mid := (left + right) / 2
		if strings.Contains(strings.ToLower(daftarBuku[mid].Judul), strings.ToLower(kataKunci)) ||
			strings.Contains(strings.ToLower(daftarBuku[mid].Penulis), strings.ToLower(kataKunci)) ||
			strings.Contains(strings.ToLower(daftarBuku[mid].Penerbit), strings.ToLower(kataKunci)) {
			hasil = append(hasil, daftarBuku[mid])
			left = mid + 1
			right = mid - 1
		} else if strings.Compare(strings.ToLower(daftarBuku[mid].Judul), strings.ToLower(kataKunci)) < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return hasil
}

// Selection Sort
func selectionSortBuku(kategori string) {
	for i := 0; i < jumlahBuku-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahBuku; j++ {
			switch kategori {
			case "judul":
				if daftarBuku[j].Judul < daftarBuku[minIdx].Judul {
					minIdx = j
				}
			case "penulis":
				if daftarBuku[j].Penulis < daftarBuku[minIdx].Penulis {
					minIdx = j
				}
			case "tahun":
				if daftarBuku[j].TahunTerbit < daftarBuku[minIdx].TahunTerbit {
					minIdx = j
				}
			}
		}
		daftarBuku[i], daftarBuku[minIdx] = daftarBuku[minIdx], daftarBuku[i]
	}
}

// Insertion Sort
func insertionSortBuku(kategori string) {
	for i := 1; i < jumlahBuku; i++ {
		key := daftarBuku[i]
		j := i - 1
		switch kategori {
		case "judul":
			for j >= 0 && daftarBuku[j].Judul > key.Judul {
				daftarBuku[j+1] = daftarBuku[j]
				j--
			}
		case "penulis":
			for j >= 0 && daftarBuku[j].Penulis > key.Penulis {
				daftarBuku[j+1] = daftarBuku[j]
				j--
			}
		case "tahun":
			for j >= 0 && daftarBuku[j].TahunTerbit > key.TahunTerbit {
				daftarBuku[j+1] = daftarBuku[j]
				j--
			}
		}
		daftarBuku[j+1] = key
	}
}


func tampilkanBukuTerfavorit() {
	insertionSortBuku("jumlahdipinjam")
	fmt.Println("5 Buku Terfavorit:")
	for i := 0; i < 5 && i < jumlahBuku; i++ {
		fmt.Printf("Judul: %s, Jumlah Dipinjam: %d\n", daftarBuku[i].Judul, daftarBuku[i].JumlahDipinjam)
	}
}


func editBukuSequentialSearch(kataKunci string, bukuBaru Buku) bool {
	for i := 0; i < jumlahBuku; i++ {
		if daftarBuku[i].ID == kataKunci {
			bukuBaru.JumlahDipinjam = daftarBuku[i].JumlahDipinjam // Tetap gunakan jumlah dipinjam yang lama
			daftarBuku[i] = bukuBaru
			return true
		}
	}
	return false
}



func hapusBuku(id string) bool {
	for i := 0; i < jumlahBuku; i++ {
		if daftarBuku[i].ID == id {
			for j := i; j < jumlahBuku-1; j++ {
				daftarBuku[j] = daftarBuku[j+1]
			}
			jumlahBuku--
			return true
		}
	}
	return false
}

func pinjamBuku(idBuku string) {
	for i := 0; i < jumlahBuku; i++ {
		if daftarBuku[i].ID == idBuku {
			if jumlahPeminjaman < maksPeminjaman {
				peminjaman := Peminjaman{
					IDBuku:        idBuku,
					TanggalPinjam: time.Now(),
				}
				daftarPeminjaman[jumlahPeminjaman] = peminjaman
				jumlahPeminjaman++
				daftarBuku[i].JumlahDipinjam++
				fmt.Println("Buku berhasil dipinjam.")
			} else {
				fmt.Println("Kapasitas peminjaman penuh.")
			}
			return
		}
	}
	fmt.Println("Buku tidak ditemukan.")
}

func kembalikanBuku(idBuku string) {
	for i := 0; i < jumlahPeminjaman; i++ {
		if daftarPeminjaman[i].IDBuku == idBuku {

			daftarPeminjaman[i].TanggalKembali = time.Now()
			hariTerlambat := int(daftarPeminjaman[i].TanggalKembali.Sub(daftarPeminjaman[i].TanggalPinjam).Hours()/24) - 7 // Batas waktu 7 hari
			if hariTerlambat > 0 {
				daftarPeminjaman[i].Denda = float64(hariTerlambat) * 5000 // Denda Rp5000 per hari
				fmt.Printf("Buku terlambat dikembalikan %d hari. Denda: Rp%.2f\n", hariTerlambat, daftarPeminjaman[i].Denda)
			} else {
				fmt.Println("Buku dikembalikan tepat waktu. Tidak ada denda.")
			}


			for j := i; j < jumlahPeminjaman-1; j++ {
				daftarPeminjaman[j] = daftarPeminjaman[j+1]
			}
			jumlahPeminjaman--

			for j := 0; j < jumlahBuku; j++ {
				if daftarBuku[j].ID == idBuku {
					daftarBuku[j].JumlahDipinjam--
					break
				}
			}
			return
		}
	}
	fmt.Println("Buku tidak ditemukan dalam daftar peminjaman.")
}


func tampilkanBukuDipinjam() {
	fmt.Println("Daftar Buku yang Sedang Dipinjam:")
	for _, peminjaman := range daftarPeminjaman[:jumlahPeminjaman] {
		fmt.Printf("ID Buku: %s, Tanggal Pinjam: %s\n", peminjaman.IDBuku, peminjaman.TanggalPinjam.Format("02-01-2006"))
	}
}


func hapusPeminjaman(idBuku string) bool {
	for i := 0; i < jumlahPeminjaman; i++ {
		if daftarPeminjaman[i].IDBuku == idBuku {
			for j := i; j < jumlahPeminjaman-1; j++ {
				daftarPeminjaman[j] = daftarPeminjaman[j+1]
			}
			jumlahPeminjaman--
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	daftarBuku[0] = Buku{ID: "B001", Judul: "Pemrograman Go", Penulis: "John Doe", Penerbit: "Gramedia", TahunTerbit: 2020}
	daftarBuku[1] = Buku{ID: "B002", Judul: "Algoritma Dasar", Penulis: "Jane Doe", Penerbit: "Gramedia", TahunTerbit: 2018}
	daftarBuku[2] = Buku{ID: "B003", Judul: "Struktur Data", Penulis: "Alice", Penerbit: "Informatika", TahunTerbit: 2021}
	daftarBuku[3] = Buku{ID: "B004", Judul: "Pemrograman Python", Penulis: "Michael Scott", Penerbit: "Tech Books", TahunTerbit: 2019}
	daftarBuku[4] = Buku{ID: "B005", Judul: "Database Management", Penulis: "Dwight Schrute", Penerbit: "IT Publisher", TahunTerbit: 2020}
	daftarBuku[5] = Buku{ID: "B006", Judul: "Web Development", Penulis: "Pam Beesly", Penerbit: "WebMaster", TahunTerbit: 2017}
	daftarBuku[6] = Buku{ID: "B007", Judul: "Machine Learning Basics", Penulis: "Jim Halpert", Penerbit: "AI Press", TahunTerbit: 2021}
	daftarBuku[7] = Buku{ID: "B008", Judul: "Cloud Computing", Penulis: "Stanley Hudson", Penerbit: "CloudTech", TahunTerbit: 2020}
	jumlahBuku = 8


	for {
		fmt.Println("\nAplikasi Perpustakaan")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Edit Buku")
		fmt.Println("3. Hapus Buku")
		fmt.Println("4. Cari Buku")
		fmt.Println("5. Pinjam Buku")
		fmt.Println("6. Kembalikan Buku")
		fmt.Println("7. Tampilkan Buku Dipinjam")
		fmt.Println("8. Tampilkan 5 Buku Terfavorit")
		fmt.Println("9. Urutkan Buku")
		fmt.Println("10. Keluar")
		fmt.Print("Masukkan pilihan: ")
		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			if jumlahBuku < maksBuku {
				var bukuBaru Buku
				fmt.Print("Masukkan ID: ")
				bukuBaru.ID, _ = reader.ReadString('\n')
				fmt.Print("Masukkan Judul: ")
				bukuBaru.Judul, _ = reader.ReadString('\n')
				fmt.Print("Masukkan Penulis: ")
				bukuBaru.Penulis, _ = reader.ReadString('\n')
				fmt.Print("Masukkan Penerbit: ")
				bukuBaru.Penerbit, _ = reader.ReadString('\n')
				fmt.Print("Masukkan Tahun Terbit: ")
				fmt.Scan(&bukuBaru.TahunTerbit)
				bukuBaru.ID = strings.TrimSpace(bukuBaru.ID)
				bukuBaru.Judul = strings.TrimSpace(bukuBaru.Judul)
				bukuBaru.Penulis = strings.TrimSpace(bukuBaru.Penulis)
				bukuBaru.Penerbit = strings.TrimSpace(bukuBaru.Penerbit)
				daftarBuku[jumlahBuku] = bukuBaru
				jumlahBuku++
				fmt.Println("Buku berhasil ditambahkan.")
			} else {
				fmt.Println("Kapasitas daftar buku penuh.")
			}
		case 2:
			fmt.Print("Masukkan kata kunci untuk mencari buku yang ingin diedit (ID Buku): ")
			kataKunci, _ := reader.ReadString('\n')
			kataKunci = strings.TrimSpace(kataKunci)
		
			var bukuBaru Buku
			fmt.Print("Masukkan ID baru: ")
			bukuBaru.ID, _ = reader.ReadString('\n')
			fmt.Print("Masukkan Judul baru: ")
			bukuBaru.Judul, _ = reader.ReadString('\n')
			fmt.Print("Masukkan Penulis baru: ")
			bukuBaru.Penulis, _ = reader.ReadString('\n')
			fmt.Print("Masukkan Penerbit baru: ")
			bukuBaru.Penerbit, _ = reader.ReadString('\n')
			fmt.Print("Masukkan Tahun Terbit baru: ")
			fmt.Scan(&bukuBaru.TahunTerbit)
			bukuBaru.ID = strings.TrimSpace(bukuBaru.ID)
			bukuBaru.Judul = strings.TrimSpace(bukuBaru.Judul)
			bukuBaru.Penulis = strings.TrimSpace(bukuBaru.Penulis)
			bukuBaru.Penerbit = strings.TrimSpace(bukuBaru.Penerbit)
		
			if editBukuSequentialSearch(kataKunci, bukuBaru) {
				fmt.Println("Buku berhasil diedit.")
			} else {
				fmt.Println("Buku tidak ditemukan.")
			}
		
		case 3:
			fmt.Print("Masukkan ID Buku yang ingin dihapus: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			if hapusBuku(id) {
				fmt.Println("Buku berhasil dihapus.")
			} else {
				fmt.Println("ID Buku tidak ditemukan.")
			}

		case 4:
			fmt.Print("Masukkan kata kunci pencarian: ")
			kataKunci, _ := reader.ReadString('\n')
			kataKunci = strings.TrimSpace(kataKunci)

			hasil := cariBukuBinary(kataKunci)
			if len(hasil) == 0 {
				fmt.Println("Buku tidak ditemukan.")
				} else {
					fmt.Println("Hasil pencarian:")
					for _, buku := range hasil {
						fmt.Printf("ID: %s, Judul: %s, Penulis: %s, Penerbit: %s, Tahun Terbit: %d\n",
						buku.ID, buku.Judul, buku.Penulis, buku.Penerbit, buku.TahunTerbit)
					}
				}
		case 5:
			fmt.Print("Masukkan ID Buku yang ingin dipinjam: ")
			idBuku, _ := reader.ReadString('\n')
			idBuku = strings.TrimSpace(idBuku)
			pinjamBuku(idBuku)

		case 6:
			fmt.Print("Masukkan ID Buku yang ingin dikembalikan: ")
			idBuku, _ := reader.ReadString('\n')
			idBuku = strings.TrimSpace(idBuku)
			kembalikanBuku(idBuku)
		case 7:
			tampilkanBukuDipinjam()
		case 8:
			tampilkanBukuTerfavorit()
		case 9:
			fmt.Print("Urutkan berdasarkan (judul/tahun/penulis): ")
			kategori, _ := reader.ReadString('\n')
			kategori = strings.TrimSpace(kategori)
			
			selectionSortBuku(kategori)
			fmt.Println("Daftar buku setelah diurutkan:")
			for _, buku := range daftarBuku[:jumlahBuku] {
				fmt.Printf("ID: %s, Judul: %s, Penulis: %s, Penerbit: %s, Tahun Terbit: %d\n",
				buku.ID, buku.Judul, buku.Penulis, buku.Penerbit, buku.TahunTerbit)
			}
		case 10:
			fmt.Println("Terima kasih telah menggunakan aplikasi perpustakaan!")
			return
		}
	}
}