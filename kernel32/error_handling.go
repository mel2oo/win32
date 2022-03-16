package kernel32

import (
	"github.com/mel2oo/win32/typedef"
)

func GetLastError() typedef.ERROR_CODE {
	ret, _, _ := procGetLastError.Call()
	return typedef.ERROR_CODE(ret)
}
