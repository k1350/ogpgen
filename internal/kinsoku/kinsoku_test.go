package kinsoku

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
	"testing"
)

type tcase struct {
	in   string
	want []string
}

func TestExecKinsoku(t *testing.T) {
	ftBin, err := ioutil.ReadFile("../../../ogpgen_files/fonts/03SmartFontUI.ttf")
	if err != nil {
		log.Fatal(err)
	}
	ft, err := truetype.Parse(ftBin)
	if err != nil {
		log.Fatal(err)
	}
	face := truetype.NewFace(ft, &truetype.Options{Size: 1.0})

	width := 5

	t.Run("行頭禁則文字", func(t *testing.T) {
		cases := []tcase{
			{
				in:   "あいうえお、かきくけこ。",
				want: []string{"あいうえ", "お、かきく", "けこ。"},
			},
			{
				in:   "あお、〕:,゠ㇷ!.。かきくけこ",
				want: []string{"あ", "お、〕:,゠", "ㇷ!.。かき", "くけこ"},
			},
		}

		assertKinsoku(t, face, width, cases)
	})
	t.Run("行末禁則文字", func(t *testing.T) {
		cases := []tcase{
			{
				in:   "「はい」「いいえ」",
				want: []string{"「はい」", "「いいえ」"},
			},
			{
				in:   "あお（[｛〔“〘«《かきくけこ",
				want: []string{"あお（", "[｛〔“〘«", "《かきくけ", "こ"},
			},
		}

		assertKinsoku(t, face, width, cases)
	})
	t.Run("分離禁則文字", func(t *testing.T) {
		cases := []tcase{
			{
				in:   "人口は約125,36万人",
				want: []string{"人口は約", "125,36万", "人"},
			},
			{
				in:   "Lorem ipsum dolor sit amet",
				want: []string{"Lorem ", "ipsum ", "dolor sit ", "amet"},
			},
		}

		assertKinsoku(t, face, width, cases)
	})
}

func assertKinsoku(t *testing.T, face font.Face, width int, cases []tcase) {
	for _, c := range cases {
		got := ExecKinsoku(face, width, c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got:\n%v\n\nwant:\n%v", strings.Join(got[:], "\n"), strings.Join(c.want[:], "\n"))
		}
	}
}
