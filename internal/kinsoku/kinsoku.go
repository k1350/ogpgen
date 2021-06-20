// Package kinsoku は禁則処理を行うパッケージです。
package kinsoku

import (
	"bytes"
	"regexp"

	"golang.org/x/image/font"
)

// invalidStart は行頭禁則文字を定義した map です。
var invalidStart = map[string]bool{
	",": true, ")": true, "]": true, "｝": true, "、": true, "〕": true, "〉": true, "》": true, "」": true, "』": true, "】": true, "〙": true, "〗": true, "〟": true, "’": true, "”": true, "｠": true, "»": true,
	"ゝ": true, "ゞ": true, "ー": true, "ァ": true, "ィ": true, "ゥ": true, "ェ": true, "ォ": true, "ッ": true, "ャ": true, "ュ": true, "ョ": true, "ヮ": true, "ヵ": true, "ヶ": true,
	"ぁ": true, "ぃ": true, "ぅ": true, "ぇ": true, "ぉ": true, "っ": true, "ゃ": true, "ゅ": true, "ょ": true, "ゎ": true, "ゕ": true, "ゖ": true, "ㇰ": true, "ㇱ": true, "ㇲ": true, "ㇳ": true, "ㇴ": true,
	"ㇵ": true, "ㇶ": true, "ㇷ": true, "ㇸ": true, "ㇹ": true, "ㇷ゚": true, "ㇺ": true, "ㇻ": true, "ㇼ": true, "ㇽ": true, "ㇾ": true, "ㇿ": true, "々": true, "〻": true,
	"‐": true, "゠": true, "–": true, "〜": true, "～": true,
	"?": true, "!": true, "‼": true, "⁇": true, "⁈": true, "⁉": true,
	"・": true, ":": true, ";": true, "/": true, "。": true, ".": true,
}

// invalidEnd は行末禁則文字を定義した map です。
var invalidEnd = map[string]bool{
	"(": true, "[": true, "｛": true, "〔": true, "〈": true, "《": true, "「": true, "『": true,
	"【": true, "〘": true, "〖": true, "〝": true, "‘": true, "“": true, "｟": true, "«": true,
}

// 分離禁則の単語を検出するための正規表現です。
var notSplit = regexp.MustCompile(`^([a-zA-Z]{2}|,[0-9]|[0-9],|[0-9]{2}|——|……|‥‥|〳〳|〴〴|〵〵)$`)

// ExecKinsoku は禁則処理を行いながら半角文字相当の文字幅 width を超えないように文字列 text を改行します。
// 改行した結果は配列で返します。
func ExecKinsoku(face font.Face, width int, text string) (words []string) {

	var tmpbuf bytes.Buffer
	for _, r := range text {

		tmpbuf.WriteString(string(r))
		for {
			if font.MeasureString(face, tmpbuf.String()).Ceil() > width {
				tmpr := []rune(tmpbuf.String())
				splitidx := determineSplitIndex(tmpr, len(tmpr)-1)
				words = append(words, string(tmpr[:splitidx]))
				tmpbuf.Reset()
				tmpbuf.WriteString(string(tmpr[splitidx:]))
			}
			// 改行位置が確定後、残った文字列が文字幅以下なら次の文字へ進む
			if font.MeasureString(face, tmpbuf.String()).Ceil() <= width {
				break
			}
		}
	}
	if tmpbuf.Len() > 0 {
		words = append(words, tmpbuf.String())
	}
	return words
}

// determineSplitIndex は、r に対して禁則処理を考慮して改行位置を決定します。
func determineSplitIndex(r []rune, defaultidx int) (idx int) {
	idx = defaultidx
	for {

		// idx の位置で改行可能なら改行位置を確定する
		if ok := canLineBreak(r, idx); ok {
			break
		}

		// 改行位置を更に前にずらす
		idx--

		// 前の行全部を次の行に回しても改行できない場合は諦めて最初の改行位置で改行する
		if idx < 1 {
			idx = defaultidx
			break
		}
	}
	return
}

// canLineBreak は、r を idx の位置で改行可能かどうかを返します。
func canLineBreak(r []rune, idx int) bool {
	nextr := r[idx:]
	// 次の行の行頭が行頭禁則文字なら改行不可
	if _, ok := invalidStart[string(nextr[:1])]; ok {
		return false
	}

	prevr := r[:idx]
	prevend := len(prevr) - 1
	// 前の行の行末が行末禁則文字なら改行不可
	if _, ok := invalidEnd[string(prevr[prevend:])]; ok {
		return false
	}

	// 分離禁則チェック
	s := string(prevr[prevend:]) + string(nextr[:1])
	if ok := notSplit.MatchString(s); ok {
		return false
	}

	return true
}
