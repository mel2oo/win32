package kernel32

import "golang.org/x/sys/windows"

var (
	dll = windows.NewLazyDLL("kernel32.dll")

	procGetLastError = dll.NewProc("GetLastError")
)
