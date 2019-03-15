package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println()
	inputPath, _ := checkArgs()

	//parse input
	fmt.Println("----- Parsing input -----")
	input := parseInput(inputPath)

	//process input
	fmt.Println("\n----- Processing input -----")
	PointsTable := buildPointsTable()

	//dividere in celle
	width, height := input.mappa.widht, input.mappa.height
	count := 1
	for ; count < input.mappa.n_max_office; count *= 2 {
		if(width > height)
	}

	//calcolare indice di bontà iniziale
	//calcolare indice di bontà path

	//generare x punti casuauli all'interno delle celle con indice di bontà minore
	//scegliere il migliore

	//connetterli

	//write output
	fmt.Println("\n----- Writing output -----")
	//writeOutput(output)

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

func buildPointsTable() map[byte]int32 {
	out := make(map[byte]int32)
	out['~'] = 80
	out['*'] = 20
	out['+'] = 15
	out['X'] = 12
	out['_'] = 10
	out['H'] = 7
	out['T'] = 5

	out['C'] = 0
	out['#'] = 200
	return out
}
