package main

import "fmt"

// Contract interface
type BangunDatar interface {
	// Methdo
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int  {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {
	persegi := Persegi{
		Sisi: 2,
	}

	fmt.Println("Luas Persegi, Sisi = 2 adalah", persegi.HitungLuas())
	fmt.Println("Luas Persegi, Sisi = 2 adalah", LuasBangunDatar(persegi))

	persegiPanjang := PersegiPanjang{
		Panjang: 2,
		Lebar: 3,
	}

	fmt.Println("Luas Persegi Panjang, Panjang = 2, Lebar = 3 adalah", persegiPanjang.HitungLuas())
	fmt.Println("Luas Persegi Panjang, Panjang = 2, Lebar = 3 adalah", LuasBangunDatar(persegiPanjang))
}

// bisa diberikan apapun parameter nya asal type mempunyai contract dengan interface dan mempunyai method HitungLuas()
func LuasBangunDatar(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}