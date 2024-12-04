package day02_test_02

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestPart02(t *testing.T) {
	// Open the file
	file, _ := os.Open("../input.txt")
	defer file.Close()

	total := 0

	// loop through each line of the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		// convert space-separated array of strings to ints
		fields := strings.Fields(scanner.Text())
		nummies := make([]int, len(fields))
		for i, field := range fields {
			nummies[i], _ = strconv.Atoi(field)
		}

		// check if list of numbers is all ascending or all descending
		// and each subsequent number is within a safe range
		if ascending(&nummies) || descending(&nummies) {
			total++
		}
	}

	t.Logf("=== total: %v ===\n", total)

	expected := 386
	if total != expected {
		t.Errorf("expected %v, got %v", expected, total)
	}
}

func ascending(list *[]int) bool {
	return isSafe(list, "asc") || isSafeDamp(list, "asc")
}

func descending(list *[]int) bool {
	return isSafe(list, "desc") || isSafeDamp(list, "desc")
}

func safeRange(x int) bool {
	return x >= 1 && x <= 3
}

func isSafe(list *[]int, dir string) bool {
	for i := 1; i < len(*list); i++ {
		curr := (*list)[i]
		prev := (*list)[i-1]

		var diff int
		var valid bool

		switch dir {
		case "asc":
			diff = curr - prev
			valid = curr < prev
		case "desc":
			diff = prev - curr
			valid = curr > prev
		}

		if valid || !safeRange(diff) {
			return false
		}
	}
	return true
}

// create a new array, removing one element at a time, and checking if that
// makes the list safe instead.
func isSafeDamp(list *[]int, dir string) bool {
	for i := 1; i <= len(*list); i++ {
		retry := slices.Delete(slices.Clone(*list), i-1, i)
		if isSafe(&retry, dir) {
			return true
		}
	}
	return false
}
