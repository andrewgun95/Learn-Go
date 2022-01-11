package main

import (
	"fmt"
)

func main() {
	t := &tree{
		x: 8,
		l: &tree{
			x: 2,
			l: &tree{
				x: 8,
			},
			r: &tree{
				x: 7,
			},
		},
		r: &tree{
			x: 6,
		},
	}

	n := searchVisible(t)
	fmt.Println(n)

	r := checkSorting(3, 4, 7, 8, 10, 11)
	fmt.Println("Single swap operations", r)
}

func checkSorting(numb ...int) bool {
	count := 0
	for i := 0; i < len(numb)-1; i++ {
		if numb[i+1] < numb[i] {
			count++
		}
	}

	return count <= 2
}

type tree struct {
	x int
	l *tree
	r *tree
}

// try
// 1 + 2 + 0
// f(3)  = 2 + 0 + 0
// f(20) = 0 + 0 + 0
// f(21) = 0 + 0 + 0
// f(10) = 0 + 0 + 0

func searchVisible(t *tree) int {
	return search(t, map[int]bool{t.x: true})
}

func search(t *tree, result map[int]bool) int {
	if t == nil {
		// do nothing
	} else {
		if t.l != nil && t.l.x > t.x {
			result[t.l.x] = true
		}
		if t.r != nil && t.r.x > t.x {
			result[t.r.x] = true
		}

		search(t.l, result)
		search(t.r, result)
	}

	return len(result)
}

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	x := c_q
	y := r_q

	var counts int32
	var direcs = [8][2]int32{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1},
	}

	for _, v := range direcs {
		var dx int32 = v[0]
		var dy int32 = v[1]

		for i, j := x+dx, y+dy; i <= n && j <= n && i > 0 && j > 0; {
			if !isObstacle(obstacles, i, j) {
				counts++
			} else {
				break
			}

			i += dx
			j += dy
		}
	}

	return counts
}

func isObstacle(obstacles [][]int32, x int32, y int32) bool {
	for _, v := range obstacles {
		if v[1] == x && v[0] == y {
			return true
		}
	}

	return false
}
