package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate_checksum(data [][]int) int {
	check_sum := 0

	for _, row := range data {
		var min, max int
		min = row[0]
		max = row[0]
		for _, col := range row {
			if col < min {
				min = col
			}
			if col > max {
				max = col
			}
		}
		diff := max - min
		check_sum += diff
	}

	return check_sum

}

func calculate_checksum_divisible_values(data [][]int) int {
	check_sum := 0

	for _, row := range data {
		len_row := len(row)

	FindPair:
		for i := 0; i < len_row; i++ {
			for j := 0; j < len_row; j++ {
				if i == j {
					continue
				}
				if row[i]%row[j] == 0 {
					check_sum += row[i] / row[j]
					break FindPair
				}
			}
		}
	}

	return check_sum

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var data [][]int

	line_num := 1

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)

		row_of_numbers := make([]int, len(fields))

		for i, f := range fields {
			number, err := strconv.Atoi(f)

			if err != nil {
				fmt.Printf("In line %d: field: %d: Can't convert %s to int\n", line_num, i+1, f)
				return
			}

			row_of_numbers[i] = number
		}

		data = append(data, row_of_numbers)
		line_num++
	}

	fmt.Printf("check_sum = %d\n", calculate_checksum(data))
	fmt.Printf("check_sum_divisible = %d\n", calculate_checksum_divisible_values(data))
}
