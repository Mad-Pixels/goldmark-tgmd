package tgmd

import (
	"bytes"
	"os"

	"github.com/yuin/goldmark/util"
)

func writeSpecialTagStart(w util.BufWriter, tag SpecialTag, prefix []byte) {
	writeWrapperArr(w.Write(tag.Bytes()))
	writeCustomBytes(w, prefix)
}

func writeSpecialTagEnd(w util.BufWriter, tag SpecialTag, postfix []byte) {
	writeCustomBytes(w, postfix)
	writeWrapperArr(w.Write(tag.Bytes()))
}

func writeNewLine(w util.BufWriter) {
	writeCustomBytes(w, []byte{NewLine.Byte()})
}

func render(w util.BufWriter, b []byte) {
	data := bytes.ReplaceAll(b, []byte{Space.Byte(), Space.Byte()}, []byte{NewLine.Byte()})
	writeCustomBytes(w, data)
}

func writeRowBytes(w util.BufWriter, data []byte) {
	writeWrapperArr(w.Write(data))
}

func writeCustomBytes(w util.BufWriter, data []byte) {
	for _, char := range data {
		if escaped, ok := escape[char]; ok {
			writeWrapperArr(w.Write(escaped))
			continue
		}
		writeWrapper(w.WriteByte(char))
	}
}

func writeWrapperArr(_ int, err error) {
	writeWrapper(err)
}

func writeWrapper(err error) {
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
	}
}
