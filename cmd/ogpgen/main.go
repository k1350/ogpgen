package main

import (
	"bufio"
	"flag"
	"image/jpeg"
	"log"
	"os"

	"github.com/k1350/ogpgen/internal/ogpgen"
)

func main() {
	text := flag.String("text", "", "Output Text")
	fpath := flag.String("fpath", "", "Font Path")
	fsize := flag.Float64("fsize", 100.0, "Font Size")
	fcolor := flag.String("fcolor", "#000000", "Font Color Code")
	bpath := flag.String("bpath", "", "Background Image Path")
	tmargin := flag.Int("tmargin", 0, "Top Margin")
	smargin := flag.Int("smargin", 0, "Side Margin")
	lspace := flag.Int("lspace", 0, "Line Space")
	out := flag.String("o", "out.jpg", "Output File Path")
	flag.Parse()

	d, err := ogpgen.NewOgpImage(
		ogpgen.Option{
			FontPath:            *fpath,
			FontSize:            *fsize,
			FontColor:           *fcolor,
			BackgroundImagePath: *bpath,
			TopMargin:           *tmargin,
			SideMargin:          *smargin,
			LineSpace:           *lspace,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	img, err := d.Draw(*text)
	if err != nil {
		log.Fatal(err)
	}

	newFile, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	b := bufio.NewWriter(newFile)
	if err = jpeg.Encode(b, img, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatal(err)
	}
}
