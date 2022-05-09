package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/mel2oo/win32/shell32"
	"github.com/mel2oo/win32/types"
)

func main() {
	var lpExecInfo types.SHELLEXECUTEINFO
	lpExecInfo.Size = types.DWORD(unsafe.Sizeof(lpExecInfo))
	lpExecInfo.Mask = types.SEE_MASK_NOCLOSEPROCESS

	verb, err := syscall.UTF16PtrFromString("runas")
	if err != nil {
		return
	}
	lpExecInfo.Verb = verb

	file, err := syscall.UTF16PtrFromString("C:\\Windows\\System32\\cmd.exe")
	if err != nil {
		return
	}
	lpExecInfo.File = file
	lpExecInfo.Show = types.SW_SHOW

	if b := shell32.ShellExecuteExW(&lpExecInfo); b == 0 {
		fmt.Println("create process error")
	} else {
		fmt.Println("create process success")
	}
}
