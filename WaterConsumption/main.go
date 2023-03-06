package main

import "fmt"

func main() {

	x := []int{2999, 1450, 2000}
	waterConsumption(3, x)

}

func waterConsumption(t int, x []int) {

	for i := 0; i < t; i++ {
		if i == t+1 {
			break
		}
		if x[i] < 2000 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
