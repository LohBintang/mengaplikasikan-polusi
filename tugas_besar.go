package main

import "fmt"

const maxData = 100

type DataPolusi struct {
    id       int
    lokasi   string
    tanggal  string
    sumber   string
    aqi      int
    kategori string
}

var data [maxData]DataPolusi
var jumlahData int
var idTerakhir int

func main() {
    isiDataDummy() 

    var pilihan int
    for {
        fmt.Println("\n===== MENU MONITORING POLUSI UDARA =====")
        fmt.Println("1. Masukkan Data Polusi")
        fmt.Println("2. Tampilkan Semua Data")
        fmt.Println("3. Cari Data Kota")
        fmt.Println("4. Urutkan AQI (Tertinggi ke Terendah)")
        fmt.Println("5. Tampilkan Kota dengan AQI Tertinggi")
        fmt.Println("6. Tampilkan Kota Berbahaya (AQI > 200)")
        fmt.Println("7. Hitung Rata-rata AQI")
        fmt.Println("8. Keluar")
        fmt.Println("9. Ubah Data")
        fmt.Println("10. Hapus Data")
        fmt.Print("Pilih menu: ")
        fmt.Scanln(&pilihan)

        switch pilihan {
        case 1:
            masukkanData()
        case 2:
            tampilkanData()
        case 3:
            menuCariData()
        case 4:
            urutkanAQI()
        case 5:
            tampilkanAQITertinggi()
        case 6:
            tampilkanKotaBerbahaya()
        case 7:
            hitungRataRataAQI()
        case 8:
            fmt.Println("Terima kasih telah menggunakan aplikasi.")
            return
        case 9:
            ubahData()
        case 10:
            hapusData()
        default:
            fmt.Println("Pilihan tidak valid.")
        }
        fmt.Println()
    }
}

func isiDataDummy() {
    dataDummy := []DataPolusi{
        {lokasi: "Jakarta", tanggal: "01-01-2025", sumber: "Kendaraan", aqi: 180},
        {lokasi: "Bandung", tanggal: "02-01-2025", sumber: "Pabrik", aqi: 210},
        {lokasi: "Surabaya", tanggal: "03-01-2025", sumber: "Pembakaran", aqi: 95},
        {lokasi: "Yogyakarta", tanggal: "04-01-2025", sumber: "Kendaraan", aqi: 60},
        {lokasi: "Medan", tanggal: "05-01-2025", sumber: "Industri", aqi: 250},
    }

    for _, d := range dataDummy {
        if jumlahData >= maxData {
            fmt.Println("Kapasitas data penuh.")
            break
        }

        idTerakhir++
        d.id = idTerakhir

        if d.aqi <= 50 {
            d.kategori = "Baik"
        } else if d.aqi <= 100 {
            d.kategori = "Sedang"
        } else if d.aqi <= 200 {
            d.kategori = "Tidak Sehat"
        } else {
            d.kategori = "Berbahaya"
        }

        data[jumlahData] = d
        jumlahData++
    }

    fmt.Println("Data dummy berhasil dimasukkan.")
}


func masukkanData() {
    if jumlahData >= maxData {
        fmt.Println("Data penuh!")
        return
    }

    fmt.Println("\n=== Masukkan Data Polusi ===")

    fmt.Print("Lokasi            : ")
    fmt.Scanln(&data[jumlahData].lokasi)

    fmt.Print("Tanggal (dd-mm-yyyy): ")
    fmt.Scanln(&data[jumlahData].tanggal)

    fmt.Print("Sumber Polusi     : ")
    fmt.Scanln(&data[jumlahData].sumber)

    fmt.Print("AQI               : ")
    fmt.Scanln(&data[jumlahData].aqi)

    aqi := data[jumlahData].aqi
    if aqi <= 50 {
        data[jumlahData].kategori = "Baik"
    } else if aqi <= 100 {
        data[jumlahData].kategori = "Sedang"
    } else if aqi <= 200 {
        data[jumlahData].kategori = "Tidak Sehat"
    } else {
        data[jumlahData].kategori = "Berbahaya"
    }

    idTerakhir++
    data[jumlahData].id = idTerakhir
    jumlahData++

    fmt.Println("Data berhasil ditambahkan.")
}

func tampilkanData() {
    if jumlahData == 0 {
        fmt.Println("Belum ada data.")
        return
    }
    fmt.Println("\n=== Daftar Data Polusi ===")
    for i := 0; i < jumlahData; i++ {
        fmt.Printf("ID: %d | Lokasi: %s, Tanggal: %s, Sumber: %s, AQI: %d, Kategori: %s\n",
            data[i].id, data[i].lokasi, data[i].tanggal, data[i].sumber, data[i].aqi, data[i].kategori)
    }
}

func menuCariData() {
    var pilih int
    fmt.Println("\n=== MENU PENCARIAN DATA KOTA ===")
    fmt.Println("Pilih metode pencarian:")
    fmt.Println("1. Sequential Search (pencarian berurutan)")
    fmt.Println("2. Binary Search (pencarian biner - data akan diurutkan dulu berdasarkan nama kota)")
    fmt.Print("Masukkan pilihan (1 atau 2): ")
    fmt.Scanln(&pilih)

    switch pilih {
    case 1:
        fmt.Println("Metode pencarian: Sequential Search")
        sequentialSearch()
    case 2:
        fmt.Println("Metode pencarian: Binary Search")
        urutkanNamaKotaAsc()
        binarySearch()
    default:
        fmt.Println("Pilihan tidak valid. Kembali ke menu utama.")
        return 
    }
}

func sequentialSearch() {
    var cari string
    fmt.Print("Masukkan nama kota yang ingin dicari (Sequential): ")
    fmt.Scan(&cari)
    ketemu := false
    for i := 0; i < jumlahData; i++ {
        if data[i].lokasi == cari {
            fmt.Printf("Ditemukan: ID: %d | Lokasi: %s, Tanggal: %s, AQI: %d, Kategori: %s\n",
                data[i].id, data[i].lokasi, data[i].tanggal, data[i].aqi, data[i].kategori)
            ketemu = true
        }
    }
    if !ketemu {
        fmt.Println("Data tidak ditemukan.")
    }
}

func binarySearch() {
    var cari string
    fmt.Print("Masukkan nama kota yang ingin dicari (Binary): ")
    fmt.Scan(&cari)

    kiri := 0
    kanan := jumlahData - 1
    ketemu := false

    for kiri <= kanan {
        tengah := (kiri + kanan) / 2
        if data[tengah].lokasi == cari {
            fmt.Printf("Ditemukan: ID: %d | Lokasi: %s, Tanggal: %s, AQI: %d, Kategori: %s\n",
                data[tengah].id, data[tengah].lokasi, data[tengah].tanggal, data[tengah].aqi, data[tengah].kategori)
            ketemu = true
            break
        } else if data[tengah].lokasi < cari {
            kiri = tengah + 1
        } else {
            kanan = tengah - 1
        }
    }

    if !ketemu {
        fmt.Println("Data tidak ditemukan.")
    }
}

func urutkanAQI() {
    for i := 0; i < jumlahData-1; i++ {
        maxIdx := i
        for j := i + 1; j < jumlahData; j++ {
            if data[j].aqi > data[maxIdx].aqi {
                maxIdx = j
            }
        }
        data[i], data[maxIdx] = data[maxIdx], data[i]
    }
    fmt.Println("Data berhasil diurutkan dari AQI tertinggi ke terendah.")
}

func tampilkanAQITertinggi() {
    if jumlahData == 0 {
        fmt.Println("Belum ada data.")
        return
    }
    maxIdx := 0
    for i := 1; i < jumlahData; i++ {
        if data[i].aqi > data[maxIdx].aqi {
            maxIdx = i
        }
    }
    fmt.Printf("Kota dengan AQI tertinggi: %s (%d)\n", data[maxIdx].lokasi, data[maxIdx].aqi)
}

func tampilkanKotaBerbahaya() {
    fmt.Println("Kota dengan kategori Berbahaya (AQI > 200):")
    ditemukan := false
    for i := 0; i < jumlahData; i++ {
        if data[i].aqi > 200 {
            fmt.Printf("- %s: AQI %d\n", data[i].lokasi, data[i].aqi)
            ditemukan = true
        }
    }
    if !ditemukan {
        fmt.Println("Tidak ada kota dengan AQI berbahaya.")
    }
}

func hitungRataRataAQI() {
    if jumlahData == 0 {
        fmt.Println("Belum ada data.")
        return
    }
    total := 0
    for i := 0; i < jumlahData; i++ {
        total += data[i].aqi
    }
    rata := float64(total) / float64(jumlahData)
    fmt.Printf("Rata-rata AQI dari seluruh data: %.2f\n", rata)
}

func urutkanNamaKotaAsc() {
    for i := 0; i < jumlahData-1; i++ {
        minIdx := i
        for j := i + 1; j < jumlahData; j++ {
            if data[j].lokasi < data[minIdx].lokasi {
                minIdx = j
            }
        }
        data[i], data[minIdx] = data[minIdx], data[i]
    }
}

func ubahData() {
    if jumlahData == 0 {
        fmt.Println("Belum ada data.")
        return
    }

    var id int
    fmt.Print("Masukkan ID data yang ingin diubah: ")
    fmt.Scanln(&id)

    index := -1
    for i := 0; i < jumlahData; i++ {
        if data[i].id == id {
            index = i
            break
        }
    }

    if index == -1 {
        fmt.Println("Data dengan ID tersebut tidak ditemukan.")
        return
    }

    fmt.Printf("Data lama: Lokasi: %s, Tanggal: %s, Sumber: %s, AQI: %d\n",
        data[index].lokasi, data[index].tanggal, data[index].sumber, data[index].aqi)

 
    fmt.Print("Masukkan lokasi baru: ")
    fmt.Scanln(&data[index].lokasi)

    fmt.Print("Masukkan tanggal baru (dd-mm-yyyy): ")
    fmt.Scanln(&data[index].tanggal)

    fmt.Print("Masukkan sumber baru: ")
    fmt.Scanln(&data[index].sumber)

    fmt.Print("Masukkan AQI baru: ")
    fmt.Scanln(&data[index].aqi)

    
    if data[index].aqi <= 50 {
        data[index].kategori = "Baik"
    } else if data[index].aqi <= 100 {
        data[index].kategori = "Sedang"
    } else if data[index].aqi <= 200 {
        data[index].kategori = "Tidak Sehat"
    } else {
        data[index].kategori = "Berbahaya"
    }

    fmt.Println("Data berhasil diubah.")
}

func hapusData() {
    if jumlahData == 0 {
        fmt.Println("Belum ada data.")
        fmt.Print("Tekan ENTER untuk kembali ke menu...")
        var dummy string
        fmt.Scanln(&dummy)
        return
    }

    var id int
    fmt.Print("Masukkan ID data yang ingin dihapus: ")
    fmt.Scan(&id)
    fmt.Scanln() 

    index := -1
    for i := 0; i < jumlahData; i++ {
        if data[i].id == id {
            index = i
            break
        }
    }

    if index == -1 {
        fmt.Println("Data dengan ID tersebut tidak ditemukan.")
        fmt.Print("Tekan ENTER untuk kembali ke menu...")
        var dummy string
        fmt.Scanln(&dummy)
        return
    }

    for i := index; i < jumlahData-1; i++ {
        data[i] = data[i+1]
    }
    jumlahData--

    fmt.Println("Data berhasil dihapus.")
    fmt.Print("Tekan ENTER untuk kembali ke menu...")
    var dummy string
    fmt.Scanln(&dummy)
}
