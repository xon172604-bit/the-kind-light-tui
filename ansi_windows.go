//go:build windows

package main

import (
	"os"
	"syscall"
	"unsafe"
)

// enableANSI turns on virtual-terminal processing so colour codes render in
// older Windows consoles. Windows Terminal already has it on.
func enableANSI() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getMode := kernel32.NewProc("GetConsoleMode")
	setMode := kernel32.NewProc("SetConsoleMode")
	h := os.Stdout.Fd()
	var mode uint32
	if r, _, _ := getMode.Call(h, uintptr(unsafe.Pointer(&mode))); r == 0 {
		return
	}
	setMode.Call(h, uintptr(mode|0x0004)) // ENABLE_VIRTUAL_TERMINAL_PROCESSING
}
