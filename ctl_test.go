package ctl

import "testing"

// test function tambah
func TestTambah(t *testing.T) {

	if Tambah("pemasukan", 2000, "transportasi", "gojek") != true {
		t.Error("fasle")
	}

	if Tambah("pengeluaran", 1000, "makan", "nasi") != true {
		t.Error("false")
	}
}

// test laporan
func TestLaporan(t *testing.T) {
	if Laporan() == 0 {
		t.Error("false")
	}
}

// test register
func TestRegister(t *testing.T){
	if Register("admin", "pass") == "null" {
		t.Error("ada error")
	}
}


// test login
func TestLogin(t *testing.T){
	if Login("admin","pass") != true {
		t.Error("ada error")
	}
}