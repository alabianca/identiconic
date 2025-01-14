package identiconic_test

import (
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/alabianca/identiconic"
)

func TestIntCreateIdenticon(t *testing.T) {
	img, err := identiconic.CreateIdenticon("a1e5577a-ce13-498e-8781-fb2f598992e3", identiconic.WithCellSize(25))
	if err != nil {
		t.Fatalf("expected no error, but got %s", err)
	}
	out, err := os.Create("test.png")
	if err != nil {
		t.Fatalf("expected no error, but got %s", err)
	}
	defer out.Close()
	if err := png.Encode(out, img); err != nil {
		t.Fatalf("expected no error, but got %s", err)
	}
}

func TestIntCreateIdenticonColorOpt(t *testing.T) {
	blue := color.RGBA{13, 117, 255, 255}
	img, err := identiconic.CreateIdenticon("177de69e-aca8-47cd-9643-5cf97727b781", identiconic.WithColor(blue))
	if err != nil {
		t.Fatalf("expected no error, but got %s", err)
	}
	out, err := os.Create("test_color.png")
	if err != nil {
		t.Fatalf("expected no error, but got %s", err)
	}
	defer out.Close()
	if err := png.Encode(out, img); err != nil {
		t.Fatalf("expected no error, but got %s", err)
	}
}
