package game

import "fmt"

type PathField = [][]int

func shortestPath(pathField PathField, cellFrom, cellTo Cell) []Cell {
	pathLen := pathField[cellTo.Row][cellTo.Row]
	result := make([]Cell, pathLen+1)
	result[0] = cellFrom
	result[pathLen] = cellTo
	cell := cellTo
	for i := pathLen - 1; i > 0; i-- {
		for _, nb := range allNeighbourCells(cell) {
			if pathField[nb.Row][nb.Col] == i {
				result[i] = nb
				cell = nb
				continue
			}
		}
	}
	return result
}

func fillPathField(field Field, cellFrom Cell) PathField {
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
			neighbours := neighbourCells(field, lcell)
			for _, c := range neighbours {
				if pathField[c.Row][c.Col] == -1 || pathField[c.Row][c.Col] > pathField[lcell.Row][lcell.Col]+1 {
					pathField[c.Row][c.Col] = pathField[lcell.Row][lcell.Col] + 1
					newLastCells = append(newLastCells, c)
				}
			}
		}
		lastCells = newLastCells
	}
	return pathField
}

func allNeighbourCells(cell Cell) []Cell {
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

func neighbourCells(field Field, cell Cell) []Cell {
	neighbours := make([]Cell, 0)

	appendNeighbour := func(cl Cell) {
		if checkBounds(cl) && field[cl.Row][cl.Col] == 0 {
			neighbours = append(neighbours, cl)
		}
	}

	appendNeighbour(Cell{cell.Row + 1, cell.Col})
	appendNeighbour(Cell{cell.Row - 1, cell.Col})
	appendNeighbour(Cell{cell.Row, cell.Col + 1})
	appendNeighbour(Cell{cell.Row, cell.Col - 1})

	return neighbours
}

func Run() {
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

	pathField := fillPathField(field, Cell{0, 0})

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
	fmt.Println("Okay111")
	path := shortestPath(pathField, Cell{0, 0}, Cell{8, 8})
	fmt.Println(path)
	fmt.Println(len(path))
}
