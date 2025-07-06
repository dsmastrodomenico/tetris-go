// game/game.go
package game

import (
	"time"

	"github.com/gdamore/tcell/v2"

	"github.com/dsmastrodomenico/tetris-go/board"
	"github.com/dsmastrodomenico/tetris-go/input"
	"github.com/dsmastrodomenico/tetris-go/shape"
)

// Game representa el estado actual del juego.
type Game struct {
	Board          *board.Board
	CurrentShape   *shape.Shape
	ShapeX, ShapeY int // Posición de la figura actual
	Screen         tcell.Screen
	InputEvents    chan input.KeyEvent
}

// NewGame inicializa un nuevo juego.
func NewGame(s tcell.Screen, events chan input.KeyEvent) *Game {
	return &Game{
		Board:       board.NewBoard(),
		Screen:      s,
		InputEvents: events,
	}
}

// StartGame inicia el bucle principal del juego.
func (g *Game) StartGame() {
	// Inicializar la primera figura
	g.CurrentShape = shape.NewRandomShape()
	g.ShapeX = board.Width/2 - 2 // Posición inicial
	g.ShapeY = 0

	ticker := time.NewTicker(500 * time.Millisecond) // Caída cada 500ms
	defer ticker.Stop()

	quit := make(chan struct{})

	go func() {
		for {
			select {
			case ev := <-g.InputEvents:
				g.handleInput(ev)
			case <-ticker.C:
				g.moveShapeDown()
			case <-quit:
				return
			}
			g.draw() // Redibujar después de cada evento o caída
		}
	}()

	// Mantener el juego corriendo hasta que se reciba una señal de salida
	<-quit // Esto es una simplificación, necesitarás un mecanismo para salir del juego.
}

// handleInput procesa la entrada del teclado.
func (g *Game) handleInput(ev input.KeyEvent) {
	newX, newY := g.ShapeX, g.ShapeY
	rotatedShape := g.CurrentShape

	switch ev.Key {
	case tcell.KeyLeft:
		newX--
	case tcell.KeyRight:
		newX++
	case tcell.KeyDown:
		newY++
	case tcell.KeyEnter, tcell.KeyUp: // Rotar
		tempShape := *g.CurrentShape // Copia para probar rotación
		tempShape.Rotate()
		rotatedShape = &tempShape
	case tcell.KeyEscape: // Salir del juego
		g.Screen.Fini()
		// Aquí deberías enviar una señal al canal `quit` para terminar el bucle principal
		return
	}

	// Verificar colisiones antes de mover/rotar
	if !g.Board.CheckCollision(rotatedShape.GetAbsoluteBlocks(newX, newY), newX, newY) {
		g.ShapeX = newX
		g.ShapeY = newY
		g.CurrentShape = rotatedShape // Actualizar si la rotación es válida
	} else if ev.Key == tcell.KeyDown {
		g.lockShape() // La figura colisionó hacia abajo, bloquearla
		g.Board.ClearLines()
		g.CurrentShape = shape.NewRandomShape() // Nueva figura
		g.ShapeX = board.Width/2 - 2
		g.ShapeY = 0
		// Aquí deberías verificar si la nueva figura colisiona inmediatamente (Game Over)
	}
}

// moveShapeDown intenta mover la figura hacia abajo.
func (g *Game) moveShapeDown() {
	if !g.Board.CheckCollision(g.CurrentShape.GetAbsoluteBlocks(g.ShapeX, g.ShapeY+1), g.ShapeX, g.ShapeY+1) {
		g.ShapeY++
	} else {
		g.lockShape()
		g.Board.ClearLines()
		g.CurrentShape = shape.NewRandomShape()
		g.ShapeX = board.Width/2 - 2
		g.ShapeY = 0
		// Verificar Game Over
	}
}

// lockShape bloquea la figura actual en el tablero.
func (g *Game) lockShape() {
	blocks := g.CurrentShape.GetAbsoluteBlocks(g.ShapeX, g.ShapeY)
	for _, p := range blocks {
		if p.X >= 0 && p.X < board.Width && p.Y >= 0 && p.Y < board.Height {
			g.Board[p.Y][p.X] = g.CurrentShape.Type + 1 // Usar 1-7 para tipos de figuras
		}
	}
}

// draw renderiza el tablero y la figura en la terminal.
func (g *Game) draw() {
	g.Screen.Clear()

	// Dibujar el tablero
	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			if g.Board[y][x] != 0 {
				g.Screen.SetContent(x, y, '█', nil, tcell.StyleDefault.Background(tcell.ColorBlue))
			} else {
				g.Screen.SetContent(x, y, ' ', nil, tcell.StyleDefault)
			}
		}
	}

	// Dibujar la figura actual
	blocks := g.CurrentShape.GetAbsoluteBlocks(g.ShapeX, g.ShapeY)
	for _, p := range blocks {
		if p.X >= 0 && p.X < board.Width && p.Y >= 0 && p.Y < board.Height {
			g.Screen.SetContent(p.X, p.Y, '█', nil, tcell.StyleDefault.Background(tcell.ColorGreen))
		}
	}
	g.Screen.Show()
}
