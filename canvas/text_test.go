package canvas_test

import (
	"image/color"
	"testing"

	"github.com/Anwert/fyne/v2"
	"github.com/Anwert/fyne/v2/canvas"
	"github.com/Anwert/fyne/v2/test"
	"github.com/Anwert/fyne/v2/theme"

	"github.com/stretchr/testify/assert"
)

func TestText_MinSize(t *testing.T) {
	text := canvas.NewText("Test", color.NRGBA{0, 0, 0, 0xff})
	min := text.MinSize()

	assert.True(t, min.Width > 0)
	assert.True(t, min.Height > 0)

	text = canvas.NewText("Test2", color.NRGBA{0, 0, 0, 0xff})
	min2 := text.MinSize()
	assert.True(t, min2.Width > min.Width)
}

func TestText_MinSize_NoMultiLine(t *testing.T) {
	text := canvas.NewText("Break", color.NRGBA{0, 0, 0, 0xff})
	min := text.MinSize()

	text = canvas.NewText("Bre\nak", color.NRGBA{0, 0, 0, 0xff})
	min2 := text.MinSize()
	assert.True(t, min2.Width > min.Width)
	assert.True(t, min2.Height == min.Height)
}

func TestText_Layout(t *testing.T) {
	test.NewApp()
	defer test.NewApp()

	for name, tt := range map[string]struct {
		text  string
		align fyne.TextAlign
		size  fyne.Size
	}{
		"short_leading_small": {
			text:  "abc",
			align: fyne.TextAlignLeading,
			size:  fyne.NewSize(1, 1),
		},
		"short_leading_large": {
			text:  "abc",
			align: fyne.TextAlignLeading,
			size:  fyne.NewSize(500, 101),
		},
		"long_leading_small": {
			text:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			align: fyne.TextAlignLeading,
			size:  fyne.NewSize(1, 1),
		},
		"long_leading_large": {
			text:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			align: fyne.TextAlignLeading,
			size:  fyne.NewSize(500, 101),
		},
		"short_center_small": {
			text:  "abc",
			align: fyne.TextAlignCenter,
			size:  fyne.NewSize(1, 1),
		},
		"short_center_large": {
			text:  "abc",
			align: fyne.TextAlignCenter,
			size:  fyne.NewSize(500, 101),
		},
		"long_center_small": {
			text:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			align: fyne.TextAlignCenter,
			size:  fyne.NewSize(1, 1),
		},
		"long_center_large": {
			text:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			align: fyne.TextAlignCenter,
			size:  fyne.NewSize(500, 101),
		},
		"short_trailing_small": {
			text:  "abc",
			align: fyne.TextAlignTrailing,
			size:  fyne.NewSize(1, 1),
		},
		"short_trailing_large": {
			text:  "abc",
			align: fyne.TextAlignTrailing,
			size:  fyne.NewSize(500, 101),
		},
		"long_trailing_small": {
			text:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			align: fyne.TextAlignTrailing,
			size:  fyne.NewSize(1, 1),
		},
		"long_trailing_large": {
			text:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			align: fyne.TextAlignTrailing,
			size:  fyne.NewSize(500, 101),
		},
	} {
		t.Run(name, func(t *testing.T) {
			text := canvas.NewText(tt.text, theme.ForegroundColor())
			text.Alignment = tt.align
			text.Resize(tt.size)

			test.AssertObjectRendersToImage(t, "text/layout_"+name+".png", text)
		})
	}
}

func TestText_CarriageReturn(t *testing.T) {
	test.NewApp()
	defer test.NewApp()

	for name, tt := range map[string]struct {
		text  string
		align fyne.TextAlign
		size  fyne.Size
	}{
		"short_leading_small": {
			text:  "\ra\rb\rc\r",
			align: fyne.TextAlignLeading,
			size:  fyne.NewSize(1, 1),
		},
		"short_leading_large": {
			text:  "\ra\rb\rc\r",
			align: fyne.TextAlignLeading,
			size:  fyne.NewSize(500, 101),
		},
		"long_leading_small": {
			text:  "\ra\rb\rc\rd\re\rf\rg\rh\ri\rj\rk\rl\rm\rn\ro\rp\rq\rr\rs\rt\ru\rv\rw\rx\ry\rz\rA\rB\rC\rD\rE\rF\rG\rH\rI\rJ\rK\rL\rM\rN\rO\rP\rQ\rR\rS\rT\rU\rV\rW\rX\rY\rZ\r",
			align: fyne.TextAlignLeading,
			size:  fyne.NewSize(1, 1),
		},
		"long_leading_large": {
			text:  "\ra\rb\rc\rd\re\rf\rg\rh\ri\rj\rk\rl\rm\rn\ro\rp\rq\rr\rs\rt\ru\rv\rw\rx\ry\rz\rA\rB\rC\rD\rE\rF\rG\rH\rI\rJ\rK\rL\rM\rN\rO\rP\rQ\rR\rS\rT\rU\rV\rW\rX\rY\rZ\r",
			align: fyne.TextAlignLeading,
			size:  fyne.NewSize(500, 101),
		},
		"short_center_small": {
			text:  "\ra\rb\rc\r",
			align: fyne.TextAlignCenter,
			size:  fyne.NewSize(1, 1),
		},
		"short_center_large": {
			text:  "\ra\rb\rc\r",
			align: fyne.TextAlignCenter,
			size:  fyne.NewSize(500, 101),
		},
		"long_center_small": {
			text:  "\ra\rb\rc\rd\re\rf\rg\rh\ri\rj\rk\rl\rm\rn\ro\rp\rq\rr\rs\rt\ru\rv\rw\rx\ry\rz\rA\rB\rC\rD\rE\rF\rG\rH\rI\rJ\rK\rL\rM\rN\rO\rP\rQ\rR\rS\rT\rU\rV\rW\rX\rY\rZ\r",
			align: fyne.TextAlignCenter,
			size:  fyne.NewSize(1, 1),
		},
		"long_center_large": {
			text:  "\ra\rb\rc\rd\re\rf\rg\rh\ri\rj\rk\rl\rm\rn\ro\rp\rq\rr\rs\rt\ru\rv\rw\rx\ry\rz\rA\rB\rC\rD\rE\rF\rG\rH\rI\rJ\rK\rL\rM\rN\rO\rP\rQ\rR\rS\rT\rU\rV\rW\rX\rY\rZ\r",
			align: fyne.TextAlignCenter,
			size:  fyne.NewSize(500, 101),
		},
		"short_trailing_small": {
			text:  "\ra\rb\rc\r",
			align: fyne.TextAlignTrailing,
			size:  fyne.NewSize(1, 1),
		},
		"short_trailing_large": {
			text:  "\ra\rb\rc\r",
			align: fyne.TextAlignTrailing,
			size:  fyne.NewSize(500, 101),
		},
		"long_trailing_small": {
			text:  "\ra\rb\rc\rd\re\rf\rg\rh\ri\rj\rk\rl\rm\rn\ro\rp\rq\rr\rs\rt\ru\rv\rw\rx\ry\rz\rA\rB\rC\rD\rE\rF\rG\rH\rI\rJ\rK\rL\rM\rN\rO\rP\rQ\rR\rS\rT\rU\rV\rW\rX\rY\rZ\r",
			align: fyne.TextAlignTrailing,
			size:  fyne.NewSize(1, 1),
		},
		"long_trailing_large": {
			text:  "\ra\rb\rc\rd\re\rf\rg\rh\ri\rj\rk\rl\rm\rn\ro\rp\rq\rr\rs\rt\ru\rv\rw\rx\ry\rz\rA\rB\rC\rD\rE\rF\rG\rH\rI\rJ\rK\rL\rM\rN\rO\rP\rQ\rR\rS\rT\rU\rV\rW\rX\rY\rZ\r",
			align: fyne.TextAlignTrailing,
			size:  fyne.NewSize(500, 101),
		},
	} {
		t.Run(name, func(t *testing.T) {
			text := canvas.NewText(tt.text, theme.ForegroundColor())
			text.Alignment = tt.align
			text.Resize(tt.size)

			test.AssertObjectRendersToImage(t, "text/layout_"+name+".png", text)
		})
	}
}
