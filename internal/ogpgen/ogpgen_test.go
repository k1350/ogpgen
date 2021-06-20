package ogpgen

import (
	"bufio"
	"image/jpeg"
	"log"
	"os"
	"testing"
)

func TestDraw(t *testing.T) {
	d, err := NewOgpImage(
		Option{
			FontPath:            "../../../ogpgen_files/fonts/03SmartFontUI.ttf",
			FontSize:            90.0,
			BackgroundImagePath: "../../../ogpgen_files/templates/template.jpg",
			TopMargin:           50,
			LineSpace:           10,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	img, err := d.Draw("あいAB!102")
	if err != nil {
		log.Fatal(err)
	}

	fn := "out.jpg"
	newFile, err := os.Create(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	b := bufio.NewWriter(newFile)
	if err = jpeg.Encode(b, img, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatal(err)
	}
	if err := os.Remove(fn); err != nil {
		log.Fatal(err)
	}
}
