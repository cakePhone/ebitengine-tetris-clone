package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawTiles(screen *ebiten.Image) {
	for x := 0; x < len(tilemap); x++ {
		for y := 0; y < len(tilemap[x]); y++ {
			if tilemap[x][y] {
				op := ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
				screen.DrawImage(tile_sprite, &op)
			}
		}
	}

	for _, pos := range tetramino_in_game.tile_positions {
		x := tetramino_in_game.pos[0] + pos[0]
		y := tetramino_in_game.pos[1] + pos[1]

		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
		op.ColorScale.ScaleWithColor(tetramino_in_game.color)
		screen.DrawImage(tile_sprite, &op)
	}
}

func MoveTetraminoDown() {
	if Collide(&tetramino_in_game) {
		for _, pos := range tetramino_in_game.tile_positions {
			x := tetramino_in_game.pos[0] + pos[0]
			y := tetramino_in_game.pos[1] + pos[1]
			tilemap[x][y] = true
		}

		Tetramino{}.Get_Random_Tetramino().Add_To_Game(&tetramino_in_game)

		if Collide(&tetramino_in_game) {
			log.Println("Game Over")
			state = -1
		}

		return
	}

	tetramino_in_game.pos[1] += 1
	log.Println(tetramino_in_game.pos)
}

func Collide(tetr *Tetramino) bool {
	for _, pos := range tetr.tile_positions {
		x := tetr.pos[0] + pos[0]
		y := tetr.pos[1] + pos[1]
		if x >= 9 || y >= 19 || tilemap[x][y + 1] {
			return true
		}
	}
	return false
}
