// Copyright 2018 chaishushan@gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"

	"github.com/chai2010/base60"
)

func main() {
	s0 := base60.Encode([]byte("你好"))
	fmt.Println(s0)

	s1 := base60.Decode("乙丑癸巳甲寅己亥丁卯甲申丁未甲午己巳")
	fmt.Println(string(s1))

	// Output:
	// 乙丑癸巳甲寅己亥丁卯甲申丁未甲午己巳
	// 你好
}
