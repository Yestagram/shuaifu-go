# Shuaifu

Yet another language written in Golang

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