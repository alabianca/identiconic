package identiconic

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"image"
	"image/color"
	"math"
	"strconv"
)

const MaxSize = 10

func getDefaultOptions() *Options {
	return &Options{
		Size:     10,
		CellSize: 20,
	}
}

// Options to provide to the CreateIdenticon func
type Options struct {
	// Determines the size of the grid
	// A larger size will result in an identicon
	// more unique to the given string
	// Note: cannot exceed MaxSize
	Size int
	// Determines the size of a given cell in the resulting image
	// ie: CellSize of 20 will result in each colored cell to be 20x20 pixels
	CellSize int
}

type OptionsFunc func(*Options)

// WithSize applies the given size option
func WithSize(size int) OptionsFunc {
	return func(o *Options) {
		o.Size = size
	}
}

// WithCellSize applies the given cell size option
func WithCellSize(size int) OptionsFunc {
	return func(o *Options) {
		o.CellSize = size
	}
}

func CreateIdenticon(from string, options ...OptionsFunc) (image.Image, error) {
	h := sha512.Sum512([]byte(from))
	hxstr := hex.EncodeToString(h[:])

	// apply options
	opts := getDefaultOptions()
	for _, opt := range options {
		opt(opts)
	}

	if opts.Size > MaxSize || opts.Size <= 0 {
		return nil, errors.New("max size out of range")
	}

	// generate our grid and initalize all cells to 0
	grid := make([][]int, opts.Size)
	for i := range grid {
		grid[i] = make([]int, opts.Size)
	}

	// the first 3 bytes in the hex string will determine the color
	r, g, b, err := extractColor(hxstr[0:6])
	if err != nil {
		return nil, err
	}

	// start at the 6th byte
	hexIdx := 6
	for i := 0; i < opts.Size; i++ {
		for j := 0; j <= opts.Size/2; j++ {
			var bit int
			bt, err := strconv.ParseInt(hxstr[hexIdx:hexIdx+2], 16, 64)
			if err != nil {
				return nil, err
			}
			if bt%2 > 0 {
				bit = 1
			}
			// set the cell to on/off
			// also do it on the mirror side
			grid[i][j] = bit
			grid[i][opts.Size-j-1] = bit
			hexIdx += 2
		}
	}

	// finally create the image
	img := image.NewRGBA(image.Rect(0, 0, opts.Size*opts.CellSize, opts.Size*opts.CellSize))
	white := color.RGBA{255, 255, 255, 255}
	for i := 0; i < opts.Size; i++ {
		for j := 0; j < opts.Size; j++ {
			bit := grid[i][j]
			col := white
			if bit == 1 {
				col = color.RGBA{uint8(r), uint8(g), uint8(b), 255}
			}
			drawCell(img, col, j*opts.CellSize, i*opts.CellSize, opts.CellSize)
		}
	}

	return img, nil
}

func drawCell(img *image.RGBA, col color.Color, offsetX, offsetY, size int) {
	for x := offsetX; x < offsetX+size; x++ {
		for y := offsetY; y < offsetY+size; y++ {
			img.Set(x, y, col)
		}
	}
}

func extractColor(str string) (r, g, b int, err error) {
	h, s, v, err := extractHsv(str)
	if err != nil {
		return
	}
	return hsvToRGB(h, s, v)
}

// extractHsv extracts hue, saturation and value
// from the first 3 bytes of the given string
func extractHsv(str string) (float64, float64, float64, error) {
	if len(str) < 6 {
		return 0, 0, 0, errors.New("invalid input str")
	}

	hue, err1 := strconv.ParseInt(str[0:2], 16, 64)
	sat, err2 := strconv.ParseInt(str[2:4], 16, 64)
	val, err3 := strconv.ParseInt(str[4:6], 16, 64)
	if err := errors.Join(err1, err2, err3); err != nil {
		return 0, 0, 0, err
	}
	// normalize values
	h := (float64(hue) / 256.0) * 365       // 0°-360°
	s := ((float64(sat) / 256.0) * 55) + 45 // 45-100
	v := ((float64(val) / 256.0) * 35) + 45 // 45-80
	return h, s, v, nil
}

func hsvToRGB(h, s, v float64) (r, g, b int, err error) {
	if h > 360 || h < 0 || s < 0 || s > 100 || v < 0 || v > 100 {
		err = errors.New("invalid input for hue, saturation or value")
		return
	}

	// normalize sat and val to values between 0 and 1
	s = s / 100
	v = v / 100

	// apply conversion formula from
	// https://www.rapidtables.com/convert/color/hsv-to-rgb.html
	hi := h / 60
	c := v * s
	x := c * (1 - math.Abs(math.Mod(hi, 2)-1))
	m := v - c

	// h can be divided up into 5 sectors
	// 0°>= && < 60°
	// 60°>= && < 120°
	// etc. until 360°
	hSector := int(hi) % 6
	var ri, gi, bi float64
	switch hSector {
	case 0:
		ri, gi, bi = c, x, 0
	case 1:
		ri, gi, bi = x, c, 0
	case 2:
		ri, gi, bi = 0, c, x
	case 3:
		ri, gi, bi = 0, x, c
	case 4:
		ri, gi, bi = x, 0, c
	case 5:
		ri, gi, bi = c, 0, x
	}

	r, g, b = int(math.Round((ri+m)*255)), int(math.Round((gi+m)*255)), int(math.Round((bi+m)*255))
	return
}
