package main

import "fmt"

func main() {

	x := []int{1, 3, 7, 10}

	fitness(4, x)

}

func fitness(t int, x []int) {
	for i := 0; i < t; i++ {
		if i == t {
			break
		}
		sum := (x[i] + x[i]) * 5
		fmt.Println(sum)
	}
}
