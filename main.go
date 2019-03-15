package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

type ByBuildWeigth []Cell

func (cells ByBuildWeigth) Len() int { return len(cells) }
func (cells ByBuildWeigth) Swap(i, j int) {
	cells[i], cells[j] = cells[j], cells[i]
}
func (cells ByBuildWeigth) Less(i, j int) bool {
	return cells[i].buildWeight > cells[j].buildWeight
}

func main() {
	fmt.Println()
	inputPath, output := checkArgs()

	//parse input
	fmt.Println("----- Parsing input -----")
	input := parseInput(inputPath)

	//process input
	fmt.Println("\n----- Processing input -----")
	PointsTable := buildPointsTable()
	fmt.Println(PointsTable['#'])

	//dividere in celle
	width, height := input.mappa.widht, input.mappa.height
	var rapp int
	if width > height {
		rapp = width / height
	} else {
		rapp = height / width
	}

	count := 1
	div := 2
	for count < input.mappa.n_max_office {

		width, height = input.mappa.widht, input.mappa.height
		if width > height {
			width /= div
			height /= (div * rapp)
		} else {
			height /= div
			width /= (div * rapp)
		}
		count = div * div * rapp
		div++
	}
	div--

	cells := [][]Cell{}
	for i := 0; i < div; i++ {
		row := make([]Cell, div)
		cells = append(cells, row)
	}

	for i := 0; i < div; i++ {
		for j := 0; j < div; j++ {
			cells[i][j] = Cell{
				start_x: (input.mappa.widht / div) * j,
				start_y: (input.mappa.height / div) * i,
				office:  Office{-1, -1},
				world_x: j,
				world_y: i,
			}
			//compute width
			if j == div-1 {
				cells[i][j].width = input.mappa.widht/div + input.mappa.widht%div
			} else {
				cells[i][j].width = input.mappa.widht / div
			}

			if i == div-1 {
				cells[i][j].height = input.mappa.height/div + input.mappa.height%div
			} else {
				cells[i][j].height = input.mappa.height / div
			}
		}
	}

	//calcolare indice di bontà iniziale
	//calcolare indice di bontà path
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells); j++ {

			buildWeight := 0
			pathWeight := 0

			for v := 0; v < cells[i][j].height; v++ {
				for w := 0; w < cells[i][j].width; w++ {
					current := input.mappa.raw[v+i*(input.mappa.height/div)][w+j*(input.mappa.widht/div)]

					if current == 'C' {
						buildWeight++
					} else {
						pathWeight += PointsTable[current]
					}
				}
			}

			cells[i][j].buildWeight = buildWeight
			cells[i][j].pathWeight = pathWeight
		}
	}

	//generare x punti casuauli all'interno delle celle con indice di bontà minore

	//scegliere il migliore
	orderedCells := make([]Cell, div*div)
	k := 0
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells); j++ {
			orderedCells[k] = cells[i][j]
			k++
		}
	}

	sort.Sort(ByBuildWeigth(orderedCells))

	for i := 0; i < input.mappa.n_max_office && orderedCells[i].buildWeight != 0; i++ {
		x := rand.Intn(orderedCells[i].width) + orderedCells[i].start_x
		y := rand.Intn(orderedCells[i].height) + orderedCells[i].start_y

		cells[orderedCells[i].world_y][orderedCells[i].world_x].office = Office{
			x: x,
			y: y,
		}
	}

	for i := 0; i < len(cells); i++ {
		fmt.Println(cells[i])
	}

	//connetterli
	omg := magic(input, cells)
	//write output
	fmt.Println("\n----- Writing output -----")

	writeOutput(omg, output)

	fmt.Println()
}

func checkArgs() (string, string) {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("ERROR: wrong number of arguments")
		fmt.Println("Usage: arg1 arg2 (arg1 is input path, arg2 is output path)\n")
		os.Exit(1)
	}

	return args[0], args[1]
}

func buildPointsTable() map[byte]int {
	out := make(map[byte]int)
	out['~'] = 80
	out['*'] = 20
	out['+'] = 15
	out['X'] = 12
	out['_'] = 10
	out['H'] = 7
	out['T'] = 5

	out['C'] = 0
	out['#'] = 250
	return out
}
