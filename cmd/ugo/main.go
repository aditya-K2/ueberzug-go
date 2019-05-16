package main

import (
	"flag"
	"image"
	"log"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/nfnt/resize"
	"gitlab.com/diamondburned/ueberzug-go"
)

func main() {
	var (
		x, y, w, h int
	)

	flag.IntVar(&x, "x", 0, "")
	flag.IntVar(&y, "y", 0, "")
	flag.IntVar(&w, "w", 0, "")
	flag.IntVar(&h, "h", 0, "")
	flag.Parse()

	log.SetPrefix("")

	img, _, err := image.Decode(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	if w != 0 || h != 0 {
		img = resize.Thumbnail(
			uint(w), uint(h), img, resize.Bilinear,
		)
	}

	i, err := ueberzug.NewImage(img, x, y)
	if err != nil {
		log.Fatalln(err)
	}

	i.Show()

	defer i.Destroy()

	select {}
}
