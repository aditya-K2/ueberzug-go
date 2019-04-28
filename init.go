package ueberzug

import (
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xwindow"
)

var (
	// X is the active X connection, initialized on
	// application start, init()
	X      *xgb.Conn
	xutil  *xgbutil.XUtil
	parent *xwindow.Window
	child  *xwindow.Window
)

func init() {
	x, err := xgb.NewConnDisplay(":0")
	if err != nil {
		panic(err)
	}

	X = x

	x.Sync()

	xgb, err := xgbutil.NewConnXgb(x)
	if err != nil {
		panic(err)
	}

	xutil = xgb

	setup := xproto.Setup(x)
	root := setup.DefaultScreen(x).Root

	// Getting the terminal parent window on startup
	// is more reliable than getting it on-demand
	parent = xwindow.New(
		xgb, mustGetActiveWindow(root),
	)

	parent.Geometry()

	w, err := xwindow.Generate(xutil)
	if err != nil {
		panic(err)
	}

	if err := w.CreateChecked(
		xgb.RootWin(),
		parent.Geom.X(),
		parent.Geom.Y(),
		parent.Geom.Width(),
		parent.Geom.Height(),
		xproto.CwOverrideRedirect, 1,
	); err != nil {
		panic(err)
	}

	if err := xproto.ReparentWindowChecked(
		x, w.Id, parent.Id, 0, 0,
	).Check(); err != nil {
		panic(err)
	}

	child = w

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
}
