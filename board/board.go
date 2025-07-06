// board/board.go
package board

const (
	Width  = 10
	Height = 20
)

// Point representa un punto (coordenada x, y) en el tablero.
type Point struct {
	X int
	Y int
}

// Board representa el tablero de juego.
// Un valor de 0 puede significar celda vacía, y otros valores el tipo de figura.
type Board [Height][Width]int

// NewBoard crea e inicializa un nuevo tablero de juego.
func NewBoard() *Board {
	b := &Board{}
	// Aquí podrías inicializar todas las celdas a 0 (vacías)
	return b
}

// ClearLines verifica y limpia las líneas completas.
// Retorna el número de líneas limpiadas.
func (b *Board) ClearLines() int {
	// Implementación para limpiar líneas
	return 0
}

// CheckCollision verifica si una figura colisiona con el tablero o con otras figuras.
// Ahora espera []board.Point
func (b *Board) CheckCollision(shapeBlocks []Point, x, y int) bool {
	// Implementación para verificar colisiones
	// Por ahora, una implementación simple que siempre dice que no hay colisión
	// para que el código compile y puedas probar.
	// Más adelante, aquí iría la lógica real.
	_ = x // Usar x para evitar el error de variable no usada
	_ = y // Usar y para evitar el error de variable no usada

	for _, p := range shapeBlocks {
		// Ejemplo de verificación de límites del tablero.
		// La lógica completa de colisión debe implementarse aquí.
		if p.Y < 0 || p.Y >= Height || p.X < 0 || p.X >= Width {
			return true // Colisión con bordes del tablero
		}
		// Si la celda en el tablero ya está ocupada por otra figura
		if b[p.Y][p.X] != 0 {
			return true // Colisión con una figura existente
		}
	}
	return false
}
