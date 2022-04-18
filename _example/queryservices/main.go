package main

import (
	"fmt"
	"unsafe"

	"github.com/mel2oo/win32/advapi32"
	"github.com/mel2oo/win32/typedef"
)

func main() {
	scmhandler := advapi32.OpenSCManager("", "", advapi32.SC_MANAGER_ALL_ACCESS)
	if scmhandler == 0 {
		return
	}
	defer advapi32.CloseServiceHandle(scmhandler)

	srvhandler := advapi32.OpenService(scmhandler, "Appinfo", advapi32.SERVICE_ALL_ACCESS)
	if srvhandler == 0 {
		return
	}
	defer advapi32.CloseServiceHandle(srvhandler)

	var ssStatus advapi32.SERVICE_STATUS_PROCESS
	var dwBytesNeeded typedef.DWORD

	if advapi32.QueryServiceStatusEx(
		srvhandler,
		advapi32.SC_STATUS_PROCESS_INFO,
		&ssStatus,
		typedef.DWORD(unsafe.Sizeof(ssStatus)),
		&dwBytesNeeded) == typedef.FALSE {
		return
	}

	fmt.Println(ssStatus.DwCurrentState)
}
