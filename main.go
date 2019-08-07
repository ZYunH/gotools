package gotools

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*----------MockHeaders----------*/
type interfaceSliceHdr struct {
	int
	h *reflect.SliceHeader
}

// Simple print function used for debug
func Print(args ...interface{}) {
	if len(args) == 0 {
		return
	}

	if !includeCRLF(args[len(args)-1]) {
		args = append(args, "\r\n")
	}

	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			if len(v) >= 2 && !(v[len(v)-2:] == "\r\n") {
				// Not include CRLF
				print(fmt.Sprint(arg), " ")
			} else {
				print(fmt.Sprint(arg))
			}
		default:
			print(fmt.Sprint(arg), " ")
		}
	}
}

func includeCRLF(i interface{}) bool {
	if i == "\r\n" {
		return true
	}
	switch v := i.(type) {
	case string:
		L := len(v)
		if L >= 2 && (v[L-1] == 10 && v[L-2] == 13) {
			return true
		}
	}
	return false
}

func DebugSlice(v interface{}) {
	mockInterface := *(*interfaceSliceHdr)(unsafe.Pointer(&v))
	Hdr := *(mockInterface.h)
	Print("Data ->", Hdr.Data)
	Print("Len ->", Hdr.Len)
	Print("Cap ->", Hdr.Cap)
}

func DebugString(s string) {
	Hdr := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	Print("Data ->", Hdr.Data)
	Print("Len ->", Hdr.Len)
}
