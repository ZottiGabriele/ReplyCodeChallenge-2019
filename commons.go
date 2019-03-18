package main

import "fmt"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func CountingSort(arr []int) []int {

	// 1. Create an array(slice) of the size of the maximum value + 1
	k := GetMaxIntArray(arr)
	count := make([]int, k+1)

	// 2. Count each element
	for i := 0; i < len(arr); i++ {
		count[arr[i]] = count[arr[i]] + 1
	}
	// fmt.Println(count)

	// 3. Add up the elements
	for i := 1; i < k+1; i++ {
		count[i] = count[i] + count[i-1]
	}
	// fmt.Println(count)

	// 4. Put them back to result
	result := make([]int, len(arr)+1)
	for j := 0; j < len(arr); j++ {
		result[count[arr[j]]] = arr[j]
		count[arr[j]] = count[arr[j]] - 1
	}
	// fmt.Println(count)

	return result
}

// Return the maximum value in an integer array.
func GetMaxIntArray(arr []int) int {
	max_num := arr[0]
	for _, elem := range arr {
		if max_num < elem {
			max_num = elem
		}
	}
	return max_num
}

// CountIntArray counts the element frequencies.
func CountIntArray(arr []int) map[int]int {
	model := make(map[int]int)
	for _, elem := range arr {
		// first element is already initialized with 0
		model[elem] += 1
	}
	return model
}

func findPath(costumer Costumer, cell Cell) string {

	out := ""
	move := costumer

	// vedo posizione costumer rispetto all'ufficio e mi muovo
	if costumer.x < cell.office.x {

		for move.x <= cell.office.x {
			out = out + "R"
			(move.x)++
		}

		if costumer.y < cell.office.y {

			diff := cell.office.y - costumer.y
			for diff > 0 {
				out = out + "D"
				diff--
			}
		} else if costumer.y > cell.office.y {
			diff := cell.office.y - costumer.y
			for diff > 0 {
				out = out + "U"
				diff--
			}
		}

	} else if costumer.x > cell.office.x {

		for move.x <= cell.office.x {
			out = out + "L"
			(move.x)++
		}

		if costumer.y < cell.office.y {

			diff := cell.office.y - costumer.y
			for diff > 0 {
				out = out + "D"
				diff--
			}
		} else if costumer.y > cell.office.y {
			diff := cell.office.y - costumer.y
			for diff > 0 {
				out = out + "U"
				diff--
			}
		}

	} else {

		if costumer.y < cell.office.y {

			diff := cell.office.y - costumer.y
			for diff > 0 {
				out = out + "D"
				diff--
			}
		} else if costumer.y > cell.office.y {
			diff := cell.office.y - costumer.y
			for diff > 0 {
				out = out + "U"
				diff--
			}
		}
	}

	return out
}

func searchCellofCostumer(costumer Costumer, cell [][]Cell, width int, height int) (Costumer, Cell) {

	var posCol, posRow int
	var temp Cell

	posCol = (costumer.x / cell[0][0].width)
	if posCol > len(cell) {
		posCol--
	}
	posRow = (costumer.y / cell[0][0].height)
	if posRow > len(cell) {
		posRow--
	}

	if cell[posRow][posCol].office.x != -1 && cell[posRow][posCol].office.y != -1 {
		return costumer, cell[posRow][posCol]

	} else {

		if posRow == 0 && posCol == 0 {

			if cell[posRow+1][posCol].buildWeight < cell[posRow][posCol+1].buildWeight {
				temp = cell[posRow+1][posCol]
			} else {
				temp = cell[posRow][posCol+1]
			}
			if temp.buildWeight > cell[posRow+1][posCol+1].buildWeight {
				temp = cell[posRow+1][posCol+1]
			}

		} else if posRow == 0 && posCol == width-1 {

			if cell[posRow+1][posCol].buildWeight < cell[posRow][posCol-1].buildWeight {
				temp = cell[posRow+1][posCol]
			} else {
				temp = cell[posRow][posCol-1]
			}
			if temp.buildWeight > cell[posRow+1][posCol-1].buildWeight {
				temp = cell[posRow+1][posCol-1]
			}

		} else if posRow == height-1 && posCol == 0 {

			if cell[posRow-1][posCol].buildWeight < cell[posRow][posCol+1].buildWeight {
				temp = cell[posRow-1][posCol]
			} else {
				temp = cell[posRow][posCol+1]
			}
			if temp.buildWeight > cell[posRow-1][posCol+1].buildWeight {
				temp = cell[posRow-1][posCol+1]
			}

		}
	}

	return costumer, temp

}

func magic(input Input, cell [][]Cell) []Path {

	arrPath := make([]Path, len(input.costumers))

	width := input.mappa.widht
	height := input.mappa.height

	for i, costumer := range input.costumers {
		fmt.Println("Processing", i, costumer)
		var out Path
		out.start_x = costumer.x
		out.start_y = costumer.y

		//fmt.Println("Searching cell for ", costumer)
		mycostumer, my_cell := searchCellofCostumer(costumer, cell, width, height)
		out.directions = findPath(mycostumer, my_cell)

		arrPath[i] = out
	}

	return arrPath
}
