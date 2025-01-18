<img src="https://github.com/alabianca/identiconic/blob/main/gopher_logo.png"/>

# Identiconic
Create [identicons](https://en.wikipedia.org/wiki/Identicon) just like GitHub.
You can use this package to create default avatars for your users or accounts.
The generated avatars will be unique to the given string (usually in the form of a UUID).
These identicons are symetrical to give them an avatar type of look.

## Usage
```go
// create the identicon from a UUID as an image.Image
img, err := identiconic.CreateIdenticon("177de69e-aca8-47cd-9643-5cf97727b781")
if err != nil {
    panic(err)
}
out, err := os.Create("out.png")
if err != nil {
    panic(err)
}
defer out.Close()
// write the image to a png file
if err := png.Encode(out, img); err != nil {
    panic(err)
}
```

### Where does the color come from
The color is computed based on the first 3 bytes of the input string. Though you can override this behavior by providing the `identiconic.WithColor(color)` option.
```go
blue := color.RGBA{13, 117, 255, 255}
img, err := identiconic.CreateIdenticon(
    "177de69e-aca8-47cd-9643-5cf97727b781", 
    identiconic.WithColor(blue),
    )
```

### Sizing
By default a grid of 10x10 cells is used to generate the identicon, but that can be overriden. 10 is the max. 
```go
// generate an identicon with a 7x7 grid
img, err := identiconic.CreateIdenticon(
    "177de69e-aca8-47cd-9643-5cf97727b781", 
    identiconic.WithSize(7),
    )
```

You can also change the default `CellSize` (defaults to 20) to create larger or smaller images
```go
// generate an identicon with a 7x7 grid
img, err := identiconic.CreateIdenticon(
    "177de69e-aca8-47cd-9643-5cf97727b781", 
    identiconic.WithCellSize(10),
    )
```

### Examples
<img src="https://github.com/alabianca/identiconic/blob/main/examples/example1.png"/>
<img src="https://github.com/alabianca/identiconic/blob/main/examples/example1_color.png"/>
<img src="https://github.com/alabianca/identiconic/blob/main/examples/example2.png"/>
<img src="https://github.com/alabianca/identiconic/blob/main/examples/example3.png"/>
<img src="https://github.com/alabianca/identiconic/blob/main/examples/example4.png"/>