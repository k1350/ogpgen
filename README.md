![tn](https://github.com/k1350/ogpgen/blob/main/example/outsample.jpg)
## インストールとビルド
必要環境

- Git
- Go (version 1.16.5 で動作確認済)

```
git clone https://github.com/k1350/ogpgen.git
cd ogpgen
go build -o ogpgen ./cmd/ogpgen
```

## 使い方
最低限の引数指定時
```
.\ogpgen -text="OGP 画像生成ツール" -bpath="/your_template/template.jpg" -fpath="/your_font/03SmartFontUI.ttf"
```

引数全指定時
```
.\ogpgen -text="OGP 画像生成ツール" -bpath="/your_template/template.jpg" -fpath="/your_font/03SmartFontUI.ttf" -fsize=90 -fcolor="#A2Ad0580" -tmargin=75 -smargin=90 -lspace=10 -o="/out_dir/output.jpg"
```

### 引数の説明
- text: 【必須】描画したい文字列。
- bpath: 【必須】背景に使用する画像ファイルのパス。JPEG でなければならない。
- fpath: 【必須】使用するフォントファイルのパス。TrueType フォントでなければならない。
- fsize: フォントサイズ。0 より大きい数でなければならない。小数点以下も指定できる。未指定の場合は 100.0 になる。
- fcolor: フォント色を6桁または8桁（透明度指定あり）のカラーコードで指定する。未指定の場合は #000000 になる。
- tmargin: 上部マージン。文字が描画される位置が、正の値を指定すると下側に、負の値を指定すると上側にずれる。未指定の場合は 0 になる。
- smargin: 左右マージン。大きい値を指定するほど文字が中央に寄る。未指定の場合は 0 になる。
- lspace: 行間。正の値を指定すると行間が広がり、負の値を指定すると詰まる。未指定の場合は 0 になる。
- o: 出力ファイル名。未指定の場合は out.jpg になる。