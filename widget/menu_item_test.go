package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Anwert/fyne/v2"
	"github.com/Anwert/fyne/v2/internal/cache"
	"github.com/Anwert/fyne/v2/theme"
)

func TestMenuItem_Disabled(t *testing.T) {
	i := fyne.NewMenuItem("Disabled", func() {})
	m := fyne.NewMenu("top", []*fyne.MenuItem{i}...)
	i.Disabled = true
	w := newMenuItem(i, NewMenu(m))
	r := cache.Renderer(w)

	assert.Equal(t, theme.DisabledColor(), r.(*menuItemRenderer).text.Color)
}
