package library

import (
	"fmt"
	"math"
	"strings"
)

// Soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() float64
	keliling() float64
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() float64 {
	return float64(s.alas*s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() float64 {
	return float64(3 * s.alas)
}

func (p persegiPanjang) luas() float64 {
	return float64(p.panjang * p.lebar)
}

func (p persegiPanjang) keliling() float64 {
	return float64(2 * (p.panjang + p.lebar))
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return float64(2*(b.panjang*b.lebar) + 2*(b.panjang*b.tinggi) + 2*(b.lebar*b.tinggi))
}

func Results() {
	fmt.Println("----- SOAL 1-----")
	segitiga := segitigaSamaSisi{
		alas:   10,
		tinggi: 6,
	}
	persegi := persegiPanjang{
		panjang: 4,
		lebar:   5,
	}
	silinder := tabung{
		jariJari: 14,
		tinggi:   8,
	}
	kubus := balok{
		panjang: 8,
		lebar:   4,
		tinggi:  6,
	}

	var bangunDatar hitungBangunDatar
	var bangunRuang hitungBangunRuang

	bangunDatar = segitiga
	fmt.Println("Segitiga Sama Sisi")
	fmt.Println("Luas Segitiga Sama Sisi:", bangunDatar.luas())
	fmt.Println("Keliling Segitiga Sama Sisi:", bangunDatar.keliling())
	fmt.Println()

	bangunDatar = persegi
	fmt.Println("Persegi Panjang")
	fmt.Println("Luas Persegi Panjang:", bangunDatar.luas())
	fmt.Println("Keliling Persegi Panjang:", bangunDatar.keliling())
	fmt.Println()

	bangunRuang = silinder
	fmt.Println("Tabung")
	fmt.Println("Volume Tabung:", bangunRuang.volume())
	fmt.Println("Luas Permukaan Tabung:", bangunRuang.luasPermukaan())
	fmt.Println()

	bangunRuang = kubus
	fmt.Println("Balok")
	fmt.Println("Volume Balok:", bangunRuang.volume())
	fmt.Println("Luas Permukaan Balok:", bangunRuang.luasPermukaan())
}

// Soal 2
type phone struct {
	name  string
	brand string
	year  int
	color []string
}

type phoneData interface {
	displayData() string
}

func (p phone) displayData() string {
	data := fmt.Sprintf("name : %s\nbrand : %s\nyear : %d\ncolor :%s", p.name, p.brand, p.year, strings.Join(p.color, ", "))
	return data
}

func data() {
	fmt.Println("----- SOAL 2 -----")
	p := phone{
		name:  "Samsung Galaxy Note 20",
		brand: "Samsung Galaxy Note 20",
		year:  2020,
		color: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}
	fmt.Println(p.displayData())
}

// Soal 3
func luasPersegi(sisi float64, showText bool) interface{} {
	if sisi == 0 {
		if showText {
			return "Maaf anda belum menginput sisi dari persegi"
		} else {
			return "nil"
		}
	}

	luas := sisi * sisi

	if showText {
		result := fmt.Sprintf("Luas persegi dengan sisi %.2f cm adalah %.2f cm", sisi, luas)
		return result
	} else {
		return luas
	}
}

func luas() {
	fmt.Println("----- SOAL 3 -----")
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))
}

// Soal 4

func hasil() {
	fmt.Println("----- SOAL 4 -----")
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	angkaPertama, ok := kumpulanAngkaPertama.([]int)
	if !ok {
		fmt.Println("kumpulanAngkaPertama bukan merupakan slice int")
		return
	}

	angkaKedua, ok := kumpulanAngkaKedua.([]int)
	if !ok {
		fmt.Println("kumpulanAngkaKedua bukan merupakan slice int")
		return
	}

	total := 0
	for _, angka := range angkaPertama {
		total += angka
	}
	for _, angka := range angkaKedua {
		total += angka
	}

	output := fmt.Sprintf("%s%s+%s=%d", prefix, stringifySlice(angkaPertama), stringifySlice(angkaKedua), total)
	fmt.Println(output)
}

func stringifySlice(slice []int) string {
	strSlice := make([]string, len(slice))
	for i, angka := range slice {
		strSlice[i] = fmt.Sprint(angka)
	}
	return strings.Join(strSlice, "+")
}
