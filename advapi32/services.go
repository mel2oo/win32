package advapi32

import (
	"errors"
	"syscall"
	"unsafe"
	"win32/typedef"
)

type SC_HANDLE typedef.HANDLE

type SCManagerAccess typedef.DWORD

const (
	GENERIC_READ                  SCManagerAccess = 0x80000000
	GENERIC_WRITE                 SCManagerAccess = 0x40000000
	GENERIC_EXECUTE               SCManagerAccess = 0x20000000
	SC_MANAGER_CONNECT            SCManagerAccess = 0x0001
	SC_MANAGER_CREATE_SERVICE     SCManagerAccess = 0x0002
	SC_MANAGER_ENUMERATE_SERVICE  SCManagerAccess = 0x0004
	SC_MANAGER_LOCK               SCManagerAccess = 0x0008
	SC_MANAGER_QUERY_LOCK_STATUS  SCManagerAccess = 0x0010
	SC_MANAGER_MODIFY_BOOT_CONFIG SCManagerAccess = 0x0020
	SC_MANAGER_ALL_ACCESS         SCManagerAccess = 0xF003F
)

type ServiceAccess typedef.DWORD

const (
	SERVICE_QUERY_CONFIG                      ServiceAccess = 0x0001
	SERVICE_CHANGE_CONFIG                     ServiceAccess = 0x0002
	SERVICE_QUERY_STATUS                      ServiceAccess = 0x0004
	SERVICE_ENUMERATE_DEPENDENTSServiceAccess ServiceAccess = 0x0008
	SERVICE_START                             ServiceAccess = 0x0010
	SERVICE_STOP                              ServiceAccess = 0x0020
	SERVICE_PAUSE_CONTINUE                    ServiceAccess = 0x0040
	SERVICE_INTERROGATE                       ServiceAccess = 0x0080
	SERVICE_USER_DEFINED_CONTROLServiceAccess ServiceAccess = 0x0100
	SERVICE_ALL_ACCESS                        ServiceAccess = 0xF01FF
	ACCESS_SYSTEM_SECURITY                    ServiceAccess = 0x01000000
	DELETE                                    ServiceAccess = 0x00010000
	READ_CONTROL                              ServiceAccess = 0x00020000
	WRITE_DAC                                 ServiceAccess = 0x00040000
	WRITE_OWNER                               ServiceAccess = 0x00080000
	SERVICE_GENERIC_READ                      ServiceAccess = 0x80000000
	SERVICE_GENERIC_WRITE                     ServiceAccess = 0x40000000
	SERVICE_GENERIC_EXECUTE                   ServiceAccess = 0x20000000
	MAXIMUM_ALLOWED                           ServiceAccess = 0x02000000
)

type SC_STATUS_TYPE typedef.UINT

const (
	SC_STATUS_PROCESS_INFO SC_STATUS_TYPE = iota
)

type SERVICE_STATUS struct {
	DwServiceType             uint32
	DwCurrentState            uint32
	DwControlsAccepted        uint32
	DwWin32ExitCode           uint32
	DwServiceSpecificExitCode uint32
	DwCheckPoint              uint32
	DwWaitHint                uint32
}

func CloseServiceHandle(hSCObject SC_HANDLE) (bool, error) {
	ret, _, _ := procCloseServiceHandle.Call(uintptr(hSCObject))
	if ret == 0 {
		return false, syscall.GetLastError()
	}
	return true, nil
}

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
		uintptr(dwDesiredAccess),
	)

	if ret == 0 {
		return 0, syscall.GetLastError()
	}

	return SC_HANDLE(ret), nil
}

func OpenService(hSCManager SC_HANDLE, lpServiceName string, dwDesiredAccess ServiceAccess) (SC_HANDLE, error) {
	a2, err := syscall.UTF16PtrFromString(lpServiceName)
	if err != nil {
		return 0, err
	}

	ret, _, _ := procOpenService.Call(
		uintptr(hSCManager),
		uintptr(unsafe.Pointer(a2)),
		uintptr(dwDesiredAccess),
	)

	if ret == 0 {
		return 0, syscall.GetLastError()
	}

	return SC_HANDLE(ret), nil
}

func QueryServiceStatus(hService SC_HANDLE, lpServiceStatus *SERVICE_STATUS) (bool, error) {
	if lpServiceStatus == nil {
		return false, errors.New("")
	}

	ret, _, _ := procQueryServiceStatus.Call(
		uintptr(hService),
		uintptr(unsafe.Pointer(lpServiceStatus)),
	)

	if ret == 0 {
		return false, syscall.GetLastError()
	}

	return true, nil
}

func QueryServiceStatusEx(hService SC_HANDLE, InfoLevel SC_STATUS_TYPE, lpBuffer *typedef.BYTE,
	cbBufSize typedef.DWORD, pcbBytesNeeded *typedef.DWORD) (bool, error) {
	ret, _, _ := procQueryServiceStatusEx.Call(
		uintptr(hService),
		uintptr(InfoLevel),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(cbBufSize),
		uintptr(unsafe.Pointer(pcbBytesNeeded)),
	)

	if ret == 0 {
		return false, syscall.GetLastError()
	}

	return true, nil
}
