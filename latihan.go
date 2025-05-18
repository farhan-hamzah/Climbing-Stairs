package main
import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	first := 1
	second := 2
	var current int

	for i := 3; i <= n; i++ {
		current = first + second
		first = second
		second = current
	}

	return current
}

func main() {
	fmt.Println(climbStairs(2)) // Output: 2
	fmt.Println(climbStairs(3)) // Output: 3
	fmt.Println(climbStairs(5)) // Output: 8
}
