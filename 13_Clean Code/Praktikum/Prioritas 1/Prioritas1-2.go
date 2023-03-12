package main

type kendaraan struct {
	totalroda       int
	kecepatanperjam int
}

type mobil struct {
	kendaraan
}

func (m *mobil) Berjalan(kecepatanBaru int) {
	m.kecepatanperjam += kecepatanBaru
}

// Fungsi dibawah dapat tidak dipakai, sehingga tidak perlu ditulis
// func (m *mobil) tambahkecepatan(kecepatanbaru int) {
// 	m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
// }

func main() {
	mobilCepat := mobil{}
	mobilCepat.Berjalan(30)

	mobilLamban := mobil{}
	mobilLamban.Berjalan(10)
}
