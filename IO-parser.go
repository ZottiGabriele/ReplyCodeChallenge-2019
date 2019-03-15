package main

import (
	"fmt"
	"os"
)

func parseInput(filePath string) Input {
	file, err := os.Open(filePath)
	check(err)
	//fmt.Println(file, err) //DELETE AFTER ARGUMENTS USAGE

	_mappa := Mappa{0, 0, 0, 0, [][]byte{}}

	_, err = fmt.Fscanf(file, "%d %d %d %d\n", &_mappa.widht, &_mappa.height, &_mappa.n_costumer, &_mappa.n_max_office)
	check(err)

	for i := 0; i < _mappa.height; i++ {
		row := make([]byte, _mappa.widht)
		_mappa.raw = append(_mappa.raw, row)

	}

	_costumers := make([]Costumer, _mappa.n_costumer)

	// read costumer
	for i := 0; i < _mappa.n_costumer; i = i + 1 {

		current := Costumer{}
		_, err = fmt.Fscanf(file, "%d %d %d\n", &current.x, &current.y, &current.points)
		check(err)

		_costumers[i] = current

		_mappa.raw[_costumers[i].y][_costumers[i].x] = 'C'
	}

	// read map
	for i := 0; i < _mappa.height; i++ {
		for j := 0; j < _mappa.widht; j++ {

			var temp byte
			if _mappa.raw[i][j] != 'C' {
				_, err = fmt.Fscanf(file, "%c", &_mappa.raw[i][j])
			} else {
				_, err = fmt.Fscanf(file, "%c", &temp)
			}
			check(err)
		}
		_, err = fmt.Fscanf(file, "\n")

	}

	// // debug map print
	// for i := 0; i < _mappa.height; i++ {
	// 	for j := 0; j < _mappa.widht; j++ {
	// 		fmt.Printf("%c", _mappa.raw[i][j])
	// 		check(err)
	// 	}
	// 	fmt.Println()
	// }
	//fmt.Println("______________________")
	//fmt.Println(_costumers)

	out := Input{_mappa, _costumers}
	return out
}

func writeOutput(paths []Path, outPath string) {
	out, err := os.Create(outPath)
	check(err)

	//out.WriteString(fmt.Sprintf("Hello World!")) //DELETE AFTER ARGUMENTS USAGE

	for i := 0; i < len(paths); i++ {
		out.WriteString(fmt.Sprintf("%d %d %s\n", paths[i].start_x, paths[i].start_y, &paths[i].directions))
	}

	// out.WriteString(fmt.Sprintf("%d\n", slideShow.n_of_slides))
	// for _, slide := range slideShow.slides {
	// 	for _, photo := range slide.photos {
	// 		out.WriteString(fmt.Sprintf("%d ", photo.ID))
	// 	}
	// 	out.WriteString(fmt.Sprintf("\n"))
	// }

	err = out.Close()
	check(err)

}
