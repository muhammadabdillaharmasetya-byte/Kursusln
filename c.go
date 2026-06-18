package main

import (
	"fmt"
)

// Struktur Data sesuai Deskripsi Tugas Besar
type Kursus struct {
	ID     int
	Nama   string
	Bidang string
}

type Peserta struct {
	IDReg         int
	Nama          string
	TanggalDaftar string
	KursusPilihan Kursus // Menjadi struct tersendiri sesuai permintaan
	IsAktif       bool   // Opsi status keaktifan peserta
}

// Konstanta kapasitas maksimal array statis
const MAX_PESERTA = 200
const MAX_KURSUS = 50

// Data Katalog Kursus (Dinamis dengan array statis dan tracker nKursus)
var Katalog [MAX_KURSUS]Kursus
var nKursus int = 3 // Diawali dengan 3 kursus bawaan

func init() {
	// Inisialisasi data awal katalog
	Katalog[0] = Kursus{ID: 1, Nama: "Dasar_Pemrograman_Go", Bidang: "Teknologi"}
	Katalog[1] = Kursus{ID: 2, Nama: "UI_UX_Design_Pemula", Bidang: "Desain"}
	Katalog[2] = Kursus{ID: 3, Nama: "Digital_Marketing_101", Bidang: "Bisnis"}
}

func main() {
	var daftarPeserta [MAX_PESERTA]Peserta
	var nPeserta int = 0
	var lastID int = 1000 // Menghasilkan ID 1001, 1002, dst.
	var pilihan int

	for {
		fmt.Println("\n==============================================")
		fmt.Println("       SISTEM PENDAFTARAN KURSUSIN (ADMIN)       ")
		fmt.Println("==============================================")
		fmt.Println("1. Kelola Data Peserta (Tambah, Ubah, Hapus)")
		fmt.Println("2. Cari Data Peserta (Sequential / Binary)")
		fmt.Println("3. Urutkan Data Peserta (Selection / Insertion)")
		fmt.Println("4. Tampilkan Statistik Bidang Minat")
		fmt.Println("5. Lihat Katalog Kursus & Semua Peserta")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu (1-6): ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			menuKelola(&daftarPeserta, &nPeserta, &lastID)
		} else if pilihan == 2 {
			menuCari(&daftarPeserta, nPeserta)
		} else if pilihan == 3 {
			menuUrutkan(&daftarPeserta, nPeserta)
		} else if pilihan == 4 {
			tampilkanStatistik(daftarPeserta, nPeserta)
		} else if pilihan == 5 {
			tampilkanKatalogDanPeserta(daftarPeserta, nPeserta)
		} else if pilihan == 6 {
			fmt.Println("\nProgram selesai. Terima kasih Koordinator KursusIn!")
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// Helper untuk mencari kursus berdasarkan ID di katalog
func cariKursusDiKatalog(id int) (Kursus, bool) {
	for i := 0; i < nKursus; i++ {
		if Katalog[i].ID == id {
			return Katalog[i], true
		}
	}
	return Kursus{}, false
}

// Helper untuk menambahkan kursus baru langsung ke dalam katalog
func tambahKursusKeKatalog() Kursus {
	if nKursus >= MAX_KURSUS {
		fmt.Println("Katalog penuh! Menggunakan kursus default pertama.")
		return Katalog[0]
	}

	var namaKursus, bidangKursus string
	fmt.Println("\n--- BUAT KURSUS BARU ---")
	fmt.Print("Masukkan Nama Kursus Baru (Gunakan_Underscore): ")
	fmt.Scan(&namaKursus)
	fmt.Print("Masukkan Bidang (Teknologi/Desain/Bisnis): ")
	fmt.Scan(&bidangKursus)

	idBaru := nKursus + 1
	Katalog[nKursus] = Kursus{ID: idBaru, Nama: namaKursus, Bidang: bidangKursus}
	nKursus++

	fmt.Printf("Kursus '%s' berhasil ditambahkan ke katalog (ID: %d)!\n", namaKursus, idBaru)
	return Katalog[nKursus-1]
}

// --- MENU 1: KELOLA DATA PESERTA ---
func menuKelola(A *[MAX_PESERTA]Peserta, n *int, lastID *int) {
	var opsi int
	fmt.Println("\n--- KELOLA DATA PESERTA ---")
	fmt.Println("1. Tambah Peserta Baru")
	fmt.Println("2. Ubah Data & Status Peserta")
	fmt.Println("3. Hapus Data Peserta")
	fmt.Print("Pilih opsi (1-3): ")
	fmt.Scan(&opsi)

	if opsi == 1 {
		if *n >= MAX_PESERTA {
			fmt.Println("Kapasitas pendaftar penuh!")
			return
		}

		// Input Data Dasar
		*lastID++
		A[*n].IDReg = *lastID

		fmt.Print("Masukkan Nama Peserta (Gunakan_Underscore): ")
		fmt.Scan(&A[*n].Nama)
		fmt.Print("Masukkan Tanggal Daftar (DD-MM-YYYY): ")
		fmt.Scan(&A[*n].TanggalDaftar)

		// Proses Pemilihan / Penambahan Kursus
		var idKursusPilihan int
		fmt.Println("\n--- Silahkan Pilih ID Kursus ---")
		for i := 0; i < nKursus; i++ {
			fmt.Printf("[%d] %s (%s)\n", Katalog[i].ID, Katalog[i].Nama, Katalog[i].Bidang)
		}
		fmt.Println("[0] -> Kursus Tidak Ada? Ketik Baru Di Sini")
		fmt.Print("Masukkan ID Kursus Pilihan: ")
		fmt.Scan(&idKursusPilihan)

		var kursusTerpilih Kursus
		if idKursusPilihan == 0 {
			// User memilih opsi membuat kursus baru secara dinamis
			kursusTerpilih = tambahKursusKeKatalog()
		} else {
			var ditemukan bool
			kursusTerpilih, ditemukan = cariKursusDiKatalog(idKursusPilihan)
			if !ditemukan {
				fmt.Println("ID Kursus tidak valid! Dialihkan ke kursus pertama.")
				kursusTerpilih = Katalog[0]
			}
		}

		A[*n].KursusPilihan = kursusTerpilih
		A[*n].IsAktif = true // Default peserta baru adalah Aktif

		*n++
		fmt.Printf("Peserta berhasil ditambahkan! ID: %d | Kursus: %s | Status: Aktif\n", *lastID, kursusTerpilih.Nama)

	} else if opsi == 2 {
		var idCari int
		fmt.Print("Masukkan ID Peserta yang akan diubah: ")
		fmt.Scan(&idCari)
		idx := -1
		for i := 0; i < *n; i++ {
			if A[i].IDReg == idCari {
				idx = i
				break
			}
		}
		if idx != -1 {
			var statusInput int
			fmt.Print("Masukkan Nama Baru (Gunakan_Underscore): ")
			fmt.Scan(&A[idx].Nama)
			fmt.Print("Masukkan Tanggal Daftar Baru (DD-MM-YYYY): ")
			fmt.Scan(&A[idx].TanggalDaftar)

			// Ubah Kursus (Bisa juga tambah kursus baru di sini)
			fmt.Println("\n--- Pilih ID Kursus Baru ---")
			for i := 0; i < nKursus; i++ {
				fmt.Printf("[%d] %s\n", Katalog[i].ID, Katalog[i].Nama)
			}
			fmt.Println("[0] -> Daftarkan Kursus Baru")
			fmt.Print("Masukkan ID Kursus Pilihan: ")
			var idKursusPilihan int
			fmt.Scan(&idKursusPilihan)

			if idKursusPilihan == 0 {
				A[idx].KursusPilihan = tambahKursusKeKatalog()
			} else {
				kursusTerpilih, ditemukan := cariKursusDiKatalog(idKursusPilihan)
				if ditemukan {
					A[idx].KursusPilihan = kursusTerpilih
				} else {
					fmt.Println("ID tidak valid, kursus tidak diubah.")
				}
			}

			// Mengubah status keaktifan peserta
			fmt.Print("Ubah Status Keaktifan (1: Aktif, 0: Tidak Aktif): ")
			fmt.Scan(&statusInput)
			if statusInput == 1 {
				A[idx].IsAktif = true
			} else if statusInput == 0 {
				A[idx].IsAktif = false
			}

			fmt.Println("Data dan status peserta berhasil diperbarui!")
		} else {
			fmt.Println("Data peserta tidak ditemukan.")
		}

	} else if opsi == 3 {
		// Logika hapus (sama seperti kode Anda sebelumnya)
		var idCari int
		fmt.Print("Masukkan ID Peserta yang akan dihapus: ")
		fmt.Scan(&idCari)
		idx := -1
		for i := 0; i < *n; i++ {
			if A[i].IDReg == idCari {
				idx = i
				break
			}
		}
		if idx != -1 {
			for i := idx; i < *n-1; i++ {
				A[i] = A[i+1]
			}
			*n--
			fmt.Println("Data peserta berhasil dihapus!")
		} else {
			fmt.Println("Data peserta tidak ditemukan.")
		}
	}
}

// --- MENU 2: CARI DATA PESERTA ---
func menuCari(A *[MAX_PESERTA]Peserta, n int) {
	var metode int
	fmt.Println("\n--- CARI DATA PESERTA ---")
	fmt.Println("1. Cari Bidang Minat Kursus (Sequential Search)")
	fmt.Println("2. Cari Nama Lengkap (Binary Search)")
	fmt.Print("Pilih metode (1-2): ")
	fmt.Scan(&metode)

	if metode == 1 {
		var keyword string
		fmt.Print("Masukkan Bidang Minat yang dicari (Teknologi/Desain/Bisnis): ")
		fmt.Scan(&keyword)

		fmt.Println("\nHasil Pencarian:")
		ketemu := false
		for i := 0; i < n; i++ {
			if A[i].KursusPilihan.Bidang == keyword {
				statusStr := "Tidak Aktif"
				if A[i].IsAktif {
					statusStr = "Aktif"
				}
				fmt.Printf("ID: %d | Nama: %s | Kursus: %s | Status: %s\n", A[i].IDReg, A[i].Nama, A[i].KursusPilihan.Nama, statusStr)
				ketemu = true
			}
		}
		if !ketemu {
			fmt.Println("Tidak ada peserta di bidang tersebut.")
		}

	} else if metode == 2 {
		var keyword string
		fmt.Print("Masukkan Nama Lengkap yang dicari: ")
		fmt.Scan(&keyword)

		low := 0
		high := n - 1
		idx := -1

		for low <= high {
			mid := (low + high) / 2
			if A[mid].Nama == keyword {
				idx = mid
				break
			} else if A[mid].Nama < keyword {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		if idx != -1 {
			statusStr := "Tidak Aktif"
			if A[idx].IsAktif {
				statusStr = "Aktif"
			}
			fmt.Println("\nData Ditemukan!")
			fmt.Printf("ID: %d | Nama: %s | Kursus: %s | Bidang: %s | Status: %s\n",
				A[idx].IDReg, A[idx].Nama, A[idx].KursusPilihan.Nama, A[idx].KursusPilihan.Bidang, statusStr)
		} else {
			fmt.Println("Data Tidak Ditemukan! (Pastikan data sudah diurutkan berdasarkan nama lewat Menu 3)")
		}
	}
}

// --- MENU 3: URUTKAN DATA PESERTA ---
func menuUrutkan(A *[MAX_PESERTA]Peserta, n int) {
	var metode int
	fmt.Println("\n--- URUTKAN DATA PESERTA ---")
	fmt.Println("1. Urutkan ID Pendaftaran Berdasarkan Angka (Selection Sort)")
	fmt.Println("2. Urutkan Nama Secara Alfabetis (Insertion Sort)")
	fmt.Print("Pilih metode (1-2): ")
	fmt.Scan(&metode)

	if metode == 1 {
		for i := 0; i < n-1; i++ {
			minIdx := i
			for j := i + 1; j < n; j++ {
				if A[j].IDReg < A[minIdx].IDReg {
					minIdx = j
				}
			}
			temp := A[i]
			A[i] = A[minIdx]
			A[minIdx] = temp
		}
		fmt.Println("Data berhasil diurutkan berdasarkan ID menggunakan Selection Sort!")

	} else if metode == 2 {
		for i := 1; i < n; i++ {
			key := A[i]
			j := i - 1
			for j >= 0 && A[j].Nama > key.Nama {
				A[j+1] = A[j]
				j--
			}
			A[j+1] = key
		}
		fmt.Println("Data berhasil diurutkan berdasarkan Nama menggunakan Insertion Sort!")
	}
}

// --- MENU 4: TAMPILKAN STATISTIK ---
func tampilkanStatistik(A [MAX_PESERTA]Peserta, n int) {
	countTekno := 0
	countDesain := 0
	countBisnis := 0
	totalAktif := 0

	for i := 0; i < n; i++ {
		if A[i].IsAktif {
			totalAktif++
			if A[i].KursusPilihan.Bidang == "Teknologi" {
				countTekno++
			} else if A[i].KursusPilihan.Bidang == "Desain" {
				countDesain++
			} else if A[i].KursusPilihan.Bidang == "Bisnis" {
				countBisnis++
			}
		}
	}

	fmt.Println("\n--- STATISTIK JUMLAH PESERTA AKTIF PER BIDANG ---")
	fmt.Printf("1. Bidang Teknologi : %d peserta aktif\n", countTekno)
	fmt.Printf("2. Bidang Desain    : %d peserta aktif\n", countDesain)
	fmt.Printf("3. Bidang Bisnis    : %d peserta aktif\n", countBisnis)
	fmt.Printf("TOTAL PESERTA AKTIF : %d (Dari total %d data terdaftar)\n", totalAktif, n)
}

// --- MENU 5: LIHAT KATALOG & SEMUA PESERTA ---
// --- MENU 5: LIHAT KATALOG & SEMUA PESERTA ---
func tampilkanKatalogDanPeserta(A [MAX_PESERTA]Peserta, n int) {
	fmt.Println("\n--- KATALOG KURSUS ---")
	for i := 0; i < nKursus; i++ {
		fmt.Printf("ID: %d | Kursus: %-25s | Bidang: %s\n", Katalog[i].ID, Katalog[i].Nama, Katalog[i].Bidang)
	}

	fmt.Println("\n--- DAFTAR SELURUH PESERTA ---")
	if n == 0 {
		fmt.Println("[ Belum ada data peserta ]")
		return
	}
	for i := 0; i < n; i++ {
		statusStr := "Aktif"
		if !A[i].IsAktif {
			statusStr = "Tidak Aktif"
		}
		// Menambahkan kolom Tgl Daftar ke dalam format output
		fmt.Printf("ID Reg: %d | Nama: %-15s | Tgl Daftar: %-12s | Kursus: %-20s | Status: %s\n",
			A[i].IDReg, A[i].Nama, A[i].TanggalDaftar, A[i].KursusPilihan.Nama, statusStr)
	}
}