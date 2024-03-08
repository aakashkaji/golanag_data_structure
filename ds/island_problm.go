package main

import "fmt"

func dfs(grid [][]byte, i, j int) {

	// todo check valid move
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}

	grid[i][j] = 0 //mark as visited

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)

}

func noIsland(grid [][]byte) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count += 1
				dfs(grid, i, j)
			}
		}
	}

	return count

}

func main() {

	mat := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'0', '1', '0', '0', '1'},
		{'1', '0', '0', '1', '1'},
		{'0', '0', '0', '0', '0'},
		{'1', '0', '1', '0', '0'},
	}

	numIslands := noIsland(mat)
	fmt.Println("Number of islands:", numIslands)
}
