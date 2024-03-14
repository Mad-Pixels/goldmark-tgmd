# goldmark-tgmd ✨

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

### Contributing 🤝
Contributions are welcome! Feel free to open issues for bugs 🐛, 
feature requests 🌈, or submit pull requests 💡. 

Acknowledgments 💖
- Thanks to the [goldmark](https://github.com/yuin/goldmark) project for providing a robust and extensible Markdown parser.
- This project is inspired by the formatting options available in Telegram.