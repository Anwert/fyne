package dialog

import (
	"image/color"

	"github.com/Anwert/fyne/v2"
	"github.com/Anwert/fyne/v2/canvas"
	"github.com/Anwert/fyne/v2/container"
	"github.com/Anwert/fyne/v2/theme"
	"github.com/Anwert/fyne/v2/widget"
)

// ProgressDialog is a simple dialog window that displays text and a progress bar.
//
// Deprecated: Use NewCustomWithoutButtons() and add a widget.ProgressBar() inside.
type ProgressDialog struct {
	*dialog

	bar *widget.ProgressBar
}

// SetValue updates the value of the progress bar - this should be between 0.0 and 1.0.
func (p *ProgressDialog) SetValue(v float64) {
	p.bar.SetValue(v)
}

// NewProgress creates a progress dialog and returns the handle.
// Using the returned type you should call Show() and then set its value through SetValue().
//
// Deprecated: Use NewCustomWithoutButtons() and add a widget.ProgressBar() inside.
func NewProgress(title, message string, parent fyne.Window) *ProgressDialog {
	d := newDialog(title, message, theme.InfoIcon(), nil /*cancel?*/, parent)
	bar := widget.NewProgressBar()
	rect := canvas.NewRectangle(color.Transparent)
	rect.SetMinSize(fyne.NewSize(200, 0))

	d.create(container.NewMax(rect, bar))
	return &ProgressDialog{d, bar}
}
