package kernel32

import (
	"github.com/mel2oo/win32/typedef"
)

// https://docs.microsoft.com/en-us/windows/win32/api/errhandlingapi/nf-errhandlingapi-getlasterror
func GetLastError() typedef.ERROR_CODE {
	ret, _, _ := procGetLastError.Call()
	return typedef.ERROR_CODE(ret)
}
