package main

import (
	"bytes"
	"errors"
	"image/png"
	"testing"
)

var (
	testVersion = Version(1)
)

func TestVersionDeterminesSize(t *testing.T) {
	table := []struct {
		version  int
		expected int
	}{
		{1, 21},
		{2, 25},
		{6, 41},
	}

	for _, test := range table {
		buffer := new(bytes.Buffer)
		GenerateQRCode(buffer, qrcode, Version(test.version))
		img, _ := png.Decode(buffer)
		if width := img.Bounds().Dx(); width != test.expected {
			t.Errorf("Version %2d, expected %3d but got %3d",
				test.version, test.expected, width)
		}
	}
}

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, qrcode, testVersion)

	if buffer.Len() == 0 {
		t.Errorf("Your QR code is empty as space... like the real Space space, " +
			"not the space between those spaces. There's nothing in there.")

		_, err := png.Decode(buffer)

		if err != nil {
			t.Errorf("Generated QRCode is not a PNG: %s", err)
		}
	}
}

type ErrorWriter struct{}

func (e *ErrorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Expected error")
}

func TestGenerateQRCodePropogatesErrors(t *testing.T) {
	w := new(ErrorWriter)
	err := GenerateQRCode(w, qrcode, testVersion)

	if err == nil || err.Error() != "Expected error" {
		t.Errorf("error not propogated correctly, got %v", err)
	}

}
