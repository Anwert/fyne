package container

import (
	"image/color"
	"testing"

	"github.com/Anwert/fyne/v2"
	"github.com/Anwert/fyne/v2/canvas"
	"github.com/Anwert/fyne/v2/driver/desktop"
	"github.com/Anwert/fyne/v2/test"
	"github.com/stretchr/testify/assert"
)

func TestSplitContainer_MinSize(t *testing.T) {
	rectA := canvas.NewRectangle(color.Black)
	rectA.SetMinSize(fyne.NewSize(10, 10))
	rectB := canvas.NewRectangle(color.Black)
	rectB.SetMinSize(fyne.NewSize(10, 10))
	t.Run("Horizontal", func(t *testing.T) {
		min := NewHSplit(rectA, rectB).MinSize()
		assert.Equal(t, rectA.MinSize().Width+rectB.MinSize().Width+dividerThickness(), min.Width)
		assert.Equal(t, fyne.Max(rectA.MinSize().Height, fyne.Max(rectB.MinSize().Height, dividerLength())), min.Height)
	})
	t.Run("Vertical", func(t *testing.T) {
		min := NewVSplit(rectA, rectB).MinSize()
		assert.Equal(t, fyne.Max(rectA.MinSize().Width, fyne.Max(rectB.MinSize().Width, dividerLength())), min.Width)
		assert.Equal(t, rectA.MinSize().Height+rectB.MinSize().Height+dividerThickness(), min.Height)
	})
}

func TestSplitContainer_Resize(t *testing.T) {
	for name, tt := range map[string]struct {
		horizontal       bool
		size             fyne.Size
		wantLeadingPos   fyne.Position
		wantLeadingSize  fyne.Size
		wantTrailingPos  fyne.Position
		wantTrailingSize fyne.Size
	}{
		"horizontal": {
			true,
			fyne.NewSize(100, 100),
			fyne.NewPos(0, 0),
			fyne.NewSize(50-dividerThickness()/2, 100),
			fyne.NewPos(50+dividerThickness()/2, 0),
			fyne.NewSize(50-dividerThickness()/2, 100),
		},
		"vertical": {
			false,
			fyne.NewSize(100, 100),
			fyne.NewPos(0, 0),
			fyne.NewSize(100, 50-dividerThickness()/2),
			fyne.NewPos(0, 50+dividerThickness()/2),
			fyne.NewSize(100, 50-dividerThickness()/2),
		},
		"horizontal insufficient width": {
			true,
			fyne.NewSize(20, 100),
			fyne.NewPos(0, 0),
			// minSize of leading is 1/3 of minSize of trailing
			fyne.NewSize((20-dividerThickness())/4, 100),
			fyne.NewPos((20-dividerThickness())/4+dividerThickness(), 0),
			fyne.NewSize((20-dividerThickness())*3/4, 100),
		},
		"vertical insufficient height": {
			false,
			fyne.NewSize(100, 20),
			fyne.NewPos(0, 0),
			// minSize of leading is 1/3 of minSize of trailing
			fyne.NewSize(100, (20-dividerThickness())/4),
			fyne.NewPos(0, (20-dividerThickness())/4+dividerThickness()),
			fyne.NewSize(100, (20-dividerThickness())*3/4),
		},
		"horizontal zero width": {
			true,
			fyne.NewSize(0, 100),
			fyne.NewPos(0, 0),
			fyne.NewSize(0, 100),
			fyne.NewPos(dividerThickness(), 0),
			fyne.NewSize(0, 100),
		},
		"horizontal zero height": {
			true,
			fyne.NewSize(100, 0),
			fyne.NewPos(0, 0),
			fyne.NewSize(50-dividerThickness()/2, 0),
			fyne.NewPos(50+dividerThickness()/2, 0),
			fyne.NewSize(50-dividerThickness()/2, 0),
		},
		"vertical zero width": {
			false,
			fyne.NewSize(0, 100),
			fyne.NewPos(0, 0),
			fyne.NewSize(0, 50-dividerThickness()/2),
			fyne.NewPos(0, 50+dividerThickness()/2),
			fyne.NewSize(0, 50-dividerThickness()/2),
		},
		"vertical zero height": {
			false,
			fyne.NewSize(100, 0),
			fyne.NewPos(0, 0),
			fyne.NewSize(100, 0),
			fyne.NewPos(0, dividerThickness()),
			fyne.NewSize(100, 0),
		},
	} {
		t.Run(name, func(t *testing.T) {
			objA := canvas.NewRectangle(color.White)
			objB := canvas.NewRectangle(color.Black)
			objA.SetMinSize(fyne.NewSize(10, 10))
			objB.SetMinSize(fyne.NewSize(30, 30))
			var c *Split
			if tt.horizontal {
				c = NewHSplit(objA, objB)
			} else {
				c = NewVSplit(objA, objB)
			}
			c.Resize(tt.size)

			assert.Equal(t, tt.wantLeadingPos, objA.Position(), "leading position")
			assert.Equal(t, tt.wantLeadingSize, objA.Size(), "leading size")
			assert.Equal(t, tt.wantTrailingPos, objB.Position(), "trailing position")
			assert.Equal(t, tt.wantTrailingSize, objB.Size(), "trailing size")
		})
	}
}

func TestSplitContainer_SetRatio(t *testing.T) {
	size := fyne.NewSize(100, 100)
	usableLength := 100 - float64(dividerThickness())

	objA := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})
	objB := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})

	t.Run("Horizontal", func(t *testing.T) {
		sc := NewHSplit(objA, objB)
		sc.Resize(size)
		t.Run("Leading", func(t *testing.T) {
			sc.SetOffset(0.75)
			sizeA := objA.Size()
			sizeB := objB.Size()
			assert.Equal(t, float32(0.75*usableLength), sizeA.Width)
			assert.Equal(t, float32(100), sizeA.Height)
			assert.Equal(t, float32(0.25*usableLength), sizeB.Width)
			assert.Equal(t, float32(100), sizeB.Height)
		})
		t.Run("Trailing", func(t *testing.T) {
			sc.SetOffset(0.25)
			sizeA := objA.Size()
			sizeB := objB.Size()
			assert.Equal(t, float32(0.25*usableLength), sizeA.Width)
			assert.Equal(t, float32(100), sizeA.Height)
			assert.Equal(t, float32(0.75*usableLength), sizeB.Width)
			assert.Equal(t, float32(100), sizeB.Height)
		})
	})
	t.Run("Vertical", func(t *testing.T) {
		sc := NewVSplit(objA, objB)
		sc.Resize(size)
		t.Run("Leading", func(t *testing.T) {
			sc.SetOffset(0.75)
			sizeA := objA.Size()
			sizeB := objB.Size()
			assert.Equal(t, float32(100), sizeA.Width)
			assert.Equal(t, float32(0.75*usableLength), sizeA.Height)
			assert.Equal(t, float32(100), sizeB.Width)
			assert.Equal(t, float32(0.25*usableLength), sizeB.Height)
		})
		t.Run("Trailing", func(t *testing.T) {
			sc.SetOffset(0.25)
			sizeA := objA.Size()
			sizeB := objB.Size()
			assert.Equal(t, float32(100), sizeA.Width)
			assert.Equal(t, float32(0.25*usableLength), sizeA.Height)
			assert.Equal(t, float32(100), sizeB.Width)
			assert.Equal(t, float32(0.75*usableLength), sizeB.Height)
		})
	})
}

func TestSplitContainer_SetRatio_limits(t *testing.T) {
	size := fyne.NewSize(50, 50)
	objA := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})
	objA.SetMinSize(size)
	objB := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})
	objB.SetMinSize(size)
	t.Run("Horizontal", func(t *testing.T) {
		sc := NewHSplit(objA, objB)
		t.Run("Leading", func(t *testing.T) {
			sc.SetOffset(1.0)
			sc.Resize(fyne.NewSize(200, 50))
			sizeA := objA.Size()
			sizeB := objB.Size()
			assert.Equal(t, 150-dividerThickness(), sizeA.Width)
			assert.Equal(t, float32(50), sizeA.Height)
			assert.Equal(t, float32(50), sizeB.Width)
			assert.Equal(t, float32(50), sizeB.Height)
		})
		t.Run("Trailing", func(t *testing.T) {
			sc.SetOffset(0.0)
			sc.Resize(fyne.NewSize(200, 50))
			sizeA := objA.Size()
			sizeB := objB.Size()
			assert.Equal(t, float32(50), sizeA.Width)
			assert.Equal(t, float32(50), sizeA.Height)
			assert.Equal(t, 150-dividerThickness(), sizeB.Width)
			assert.Equal(t, float32(50), sizeB.Height)
		})
	})
	t.Run("Vertical", func(t *testing.T) {
		sc := NewVSplit(objA, objB)
		t.Run("Leading", func(t *testing.T) {
			sc.SetOffset(1.0)
			sc.Resize(fyne.NewSize(50, 200))
			sizeA := objA.Size()
			sizeB := objB.Size()
			assert.Equal(t, float32(50), sizeA.Width)
			assert.Equal(t, 150-dividerThickness(), sizeA.Height)
			assert.Equal(t, float32(50), sizeB.Width)
			assert.Equal(t, float32(50), sizeB.Height)
		})
		t.Run("Trailing", func(t *testing.T) {
			sc.SetOffset(0.0)
			sc.Resize(fyne.NewSize(50, 200))
			sizeA := objA.Size()
			sizeB := objB.Size()
			assert.Equal(t, float32(50), sizeA.Width)
			assert.Equal(t, float32(50), sizeA.Height)
			assert.Equal(t, float32(50), sizeB.Width)
			assert.Equal(t, 150-dividerThickness(), sizeB.Height)
		})
	})
}

func TestSplitContainer_swap_contents(t *testing.T) {
	dl := dividerLength()
	dt := dividerThickness()
	initialWidth := 10 + 10 + dt
	initialHeight := fyne.Max(10, dl)
	expectedWidth := 100 + 10 + dt
	expectedHeight := fyne.Max(100, dl)

	objA := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})
	objA.SetMinSize(fyne.NewSize(10, 10))
	objB := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})
	objB.SetMinSize(fyne.NewSize(10, 10))
	objC := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})
	objC.SetMinSize(fyne.NewSize(100, 100))
	t.Run("Leading", func(t *testing.T) {
		sc := NewHSplit(objA, objB)
		min := sc.MinSize()
		assert.Equal(t, float32(initialWidth), min.Width)
		assert.Equal(t, float32(initialHeight), min.Height)
		sc.Leading = objC
		sc.Refresh()
		min = sc.MinSize()
		assert.Equal(t, float32(expectedWidth), min.Width)
		assert.Equal(t, float32(expectedHeight), min.Height)
	})
	t.Run("Trailing", func(t *testing.T) {
		sc := NewHSplit(objA, objB)
		min := sc.MinSize()
		assert.Equal(t, float32(initialWidth), min.Width)
		assert.Equal(t, float32(initialHeight), min.Height)
		sc.Trailing = objC
		sc.Refresh()
		min = sc.MinSize()
		assert.Equal(t, float32(expectedWidth), min.Width)
		assert.Equal(t, float32(expectedHeight), min.Height)
	})
}

func TestSplitContainer_divider_cursor(t *testing.T) {
	t.Run("Horizontal", func(t *testing.T) {
		divider := newDivider(&Split{Horizontal: true})
		assert.Equal(t, desktop.HResizeCursor, divider.Cursor())
	})
	t.Run("Vertical", func(t *testing.T) {
		divider := newDivider(&Split{Horizontal: false})
		assert.Equal(t, desktop.VResizeCursor, divider.Cursor())
	})
}

func TestSplitContainer_divider_drag(t *testing.T) {
	size := fyne.NewSize(10, 10)
	objA := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})
	objA.SetMinSize(size)
	objB := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})
	objB.SetMinSize(size)
	t.Run("Horizontal", func(t *testing.T) {
		split := NewHSplit(objA, objB)
		split.Resize(fyne.NewSize(108, 108))
		divider := test.WidgetRenderer(split).(*splitContainerRenderer).divider
		assert.Equal(t, 0.5, split.Offset)

		divider.Dragged(&fyne.DragEvent{
			PointEvent: fyne.PointEvent{Position: fyne.NewPos(13, 9)},
			Dragged:    fyne.NewDelta(10, 0),
		})
		assert.Equal(t, 0.6, split.Offset)

		divider.DragEnd()
		assert.Equal(t, 0.6, split.Offset)

		// test max limit
		divider.Dragged(&fyne.DragEvent{
			PointEvent: fyne.PointEvent{Position: fyne.NewPos(53, 9)},
			Dragged:    fyne.NewDelta(54, 0),
		})
		divider.Dragged(&fyne.DragEvent{
			PointEvent: fyne.PointEvent{Position: fyne.NewPos(53, 9)},
			Dragged:    fyne.NewDelta(54, 0),
		})
		assert.Equal(t, 0.9, split.Offset)

		// test returns to position after over-drag
		divider.Dragged(&fyne.DragEvent{
			PointEvent: fyne.PointEvent{Position: fyne.NewPos(-97, 9)},
			Dragged:    fyne.NewDelta(-108, 0),
		})
		divider.DragEnd()
		assert.Equal(t, 0.6, split.Offset)
	})
	t.Run("Vertical", func(t *testing.T) {
		split := NewVSplit(objA, objB)
		split.Resize(fyne.NewSize(108, 108))
		divider := test.WidgetRenderer(split).(*splitContainerRenderer).divider
		assert.Equal(t, 0.5, split.Offset)

		divider.Dragged(&fyne.DragEvent{
			PointEvent: fyne.PointEvent{Position: fyne.NewPos(13, 9)},
			Dragged:    fyne.NewDelta(0, 10),
		})
		assert.Equal(t, 0.6, split.Offset)

		divider.DragEnd()
		assert.Equal(t, 0.6, split.Offset)

		// test max limit
		divider.Dragged(&fyne.DragEvent{
			PointEvent: fyne.PointEvent{Position: fyne.NewPos(9, 53)},
			Dragged:    fyne.NewDelta(0, 54),
		})
		divider.Dragged(&fyne.DragEvent{
			PointEvent: fyne.PointEvent{Position: fyne.NewPos(9, 53)},
			Dragged:    fyne.NewDelta(0, 54),
		})
		assert.Equal(t, 0.9, split.Offset)

		// test returns to position after over-drag
		divider.Dragged(&fyne.DragEvent{
			PointEvent: fyne.PointEvent{Position: fyne.NewPos(9, -97)},
			Dragged:    fyne.NewDelta(0, -108),
		})
		divider.DragEnd()
		assert.Equal(t, 0.6, split.Offset)
	})
}

func TestSplitContainer_divider_drag_StartOffsetLessThanMinSize(t *testing.T) {
	size := fyne.NewSize(30, 30)
	objA := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})
	objA.SetMinSize(size)
	objB := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})
	objB.SetMinSize(size)
	t.Run("Horizontal", func(t *testing.T) {
		split := NewHSplit(objA, objB)
		split.Resize(fyne.NewSize(108, 108))
		divider := test.WidgetRenderer(split).(*splitContainerRenderer).divider
		t.Run("Leading", func(t *testing.T) {
			split.SetOffset(0.1)

			divider.Dragged(&fyne.DragEvent{
				Dragged: fyne.NewDelta(10, 0),
			})
			divider.DragEnd()

			assert.Equal(t, 0.4, split.Offset)
		})
		t.Run("Trailing", func(t *testing.T) {
			split.SetOffset(0.9)

			divider.Dragged(&fyne.DragEvent{
				Dragged: fyne.NewDelta(-10, 0),
			})
			divider.DragEnd()

			assert.Equal(t, 0.6, split.Offset)
		})
	})
	t.Run("Vertical", func(t *testing.T) {
		split := NewVSplit(objA, objB)
		split.Resize(fyne.NewSize(108, 108))
		divider := test.WidgetRenderer(split).(*splitContainerRenderer).divider
		t.Run("Leading", func(t *testing.T) {
			split.SetOffset(0.1)

			divider.Dragged(&fyne.DragEvent{
				Dragged: fyne.NewDelta(0, 10),
			})
			divider.DragEnd()

			assert.Equal(t, 0.4, split.Offset)
		})
		t.Run("Trailing", func(t *testing.T) {
			split.SetOffset(0.9)

			divider.Dragged(&fyne.DragEvent{
				Dragged: fyne.NewDelta(0, -10),
			})
			divider.DragEnd()

			assert.Equal(t, 0.6, split.Offset)
		})
	})
}

func TestSplitContainer_divider_hover(t *testing.T) {
	t.Run("Horizontal", func(t *testing.T) {
		divider := newDivider(&Split{Horizontal: true})
		assert.False(t, divider.hovered)

		divider.MouseIn(&desktop.MouseEvent{})
		assert.True(t, divider.hovered)

		divider.MouseOut()
		assert.False(t, divider.hovered)
	})
	t.Run("Vertical", func(t *testing.T) {
		divider := newDivider(&Split{Horizontal: false})
		assert.False(t, divider.hovered)

		divider.MouseIn(&desktop.MouseEvent{})
		assert.True(t, divider.hovered)

		divider.MouseOut()
		assert.False(t, divider.hovered)
	})
}

func TestSplitContainer_divider_MinSize(t *testing.T) {
	t.Run("Horizontal", func(t *testing.T) {
		divider := newDivider(&Split{Horizontal: true})
		min := divider.MinSize()
		assert.Equal(t, dividerThickness(), min.Width)
		assert.Equal(t, dividerLength(), min.Height)
	})
	t.Run("Vertical", func(t *testing.T) {
		divider := newDivider(&Split{Horizontal: false})
		min := divider.MinSize()
		assert.Equal(t, dividerLength(), min.Width)
		assert.Equal(t, dividerThickness(), min.Height)
	})
}

func TestSplitContainer_Hidden(t *testing.T) {
	dt := dividerThickness()
	size := fyne.NewSize(10, 10)
	objA := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})
	objA.SetMinSize(size)
	objB := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})

	t.Run("Horizontal_Leading", func(t *testing.T) {
		sc := NewHSplit(objA, objB)
		sc.Resize(fyne.NewSize(100, 100))
		assert.Equal(t, 50-(dt/2), sc.Leading.Size().Width)

		objA.Hide()
		sc.SetOffset(0)
		assert.Equal(t, float32(0), sc.Leading.Size().Width)
	})
	t.Run("Horizontal_Trailing", func(t *testing.T) {
		sc := NewHSplit(objA, objB)
		sc.Resize(fyne.NewSize(100, 100))
		assert.Equal(t, 50-(dt/2), sc.Trailing.Size().Width)

		objB.Hide()
		sc.SetOffset(1)
		assert.Equal(t, float32(0), sc.Trailing.Size().Width)
	})
	t.Run("Vertical_Leading", func(t *testing.T) {
		sc := NewVSplit(objA, objB)
		sc.Resize(fyne.NewSize(100, 100))
		assert.Equal(t, 50-(dt/2), sc.Leading.Size().Height)

		objA.Hide()
		sc.SetOffset(0)
		assert.Equal(t, float32(0), sc.Leading.Size().Height)
	})
	t.Run("Vertical_Trailing", func(t *testing.T) {
		sc := NewVSplit(objA, objB)
		sc.Resize(fyne.NewSize(100, 100))
		assert.Equal(t, 50-(dt/2), sc.Trailing.Size().Height)

		objB.Hide()
		sc.SetOffset(1)
		assert.Equal(t, float32(0), sc.Trailing.Size().Height)
	})
}
