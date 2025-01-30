package exercises

import "fmt"

func Fibonacci(n int) {
	// 0, 1, 1, 2, 3, 4, 4
	i1 := 0
	fmt.Print(i1, " ")
	i2 := 1
	fmt.Print(i2, " ")
	var num int
	for i := 0; i < n-2; i++ {
		num = i1 + i2
		fmt.Print(num, " ")
		i1 = i2
		i2 = num
	}
}
