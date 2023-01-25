package ctl

var (
	Selisih, AllIn, AllOut int
)

func Tambah(jenis string, jml int, kategori string, note string) bool {
	if jenis == "pemasukan" {
		// save to database
		AllIn += jml
		return true
	} else if jenis == "pengeluaran" {
		// save to database
		AllIn -= jml
		return true
	} else if jenis == "" {
		return false
	}
	return false
}

func Laporan() int {
	Selisih := AllIn - AllOut
	if Selisih == 0 || Selisih <= 0 {
		return 0
	}
	return Selisih
}