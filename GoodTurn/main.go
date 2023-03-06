package main

import "fmt"

//The code is doing the following:
//1. Scan the size of the array
//2. Loop through the array and scan in two numbers
//3. Check if the sum of the two numbers is greater than 6
//4. If it is, print "YES"
//5. If not, print "NO"

func main() {

	x := []int{1, 3, 4, 2}
	y := []int{4, 4, 2, 6}
	goodTurn(4, x, y)

}

func goodTurn(t int, x, y []int) {
	for i := 0; i < t; i++ {
		if i == t+1 {
			break
		}
		if x[i]+y[i] <= 6 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
