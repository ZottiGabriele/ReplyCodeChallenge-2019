package main

type Mappa struct {
	widht        int // larghezza
	height       int // altezza
	n_costumer   int // # costumer
	n_max_office int // # max office

	raw [][]byte
}

type Costumer struct {
	x      int // colonne
	y      int // righe
	points int
}

type Path struct {
	start_x    int
	start_y    int
	directions string
	points     int
}

type Input struct {
	mappa     Mappa
	costumers []Costumer
}
