# RGB color package in go

Package RGB provides a RGBA color type, which is a uint32, with a set of methods to manipulate the color.

RGB type is a 32bits color: 24bits for the RGB color itself and 8 bits for the alpha channel (transparency/opacity)

[![Go Reference](https://pkg.go.dev/badge/github.com/larry868/rgb/v2.svg)](https://pkg.go.dev/github.com/larry868/rgb/v2)

## Usage

There is 4 ways to initialize a new RGBA color

```go
// 1st way: directly with its 32bits exa value
blue1 := rgb.Color(0x0d6efdff)

// 2nd way: calling the NewRGB factory
blue2 := rgb.MakeRGB(13,110,253)

// 3rd way: calling the NewRGBA factory
blue3 := rgb.MakeRGBA(13,110,253,255)

// 4th way: calling the NewHexa factory
blue4 := rgb.ParseHexa("#0d6efdff")
if blue4 == nil {
    return errors.New("fail parsing the hexa color value")
}
```	

When you call the NewHexa factory, possible string formats are:

- '#A' means a color with R="0xAA", G="0xAA", B="0xAA", and a full opacity
- '#ABC' means a color with R="0xAA", G="0xBB", B="0xCC", and a full opacity
- '#ABCD' means a color with R="0xAA", G="0xBB", B="0xCC", and an opacity of "0xDD"
- '#ABCDEF' means a color with R="0xAB", G="0xCD", B="0xEF", and a full opacity
- '#ABCDEF88' means a color with R="0xAB", G="0xCD", B="0xEF", and an opacity of "0x88"

Then you can darken, lighten, convert to a grayscale, extracts the RGBA components or get the Hexa string representation of the color.

```go
    str := rgb.NewRGB(10,10,10).Darken(0.8).Hexa()
```

Or if you which you can use predefined boostrap const color
```go
    str := bootstrapcolor.Pink.Clone().Darken(0.8).Hexa()
```

## Installing 

```bash
go get -u github.com/larry868/rgb
```

## Change Log

- v1.1.0: migration to larry868 and go 1.23
- v1.0.0: initial version 

## Licence

[MIT](LICENSE)
