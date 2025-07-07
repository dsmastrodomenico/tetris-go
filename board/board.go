// board/board.go
package board

const (
	Width  = 10 // Ancho del tablero de Tetris
	Height = 20 // Alto del tablero de Tetris
)

// Point representa un punto (coordenada x, y) en el tablero.
type Point struct {
	X int
	Y int
}

// Board representa el tablero de juego.
// Un valor de 0 significa celda vacía; otros valores indican el tipo de figura.
type Board [Height][Width]int

// NewBoard crea e inicializa un nuevo tablero de juego.
func NewBoard() *Board {
	b := &Board{}
	// Todas las celdas se inicializan a 0 (vacías) por defecto en Go.
	return b
}

// ClearLines verifica y limpia las líneas completas.
// Retorna el número de líneas limpiadas.
func (b *Board) ClearLines() int {
	linesCleared := 0
	// Recorrer el tablero de abajo hacia arriba
	for y := Height - 1; y >= 0; y-- {
		isFull := true
		for x := 0; x < Width; x++ {
			if b[y][x] == 0 {
				isFull = false
				break
			}
		}

		if isFull {
			linesCleared++
			// Mover todas las líneas de arriba una posición hacia abajo
			for moveY := y; moveY > 0; moveY-- {
				for x := 0; x < Width; x++ {
					b[moveY][x] = b[moveY-1][x]
				}
			}
			// Limpiar la fila superior (ahora es nueva)
			for x := 0; x < Width; x++ {
				b[0][x] = 0
			}
			y++ // Volver a verificar la misma línea, ya que ha sido rellenada desde arriba
		}
	}
	return linesCleared
}

// CheckCollision verifica si una figura, en una posición dada (offsetX, offsetY),
// colisionaría con los límites del tablero o con bloques ya existentes.
func (b *Board) CheckCollision(shapeBlocks []Point, offsetX, offsetY int) bool {
	for _, p := range shapeBlocks {
		// Calcular la posición absoluta de cada bloque de la figura
		absX := p.X + offsetX
		absY := p.Y + offsetY

		// 1. Colisión con los límites laterales o inferiores del tablero
		if absX < 0 || absX >= Width || absY >= Height {
			return true // Fuera de los límites del tablero
		}

		// 2. Colisión con la parte superior del tablero (permite aparecer)
		// Si absY es menor que 0 (la figura está parcialmente fuera por arriba),
		// no es una colisión válida de movimiento, solo de aparición.
		// Solo verificamos colisiones con bloques existentes si absY está dentro del tablero.
		if absY >= 0 {
			if b[absY][absX] != 0 {
				return true // Colisión con un bloque existente en el tablero
			}
		}
	}
	return false
}
