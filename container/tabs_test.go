package container

import (
	"testing"

	"github.com/Anwert/fyne/v2/canvas"
	"github.com/Anwert/fyne/v2/internal/cache"
	"github.com/Anwert/fyne/v2/theme"
	"github.com/stretchr/testify/assert"
)

func TestTabButton_Icon_Change(t *testing.T) {
	b := &tabButton{icon: theme.CancelIcon()}
	r := cache.Renderer(b)
	icon := r.Objects()[3].(*canvas.Image)
	oldResource := icon.Resource

	b.icon = theme.ConfirmIcon()
	b.Refresh()
	assert.NotEqual(t, oldResource, icon.Resource)
}
