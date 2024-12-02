package day01_test

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestPart01(t *testing.T) {
	// Open the file
	file, _ := os.Open("input.txt")
	defer file.Close()

	// create 2 in-memory arrays, one per column
	var col1 []int
	var col2 []int
	total := 0

	// loop through each line of the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		// convert each string to int
		val1, _ := strconv.Atoi(fields[0])
		val2, _ := strconv.Atoi(fields[1])

		// push converted values onto arrays
		col1 = append(col1, val1)
		col2 = append(col2, val2)
	}

	// sort columns lowest to highest1	`Q`
	sort.Ints(col1)
	sort.Ints(col2)

	// compare distance between lowest values
	for i := 0; i < len(col1); i++ {
		val1, val2 := col1[i], col2[i]

		// go authors don't "believe" in an abs() function?
		// sigh
		var diff int
		if val1 >= val2 {
			diff = val1 - val2
		} else {
			diff = val2 - val1
		}

		total += diff

		fmt.Printf("%v -> %v (+ %v) = %v\n", val1, val2, diff, total)
	}

	fmt.Printf("=== total: %v ===\n", total)

	if total != 936063 {
		t.Errorf("expected 936063, got %v", total)
	}
}
