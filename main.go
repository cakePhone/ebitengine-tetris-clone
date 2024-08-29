package main

import (
	_ "image"
	"log"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game implements ebiten.Game interface.
type Game struct{}

var state int = 0

const (
	tileSize = 32
)

const (
	screenHeight = 20 * tileSize
	screenWidth  = 10 * tileSize
)

var tetramino_in_game Tetramino = Tetramino{}.Get_Random_Tetramino()

var tilemap [screenWidth / tileSize][screenHeight / tileSize]bool

var ticks_elapsed int

var tile_sprite *ebiten.Image

func init() {
	var err error
	tile_sprite, _, err = ebitenutil.NewImageFromFile("assets/tile.png")

	Tetramino{}.Get_Random_Tetramino().Add_To_Game(&tetramino_in_game)

	if err != nil {
		log.Fatal(err)
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		tetramino_in_game.pos[0] += 1
	}

	// Game Over State
	if state == -1 {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			state = 0
			tilemap = [screenWidth / tileSize][screenHeight / tileSize]bool{}
			Tetramino{}.Get_Random_Tetramino().Add_To_Game(&tetramino_in_game)
		}

		return nil
	}

	ticks_elapsed += 1

	if ticks_elapsed%20 == 0 {
		MoveTetraminoDown()

		ticks_elapsed = 0
	}


	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	screen.Fill(color.RGBA{0x18, 0x18, 0x18, 0xff})

	DrawTiles(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
