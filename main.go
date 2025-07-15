package main

import "fmt"

func main() {
	soalNomor1(5)
	fmt.Println()
	soalNomor2(5)
	fmt.Println()
	soalNomor3(5)
}

func soalNomor1(n int) {
	for i := 0; i < n; i ++ {
		for j, k := n, 1; j > 0; j, k = j-1, k+1 {
			if i % 2 == 0 {
				fmt.Print(j)
			} else {
				fmt.Print(k)
			}
		} 
		fmt.Println()
	}
}

func soalNomor2(n int) {
	for i := 0; i < n; i++ {
		for j, k := 0, 0; j < n; j, k = j+1, k+1 {
			if j%2 == 0 {
				fmt.Print(i + 1)
			} else {
				fmt.Print(n - i)
			}
		}
		fmt.Println()
	}
}

func soalNomor3(n int) {
	for i := 0; i < n; i++ {
		// kehabisan waktu
		res1 := []int{}
		res2 := []int{}
		if i % 2 == 0 {
			for j := 1; j < n; j++ {
				if j % 2 == 0 {
					res1 = append(res1, j)
				}
			}
		} else {
			for j := 1; j < n; j++ {
				if j % 2 != 0 {
					res2 = append(res2, j)
				}
			}
		}
		// perlu ada perbaikan 
		// perlu print dari array yang sudah dibuat
	}
}