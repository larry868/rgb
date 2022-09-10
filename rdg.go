// Copyright 2022 by lolorenzo77. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

/*
Package rgb provides a RGBA color type, which is a uint32, with a set of methods to manipulate the color.

Color type is a 32bits color: 24bits for the RGB color itself and 8 bits for the alpha channel (transparency/opacity)
*/
package rgb

import (
	"fmt"
	"strconv"
	"strings"
)

// Color is a 32bits color: 24bits for the Color color itself and 8 bits for the alpha channel (transparency/opacity)
type Color uint32

// NewRGB create a 24bits color with 100% opacity
func NewRGB(R uint8, G uint8, B uint8) *Color {
	return NewRGBA(R, G, B, 255)
}

// NewRGBA create a 24bits RGB color with an Alpha channel for the opacity.
// Alpha varies from 0 (full transprency) to 255 (full opacity)
func NewRGBA(R uint8, G uint8, B uint8, A uint8) *Color {
	var c = new(Color)
	*c = Color(((uint32(R)<<8+uint32(G))<<8+uint32(B))<<8 + uint32(A))
	return c
}

// NewHexa create a 24bits color with 100% opacity.
//
// Accepted string formats:
//
//	the string can start with '#' but it s optional
//	'#A' means a color with R="0xAA", G="0xAA", B="0xAA", and a full opacity
//	'#ABC' means a color with R="0xAA", G="0xBB", B="0xCC", and a full opacity
//	'#ABCD' means a color with R="0xAA", G="0xBB", B="0xCC", and an opacity of "0xDD"
//	'#ABCDEF' means a color with R="0xAB", G="0xCD", B="0xEF", and a full opacity
//	'#ABCDEF88' means a color with R="0xAB", G="0xCD", B="0xEF", and an opacity of "0x88"
//
// return nil if unable to convert the string in a valid color
func NewHexa(hex string) *Color {
	l := len(hex)
	if l > 0 && string(hex[0]) == "#" {
		hex = hex[1:]
		l -= 1
	}
	switch l {
	case 1:
		hex = strings.Repeat(hex, 6) + "FF"
	case 3:
		hex = string(hex[0]) + string(hex[0]) + string(hex[1]) + string(hex[1]) + string(hex[2]) + string(hex[2]) + "FF"
	case 4:
		hex = string(hex[0]) + string(hex[0]) + string(hex[1]) + string(hex[1]) + string(hex[2]) + string(hex[2]) + string(hex[3]) + string(hex[3])
	case 6:
		hex = hex + "FF"
	case 8:
		break
	default:
		return nil
	}

	num, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return nil
	}
	c := Color(num)
	return &c
}

// RGB returns the R, G, B components of the color
func (col Color) RGB() (R uint8, G uint8, B uint8) {
	R = uint8(col >> 24)
	G = uint8((col >> 16) & 0x00FF)
	B = uint8((col >> 8) & 0x0000FF)
	return
}

// RGB returns the R, G, B components of the color and the value of the Alpha chanel
func (col Color) RGBA() (R uint8, G uint8, B uint8, A uint8) {
	R = uint8(col >> 24)
	G = uint8((col >> 16) & 0x00FF)
	B = uint8((col >> 8) & 0x0000FF)
	A = uint8(col & 0x000000FF)
	return
}

// Hexa returns a 8 length char hexadecimal value of the color, starting with #
func (col Color) Hexa() string {
	str := "#"
	str += fmt.Sprintf("%08X", uint32(col))
	return str
}

// Clone duplicates the color
func (col Color) Clone() *Color {
	return &col
}

// Create a ligher color, a tint.
// The factor values is between 0 and 1. Alpha chanel stays unchanged.
//
//	A factor >= 1 returns WHITE
//	A factor <= 0 returns an unchanged color
func (col Color) Lighten(factor float32) *Color {
	if factor < 0 {
		factor = 0
	}
	if factor > 1 {
		factor = 1
	}
	r, g, b, a := col.RGBA()
	rf := float32(r) + ((255 - float32(r)) * factor)
	gf := float32(g) + ((255 - float32(g)) * factor)
	bf := float32(b) + ((255 - float32(b)) * factor)
	c := NewRGBA(uint8(rf), uint8(gf), uint8(bf), a)
	return c
}

// Create a darker color, usefull to create a shade for example.
// The factor values is between 0 and 1. Alpha chanel stays unchanged.
//
//	A factor >= 1 returns an unchanged color
//	A factor <= 0 returns BLACK
func (col Color) Darken(factor float32) *Color {
	if factor < 0 {
		factor = 0
	}
	if factor > 1 {
		factor = 1
	}
	r, g, b, a := col.RGBA()
	rf := float32(r) * (1 - factor)
	gf := float32(g) * (1 - factor)
	bf := float32(b) * (1 - factor)
	c := NewRGBA(uint8(rf), uint8(gf), uint8(bf), a)
	return c
}

// Clone the color setting the alpha chanel with an opacity factor.
//
//		opacityfactor >= 1 set the alpha chanel to 255, fully opaque
//	 opacityfactor <= 0 set the alpha chanel to 0, fully transparent
func (col Color) Alpha(opacityfactor float32) *Color {
	r, g, b, _ := col.RGBA()
	aa := float32(255) * opacityfactor
	return NewRGBA(r, g, b, uint8(aa))
}

// GrayScale returns a grayscale of the color
// https://www.dynamsoft.com/blog/insights/image-processing/image-processing-101-color-space-conversion/
func (col Color) GrayScale() *Color {
	r, g, b := col.RGB()
	rs := float32(r) * 0.299
	rg := float32(g) * 0.587
	rb := float32(b) * 0.114
	return NewRGB(uint8(rs), uint8(rg), uint8(rb))
}
