package dialog

import (
	"image/color"

	"github.com/Anwert/fyne/v2"
	"github.com/Anwert/fyne/v2/canvas"
	"github.com/Anwert/fyne/v2/container"
	"github.com/Anwert/fyne/v2/theme"
	"github.com/Anwert/fyne/v2/widget"
)

// ProgressInfiniteDialog is a simple dialog window that displays text and a infinite progress bar.
//
// Deprecated: Use NewCustomWithoutButtons() and add a widget.ProgressBarInfinite() inside.
type ProgressInfiniteDialog struct {
	*dialog

	bar *widget.ProgressBarInfinite
}

// NewProgressInfinite creates a infinite progress dialog and returns the handle.
// Using the returned type you should call Show().
//
// Deprecated: Use NewCustomWithoutButtons() and add a widget.ProgressBarInfinite() inside.
func NewProgressInfinite(title, message string, parent fyne.Window) *ProgressInfiniteDialog {
	d := newDialog(title, message, theme.InfoIcon(), nil /*cancel?*/, parent)
	bar := widget.NewProgressBarInfinite()
	rect := canvas.NewRectangle(color.Transparent)
	rect.SetMinSize(fyne.NewSize(200, 0))

	d.create(container.NewMax(rect, bar))
	return &ProgressInfiniteDialog{d, bar}
}

// Hide this dialog and stop the infinite progress goroutine
func (d *ProgressInfiniteDialog) Hide() {
	d.bar.Hide()
	d.dialog.Hide()
}
