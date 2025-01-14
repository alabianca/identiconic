package identiconic

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"testing"
)

func TestCreateIdenticon(t *testing.T) {
	_, err := CreateIdenticon("177de69e-aca8-47cd-9643-5cf97727b781")
	if err != nil {
		t.Fatalf("expected error nil, but got %s", err)
	}
}

func TestExtractHsv(t *testing.T) {
	h := sha512.Sum512([]byte("177de69e-aca8-47cd-9643-5cf97727b781"))
	str := hex.EncodeToString(h[:])
	extractHsv(str)
}

func TestHsvToRGB(t *testing.T) {
	cases := []struct {
		hue     float64
		sat     float64
		val     float64
		wantR   int
		wantG   int
		wantB   int
		wantErr error
	}{
		{
			hue:   60.0,
			sat:   45.0,
			val:   60.0,
			wantR: 153,
			wantG: 153,
			wantB: 84,
		},
		{
			hue:   35.0,
			sat:   85.0,
			val:   20.0,
			wantR: 51,
			wantG: 33,
			wantB: 8,
		},
		{
			hue:   150.0,
			sat:   85.0,
			val:   20.0,
			wantR: 8,
			wantG: 51,
			wantB: 29,
		},
		{
			hue:   190.0,
			sat:   85.0,
			val:   20.0,
			wantR: 8,
			wantG: 44,
			wantB: 51,
		},
		{
			hue:   250.0,
			sat:   85.0,
			val:   20.0,
			wantR: 15,
			wantG: 8,
			wantB: 51,
		},
		{
			hue:   310.0,
			sat:   85.0,
			val:   20.0,
			wantR: 51,
			wantG: 8,
			wantB: 44,
		},
		{
			hue:   360.0,
			sat:   85.0,
			val:   20.0,
			wantR: 51,
			wantG: 8,
			wantB: 8,
		},
		{
			hue:   0.0,
			sat:   85.0,
			val:   20.0,
			wantR: 51,
			wantG: 8,
			wantB: 8,
		},
		{
			hue:   0.0,
			sat:   0.0,
			val:   0.0,
			wantR: 0,
			wantG: 0,
			wantB: 0,
		},
		{
			hue:   360.0,
			sat:   100,
			val:   100,
			wantR: 255,
			wantG: 0,
			wantB: 0,
		},
		{
			hue:   0.0,
			sat:   100,
			val:   100,
			wantR: 255,
			wantG: 0,
			wantB: 0,
		},
		{
			hue:     -1.0,
			sat:     100,
			val:     100,
			wantErr: errors.New("invalid input for hue, saturation or value"),
		},
		{
			hue:     361.0,
			sat:     100,
			val:     100,
			wantErr: errors.New("invalid input for hue, saturation or value"),
		},
		{
			hue:     360.0,
			sat:     101,
			val:     100,
			wantErr: errors.New("invalid input for hue, saturation or value"),
		},
		{
			hue:     360.0,
			sat:     -1,
			val:     100,
			wantErr: errors.New("invalid input for hue, saturation or value"),
		},
		{
			hue:     360.0,
			sat:     100,
			val:     101,
			wantErr: errors.New("invalid input for hue, saturation or value"),
		},
		{
			hue:     360.0,
			sat:     100,
			val:     -1,
			wantErr: errors.New("invalid input for hue, saturation or value"),
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("hsvToRGB(%f,%f,%f)", c.hue, c.sat, c.val), func(t *testing.T) {
			r, g, b, err := hsvToRGB(c.hue, c.sat, c.val)
			if err != nil && c.wantErr == nil {
				t.Fatalf("expected no error, but got %s", err)
			}
			if err != nil && c.wantErr.Error() != err.Error() {
				t.Fatalf("expected error %s, but got %s", c.wantErr, err)
			}
			if r != c.wantR {
				t.Fatalf("expected r to be %d, but got %d\n", c.wantR, r)
			}
			if g != c.wantG {
				t.Fatalf("expected g to be %d, but got %d\n", c.wantG, g)
			}
			if b != c.wantB {
				t.Fatalf("expected b to be %d, but got %d\n", c.wantB, b)
			}
		})
	}

}
