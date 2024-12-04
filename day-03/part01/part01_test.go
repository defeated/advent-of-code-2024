package day03_test_01

import (
	"os"
	"regexp"
	"strconv"
	"testing"
)

func TestPart01(t *testing.T) {
	// Open the file
	file, _ := os.ReadFile("../input.txt")

	total := 0

	// parse string for strict mul(x,y) pattern and capture groups for each
	// numeric arg
	mul := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// get all captured groups from each match in the file
	// and multiple the two args together
	for _, res := range mul.FindAllSubmatch(file, -1) {
		x, _ := strconv.Atoi(string(res[1]))
		y, _ := strconv.Atoi(string(res[2]))

		total += x * y
	}

	t.Logf("=== total: %v ===\n", total)

	expected := 166905464
	if total != expected {
		t.Errorf("expected %v, got %v", expected, total)
	}
}
