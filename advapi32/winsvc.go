//go:build windows
// +build windows

// https://docs.microsoft.com/en-us/windows/win32/api/winsvc/

package advapi32

import (
	"syscall"
	"unsafe"

	"github.com/mel2oo/win32/typedef"
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

type ServiceType typedef.DWORD

const (
	SERVICE_KERNEL_DRIVER       ServiceType = 0x00000001
	SERVICE_FILE_SYSTEM_DRIVER  ServiceType = 0x00000002
	SERVICE_ADAPTER             ServiceType = 0x00000004
	SERVICE_RECOGNIZER_DRIVER   ServiceType = 0x00000008
	SERVICE_DRIVER              ServiceType = 0x0000000B
	SERVICE_WIN32_OWN_PROCESS   ServiceType = 0x00000010
	SERVICE_WIN32_SHARE_PROCESS ServiceType = 0x00000020
	SERVICE_INTERACTIVE_PROCESS ServiceType = 0x00000100
	SERVICE_WIN32               ServiceType = 0x00000030
	SERVICE_NO_CHANGE           ServiceType = 0xffffffff
)

type ServiceCurrentState typedef.DWORD

const (
	SERVICE_STOPPED          ServiceCurrentState = 0x00000001
	SERVICE_START_PENDING    ServiceCurrentState = 0x00000002
	SERVICE_STOP_PENDING     ServiceCurrentState = 0x00000003
	SERVICE_RUNNING          ServiceCurrentState = 0x00000004
	SERVICE_CONTINUE_PENDING ServiceCurrentState = 0x00000005
	SERVICE_PAUSE_PENDING    ServiceCurrentState = 0x00000006
	SERVICE_PAUSED           ServiceCurrentState = 0x00000007
)

type ServiceAcceptControls typedef.DWORD

const (
	SERVICE_ACCEPT_STOP                  ServiceAcceptControls = 0x00000001
	SERVICE_ACCEPT_PAUSE_CONTINUE        ServiceAcceptControls = 0x00000002
	SERVICE_ACCEPT_SHUTDOWN              ServiceAcceptControls = 0x00000004
	SERVICE_ACCEPT_PARAMCHANGE           ServiceAcceptControls = 0x00000008
	SERVICE_ACCEPT_NETBINDCHANGE         ServiceAcceptControls = 0x00000010
	SERVICE_ACCEPT_HARDWAREPROFILECHANGE ServiceAcceptControls = 0x00000020
	SERVICE_ACCEPT_POWEREVENT            ServiceAcceptControls = 0x00000040
	SERVICE_ACCEPT_SESSIONCHANGE         ServiceAcceptControls = 0x00000080
	SERVICE_ACCEPT_PRESHUTDOWN           ServiceAcceptControls = 0x00000100
	SERVICE_ACCEPT_TIMECHANGE            ServiceAcceptControls = 0x00000200
	SERVICE_ACCEPT_TRIGGEREVENT          ServiceAcceptControls = 0x00000400
)

type ServiceFlags typedef.DWORD

const (
	SERVICE_RUNS_IN_SYSTEM_PROCESS ServiceFlags = 0x00000001
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

type SERVICE_STATUS_PROCESS struct {
	DwServiceType             ServiceType
	DwCurrentState            ServiceCurrentState
	DwControlsAccepted        ServiceAcceptControls
	DwWin32ExitCode           typedef.DWORD
	DwServiceSpecificExitCode typedef.DWORD
	DwCheckPoint              typedef.DWORD
	DwWaitHint                typedef.DWORD
	DwProcessId               typedef.DWORD
	DwServiceFlags            ServiceFlags
}

// https://docs.microsoft.com/en-us/windows/win32/api/winsvc/nf-winsvc-closeservicehandle
func CloseServiceHandle(hSCObject SC_HANDLE) typedef.BOOL {
	ret, _, _ := procCloseServiceHandle.Call(uintptr(hSCObject))
	return typedef.BOOL(ret)
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

	ret, _, _ := procOpenSCManager.Call(
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

	ret, _, _ := procOpenService.Call(
		uintptr(hSCManager),
		uintptr(unsafe.Pointer(a1)),
		uintptr(dwDesiredAccess),
	)

	return SC_HANDLE(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winsvc/nf-winsvc-queryservicestatus
func QueryServiceStatus(hService SC_HANDLE, lpServiceStatus *SERVICE_STATUS) typedef.BOOL {
	ret, _, _ := procQueryServiceStatus.Call(
		uintptr(hService),
		uintptr(unsafe.Pointer(lpServiceStatus)),
	)

	return typedef.BOOL(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winsvc/nf-winsvc-queryservicestatusex
func QueryServiceStatusEx(hService SC_HANDLE, InfoLevel SC_STATUS_TYPE, lpBuffer *SERVICE_STATUS_PROCESS,
	cbBufSize typedef.DWORD, pcbBytesNeeded *typedef.DWORD) typedef.BOOL {
	ret, _, _ := procQueryServiceStatusEx.Call(
		uintptr(hService),
		uintptr(InfoLevel),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(cbBufSize),
		uintptr(unsafe.Pointer(pcbBytesNeeded)),
	)

	return typedef.BOOL(ret)
}
