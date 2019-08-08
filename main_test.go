package gotools

import (
	"bytes"
	"testing"
)

func TestPrint(t *testing.T) {
	Print("Test1", "Test2")
	Print("=============")
	Print("Test1", "\r\n", "Test2")
	Print("=============")
	Print("Test1", "\r\n", "Test2", "\r\n")
	Print("=============")
	Print("Test1\r\n", "Test2\r\n")
	Print("=============")
	Print(123, 456)
	Print("=============")
	Print("", "t")
	Print("=============")
}

func TestDebug(t *testing.T) {
	DebugSlice(make([]string, 10, 12))
	Print("=============")
	DebugString("gotools")
	Print("=============")
	DebugBuffer(bytes.NewBuffer(make([]byte, 10, 12)))
	Print("=============")
}
