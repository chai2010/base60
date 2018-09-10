// Copyright 2018 chaishushan@gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package base60 implements base60(天干地支) encoding.
package base60

import (
	"math/big"
	"strings"
)

var (
	base60Table      = [60]string{}
	base60IndexTable = map[string]int{}
)

func init() {
	var (
		天干 = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
		地支 = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	)

	for i := 0; i < 60; i++ {
		base60Table[i] = 天干[i%10] + 地支[i%12]
		base60IndexTable[base60Table[i]] = i
	}
}

func Encode(ba []byte) string {
	if len(ba) == 0 {
		return ""
	}

	ri := len(ba) * 138 / 100
	ra := make([]byte, ri+1)

	x := new(big.Int).SetBytes(ba) // ba is big-endian
	x.Abs(x)
	y := big.NewInt(60)
	m := new(big.Int)

	for x.Sign() > 0 {
		x, m = x.DivMod(x, y, m)
		ra[ri] = byte(m.Int64())
		ri--
	}

	// Leading zeroes encoded as base60Table zeros
	for i := 0; i < len(ba); i++ {
		if ba[i] != 0 {
			break
		}
		ra[ri] = 0
		ri--
	}
	ra = ra[ri+1:]

	var b strings.Builder
	for i := 0; i < len(ra); i++ {
		b.WriteString(base60Table[ra[i]])
	}

	return b.String()
}

func Decode(s string) []byte {
	var ba = []rune(s)

	if len(ba) == 0 {
		return nil
	}

	x := new(big.Int)
	y := big.NewInt(60)
	z := new(big.Int)

	for i := 0; i < len(ba); i += 2 {
		v := base60IndexTable[string(ba[i:i+2])]
		z.SetInt64(int64(v))
		x.Mul(x, y)
		x.Add(x, z)
	}
	xa := x.Bytes()

	// Restore leading zeros
	i := 0
	for i < len(ba) && string(ba[i:i+2]) == "甲子" {
		i += 2
	}

	ra := make([]byte, (i/2)+len(xa))
	copy(ra[(i/2):], xa)
	return ra
}
