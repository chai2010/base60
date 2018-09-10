// Copyright 2018 chaishushan@gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base60

import (
	"testing"
)

type base60Test struct {
	en, de string
}

var base60Golden = []base60Test{
	{"", ""},
}

func TestEncodeBase60(t *testing.T) {
	for _, g := range base60Golden {
		s := string(Encode([]byte(g.en)))
		if s != g.de {
			t.Errorf("Bad EncodeBase60. Need=%v, Got=%v", g.de, s)
		}
	}
}

func TestDecodeBase60(t *testing.T) {
	t.Skip("TODO")
	for _, g := range base60Golden {
		s := string(Decode(g.de))
		if s != g.en {
			t.Errorf("Bad DecodeBase60. Need=%v, Got=%v", g.en, s)
		}
	}
}
