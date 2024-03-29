package freesans

import (
	"tinygo.org/x/tinyfont"
)

var Oblique9pt7b = tinyfont.Font{
	BBox: [4]int8{20, 18, -1, -13},
	Glyphs: []tinyfont.Glyph{
		/*   */ tinyfont.Glyph{Rune: 32, Width: 0x0, Height: 0x0, XAdvance: 0x5, XOffset: 0, YOffset: 1, Bitmaps: []uint8{}},
		/* ! */ tinyfont.Glyph{Rune: 33, Width: 0x5, Height: 0xd, XAdvance: 0x5, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x10, 0x84, 0x22, 0x10, 0x84, 0x42, 0x10, 0x8, 0x0}},
		/* " */ tinyfont.Glyph{Rune: 34, Width: 0x5, Height: 0x4, XAdvance: 0x6, XOffset: 3, YOffset: -12, Bitmaps: []uint8{0xde, 0xe5, 0x20}},
		/* # */ tinyfont.Glyph{Rune: 35, Width: 0xb, Height: 0xd, XAdvance: 0xa, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x6, 0x40, 0x88, 0x13, 0x6, 0x43, 0xfe, 0x32, 0x4, 0x40, 0x98, 0x32, 0x1f, 0xf0, 0x98, 0x22, 0x4, 0xc0}},
		/* $ */ tinyfont.Glyph{Rune: 36, Width: 0xb, Height: 0x10, XAdvance: 0xa, XOffset: 1, YOffset: -13, Bitmaps: []uint8{0x2, 0x1, 0xf8, 0x6b, 0x99, 0x33, 0x40, 0x68, 0xf, 0x0, 0xf8, 0x7, 0xc1, 0x1b, 0x23, 0x64, 0x4e, 0x98, 0xfc, 0x4, 0x0, 0x80}},
		/* % */ tinyfont.Glyph{Rune: 37, Width: 0xf, Height: 0xd, XAdvance: 0x10, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x3c, 0x8, 0xcc, 0x23, 0x18, 0x86, 0x32, 0xc, 0x64, 0x19, 0x90, 0x1e, 0x40, 0x1, 0x1e, 0x2, 0x66, 0x9, 0x8c, 0x23, 0x18, 0x86, 0x62, 0x7, 0x80}},
		/* & */ tinyfont.Glyph{Rune: 38, Width: 0xa, Height: 0xd, XAdvance: 0xc, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0xf, 0x6, 0x63, 0x18, 0xc6, 0x3f, 0x7, 0x3, 0xc1, 0xb3, 0xc7, 0xb0, 0xcc, 0x33, 0x3e, 0x79, 0x80}},
		/* ' */ tinyfont.Glyph{Rune: 39, Width: 0x2, Height: 0x4, XAdvance: 0x3, XOffset: 3, YOffset: -12, Bitmaps: []uint8{0xfa}},
		/* ( */ tinyfont.Glyph{Rune: 40, Width: 0x7, Height: 0x11, XAdvance: 0x6, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x4, 0x10, 0x60, 0x83, 0x4, 0x18, 0x30, 0xc1, 0x83, 0x6, 0xc, 0x18, 0x10, 0x30, 0x20}},
		/* ) */ tinyfont.Glyph{Rune: 41, Width: 0x7, Height: 0x11, XAdvance: 0x6, XOffset: -1, YOffset: -12, Bitmaps: []uint8{0x8, 0x18, 0x10, 0x30, 0x60, 0xc1, 0x83, 0x6, 0x18, 0x30, 0x41, 0x82, 0xc, 0x10, 0x40}},
		/* * */ tinyfont.Glyph{Rune: 42, Width: 0x6, Height: 0x5, XAdvance: 0x7, XOffset: 3, YOffset: -12, Bitmaps: []uint8{0x19, 0x73, 0x16, 0x48}},
		/* + */ tinyfont.Glyph{Rune: 43, Width: 0x9, Height: 0x8, XAdvance: 0xb, XOffset: 2, YOffset: -7, Bitmaps: []uint8{0x4, 0x4, 0x2, 0x1f, 0xf0, 0x80, 0x80, 0x40, 0x20}},
		/* , */ tinyfont.Glyph{Rune: 44, Width: 0x3, Height: 0x5, XAdvance: 0x5, XOffset: 1, YOffset: -1, Bitmaps: []uint8{0x6d, 0x28}},
		/* - */ tinyfont.Glyph{Rune: 45, Width: 0x4, Height: 0x1, XAdvance: 0x6, XOffset: 2, YOffset: -4, Bitmaps: []uint8{0xf0}},
		/* . */ tinyfont.Glyph{Rune: 46, Width: 0x2, Height: 0x1, XAdvance: 0x5, XOffset: 2, YOffset: 0, Bitmaps: []uint8{0xc0}},
		/* / */ tinyfont.Glyph{Rune: 47, Width: 0x8, Height: 0xd, XAdvance: 0x5, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0x1, 0x2, 0x4, 0x4, 0x8, 0x8, 0x10, 0x10, 0x20, 0x20, 0x40, 0x40, 0x80}},
		/* 0 */ tinyfont.Glyph{Rune: 48, Width: 0x9, Height: 0xd, XAdvance: 0xa, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0xf, 0x19, 0xc8, 0x6c, 0x36, 0x1a, 0xf, 0x5, 0x86, 0xc3, 0x61, 0xb1, 0x9c, 0x87, 0x80}},
		/* 1 */ tinyfont.Glyph{Rune: 49, Width: 0x5, Height: 0xd, XAdvance: 0xa, XOffset: 4, YOffset: -12, Bitmaps: []uint8{0x8, 0xcd, 0xe3, 0x18, 0xc4, 0x23, 0x18, 0xc4, 0x0}},
		/* 2 */ tinyfont.Glyph{Rune: 50, Width: 0xb, Height: 0xd, XAdvance: 0xa, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x7, 0x83, 0x1c, 0x41, 0x98, 0x30, 0x6, 0x1, 0x80, 0x60, 0x38, 0x1c, 0x6, 0x1, 0x80, 0x20, 0xf, 0xf8}},
		/* 3 */ tinyfont.Glyph{Rune: 51, Width: 0xa, Height: 0xd, XAdvance: 0xa, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0xf, 0x86, 0x73, 0xc, 0x83, 0x0, 0xc0, 0x60, 0xe0, 0x6, 0x1, 0xb0, 0x6c, 0x13, 0x8c, 0x7c, 0x0}},
		/* 4 */ tinyfont.Glyph{Rune: 52, Width: 0x9, Height: 0xd, XAdvance: 0xa, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x0, 0x80, 0xc0, 0xe0, 0xa0, 0x90, 0x98, 0x8c, 0x86, 0xff, 0x81, 0x1, 0x80, 0xc0, 0x60}},
		/* 5 */ tinyfont.Glyph{Rune: 53, Width: 0xb, Height: 0xd, XAdvance: 0xa, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0xf, 0xc3, 0x0, 0x40, 0x8, 0x3, 0x0, 0x7f, 0x1c, 0x70, 0x6, 0x0, 0xc0, 0x1b, 0x6, 0x71, 0x87, 0xe0}},
		/* 6 */ tinyfont.Glyph{Rune: 54, Width: 0xa, Height: 0xd, XAdvance: 0xa, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0xf, 0x86, 0x73, 0xd, 0x80, 0x60, 0x1f, 0xcf, 0x3b, 0x86, 0xc1, 0xb0, 0x6c, 0x33, 0x98, 0x3c, 0x0}},
		/* 7 */ tinyfont.Glyph{Rune: 55, Width: 0xa, Height: 0xd, XAdvance: 0xa, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x7f, 0xc0, 0x20, 0x10, 0xc, 0x6, 0x1, 0x0, 0x80, 0x60, 0x10, 0xc, 0x2, 0x1, 0x80, 0x40, 0x0}},
		/* 8 */ tinyfont.Glyph{Rune: 56, Width: 0xa, Height: 0xd, XAdvance: 0xa, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0xf, 0x86, 0x73, 0xc, 0xc3, 0x30, 0xcc, 0x61, 0xe1, 0x86, 0x41, 0xb0, 0x6c, 0x13, 0x8c, 0x3e, 0x0}},
		/* 9 */ tinyfont.Glyph{Rune: 57, Width: 0xa, Height: 0xd, XAdvance: 0xa, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0xf, 0x6, 0x73, 0xd, 0x83, 0x60, 0xd8, 0x77, 0x3c, 0xfe, 0x1, 0x80, 0x6c, 0x33, 0x98, 0x7c, 0x0}},
		/* : */ tinyfont.Glyph{Rune: 58, Width: 0x4, Height: 0x9, XAdvance: 0x5, XOffset: 2, YOffset: -8, Bitmaps: []uint8{0x30, 0x0, 0x0, 0x0, 0xc0}},
		/* ; */ tinyfont.Glyph{Rune: 59, Width: 0x5, Height: 0xc, XAdvance: 0x5, XOffset: 1, YOffset: -8, Bitmaps: []uint8{0x18, 0x0, 0x0, 0x0, 0xc, 0x62, 0x11, 0x0}},
		/* < */ tinyfont.Glyph{Rune: 60, Width: 0x9, Height: 0x9, XAdvance: 0xb, XOffset: 2, YOffset: -8, Bitmaps: []uint8{0x0, 0x1, 0xc3, 0x8f, 0xc, 0x7, 0x0, 0xe0, 0x1e, 0x1, 0x0}},
		/* = */ tinyfont.Glyph{Rune: 61, Width: 0xa, Height: 0x4, XAdvance: 0xb, XOffset: 1, YOffset: -5, Bitmaps: []uint8{0x7f, 0xc0, 0x0, 0x3, 0xfe}},
		/* > */ tinyfont.Glyph{Rune: 62, Width: 0x9, Height: 0x9, XAdvance: 0xb, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x40, 0x3c, 0x3, 0x80, 0x70, 0x18, 0x78, 0xe1, 0xc0, 0x0, 0x0}},
		/* ? */ tinyfont.Glyph{Rune: 63, Width: 0x9, Height: 0xd, XAdvance: 0xa, XOffset: 3, YOffset: -12, Bitmaps: []uint8{0x1f, 0x30, 0xd0, 0x78, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x0, 0x0, 0x0, 0x6, 0x0}},
		/* @ */ tinyfont.Glyph{Rune: 64, Width: 0x12, Height: 0x10, XAdvance: 0x12, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x0, 0xfe, 0x0, 0xc0, 0xe0, 0xc0, 0x18, 0x61, 0xd3, 0x31, 0x9c, 0xd8, 0xc2, 0x36, 0x31, 0x8f, 0x18, 0x67, 0xc6, 0x11, 0xb1, 0x8c, 0xcc, 0x67, 0x63, 0xe, 0xf0, 0x60, 0x0, 0x1c, 0x0, 0x1, 0x81, 0x0, 0x1f, 0xc0}},
		/* A */ tinyfont.Glyph{Rune: 65, Width: 0xc, Height: 0xd, XAdvance: 0xc, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0x1, 0xc0, 0x1c, 0x3, 0xc0, 0x24, 0x6, 0x60, 0x46, 0xc, 0x61, 0x86, 0x1f, 0xe3, 0x6, 0x20, 0x26, 0x3, 0x40, 0x30}},
		/* B */ tinyfont.Glyph{Rune: 66, Width: 0xc, Height: 0xd, XAdvance: 0xc, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1f, 0xe1, 0x87, 0x30, 0x33, 0x3, 0x30, 0x23, 0x6, 0x3f, 0xc6, 0x6, 0x60, 0x66, 0x6, 0x60, 0x66, 0xc, 0x7f, 0x80}},
		/* C */ tinyfont.Glyph{Rune: 67, Width: 0xc, Height: 0xd, XAdvance: 0xd, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x7, 0xc1, 0x86, 0x30, 0x32, 0x3, 0x60, 0x4, 0x0, 0xc0, 0xc, 0x0, 0xc0, 0x6c, 0x6, 0xc0, 0xc6, 0x18, 0x3e, 0x0}},
		/* D */ tinyfont.Glyph{Rune: 68, Width: 0xd, Height: 0xd, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1f, 0xe0, 0xc1, 0x84, 0x6, 0x60, 0x33, 0x1, 0x98, 0xc, 0x80, 0x64, 0x2, 0x60, 0x33, 0x1, 0x98, 0x18, 0x81, 0x87, 0xf0, 0x0}},
		/* E */ tinyfont.Glyph{Rune: 69, Width: 0xc, Height: 0xd, XAdvance: 0xc, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1f, 0xf1, 0x80, 0x10, 0x3, 0x0, 0x30, 0x3, 0x0, 0x3f, 0xe2, 0x0, 0x60, 0x6, 0x0, 0x60, 0x4, 0x0, 0x7f, 0xc0}},
		/* F */ tinyfont.Glyph{Rune: 70, Width: 0xc, Height: 0xd, XAdvance: 0xb, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1f, 0xf1, 0x80, 0x10, 0x3, 0x0, 0x30, 0x3, 0x0, 0x3f, 0xc2, 0x0, 0x60, 0x6, 0x0, 0x60, 0x4, 0x0, 0x40, 0x0}},
		/* G */ tinyfont.Glyph{Rune: 71, Width: 0xd, Height: 0xd, XAdvance: 0xe, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x7, 0xe0, 0xe1, 0x8c, 0x6, 0xc0, 0x36, 0x0, 0x60, 0x3, 0x7, 0xf8, 0x2, 0xc0, 0x36, 0x1, 0x98, 0x1c, 0xe1, 0xc1, 0xf2, 0x0}},
		/* H */ tinyfont.Glyph{Rune: 72, Width: 0xd, Height: 0xd, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x18, 0x8, 0xc0, 0xc4, 0x6, 0x60, 0x33, 0x1, 0x18, 0x18, 0xff, 0xc4, 0x6, 0x60, 0x23, 0x1, 0x18, 0x18, 0x80, 0xc4, 0x6, 0x0}},
		/* I */ tinyfont.Glyph{Rune: 73, Width: 0x4, Height: 0xd, XAdvance: 0x5, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x33, 0x32, 0x26, 0x66, 0x44, 0xcc, 0xc0}},
		/* J */ tinyfont.Glyph{Rune: 74, Width: 0xa, Height: 0xd, XAdvance: 0x9, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x0, 0xc0, 0x60, 0x18, 0x6, 0x1, 0x80, 0x60, 0x30, 0xc, 0x3, 0x30, 0xcc, 0x63, 0x18, 0x7c, 0x0}},
		/* K */ tinyfont.Glyph{Rune: 75, Width: 0xe, Height: 0xd, XAdvance: 0xc, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x18, 0x18, 0x60, 0xc1, 0xe, 0xc, 0x60, 0x33, 0x0, 0xd8, 0x3, 0xf0, 0xc, 0xc0, 0x61, 0x81, 0x86, 0x6, 0xc, 0x10, 0x30, 0x40, 0x60}},
		/* L */ tinyfont.Glyph{Rune: 76, Width: 0x9, Height: 0xd, XAdvance: 0xa, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x18, 0xc, 0x4, 0x6, 0x3, 0x1, 0x80, 0xc0, 0x40, 0x60, 0x30, 0x18, 0x8, 0x7, 0xf8}},
		/* M */ tinyfont.Glyph{Rune: 77, Width: 0x10, Height: 0xd, XAdvance: 0xf, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x18, 0x6, 0x18, 0xe, 0x18, 0xe, 0x34, 0x1e, 0x34, 0x36, 0x34, 0x34, 0x24, 0x64, 0x24, 0x6c, 0x64, 0xcc, 0x64, 0x8c, 0x65, 0x88, 0x43, 0x8, 0x43, 0x18}},
		/* N */ tinyfont.Glyph{Rune: 78, Width: 0xd, Height: 0xd, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x18, 0x8, 0xe0, 0x47, 0x6, 0x6c, 0x33, 0x61, 0x99, 0x8, 0x8c, 0xc4, 0x66, 0x61, 0xb3, 0xd, 0x18, 0x38, 0x81, 0xc4, 0x6, 0x0}},
		/* O */ tinyfont.Glyph{Rune: 79, Width: 0xd, Height: 0xd, XAdvance: 0xe, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x7, 0xc0, 0xc3, 0x8c, 0xe, 0xc0, 0x36, 0x1, 0xe0, 0xf, 0x0, 0x78, 0x3, 0xc0, 0x36, 0x1, 0xb8, 0x18, 0xe1, 0x81, 0xf0, 0x0}},
		/* P */ tinyfont.Glyph{Rune: 80, Width: 0xc, Height: 0xd, XAdvance: 0xc, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1f, 0xe1, 0x83, 0x10, 0x33, 0x3, 0x30, 0x33, 0x6, 0x3f, 0xc2, 0x0, 0x60, 0x6, 0x0, 0x60, 0x4, 0x0, 0x40, 0x0}},
		/* Q */ tinyfont.Glyph{Rune: 81, Width: 0xd, Height: 0xe, XAdvance: 0xe, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x7, 0xc0, 0xc3, 0x8c, 0xe, 0xc0, 0x36, 0x1, 0xe0, 0xf, 0x0, 0x78, 0x3, 0xc0, 0x36, 0x9, 0xb8, 0x78, 0xe3, 0x81, 0xf6, 0x0, 0x10}},
		/* R */ tinyfont.Glyph{Rune: 82, Width: 0xd, Height: 0xd, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1f, 0xf0, 0xc0, 0xc4, 0x6, 0x60, 0x33, 0x1, 0x18, 0x18, 0xff, 0x4, 0xc, 0x60, 0x63, 0x3, 0x18, 0x18, 0x80, 0xc4, 0x6, 0x0}},
		/* S */ tinyfont.Glyph{Rune: 83, Width: 0xc, Height: 0xd, XAdvance: 0xc, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x7, 0xc1, 0x87, 0x30, 0x33, 0x3, 0x30, 0x3, 0xc0, 0xf, 0xc0, 0x1e, 0x0, 0x6c, 0x6, 0xc0, 0x46, 0xc, 0x3f, 0x0}},
		/* T */ tinyfont.Glyph{Rune: 84, Width: 0xa, Height: 0xd, XAdvance: 0xb, XOffset: 3, YOffset: -12, Bitmaps: []uint8{0xff, 0xc3, 0x0, 0xc0, 0x20, 0x18, 0x6, 0x1, 0x80, 0x60, 0x10, 0xc, 0x3, 0x0, 0xc0, 0x20, 0x0}},
		/* U */ tinyfont.Glyph{Rune: 85, Width: 0xc, Height: 0xd, XAdvance: 0xd, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x30, 0x13, 0x3, 0x20, 0x36, 0x3, 0x60, 0x26, 0x6, 0x60, 0x64, 0x6, 0xc0, 0x6c, 0x4, 0xc0, 0xce, 0x18, 0x3e, 0x0}},
		/* V */ tinyfont.Glyph{Rune: 86, Width: 0xb, Height: 0xd, XAdvance: 0xc, XOffset: 3, YOffset: -12, Bitmaps: []uint8{0xc0, 0x78, 0xb, 0x3, 0x20, 0xc4, 0x18, 0xc6, 0x18, 0x83, 0x30, 0x64, 0xd, 0x80, 0xa0, 0x1c, 0x3, 0x0}},
		/* W */ tinyfont.Glyph{Rune: 87, Width: 0x10, Height: 0xd, XAdvance: 0x11, XOffset: 3, YOffset: -12, Bitmaps: []uint8{0xc1, 0x83, 0xc1, 0x83, 0xc3, 0x86, 0xc2, 0x86, 0xc6, 0x84, 0xc4, 0x8c, 0xcc, 0xc8, 0xc8, 0xd8, 0xd8, 0xd0, 0xd0, 0xf0, 0x70, 0xe0, 0x60, 0xe0, 0x60, 0xe0}},
		/* X */ tinyfont.Glyph{Rune: 88, Width: 0xe, Height: 0xd, XAdvance: 0xc, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0xc, 0xc, 0x30, 0x60, 0x63, 0x1, 0x98, 0x2, 0xc0, 0xe, 0x0, 0x38, 0x1, 0xe0, 0xc, 0x80, 0x33, 0x1, 0x8c, 0xc, 0x18, 0x60, 0x60}},
		/* Y */ tinyfont.Glyph{Rune: 89, Width: 0xc, Height: 0xd, XAdvance: 0xc, XOffset: 3, YOffset: -12, Bitmaps: []uint8{0xc0, 0x66, 0xc, 0x60, 0xc2, 0x18, 0x33, 0x3, 0x60, 0x1c, 0x1, 0x80, 0x18, 0x1, 0x80, 0x18, 0x1, 0x0, 0x30, 0x0}},
		/* Z */ tinyfont.Glyph{Rune: 90, Width: 0xc, Height: 0xd, XAdvance: 0xb, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1f, 0xf0, 0x7, 0x0, 0xe0, 0xc, 0x1, 0x80, 0x30, 0x6, 0x0, 0xc0, 0x18, 0x3, 0x0, 0x60, 0xc, 0x0, 0xff, 0xc0}},
		/* [ */ tinyfont.Glyph{Rune: 91, Width: 0x7, Height: 0x11, XAdvance: 0x5, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0xe, 0x10, 0x20, 0x41, 0x2, 0x4, 0x8, 0x20, 0x40, 0x81, 0x4, 0x8, 0x10, 0x20, 0xe0}},
		/* \ */ tinyfont.Glyph{Rune: 92, Width: 0x2, Height: 0xd, XAdvance: 0x5, XOffset: 3, YOffset: -12, Bitmaps: []uint8{0xaa, 0xa9, 0x55, 0x40}},
		/* ] */ tinyfont.Glyph{Rune: 93, Width: 0x7, Height: 0x11, XAdvance: 0x5, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0xe, 0x8, 0x10, 0x20, 0x41, 0x2, 0x4, 0x8, 0x20, 0x40, 0x81, 0x4, 0x8, 0x10, 0xe0}},
		/* ^ */ tinyfont.Glyph{Rune: 94, Width: 0x7, Height: 0x7, XAdvance: 0x8, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0xc, 0x18, 0x51, 0xa2, 0x4c, 0x50, 0x80}},
		/* _ */ tinyfont.Glyph{Rune: 95, Width: 0xb, Height: 0x1, XAdvance: 0xa, XOffset: -1, YOffset: 3, Bitmaps: []uint8{0xff, 0xe0}},
		/* ` */ tinyfont.Glyph{Rune: 96, Width: 0x3, Height: 0x3, XAdvance: 0x6, XOffset: 3, YOffset: -12, Bitmaps: []uint8{0xc8, 0x80}},
		/* a */ tinyfont.Glyph{Rune: 97, Width: 0xa, Height: 0xa, XAdvance: 0xa, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xf, 0x86, 0x33, 0xc, 0x3, 0x3, 0xdf, 0xee, 0xb, 0x2, 0xc1, 0x9f, 0xe0}},
		/* b */ tinyfont.Glyph{Rune: 98, Width: 0xa, Height: 0xd, XAdvance: 0xa, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x10, 0x4, 0x1, 0x0, 0xdc, 0x39, 0x88, 0x32, 0xd, 0x83, 0x40, 0xd0, 0x64, 0x1b, 0x8c, 0xbc, 0x0}},
		/* c */ tinyfont.Glyph{Rune: 99, Width: 0x9, Height: 0xa, XAdvance: 0x9, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x1f, 0x18, 0xd8, 0x6c, 0xc, 0x6, 0x3, 0x1, 0x86, 0x66, 0x3e, 0x0}},
		/* d */ tinyfont.Glyph{Rune: 100, Width: 0xb, Height: 0xd, XAdvance: 0xa, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x0, 0x20, 0x8, 0x1, 0xf, 0x23, 0x14, 0xc1, 0x18, 0x26, 0x4, 0xc0, 0x98, 0x23, 0x4, 0x71, 0x87, 0xd0}},
		/* e */ tinyfont.Glyph{Rune: 101, Width: 0xa, Height: 0xa, XAdvance: 0xa, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xf, 0xc, 0x76, 0xd, 0x83, 0xff, 0xf0, 0xc, 0x3, 0x6, 0x63, 0xf, 0x80}},
		/* f */ tinyfont.Glyph{Rune: 102, Width: 0x6, Height: 0xd, XAdvance: 0x5, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1c, 0xc2, 0x1e, 0x20, 0x84, 0x10, 0x41, 0x4, 0x20, 0x80}},
		/* g */ tinyfont.Glyph{Rune: 103, Width: 0xa, Height: 0xe, XAdvance: 0xa, XOffset: 0, YOffset: -9, Bitmaps: []uint8{0xf, 0x46, 0x33, 0xc, 0xc1, 0x60, 0xd8, 0x26, 0x9, 0x86, 0x71, 0x8f, 0xe0, 0x10, 0x4, 0xc2, 0x1f, 0x0}},
		/* h */ tinyfont.Glyph{Rune: 104, Width: 0xa, Height: 0xd, XAdvance: 0xa, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x10, 0x4, 0x1, 0x0, 0x9f, 0x39, 0x88, 0x22, 0x9, 0x2, 0x40, 0x90, 0x44, 0x12, 0x4, 0x81, 0x0}},
		/* i */ tinyfont.Glyph{Rune: 105, Width: 0x4, Height: 0xd, XAdvance: 0x4, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x10, 0x2, 0x22, 0x64, 0x44, 0x48, 0x80}},
		/* j */ tinyfont.Glyph{Rune: 106, Width: 0x6, Height: 0x11, XAdvance: 0x4, XOffset: -1, YOffset: -12, Bitmaps: []uint8{0x4, 0x0, 0x1, 0x8, 0x20, 0x82, 0x8, 0x41, 0x4, 0x10, 0x42, 0x8, 0xe0}},
		/* k */ tinyfont.Glyph{Rune: 107, Width: 0x9, Height: 0xd, XAdvance: 0x9, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x10, 0x8, 0x4, 0x4, 0x32, 0x31, 0x20, 0xa0, 0xb8, 0x6c, 0x22, 0x11, 0x90, 0xc8, 0x30}},
		/* l */ tinyfont.Glyph{Rune: 108, Width: 0x4, Height: 0xd, XAdvance: 0x4, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x11, 0x22, 0x22, 0x64, 0x44, 0x48, 0x80}},
		/* m */ tinyfont.Glyph{Rune: 109, Width: 0xf, Height: 0xa, XAdvance: 0xf, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x2f, 0x3c, 0x63, 0x8c, 0x86, 0x19, 0x8, 0x44, 0x10, 0x88, 0x21, 0x10, 0x82, 0x21, 0x4, 0x82, 0x11, 0x4, 0x20}},
		/* n */ tinyfont.Glyph{Rune: 110, Width: 0xa, Height: 0xb, XAdvance: 0xa, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0x0, 0xb, 0xf3, 0x18, 0x82, 0x20, 0x90, 0x24, 0x9, 0x4, 0x41, 0x20, 0x48, 0x10}},
		/* o */ tinyfont.Glyph{Rune: 111, Width: 0xa, Height: 0xa, XAdvance: 0xa, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xf, 0xc, 0x76, 0xd, 0x83, 0xc0, 0xf0, 0x3c, 0x1b, 0x6, 0xe3, 0xf, 0x0}},
		/* p */ tinyfont.Glyph{Rune: 112, Width: 0xb, Height: 0xe, XAdvance: 0xa, XOffset: 0, YOffset: -9, Bitmaps: []uint8{0x17, 0xc3, 0x1c, 0x41, 0x98, 0x32, 0x6, 0x40, 0xc8, 0x33, 0x6, 0x71, 0x8b, 0xc1, 0x0, 0x20, 0x8, 0x1, 0x0, 0x0}},
		/* q */ tinyfont.Glyph{Rune: 113, Width: 0xa, Height: 0xe, XAdvance: 0xa, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x1e, 0xcc, 0x66, 0x9, 0x82, 0xc0, 0xb0, 0x4c, 0x13, 0x4, 0x63, 0xf, 0xc0, 0x20, 0x8, 0x2, 0x0, 0x80}},
		/* r */ tinyfont.Glyph{Rune: 114, Width: 0x7, Height: 0xa, XAdvance: 0x6, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x2c, 0x60, 0x81, 0x4, 0x8, 0x10, 0x20, 0x81, 0x0}},
		/* s */ tinyfont.Glyph{Rune: 115, Width: 0x8, Height: 0xa, XAdvance: 0x9, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x1e, 0x33, 0x63, 0x60, 0x70, 0x1e, 0x3, 0xc3, 0xc6, 0x7c}},
		/* t */ tinyfont.Glyph{Rune: 116, Width: 0x4, Height: 0xc, XAdvance: 0x5, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0x22, 0xf2, 0x44, 0x44, 0xcc, 0xce}},
		/* u */ tinyfont.Glyph{Rune: 117, Width: 0x9, Height: 0xa, XAdvance: 0xa, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0x21, 0x20, 0x90, 0x48, 0x24, 0x12, 0x13, 0x9, 0x84, 0xe6, 0x3e, 0x0}},
		/* v */ tinyfont.Glyph{Rune: 118, Width: 0x9, Height: 0xa, XAdvance: 0x9, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0xc1, 0xe1, 0xb0, 0xc8, 0xc4, 0x43, 0x61, 0xa0, 0xf0, 0x70, 0x18, 0x0}},
		/* w */ tinyfont.Glyph{Rune: 119, Width: 0xd, Height: 0xa, XAdvance: 0xd, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0xc7, 0x1e, 0x38, 0xb3, 0xcd, 0x96, 0x4c, 0xb6, 0x6d, 0xb1, 0x4d, 0xe, 0x78, 0x63, 0x83, 0x1c, 0x0}},
		/* x */ tinyfont.Glyph{Rune: 120, Width: 0xb, Height: 0xa, XAdvance: 0x9, XOffset: 0, YOffset: -9, Bitmaps: []uint8{0x10, 0xc3, 0x10, 0x24, 0x7, 0x80, 0xe0, 0x1c, 0x7, 0x81, 0x90, 0x23, 0x8, 0x20}},
		/* y */ tinyfont.Glyph{Rune: 121, Width: 0xb, Height: 0xe, XAdvance: 0x9, XOffset: 0, YOffset: -9, Bitmaps: []uint8{0x30, 0x46, 0x18, 0x42, 0x8, 0xc1, 0x10, 0x24, 0x7, 0x80, 0xe0, 0x1c, 0x3, 0x0, 0x60, 0x8, 0x3, 0x1, 0xc0, 0x0}},
		/* z */ tinyfont.Glyph{Rune: 122, Width: 0x9, Height: 0xa, XAdvance: 0x9, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x3f, 0x80, 0x80, 0x80, 0xc0, 0xc0, 0xc0, 0xc0, 0xc0, 0xc0, 0x7f, 0x0}},
		/* { */ tinyfont.Glyph{Rune: 123, Width: 0x5, Height: 0x11, XAdvance: 0x6, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x18, 0x88, 0x42, 0x10, 0x88, 0xc3, 0x18, 0x88, 0x42, 0x18, 0xe0}},
		/* | */ tinyfont.Glyph{Rune: 124, Width: 0x4, Height: 0x11, XAdvance: 0x5, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x11, 0x22, 0x22, 0x24, 0x44, 0x4c, 0x88, 0x88, 0x0}},
		/* } */ tinyfont.Glyph{Rune: 125, Width: 0x5, Height: 0x11, XAdvance: 0x6, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0x38, 0xc2, 0x10, 0x88, 0xc6, 0x18, 0x88, 0x42, 0x10, 0x88, 0xc0}},
		/* ~ */ tinyfont.Glyph{Rune: 126, Width: 0x9, Height: 0x3, XAdvance: 0xb, XOffset: 2, YOffset: -7, Bitmaps: []uint8{0x70, 0x4e, 0x41, 0xc0}},
	},

	YAdvance: 0x16,
}
