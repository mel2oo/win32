package kernel32

import (
	"github.com/mel2oo/win32/types"
)

// https://docs.microsoft.com/en-us/windows/win32/api/errhandlingapi/nf-errhandlingapi-getlasterror
func GetLastError() types.ERROR_CODE {
	ret, _, _ := procGetLastError.Call()
	return types.ERROR_CODE(ret)
}
