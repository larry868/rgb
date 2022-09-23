// Copyright 2022 by lolorenzo77. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

package rgb

import "fmt"

func ExampleColor_RGB() {
	blue := Color(0x0d6efdff)
	r, g, b := blue.RGB()
	fmt.Printf("blue R=%d G=%d B=%d\n", r, g, b)
	// Output: blue R=13 G=110 B=253
}

func ExampleColor_Lighten() {
	lightblue := MakeRGB(13, 110, 253).Lighten(0.5)
	r, g, b := lightblue.RGB()
	fmt.Printf("lightblue R=%d G=%d B=%d\n", r, g, b)
	// Output: lightblue R=134 G=182 B=254
}

func ExampleColor_Hexa() {
	blue := MakeRGB(13, 110, 253)
	transparentblue := blue.Opacify(0.8)
	fmt.Printf("blue=%s, transparent blue=%s\n", blue.Hexa(), transparentblue.Hexa())
	// Output: blue=#0D6EFDFF, transparent blue=#0D6EFDCC
}
