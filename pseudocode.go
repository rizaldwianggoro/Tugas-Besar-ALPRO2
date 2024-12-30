Start
Deklarasi:
    Tipe data Buku:
        - ID: string
        - Judul: string
        - Penulis: string
        - Penerbit: string
        - TahunTerbit: integer
        - JumlahDipinjam: integer

    Tipe data Peminjaman:
        - IDBuku: string
        - TanggalPinjam: tipe waktu
        - TanggalKembali: tipe waktu
        - Denda: float

    Konstanta:
        - maksBuku: integer = 100
        - maksPeminjaman: integer = 100

	Declare daftarBuku as an array of Buku with a maximum size of maksBuku
	Declare jumlahBuku as an integer initialized to 0
	Declare daftarPeminjaman as an array of Peminjaman with a maximum size of maksPeminjaman
	Declare jumlahPeminjaman as an integer initialized to 0
		

//Binary Search
Deklarasi:
    left, right, mid: integer
    hasil: array of Buku

Algoritma:
    hasil ← array kosong
    left ← 0
    right ← jumlahBuku - 1

    while left ≤ right do
        mid ← (left + right) div 2
        if kataKunci ditemukan dalam daftarBuku[mid].Judul atau
           kataKunci ditemukan dalam daftarBuku[mid].Penulis atau
           kataKunci ditemukan dalam daftarBuku[mid].Penerbit then
            Tambahkan daftarBuku[mid] ke hasil
            left ← mid + 1
            right ← mid - 1
        else if daftarBuku[mid].Judul secara leksikografis lebih kecil dari kataKunci then
            left ← mid + 1
        else
            right ← mid - 1
        end if
    end while
    return hasil
end procedure

// Selection Sort
Function selectionSortBuku(kategori):
    For i from 0 to jumlahBuku - 2:
        minIdx = i
        For j from i + 1 to jumlahBuku - 1:
            If kategori is "judul":
                If daftarBuku[j].Judul is less than daftarBuku[minIdx].Judul:
                    minIdx = j
            If kategori is "penulis":
                If daftarBuku[j].Penulis is less than daftarBuku[minIdx].Penulis:
                    minIdx = j
            If kategori is "tahun":
                If daftarBuku[j].TahunTerbit is less than daftarBuku[minIdx].TahunTerbit:
                    minIdx = j
        Swap daftarBuku[i] with daftarBuku[minIdx]

//Insertion Sort
Function insertionSortBuku(kategori):
    For i from 1 to jumlahBuku - 1:
        key = daftarBuku[i]
        j = i - 1
        Switch kategori:
            Case "judul":
                While j >= 0 and daftarBuku[j].Judul > key.Judul:
                    daftarBuku[j + 1] = daftarBuku[j]
                    j = j - 1
            Case "penulis":
                While j >= 0 and daftarBuku[j].Penulis > key.Penulis:
                    daftarBuku[j + 1] = daftarBuku[j]
                    j = j - 1
            Case "tahun":
                While j >= 0 and daftarBuku[j].TahunTerbit > key.TahunTerbit:
                    daftarBuku[j + 1] = daftarBuku[j]
                    j = j - 1
        daftarBuku[j + 1] = key

//Buku Favorit
Function tampilkanBukuTerfavorit:
    Call insertionSortBuku with parameter "jumlahdipinjam"
    Print "5 Buku Terfavorit:"
    For i from 0 to 4 (or until i < jumlahBuku):
        Print "Judul: " + daftarBuku[i].Judul + ", Jumlah Dipinjam: " + daftarBuku[i].JumlahDipinjam

// Fungsi untuk mengedit data peminjaman buku
Function editPeminjaman(idBuku, tanggalBaru):
    For each index i from 0 to jumlahPeminjaman - 1:
        If daftarPeminjaman[i].IDBuku is equal to idBuku:
            Print "Mengubah tanggal pinjam buku dengan ID" + idBuku
            Set daftarPeminjaman[i].TanggalPinjam to tanggalBaru
            Return true
    Return false

//Fungsi Hapus Peminjaman
Function hapusPeminjaman(idBuku):
    For each index i from 0 to jumlahPeminjaman - 1:
        If daftarPeminjaman[i].IDBuku is equal to idBuku:
            For each index j from i to jumlahPeminjaman - 2:
                Set daftarPeminjaman[j] to daftarPeminjaman[j+1]
            Decrease jumlahPeminjaman by 1
            Return true
    Return false

//Fungsi Edit Buku Squential
Function editBukuSequentialSearch(kataKunci, bukuBaru):
    For each index i from 0 to jumlahBuku - 1:
        If daftarBuku[i].ID is equal to kataKunci:
            Set bukuBaru.JumlahDipinjam to daftarBuku[i].JumlahDipinjam
            Set daftarBuku[i] to bukuBaru
            Return true
    Return false

// Fungsi Hapus Buku
Function hapusBuku(id):
    For each index i from 0 to jumlahBuku - 1:
        If daftarBuku[i].ID is equal to id:
            For each index j from i to jumlahBuku - 2:
                Set daftarBuku[j] to daftarBuku[j+1]
            Decrease jumlahBuku by 1
            Return true
    Return false

//Fungsi Pinjam Buku
Function pinjamBuku(idBuku):
    For each index i from 0 to jumlahBuku - 1:
        If daftarBuku[i].ID is equal to idBuku:
            If jumlahPeminjaman is less than maksPeminjaman:
                Create a new Peminjaman object with IDBuku set to idBuku and TanggalPinjam set to the current time
                Add the new peminjaman to daftarPeminjaman
                Increase jumlahPeminjaman by 1
                Increment daftarBuku[i].JumlahDipinjam by 1
                Print "Buku berhasil dipinjam."
            Else:
                Print "Kapasitas peminjaman penuh."
            Return
    Print "Buku tidak ditemukan."

//Fungsi Kembalikan Buku
Function kembalikanBuku(idBuku):
    For each index i from 0 to jumlahPeminjaman - 1:
        If daftarPeminjaman[i].IDBuku is equal to idBuku:
            Set daftarPeminjaman[i].TanggalKembali to the current time
            Calculate hariTerlambat as the difference between TanggalKembali and TanggalPinjam in days minus 7
            If hariTerlambat is greater than 0:
                Calculate denda as hariTerlambat * 5000
                Print "Buku terlambat dikembalikan" + hariTerlambat + "hari. Denda: Rp" + denda
            Else:
                Print "Buku dikembalikan tepat waktu. Tidak ada denda."
            
            For each index j from i to jumlahPeminjaman - 2:
                Set daftarPeminjaman[j] to daftarPeminjaman[j+1]
            Decrease jumlahPeminjaman by 1

            For each index j from 0 to jumlahBuku - 1:
                If daftarBuku[j].ID is equal to idBuku:
                    Decrease daftarBuku[j].JumlahDipinjam by 1
                    Return
    Print "Buku tidak ditemukan dalam daftar peminjaman."

//Fungsi Tampilan Buku Pinjaman
Function tampilkanBukuDipinjam:
    Print "Daftar Buku yang Sedang Dipinjam:"
    For each peminjaman in daftarPeminjaman up to jumlahPeminjaman:
        Print "ID Buku: " + peminjaman.IDBuku + ", Tanggal Pinjam: " + TanggalPinjam formatted as "DD-MM-YYYY"

Function main():
    Initialize reader for input
    Initialize daftarBuku with sample books
    Set jumlahBuku to 8

    Loop indefinitely:
        Print "Aplikasi Perpustakaan"
        Print menu options (1 to 12)
        Prompt user to input their choice (pilihan)
        
        Switch on pilihan:
            Case 1: Tambah Buku
                If jumlahBuku < maksBuku:
                    Create a new book (bukuBaru)
                    Prompt user for book details (ID, Judul, Penulis, Penerbit, Tahun Terbit)
                    Trim spaces from the input values
                    Add bukuBaru to daftarBuku and increment jumlahBuku
                    Print "Buku berhasil ditambahkan."
                Else:
                    Print "Kapasitas daftar buku penuh."         

            Case 2: Edit Buku
                Prompt user for the ID of the book to edit (kataKunci)
                Create a new book (bukuBaru) and prompt user for new details
                Trim spaces from the input values
                Call editBukuSequentialSearch with kataKunci and bukuBaru
                If editBukuSequentialSearch returns true, print "Buku berhasil diedit."
                Else, print "Buku tidak ditemukan."

            Case 3: Hapus Buku
                Prompt user for the ID of the book to delete
                Call hapusBuku with the entered ID
                If hapusBuku returns true, print "Buku berhasil dihapus."
                Else, print "ID Buku tidak ditemukan."

            Case 4: Cari Buku
                Prompt user for a search keyword (kataKunci)
                Call cariBukuBinary with kataKunci
                If no results found, print "Buku tidak ditemukan."
                Else, print the search results

            Case 5: Pinjam Buku
                Prompt user for the book ID to borrow
                Call pinjamBuku with the entered ID

            Case 6: Edit Buku Pinjaman
                Prompt user for the ID of the loan to edit
                Prompt user for the new return date (tanggalBaru)
                If the date format is incorrect, print "Format tanggal salah."
                Else, call editPeminjaman with the ID and new date
                If editPeminjaman returns true, print "Data peminjaman berhasil diedit."
                Else, print "Data peminjaman tidak ditemukan."

            Case 7: Hapus Buku Pinjaman
                Prompt user for the ID of the loan to delete
                Call hapusPeminjaman with the entered ID
                If hapusPeminjaman returns true, print "Data peminjaman berhasil dihapus."
                Else, print "Data peminjaman tidak ditemukan."

            Case 8: Kembalikan Buku
                Prompt user for the book ID to return
                Call kembalikanBuku with the entered ID

            Case 9: Tampilankan Buku Dipinjam
                Call tampilkanBukuDipinjam

            Case 10: Tampilkan 5 Buku Terfavorit
                Call tampilkanBukuTerfavorit

            Case 11: Urutkan Buku
                Prompt user for the sorting category (judul/tahun/penulis)
                Call selectionSortBuku with the selected category
                Print the sorted list of books

            Case 12: Keluar
                Print "Terima kasih telah menggunakan aplikasi perpustakaan!"
                Return
End
