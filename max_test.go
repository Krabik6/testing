package main

import "testing"

func TestMax(t *testing.T) {
	// Arrange
	testTable := []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{-32, -16, -8, -4, -2, 0, 2, 4, 8, 16, 32, 64},
			expected: 64,
		},
		{
			numbers:  []int{-32, -16},
			expected: -16,
		},
		{
			numbers:  []int{3, 4, 5},
			expected: 5,
		},
		{
			numbers:  []int{-12222222, 123, 1234},
			expected: 1234,
		},
		{
			numbers:  []int{-12222222, 123, 1234, 1234},
			expected: 1234,
		},
		{
			numbers:  []int{},
			expected: 0,
		},
	}

	// Act
	for _, testCase := range testTable {
		result := Max(testCase.numbers)

		// Assert
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expected %d, got %d", testCase.expected, result)
		}
	}

}

func TestMaxIndex(t *testing.T) {
	// Arrange
	testTable := []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{-32, -16, -8, -4, -2, 0, 2, 4, 8, 16, 32, 64},
			expected: 11,
		},
		{
			numbers:  []int{-32, -16},
			expected: 1,
		},
		{
			numbers:  []int{3, 4, 5},
			expected: 2,
		},
		{
			numbers:  []int{-12222222, 123, 1234},
			expected: 2,
		},
		{
			numbers:  []int{-12222222, 123, 1234, 1234},
			expected: 2,
		},
		{
			numbers:  []int{},
			expected: 0,
		},
	}

	// Act
	for _, testCase := range testTable {
		result := MaxIndex(testCase.numbers)

		// Assert
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expected %d, got %d", testCase.expected, result)
		}
	}
}
