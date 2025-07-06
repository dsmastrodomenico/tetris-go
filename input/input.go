// input/input.go
package input

import (
	"log"

	"github.com/gdamore/tcell/v2" // Una buena opción para manejo de terminal
)

// KeyEvent representa un evento de teclado.
type KeyEvent struct {
	Key  tcell.Key
	Rune rune
}

// StartKeyListener inicia un goroutine para escuchar eventos de teclado.
// Envía los eventos al canal 'events'.
func StartKeyListener(events chan<- KeyEvent) {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err = s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				events <- KeyEvent{Key: ev.Key(), Rune: ev.Rune()}
			}
		}
	}()
}
