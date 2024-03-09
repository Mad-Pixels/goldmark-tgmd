package main

import (
	"bytes"
	"fmt"
	"os"

	tgmd "github.com/Mad-Pixels/goldmark-tgmd"
)

func main() {
	var buf bytes.Buffer
	content, _ := os.ReadFile("./source.md")
	md := tgmd.TGMD()

	// change some configs for example:
	tgmd.Config.UpdatePrimaryListBullet('◦')
	tgmd.Config.UpdateHeading1(tgmd.Element{
		Style:   tgmd.BoldTg,
		Prefix:  "!!!",
		Postfix: "!!!",
	})

	_ = md.Convert(content, &buf)
	fmt.Println(buf.String())
}
