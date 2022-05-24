// https://docs.microsoft.com/en-us/windows/win32/api/winbase/

package winbase

import (
	"unsafe"

	"github.com/mel2oo/win32/advapi32"
	"github.com/mel2oo/win32/types"
)

func LookupAccountSid(lpSystemName types.LPCWSTR, Sid []byte,
	Name types.LPWSTR, cchName types.LPDWORD, ReferencedDomainName types.LPWSTR,
	cchReferencedDomainName types.LPDWORD, peUs types.PSID_NAME_USE) types.BOOL {
	ret, _, _ := advapi32.ProcLookupAccountSid.Call(
		uintptr(unsafe.Pointer(lpSystemName)),
		uintptr(unsafe.Pointer(&Sid[0])),
		uintptr(unsafe.Pointer(Name)),
		uintptr(unsafe.Pointer(cchName)),
		uintptr(unsafe.Pointer(ReferencedDomainName)),
		uintptr(unsafe.Pointer(cchReferencedDomainName)),
		uintptr(unsafe.Pointer(peUs)),
	)
	return types.BOOL(ret)
}
