// https://docs.microsoft.com/en-us/windows/win32/api/ip2string/
package ip2string

import (
	"syscall"
	"unsafe"

	"github.com/mel2oo/win32/ntdll"
)

// https://docs.microsoft.com/en-us/windows/win32/api/ip2string/nf-ip2string-rtlipv6addresstostringw
func RtlIpv6AddressToString(buffer []byte, ipv6 []uint16) string {
	_, _, _ = ntdll.ProcRtlIpv6AddressToString.Call(
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(unsafe.Pointer(&ipv6[0])),
	)

	return syscall.UTF16ToString(ipv6)
}
