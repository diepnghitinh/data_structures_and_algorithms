package main

import "fmt"

/*
Công ty Anfin
Đề bài cho biết tìm số đảo trong 1 mảng 2 chiều

{1, 1, 0, 0, 0},
{0, 1, 0, 0, 1},
{1, 0, 0, 1, 1},
{0, 0, 0, 0, 0},
{1, 0, 1, 0, 1},

có 5 đảo, dc tính theo 8 hướng

*/

const row, column = 5, 5

func getDirections() [][]int {
	return [][]int{
		{-1, 0}, // Top
		{1, 0},  // Bottom
		{0, 1},  // Right
		{0, -1}, // Left

		{-1, -1}, // Top Left
		{-1, 1},  // Top Right
		{1, -1},  // Bottom Left
		{1, 1},   // Bottom Right
	}
}

func DFS(grid [][]int, visited [][]bool, i int, j int) {

	var stack = NewStack()

	//Push to stack and begin
	if !(visited[i][j]) && grid[i][j] == 1 {
		stack.Push(&Node{x: i, y: j, Value: grid[i][j]})
	}

	// Stack !IsEmpty
	for !(stack.count == 0) {
		popVal := stack.Pop()

		if grid[popVal.x][popVal.y] == 0 {
			continue
		}

		visited[popVal.x][popVal.y] = true

		for _, direction := range getDirections() {
			dx, dy := direction[0], direction[1]

			x := popVal.x + dx
			y := popVal.y + dy

			// Boundary check.
			if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
				continue
			}

			if visited[x][y] || grid[x][y] == 0 {
				continue
			}

			stack.Push(&Node{x: x, y: y, Value: grid[x][y]})
			visited[x][y] = true
		}

	}
}

func main() {

	// Init variable island
	var a = [][]int{
		{1, 1, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{1, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1},
	}
	visited := make([][]bool, len(a))
	for row := range visited {
		visited[row] = make([]bool, len(a[0]))
	}

	// Foreach cell
	var i, j, numIsland int
	for i = 0; i < row; i++ {
		for j = 0; j < column; j++ {

			if visited[i][j] {
				continue
			}

			if a[i][j] == 1 {
				numIsland++
				DFS(a, visited, i, j)
			}

		}
	}

	fmt.Println("Num island:")
	fmt.Println(numIsland)

}
