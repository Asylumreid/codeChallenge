package main

// Using mathematical formula to calculate the summation to n, O(1) time complexity
func sum_to_n_a(n int) int {
	if n <= 0 {
		return 0
	}
	return n * (n + 1) / 2
}

// Using a loop to calculate the summation to n, O(n) time complexity
func sum_to_n_b(n int) int {
	if n <= 0 {
		return 0
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// Using recursion to calculate summation to n, O(n) time complexity - can cause a stack overflow for large n
func sum_to_n_c(n int) int {
	if n <= 0 {
		return 0
	}
	return n + sum_to_n_c(n-1)
}
