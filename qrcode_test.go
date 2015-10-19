package main

import (
	"bytes"
	"image/png"
	"testing"
)

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "617-1234")

	if buffer.Len() == 0 {
		t.Errorf("Your QR code is empty as space... like the real Space space, " +
			"not the space between those spaces. There's nothing in there.")

		_, err := png.Decode(buffer)

		if err != nil {
			t.Errorf("Generated QRCode is not a PNG: %s", err)
		}
	}
}
