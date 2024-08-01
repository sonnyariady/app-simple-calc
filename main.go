package main

import (
	"fmt"

	sclc "github.com/sonnyariady/go-simple-calc"
)

func main() {
	a := float32(12)
	b := float32(4)
	fmt.Println(sclc.Jumlah(a, b))
	fmt.Println(sclc.Kali(a, b))
	fmt.Println(sclc.Bagi(a, b))
	fmt.Println(sclc.Kurang(a, b))
}
