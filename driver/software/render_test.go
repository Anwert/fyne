package software

import (
	"image/color"
	"testing"

	"github.com/Anwert/fyne/v2"
	"github.com/Anwert/fyne/v2/canvas"
	"github.com/Anwert/fyne/v2/container"
	"github.com/Anwert/fyne/v2/test"
	"github.com/Anwert/fyne/v2/theme"
	"github.com/Anwert/fyne/v2/widget"
)

func TestRender(t *testing.T) {
	obj := widget.NewLabel("Hi")
	test.AssertImageMatches(t, "label_dark.png", Render(obj, theme.DarkTheme()))
	test.AssertImageMatches(t, "label_light.png", Render(obj, theme.LightTheme()))
}

func TestRender_State(t *testing.T) {
	obj := widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {})
	test.AssertImageMatches(t, "button.png", Render(obj, theme.DarkTheme()))

	obj.Importance = widget.HighImportance
	obj.Refresh()
	test.AssertImageMatches(t, "button_important.png", Render(obj, theme.DarkTheme()))
}

func TestRender_Focus(t *testing.T) {
	obj := widget.NewEntry()
	test.AssertImageMatches(t, "entry.png", Render(obj, theme.DarkTheme()))

	obj.FocusGained()
	test.AssertImageMatches(t, "entry_focus.png", Render(obj, theme.DarkTheme()))
}

func TestRenderCanvas(t *testing.T) {
	obj := container.NewAppTabs(
		container.NewTabItem("Tab 1", container.NewVBox(
			widget.NewLabel("Label"),
			widget.NewButton("Button", func() {}),
		)))

	c := NewCanvas()
	c.SetContent(obj)

	if fyne.CurrentDevice().IsMobile() {
		test.AssertImageMatches(t, "canvas_mobile.png", RenderCanvas(c, theme.LightTheme()))
	} else {
		test.AssertImageMatches(t, "canvas.png", RenderCanvas(c, theme.LightTheme()))
	}
}

func TestRender_ImageSize(t *testing.T) {
	image := canvas.NewImageFromFile("../../theme/icons/fyne.png")
	image.FillMode = canvas.ImageFillOriginal
	bg := canvas.NewCircle(color.NRGBA{255, 0, 0, 128})
	bg.StrokeColor = color.White
	bg.StrokeWidth = 5

	c := container.NewStack(image, bg)

	test.AssertImageMatches(t, "image_size.png", Render(c, theme.LightTheme()))
}
