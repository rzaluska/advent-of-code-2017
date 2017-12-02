package main

import "testing"

func TestDayTwoTestPartOne(t *testing.T) {
	cases := []struct {
		in  [][]int
		out int
	}{
		{[][]int{{5, 1, 9, 5}, {7, 3, 5}, {2, 4, 6, 8}}, 18},
	}

	for _, c := range cases {
		result := calculate_checksum(c.in)

		if result != c.out {
			t.Errorf("calculate_checksum(%#v) got %d but expected %d\n", c.in, result, c.out)
		}
	}
}

func TestDayTwoTestPartTwo(t *testing.T) {
	cases := []struct {
		in  [][]int
		out int
	}{
		{[][]int{{5, 9, 2, 8}, {9, 4, 7, 3}, {3, 8, 6, 5}}, 9},
	}

	for _, c := range cases {
		result := calculate_checksum_divisible_values(c.in)

		if result != c.out {
			t.Errorf("calculate_checksum_divisible_values(%#v) got %d but expected %d\n", c.in, result, c.out)
		}
	}
}
