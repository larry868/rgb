// Copyright 2022 by lolorenzo77. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

/*
Package rgb provides a RGBA color type, which is a uint32, with a set of methods to manipulate the color.

Color type is a 32bits color: 24bits for the RGB color itself and 8 bits for the alpha channel (transparency/opacity)
*/
package rgb

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Basic RGB Colors
const (
	Red     = Color(0xff0000ff)
	Green   = Color(0x00ff00ff)
	Blue    = Color(0x0000ffff)
	Yellow  = Color(0xffff00ff)
	Cyan    = Color(0x00ffffff)
	Magenta = Color(0xff00ffff)
	Black   = Color(0x000000ff)
	Silver  = Color(0xc0c0c0ff)
	Gray    = Color(0x808080ff)
	White   = Color(0xffffffff)
	None    = Color(0x00000000)
)

// Color is a 32bits color: 24bits for the Color color itself and 8 bits for the alpha channel (transparency/opacity)
type Color uint32

// MakeRGB create a 24bits color with 100% opacity
func MakeRGB(R uint8, G uint8, B uint8) Color {
	return MakeRGBA(R, G, B, 255)
}

// MakeRGBA create a 24bits RGB color with an Alpha channel for the opacity.
// Alpha varies from 0 (full transprency) to 255 (full opacity)
func MakeRGBA(R uint8, G uint8, B uint8, A uint8) Color {
	return Color(((uint32(R)<<8+uint32(G))<<8+uint32(B))<<8 + uint32(A))
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
func (c Color) RGB() (R uint8, G uint8, B uint8) {
	R = uint8(c >> 24)
	G = uint8((c >> 16) & 0x00FF)
	B = uint8((c >> 8) & 0x0000FF)
	return
}

// RGB returns the R, G, B components of the color and the value of the Alpha chanel
func (c Color) RGBA() (R uint8, G uint8, B uint8, A uint8) {
	R = uint8(c >> 24)
	G = uint8((c >> 16) & 0x00FF)
	B = uint8((c >> 8) & 0x0000FF)
	A = uint8(c & 0x000000FF)
	return
}

// Hexa returns a 8 length char hexadecimal value of the color, starting with #
func (c Color) Hexa() string {
	str := "#"
	str += fmt.Sprintf("%08X", uint32(c))
	return str
}

// Returns the Red component of the color
func (c Color) Red() uint8 {
	return uint8(c >> 24)
}

// Returns the Green component of the color
func (c Color) Green() uint8 {
	return uint8((c >> 16) & 0x00FF)
}

// Returns the Blue component of the color
func (c Color) Blue() uint8 {
	return uint8((c >> 8) & 0x0000FF)
}

// Returns the Alpha component of the color
func (c Color) Alpha() uint8 {
	return uint8(c & 0x000000FF)
}

// Create a ligher color, a tint.
// The factor values is between 0 and 1. Alpha chanel stays unchanged.
//
//	A factor >= 1 returns WHITE
//	A factor <= 0 returns an unchanged color
func (c Color) Lighten(factor float32) Color {
	if factor < 0 {
		factor = 0
		log.Printf("Color.Lighten: factor out of range: %v\n", factor)
	}
	if factor > 1 {
		factor = 1
		log.Printf("Color.Lighten: factor out of range: %v\n", factor)
	}
	r, g, b, a := c.RGBA()
	rf := float32(r) + ((255 - float32(r)) * factor)
	gf := float32(g) + ((255 - float32(g)) * factor)
	bf := float32(b) + ((255 - float32(b)) * factor)
	return MakeRGBA(uint8(rf), uint8(gf), uint8(bf), a)
}

// Create a darker color, usefull to create a shade for example.
// The factor values is between 0 and 1. Alpha chanel stays unchanged.
//
//	A factor >= 1 returns an unchanged color
//	A factor <= 0 returns BLACK
func (c Color) Darken(factor float32) Color {
	if factor < 0 {
		factor = 0
		log.Printf("Color.Darken: factor out of range: %v\n", factor)
	}
	if factor > 1 {
		factor = 1
		log.Printf("Color.Darken: factor out of range: %v\n", factor)
	}
	r, g, b, a := c.RGBA()
	rf := float32(r) * (1 - factor)
	gf := float32(g) * (1 - factor)
	bf := float32(b) * (1 - factor)
	return MakeRGBA(uint8(rf), uint8(gf), uint8(bf), a)
}

// Clone the color setting the alpha chanel with an opacity factor.
//
//	opacityfactor >= 1 set the alpha chanel to 255, fully opaque
//	opacityfactor <= 0 set the alpha chanel to 0, fully transparent
func (c Color) Opacify(opacityfactor float32) Color {
	if opacityfactor < 0.0 || opacityfactor > 255.0 {
		log.Printf("Color.Alpha: opacityfactor out of range: %v\n", opacityfactor)
	}
	r, g, b, _ := c.RGBA()
	aa := float32(255) * opacityfactor
	return MakeRGBA(r, g, b, uint8(aa))
}

// GrayScale color of c
//
// https://www.dynamsoft.com/blog/insights/image-processing/image-processing-101-color-space-conversion/
func (c Color) GrayScale() Color {
	r, g, b, a := c.RGBA()
	rs := float32(r) * 0.299
	rg := float32(g) * 0.587
	rb := float32(b) * 0.114
	return MakeRGBA(uint8(rs), uint8(rg), uint8(rb), a)
}
