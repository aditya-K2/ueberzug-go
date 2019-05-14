package ueberzug

import (
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

const (
	atomActiveWindow = "_NET_ACTIVE_WINDOW"
)

func getActiveWindow(root xproto.Window) (xproto.Window, error) {
	i := xproto.InternAtom(
		x, true,
		uint16(len(atomActiveWindow)),
		atomActiveWindow,
	)

	a, err := i.Reply()
	if err != nil {
		return 0, err
	}

	// https://github.com/BurntSushi/xgb/blob/master/examples/get-active-window/main.go#L44
	c := xproto.GetProperty(
		x, false,
		root, a.Atom,
		xproto.GetPropertyTypeAny,
		0, (1<<32)-1,
	)

	r, err := c.Reply()
	if err != nil {
		panic(err)
	}

	return xproto.Window(xgb.Get32(r.Value)), nil
}
