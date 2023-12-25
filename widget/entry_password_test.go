package widget_test

import (
	"testing"

	"github.com/Anwert/fyne/v2"
	"github.com/Anwert/fyne/v2/canvas"
	"github.com/Anwert/fyne/v2/container"
	"github.com/Anwert/fyne/v2/test"
	"github.com/Anwert/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

func TestNewPasswordEntry(t *testing.T) {
	p := widget.NewPasswordEntry()
	p.Text = "visible"
	r := test.WidgetRenderer(p)

	cont := r.Objects()[2].(*container.Scroll).Content.(fyne.Widget)
	r = test.WidgetRenderer(cont)
	rich := r.Objects()[1].(*widget.RichText)
	r = test.WidgetRenderer(rich)

	assert.Equal(t, "•••••••", r.Objects()[0].(*canvas.Text).Text)
}
