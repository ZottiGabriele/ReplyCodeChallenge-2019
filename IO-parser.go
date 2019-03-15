package main

import (
	"fmt"
	"os"
)

func parseInput(filePath string) /*return type if needed*/ {
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

	// per ogni riga ho un customer
	for i := 0; i < _mappa.n_costumer; i = i + 1 {

		current := Costumer{}
		_, err = fmt.Fscanf(file, "%d %d %d\n", &current.x, &current.y, &current.points)
		_costumers[i] = current

		_mappa.raw[_costumers[i].y][_costumers[i].x] = 'C'
	}

	for i := 0; i < _mappa.widht; i++ {
		for j := 0; j < _mappa.height; j++ {

		}
	}
	/*
		for i := 0; i < _mappa.height; i++ {
			for j := 0; j < _mappa.widht; j++ {
				fmt.Print(_mappa.raw[i][j])
			}
			fmt.Println()
		}
	*/

	fmt.Println(_costumers)
	// out := PhotoCollection{n_of_photos: N}
	// out.photos = make([]Photo, N)

	// for i := 0; i < N; i++ {
	// 	current := Photo{ID: i}

	// 	_, err = fmt.Fscanf(file, "%c", &current.orient)
	// 	check(err)

	// 	_, err = fmt.Fscanf(file, "%d", &current.n_of_tags)
	// 	check(err)

	// 	for j := 0; j < current.n_of_tags; j++ {
	// 		currentTag := ""
	// 		_, err = fmt.Fscanf(file, "%s", &currentTag)
	// 		check(err)
	// 		current.tags = append(current.tags, currentTag)
	// 	}
	// 	out.photos[i] = current
	// }

	// return out
}

func writeOutput( /*OutputType, */ outPath string) {
	out, err := os.Create(outPath)
	check(err)

	out.WriteString(fmt.Sprintf("Hello World!")) //DELETE AFTER ARGUMENTS USAGE

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
