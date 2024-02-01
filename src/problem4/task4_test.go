package main

import "testing"

// Test cases for sum_to_n_a function
func TestSumToNA(t *testing.T) {
	if sum_to_n_a(5) != 15 {
		t.Error("sum_to_n_a(5) should be 15")
	}
	if sum_to_n_a(325) != 52975 {
		t.Error("sum_to_n_a(10) should be 52975")
	}
}

// Test cases for sum_to_n_b function
func TestSumToNB(t *testing.T) {
	if sum_to_n_b(5) != 15 {
		t.Error("sum_to_n_b(5) should be 15")
	}
	if sum_to_n_b(325) != 52975 {
		t.Error("sum_to_n_b(10) should be 52975")
	}
}

// Test cases for sum_to_n_c function
func TestSumToNC(t *testing.T) {
	if sum_to_n_c(5) != 15 {
		t.Error("sum_to_n_c(5) should be 15")
	}
	if sum_to_n_c(325) != 52975 {
		t.Error("sum_to_n_c(10) should be 52975")
	}
}
