package main

import (
	"fmt"
	"os"
)

func parseInput(filePath string) /*return type if needed*/ {
	file, err := os.Open(filePath)
	check(err)
	fmt.Println(file, err) //DELETE AFTER ARGUMENTS USAGE

	N := 0
	_, err = fmt.Fscanf(file, "%d", &N)
	check(err)

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
