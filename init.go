package ueberzug

import (
	"time"

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

	g, err := xwindow.Create(xutil, xgb.RootWin())
	if err != nil {
		panic(err)
	}

	child = g
	child.Create(
		child.Id,
		0, 0, 200, 200,
		xproto.CwBackPixel, 0xffffff,
	)

	child.Map()

	time.Sleep(8 * time.Second)

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
