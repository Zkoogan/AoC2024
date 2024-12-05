package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	sum_of_mid_pages := 0

	file := flag.String("file", "aoc_5_2024_test.txt", "the test data")

	flag.Parse()

	content, err := os.ReadFile(*file)

	if err != nil {
		log.Fatal("Error reading file: " + err.Error())
	}

	separated_data := strings.Split(string(content), "\r\n\r\n")

	page_order := strings.Split(separated_data[0], "\r\n")
	updates := strings.Split(separated_data[1], "\r\n")

	page_rules := create_page_rules(page_order)

	invalid_updates := []string{}

	for _, manual_order := range updates {
		is_valid, middleValue := processManualOrder(manual_order, page_rules)
		if is_valid {
			sum_of_mid_pages += middleValue
		} else {
			invalid_updates = append(invalid_updates, manual_order)
		}
	}
	fmt.Println(sum_of_mid_pages)
	calculate_invalid_sums(invalid_updates, page_rules)
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func create_page_rules(page_order []string) map[string][]string {
	page_rules := make(map[string][]string)

	for _, rule := range page_order {
		key_value := strings.Split(rule, "|")

		if _, exists := page_rules[key_value[0]]; !exists {
			page_rules[key_value[0]] = []string{}
		}
		page_rules[key_value[0]] = append(page_rules[key_value[0]], key_value[1])
	}
	return page_rules
}

func calculate_invalid_sums(invalid_updates []string, page_rules map[string][]string) {
	final_invalid_sum := 0
	for _, update := range invalid_updates {
		order := make(map[int]string)
		update_list := strings.Split(update, ",")
		for _, page := range update_list {
			key_value := 0
			for _, child := range page_rules[page] {
				if contains(update_list, child) {
					key_value++
				}
			}
			order[key_value] = page
		}
		final_invalid_sum_part, _ := strconv.ParseInt(order[len(order)/2], 10, 32)
		final_invalid_sum += int(final_invalid_sum_part)
	}
	fmt.Println(final_invalid_sum)
}

func processManualOrder(manual_order string, page_rules map[string][]string) (bool, int) {
	pages_of_manual := strings.Split(manual_order, ",")
	current_page := pages_of_manual[0]

	for i := 1; i < len(pages_of_manual); i++ {
		if contains(page_rules[current_page], pages_of_manual[i]) {
			current_page = pages_of_manual[i]
		} else {
			return false, 0
		}
	}

	middleIndex := len(pages_of_manual) / 2
	middleValue, _ := strconv.ParseInt(pages_of_manual[middleIndex], 10, 32)
	return true, int(middleValue)
}
