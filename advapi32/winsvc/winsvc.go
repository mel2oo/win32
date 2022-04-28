//go:build windows
// +build windows

// https://docs.microsoft.com/en-us/windows/win32/api/winsvc/

package winsvc

import (
	"syscall"
	"unsafe"

	"github.com/mel2oo/win32/advapi32"
	"github.com/mel2oo/win32/types"
)

// https://docs.microsoft.com/en-us/windows/win32/api/winsvc/nf-winsvc-closeservicehandle
func CloseServiceHandle(hSCObject SC_HANDLE) types.BOOL {
	ret, _, _ := advapi32.ProcCloseServiceHandle.Call(uintptr(hSCObject))
	return types.BOOL(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winsvc/nf-winsvc-openscmanagerw
func OpenSCManager(lpMachineName, lpDatabaseName string, dwDesiredAccess SCManagerAccess) SC_HANDLE {
	var p1, p2 uintptr

	if len(lpMachineName) > 0 {
		a1, err := syscall.UTF16PtrFromString(lpMachineName)
		if err != nil {
			return 0
		}
		p1 = uintptr(unsafe.Pointer(a1))
	}

	if len(lpDatabaseName) > 0 {
		a2, err := syscall.UTF16PtrFromString(lpDatabaseName)
		if err != nil {
			return 0
		}
		p2 = uintptr(unsafe.Pointer(a2))
	}

	ret, _, _ := advapi32.ProcOpenSCManager.Call(
		p1,
		p2,
		uintptr(dwDesiredAccess),
	)

	return SC_HANDLE(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winsvc/nf-winsvc-openservicew
func OpenService(hSCManager SC_HANDLE, lpServiceName string, dwDesiredAccess ServiceAccess) SC_HANDLE {
	a1, err := syscall.UTF16PtrFromString(lpServiceName)
	if err != nil {
		return 0
	}

	ret, _, _ := advapi32.ProcOpenService.Call(
		uintptr(hSCManager),
		uintptr(unsafe.Pointer(a1)),
		uintptr(dwDesiredAccess),
	)

	return SC_HANDLE(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winsvc/nf-winsvc-queryservicestatus
func QueryServiceStatus(hService SC_HANDLE, lpServiceStatus *SERVICE_STATUS) types.BOOL {
	ret, _, _ := advapi32.ProcQueryServiceStatus.Call(
		uintptr(hService),
		uintptr(unsafe.Pointer(lpServiceStatus)),
	)

	return types.BOOL(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winsvc/nf-winsvc-queryservicestatusex
func QueryServiceStatusEx(hService SC_HANDLE, InfoLevel SC_STATUS_TYPE, lpBuffer *SERVICE_STATUS_PROCESS,
	cbBufSize types.DWORD, pcbBytesNeeded *types.DWORD) types.BOOL {
	ret, _, _ := advapi32.ProcQueryServiceStatusEx.Call(
		uintptr(hService),
		uintptr(InfoLevel),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(cbBufSize),
		uintptr(unsafe.Pointer(pcbBytesNeeded)),
	)

	return types.BOOL(ret)
}
