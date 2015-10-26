package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
)

var (
	qrcode = "617-1234"
)

// Version is a parameter of all qrcodes that meet specific
// size expectations
type Version int8

// GenerateQRCode creates a qrcode
func GenerateQRCode(w io.Writer, code string, version Version) error {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	return png.Encode(w, img)
}

func main() {
	fmt.Println("Hello QR Code.")

	file, _ := os.Create("qrcode.png")
	defer file.Close()

	GenerateQRCode(file, qrcode, Version(1))
}
