package advapi32

import "win32/typedef"

type SC_HANDLE typedef.HANDLE
type SCManagerAccess typedef.DWORD

const (
	GENERIC_READ                  SCManagerAccess = 0x80000000
	GENERIC_WRITE                 SCManagerAccess = 0x40000000
	GENERIC_EXECUTE               SCManagerAccess = 0x20000000
	SC_MANAGER_CONNECT            SCManagerAccess = 0x0001
	SC_MANAGER_CREATE_SERVICE     SCManagerAccess = 0x0002
	SC_MANAGER_ENUMERATE_SERVICE  SCManagerAccess = 0x0004
	SC_MANAGER_LOCK               SCManagerAccess = 0x0008
	SC_MANAGER_QUERY_LOCK_STATUS  SCManagerAccess = 0x0010
	SC_MANAGER_MODIFY_BOOT_CONFIG SCManagerAccess = 0x0020
	SC_MANAGER_ALL_ACCESS         SCManagerAccess = 0xF003F
)
