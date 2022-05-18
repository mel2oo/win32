package kernel32

import (
	"github.com/mel2oo/win32/types"
)

// https://docs.microsoft.com/en-us/windows/win32/api/errhandlingapi/nf-errhandlingapi-getlasterror
func GetLastError() types.ErrorCode {
	ret, _, _ := procGetLastError.Call()
	return types.ErrorCode(ret)
}
