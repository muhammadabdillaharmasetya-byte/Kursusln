# KursusIn - Sistem Pendaftaran Kursus Berbasis CLI

## Deskripsi Program

KursusIn merupakan aplikasi berbasis Command Line Interface (CLI) yang dikembangkan menggunakan bahasa pemrograman Go (Golang). Program ini dirancang untuk membantu pengelolaan data peserta kursus secara sederhana, mulai dari proses pendaftaran peserta, pengelolaan katalog kursus, pencarian data, pengurutan data, hingga penyajian statistik peserta berdasarkan bidang minat.

Program ini dibuat sebagai implementasi konsep algoritma dan pemrograman yang mencakup penggunaan struktur data, algoritma pencarian (searching), algoritma pengurutan (sorting), serta pengelolaan data menggunakan array statis dan struct.

---

## Tujuan Program

1. Menerapkan konsep algoritma dan pemrograman dalam sebuah studi kasus nyata.
2. Mengembangkan aplikasi pengelolaan pendaftaran kursus berbasis CLI.
3. Melatih kemampuan analisis, perancangan, dan implementasi program menggunakan bahasa Go.
4. Mengimplementasikan algoritma searching dan sorting pada data peserta kursus.
5. Menyediakan sistem sederhana untuk mengelola data peserta dan kursus.

---

## Fitur Aplikasi

### Kelola Data Peserta
- Menambahkan peserta baru.
- Mengubah data peserta.
- Menghapus data peserta.
- Mengubah status keaktifan peserta (Aktif/Tidak Aktif).

### Kelola Katalog Kursus
- Menampilkan daftar kursus yang tersedia.
- Menambahkan kursus baru ke dalam katalog.

### Pencarian Data Peserta
- Sequential Search berdasarkan bidang minat kursus.
- Binary Search berdasarkan nama peserta.

### Pengurutan Data Peserta
- Selection Sort berdasarkan ID pendaftaran.
- Insertion Sort berdasarkan nama peserta.

### Statistik Peserta
- Menampilkan jumlah peserta aktif berdasarkan bidang kursus.
- Menampilkan total peserta aktif.

### Menampilkan Data
- Menampilkan seluruh katalog kursus.
- Menampilkan seluruh data peserta yang terdaftar.

---

## Struktur Data

### Struct Kursus

```go
type Kursus struct {
    ID     int
    Nama   string
    Bidang string
}
```

### Struct Peserta

```go
type Peserta struct {
    IDReg         int
    Nama          string
    TanggalDaftar string
    KursusPilihan Kursus
    IsAktif       bool
}
```

---

## Struktur Penyimpanan Data

Program menggunakan array statis sebagai media penyimpanan data.

### Array Data Peserta

```go
var daftarPeserta [MAX_PESERTA]Peserta
```

### Array Katalog Kursus

```go
var Katalog [MAX_KURSUS]Kursus
```

### Konstanta Kapasitas

```go
const MAX_PESERTA = 200
const MAX_KURSUS = 50
```

### Tracker Jumlah Data

```go
var nPeserta int
var nKursus int
```

---

## Cara Menjalankan Program

### Persyaratan

Pastikan Go (Golang) telah terinstal pada komputer.

Cek instalasi:

```bash
go version
```

### Menjalankan Program

Clone repository:

```bash
git clone https://github.com/muhammadabdillaharmasetya-byte/KursusIn.git
```

Masuk ke folder project:

```bash
cd KursusIn
```

Jalankan program:

```bash
go run .
```

atau

```bash
go run main.go
```

---

## Panduan Pengguna

### Menu Utama

```text
1. Kelola Data Peserta
2. Cari Data Peserta
3. Urutkan Data Peserta
4. Tampilkan Statistik Bidang Minat
5. Lihat Katalog Kursus & Semua Peserta
6. Keluar
```

### Menambah Peserta

1. Pilih menu **Kelola Data Peserta**.
2. Pilih **Tambah Peserta Baru**.
3. Masukkan nama peserta.
4. Masukkan tanggal pendaftaran.
5. Pilih kursus yang tersedia atau tambahkan kursus baru.
6. Sistem akan memberikan ID pendaftaran secara otomatis.

### Mengubah Data Peserta

1. Pilih menu **Kelola Data Peserta**.
2. Pilih **Ubah Data Peserta**.
3. Masukkan ID peserta.
4. Ubah data yang diinginkan.
5. Simpan perubahan.

### Menghapus Data Peserta

1. Pilih menu **Kelola Data Peserta**.
2. Pilih **Hapus Data Peserta**.
3. Masukkan ID peserta.
4. Data peserta akan dihapus dari sistem.

### Mencari Data Peserta

- Cari berdasarkan bidang minat menggunakan Sequential Search.
- Cari berdasarkan nama menggunakan Binary Search.

### Mengurutkan Data Peserta

- Urutkan berdasarkan ID menggunakan Selection Sort.
- Urutkan berdasarkan nama menggunakan Insertion Sort.

---

## Algoritma yang Digunakan

### Sequential Search

Digunakan untuk mencari peserta berdasarkan bidang minat kursus.

**Kompleksitas Waktu:** O(n)

### Binary Search

Digunakan untuk mencari peserta berdasarkan nama.

**Syarat:** Data harus sudah diurutkan berdasarkan nama.

**Kompleksitas Waktu:** O(log n)

### Selection Sort

Digunakan untuk mengurutkan data berdasarkan ID pendaftaran.

**Kompleksitas Waktu:** O(n²)

### Insertion Sort

Digunakan untuk mengurutkan data berdasarkan nama peserta.

**Kompleksitas Waktu:** O(n²)

---

## Catatan Teknis

- Bahasa Pemrograman: Go (Golang)
- Interface: Command Line Interface (CLI)
- Penyimpanan menggunakan array statis
- Maksimum data peserta: 200
- Maksimum data kursus: 50
- Data tidak tersimpan permanen setelah program ditutup
- Program menerapkan:
  - Struct
  - Array
  - Function
  - Searching
  - Sorting
  - Percabangan
  - Perulangan

---

## Pengembang

Proyek ini dikembangkan sebagai Tugas Besar Mata Kuliah Algoritma dan Pemrograman.

**Nama Aplikasi:** KursusIn  
**Bahasa Pemrograman:** Go (Golang)
