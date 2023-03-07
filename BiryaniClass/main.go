package main

import "fmt"

func main() {

	x := []int{1, 1, 2, 2}
	y := []int{10, 15, 10, 15}
	biryaniClass(4, x, y)

}

func biryaniClass(t int, x, y []int) {
	for i := 0; i < t; i++ {
		if i == t {
			break
		}
		res := x[i] * y[i]
		fmt.Println(res)
	}
}
