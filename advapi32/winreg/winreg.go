//go:build windows
// +build windows

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/

package winreg

import (
	"syscall"
	"unsafe"

	"github.com/mel2oo/win32/advapi32"
	"github.com/mel2oo/win32/types"
)

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-abortsystemshutdownw
func AbortSystemShutdown(lpMachineName string) types.BOOL {
	a1, err := syscall.UTF16PtrFromString(lpMachineName)
	if err != nil {
		return types.BOOL(0)
	}

	ret, _, _ := advapi32.ProcAbortSystemShutdown.Call(uintptr(unsafe.Pointer(a1)))

	return types.BOOL(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-initiateshutdownw
func InitiateShutdown(lpMachineName, lpMessage string, dwGracePeriod, dwShutdownFlags, dwReason types.DWORD) types.DWORD {
	a1, err := syscall.UTF16PtrFromString(lpMachineName)
	if err != nil {
		return types.DWORD(types.ERROR_INVALID_PARAMETER)
	}

	a2, err := syscall.UTF16PtrFromString(lpMessage)
	if err != nil {
		return types.DWORD(types.ERROR_INVALID_PARAMETER)
	}

	ret, _, _ := advapi32.ProcInitiateShutdown.Call(
		uintptr(unsafe.Pointer(a1)),
		uintptr(unsafe.Pointer(a2)),
		uintptr(dwGracePeriod),
		uintptr(dwShutdownFlags),
		uintptr(dwReason),
	)

	return types.DWORD(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-initiatesystemshutdownw
func InitiateSystemShutdown(lpMachineName, lpMessage string, dwTimeout types.DWORD,
	bForceAppsClosed, bRebootAfterShutdown types.BOOL) types.BOOL {
	a1, err := syscall.UTF16PtrFromString(lpMachineName)
	if err != nil {
		return types.BOOL(0)
	}

	a2, err := syscall.UTF16PtrFromString(lpMessage)
	if err != nil {
		return types.BOOL(0)
	}

	ret, _, _ := advapi32.ProcInitiateSystemShutdown.Call(
		uintptr(unsafe.Pointer(a1)),
		uintptr(unsafe.Pointer(a2)),
		uintptr(dwTimeout),
		uintptr(bForceAppsClosed),
		uintptr(bRebootAfterShutdown),
	)

	return types.BOOL(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-initiatesystemshutdownexw
func InitiateSystemShutdownEx(lpMachineName, lpMessage string, dwTimeout types.DWORD,
	bForceAppsClosed, bRebootAfterShutdown types.BOOL, dwReason types.DWORD) types.BOOL {
	a1, err := syscall.UTF16PtrFromString(lpMachineName)
	if err != nil {
		return types.BOOL(0)
	}

	a2, err := syscall.UTF16PtrFromString(lpMessage)
	if err != nil {
		return types.BOOL(0)
	}

	ret, _, _ := advapi32.ProcInitiateSystemShutdownEx.Call(
		uintptr(unsafe.Pointer(a1)),
		uintptr(unsafe.Pointer(a2)),
		uintptr(dwTimeout),
		uintptr(bForceAppsClosed),
		uintptr(bRebootAfterShutdown),
		uintptr(dwReason),
	)

	return types.BOOL(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regclosekey
func RegCloseKey(hKey HKEY) types.LSTATUS {
	ret, _, _ := advapi32.ProcRegCloseKey.Call(uintptr(hKey))
	return types.LSTATUS(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regconnectregistryw
func RegConnectRegistry(lpMachineName string, hKey HKEY, phkResult *HKEY) types.LSTATUS {
	a1, err := syscall.UTF16PtrFromString(lpMachineName)
	if err != nil {
		return types.LSTATUS(types.ERROR_INVALID_PARAMETER)
	}

	ret, _, _ := advapi32.ProcRegConnectRegistry.Call(
		uintptr(unsafe.Pointer(a1)),
		uintptr(hKey),
		uintptr(unsafe.Pointer(phkResult)),
	)

	return types.LSTATUS(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regcopytreew
func RegCopyTree(hKeySrc HKEY, lpSubKey string, hKeyDest HKEY) types.LSTATUS {
	a1, err := syscall.UTF16PtrFromString(lpSubKey)
	if err != nil {
		return types.LSTATUS(types.ERROR_INVALID_PARAMETER)
	}

	ret, _, _ := advapi32.ProcRegCopyTree.Call(
		uintptr(hKeySrc),
		uintptr(unsafe.Pointer(a1)),
		uintptr(hKeyDest),
	)

	return types.LSTATUS(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regcreatekeyexw
func RegCreateKeyEx(
	hKey HKEY, lpSubKey string,
	Reserved types.DWORD, lpClass string,
	dwOptions types.DWORD, samDesired REGSAM,
	lpSecurityAttributes types.LPSECURITY_ATTRIBUTES,
	phkResult PHKEY, lpdwDisposition types.LPDWORD) types.LSTATUS {
	a1, err := syscall.UTF16PtrFromString(lpSubKey)
	if err != nil {
		return types.LSTATUS(types.ERROR_INVALID_PARAMETER)
	}

	a2, err := syscall.UTF16PtrFromString(lpClass)
	if err != nil {
		return types.LSTATUS(types.ERROR_INVALID_PARAMETER)
	}

	ret, _, _ := advapi32.ProcRegCreateKeyEx.Call(
		uintptr(hKey),
		uintptr(unsafe.Pointer(a1)),
		uintptr(Reserved),
		uintptr(unsafe.Pointer(a2)),
		uintptr(dwOptions),
		uintptr(samDesired),
		uintptr(unsafe.Pointer(lpSecurityAttributes)),
		uintptr(unsafe.Pointer(phkResult)),
		uintptr(unsafe.Pointer(lpdwDisposition)),
	)

	return types.LSTATUS(ret)
}
