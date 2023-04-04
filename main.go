package main

import (
	"image/color"
	"time"

	"github.com/conejoninja/screenshotter/screenshot"

	qrcode "github.com/skip2/go-qrcode"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/gophers"
)

const (
	WIDTH  = 320
	HEIGHT = 240
)

const (
	BLACK = iota
	WHITE
	RED
)

var colors = []color.RGBA{
	color.RGBA{0, 0, 0, 255},
	color.RGBA{255, 255, 255, 255},
	color.RGBA{255, 0, 0, 255},
}

var rainbow []color.RGBA
var display screenshot.Device

func main() {
	rainbow = make([]color.RGBA, 256)
	for i := 0; i < 256; i++ {
		rainbow[i] = getRainbowRGB(uint8(i))
	}

	display = screenshot.NewScreen(WIDTH, HEIGHT)

	/*myNameIsRainbow("@_CONEJO")
	display.Screenshot()

	/*scroll("This badge", "runs", "TINYGO")
	for i := int16(10); i < WIDTH; i++ {
		display.SetScroll(i)
		display.Screenshot()
	} */

	//QR("https://gopherbadge.com")
	//display.Screenshot()

	blinkyRainbow("GopherCon EU", "Berlin, June 26-29")
}

func myNameIs(name string) {
	display.FillScreen(colors[WHITE])

	var r int16 = 10

	// black corners detail
	display.FillRectangle(0, 0, r, r, colors[BLACK])
	display.FillRectangle(0, HEIGHT-r, r, r, colors[BLACK])
	display.FillRectangle(WIDTH-r, 0, r, r, colors[BLACK])
	display.FillRectangle(WIDTH-r, HEIGHT-r, r, r, colors[BLACK])

	// round corners
	tinydraw.FilledCircle(&display, r, r, r, colors[RED])
	tinydraw.FilledCircle(&display, WIDTH-r-1, r, r, colors[RED])
	tinydraw.FilledCircle(&display, r, HEIGHT-r-1, r, colors[RED])
	tinydraw.FilledCircle(&display, WIDTH-r-1, HEIGHT-r-1, r, colors[RED])

	// top band
	display.FillRectangle(r, 0, WIDTH-2*r-1, r, colors[RED])
	display.FillRectangle(0, r, WIDTH, 54, colors[RED])

	// bottom band
	display.FillRectangle(r, HEIGHT-r-1, WIDTH-2*r-1, r+1, colors[RED])
	display.FillRectangle(0, HEIGHT-3*r-1, WIDTH, 2*r, colors[RED])

	// top text : my NAME is
	w32, _ := tinyfont.LineWidth(&freesans.Bold18pt7b, "HELLO")
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 34, "HELLO", colors[WHITE])

	w32, _ = tinyfont.LineWidth(&freesans.Oblique9pt7b, "my name is")
	tinyfont.WriteLine(&display, &freesans.Oblique9pt7b, (WIDTH-int16(w32))/2, 54, "my name is", colors[WHITE])

	// middle text
	w32, _ = tinyfont.LineWidth(&freesans.Bold24pt7b, name)
	if w32 < 300 {
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32))/2, 140, name, colors[BLACK])
	} else {
		w32, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, name)
		if w32 < 300 {
			tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 140, name, colors[BLACK])
		} else {
			w32, _ = tinyfont.LineWidth(&freesans.Bold12pt7b, name)
			if w32 < 300 {
				tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32))/2, 140, name, colors[BLACK])
			} else {
				w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, name)
				tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, 140, name, colors[BLACK])
			}
		}
	}

	// gophers
	tinyfont.WriteLineColors(&display, &gophers.Regular58pt, WIDTH-84, 208, "BE", []color.RGBA{getRainbowRGB(100), getRainbowRGB(200)})
}

func myNameIsRainbow(name string) {
	myNameIs(name)

	w32, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, name)
	size := 24
	if w32 < 300 {
		size = 24
	} else {
		w32, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, name)
		if w32 < 300 {
			size = 18
		} else {
			w32, _ = tinyfont.LineWidth(&freesans.Bold12pt7b, name)
			if w32 < 300 {
				size = 12
			} else {
				w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, name)
				size = 9
			}
		}
	}
	for i := 0; i < 230; i++ {
		if size == 24 {
			tinyfont.WriteLineColors(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32))/2, 140, name, rainbow[i:])
		} else if size == 18 {
			tinyfont.WriteLineColors(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 140, name, rainbow[i:])
		} else if size == 12 {
			tinyfont.WriteLineColors(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32))/2, 140, name, rainbow[i:])
		} else {
			tinyfont.WriteLineColors(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, 140, name, rainbow[i:])
		}
		i += 2

		display.Screenshot()
	}
}

func blinkyRainbow(topline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text so we could center them later
	w32top, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, topline)
	sizetop := 24
	if w32top < 300 {
		sizetop = 24
	} else {
		w32top, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, topline)
		if w32top < 300 {
			sizetop = 18
		} else {
			w32top, _ = tinyfont.LineWidth(&freesans.Bold12pt7b, topline)
			if w32top < 300 {
				sizetop = 12
			} else {
				w32top, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, topline)
				sizetop = 9
			}
		}
	}

	w32bottom, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, bottomline)
	sizebottom := 24
	if w32bottom < 300 {
		sizebottom = 24
	} else {
		w32bottom, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, bottomline)
		if w32bottom < 300 {
			sizebottom = 18
		} else {
			w32bottom, _ = tinyfont.LineWidth(&freesans.Bold12pt7b, bottomline)
			if w32bottom < 300 {
				sizebottom = 12
			} else {
				w32bottom, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, bottomline)
				sizebottom = 9
			}
		}
	}

	for i := int16(0); i < 20; i++ {
		// show black text
		if sizetop == 24 {
			tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32top))/2, 100, topline, getRainbowRGB(uint8(i*12)))
		} else if sizetop == 18 {
			tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32top))/2, 100, topline, getRainbowRGB(uint8(i*12)))
		} else if sizetop == 12 {
			tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32top))/2, 100, topline, getRainbowRGB(uint8(i*12)))
		} else {
			tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32top))/2, 100, topline, getRainbowRGB(uint8(i*12)))
		}

		if sizebottom == 24 {
			tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(uint8(i*12)))
		} else if sizebottom == 18 {
			tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(uint8(i*12)))
		} else if sizebottom == 12 {
			tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(uint8(i*12)))
		} else {
			tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(uint8(i*12)))
		}

		display.Screenshot()
	}
}

func scroll(topline, middleline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text, so we could center them later
	w32top, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, topline)
	sizetop := 24
	if w32top < 300 {
		sizetop = 24
	} else {
		w32top, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, topline)
		if w32top < 300 {
			sizetop = 18
		} else {
			w32top, _ = tinyfont.LineWidth(&freesans.Bold12pt7b, topline)
			if w32top < 300 {
				sizetop = 12
			} else {
				w32top, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, topline)
				sizetop = 9
			}
		}
	}

	w32middle, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, middleline)
	sizemiddle := 24
	if w32middle < 300 {
		sizemiddle = 24
	} else {
		w32middle, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, middleline)
		if w32middle < 300 {
			sizemiddle = 18
		} else {
			w32middle, _ = tinyfont.LineWidth(&freesans.Bold12pt7b, middleline)
			if w32middle < 300 {
				sizemiddle = 12
			} else {
				w32middle, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, middleline)
				sizemiddle = 9
			}
		}
	}

	w32bottom, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, bottomline)
	sizebottom := 24
	if w32bottom < 300 {
		sizebottom = 24
	} else {
		w32bottom, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, bottomline)
		if w32bottom < 300 {
			sizebottom = 18
		} else {
			w32bottom, _ = tinyfont.LineWidth(&freesans.Bold12pt7b, bottomline)
			if w32bottom < 300 {
				sizebottom = 12
			} else {
				w32bottom, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, bottomline)
				sizebottom = 9
			}
		}
	}

	if sizetop == 24 {
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32top))/2, 70, topline, getRainbowRGB(200))
	} else if sizetop == 18 {
		tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32top))/2, 70, topline, getRainbowRGB(200))
	} else if sizetop == 12 {
		tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32top))/2, 70, topline, getRainbowRGB(200))
	} else {
		tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32top))/2, 70, topline, getRainbowRGB(200))
	}

	if sizemiddle == 24 {
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32middle))/2, 120, middleline, getRainbowRGB(80))
	} else if sizemiddle == 18 {
		tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32middle))/2, 120, middleline, getRainbowRGB(80))
	} else if sizemiddle == 12 {
		tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32middle))/2, 120, middleline, getRainbowRGB(80))
	} else {
		tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32middle))/2, 120, middleline, getRainbowRGB(80))
	}

	if sizebottom == 24 {
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(120))
	} else if sizebottom == 18 {
		tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(120))
	} else if sizebottom == 12 {
		tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(120))
	} else {
		tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(120))
	}

	display.SetScrollArea(0, 0)
	for k := 0; k < 4; k++ {
		for i := int16(319); i >= 0; i-- {

			display.SetScroll(i)
			time.Sleep(10 * time.Millisecond)
		}
	}
	display.SetScroll(0)
	display.StopScroll()
}

func logo() {
	bgColor := color.RGBA{109, 0, 140, 255}
	white := color.RGBA{255, 255, 255, 255}
	display.FillScreen(bgColor)

	display.FillRectangle(6, 166, 308, 21, white)

	tinydraw.FilledCircle(&display, 282, 130, 9, white)
	tinydraw.Line(&display, 259, 110, 298, 149, bgColor)
	tinydraw.Line(&display, 260, 110, 299, 149, bgColor)
	tinydraw.Line(&display, 261, 110, 300, 149, bgColor)
	tinydraw.Line(&display, 262, 110, 301, 149, bgColor)
	tinydraw.Line(&display, 263, 110, 302, 149, bgColor)
	tinydraw.Line(&display, 264, 110, 303, 149, bgColor)
	tinydraw.Line(&display, 265, 110, 304, 149, bgColor)

	display.FillRectangle(250, 98, 11, 63, white)
	display.FillRectangle(250, 98, 44, 11, white)

	display.FillRectangle(270, 150, 44, 11, white)
	display.FillRectangle(303, 98, 11, 63, white)

	tinyfont.WriteLine(&display, &freesans.Regular18pt7b, 6, 109, "Purple", white)
	tinyfont.WriteLine(&display, &freesans.Regular18pt7b, 6, 153, "Hardware by", white)

}

func getRainbowRGB(i uint8) color.RGBA {
	if i < 85 {
		return color.RGBA{i * 3, 255 - i*3, 0, 255}
	} else if i < 170 {
		i -= 85
		return color.RGBA{255 - i*3, 0, i * 3, 255}
	}
	i -= 170
	return color.RGBA{0, i * 3, 255 - i*3, 255}
}

func QR(text string) {
	qr, err := qrcode.New(text, qrcode.Medium)
	if err != nil {
		println(err, 123)
	}

	qrbytes := qr.Bitmap()
	size := int16(len(qrbytes))

	factor := int16(HEIGHT / len(qrbytes))

	bx := (WIDTH - size*factor) / 2
	by := (HEIGHT - size*factor) / 2
	display.FillScreen(color.RGBA{109, 0, 140, 255})
	for y := int16(0); y < size; y++ {
		for x := int16(0); x < size; x++ {
			if qrbytes[y][x] {
				display.FillRectangle(bx+x*factor, by+y*factor, factor, factor, colors[0])
			} else {
				display.FillRectangle(bx+x*factor, by+y*factor, factor, factor, colors[1])
			}
		}
	}

}
