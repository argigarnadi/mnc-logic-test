package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input total belanja : ")
	scanner.Scan()
	totalBelanja, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Input jumlah uang dibayar : ")
	scanner.Scan()
	jumlahUang, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}

	kembalian, ok := HitungKembalian(totalBelanja, jumlahUang)
	if !ok {
		fmt.Println("False, kurang bayar")
		return
	}

	fmt.Println("Pecahan Uang :")
	for _, k := range kembalian {
		if k.Nominal < 1000 {
			fmt.Println(fmt.Sprintf("%d koin %d", k.Jumlah, k.Nominal))
		} else {
			fmt.Println(fmt.Sprintf("%d lembar %d", k.Jumlah, k.Nominal))
		}

	}
}

type Kembalian struct {
	Nominal int
	Jumlah  int
}

func HitungKembalian(totalBelanja, uangBayar int) ([]Kembalian, bool) {
	if uangBayar < totalBelanja {
		return nil, false
	}

	kembalian := uangBayar - totalBelanja
	fmt.Println(fmt.Sprintf("kembalian yang harus diberikan kasir : %d,", kembalian))
	kembalian = (kembalian / 100) * 100
	fmt.Println(fmt.Sprintf("dibulatkan menjadi : %d", kembalian))
	pecahanUang := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}

	var result []Kembalian
	for _, p := range pecahanUang {
		jumlah := kembalian / p
		kembalian -= jumlah * p
		if jumlah > 0 {
			result = append(result, Kembalian{p, jumlah})
		}
	}

	return result, true
}
