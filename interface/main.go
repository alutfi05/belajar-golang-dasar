package main

import "fmt"

// Contract interface
type Luas interface {
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

	persegiPanjang := PersegiPanjang{
		Panjang: 2,
		Lebar: 3,
	}

	fmt.Println("Luas Persegi Panjang, Panjang = 2, Lebar = 3 adalah", persegiPanjang.HitungLuas())
}