package main

import (
	"fmt"
)

//返回a、b中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	x := 3
	y := 4
	z := 5
	max_xy := max(x, y)
	max_xz := max(x, z)

	fmt.Println("max(%d,%d) = %d\n", x, y, max_xy)
	fmt.Println("max(%d,%d) = %d\n", x, z, max_xz)
	fmt.Println("max(%d,%d) = %d\n", y, z, max(y, z))

}
