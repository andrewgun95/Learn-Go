package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"unicode"
)

type cell struct {
	x, y int
	char byte
}

func main() {
	pwd, _ := os.Getwd()

	path := pwd + "/exercise"

	data, err := readLines(path + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	n := 1
	var grid [][]byte
	var rows, cols int
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
				// Initialize grid
				for y := 0; y < len(grid); y++ {
					arr := []byte(data[i+y])
					grid[y] = arr
				}

				// Get all position of all characters in each line
				charPositions := getCharacters(grid)

				// Find the nearest characters of the current characters
				faction := map[byte]int{}
				contest := map[string]int{}
				if len(charPositions) > 0 {
					visited := []cell{}
					for _, pos := range charPositions {
						if !find(visited, pos) {
							v, result := findAllies(pos, grid)
							if result {
								if _, ok := faction[pos.char]; ok {
									faction[pos.char]++
								} else {
									faction[pos.char] = 1
								}
							} else {
								if _, ok := contest["contested"]; ok {
									contest["contested"]++
								} else {
									contest["contested"] = 1
								}
							}

							visited = v
						}
					}
				}

				fmt.Println(printGrid(grid))
				// Sorting the keys
				keys := make([]int, 0, len(faction))
				for k := range faction {
					keys = append(keys, int(k))
				}
				sort.Ints(keys)

				for _, k := range keys {
					fmt.Printf("Faction %c %d\n", k, faction[byte(k)])
				}

				fmt.Println("Contested", contest["contested"])

				n++

				i += len(grid)
			}
		} else {
			i++
		}
	}

}

func findAllies(current cell, grid [][]byte) ([]cell, bool) { // Visited cells
	visited := []cell{current}
	stack := []cell{current}

	var currChar byte = grid[current.y][current.x]
	var result bool = true
	for {
		nextPos := getNext(current, grid, visited)
		if len(nextPos) > 0 {
			next := nextPos[rand.Intn(len(nextPos))]

			visited = append(visited, next)

			var nextChar byte = grid[next.y][next.x]
			// Find enemy
			if (nextChar >= 97 && nextChar <= 122) && (nextChar != currChar) { // Lower character 97 to 122
				result = false
			}

			current = next
			stack = append(stack, current)
		} else {
			index := len(stack) - 1
			if index == -1 {
				return visited, result
			}

			last := stack[index]
			current = last
			// Remove last element in stack
			stack = stack[:index]
		}
	}
}

func getNext(current cell, grid [][]byte, visited []cell) []cell {
	result := []cell{}
	if current.y+1 < len(grid) {
		nextChar := grid[current.y+1][current.x]
		nextPos := cell{x: current.x, y: current.y + 1, char: nextChar}
		if nextChar != 35 && !find(visited, nextPos) { // Is not wall and not yet visited
			result = append(result, nextPos)
		}
	}
	if current.y-1 > -1 {
		nextChar := grid[current.y-1][current.x]
		nextPos := cell{x: current.x, y: current.y - 1, char: nextChar}
		if nextChar != 35 && !find(visited, nextPos) {
			result = append(result, nextPos)
		}
	}
	if current.x+1 < len(grid[current.y]) {
		nextChar := grid[current.y][current.x+1]
		nextPos := cell{x: current.x + 1, y: current.y, char: nextChar}
		if nextChar != 35 && !find(visited, nextPos) {
			result = append(result, nextPos)
		}
	}
	if current.x-1 > -1 {
		nextChar := grid[current.y][current.x-1]
		nextPos := cell{x: current.x - 1, y: current.y, char: nextChar}
		if nextChar != 35 && !find(visited, nextPos) {
			result = append(result, nextPos)
		}
	}
	return result
}

func getCharacters(grid [][]byte) []cell {
	result := []cell{}
	for yChar, _ := range grid {
		for xChar, _ := range grid[yChar] {
			v := grid[yChar][xChar]
			if v >= 97 && v <= 122 { // Lower character 97 to 122
				result = append(result, cell{x: xChar, y: yChar, char: v})
			}
		}
	}

	return result
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
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

func find(slice []cell, val cell) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func printGrid(grid [][]byte) string {
	str := ""
	for i := 0; i < len(grid); i++ {
		str += fmt.Sprintf("\t%d\t%c", i, grid[i])
		if i != len(grid)-1 {
			str += "\n"
		}
	}
	return str
}
