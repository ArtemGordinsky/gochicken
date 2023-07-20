package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"log"
)

// Define constants for better readability.
const (
	screenWidth, screenHeight = 640, 480
	chickenSize               = 50
	jumpVelocity              = 4
)

// Game holds the game state.
type Game struct {
	chickenX, chickenY float64
	vy                 float64 // Vertical speed.
	chickenImage       *ebiten.Image
	directionRight     bool
}

// NewGame initializes a new game state.
func NewGame() *Game {
	g := &Game{}
	g.chickenX = 0
	g.chickenY = screenHeight - chickenSize // Place the chicken at the bottom.
	g.loadImages()
	return g
}

// Update progresses the game state one tick.
func (g *Game) Update() error {
	g.handleInput()
	g.applyPhysics()
	return nil
}

// handleInput deals with any user input.
func (g *Game) handleInput() {
	// Move the chicken to left.
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.chickenX -= 2
		if g.chickenX < -chickenSize {
			g.chickenX = screenWidth // Move to right edge if off the left edge
		}
		g.directionRight = false
	}
	// Move the chicken to right.
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.chickenX += 2
		if g.chickenX > screenWidth {
			g.chickenX = -chickenSize // Move to left edge if off the right edge
		}
		g.directionRight = true
	}
	// Make the chicken jump.
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.vy = -jumpVelocity
	}
}

// applyPhysics applies the game physics.
func (g *Game) applyPhysics() {
	// Apply gravity.
	g.vy += 0.2
	g.chickenY += g.vy
	// Chicken hits the ground.
	if g.chickenY > screenHeight-chickenSize {
		g.chickenY = screenHeight - chickenSize
		g.vy = 0
	}
}

// loadImages loads all the necessary images.
func (g *Game) loadImages() {
	if g.chickenImage == nil {
		chickenImage, _, err := ebitenutil.NewImageFromFile("assets/chicken.png")
		if err != nil {
			log.Fatalf("Failed to load chicken image: %v", err)
		}
		g.chickenImage = chickenImage
	}
}

// Draw renders the game state.
func (g *Game) Draw(screen *ebiten.Image) {
	// Fill the background with white color.
	screen.Fill(color.White)

	op := &ebiten.DrawImageOptions{}
	if g.directionRight {
		op.GeoM.Scale(-0.5, 0.5) // Flip the chicken image when it moves to right.
	} else {
		op.GeoM.Scale(0.5, 0.5) // Do not flip the chicken image when it moves to left.
	}
	op.GeoM.Translate(g.chickenX, g.chickenY)
	screen.DrawImage(g.chickenImage, op)
}

// Layout sets the game screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// main is the entry point of the game.
func main() {
	game := NewGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Chicken Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
