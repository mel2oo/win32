package kernel32

import (
	"github.com/switch-li/win32/typedef"
)

func GetLastError() typedef.ERROR_CODE {
	ret, _, _ := procGetLastError.Call()
	return typedef.ERROR_CODE(ret)
}
