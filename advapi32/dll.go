package advapi32

import (
	"golang.org/x/sys/windows"
)

var (
	dll               = windows.NewLazyDLL("advapi32.dll")
	procOpenSCManager = dll.NewProc("OpenSCManagerW")
)
