package ueberzug

import (
	"image"

	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/BurntSushi/xgbutil/xwindow"
)

// Image is the structure for the image
type Image struct {
	*xgraphics.Image
	win *xwindow.Window
}

// NewImage makes a new image
func NewImage(img image.Image, X, Y int) *Image {
	bounds := img.Bounds()

	// Make a new Image
	i := &Image{
		Image: xgraphics.NewConvert(xutil, img),
		win: newChildWindow(
			X, Y,
			bounds.Dx(),
			bounds.Dy(),
		),
	}

	i.XSurfaceSet(i.win.Id)
	i.XDraw()
	i.XPaint(i.win.Id)

	i.Show()

	return i
}

// Show shows the image
func (i *Image) Show() {
	i.win.Map()
}

// Clear clears the image
func (i *Image) Clear() {
	i.win.Unmap()
}

// Destroy destroys the image and window, freeing up
// resources
func (i *Image) Destroy() {
	i.Destroy()
	i.win.Destroy()
}
