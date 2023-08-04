package game

import "fmt"

func ShortestPath(pathField PathField, cellFrom, cellTo Cell) []Cell {
	pathLen := pathField[cellTo.Row][cellTo.Row]
	result := make([]Cell, pathLen+1)
	result[0] = cellFrom
	result[pathLen] = cellTo
	cell := cellTo
	for i := pathLen - 1; i > 0; i-- {
		for _, nb := range neighbourCells(cell) {
			if pathField[nb.Row][nb.Col] == i {
				result[i] = nb
				cell = nb
				continue
			}
		}
	}
	return result
}

func FillPathField(field Field, cellFrom Cell) PathField {
	pathField := make([][]int, FIELD_SIZE)
	for i := 0; i < FIELD_SIZE; i++ {
		pathField[i] = make([]int, FIELD_SIZE)
	}

	for r, cols := range field {
		for c, _ := range cols {
			pathField[r][c] = -1
		}
	}

	pathField[cellFrom.Col][cellFrom.Row] = 0
	lastCells := []Cell{cellFrom}

	for len(lastCells) != 0 {
		newLastCells := make([]Cell, 0)
		for _, lcell := range lastCells {
			neighbours := neighbourCells(lcell)
			for _, c := range neighbours {
				if field[c.Row][c.Col] == 0 && (pathField[c.Row][c.Col] == -1 || pathField[c.Row][c.Col] > pathField[lcell.Row][lcell.Col]+1) {
					pathField[c.Row][c.Col] = pathField[lcell.Row][lcell.Col] + 1
					newLastCells = append(newLastCells, c)
				}
			}
		}
		lastCells = newLastCells
	}
	return pathField
}

func neighbourCells(cell Cell) []Cell {
	neighbours := make([]Cell, 0)

	appendNeighbour := func(cl Cell) {
		if checkBounds(cl) {
			neighbours = append(neighbours, cl)
		}
	}

	appendNeighbour(Cell{cell.Row + 1, cell.Col})
	appendNeighbour(Cell{cell.Row - 1, cell.Col})
	appendNeighbour(Cell{cell.Row, cell.Col + 1})
	appendNeighbour(Cell{cell.Row, cell.Col - 1})

	return neighbours
}

func TestPath() {
	field := [][]uint8{
		{0, 0, 1, 0, 0, 0, 0, 1, 0},
		{1, 0, 1, 0, 0, 1, 0, 1, 0},
		{0, 0, 1, 0, 0, 1, 0, 1, 0},
		{0, 1, 0, 0, 1, 1, 0, 1, 0},
		{0, 0, 0, 0, 1, 1, 0, 1, 0},
		{1, 1, 1, 0, 1, 1, 0, 1, 0},
		{0, 0, 0, 1, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0},
	}

	pathField := FillPathField(field, Cell{0, 0})

	for _, cols := range field {
		for _, val := range cols {
			fmt.Printf("%4d", val)
		}
		fmt.Println()
	}

	fmt.Println()

	for _, cols := range pathField {
		for _, val := range cols {
			fmt.Printf("%4d", val)
		}
		fmt.Println()
	}
	fmt.Println("Okay222")
	path := ShortestPath(pathField, Cell{0, 0}, Cell{8, 8})
	fmt.Println(path)
	fmt.Println(len(path))
}
