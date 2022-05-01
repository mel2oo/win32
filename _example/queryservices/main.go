package main

import (
	"fmt"
	"unsafe"

	"github.com/mel2oo/win32/advapi32/winsvc"
	"github.com/mel2oo/win32/types"
)

func main() {
	scmhandler := winsvc.OpenSCManager("", "", winsvc.SC_MANAGER_ALL_ACCESS)
	if scmhandler == 0 {
		return
	}
	defer winsvc.CloseServiceHandle(scmhandler)

	srvhandler := winsvc.OpenService(scmhandler, "Appinfo", winsvc.SERVICE_ALL_ACCESS)
	if srvhandler == 0 {
		return
	}
	defer winsvc.CloseServiceHandle(srvhandler)

	var ssStatus winsvc.SERVICE_STATUS_PROCESS
	var dwBytesNeeded types.DWORD

	if winsvc.QueryServiceStatusEx(
		srvhandler,
		winsvc.SC_STATUS_PROCESS_INFO,
		&ssStatus,
		types.DWORD(unsafe.Sizeof(ssStatus)),
		&dwBytesNeeded) == 0 {
		return
	}

	fmt.Println(ssStatus.DwCurrentState)
}
