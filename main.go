package gotools

import (
	"fmt"
	"unsafe"
)

var sizeOfInt = unsafe.Sizeof(0)

// Simple print function used for debug
func DebugPrint(args ...interface{}) {
	if !includeCRLF(args[len(args)-1]) {
		args = append(args, "\r\n")
	}

	for _, v := range args {
		switch v.(type) {
		case string:
			interfaceAddr := uintptr(unsafe.Pointer(&v))
			dataPtr := interfaceAddr + sizeOfInt
			dataAddr := *(*uintptr)(unsafe.Pointer(dataPtr)) // type:string

			dataAsSlice := *(*[]byte)(unsafe.Pointer(dataAddr))
			dataAsSliceLen := *(*int)(unsafe.Pointer(dataAddr + sizeOfInt))

			if dataAsSliceLen >= 2 && !(dataAsSlice[dataAsSliceLen-1] == 10 && dataAsSlice[dataAsSliceLen-2] == 13) {
				// Not include CRLF
				print(fmt.Sprint(v), " ")
			} else {
				print(fmt.Sprint(v))
			}
		default:
			print(fmt.Sprint(v), " ")
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

// Not used, for read only
func toAddress(v interface{}) uintptr {
	return uintptr(unsafe.Pointer(&v))
}

// Not used
func toIntPointer(address uintptr) *int {
	return (*int)(unsafe.Pointer(address))
}

// Not used
func toInt(address uintptr) int {
	return *(*int)(unsafe.Pointer(address))
}
