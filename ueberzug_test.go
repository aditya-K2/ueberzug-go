package ueberzug

import (
	"image"
	_ "image/jpeg"
	"net/http"
	"testing"
	"time"

	"github.com/nfnt/resize"
)

func TestImage(t *testing.T) {
	r, err := http.Get("https://golang.org/doc/gopher/pencil/gophermega.jpg")
	if err != nil {
		t.Fatal(err)
	}

	defer r.Body.Close()

	img, _, err := image.Decode(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	img = resize.Thumbnail(
		300, 300,
		img,
		resize.Bilinear,
	)

	i := NewImage(img, 10, 10)
	defer i.Clear()

	time.Sleep(5 * time.Second)
}
