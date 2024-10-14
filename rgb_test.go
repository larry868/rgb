// Copyright 2022-2024 by larry868. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

package rgb

import "testing"

func TestRGB(t *testing.T) {

	var r, g, b, a uint8
	var str string

	col1 := MakeRGB(10, 100, 255)
	r, g, b = col1.RGB()
	if r != 10 || g != 100 || b != 255 {
		t.Errorf("RGB fails")
	}
	_, _, _, a = col1.RGBA()
	if a != 255 {
		t.Errorf("RGBA fails")
	}
	str = col1.Hexa()
	if str != "#0A64FFFF" {
		t.Errorf("String fails")
	}

	col2 := MakeRGBA(10, 100, 255, 100)
	r, g, b = col2.RGB()
	if r != 10 || g != 100 || b != 255 {
		t.Errorf("RGB fails")
	}
	_, _, _, a = col2.RGBA()
	if a != 100 {
		t.Errorf("RGBA fail")
	}
	str = col2.Hexa()
	if str != "#0A64FF64" {
		t.Errorf("String fails")
	}

	testset := []string{"#A", "#ABC", "#ABCD", "#ABCDEF", "#ABCDEF09"}
	testwant := []string{"#AAAAAAFF", "#AABBCCFF", "#AABBCCDD", "#ABCDEFFF", "#ABCDEF09"}
	for i, val := range testset {
		col := ParseHexa(val)
		if col == nil || col.Hexa() != testwant[i] {
			t.Errorf("Newhexa fails %q", val)
		}
	}

}
