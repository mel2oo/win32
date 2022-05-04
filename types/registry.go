package types

// HKEY
type HKEY HANDLE

const (
	HKEY_CLASSES_ROOT                = -2147483648
	HKEY_CURRENT_USER                = -2147483647
	HKEY_LOCAL_MACHINE               = -2147483646
	HKEY_USERS                       = -2147483645
	HKEY_PERFORMANCE_DATA            = -2147483644
	HKEY_PERFORMANCE_TEXT            = -2147483568
	HKEY_PERFORMANCE_NLSTEXT         = -2147483552
	HKEY_CURRENT_CONFIG              = -2147483643
	HKEY_DYN_DATA                    = -2147483642
	HKEY_CURRENT_USER_LOCAL_SETTINGS = -2147483641
)

type PHKEY *HKEY

// REGSAM
type REGSAM DWORD

const (
	KEY_DELETE                 REGSAM = 0x00010000
	KEY_READ_CONTROL           REGSAM = 0x00020000
	KEY_WRITE_DAC              REGSAM = 0x00040000
	KEY_WRITE_OWNER            REGSAM = 0x00080000
	KEY_ACCESS_SYSTEM_SECURITY REGSAM = 0x01000000
	KEY_QUERY_VALUE            REGSAM = 0x0001
	KEY_SET_VALUE              REGSAM = 0x0002
	KEY_CREATE_SUB_KEY         REGSAM = 0x0004
	KEY_ENUMERATE_SUB_KEYS     REGSAM = 0x0008
	KEY_NOTIFY                 REGSAM = 0x0010
	KEY_CREATE_LINK            REGSAM = 0x0020
	KEY_WOW64_32KEY            REGSAM = 0x0200
	KEY_WOW64_64KEY            REGSAM = 0x0100
	KEY_WOW64_RES              REGSAM = 0x0300
	KEY_READ                   REGSAM = 0x20019
	KEY_WRITE                  REGSAM = 0x20006
	KEY_ALL_ACCESS             REGSAM = 0xF003F
	KEY_MAXIMUM_ALLOWED        REGSAM = 0x02000000
)

// RegType
type RegType DWORD

const (
	REG_NONE                       RegType = 0
	REG_SZ                         RegType = 1
	REG_EXPAND_SZ                  RegType = 2
	REG_BINARY                     RegType = 3
	REG_DWORD                      RegType = 4
	REG_DWORD_BIG_ENDIAN           RegType = 5
	REG_LINK                       RegType = 6
	REG_MULTI_SZ                   RegType = 7
	REG_RESOURCE_LIST              RegType = 8
	REG_FULL_RESOURCE_DESCRIPTOR   RegType = 9
	REG_RESOURCE_REQUIREMENTS_LIST RegType = 10
	REG_QWORD                      RegType = 11
)

type RegTypeULONG RegType
type RegTypePULONG *RegTypeULONG

// RegOptions
type RegOptions DWORD

const (
	REG_OPTION_VOLATILE       RegOptions = 0x00000001
	REG_OPTION_CREATE_LINK    RegOptions = 0x00000002
	REG_OPTION_BACKUP_RESTORE RegOptions = 0x00000004
	REG_OPTION_OPEN_LINK      RegOptions = 0x00000008
)

// RegDisposition
type RegDisposition DWORD

const (
	REG_CREATED_NEW_KEY     RegDisposition = 0x00000001
	REG_OPENED_EXISTING_KEY RegDisposition = 0x00000002
)

type RegDispositionULONG RegDisposition
type RegDispositionPULONG *RegDispositionULONG

// RegMuiFlags
type RegMuiFlags DWORD

const (
	REG_MUI_STRING_TRUNCATE RegMuiFlags = 0x00000001
)

// RegNotifyFlags
type RegNotifyFlags DWORD

const (
	REG_NOTIFY_CHANGE_NAME       RegNotifyFlags = 0x00000001
	REG_NOTIFY_CHANGE_ATTRIBUTES RegNotifyFlags = 0x00000002
	REG_NOTIFY_CHANGE_LAST_SET   RegNotifyFlags = 0x00000004
	REG_NOTIFY_CHANGE_SECURITY   RegNotifyFlags = 0x00000008
	REG_NOTIFY_THREAD_AGNOSTIC   RegNotifyFlags = 0x10000000
)
