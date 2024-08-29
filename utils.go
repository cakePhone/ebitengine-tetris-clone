package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawTiles(screen *ebiten.Image) {
	for _, tetr := range tetraminos_in_game {
		for _, pos := range tetr.tile_positions {
			x := tetr.pos[0] + pos[0]
			y := tetr.pos[1] + pos[1]
			
			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
			op.ColorScale.ScaleWithColor(tetr.color)
			screen.DrawImage(tile_sprite, &op)
		}
	}
}

func MoveTetraminosDown() {
	for i := range tetraminos_in_game {
		if tetraminos_in_game[i].placed {
			continue
		}

		lowestTile := FindLowestTile(&tetraminos_in_game[i])

		if lowestTile[1]+1 >= screenHeight/tileSize {
			tetraminos_in_game[i].placed = true

			Tetramino{}.Get_Random_Tetramino().Add_To_Game(&tetraminos_in_game)

			return
		}

		tetraminos_in_game[i].pos[1] += 1
		log.Println(tetraminos_in_game[i].pos)
	}
}

func FindLowestTile(tetr *Tetramino) [2]int {
	lowestTile := tetr.pos[1]
	for _, pos := range tetr.tile_positions {
		y := tetr.pos[1] + pos[1]
		if y > lowestTile {
			lowestTile = y
		}
	}
	return [2]int{tetr.pos[0], lowestTile}
}
