package ueberzug

import (
	"image"
	_ "image/jpeg"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/nfnt/resize"
)

func TestImage(t *testing.T) {
	img1, err := getImage("https://golang.org/doc/gopher/pencil/gophermega.jpg")
	if err != nil {
		t.Fatal(err)
	}

	println("Displaying img1")

	i, err := NewImage(img1, 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	defer i.Clear()
	defer i.Destroy()

	img2, err := getImage("_testdata/thonk.png")
	if err != nil {
		t.Fatal(err)
	}

	println("Displaying img2")

	j, err := NewImage(img2, 50, 75)
	if err != nil {
		t.Fatal(err)
	}

	defer j.Clear()
	defer j.Destroy()

	time.Sleep(5 * time.Second)
}

func getImage(url string) (image.Image, error) {
	var reader io.Reader

	if strings.HasPrefix(url, "http") {
		r, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer r.Body.Close()

		reader = r.Body
	} else {
		f, err := os.Open(url)
		if err != nil {
			return nil, err
		}

		defer f.Close()

		reader = f
	}

	img, _, err := image.Decode(reader)
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
