# go-a1

[![test](https://github.com/kenkyu392/go-a1/workflows/test/badge.svg?branch=master)](https://github.com/kenkyu392/go-a1)
[![codecov](https://codecov.io/gh/kenkyu392/go-a1/branch/master/graph/badge.svg)](https://codecov.io/gh/kenkyu392/go-a1)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-00ADD8?logo=go)](https://pkg.go.dev/github.com/kenkyu392/go-a1)
[![go report card](https://goreportcard.com/badge/github.com/kenkyu392/go-a1)](https://goreportcard.com/report/github.com/kenkyu392/go-a1)
[![license](https://img.shields.io/github/license/kenkyu392/go-a1)](LICENSE)


## Installation

```
go get -u github.com/kenkyu392/go-a1
```

## Usage

```go
package main

import (
	"log"

	"github.com/kenkyu392/go-a1"
)

func main() {
	n, err := a1.Atoi("AZ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n) // -> 52

	s, err := a1.Itoa(n)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s) // -> AZ

	v, err := a1.Cell(8, 88)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(v) // -> CJ8
}
```

## License

[MIT](LICENSE)
