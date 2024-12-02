package day01_test

import (
	"bufio"
	"math"
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

	// compare distance between lowest values
	for i := 0; i < len(col1); i++ {
		val1, val2 := col1[i], col2[i]
		diff := math.Abs(val1 - val2)
		total += diff

		t.Logf("%f -> %f (+ %f) = %f\n", val1, val2, diff, total)
	}

	t.Logf("=== total: %f ===\n", total)

	if total != 936063 {
		t.Errorf("expected 936063, got %f", total)
	}
}
