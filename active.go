package ueberzug

import (
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

const (
	atomActiveWindow = "_NET_ACTIVE_WINDOW"
)

func mustGetActiveWindowAtom() xproto.Atom {
	c := xproto.InternAtom(
		x, true,
		uint16(len(atomActiveWindow)),
		atomActiveWindow,
	)

	a, err := c.Reply()
	if err != nil {
		panic(err)
	}

	return a.Atom
}

func mustGetActiveWindow(root xproto.Window) xproto.Window {
	// https://github.com/BurntSushi/xgb/blob/master/examples/get-active-window/main.go#L44
	c := xproto.GetProperty(
		x, false,
		root, mustGetActiveWindowAtom(),
		xproto.GetPropertyTypeAny,
		0, (1<<32)-1,
	)

	r, err := c.Reply()
	if err != nil {
		panic(err)
	}

	return xproto.Window(xgb.Get32(r.Value))
}
