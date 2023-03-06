package main

import "fmt"

func main() {

	x := []int{1, 4, 99}
	squats(3, x)

}

func squats(t int, x []int) {

	for i := 0; i < t; i++ {
		if i == t+1 {
			break
		}
		sum := x[i] * 15
		fmt.Println(sum)
	}
}
