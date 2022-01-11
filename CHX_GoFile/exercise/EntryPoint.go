package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	pwd, _ := os.Getwd()

	path := pwd + "/exercise"

	data, err := readLines(path + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var grid [][]byte
	var rows, cols int
	n := 1
	for i := 0; i < len(data); {
		if i > 0 {
			if isInt(data[i]) {
				rows, _ = strconv.Atoi(data[i])
				cols, _ = strconv.Atoi(data[i+1])
				grid = make([][]byte, rows)
				for j := range grid {
					grid[j] = make([]byte, cols)
				}
				i += 2
			} else {
				for j := 0; j < len(grid); j++ {
					arr := []byte(data[i+j])
					grid[j] = arr
				}

				word := []byte(data[i+len(grid)])
				result := wordSearch(grid, word)

				fmt.Printf("%d\tWord %q : %d\n", n, word, result)
				fmt.Println(printGrid(grid))
				n++

				i += (len(grid) + 1)
			}
		} else {
			i++
		}
	}
}

func printGrid(grid [][]byte) string {
	str := ""
	for i := 0; i < len(grid); i++ {
		str += fmt.Sprintf("\t%d %c", i, grid[i])
		if i != len(grid)-1 {
			str += "\n"
		}
	}
	return str
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func wordSearch(grid [][]byte, word []byte) int {
	resultSearch := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == word[0] {
				resultSearch += searchWordInGrid(i, j, grid, word)
			}
		}
	}
	return resultSearch
}

func searchWordInGrid(i int, j int, grid [][]byte, word []byte) int {
	var counts [8]int
	for x := 1; x < len(word); x++ {
		// vertical
		if i+(len(word)-1) < len(grid) {
			if grid[i+x][j] == word[x] {
				counts[0]++
			}
		}
		if i-(len(word)-1) > -1 {
			if grid[i-x][j] == word[x] {
				counts[1]++
			}
		}
		// horizontal
		if j+(len(word)-1) < len(grid[i]) {
			if grid[i][j+x] == word[x] {
				counts[2]++
			}
		}
		if j-(len(word)-1) > -1 {
			if grid[i][j-x] == word[x] {
				counts[3]++
			}
		}
		// diagonal
		if (i+(len(word)-1) < len(grid)) && (j+(len(word)-1) < len(grid[i])) {
			if grid[i+x][j+x] == word[x] {
				counts[4]++
			}
		}
		if (i-(len(word)-1) > -1) && (j-(len(word)-1) > -1) {
			if grid[i-x][j-x] == word[x] {
				counts[5]++
			}
		}
		if (i+(len(word)-1) < len(grid)) && (j-(len(word)-1) > -1) {
			if grid[i+x][j-x] == word[x] {
				counts[6]++
			}
		}
		if (i-(len(word)-1) > -1) && (j+(len(word)-1) < len(grid[i])) {
			if grid[i-x][j+x] == word[x] {
				counts[7]++
			}
		}
	}

	result := 0
	for i := 0; i < len(counts); i++ {
		result += counts[i] / (len(word) - 1)
	}
	return result
}
