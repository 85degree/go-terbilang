# Terbilang in Go

## Synopsis
Package terbilang is a package for styling int value to terbilang formated.

## Code Example

```
package main

import (
	"fmt"
	"github.com/apung/go-terbilang"
)

func main() {
	var angka int
	angka = 109209
	fmt.Println(terbilang.ToTerbilang(angka))
  fmt.Println(terbilang.ToTerbilangRp(angka))
  fmt.Println(terbilang.ToTerbilangSuffix("ekor",angka))
}
```

## Installation

As easy as `go get -u github.com/apung/go-terbilang`

## Contributors

You can always contribute this package and create pull requests. Just give me a note

## License

This package is released under the MIT License (MIT). Please read [LICENSE.md](https://github.com/apung/go-terbilang/LICENSE.md)
