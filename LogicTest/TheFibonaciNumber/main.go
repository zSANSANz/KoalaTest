package main

import "fmt"

func main() {
	fmt.Println(fibonacciTopDown(5))
	fmt.Println(fibonacciButtomUp(5))
}

func fibonacciTopDown(n int) int {
	if n== 0 || n == 1 {
		return n
	}
	return fibonacciTopDown(n-1) + fibonacciTopDown(n-2)
}

func fibonacciButtomUp(n int) int {
	if n== 0 || n == 1 {
		return n
	}
	var a int = 0

	var b int = 1

	temp := 0
	for i:=2; i<n; i++ {
		temp = a + b
		a = b
		b = temp	
	}
	return b
}