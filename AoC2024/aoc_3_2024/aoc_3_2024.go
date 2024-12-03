package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	re := regexp.MustCompile(`mul\((\d{1,3},\d{1,3})\)`)
	do_re := regexp.MustCompile(`do\(\)`)
	dont_re := regexp.MustCompile(`don't\(\)`)
	product := 0

	file := flag.String("file", "aoc_3_2024_test.txt", "the test data")
	isPart1 := flag.Bool("isPart1", true, "feature flag to enable part 2 parsing")

	flag.Parse()

	content, err := os.ReadFile(*file)

	if err != nil {
		log.Fatal(err)
	}

	string_content := string(content)

	matches := re.FindAllStringSubmatch(string_content, -1)
	if matches == nil {
		log.Fatal("No matches found")
	}

	do_start_indices := []int{}
	for _, indices := range do_re.FindAllStringIndex(string_content, -1) {
		do_start_indices = append(do_start_indices, indices[0])
	}

	dont_start_indices := []int{}
	for _, indices := range dont_re.FindAllStringIndex(string_content, -1) {
		dont_start_indices = append(dont_start_indices, indices[0])
	}

	mul_start_indices := []int{}
	for _, indices := range re.FindAllStringIndex(string_content, -1) {
		mul_start_indices = append(mul_start_indices, indices[0])
	}

	for index, element := range matches {
		if *isPart1 {
			product += int(getProductOfMatch(element[1]))
		} else {

			do_index, found_do := largestSmallerNumber(do_start_indices, mul_start_indices[index])
			dont_index, found_dont := largestSmallerNumber(dont_start_indices, mul_start_indices[index])

			if found_do && found_dont {
				if do_index > dont_index {
					product += int(getProductOfMatch(element[1]))
				}
			} else if found_dont {
			} else {
				product += int(getProductOfMatch(element[1]))
			}
		}
	}

	fmt.Println(product)

}

func getProductOfMatch(match string) int64 {
	number_strings := strings.Split(match, ",")

	x, err := strconv.ParseInt(number_strings[0], 10, 32)

	if err != nil {
		log.Fatal("Error found when parsing numbers" + string(x))
	}

	y, err2 := strconv.ParseInt(number_strings[1], 10, 32)

	if err2 != nil {
		log.Fatal("Error found when parsing numbers" + string(y))
	}

	return x * y
}

func largestSmallerNumber(sortedList []int, x int) (int, bool) {

	if len(sortedList) == 0 || x <= sortedList[0] {
		return 0, false
	}

	low, high := 0, len(sortedList)-1
	var result int
	found := false

	for low <= high {
		mid := low + (high-low)/2
		if sortedList[mid] < x {
			result = sortedList[mid]
			found = true
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result, found
}
