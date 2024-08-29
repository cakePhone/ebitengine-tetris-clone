package main

import (
	"image/color"
	"math/rand"

	"golang.org/x/image/colornames"
)

type Tetramino struct {
	pos            [2]int
	color          color.RGBA
	tile_positions [4][2]int
}

func (tetr Tetramino) Add_To_Game(tetraminos_in_game *Tetramino) {
	tetramino_in_game = tetr.Get_Random_Tetramino()
}

func (tetr Tetramino) Get_Random_Tetramino() Tetramino {
	switch rand.Intn(7) {
	case 0:
		return tetr.New_I_Tetramino()
	case 1:
		return tetr.New_J_Tetramino()
	case 2:
		return tetr.New_L_Tetramino()
	case 3:
		return tetr.New_O_Tetramino()
	case 4:
		return tetr.New_S_Tetramino()
	case 5:
		return tetr.New_T_Tetramino()
	case 6:
		return tetr.New_Z_Tetramino()
	default:
		return tetr
	}
}

func (tetr Tetramino) New_L_Tetramino() Tetramino {
	return Tetramino{
		pos:   [2]int{4, 0},
		color: colornames.Red,
		tile_positions: [4][2]int{
			{0, 1},
			{0, 2},
			{0, 0},
			{1, 2},
		},
	}
}

func (tetr Tetramino) New_J_Tetramino() Tetramino {
	return Tetramino{
		pos:   [2]int{4, 0},
		color: colornames.Blue,
		tile_positions: [4][2]int{
			{0, 1},
			{0, 2},
			{0, 0},
			{1, 0},
		},
	}
}

func (tetr Tetramino) New_S_Tetramino() Tetramino {
	return Tetramino{
		pos:   [2]int{4, 0},
		color: colornames.Green,
		tile_positions: [4][2]int{
			{0, 1},
			{0, 2},
			{1, 1},
			{1, 0},
		},
	}
}

func (tetr Tetramino) New_T_Tetramino() Tetramino {
	return Tetramino{
		pos:   [2]int{4, 0},
		color: colornames.Yellow,
		tile_positions: [4][2]int{
			{0, 1},
			{0, 2},
			{1, 1},
			{0, 0},
		},
	}
}

func (tetr Tetramino) New_O_Tetramino() Tetramino {
	return Tetramino{
		pos:   [2]int{4, 0},
		color: colornames.Purple,
		tile_positions: [4][2]int{
			{0, 1},
			{0, 2},
			{1, 1},
			{1, 2},
		},
	}
}

func (tetr Tetramino) New_I_Tetramino() Tetramino {
	return Tetramino{
		pos:   [2]int{4, 0},
		color: colornames.Cyan,
		tile_positions: [4][2]int{
			{0, 1},
			{0, 2},
			{0, 3},
			{0, 0},
		},
	}
}

func (tetr Tetramino) New_Z_Tetramino() Tetramino {
	return Tetramino{
		pos:   [2]int{4, 0},
		color: colornames.Orange,
		tile_positions: [4][2]int{
			{0, 1},
			{0, 2},
			{1, 1},
			{1, 0},
		},
	}
}