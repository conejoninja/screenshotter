package screenshot

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"time"
)

type Device struct {
	w, h, k int16
	buffer  [][]color.RGBA
}

func NewScreen(w, h int16) Device {
	buffer := make([][]color.RGBA, w)
	for i := int16(0); i < w; i++ {
		buffer[i] = make([]color.RGBA, h)
	}
	return Device{
		w:      w,
		h:      h,
		buffer: buffer,
		k:      0,
	}
}

func (d *Device) Size() (x, y int16) {
	return d.w, d.h
}

func (d *Device) SetPixel(x, y int16, c color.RGBA) {
	if x < 0 || y < 0 || x >= d.w || y >= d.h {
		return
	}
	d.buffer[x][y] = c
}

func (d *Device) Display() error {
	return nil
}

func (d *Device) FillScreen(c color.RGBA) {
	for i := int16(0); i < d.w; i++ {
		for j := int16(0); j < d.h; j++ {
			d.buffer[i][j] = c
		}
	}
}

func (d *Device) FillRectangle(x, y, w, h int16, c color.RGBA) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			d.buffer[i][j] = c
		}
	}
}

func (d *Device) SetScrollArea(x, y int16) {
}
func (d *Device) SetScroll(k int16) {
	d.k = k
}
func (d *Device) StopScroll() {
}

func (d *Device) Screenshot() {
	img := image.NewRGBA(image.Rect(0, 0, int(d.w), int(d.h)))
	for i := 0; i < int(d.w); i++ {
		z := i - int(d.k)
		if z < 0 {
			z += int(d.w)
		}

		for j := 0; j < int(d.h); j++ {
			img.Set(z, j, d.buffer[i][j])
		}
	}
	unixtime := strconv.Itoa(int(time.Now().UnixMicro()))
	f, _ := os.Create( /*time.DateTime + "_" +*/ unixtime + ".png")
	png.Encode(f, img)
}
