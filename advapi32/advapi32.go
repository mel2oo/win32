package advapi32

import (
	"syscall"
	"unsafe"
)

func OpenSCManager(lpMachineName, lpDatabaseName string, dwDesiredAccess SCManagerAccess) (SC_HANDLE, error) {
	a1, err := syscall.UTF16PtrFromString(lpMachineName)
	if err != nil {
		return 0, err
	}

	a2, err := syscall.UTF16PtrFromString(lpDatabaseName)
	if err != nil {
		return 0, err
	}

	ret, _, _ := procOpenSCManager.Call(
		uintptr(unsafe.Pointer(a1)),
		uintptr(unsafe.Pointer(a2)),
		uintptr(dwDesiredAccess))

	if ret == 0 {
		return 0, syscall.GetLastError()
	}

	return SC_HANDLE(ret), nil
}
