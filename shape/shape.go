// shape/shape.go
package shape

import (
	"math/rand"

	"github.com/dsmastrodomenico/tetris-go/board" // Importa el paquete board
)

// Shape representa una figura de Tetris.
type Shape struct {
	Type     int
	Rotation int
	Blocks   []board.Point // Ahora usa board.Point
}

// Definiciones de las formas de Tetris estándar (ejemplo para la forma 'I')
var Tetrominos = [][]board.Point{ // Usa board.Point aquí también
	// Forma I
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
	// Forma O
	{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
	// ... y así para las otras 5 figuras
}

// NewRandomShape crea una nueva figura aleatoria.
func NewRandomShape() *Shape {
	shapeType := rand.Intn(len(Tetrominos))
	return &Shape{
		Type:     shapeType,
		Rotation: 0,
		Blocks:   Tetrominos[shapeType],
	}
}

// Rotate rota la figura 90 grados en sentido horario.
func (s *Shape) Rotate() {
	// Implementación de la lógica de rotación
	s.Rotation = (s.Rotation + 1) % 4
	// (Aquí se actualizarían los 'Blocks' de la figura según la nueva rotación)
}

// GetAbsoluteBlocks retorna las coordenadas absolutas de los bloques
// de la figura en el tablero, dada su posición (x, y).
// Retorna []board.Point
func (s *Shape) GetAbsoluteBlocks(x, y int) []board.Point {
	absBlocks := make([]board.Point, len(s.Blocks))
	for i, p := range s.Blocks {
		absBlocks[i] = board.Point{X: x + p.X, Y: y + p.Y}
	}
	return absBlocks
}
