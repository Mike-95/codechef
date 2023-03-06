package main

import "fmt"

func main() {
	x := []int{4, 8, 5, 2}
	y := []int{2, 7, 1, 1}
	savingTax(4, x, y)

}

func savingTax(t int, x, y []int) {
	for i := 0; i < t; i++ {
		if i == t+1 {
			break
		}
		sum := x[i] - y[i]
		fmt.Println(sum)
	}
}
