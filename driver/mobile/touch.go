package mobile

import "github.com/Anwert/fyne/v2"

// TouchEvent contains data relating to mobile touch events
type TouchEvent struct {
	fyne.PointEvent
}

// Touchable represents mobile touch events that can be sent to CanvasObjects
type Touchable interface {
	TouchDown(*TouchEvent)
	TouchUp(*TouchEvent)
	TouchCancel(*TouchEvent)
}
