package main

func add(x, y int) (z int) {
	defer func() {
		println(z)
	}()
	z = x + y
	return z + 200
}

func main() {
	println(add(1, 2))
}
