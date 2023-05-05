package main

import "fmt"

func main() {
	x := som(1, 2, 3)
	y := multi(10, 10)
	w := subt(5, 10)
	z := div(20)
	fmt.Println(x, y, w, z)
}

func som(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func subt(i ...int) int {
	total := 0
	for _, v := range i {
		total = v - total
	}
	return total
}

func multi(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}

func div(i ...int) int {
	total := 1
	for _, v := range i {
		total = v / total
	}
	return total
}
