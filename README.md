# colors ![Tests](https://github.com/dbackowski/colors/actions/workflows/test.yml/badge.svg)

Very simple GO library for rendering colors in the terminal.<br />
Does not currently support the windows platform.

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
	fmt.Println(colors.Colorize("FgBlack", colors.FgBlack))
	fmt.Println(colors.Colorize("FgRed", colors.FgRed))
	fmt.Println(colors.Colorize("FgGreen", colors.FgGreen))
	fmt.Println(colors.Colorize("FgYellow", colors.FgYellow))
	fmt.Println(colors.Colorize("FgBlue", colors.FgBlue))
	fmt.Println(colors.Colorize("FgMagenta", colors.FgMagenta))
	fmt.Println(colors.Colorize("FgCyan", colors.FgCyan))
	fmt.Println(colors.Colorize("FgWhite", colors.FgWhite))

	fmt.Println(colors.Colorize("BgBlack", colors.Default, colors.BgBlack))
	fmt.Println(colors.Colorize("BgRed", colors.Default, colors.BgRed))
	fmt.Println(colors.Colorize("BgGreen", colors.Default, colors.BgGreen))
	fmt.Println(colors.Colorize("BgYellow", colors.Default, colors.BgYellow))
	fmt.Println(colors.Colorize("BgBlue", colors.Default, colors.BgBlue))
	fmt.Println(colors.Colorize("BgMagenta", colors.Default, colors.BgMagenta))
	fmt.Println(colors.Colorize("BgCyan", colors.Default, colors.BgCyan))
	fmt.Println(colors.Colorize("BgWhite", colors.Default, colors.BgWhite))
}

```

![screenshot](https://i.imgur.com/LheCsFW.png)
