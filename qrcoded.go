package main

import (
	"fmt"
  "io"
  "os"
  "image"
  "image/png"
)

func GenerateQRCode(w io.Writer, code string) {
 	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
  _ = png.Encode(w, img)
 }

func main() {
	fmt.Println("Hello QR Code.")

  file, _ := os.Create("qrcode.png")
  defer file.Close()

	GenerateQRCode(file, "617-1234")
}
