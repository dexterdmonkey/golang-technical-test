package main

import (
	"fmt"
)

// 1. Function to check if a string is a palindrome
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// 2. Function to find the maximum value in an integer array without using built-in max() function
func findMax(arr []int) int {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

// 3. Function to print a triangle pattern with '*'
func printTriangle(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 4. Recursive function to calculate factorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("Is 'radar' a palindrome?", isPalindrome("radar"))
	fmt.Println("Is 'hello' a palindrome?", isPalindrome("hello"))
	fmt.Println("Is 'madam' a palindrome?", isPalindrome("madam"))

	arr1 := []int{3, 5, 1, 9, 2}
	fmt.Println("Maximum value in array", arr1, "is:", findMax(arr1))
	arr2 := []int{1, 5, 1, 0, 5}
	fmt.Println("Maximum value in array", arr2, "is:", findMax(arr2))

	fmt.Println("Triangle pattern of size 5:")
	printTriangle(5)
	fmt.Println("Triangle pattern of size 8:")
	printTriangle(8)

	fmt.Println("Factorial of 5 is:", factorial(5))
	fmt.Println("Factorial of 8 is:", factorial(8))

}
