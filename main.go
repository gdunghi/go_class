package main



var plush = add

func add(x int, y int) int {
	return x + y
}

func main() {
	println(plush(43,12))
}
