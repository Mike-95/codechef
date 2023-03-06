package main

import "fmt"

func main() {
	x := []int{21, 25, 22, 20, 28}
	y := []int{34, 31, 29, 40, 29}
	a := []int{30, 31, 25, 15, 28}

	ageLimit(5, x, y, a)
}

func ageLimit(t int, x, y, a []int) {
	for i := 0; i < t; i++ {
		if i == t {
			break
		}
		//var x, y, a int
		//fmt.Scanf("%d", &x)
		//fmt.Scanf("%d", &y)
		//fmt.Scanf("%d", &a)

		if a[i] >= x[i] && a[i] < y[i] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
