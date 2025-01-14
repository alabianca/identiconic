package identiconic_test

import (
	"image/png"
	"os"
	"testing"

	"github.com/alabianca/identiconic"
)

func TestIntCreateIdenticon(t *testing.T) {
	img, err := identiconic.CreateIdenticon("177de69e-aca8-47cd-9643-5cf97727b781")
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
