package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println()
	input, _ := checkArgs()

	//parse input
	fmt.Println("----- Parsing input -----")
	parseInput(input)

	//process input
	fmt.Println("\n----- Processing input -----")

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
