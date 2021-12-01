// package ogpgen は OGP 画像を生成するパッケージです。
package ogpgen

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	c "github.com/k1350/ogpgen/internal/color"
	e "github.com/k1350/ogpgen/internal/errors"
	"github.com/k1350/ogpgen/internal/kinsoku"
	"github.com/pkg/errors"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// OgpImage は OGP 画像生成に必要なフィールド等を定義した構造体です。
type OgpImage struct {
	Font            *truetype.Font
	BackgroundImage image.Image
	FontSize        float64
	FontColor       color.NRGBA
	TopMargin       int
	SideMargin      int
	LineSpace       int
}

// Option は NewOgpImage の引数を定義したコンストラクタです。
// FontPath は OGP 画像に描画する文字列の TrueType フォントのパスを指定します。
// FontSize は OGP 画像に描画する文字列のフォントサイズを指定します。
// BackgroudImagePath は OGP 画像のベースになる jpeg 画像のパスを指定します。
// TopMargin は 文字列描画時の上部のマージンを微調整したい場合に指定します。
// SideMargin は 文字列描画時の左右のマージンを調整したい場合に指定します。
// LineSpace は 文字列描画時の行間を調整したい場合に指定します。
type Option struct {
	FontPath            string
	FontSize            float64
	FontColor           string
	BackgroundImagePath string
	TopMargin           int
	SideMargin          int
	LineSpace           int
}

// NewOgpImage は OgpImage のコンストラクタです。
func NewOgpImage(opt Option) (o *OgpImage, err error) {
	o = &OgpImage{}
	err = nil

	ftBin, err := ioutil.ReadFile(opt.FontPath)
	if err != nil {
		err = errors.Wrap(e.ErrorFontReadFailed, err.Error())
		return
	}
	ft, err := truetype.Parse(ftBin)
	if err != nil {
		err = errors.Wrap(e.ErrorFontParseFailed, err.Error())
		return
	}
	o.Font = ft

	if opt.FontSize <= 0.0 {
		err = e.ErrorFontSizeNegative
		return
	}
	o.FontSize = opt.FontSize

	if opt.FontColor == "" {
		opt.FontColor = "#000000"
	}
	c, err := c.ConvertHexToNRGBA(opt.FontColor)
	if err != nil {
		return
	}
	o.FontColor = c

	file, err := os.Open(opt.BackgroundImagePath)
	if err != nil {
		err = errors.Wrap(e.ErrorBackgroundImageReadFailed, err.Error())
		return
	}
	defer file.Close()
	img, err := jpeg.Decode(file)
	if err != nil {
		err = errors.Wrap(e.ErrorBackgroundImageDecodeFailed, err.Error())
		return
	}
	o.BackgroundImage = img

	o.TopMargin = opt.TopMargin
	o.SideMargin = opt.SideMargin
	o.LineSpace = opt.LineSpace

	return
}

// Draw は OGP 画像のベース画像に指定された文字列 text を描画します。
func (o *OgpImage) Draw(text string) (img *image.RGBA, err error) {
	rect := image.Rectangle{image.Pt(0, 0), o.BackgroundImage.Bounds().Size()}
	img = image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), o.BackgroundImage, image.ZP, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(o.Font)
	c.SetFontSize(o.FontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.NewUniform(o.FontColor))
	c.SetHinting(font.HintingNone)

	imageWidth := o.BackgroundImage.Bounds().Dx()
	imageHeight := o.BackgroundImage.Bounds().Dy()
	face := truetype.NewFace(o.Font, &truetype.Options{Size: o.FontSize})

	words := kinsoku.ExecKinsoku(face, (imageWidth - o.SideMargin*2), text)

	rows := len(words)
	for i, v := range words {
		textWidth := font.MeasureString(face, v).Ceil()
		dot := fixed.Point26_6{X: fixed.I((imageWidth - textWidth) / 2), Y: fixed.I((imageHeight-face.Metrics().Height.Ceil()*rows)/2 + (i+1)*(face.Metrics().Height.Ceil()+o.LineSpace) + o.TopMargin)}
		_, err = c.DrawString(v, dot)
		if err != nil {
			err = errors.Wrap(e.ErrorDrawFailed, err.Error())
			return
		}
	}

	return
}
