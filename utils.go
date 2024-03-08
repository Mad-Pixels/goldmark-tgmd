package tgmd

import (
	"reflect"
	"unsafe"
)

// StringToBytes convert a string to a byte slice.
func StringToBytes(v string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&v))

	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return
}
