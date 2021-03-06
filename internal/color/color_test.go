package color

import (
	"image/color"
	"reflect"
	"testing"
)

type tcase struct {
	in   string
	want color.NRGBA
}

func TestConvertHexToNRGBA(t *testing.T) {
	cases := []tcase{
		{
			in: "#FF0000",
			want: color.NRGBA{
				R: uint8(255),
				G: uint8(0),
				B: uint8(0),
				A: uint8(255),
			},
		},
		{
			in: "#00FF00",
			want: color.NRGBA{
				R: uint8(0),
				G: uint8(255),
				B: uint8(0),
				A: uint8(255),
			},
		},
		{
			in: "#0000FF",
			want: color.NRGBA{
				R: uint8(0),
				G: uint8(0),
				B: uint8(255),
				A: uint8(255),
			},
		},
		{
			in: "#FEC0AD80",
			want: color.NRGBA{
				R: uint8(254),
				G: uint8(192),
				B: uint8(173),
				A: uint8(128),
			},
		},
	}

	for _, c := range cases {
		got, _ := ConvertHexToNRGBA(c.in)
		assertColor(t, got, c.want)
	}
}

func assertColor(t *testing.T, c color.NRGBA, wc color.NRGBA) {
	var got [4]uint32
	got[0], got[1], got[2], got[3] = c.RGBA()
	var want [4]uint32
	want[0], want[1], want[2], want[3] = wc.RGBA()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:\n%v\n\nwant:\n%v", got, want)
	}
}
