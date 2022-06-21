# colors

Very simple go library for rendering colors in terminal.

### Install

```
go get github.com/dbackowski/colors
```

### Usage

```go
package main

import (
	"fmt"

	"github.com/dbackowski/colors"
)

func main() {
	fmt.Println(colors.Colorize("green text on red background", colors.FgGreen, colors.BgRed))
	fmt.Println(colors.Colorize("green text", colors.FgGreen))
}
```

![screenshot](https://i.imgur.com/ygmsspZ.png)
