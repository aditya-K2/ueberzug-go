package ueberzug

import (
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xwindow"
)

var (
	x      *xgb.Conn
	xutil  *xgbutil.XUtil
	parent *xwindow.Window
)

// Initialize initalizes an X connection
func Initialize() error {
	if x != nil {
		return nil
	}

	c, err := xgb.NewConn()
	if err != nil {
		return err
	}

	x = c

	x.Sync()

	xgb, err := xgbutil.NewConnXgb(x)
	if err != nil {
		panic(err)
	}

	xutil = xgb

	setup := xproto.Setup(x)
	root := setup.DefaultScreen(x).Root

	active, err := getActiveWindow(root)
	if err != nil {
		return err
	}

	// Getting the terminal parent window on startup
	// is more reliable than getting it on-demand
	parent = xwindow.New(
		xgb, active,
	)

	/*
		parent.Geometry()

		// Listen to resizes on the parent
		parent.Listen(xproto.EventMaskSubstructureRedirect)

		// Callback to resize events
		xevent.ClientMessageFun(
			func(X *xgbutil.XUtil, ev xevent.ClientMessageEvent) {
				g, err := parent.Geometry()
				if err != nil {
					panic(err)
				}

				parent.Geom = g

				child.Resize(g.Width(), g.Height())
			},
		).Connect(xutil, parent.Id)
	*/

	return nil
}

// Close frees things
func Close() {
	x.Close()
}
