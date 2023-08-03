package game

// type Cell struct {
// 	row, col int
// 	path     [][]int
// }

// func findShortestPath(matrix [][]int, startRow, startCol, endRow, endCol int) [][]int {
// 	queue := []Cell{{startRow, startCol, [][]int{{startRow, startCol}}}}
// 	visited := make([][]bool, len(matrix))
// 	for i := range visited {
// 		visited[i] = make([]bool, len(matrix[0]))
// 	}

// 	for len(queue) > 0 {
// 		cell := queue[0]
// 		queue = queue[1:]

// 		if cell.row == endRow && cell.col == endCol {
// 			return cell.path
// 		}

// 		if cell.row < 0 || cell.row >= len(matrix) ||
// 			cell.col < 0 || cell.col >= len(matrix[0]) ||
// 			matrix[cell.row][cell.col] == 1 || visited[cell.row][cell.col] {
// 			continue
// 		}

// 		visited[cell.row][cell.col] = true

// 		// Add neighboring cells to the queue
// 		neighs := []struct{ row, col int }{{cell.row + 1, cell.col}, {cell.row - 1, cell.col},
// 			{cell.row, cell.col + 1}, {cell.row, cell.col - 1}}

// 		for _, neigh := range neighs {
// 			newPath := append([][]int{}, cell.path...)
// 			newPath = append(newPath, []int{neigh.row, neigh.col})

// 			queue = append(queue, Cell{neigh.row, neigh.col, newPath})
// 		}
// 	}

// 	return nil
// }

// func main() {
// 	// matrix := [][]int{
// 	// 	{0, 1, 0, 0, 0},
// 	// 	{0, 0, 0, 1, 0},
// 	// 	{1, 1, 0, 0, 0},
// 	// 	{0, 0, 0, 0, 0},
// 	// 	{1, 1, 0, 1, 0},
// 	// }

// 	matrix := [][]int{
// 		{0, 0, 1, 0, 0, 0, 0, 1, 0},
// 		{1, 0, 1, 0, 0, 1, 0, 1, 0},
// 		{0, 0, 1, 0, 0, 1, 0, 1, 0},
// 		{0, 1, 0, 0, 1, 1, 0, 1, 0},
// 		{0, 0, 0, 0, 1, 1, 0, 1, 0},
// 		{1, 1, 1, 0, 1, 1, 0, 1, 0},
// 		{0, 0, 0, 1, 0, 1, 0, 1, 0},
// 		{0, 0, 0, 0, 0, 1, 0, 1, 0},
// 		{0, 0, 0, 0, 0, 1, 0, 0, 0},
// 	}

// 	startRow, startCol := 0, 0
// 	endRow, endCol := 8, 8

// 	path := findShortestPath(matrix, startRow, startCol, endRow, endCol)
// 	if path != nil {
// 		fmt.Println("Shortest path found:")
// 		for _, coord := range path {
// 			fmt.Printf("(%d, %d) ", coord[0], coord[1])
// 		}
// 	} else {
// 		fmt.Println("No path found.")
// 	}
// }
