package day01_test

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestPart02(t *testing.T) {
	// Open the file
	file, _ := os.Open("input.txt")
	defer file.Close()

	// create 2 in-memory arrays, one per column
	var col1 []float64
	var col2 []float64
	var total float64

	// loop through each line of the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		// convert each string to int
		val1, _ := strconv.ParseFloat(fields[0], 64)
		val2, _ := strconv.ParseFloat(fields[1], 64)

		// push converted values onto arrays
		col1 = append(col1, val1)
		col2 = append(col2, val2)
	}

	// sort columns lowest to highest
	sort.Float64s(col1)
	sort.Float64s(col2)

	// calculate similarity between values in columns
	for i := 0; i < len(col1); i++ {
		var occurs float64
		val1 := col1[i]

		// count number of times value from column 1 occur in column 2
		for z := 0; z < len(col2); z++ {
			val2 := col2[z]
			if val1 == val2 {
				occurs += 1
			}
		}

		total += val1 * occurs

		t.Logf("%f (* %f) = %f\n", val1, occurs, total)
	}

	t.Logf("=== total: %f ===\n", total)

	if total != 23150395 {
		t.Errorf("expected 23150395, got %f", total)
	}
}
