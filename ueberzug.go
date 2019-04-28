package ueberzug

import (
	"image"

	"github.com/BurntSushi/xgbutil/xgraphics"
)

// Image is the structure for the image
type Image struct {
	*xgraphics.Image
	X, Y int
}

// NewImage makes a new image
func NewImage(i image.Image, X, Y int) *Image {
	// Make a new Image
	img := &Image{
		Image: xgraphics.NewConvert(xutil, i),
		X:     X,
		Y:     Y,
	}

	img.Show()

	return img
}

// Show shows the image
func (i *Image) Show() {
	i.XSurfaceSet(child.Id)
	i.XDraw()
	i.XPaint(child.Id)

	child.Map()

	/*
		// Generate theh pixmap
		if err := i.image.CreatePixmap(); err != nil {
			panic(err)
		}

		// Draw the image onto the child window
		i.image.XExpPaint(child.Id, i.X, i.Y)
	*/
}

// Clear destroys the image
func (i *Image) Clear() {
	i.Destroy()
}
