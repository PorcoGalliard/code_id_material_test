package main

import "fmt"

func main() {
	soalNomor1(3)
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

}

// func soalNomor3(n int) {

// }