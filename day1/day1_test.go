package main

import "testing"

func TestExaplesPartOne(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	}

	for _, tc := range cases {
		int_table := to_int_table(tc.in)
		result := sum_digits_part_one(int_table)
		if result != tc.out {
			t.Errorf("sum_digits_part_one(%#v) resulted in %d but expected %d", int_table, result, tc.out)
		}
	}
}

func TestExaplesPartTwo(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}

	for _, tc := range cases {
		int_table := to_int_table(tc.in)
		result := sum_digits_part_two(int_table)
		if result != tc.out {
			t.Errorf("sum_digits_part_two(%#v) resulted in %d but expected %d", int_table, result, tc.out)
		}
	}
}
