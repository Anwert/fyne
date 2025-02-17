package canvas

import "github.com/Anwert/fyne/v2"

// Refresh instructs the containing canvas to refresh the specified obj.
func Refresh(obj fyne.CanvasObject) {
	if fyne.CurrentApp() == nil || fyne.CurrentApp().Driver() == nil {
		return
	}

	c := fyne.CurrentApp().Driver().CanvasForObject(obj)
	if c != nil {
		c.Refresh(obj)
	}
}

// repaint instructs the containing canvas to redraw, even if nothing changed.
func repaint(obj fyne.CanvasObject) {
	if fyne.CurrentApp() == nil || fyne.CurrentApp().Driver() == nil {
		return
	}

	c := fyne.CurrentApp().Driver().CanvasForObject(obj)
	if c != nil {
		if paint, ok := c.(interface{ SetDirty() }); ok {
			paint.SetDirty()
		}
	}
}
