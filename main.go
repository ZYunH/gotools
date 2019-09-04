package gotools

import (
	"bytes"
	"fmt"
	"reflect"
	"unsafe"
)

/*----------MockHeaders----------*/
type interfaceSliceHdr struct {
	int
	h *reflect.SliceHeader
}

type mockBuffer struct {
	Buf      []byte // contents are the bytes buf[off : len(buf)]
	Off      int    // read at &buf[off], write at &buf[len(buf)]
	LastRead int8   // last read operation, so that Unread* can work correctly.
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
	Hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	Print("Data ->", Hdr.Data)
	Print("Len ->", Hdr.Len)
}

/*DebugBuffer*/
const (
	opRead      int8 = -1 // Any other read operation.
	opInvalid   int8 = 0  // Non-read operation.
	opReadRune1 int8 = 1  // Read rune of size 1.
	opReadRune2 int8 = 2  // Read rune of size 2.
	opReadRune3 int8 = 3  // Read rune of size 3.
	opReadRune4 int8 = 4  // Read rune of size 4.
)

func DebugBuffer(b *bytes.Buffer) {
	lookupBuffer := *(*mockBuffer)(unsafe.Pointer(b))
	buf := (*reflect.SliceHeader)(unsafe.Pointer(&lookupBuffer.Buf))
	Print("Buffer.buf.Data", "->", buf.Data)
	Print("Buffer.buf.Len", "->", buf.Len)
	Print("Buffer.buf.Cap", "->", buf.Cap)
	Print("Buffer.off", "->", lookupBuffer.Off)

	var lastReadText string
	if lookupBuffer.LastRead == opRead {
		lastReadText = "(opRead)"
	} else if lookupBuffer.LastRead == opInvalid {
		lastReadText = "(opInvalid)"
	} else if lookupBuffer.LastRead == opReadRune1 {
		lastReadText = "(opReadRune1)"
	} else if lookupBuffer.LastRead == opReadRune2 {
		lastReadText = "(opReadRune2)"
	} else if lookupBuffer.LastRead == opReadRune3 {
		lastReadText = "(opReadRune3)"
	} else if lookupBuffer.LastRead == opReadRune4 {
		lastReadText = "(opReadRune4)"
	}
	Print("Buffer.lastRead", "->", lookupBuffer.LastRead, lastReadText)

}
