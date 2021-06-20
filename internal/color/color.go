// package color は 16 進数のカラーコードから color.RGBA を生成するパッケージです。
package color

import (
	"errors"
	"image/color"
	"regexp"
	"strconv"
)

// rgbaFormat はアルファ値を含むカラーコードを判定するための正規表現です。
var rgbaFormat = regexp.MustCompile(`^#[0-9a-fA-F]{2}[0-9a-fA-F]{2}[0-9a-fA-F]{2}[0-9a-fA-F]{2}$`)

// rgbFormat はアルファ値を含まないカラーコードを判定するための正規表現です。
var rgbFormat = regexp.MustCompile(`^#[0-9a-fA-F]{2}[0-9a-fA-F]{2}[0-9a-fA-F]{2}$`)

// ConvertHexToRGBA は 16 進数のカラーコードから color.RGBA を生成して返します。
// hex はカラーコードです。#333 のような省略形は対応していません。
func ConvertHexToRGBA(hex string) (color.RGBA, error) {

	if ok := rgbaFormat.MatchString(hex); ok {
		return parseColor(hex, true)
	}
	if ok := rgbFormat.MatchString(hex); ok {
		return parseColor(hex, false)
	}
	return color.RGBA{0, 0, 0, 0}, errors.New("invalid color format")
}

// parseColor は 16 進数のカラーコードから color.RGBA を生成して返します。
// hex はカラーコードです。#333 のような省略形は対応していません。
// existAlpha はアルファ値を含むかどうかを指定します。true の場合はアルファ値を読み取り、false の場合はアルファ値は FF とします。
func parseColor(hex string, existAlpha bool) (c color.RGBA, err error) {
	err = nil
	var a int64
	a = 255
	if existAlpha {
		a, err = strconv.ParseInt(hex[7:9], 16, 64)
		if err != nil {
			return c, err
		}
	}
	co := float64(a) / 255

	r, err := strconv.ParseInt(hex[1:3], 16, 64)
	if err != nil {
		return c, err
	}
	g, err := strconv.ParseInt(hex[3:5], 16, 64)
	if err != nil {
		return c, err
	}
	b, err := strconv.ParseInt(hex[5:7], 16, 64)
	if err != nil {
		return c, err
	}

	c.R = uint8(float64(r) * co)
	c.G = uint8(float64(g) * co)
	c.B = uint8(float64(b) * co)
	c.A = uint8(a)
	return
}
