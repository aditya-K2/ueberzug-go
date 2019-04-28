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
	img1, err := getImage("https://golang.org/doc/gopher/pencil/gophermega.jpg")
	if err != nil {
		t.Fatal(err)
	}

	i := NewImage(img1, 0, 0)
	defer i.Clear()

	img2, err := getImage("https://golang.org/doc/gopher/pencil/gophermega.jpg")
	if err != nil {
		t.Fatal(err)
	}

	j := NewImage(img2, 50, 75)
	defer j.Clear()

	time.Sleep(5 * time.Second)
}

func getImage(url string) (image.Image, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	img, _, err := image.Decode(r.Body)
	if err != nil {
		return nil, err
	}

	img = resize.Thumbnail(
		300, 300,
		img,
		resize.Bilinear,
	)

	return img, nil
}
