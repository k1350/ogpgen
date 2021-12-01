// package color は 16 進数のカラーコードから color.NRGBA を生成するパッケージです。
package color

import (
	"image/color"
	"regexp"
	"strconv"

	e "github.com/k1350/ogpgen/internal/errors"
	"github.com/pkg/errors"
)

// rgbaFormat はアルファ値を含むカラーコードを判定するための正規表現です。
var rgbaFormat = regexp.MustCompile(`^#[0-9a-fA-F]{2}[0-9a-fA-F]{2}[0-9a-fA-F]{2}[0-9a-fA-F]{2}$`)

// rgbFormat はアルファ値を含まないカラーコードを判定するための正規表現です。
var rgbFormat = regexp.MustCompile(`^#[0-9a-fA-F]{2}[0-9a-fA-F]{2}[0-9a-fA-F]{2}$`)

// ConvertHexToNRGBA は 16 進数のカラーコードから color.NRGBA を生成して返します。
// hex はカラーコードです。#333 のような省略形は対応していません。
func ConvertHexToNRGBA(hex string) (color.NRGBA, error) {

	if ok := rgbaFormat.MatchString(hex); ok {
		return parseColor(hex, true)
	}
	if ok := rgbFormat.MatchString(hex); ok {
		return parseColor(hex, false)
	}
	return color.NRGBA{0, 0, 0, 0}, e.ErrorInvalidColorFormat
}

// parseColor は 16 進数のカラーコードから color.NRGBA を生成して返します。
// hex はカラーコードです。#333 のような省略形は対応していません。
// existAlpha はアルファ値を含むかどうかを指定します。true の場合はアルファ値を読み取り、false の場合はアルファ値は FF とします。
func parseColor(hex string, existAlpha bool) (c color.NRGBA, err error) {
	err = nil
	var a int64
	a = 255
	if existAlpha {
		a, err = strconv.ParseInt(hex[7:9], 16, 64)
		if err != nil {
			err = errors.Wrap(e.ErrorInvalidColorFormat, err.Error())
			return
		}
	}

	r, err := strconv.ParseInt(hex[1:3], 16, 64)
	if err != nil {
		err = errors.Wrap(e.ErrorInvalidColorFormat, err.Error())
		return c, err
	}
	g, err := strconv.ParseInt(hex[3:5], 16, 64)
	if err != nil {
		err = errors.Wrap(e.ErrorInvalidColorFormat, err.Error())
		return c, err
	}
	b, err := strconv.ParseInt(hex[5:7], 16, 64)
	if err != nil {
		err = errors.Wrap(e.ErrorInvalidColorFormat, err.Error())
		return c, err
	}

	c.R = uint8(r)
	c.G = uint8(g)
	c.B = uint8(b)
	c.A = uint8(a)
	return
}
