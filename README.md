# parse-author

[![Go Reference](https://pkg.go.dev/badge/github.com/ImSingee/parse-author.svg)](https://pkg.go.dev/github.com/ImSingee/parse-author) [![Test Status](https://github.com/ImSingee/parse-author/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/ImSingee/parse-author/actions/workflows/test.yml?query=branch%3Amaster) [![codecov](https://codecov.io/gh/ImSingee/parse-author/branch/master/graph/badge.svg?token=RWV4ZYS1DH)](https://codecov.io/gh/ImSingee/parse-author) [![Go Report Card](https://goreportcard.com/badge/github.com/ImSingee/parse-author)](https://goreportcard.com/report/github.com/ImSingee/parse-author)


Golang port of [author-regex](https://github.com/jonschlinkert/author-regex), and provide an author parser

## Install

```bash
go get -u github.com/ImSingee/parse-author
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/ImSingee/parse-author"
)

func main() {
	fmt.Println(pauthor.RegexStr)

	re := pauthor.Regex()
	fmt.Println(re.FindAllStringSubmatch(`Jon Schlinkert <jon.schlinkert@sellside.com> (https://github.com/jonschlinkert)`, -1))
}
```