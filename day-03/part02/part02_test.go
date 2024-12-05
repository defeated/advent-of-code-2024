package day03_test_02

import (
	"os"
	"regexp"
	"strconv"
	"testing"
)

func TestPart02(t *testing.T) {
	// Open the file
	file, _ := os.ReadFile("../input.txt")
	total := 0

	// parse string for strict mul(x,y) pattern and capture groups for each
	// numeric arg
	mul := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|(do\(\)|don't\(\))`)

	// get all captured groups from each match in the file
	// and multiple the two args together
	doing := true
	for _, res := range mul.FindAllString(string(file), -1) {

		// check if stop token encountered, or keep going
		switch res {
		case "don't()":
			doing = false
		case "do()":
			doing = true
		}

		// if keep going, parse and multiply args, then add to total
		if doing {
			sub := mul.FindStringSubmatch(res)
			x, _ := strconv.Atoi(sub[1])
			y, _ := strconv.Atoi(sub[2])

			total += x * y
		}
	}

	t.Logf("=== total: %v ===\n", total)

	expected := 72948684
	if total != expected {
		t.Errorf("expected %v, got %v", expected, total)
	}
}
