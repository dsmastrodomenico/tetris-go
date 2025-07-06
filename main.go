// main.go
package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2" // Biblioteca para la interacción con la terminal

	"github.com/dsmastrodomenico/tetris-go/game" // Asegúrate de que "tetris-go" sea el nombre de tu módulo
	"github.com/dsmastrodomenico/tetris-go/input"
)

func main() {
	// Inicializa la pantalla de tcell para la interacción con la terminal
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err = s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	defer s.Fini() // Asegura que la pantalla se limpie al salir

	// Configura la salida del log a un archivo o a stdout para depuración
	// Por ahora, lo dirigimos a os.Stderr para no interferir con la pantalla del juego
	log.SetOutput(os.Stderr) // Puedes cambiarlo a un archivo de log si prefieres: os.OpenFile("game.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	// Canal para eventos de teclado
	inputEvents := make(chan input.KeyEvent)
	go input.StartKeyListener(inputEvents) // Inicia la goroutine para escuchar el teclado

	// Crea una nueva instancia del juego
	g := game.NewGame(s, inputEvents)

	// Inicia el bucle principal del juego
	g.StartGame()

	// Por ahora, el juego terminará cuando la goroutine de StartGame decida.
	// Más adelante, podrías tener un canal para señales de salida del juego.
}
