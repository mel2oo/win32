package main

import (
	"fmt"
	"unsafe"

	"github.com/mel2oo/win32/advapi32/winsvc"
	"github.com/mel2oo/win32/types"
)

func main() {
	scmhandler := winsvc.OpenSCManager("", "", types.SC_MANAGER_ALL_ACCESS)
	if scmhandler == 0 {
		return
	}
	defer winsvc.CloseServiceHandle(scmhandler)

	srvhandler := winsvc.OpenService(scmhandler, "Appinfo", types.SERVICE_ALL_ACCESS)
	if srvhandler == 0 {
		return
	}
	defer winsvc.CloseServiceHandle(srvhandler)

	var ssStatus types.SERVICE_STATUS_PROCESS
	var dwBytesNeeded types.DWORD

	if winsvc.QueryServiceStatusEx(
		srvhandler,
		types.SC_STATUS_PROCESS_INFO,
		&ssStatus,
		types.DWORD(unsafe.Sizeof(ssStatus)),
		&dwBytesNeeded) == 0 {
		return
	}

	fmt.Println(ssStatus.CurrentState)
}
