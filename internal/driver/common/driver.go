package common

import (
	"github.com/Anwert/fyne/v2"
	"github.com/Anwert/fyne/v2/internal/cache"
)

// CanvasForObject returns the canvas for the specified object.
func CanvasForObject(obj fyne.CanvasObject) fyne.Canvas {
	return cache.GetCanvasForObject(obj)
}
