- *赞助 BTC: 1Cbd6oGAUUyBi7X7MaR4np4nTmQZXVgkCW*
- *赞助 ETH: 0x623A3C3a72186A6336C79b18Ac1eD36e1c71A8a6*

----

# base60天干地支编码

- godoc: https://godoc.org/github.com/chai2010/base60

## Install

1. go get github.com/chai2010/base60
1. go run hello.go

## Example

```go
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
```

## BUGS

Report bugs to <chaishushan@gmail.com>.

Thanks!
