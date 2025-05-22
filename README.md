<picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://github.com/Mad-Pixels/.github/raw/main/profile/banner.png">
    <source media="(prefers-color-scheme: light)" srcset="https://github.com/Mad-Pixels/.github/raw/main/profile/banner.png">
    <img
        alt="MadPixels"
        src="https://github.com/Mad-Pixels/.github/raw/main/profile/banner.png">
</picture>

# goldmark-tgmd ✨

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.19-blue.svg)](https://golang.org)

goldmark-tgmd is an extension for the goldmark Markdown parser 
that adds support for Telegram-specific Markdown features 🚀. 
This library allows developers to render Markdown text according 
to Telegram's formatting options, making it easier to create content 
for bots 🤖 and applications integrated with Telegram 📱.

## Features 🌟

- Support for Telegram Markdown features including custom entities like hidden text and strikethrough text. 📝
- Easy integration with goldmark-based projects. 🔌
- Extensible architecture for further customizations. 🔨

## Getting Started 🚀
### Prerequisites 📋
- Go 1.19 or higher

### Installation 💽
To install goldmark-tgmd, use the following go get command:
```shell
go get github.com/Mad-Pixels/goldmark-tgmd
```

### Usage 🛠️
```go
package main

import (
   "bytes"
   "fmt"
   "os"
   
   tgmd "github.com/Mad-Pixels/goldmark-tgmd"
)

func main() {
   var buf bytes.Buffer
   content, _ := os.ReadFile("./examples/source.md")
   
   md := tgmd.TGMD()
   _ = md.Convert(content, &buf)
   fmt.Println(buf.String())
}
```

You can try [example](./example)

# Contributing
We're open to any new ideas and contributions. We also have some rules and taboos here, so please read this page and our [Code of Conduct](/CODE_OF_CONDUCT.md) carefully.

## I want to report an issue
If you've found an issue and want to report it, please check our [Issues](https://github.com/Mad-Pixels/goldmark-tgmd/issues) page.