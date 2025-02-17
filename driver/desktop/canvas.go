package desktop

import "github.com/Anwert/fyne/v2"

// Canvas defines the desktop specific extensions to a fyne.Canvas.
type Canvas interface {
	OnKeyDown() func(*fyne.KeyEvent)
	SetOnKeyDown(func(*fyne.KeyEvent))
	OnKeyUp() func(*fyne.KeyEvent)
	SetOnKeyUp(func(*fyne.KeyEvent))
}
