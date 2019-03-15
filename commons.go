package main

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
			out := append(out, 'R')
			(move.x)++
		}

		if costumer.y < cell.office.y {

			diff := cell.office.y - costumer.y
			for diff > 0 {
				out := append(out, 'D')
			}
		} else if costumer.y > cell.office.y {
			diff := cell.office.y - costumer.y
			for diff > 0 {
				out := append(out, 'U')
			}
		}

	} else if costumer.x > cell.office.x {

		for move.x <= cell.office.x {
			out := append(out, 'L')
			(move.x)++
		}

		if costumer.y < cell.office.y {

			diff := cell.office.y - costumer.y
			for diff > 0 {
				out := append(out, 'D')
			}
		} else if costumer.y > cell.office.y {
			diff := cell.office.y - costumer.y
			for diff > 0 {
				out := append(out, 'U')
			}
		}

	} else {

		if costumer.y < cell.office.y {

			diff := cell.office.y - costumer.y
			for diff > 0 {
				out := append(out, 'D')
			}
		} else if costumer.y > cell.office.y {
			diff := cell.office.y - costumer.y
			for diff > 0 {
				out := append(out, 'U')
			}
		}
	}
}

func searchCell_ofCostumer(costumer Costumer, cell [][]Cell) (Costumer, Cell) {

	posCol := floor(costumer.x / cell.width)
	posRow := floor(costumer.y / cell.height)

	return costumer, cell[posRow][posCol]
}

func magic(Input input, cell [][]Cell) Path {

	var arrPath []Path

	for i, costumer := range input.costumers {

		var out Path
		out.start_x = costumer.x
		out.start_y = costumer.y

		my_cell := searchCell_ofCostumer(costumer, cell)
		out.directions = findPath(costumer, my_cell)

	}
}
