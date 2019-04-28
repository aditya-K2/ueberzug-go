package ueberzug

import "github.com/BurntSushi/xgbutil/xwindow"

func newParentWindow() *xwindow.Window {
	w, err := xwindow.Create(xutil, parent.Id)
	if err != nil {
		panic(err)
	}

	return w
}
