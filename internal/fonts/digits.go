// tinyfontgen --package fonts --fontname Digits digits.bdf --output internal/fonts/digits.go --all --verbose

package fonts

import (
	"tinygo.org/x/tinyfont"
)

var Digits = tinyfont.Font{
	Glyphs: []tinyfont.Glyph{
		/* 0 */ tinyfont.Glyph{Rune: 0x30, Width: 0xc, Height: 0x14, XAdvance: 0xf, XOffset: 3, YOffset: -20, Bitmaps: []uint8{0xff, 0xff, 0xff, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x3, 0xff, 0xff, 0xff}},
		/* 1 */ tinyfont.Glyph{Rune: 0x31, Width: 0x6, Height: 0x14, XAdvance: 0xb, XOffset: 5, YOffset: -20, Bitmaps: []uint8{0xff, 0xf0, 0xc3, 0xc, 0x30, 0xc3, 0xc, 0x30, 0xc3, 0xc, 0x30, 0xc3, 0xc, 0x30, 0xc3}},
		/* 2 */ tinyfont.Glyph{Rune: 0x32, Width: 0xb, Height: 0x14, XAdvance: 0xe, XOffset: 3, YOffset: -20, Bitmaps: []uint8{0xff, 0xff, 0xfc, 0x1, 0x80, 0x30, 0x6, 0x0, 0xc0, 0x18, 0x3, 0x0, 0x60, 0xf, 0xff, 0xff, 0xfc, 0x1, 0x80, 0x30, 0x6, 0x0, 0xc0, 0x18, 0x3, 0xff, 0xff, 0xf0}},
		/* 3 */ tinyfont.Glyph{Rune: 0x33, Width: 0xc, Height: 0x14, XAdvance: 0xf, XOffset: 3, YOffset: -20, Bitmaps: []uint8{0xff, 0xff, 0xff, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3, 0xff, 0xff, 0xff, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3, 0xff, 0xff, 0xff}},
		/* 4 */ tinyfont.Glyph{Rune: 0x34, Width: 0xf, Height: 0x14, XAdvance: 0xf, XOffset: 0, YOffset: -20, Bitmaps: []uint8{0x18, 0x6, 0x30, 0xc, 0x60, 0x18, 0xc0, 0x31, 0x80, 0x63, 0x0, 0xc6, 0x1, 0x8c, 0x3, 0x18, 0x6, 0x30, 0xc, 0x7f, 0xf8, 0xff, 0xf0, 0x0, 0x60, 0x0, 0xc0, 0x1, 0x80, 0x3, 0x0, 0x6, 0x0, 0xc, 0x0, 0x18, 0x0, 0x30}},
		/* 5 */ tinyfont.Glyph{Rune: 0x35, Width: 0xf, Height: 0x14, XAdvance: 0xf, XOffset: 0, YOffset: -20, Bitmaps: []uint8{0x1f, 0xfe, 0x3f, 0xfc, 0x60, 0x0, 0xc0, 0x1, 0x80, 0x3, 0x0, 0x6, 0x0, 0xc, 0x0, 0x1f, 0xfe, 0x3f, 0xfc, 0x0, 0x18, 0x0, 0x30, 0x0, 0x60, 0x0, 0xc0, 0x1, 0x80, 0x3, 0x0, 0x6, 0x0, 0xc, 0x7f, 0xf8, 0xff, 0xf0}},
		/* 6 */ tinyfont.Glyph{Rune: 0x36, Width: 0xc, Height: 0x14, XAdvance: 0xf, XOffset: 3, YOffset: -20, Bitmaps: []uint8{0xff, 0xff, 0xff, 0xc0, 0xc, 0x0, 0xc0, 0xc, 0x0, 0xc0, 0xc, 0x0, 0xff, 0xff, 0xff, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x3, 0xff, 0xff, 0xff}},
		/* 7 */ tinyfont.Glyph{Rune: 0x37, Width: 0xc, Height: 0x14, XAdvance: 0xf, XOffset: 3, YOffset: -20, Bitmaps: []uint8{0xff, 0xff, 0xff, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3}},
		/* 8 */ tinyfont.Glyph{Rune: 0x38, Width: 0xf, Height: 0x14, XAdvance: 0xf, XOffset: 0, YOffset: -20, Bitmaps: []uint8{0x1f, 0xfe, 0x3f, 0xfc, 0x60, 0x18, 0xc0, 0x31, 0x80, 0x63, 0x0, 0xc6, 0x1, 0x8c, 0x3, 0x1f, 0xfe, 0x3f, 0xfc, 0x60, 0x18, 0xc0, 0x31, 0x80, 0x63, 0x0, 0xc6, 0x1, 0x8c, 0x3, 0x18, 0x6, 0x30, 0xc, 0x7f, 0xf8, 0xff, 0xf0}},
		/* 9 */ tinyfont.Glyph{Rune: 0x39, Width: 0xc, Height: 0x14, XAdvance: 0xf, XOffset: 3, YOffset: -20, Bitmaps: []uint8{0xff, 0xff, 0xff, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x3, 0xff, 0xff, 0xff, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3, 0x0, 0x30, 0x3, 0xff, 0xff, 0xff}},
		/* : */ tinyfont.Glyph{Rune: 0x3a, Width: 0x2, Height: 0x8, XAdvance: 0xa, XOffset: 8, YOffset: -14, Bitmaps: []uint8{0xf0, 0xf}},
	},

	YAdvance: 0xb,
}
