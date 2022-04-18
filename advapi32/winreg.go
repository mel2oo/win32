//go:build windows
// +build windows

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/

package advapi32

import (
	"syscall"
	"unsafe"

	"github.com/mel2oo/win32/typedef"
)

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-abortsystemshutdownw
func AbortSystemShutdown(lpMachineName string) typedef.BOOL {
	a1, err := syscall.UTF16PtrFromString(lpMachineName)
	if err != nil {
		return 0
	}

	ret, _, _ := procAbortSystemShutdown.Call(uintptr(unsafe.Pointer(a1)))

	return typedef.BOOL(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-initiateshutdownw
func InitiateShutdown(lpMachineName, lpMessage string, dwGracePeriod, dwShutdownFlags, dwReason typedef.DWORD) typedef.DWORD {
	a1, err := syscall.UTF16PtrFromString(lpMachineName)
	if err != nil {
		return 0
	}

	a2, err := syscall.UTF16PtrFromString(lpMessage)
	if err != nil {
		return 0
	}

	ret, _, _ := procInitiateShutdown.Call(
		uintptr(unsafe.Pointer(a1)),
		uintptr(unsafe.Pointer(a2)),
		uintptr(dwGracePeriod),
		uintptr(dwShutdownFlags),
		uintptr(dwReason),
	)

	return typedef.DWORD(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-initiatesystemshutdownw
func InitiateSystemShutdown(lpMachineName, lpMessage string, dwTimeout typedef.DWORD,
	bForceAppsClosed, bRebootAfterShutdown typedef.BOOL) typedef.BOOL {
	a1, err := syscall.UTF16PtrFromString(lpMachineName)
	if err != nil {
		return 0
	}

	a2, err := syscall.UTF16PtrFromString(lpMessage)
	if err != nil {
		return 0
	}

	ret, _, _ := procInitiateSystemShutdown.Call(
		uintptr(unsafe.Pointer(a1)),
		uintptr(unsafe.Pointer(a2)),
		uintptr(dwTimeout),
		uintptr(bForceAppsClosed),
		uintptr(bRebootAfterShutdown),
	)

	return typedef.BOOL(ret)
}
