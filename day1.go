package main

import (
	"fmt"
	"os"
	"unicode"
)

func to_int_table(s string) []int {
	t := make([]int, len(s))

	for i, r := range s {
		t[i] = int(r - '0')
	}

	return t
}

func validate_string(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true

}

func sum_digits_part_one(t []int) int {
	return sum_when_match_at_offset(t, 1)
}

func sum_digits_part_two(t []int) int {
	table_len := len(t)
	half_len := table_len / 2

	return sum_when_match_at_offset(t, half_len)
}

func sum_when_match_at_offset(t []int, offset int) int {
	sum := 0
	table_len := len(t)
	for i := 0; i < table_len; i++ {
		j := (i + offset) % table_len
		if t[i] == t[j] {
			sum += t[i]
		}
	}
	return sum
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage:\n%s captcha\n", os.Args[0])
		return
	}

	if !validate_string(os.Args[1]) {
		fmt.Printf("Input string %s is not a captcha\n", os.Args[1])
	}

	t := to_int_table(os.Args[1])

	sum := sum_digits_part_one(t)
	fmt.Printf("part_one = %d\n", sum)

	sum = sum_digits_part_two(t)
	fmt.Printf("part_two = %d\n", sum)
}
