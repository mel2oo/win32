//go:build windows
// +build windows

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/

package advapi32

import (
	"syscall"
	"unsafe"

	"github.com/mel2oo/win32/typedef"
)

type REGSAM typedef.DWORD

const (
	REGSAM_DELETE                 REGSAM = 0x00010000
	REGSAM_READ_CONTROL           REGSAM = 0x00020000
	REGSAM_WRITE_DAC              REGSAM = 0x00040000
	REGSAM_WRITE_OWNER            REGSAM = 0x00080000
	REGSAM_ACCESS_SYSTEM_SECURITY REGSAM = 0x01000000
	REGSAM_KEY_QUERY_VALUE        REGSAM = 0x0001
	REGSAM_KEY_SET_VALUE          REGSAM = 0x0002
	REGSAM_KEY_CREATE_SUB_KEY     REGSAM = 0x0004
	REGSAM_KEY_ENUMERATE_SUB_KEYS REGSAM = 0x0008
	REGSAM_KEY_NOTIFY             REGSAM = 0x0010
	REGSAM_KEY_CREATE_LINK        REGSAM = 0x0020
	REGSAM_KEY_WOW64_32KEY        REGSAM = 0x0200
	REGSAM_KEY_WOW64_64KEY        REGSAM = 0x0100
	REGSAM_KEY_WOW64_RES          REGSAM = 0x0300
	REGSAM_KEY_READ               REGSAM = 0x20019
	REGSAM_KEY_WRITE              REGSAM = 0x20006
	REGSAM_KEY_ALL_ACCESS         REGSAM = 0xF003F
	REGSAM_MAXIMUM_ALLOWED        REGSAM = 0x02000000
)

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-abortsystemshutdownw
func AbortSystemShutdown(lpMachineName string) typedef.BOOL {
	a1, err := syscall.UTF16PtrFromString(lpMachineName)
	if err != nil {
		return typedef.FALSE
	}

	ret, _, _ := procAbortSystemShutdown.Call(uintptr(unsafe.Pointer(a1)))

	return typedef.BOOL(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-initiateshutdownw
func InitiateShutdown(lpMachineName, lpMessage string, dwGracePeriod, dwShutdownFlags, dwReason typedef.DWORD) typedef.DWORD {
	a1, err := syscall.UTF16PtrFromString(lpMachineName)
	if err != nil {
		return typedef.DWORD(typedef.ERROR_INVALID_PARAMETER)
	}

	a2, err := syscall.UTF16PtrFromString(lpMessage)
	if err != nil {
		return typedef.DWORD(typedef.ERROR_INVALID_PARAMETER)
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
		return typedef.FALSE
	}

	a2, err := syscall.UTF16PtrFromString(lpMessage)
	if err != nil {
		return typedef.FALSE
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

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-initiatesystemshutdownexw
func InitiateSystemShutdownEx(lpMachineName, lpMessage string, dwTimeout typedef.DWORD,
	bForceAppsClosed, bRebootAfterShutdown typedef.BOOL, dwReason typedef.DWORD) typedef.BOOL {
	a1, err := syscall.UTF16PtrFromString(lpMachineName)
	if err != nil {
		return typedef.FALSE
	}

	a2, err := syscall.UTF16PtrFromString(lpMessage)
	if err != nil {
		return typedef.FALSE
	}

	ret, _, _ := procInitiateSystemShutdownEx.Call(
		uintptr(unsafe.Pointer(a1)),
		uintptr(unsafe.Pointer(a2)),
		uintptr(dwTimeout),
		uintptr(bForceAppsClosed),
		uintptr(bRebootAfterShutdown),
		uintptr(dwReason),
	)

	return typedef.BOOL(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regclosekey
func RegCloseKey(hKey typedef.HKEY) typedef.LSTATUS {
	ret, _, _ := procRegCloseKey.Call(uintptr(hKey))
	return typedef.LSTATUS(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regconnectregistryw
func RegConnectRegistry(lpMachineName string, hKey typedef.HKEY, phkResult *typedef.HKEY) typedef.LSTATUS {
	a1, err := syscall.UTF16PtrFromString(lpMachineName)
	if err != nil {
		return typedef.LSTATUS(typedef.ERROR_INVALID_PARAMETER)
	}

	ret, _, _ := procRegConnectRegistry.Call(
		uintptr(unsafe.Pointer(a1)),
		uintptr(hKey),
		uintptr(unsafe.Pointer(phkResult)),
	)

	return typedef.LSTATUS(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regcopytreew
func RegCopyTree(hKeySrc typedef.HKEY, lpSubKey string, hKeyDest typedef.HKEY) typedef.LSTATUS {
	a1, err := syscall.UTF16PtrFromString(lpSubKey)
	if err != nil {
		return typedef.LSTATUS(typedef.ERROR_INVALID_PARAMETER)
	}

	ret, _, _ := procRegCopyTree.Call(
		uintptr(hKeySrc),
		uintptr(unsafe.Pointer(a1)),
		uintptr(hKeyDest),
	)

	return typedef.LSTATUS(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regcreatekeyexw
func RegCreateKeyEx(
	hKey typedef.HKEY, lpSubKey string,
	Reserved typedef.DWORD, lpClass string,
	dwOptions typedef.DWORD, samDesired REGSAM,
	lpSecurityAttributes typedef.LPSECURITY_ATTRIBUTES,
	phkResult typedef.PHKEY, lpdwDisposition typedef.LPDWORD) typedef.LSTATUS {
	a1, err := syscall.UTF16PtrFromString(lpSubKey)
	if err != nil {
		return typedef.LSTATUS(typedef.ERROR_INVALID_PARAMETER)
	}

	a2, err := syscall.UTF16PtrFromString(lpClass)
	if err != nil {
		return typedef.LSTATUS(typedef.ERROR_INVALID_PARAMETER)
	}

	ret, _, _ := procRegCreateKeyEx.Call(
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

	return typedef.LSTATUS(ret)
}
