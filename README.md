# Shuaifu

Yet another language written in Golang

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/yestagram/shuaifu-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/yestagram/shuaifu-go?style=flat-square)](https://goreportcard.com/report/github.com/yestagram/shuaifu-go)
[![GitHub](https://img.shields.io/github/license/yestagram/shuaifu-go?color=blue&style=flat-square)](LICENSE)


## Install
 
```shell
go get github.com/yestagram/shuaifu-go
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/yestagram/shuaifu-go"
)

func main() {
	shuaifu.Believe() // print believe words
	fmt.Println(shuaifu.Version()) // print version
	fmt.Println(shuaifu.Translate("woxbyhukfu!")) // translate into pinyin
}
```