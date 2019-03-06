package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			fmt.Println("nilai :", i, e)
			switch {
			case i == 0:
				min, max = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{1, 2, 2, 4, 5}
	var mi, mx = getMinMax(numbers)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers, mi, mx)
}