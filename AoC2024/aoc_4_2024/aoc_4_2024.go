package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Tuple struct {
	x, y int
}

type Node struct {
	pos              Tuple
	relevantNeighbor []Node
	direction        Tuple
}

type LetterValue struct {
	X int32
	S int32
	A int32
	M int32
}

func main() {

	file := flag.String("file", "aoc_4_2024_test.txt", "the test data")

	flag.Parse()

	letterValues := LetterValue{X: 88, S: 83, A: 65, M: 77}

	content, err := os.ReadFile(*file)

	if err != nil {
		log.Fatal("Error reading file: " + err.Error())
	}

	found_start_pos_1 := []Node{}
	found_start_pos_2 := []Tuple{}

	matrix := stringToCharMatrix(string(content))
	for i, row := range matrix {
		for j, col := range row {
			if col == letterValues.X {
				found_start_pos_1 = append(found_start_pos_1, Node{pos: Tuple{x: j, y: i}, relevantNeighbor: []Node{}, direction: Tuple{}})
			}
			if col == letterValues.A {
				found_start_pos_2 = append(found_start_pos_2, Tuple{x: j, y: i})
			}
		}
	}

	solve_part_1(letterValues, matrix, found_start_pos_1)
	solve_part_2(matrix, found_start_pos_2)
}

func solve_part_2(matrix [][]rune, found_start_pos []Tuple) {
	num_valid_x_mas := 0
	for _, x := range found_start_pos {
		if isValidMasX(matrix, x) {
			num_valid_x_mas += 1
		}
	}

	fmt.Println(num_valid_x_mas)
}

func isValidMasX(matrix [][]rune, pos Tuple) bool {
	if pos.y >= 1 && pos.y < len(matrix)-1 && pos.x >= 1 && pos.x < len(matrix[pos.y])-1 {
		corners := []Tuple{
			{-1, -1},
			{-1, 1},
			{1, -1},
			{1, 1},
		}

		corner_product := 1

		for _, corner := range corners {
			corner_product *= int(matrix[pos.y+corner.y][pos.x+corner.x])
		}

		if corner_product != 40844881 { //88^2 * 77^2 checks that there are 2 corners of the box surrounding an A that has 2 M's and 2 S's, should use letterValue types
			return false
		}

		if int(matrix[pos.y+corners[0].y][pos.x+corners[0].x])*int(matrix[pos.y+corners[3].y][pos.x+corners[3].x]) != 6391 { //88*77 checks that the M's and S's are not on the diagonal, should use letterValue types
			return false
		}

		return true
	}
	return false
}

func solve_part_1(letterValues LetterValue, matrix [][]rune, found_start_pos []Node) {
	xmas_sum := 0

	for _, found_x := range found_start_pos {
		found_x.relevantNeighbor = findStartPoints(matrix, found_x.pos, letterValues.M)
		for _, found_m := range found_x.relevantNeighbor {
			found_m.relevantNeighbor = findSurroundingOccurrences(matrix, found_m.pos, found_m.direction, letterValues.A)
			for _, found_a := range found_m.relevantNeighbor {
				found_a.relevantNeighbor = findSurroundingOccurrences(matrix, found_a.pos, found_a.direction, letterValues.S)
				xmas_sum += len(found_a.relevantNeighbor)
			}
		}
	}
	fmt.Println(xmas_sum)
}

func stringToCharMatrix(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	return matrix
}

func findStartPoints(matrix [][]rune, pos Tuple, target rune) []Node {
	directions := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	var occurrences []Node

	for _, d := range directions {
		newX, newY := pos.x+d[1], pos.y+d[0]

		if newY >= 0 && newY < len(matrix) && newX >= 0 && newX < len(matrix[newY]) {
			if matrix[newY][newX] == target {
				occurrences = append(occurrences, Node{pos: Tuple{x: newX, y: newY}, relevantNeighbor: []Node{}, direction: Tuple{d[1], d[0]}})
			}
		}
	}
	return occurrences
}

func findSurroundingOccurrences(matrix [][]rune, pos Tuple, direction Tuple, target rune) []Node {
	var occurrences []Node

	if direction.x != 0 || direction.y != 0 {
		newX, newY := pos.x+direction.x, pos.y+direction.y
		if newY >= 0 && newY < len(matrix) && newX >= 0 && newX < len(matrix[newY]) {
			if matrix[newY][newX] == target {
				occurrences = append(occurrences, Node{pos: Tuple{x: newX, y: newY}, relevantNeighbor: []Node{}, direction: Tuple{newX - pos.x, newY - pos.y}})
			}
		}
		return occurrences
	}

	return occurrences
}
