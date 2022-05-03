package types

import "syscall"

type (
	HINSTANCE                   HMODULE
	WPARAM                      UINT_PTR
	LPARAM                      LONG_PTR
	LRESULT                     LONG_PTR
	BSTR                        PWCHAR
	HDEVINFO                    PVOID
	PIO_APC_ROUTINE             LPVOID
	FARPROC                     LPVOID
	PSID                        PVOID
	PVECTORED_EXCEPTION_HANDLER LPVOID
	ATOM                        WORD
	PSLIST_ENTRY                LPVOID
	KAFFINITY                   ULONG_PTR
	LPTHREAD_START_ROUTINE      LPVOID
	STDAPI                      HRESULT
	SCODE                       HRESULT
)

// Wireless LAN Variables
type (
	WLAN_REASON_CODE  DWORD
	PWLAN_REASON_CODE *WLAN_REASON_CODE
)

// LANGID
type LANGID WORD

const (
	LANG_SYSTEM_DEFAULT LANGID = 0x0800
	LANG_USER_DEFAULT   LANGID = 0x0400
)

// LCID
type (
	LCID  DWORD
	PLCID *LCID
)

const (
	LOCALE_SYSTEM_DEFAULT     LCID = 0x0800
	LOCALE_USER_DEFAULT       LCID = 0x0400
	LOCALE_CUSTOM_DEFAULT     LCID = 0x0c00
	LOCALE_CUSTOM_UNSPECIFIED LCID = 0x1000
	LOCALE_CUSTOM_UI_DEFAULT  LCID = 0x1400
	LOCALE_NEUTRAL            LCID = 0x0000
	LOCALE_INVARIANT          LCID = 0x007f
)

// LIST_ENTRY
type LIST_ENTRY struct {
	Flink PSLIST_ENTRY
	Blink PSLIST_ENTRY
}
type PLIST_ENTRY LIST_ENTRY

// Unions, Bitfields, Structures
type PSLIST_HEADER LPVOID

// Structures - 32/64 Differences
type PCONTEXT LPVOID

// POINT
type POINT struct {
	X LONG
	Y LONG
}

type (
	LPPOINT *POINT
	POINTL  POINT
)

// DMORIENT
type DMORIENT SHORT

const (
	DMORIENT_PORTRAIT  DMORIENT = 1
	DMORIENT_LANDSCAPE DMORIENT = 2
)

// DMPAPER
type DMPAPER SHORT

const (
	DMPAPER_LETTER                        DMPAPER = 1
	DMPAPER_LETTERSMALL                   DMPAPER = 2
	DMPAPER_TABLOID                       DMPAPER = 3
	DMPAPER_LEDGER                        DMPAPER = 4
	DMPAPER_LEGAL                         DMPAPER = 5
	DMPAPER_STATEMENT                     DMPAPER = 6
	DMPAPER_EXECUTIVE                     DMPAPER = 7
	DMPAPER_A3                            DMPAPER = 8
	DMPAPER_A4                            DMPAPER = 9
	DMPAPER_A4SMALL                       DMPAPER = 10
	DMPAPER_A5                            DMPAPER = 11
	DMPAPER_B4                            DMPAPER = 12
	DMPAPER_B5                            DMPAPER = 13
	DMPAPER_FOLIO                         DMPAPER = 14
	DMPAPER_QUARTO                        DMPAPER = 15
	DMPAPER_10X14                         DMPAPER = 16
	DMPAPER_11X17                         DMPAPER = 17
	DMPAPER_NOTE                          DMPAPER = 18
	DMPAPER_ENV_9                         DMPAPER = 19
	DMPAPER_ENV_10                        DMPAPER = 20
	DMPAPER_ENV_11                        DMPAPER = 21
	DMPAPER_ENV_12                        DMPAPER = 22
	DMPAPER_ENV_14                        DMPAPER = 23
	DMPAPER_CSHEET                        DMPAPER = 24
	DMPAPER_DSHEET                        DMPAPER = 25
	DMPAPER_ESHEET                        DMPAPER = 26
	DMPAPER_ENV_DL                        DMPAPER = 27
	DMPAPER_ENV_C5                        DMPAPER = 28
	DMPAPER_ENV_C3                        DMPAPER = 29
	DMPAPER_ENV_C4                        DMPAPER = 30
	DMPAPER_ENV_C6                        DMPAPER = 31
	DMPAPER_ENV_C65                       DMPAPER = 32
	DMPAPER_ENV_B4                        DMPAPER = 33
	DMPAPER_ENV_B5                        DMPAPER = 34
	DMPAPER_ENV_B6                        DMPAPER = 35
	DMPAPER_ENV_ITALY                     DMPAPER = 36
	DMPAPER_ENV_MONARCH                   DMPAPER = 37
	DMPAPER_ENV_PERSONAL                  DMPAPER = 38
	DMPAPER_FANFOLD_US                    DMPAPER = 39
	DMPAPER_FANFOLD_STD_GERMAN            DMPAPER = 40
	DMPAPER_FANFOLD_LGL_GERMAN            DMPAPER = 41
	DMPAPER_ISO_B4                        DMPAPER = 42
	DMPAPER_JAPANESE_POSTCARD             DMPAPER = 43
	DMPAPER_9X11                          DMPAPER = 44
	DMPAPER_10X11                         DMPAPER = 45
	DMPAPER_15X11                         DMPAPER = 46
	DMPAPER_ENV_INVITE                    DMPAPER = 47
	DMPAPER_RESERVED_48                   DMPAPER = 48
	DMPAPER_RESERVED_49                   DMPAPER = 49
	DMPAPER_LETTER_EXTRA                  DMPAPER = 50
	DMPAPER_LEGAL_EXTRA                   DMPAPER = 51
	DMPAPER_TABLOID_EXTRA                 DMPAPER = 52
	DMPAPER_A4_EXTRA                      DMPAPER = 53
	DMPAPER_LETTER_TRANSVERSE             DMPAPER = 54
	DMPAPER_A4_TRANSVERSE                 DMPAPER = 55
	DMPAPER_LETTER_EXTRA_TRANSVERSE       DMPAPER = 56
	DMPAPER_A_PLUS                        DMPAPER = 57
	DMPAPER_B_PLUS                        DMPAPER = 58
	DMPAPER_LETTER_PLUS                   DMPAPER = 59
	DMPAPER_A4_PLUS                       DMPAPER = 60
	DMPAPER_A5_TRANSVERSE                 DMPAPER = 61
	DMPAPER_B5_TRANSVERSE                 DMPAPER = 62
	DMPAPER_A3_EXTRA                      DMPAPER = 63
	DMPAPER_A5_EXTRA                      DMPAPER = 64
	DMPAPER_B5_EXTRA                      DMPAPER = 65
	DMPAPER_A2                            DMPAPER = 66
	DMPAPER_A3_TRANSVERSE                 DMPAPER = 67
	DMPAPER_A3_EXTRA_TRANSVERSE           DMPAPER = 68
	DMPAPER_DBL_JAPANESE_POSTCARD         DMPAPER = 69
	DMPAPER_A6                            DMPAPER = 70
	DMPAPER_JENV_KAKU2                    DMPAPER = 71
	DMPAPER_JENV_KAKU3                    DMPAPER = 72
	DMPAPER_JENV_CHOU3                    DMPAPER = 73
	DMPAPER_JENV_CHOU4                    DMPAPER = 74
	DMPAPER_LETTER_ROTATED                DMPAPER = 75
	DMPAPER_A3_ROTATED                    DMPAPER = 76
	DMPAPER_A4_ROTATED                    DMPAPER = 77
	DMPAPER_A5_ROTATED                    DMPAPER = 78
	DMPAPER_B4_JIS_ROTATED                DMPAPER = 79
	DMPAPER_B5_JIS_ROTATED                DMPAPER = 80
	DMPAPER_JAPANESE_POSTCARD_ROTATED     DMPAPER = 81
	DMPAPER_DBL_JAPANESE_POSTCARD_ROTATED DMPAPER = 82
	DMPAPER_A6_ROTATED                    DMPAPER = 83
	DMPAPER_JENV_KAKU2_ROTATED            DMPAPER = 84
	DMPAPER_JENV_KAKU3_ROTATED            DMPAPER = 85
	DMPAPER_JENV_CHOU3_ROTATED            DMPAPER = 86
	DMPAPER_JENV_CHOU4_ROTATED            DMPAPER = 87
	DMPAPER_B6_JIS                        DMPAPER = 88
	DMPAPER_B6_JIS_ROTATED                DMPAPER = 89
	DMPAPER_12X11                         DMPAPER = 90
	DMPAPER_JENV_YOU4                     DMPAPER = 91
	DMPAPER_JENV_YOU4_ROTATED             DMPAPER = 92
	DMPAPER_P16K                          DMPAPER = 93
	DMPAPER_P32K                          DMPAPER = 94
	DMPAPER_P32KBIG                       DMPAPER = 95
	DMPAPER_PENV_1                        DMPAPER = 96
	DMPAPER_PENV_2                        DMPAPER = 97
	DMPAPER_PENV_3                        DMPAPER = 98
	DMPAPER_PENV_4                        DMPAPER = 99
	DMPAPER_PENV_5                        DMPAPER = 100
	DMPAPER_PENV_6                        DMPAPER = 101
	DMPAPER_PENV_7                        DMPAPER = 102
	DMPAPER_PENV_8                        DMPAPER = 103
	DMPAPER_PENV_9                        DMPAPER = 104
	DMPAPER_PENV_10                       DMPAPER = 105
	DMPAPER_P16K_ROTATED                  DMPAPER = 106
	DMPAPER_P32K_ROTATED                  DMPAPER = 107
	DMPAPER_P32KBIG_ROTATED               DMPAPER = 108
	DMPAPER_PENV_1_ROTATED                DMPAPER = 109
	DMPAPER_PENV_2_ROTATED                DMPAPER = 110
	DMPAPER_PENV_3_ROTATED                DMPAPER = 111
	DMPAPER_PENV_4_ROTATED                DMPAPER = 112
	DMPAPER_PENV_5_ROTATED                DMPAPER = 113
	DMPAPER_PENV_6_ROTATED                DMPAPER = 114
	DMPAPER_PENV_7_ROTATED                DMPAPER = 115
	DMPAPER_PENV_8_ROTATED                DMPAPER = 116
	DMPAPER_PENV_9_ROTATED                DMPAPER = 117
	DMPAPER_PENV_10_ROTATED               DMPAPER = 118
)

// DMBIN
type DMBIN SHORT

const (
	DMBIN_UPPER         DMBIN = 1
	DMBIN_LOWER         DMBIN = 2
	DMBIN_MIDDLE        DMBIN = 3
	DMBIN_MANUAL        DMBIN = 4
	DMBIN_ENVELOPE      DMBIN = 5
	DMBIN_ENVMANUAL     DMBIN = 6
	DMBIN_AUTO          DMBIN = 7
	DMBIN_TRACTOR       DMBIN = 8
	DMBIN_SMALLFMT      DMBIN = 9
	DMBIN_LARGEFMT      DMBIN = 10
	DMBIN_LARGECAPACITY DMBIN = 11
	DMBIN_CASSETTE      DMBIN = 14
	DMBIN_FORMSOURCE    DMBIN = 15
)

// DMRES
type DMRES SHORT

const (
	DMRES_DRAFT  DMRES = -1
	DMRES_LOW    DMRES = -2
	DMRES_MEDIUM DMRES = -3
	DMRES_HIGH   DMRES = -4
)

// DEVMODE_u1_s1
type DEVMODE_u1_s1 struct {
	DmOrientation   DMORIENT
	DmPaperSize     DMPAPER
	DmPaperLength   SHORT
	DmPaperWidth    SHORT
	DmScale         SHORT
	DmCopies        SHORT
	DmDefaultSource DMBIN
	DmPrintQuality  DMRES
}

// DMDO
type DMDO DWORD

const (
	DMDO_DEFAULT DMDO = 0
	DMDO_90      DMDO = 1
	DMDO_180     DMDO = 2
	DMDO_270     DMDO = 3
)

// DMDFO
type DMDFO DWORD

const (
	DMDFO_DEFAULT DMDFO = 0
	DMDFO_STRETCH DMDFO = 1
	DMDFO_CENTER  DMDFO = 2
)

// DEVMODE_u1_s2
type DEVMODE_u1_s2 struct {
	DmPosition           POINTL
	DmDisplayOrientation DMDO
	DmDisplayFixedOutput DMDFO
}

// DEVMODE_u1
type DEVMODE_u1 struct {
	DEVMODE_u1_s1
	DEVMODE_u1_s2
}

// DMDISPLAYFLAGS
type DMDISPLAYFLAGS DWORD

const (
	DM_GRAYSCALE            DMDISPLAYFLAGS = 0x00000001
	DM_INTERLACED           DMDISPLAYFLAGS = 0x00000002
	DMDISPLAYFLAGS_TEXTMODE DMDISPLAYFLAGS = 0x00000004
)

// DMNUP
type DMNUP DWORD

const (
	DMNUP_SYSTEM DMNUP = 1
	DMNUP_ONEUP  DMNUP = 2
)

// DEVMODE_u2
type DEVMODE_u2 struct {
	DmDisplayFlags DMDISPLAYFLAGS
	DmNup          DMNUP
}

// DM_Fields
type DM_Fields DWORD

const (
	DM_ORIENTATION        DM_Fields = 0x00000001
	DM_PAPERSIZE          DM_Fields = 0x00000002
	DM_PAPERLENGTH        DM_Fields = 0x00000004
	DM_PAPERWIDTH         DM_Fields = 0x00000008
	DM_SCALE              DM_Fields = 0x00000010
	DM_POSITION           DM_Fields = 0x00000020
	DM_NUP                DM_Fields = 0x00000040
	DM_DISPLAYORIENTATION DM_Fields = 0x00000080
	DM_COPIES             DM_Fields = 0x00000100
	DM_DEFAULTSOURCE      DM_Fields = 0x00000200
	DM_PRINTQUALITY       DM_Fields = 0x00000400
	DM_COLOR              DM_Fields = 0x00000800
	DM_DUPLEX             DM_Fields = 0x00001000
	DM_YRESOLUTION        DM_Fields = 0x00002000
	DM_TTOPTION           DM_Fields = 0x00004000
	DM_COLLATE            DM_Fields = 0x00008000
	DM_FORMNAME           DM_Fields = 0x00010000
	DM_LOGPIXELS          DM_Fields = 0x00020000
	DM_BITSPERPEL         DM_Fields = 0x00040000
	DM_PELSWIDTH          DM_Fields = 0x00080000
	DM_PELSHEIGHT         DM_Fields = 0x00100000
	DM_DISPLAYFLAGS       DM_Fields = 0x00200000
	DM_DISPLAYFREQUENCY   DM_Fields = 0x00400000
	DM_ICMMETHOD          DM_Fields = 0x00800000
	DM_ICMINTENT          DM_Fields = 0x01000000
	DM_MEDIATYPE          DM_Fields = 0x02000000
	DM_DITHERTYPE         DM_Fields = 0x04000000
	DM_PANNINGWIDTH       DM_Fields = 0x08000000
	DM_PANNINGHEIGHT      DM_Fields = 0x10000000
	DM_DISPLAYFIXEDOUTPUT DM_Fields = 0x20000000
)

// DMCOLOR
type DMCOLOR SHORT

const (
	DMCOLOR_MONOCHROME DMCOLOR = 1
	DMCOLOR_COLOR      DMCOLOR = 2
)

// DMDUP
type DMDUP SHORT

const (
	DMDUP_SIMPLEX    DMDUP = 1
	DMDUP_VERTICAL   DMDUP = 2
	DMDUP_HORIZONTAL DMDUP = 3
)

// DMTT
type DMTT SHORT

const (
	DMTT_BITMAP           DMTT = 1
	DMTT_DOWNLOAD         DMTT = 2
	DMTT_SUBDEV           DMTT = 3
	DMTT_DOWNLOAD_OUTLINE DMTT = 4
)

// DMCOLLATE
type DMCOLLATE SHORT

const (
	DMCOLLATE_FALSE DMCOLLATE = 0
	DMCOLLATE_TRUE  DMCOLLATE = 1
)

// DMICMMETHOD
type DMICMMETHOD DWORD

const (
	DMICMMETHOD_NONE   DMICMMETHOD = 1
	DMICMMETHOD_SYSTEM DMICMMETHOD = 2
	DMICMMETHOD_DRIVER DMICMMETHOD = 3
	DMICMMETHOD_DEVICE DMICMMETHOD = 4
)

// DMICM
type DMICM DWORD

const (
	DMICM_SATURATE         DMICM = 1
	DMICM_CONTRAST         DMICM = 2
	DMICM_COLORIMETRIC     DMICM = 3
	DMICM_ABS_COLORIMETRIC DMICM = 4
)

// DMMEDIA
type DMMEDIA DWORD

const (
	DMMEDIA_STANDARD     DMMEDIA = 1
	DMMEDIA_TRANSPARENCY DMMEDIA = 2
	DMMEDIA_GLOSSY       DMMEDIA = 3
)

// DMDITHER
type DMDITHER DWORD

const (
	DMDITHER_NONE           DMDITHER = 1
	DMDITHER_COARSE         DMDITHER = 2
	DMDITHER_FINE           DMDITHER = 3
	DMDITHER_LINEART        DMDITHER = 4
	DMDITHER_ERRORDIFFUSION DMDITHER = 5
	DMDITHER_RESERVED6      DMDITHER = 6
	DMDITHER_RESERVED7      DMDITHER = 7
	DMDITHER_RESERVED8      DMDITHER = 8
	DMDITHER_RESERVED9      DMDITHER = 9
	DMDITHER_GRAYSCALE      DMDITHER = 10
)

// DEVMODE
type DEVMODE struct {
	DmDeviceName       [32]WCHAR
	DmSpecVersion      WORD
	DmDriverVersion    WORD
	DmSize             WORD
	DmDriverExtra      WORD
	DmFields           DM_Fields
	DEVMODE_u1         [12]byte
	DmColor            DMCOLOR
	DmDuplex           DMDUP
	DmYResolution      SHORT
	DmTTOption         DMTT
	DmCollate          DMCOLLATE
	DmFormName         [32]WCHAR
	DmLogPixels        WORD
	DmBitsPerPel       DWORD
	DmPelsWidth        DWORD
	DmPelsHeight       DWORD
	DEVMODE_u2         [4]byte
	DmDisplayFrequency DWORD
	DmICMMethod        DMICMMETHOD
	DmICMIntent        DMICM
	DmMediaType        DMMEDIA
	DmDitherType       DMDITHER
	DmReserved1        DWORD
	DmReserved2        DWORD
	DmPanningWidth     DWORD
	DmPanningHeight    DWORD
}

// HANDLE
type (
	HANDLE   UINT_PTR
	PHANDLE  *HANDLE
	LPHANDLE *HANDLE

	FILE_HANDLE  HANDLE
	HACCEL       HANDLE
	HRSRC        HANDLE
	HWND         HANDLE
	HGLOBAL      HANDLE
	HKL          HANDLE
	HBITMAP      HANDLE
	HENHMETAFILE HANDLE

	ProcessHandle HANDLE
	ThreadHandle  HANDLE
)

// LARGE_INTEGER
type LARGE_INTEGER_s struct {
	LowPart  DWORD
	HighPart LONG
}

type (
	LARGE_INTEGER  LARGE_INTEGER_s
	PLARGE_INTEGER *LARGE_INTEGER
)

// ULARGE_INTEGER
type ULARGE_INTEGER_s struct {
	LowPart  DWORD
	HighPart LONG
}

type (
	ULARGE_INTEGER  ULARGE_INTEGER_s
	PULARGE_INTEGER *ULARGE_INTEGER
)

// FILETIME
type FILETIME struct {
	DwLowDateTime  DWORD
	DwHighDateTime DWORD
}

type (
	PFILETIME  *FILETIME
	LPFILETIME *FILETIME
)

type (
	// GUID
	GUID    syscall.GUID
	PGUID   *GUID
	LPCGUID *GUID
	LPGUID  *GUID
	REFGUID *GUID

	// CLSID
	CLSID    GUID
	LPCLSID  *CLSID
	REFCLSID *CLSID

	// UUID
	UUID GUID

	// IID
	IID    GUID
	REFIID *IID
	LPIID  *IID

	// LUID
	LUID struct {
		LowPart  DWORD
		HighPart LONG
	}
	PLUID *LUID

	// SLID
	SLID GUID
)

// ACCESS_MASK
type (
	ACCESS_MASK         DWORD
	PACCESS_MASK        *ACCESS_MASK
	ACCESS_MASK_DWORD   ACCESS_MASK
	ACCESS_MASK_LPDWORD *ACCESS_MASK
)

const (
	GENERIC_READ             ACCESS_MASK = 0x80000000
	GENERIC_WRITE            ACCESS_MASK = 0x40000000
	GENERIC_EXECUTE          ACCESS_MASK = 0x20000000
	GENERIC_ALL              ACCESS_MASK = 0x10000000
	MAXIMUM_ALLOWED          ACCESS_MASK = 0x02000000
	ACCESS_SYSTEM_SECURITY   ACCESS_MASK = 0x01000000
	SYNCHRONIZE              ACCESS_MASK = 0x00100000
	STANDARD_RIGHTS_REQUIRED ACCESS_MASK = 0x000F0000
	STANDARD_RIGHTS_ALL      ACCESS_MASK = 0x001F0000
	WRITE_OWNER              ACCESS_MASK = 0x00080000
	WRITE_DAC                ACCESS_MASK = 0x00040000
	READ_CONTROL             ACCESS_MASK = 0x00020000
	DELETE                   ACCESS_MASK = 0x00010000
	SPECIFIC_RIGHTS_ALL      ACCESS_MASK = 0x0000FFFF
)

// THREAD_ACCESS_MASK
type THREAD_ACCESS_MASK ACCESS_MASK

const (
	THREAD_TERMINATE                 THREAD_ACCESS_MASK = 0x0001
	THREAD_SUSPEND_RESUME            THREAD_ACCESS_MASK = 0x0002
	THREAD_GET_CONTEXT               THREAD_ACCESS_MASK = 0x0008
	THREAD_SET_CONTEXT               THREAD_ACCESS_MASK = 0x0010
	THREAD_QUERY_INFORMATION         THREAD_ACCESS_MASK = 0x0040
	THREAD_SET_INFORMATION           THREAD_ACCESS_MASK = 0x0020
	THREAD_SET_THREAD_TOKEN          THREAD_ACCESS_MASK = 0x0080
	THREAD_IMPERSONATE               THREAD_ACCESS_MASK = 0x0100
	THREAD_DIRECT_IMPERSONATION      THREAD_ACCESS_MASK = 0x0200
	THREAD_SET_LIMITED_INFORMATION   THREAD_ACCESS_MASK = 0x0400
	THREAD_QUERY_LIMITED_INFORMATION THREAD_ACCESS_MASK = 0x0800
	THREAD_ALL_ACCESS                THREAD_ACCESS_MASK = 0x1FFFFF
)

// PROCESS_ACCESS_MASK
type PROCESS_ACCESS_MASK ACCESS_MASK

const (
	PROCESS_TERMINATE                 PROCESS_ACCESS_MASK = 0x0001
	PROCESS_CREATE_THREAD             PROCESS_ACCESS_MASK = 0x0002
	PROCESS_SET_SESSIONID             PROCESS_ACCESS_MASK = 0x0004
	PROCESS_VM_OPERATION              PROCESS_ACCESS_MASK = 0x0008
	PROCESS_VM_READ                   PROCESS_ACCESS_MASK = 0x0010
	PROCESS_VM_WRITE                  PROCESS_ACCESS_MASK = 0x0020
	PROCESS_DUP_HANDLE                PROCESS_ACCESS_MASK = 0x0040
	PROCESS_CREATE_PROCESS            PROCESS_ACCESS_MASK = 0x0080
	PROCESS_SET_QUOTA                 PROCESS_ACCESS_MASK = 0x0100
	PROCESS_SET_INFORMATION           PROCESS_ACCESS_MASK = 0x0200
	PROCESS_QUERY_INFORMATION         PROCESS_ACCESS_MASK = 0x0400
	PROCESS_SUSPEND_RESUME            PROCESS_ACCESS_MASK = 0x0800
	PROCESS_QUERY_LIMITED_INFORMATION PROCESS_ACCESS_MASK = 0x1000
	PROCESS_ALL_ACCESS                PROCESS_ACCESS_MASK = 0x1fffff
)

// FILE_ACCESS_MASK
type FILE_ACCESS_MASK ACCESS_MASK

const (
	FILE_WRITE_DATA           FILE_ACCESS_MASK = 0x0002
	FILE_ADD_FILE             FILE_ACCESS_MASK = 0x0002
	FILE_APPEND_DATA          FILE_ACCESS_MASK = 0x0004
	FILE_ADD_SUBDIRECTORY     FILE_ACCESS_MASK = 0x0004
	FILE_CREATE_PIPE_INSTANCE FILE_ACCESS_MASK = 0x0004
	FILE_ALL_ACCESS           FILE_ACCESS_MASK = 0x001f01ff
	FILE_DELETE_CHILD         FILE_ACCESS_MASK = 0x0040
	FILE_EXECUTE              FILE_ACCESS_MASK = 0x0020
	FILE_TRAVERSE             FILE_ACCESS_MASK = 0x0020
	FILE_READ_DATA            FILE_ACCESS_MASK = 0x0001
	FILE_LIST_DIRECTORY       FILE_ACCESS_MASK = 0x0001
	FILE_READ_ATTRIBUTES      FILE_ACCESS_MASK = 0x0080
	FILE_READ_EA              FILE_ACCESS_MASK = 0x0008
	FILE_WRITE_ATTRIBUTES     FILE_ACCESS_MASK = 0x0100
	FILE_WRITE_EA             FILE_ACCESS_MASK = 0x0010
	FILE_GENERIC_READ         FILE_ACCESS_MASK = 0x00120089
	FILE_GENERIC_WRITE        FILE_ACCESS_MASK = 0x00120116
	FILE_GENERIC_EXECUTE      FILE_ACCESS_MASK = 0x001200A0
)

// SECURITY_INFORMATION
type (
	SECURITY_INFORMATION  UINT
	PSECURITY_INFORMATION *SECURITY_INFORMATION
)

const (
	OWNER_SECURITY_INFORMATION            SECURITY_INFORMATION = 0x00000001
	GROUP_SECURITY_INFORMATION            SECURITY_INFORMATION = 0x00000002
	DACL_SECURITY_INFORMATION             SECURITY_INFORMATION = 0x00000004
	SACL_SECURITY_INFORMATION             SECURITY_INFORMATION = 0x00000008
	LABEL_SECURITY_INFORMATION            SECURITY_INFORMATION = 0x00000010
	PROTECTED_DACL_SECURITY_INFORMATION   SECURITY_INFORMATION = 0x80000000
	PROTECTED_SACL_SECURITY_INFORMATION   SECURITY_INFORMATION = 0x40000000
	UNPROTECTED_DACL_SECURITY_INFORMATION SECURITY_INFORMATION = 0x20000000
	UNPROTECTED_SACL_SECURITY_INFORMATION SECURITY_INFORMATION = 0x10000000
)

// SYSTEMTIME
type SYSTEMTIME struct {
	WYear         WORD
	WMonth        WORD
	WDayOfWeek    WORD
	WDay          WORD
	WHour         WORD
	WMinute       WORD
	WSecond       WORD
	WMilliseconds WORD
}

type (
	PSYSTEMTIME  *SYSTEMTIME
	LPSYSTEMTIME PSYSTEMTIME
)

// JOBOBJECTINFOCLASS
type JOBOBJECTINFOCLASS UINT

const (
	JobObjectBasicAccountingInformation         JOBOBJECTINFOCLASS = 1
	JobObjectBasicLimitInformation              JOBOBJECTINFOCLASS = 2
	JobObjectBasicProcessIdList                 JOBOBJECTINFOCLASS = 3
	JobObjectBasicUIRestrictions                JOBOBJECTINFOCLASS = 4
	JobObjectSecurityLimitInformation           JOBOBJECTINFOCLASS = 5
	JobObjectEndOfJobTimeInformation            JOBOBJECTINFOCLASS = 6
	JobObjectAssociateCompletionPortInformation JOBOBJECTINFOCLASS = 7
	JobObjectBasicAndIoAccountingInformation    JOBOBJECTINFOCLASS = 8
	JobObjectExtendedLimitInformation           JOBOBJECTINFOCLASS = 9
	JobObjectJobSetInformation                  JOBOBJECTINFOCLASS = 10
	JobObjectGroupInformation                   JOBOBJECTINFOCLASS = 11
	JobObjectNotificationLimitInformation       JOBOBJECTINFOCLASS = 12
	JobObjectLimitViolationInformation          JOBOBJECTINFOCLASS = 13
	JobObjectGroupInformationEx                 JOBOBJECTINFOCLASS = 14
	JobObjectCpuRateControlInformation          JOBOBJECTINFOCLASS = 15
	JobObjectCompletionFilter                   JOBOBJECTINFOCLASS = 16
	JobObjectCompletionCounter                  JOBOBJECTINFOCLASS = 17
	JobObjectFreezeInformation                  JOBOBJECTINFOCLASS = 18
	JobObjectExtendedAccountingInformation      JOBOBJECTINFOCLASS = 19
	JobObjectWakeInformation                    JOBOBJECTINFOCLASS = 20
	JobObjectIdleAwareInformation               JOBOBJECTINFOCLASS = 21
	JobObjectSchedulingRankBiasInformation      JOBOBJECTINFOCLASS = 22
	JobObjectTimerVirtualizationInformation     JOBOBJECTINFOCLASS = 23
	JobObjectCycleTimeNotification              JOBOBJECTINFOCLASS = 24
)

// PROCESSINFOCLASS
type PROCESSINFOCLASS UINT

const (
	ProcessBasicInformation                PROCESSINFOCLASS = 0
	ProcessQuotaLimits                     PROCESSINFOCLASS = 1
	ProcessIoCounters                      PROCESSINFOCLASS = 2
	ProcessVmCounters                      PROCESSINFOCLASS = 3
	ProcessTimes                           PROCESSINFOCLASS = 4
	ProcessBasePriority                    PROCESSINFOCLASS = 5
	ProcessRaisePriority                   PROCESSINFOCLASS = 6
	ProcessDebugPort                       PROCESSINFOCLASS = 7
	ProcessExceptionPort                   PROCESSINFOCLASS = 8
	ProcessAccessToken                     PROCESSINFOCLASS = 9
	ProcessLdtInformation                  PROCESSINFOCLASS = 10
	ProcessLdtSize                         PROCESSINFOCLASS = 11
	ProcessDefaultHardErrorMode            PROCESSINFOCLASS = 12
	ProcessIoPortHandlers                  PROCESSINFOCLASS = 13
	ProcessPooledUsageAndLimits            PROCESSINFOCLASS = 14
	ProcessWorkingSetWatch                 PROCESSINFOCLASS = 15
	ProcessUserModeIOPL                    PROCESSINFOCLASS = 16
	ProcessEnableAlignmentFaultFixup       PROCESSINFOCLASS = 17
	ProcessPriorityClass                   PROCESSINFOCLASS = 18
	ProcessWx86Information                 PROCESSINFOCLASS = 19
	ProcessHandleCount                     PROCESSINFOCLASS = 20
	ProcessAffinityMask                    PROCESSINFOCLASS = 21
	ProcessPriorityBoost                   PROCESSINFOCLASS = 22
	ProcessDeviceMap                       PROCESSINFOCLASS = 23
	ProcessSessionInformation              PROCESSINFOCLASS = 24
	ProcessForegroundInformation           PROCESSINFOCLASS = 25
	ProcessWow64Information                PROCESSINFOCLASS = 26
	ProcessImageFileName                   PROCESSINFOCLASS = 27
	ProcessLUIDDeviceMapsEnabled           PROCESSINFOCLASS = 28
	ProcessBreakOnTermination              PROCESSINFOCLASS = 29
	ProcessDebugObjectHandle               PROCESSINFOCLASS = 30
	ProcessDebugFlags                      PROCESSINFOCLASS = 31
	ProcessHandleTracing                   PROCESSINFOCLASS = 32
	ProcessIoPriority                      PROCESSINFOCLASS = 33
	ProcessExecuteFlags                    PROCESSINFOCLASS = 34
	ProcessTlsInformation                  PROCESSINFOCLASS = 35
	ProcessCookie                          PROCESSINFOCLASS = 36
	ProcessImageInformation                PROCESSINFOCLASS = 37
	ProcessCycleTime                       PROCESSINFOCLASS = 38
	ProcessPagePriority                    PROCESSINFOCLASS = 39
	ProcessInstrumentationCallback         PROCESSINFOCLASS = 40
	ProcessThreadStackAllocation           PROCESSINFOCLASS = 41
	ProcessWorkingSetWatchEx               PROCESSINFOCLASS = 42
	ProcessImageFileNameWin32              PROCESSINFOCLASS = 43
	ProcessImageFileMapping                PROCESSINFOCLASS = 44
	ProcessAffinityUpdateMode              PROCESSINFOCLASS = 45
	ProcessMemoryAllocationMode            PROCESSINFOCLASS = 46
	ProcessGroupInformation                PROCESSINFOCLASS = 47
	ProcessTokenVirtualizationEnabled      PROCESSINFOCLASS = 48
	ProcessConsoleHostProcess              PROCESSINFOCLASS = 49
	ProcessWindowInformation               PROCESSINFOCLASS = 50
	ProcessHandleInformation               PROCESSINFOCLASS = 51
	ProcessMitigationPolicy                PROCESSINFOCLASS = 52
	ProcessDynamicFunctionTableInformation PROCESSINFOCLASS = 53
	ProcessHandleCheckingMode              PROCESSINFOCLASS = 54
	ProcessKeepAliveCount                  PROCESSINFOCLASS = 55
	ProcessRevokeFileHandles               PROCESSINFOCLASS = 56
)

// THREADINFOCLASS
type THREADINFOCLASS UINT

const (
	ThreadBasicInformation          THREADINFOCLASS = 0
	ThreadTimes                     THREADINFOCLASS = 1
	ThreadPriority                  THREADINFOCLASS = 2
	ThreadBasePriority              THREADINFOCLASS = 3
	ThreadAffinityMask              THREADINFOCLASS = 4
	ThreadImpersonationToken        THREADINFOCLASS = 5
	ThreadDescriptorTableEntry      THREADINFOCLASS = 6
	ThreadEnableAlignmentFaultFixup THREADINFOCLASS = 7
	ThreadEventPair_Reusable        THREADINFOCLASS = 8
	ThreadQuerySetWin32StartAddress THREADINFOCLASS = 9
	ThreadZeroTlsCell               THREADINFOCLASS = 10
	ThreadPerformanceCount          THREADINFOCLASS = 11
	ThreadAmILastThread             THREADINFOCLASS = 12
	ThreadIdealProcessor            THREADINFOCLASS = 13
	ThreadPriorityBoost             THREADINFOCLASS = 14
	ThreadSetTlsArrayAddress        THREADINFOCLASS = 15
	ThreadIsIoPending               THREADINFOCLASS = 16
	ThreadHideFromDebugger          THREADINFOCLASS = 17
	ThreadBreakOnTermination        THREADINFOCLASS = 18
	ThreadSwitchLegacyState         THREADINFOCLASS = 19
	ThreadIsTerminated              THREADINFOCLASS = 20
	ThreadLastSystemCall            THREADINFOCLASS = 21
	ThreadIoPriority                THREADINFOCLASS = 22
	ThreadCycleTime                 THREADINFOCLASS = 23
	ThreadPagePriority              THREADINFOCLASS = 24
	ThreadActualBasePriority        THREADINFOCLASS = 25
	ThreadTebInformation            THREADINFOCLASS = 26
	ThreadCSwitchMon                THREADINFOCLASS = 27
	ThreadCSwitchPmu                THREADINFOCLASS = 28
	ThreadWow64Context              THREADINFOCLASS = 29
	ThreadGroupInformation          THREADINFOCLASS = 30
	ThreadUmsInformation            THREADINFOCLASS = 31
	ThreadCounterProfiling          THREADINFOCLASS = 32
	ThreadIdealProcessorEx          THREADINFOCLASS = 33
	ThreadCpuAccountingInformation  THREADINFOCLASS = 34
	ThreadSwitchStackCheck          THREADINFOCLASS = 35
)

// GET_FILEEX_INFO_LEVELS
type GET_FILEEX_INFO_LEVELS UINT

const (
	GetFileExInfoStandard GET_FILEEX_INFO_LEVELS = 0
)

// UNICODE_STRING
type UNICODE_STRING struct {
	Length        USHORT
	MaximumLength USHORT
	Buffer        PWSTR
}

type (
	PUNICODE_STRING  *UNICODE_STRING
	PCUNICODE_STRING *UNICODE_STRING
)

// STRING
type STRING struct {
	Length        USHORT
	MaximumLength USHORT
	Buffer        PCHAR
}

type (
	PSTRING *STRING
	// ANSI_STRING
	ANSI_STRING   STRING
	PANSI_STRING  *ANSI_STRING
	PCANSI_STRING *ANSI_STRING
	// OEM_STRING
	OEM_STRING   STRING
	POEM_STRING  *OEM_STRING
	PCOEM_STRING *OEM_STRING
)

// WinMsg
type WinMsg UINT

const (
	WM_NULL                    WinMsg = 0x0000
	WM_CREATE                  WinMsg = 0x0001
	WM_DESTROY                 WinMsg = 0x0002
	WM_MOVE                    WinMsg = 0x0003
	WM_SIZEWAIT                WinMsg = 0x0004
	WM_SIZE                    WinMsg = 0x0005
	WM_ACTIVATE                WinMsg = 0x0006
	WM_SETFOCUS                WinMsg = 0x0007
	WM_KILLFOCUS               WinMsg = 0x0008
	WM_SETVISIBLE              WinMsg = 0x0009
	WM_ENABLE                  WinMsg = 0x000A
	WM_SETREDRAW               WinMsg = 0x000B
	WM_SETTEXT                 WinMsg = 0x000C
	WM_GETTEXT                 WinMsg = 0x000D
	WM_GETTEXTLENGTH           WinMsg = 0x000E
	WM_PAINT                   WinMsg = 0x000F
	WM_CLOSE                   WinMsg = 0x0010
	WM_QUERYENDSESSION         WinMsg = 0x0011
	WM_QUIT                    WinMsg = 0x0012
	WM_QUERYOPEN               WinMsg = 0x0013
	WM_ERASEBKGND              WinMsg = 0x0014
	WM_SYSCOLORCHANGE          WinMsg = 0x0015
	WM_ENDSESSION              WinMsg = 0x0016
	WM_SYSTEMERROR             WinMsg = 0x0017
	WM_SHOWWINDOW              WinMsg = 0x0018
	WM_CTLCOLOR                WinMsg = 0x0019
	WM_SETTINGCHANGE           WinMsg = 0x001A
	WM_DEVMODECHANGE           WinMsg = 0x001B
	WM_ACTIVATEAPP             WinMsg = 0x001C
	WM_FONTCHANGE              WinMsg = 0x001D
	WM_TIMECHANGE              WinMsg = 0x001E
	WM_CANCELMODE              WinMsg = 0x001F
	WM_SETCURSOR               WinMsg = 0x0020
	WM_MOUSEACTIVATE           WinMsg = 0x0021
	WM_CHILDACTIVATE           WinMsg = 0x0022
	WM_QUEUESYNC               WinMsg = 0x0023
	WM_GETMINMAXINFO           WinMsg = 0x0024
	WM_PAINTICON               WinMsg = 0x0026
	WM_ICONERASEBKGND          WinMsg = 0x0027
	WM_NEXTDLGCTL              WinMsg = 0x0028
	WM_ALTTABACTIVE            WinMsg = 0x0029
	WM_SPOOLERSTATUS           WinMsg = 0x002A
	WM_DRAWITEM                WinMsg = 0x002B
	WM_MEASUREITEM             WinMsg = 0x002C
	WM_DELETEITEM              WinMsg = 0x002D
	WM_VKEYTOITEM              WinMsg = 0x002E
	WM_CHARTOITEM              WinMsg = 0x002F
	WM_SETFONT                 WinMsg = 0x0030
	WM_GETFONT                 WinMsg = 0x0031
	WM_SETHOTKEY               WinMsg = 0x0032
	WM_GETHOTKEY               WinMsg = 0x0033
	WM_FILESYSCHANGE           WinMsg = 0x0034
	WM_ISACTIVEICON            WinMsg = 0x0035
	WM_QUERYPARKICON           WinMsg = 0x0036
	WM_QUERYDRAGICON           WinMsg = 0x0037
	WM_QUERYSAVESTATE          WinMsg = 0x0038
	WM_COMPAREITEM             WinMsg = 0x0039
	WM_TESTING                 WinMsg = 0x003A
	WM_GETOBJECT               WinMsg = 0x003D
	WM_ACTIVATESHELLWINDOW     WinMsg = 0x003E
	WM_COMPACTING              WinMsg = 0x0041
	WM_COMMNOTIFY              WinMsg = 0x0044
	WM_WINDOWPOSCHANGING       WinMsg = 0x0046
	WM_WINDOWPOSCHANGED        WinMsg = 0x0047
	WM_POWER                   WinMsg = 0x0048
	WM_COPYDATA                WinMsg = 0x004A
	WM_CANCELJOURNAL           WinMsg = 0x004B
	WM_NOTIFY                  WinMsg = 0x004E
	WM_INPUTLANGCHANGEREQUEST  WinMsg = 0x0050
	WM_INPUTLANGCHANGE         WinMsg = 0x0051
	WM_TCARD                   WinMsg = 0x0052
	WM_HELP                    WinMsg = 0x0053
	WM_USERCHANGED             WinMsg = 0x0054
	WM_NOTIFYFORMAT            WinMsg = 0x0055
	WM_CONTEXTMENU             WinMsg = 0x007B
	WM_STYLECHANGING           WinMsg = 0x007C
	WM_STYLECHANGED            WinMsg = 0x007D
	WM_DISPLAYCHANGE           WinMsg = 0x007E
	WM_GETICON                 WinMsg = 0x007F
	WM_SETICON                 WinMsg = 0x0080
	WM_NCCREATE                WinMsg = 0x0081
	WM_NCDESTROY               WinMsg = 0x0082
	WM_NCCALCSIZE              WinMsg = 0x0083
	WM_NCHITTEST               WinMsg = 0x0084
	WM_NCPAINT                 WinMsg = 0x0085
	WM_NCACTIVATE              WinMsg = 0x0086
	WM_GETDLGCODE              WinMsg = 0x0087
	WM_SYNCPAINT               WinMsg = 0x0088
	WM_SYNCTASK                WinMsg = 0x0089
	WM_NCMOUSEMOVE             WinMsg = 0x00A0
	WM_NCLBUTTONDOWN           WinMsg = 0x00A1
	WM_NCLBUTTONUP             WinMsg = 0x00A2
	WM_NCLBUTTONDBLCLK         WinMsg = 0x00A3
	WM_NCRBUTTONDOWN           WinMsg = 0x00A4
	WM_NCRBUTTONUP             WinMsg = 0x00A5
	WM_NCRBUTTONDBLCLK         WinMsg = 0x00A6
	WM_NCMBUTTONDOWN           WinMsg = 0x00A7
	WM_NCMBUTTONUP             WinMsg = 0x00A8
	WM_NCMBUTTONDBLCLK         WinMsg = 0x00A9
	WM_NCXBUTTONDOWN           WinMsg = 0x00AB
	WM_NCXBUTTONUP             WinMsg = 0x00AC
	WM_NCXBUTTONDBLCLK         WinMsg = 0x00AD
	WM_INPUT_DEVICE_CHANGE     WinMsg = 0x00FE
	WM_INPUT                   WinMsg = 0x00FF
	WM_KEYDOWN                 WinMsg = 0x0100
	WM_KEYUP                   WinMsg = 0x0101
	WM_CHAR                    WinMsg = 0x0102
	WM_DEADCHAR                WinMsg = 0x0103
	WM_SYSKEYDOWN              WinMsg = 0x0104
	WM_SYSKEYUP                WinMsg = 0x0105
	WM_SYSCHAR                 WinMsg = 0x0106
	WM_SYSDEADCHAR             WinMsg = 0x0107
	WM_UNICHAR                 WinMsg = 0x0109
	WM_IME_STARTCOMPOSITION    WinMsg = 0x010D
	WM_IME_ENDCOMPOSITION      WinMsg = 0x010E
	WM_IME_COMPOSITION         WinMsg = 0x010F
	WM_INITDIALOG              WinMsg = 0x0110
	WM_COMMAND                 WinMsg = 0x0111
	WM_SYSCOMMAND              WinMsg = 0x0112
	WM_TIMER                   WinMsg = 0x0113
	WM_HSCROLL                 WinMsg = 0x0114
	WM_VSCROLL                 WinMsg = 0x0115
	WM_INITMENU                WinMsg = 0x0116
	WM_INITMENUPOPUP           WinMsg = 0x0117
	WM_SYSTIMER                WinMsg = 0x0118
	WM_GESTURE                 WinMsg = 0x0119
	WM_GESTURENOTIFY           WinMsg = 0x011A
	WM_MENUSELECT              WinMsg = 0x011F
	WM_MENUCHAR                WinMsg = 0x0120
	WM_ENTERIDLE               WinMsg = 0x0121
	WM_MENURBUTTONUP           WinMsg = 0x0122
	WM_MENUDRAG                WinMsg = 0x0123
	WM_MENUGETOBJECT           WinMsg = 0x0124
	WM_UNINITMENUPOPUP         WinMsg = 0x0125
	WM_MENUCOMMAND             WinMsg = 0x0126
	WM_CHANGEUISTATE           WinMsg = 0x0127
	WM_UPDATEUISTATE           WinMsg = 0x0128
	WM_QUERYUISTATE            WinMsg = 0x0129
	WM_LBTRACKPOINT            WinMsg = 0x0131
	WM_CTLCOLORMSGBOX          WinMsg = 0x0132
	WM_CTLCOLOREDIT            WinMsg = 0x0133
	WM_CTLCOLORLISTBOX         WinMsg = 0x0134
	WM_CTLCOLORBTN             WinMsg = 0x0135
	WM_CTLCOLORDLG             WinMsg = 0x0136
	WM_CTLCOLORSCROLLBAR       WinMsg = 0x0137
	WM_CTLCOLORSTATIC          WinMsg = 0x0138
	MN_GETHMENU                WinMsg = 0x01E1
	WM_MOUSEMOVE               WinMsg = 0x0200
	WM_LBUTTONDOWN             WinMsg = 0x0201
	WM_LBUTTONUP               WinMsg = 0x0202
	WM_LBUTTONDBLCLK           WinMsg = 0x0203
	WM_RBUTTONDOWN             WinMsg = 0x0204
	WM_RBUTTONUP               WinMsg = 0x0205
	WM_RBUTTONDBLCLK           WinMsg = 0x0206
	WM_MBUTTONDOWN             WinMsg = 0x0207
	WM_MBUTTONUP               WinMsg = 0x0208
	WM_MBUTTONDBLCLK           WinMsg = 0x0209
	WM_MOUSEWHEEL              WinMsg = 0x020A
	WM_XBUTTONDOWN             WinMsg = 0x020B
	WM_XBUTTONUP               WinMsg = 0x020C
	WM_XBUTTONDBLCLK           WinMsg = 0x020D
	WM_MOUSEHWHEEL             WinMsg = 0x020E
	WM_PARENTNOTIFY            WinMsg = 0x0210
	WM_ENTERMENULOOP           WinMsg = 0x0211
	WM_EXITMENULOOP            WinMsg = 0x0212
	WM_NEXTMENU                WinMsg = 0x0213
	WM_SIZING                  WinMsg = 0x0214
	WM_CAPTURECHANGED          WinMsg = 0x0215
	WM_MOVING                  WinMsg = 0x0216
	WM_POWERBROADCAST          WinMsg = 0x0218
	WM_DEVICECHANGE            WinMsg = 0x0219
	WM_MDICREATE               WinMsg = 0x0220
	WM_MDIDESTROY              WinMsg = 0x0221
	WM_MDIACTIVATE             WinMsg = 0x0222
	WM_MDIRESTORE              WinMsg = 0x0223
	WM_MDINEXT                 WinMsg = 0x0224
	WM_MDIMAXIMIZE             WinMsg = 0x0225
	WM_MDITILE                 WinMsg = 0x0226
	WM_MDICASCADE              WinMsg = 0x0227
	WM_MDIICONARRANGE          WinMsg = 0x0228
	WM_MDIGETACTIVE            WinMsg = 0x0229
	WM_DROPOBJECT              WinMsg = 0x022A
	WM_QUERYDROPOBJECT         WinMsg = 0x022B
	WM_BEGINDRAG               WinMsg = 0x022C
	WM_DRAGLOOP                WinMsg = 0x022D
	WM_DRAGSELECT              WinMsg = 0x022E
	WM_DRAGMOVE                WinMsg = 0x022F
	WM_MDISETMENU              WinMsg = 0x0230
	WM_ENTERSIZEMOVE           WinMsg = 0x0231
	WM_EXITSIZEMOVE            WinMsg = 0x0232
	WM_DROPFILES               WinMsg = 0x0233
	WM_MDIREFRESHMENU          WinMsg = 0x0234
	WM_POINTERDEVICECHANGE     WinMsg = 0x0238
	WM_POINTERDEVICEINRANGE    WinMsg = 0x0239
	WM_POINTERDEVICEOUTOFRANGE WinMsg = 0x023A
	WM_TOUCH                   WinMsg = 0x0240
	WM_NCPOINTERUPDATE         WinMsg = 0x0241
	WM_NCPOINTERDOWN           WinMsg = 0x0242
	WM_NCPOINTERUP             WinMsg = 0x0243
	WM_POINTERUPDATE           WinMsg = 0x0245
	WM_POINTERDOWN             WinMsg = 0x0246
	WM_POINTERUP               WinMsg = 0x0247
	WM_POINTERENTER            WinMsg = 0x0249
	WM_POINTERLEAVE            WinMsg = 0x024A
	WM_POINTERACTIVATE         WinMsg = 0x024B
	WM_POINTERCAPTURECHANGED   WinMsg = 0x024C
	WM_TOUCHHITTESTING         WinMsg = 0x024D
	WM_POINTERWHEEL            WinMsg = 0x024E
	WM_POINTERHWHEEL           WinMsg = 0x024F
	// IME
	WM_IME_REPORT          WinMsg = 0x0280
	WM_IME_SETCONTEXT      WinMsg = 0x0281
	WM_IME_NOTIFY          WinMsg = 0x0282
	WM_IME_CONTROL         WinMsg = 0x0283
	WM_IME_COMPOSITIONFULL WinMsg = 0x0284
	WM_IME_SELECT          WinMsg = 0x0285
	WM_IME_CHAR            WinMsg = 0x0286
	WM_IME_REQUEST         WinMsg = 0x0288
	WM_IME_KEYDOWN         WinMsg = 0x0290
	WM_IME_KEYUP           WinMsg = 0x0291

	WM_MOUSEHOVER                      WinMsg = 0x02A1
	WM_MOUSELEAVE                      WinMsg = 0x02A3
	WM_NCMOUSEHOVER                    WinMsg = 0x02A0
	WM_NCMOUSELEAVE                    WinMsg = 0x02A2
	WM_WTSSESSION_CHANGE               WinMsg = 0x02B1
	WM_TABLET_ADDED                    WinMsg = 0x02c8
	WM_TABLET_DELETED                  WinMsg = 0x02c9
	WM_TABLET_FLICK                    WinMsg = 0x02cb
	WM_TABLET_QUERYSYSTEMGESTURESTATUS WinMsg = 0x02cc
	WM_CUT                             WinMsg = 0x0300
	WM_COPY                            WinMsg = 0x0301
	WM_PASTE                           WinMsg = 0x0302
	WM_CLEAR                           WinMsg = 0x0303
	WM_UNDO                            WinMsg = 0x0304
	WM_RENDERFORMAT                    WinMsg = 0x0305
	WM_RENDERALLFORMATS                WinMsg = 0x0306
	WM_DESTROYCLIPBOARD                WinMsg = 0x0307
	WM_DRAWCLIPBOARD                   WinMsg = 0x0308
	WM_PAINTCLIPBOARD                  WinMsg = 0x0309
	WM_VSCROLLCLIPBOARD                WinMsg = 0x030A
	WM_SIZECLIPBOARD                   WinMsg = 0x030B
	WM_ASKCBFORMATNAME                 WinMsg = 0x030C
	WM_CHANGECBCHAIN                   WinMsg = 0x030D
	WM_HSCROLLCLIPBOARD                WinMsg = 0x030E
	WM_QUERYNEWPALETTE                 WinMsg = 0x030F
	WM_PALETTEISCHANGING               WinMsg = 0x0310
	WM_PALETTECHANGED                  WinMsg = 0x0311
	WM_HOTKEY                          WinMsg = 0x0312
	WM_PRINT                           WinMsg = 0x0317
	WM_PRINTCLIENT                     WinMsg = 0x0318
	WM_APPCOMMAND                      WinMsg = 0x0319
	WM_THEMECHANGED                    WinMsg = 0x031A
	WM_CLIPBOARDUPDATE                 WinMsg = 0x031D
	WM_DWMCOMPOSITIONCHANGED           WinMsg = 0x031E
	WM_DWMNCRENDERINGCHANGED           WinMsg = 0x031F
	WM_DWMCOLORIZATIONCOLORCHANGED     WinMsg = 0x0320
	WM_DWMWINDOWMAXIMIZEDCHANGE        WinMsg = 0x0321
	WM_GETTITLEBARINFOEX               WinMsg = 0x033F
	WM_QUERYAFXWNDPROC                 WinMsg = 0x0360
	WM_SIZEPARENT                      WinMsg = 0x0361
	WM_SETMESSAGESTRING                WinMsg = 0x0362
	WM_IDLEUPDATECMDUI                 WinMsg = 0x0363
	WM_INITIALUPDATE                   WinMsg = 0x0364
	WM_COMMANDHELP                     WinMsg = 0x0365
	WM_HELPHITTEST                     WinMsg = 0x0366
	WM_EXITHELPMODE                    WinMsg = 0x0367
	WM_RECALCPARENT                    WinMsg = 0x0368
	WM_SIZECHILD                       WinMsg = 0x0369
	WM_KICKIDLE                        WinMsg = 0x036A
	WM_QUERYCENTERWND                  WinMsg = 0x036B
	WM_DISABLEMODAL                    WinMsg = 0x036C
	WM_FLOATSTATUS                     WinMsg = 0x036D
	WM_ACTIVATETOPLEVEL                WinMsg = 0x036E
	WM_QUERY3DCONTROLS                 WinMsg = 0x036F
	WM_SOCKET_NOTIFY                   WinMsg = 0x0373
	WM_SOCKET_DEAD                     WinMsg = 0x0374
	WM_POPMESSAGESTRING                WinMsg = 0x0375
	WM_HELPPROMPTADDR                  WinMsg = 0x0376
	WM_QUEUE_SENTINEL                  WinMsg = 0x0379
	WM_RESERVED_037C                   WinMsg = 0x037C
	WM_RESERVED_037D                   WinMsg = 0x037D
	WM_RESERVED_037E                   WinMsg = 0x037E
	WM_FORWARDMSG                      WinMsg = 0x037F
	WM_USER                            WinMsg = 0x0400
	// Edit
	EM_GETSEL              WinMsg = 0x00B0
	EM_SETSEL              WinMsg = 0x00B1
	EM_GETRECT             WinMsg = 0x00B2
	EM_SETRECT             WinMsg = 0x00B3
	EM_SETRECTNP           WinMsg = 0x00B4
	EM_SCROLL              WinMsg = 0x00B5
	EM_LINESCROLL          WinMsg = 0x00B6
	EM_SCROLLCARET         WinMsg = 0x00B7
	EM_GETMODIFY           WinMsg = 0x00B8
	EM_SETMODIFY           WinMsg = 0x00B9
	EM_GETLINECOUNT        WinMsg = 0x00BA
	EM_LINEINDEX           WinMsg = 0x00BB
	EM_SETHANDLE           WinMsg = 0x00BC
	EM_GETHANDLE           WinMsg = 0x00BD
	EM_GETTHUMB            WinMsg = 0x00BE
	EM_LINELENGTH          WinMsg = 0x00C1
	EM_REPLACESEL          WinMsg = 0x00C2
	EM_GETLINE             WinMsg = 0x00C4
	EM_SETLIMITTEXT        WinMsg = 0x00C5
	EM_CANUNDO             WinMsg = 0x00C6
	EM_UNDO                WinMsg = 0x00C7
	EM_FMTLINES            WinMsg = 0x00C8
	EM_LINEFROMCHAR        WinMsg = 0x00C9
	EM_SETTABSTOPS         WinMsg = 0x00CB
	EM_SETPASSWORDCHAR     WinMsg = 0x00CC
	EM_EMPTYUNDOBUFFER     WinMsg = 0x00CD
	EM_GETFIRSTVISIBLELINE WinMsg = 0x00CE
	EM_SETREADONLY         WinMsg = 0x00CF
	EM_SETWORDBREAKPROC    WinMsg = 0x00D0
	EM_GETWORDBREAKPROC    WinMsg = 0x00D1
	EM_GETPASSWORDCHAR     WinMsg = 0x00D2
	EM_SETMARGINS          WinMsg = 0x00D3
	EM_GETMARGINS          WinMsg = 0x00D4
	EM_GETLIMITTEXT        WinMsg = 0x00D5
	EM_POSFROMCHAR         WinMsg = 0x00D6
	EM_CHARFROMPOS         WinMsg = 0x00D7
	EM_SETIMESTATUS        WinMsg = 0x00D8
	EM_GETIMESTATUS        WinMsg = 0x00D9
	// Button
	BM_GETCHECK     WinMsg = 0x00F0
	BM_SETCHECK     WinMsg = 0x00F1
	BM_GETSTATE     WinMsg = 0x00F2
	BM_SETSTATE     WinMsg = 0x00F3
	BM_SETSTYLE     WinMsg = 0x00F4
	BM_CLICK        WinMsg = 0x00F5
	BM_GETIMAGE     WinMsg = 0x00F6
	BM_SETIMAGE     WinMsg = 0x00F7
	BM_SETDONTCLICK WinMsg = 0x00F8
	// Combo Box
	CB_GETEDITSEL            WinMsg = 0x0140
	CB_LIMITTEXT             WinMsg = 0x0141
	CB_SETEDITSEL            WinMsg = 0x0142
	CB_ADDSTRING             WinMsg = 0x0143
	CB_DELETESTRING          WinMsg = 0x0144
	CB_DIR                   WinMsg = 0x0145
	CB_GETCOUNT              WinMsg = 0x0146
	CB_GETCURSEL             WinMsg = 0x0147
	CB_GETLBTEXT             WinMsg = 0x0148
	CB_GETLBTEXTLEN          WinMsg = 0x0149
	CB_INSERTSTRING          WinMsg = 0x014A
	CB_RESETCONTENT          WinMsg = 0x014B
	CB_FINDSTRING            WinMsg = 0x014C
	CB_SELECTSTRING          WinMsg = 0x014D
	CB_SETCURSEL             WinMsg = 0x014E
	CB_SHOWDROPDOWN          WinMsg = 0x014F
	CB_GETITEMDATA           WinMsg = 0x0150
	CB_SETITEMDATA           WinMsg = 0x0151
	CB_GETDROPPEDCONTROLRECT WinMsg = 0x0152
	CB_SETITEMHEIGHT         WinMsg = 0x0153
	CB_GETITEMHEIGHT         WinMsg = 0x0154
	CB_SETEXTENDEDUI         WinMsg = 0x0155
	CB_GETEXTENDEDUI         WinMsg = 0x0156
	CB_GETDROPPEDSTATE       WinMsg = 0x0157
	CB_FINDSTRINGEXACT       WinMsg = 0x0158
	CB_SETLOCALE             WinMsg = 0x0159
	CB_GETLOCALE             WinMsg = 0x015A
	CB_GETTOPINDEX           WinMsg = 0x015b
	CB_SETTOPINDEX           WinMsg = 0x015c
	CB_GETHORIZONTALEXTENT   WinMsg = 0x015d
	CB_SETHORIZONTALEXTENT   WinMsg = 0x015e
	CB_GETDROPPEDWIDTH       WinMsg = 0x015f
	CB_SETDROPPEDWIDTH       WinMsg = 0x0160
	CB_INITSTORAGE           WinMsg = 0x0161
	CB_MULTIPLEADDSTRING     WinMsg = 0x0163
	CB_GETCOMBOBOXINFO       WinMsg = 0x0164
	// List Box
	LB_ADDSTRING           WinMsg = 0x0180
	LB_INSERTSTRING        WinMsg = 0x0181
	LB_DELETESTRING        WinMsg = 0x0182
	LB_SELITEMRANGEEX      WinMsg = 0x0183
	LB_RESETCONTENT        WinMsg = 0x0184
	LB_SETSEL              WinMsg = 0x0185
	LB_SETCURSEL           WinMsg = 0x0186
	LB_GETSEL              WinMsg = 0x0187
	LB_GETCURSEL           WinMsg = 0x0188
	LB_GETTEXT             WinMsg = 0x0189
	LB_GETTEXTLEN          WinMsg = 0x018A
	LB_GETCOUNT            WinMsg = 0x018B
	LB_SELECTSTRING        WinMsg = 0x018C
	LB_DIR                 WinMsg = 0x018D
	LB_GETTOPINDEX         WinMsg = 0x018E
	LB_FINDSTRING          WinMsg = 0x018F
	LB_GETSELCOUNT         WinMsg = 0x0190
	LB_GETSELITEMS         WinMsg = 0x0191
	LB_SETTABSTOPS         WinMsg = 0x0192
	LB_GETHORIZONTALEXTENT WinMsg = 0x0193
	LB_SETHORIZONTALEXTENT WinMsg = 0x0194
	LB_SETCOLUMNWIDTH      WinMsg = 0x0195
	LB_ADDFILE             WinMsg = 0x0196
	LB_SETTOPINDEX         WinMsg = 0x0197
	LB_GETITEMRECT         WinMsg = 0x0198
	LB_GETITEMDATA         WinMsg = 0x0199
	LB_SETITEMDATA         WinMsg = 0x019A
	LB_SELITEMRANGE        WinMsg = 0x019B
	LB_SETANCHORINDEX      WinMsg = 0x019C
	LB_GETANCHORINDEX      WinMsg = 0x019D
	LB_SETCARETINDEX       WinMsg = 0x019E
	LB_GETCARETINDEX       WinMsg = 0x019F
	LB_SETITEMHEIGHT       WinMsg = 0x01A0
	LB_GETITEMHEIGHT       WinMsg = 0x01A1
	LB_FINDSTRINGEXACT     WinMsg = 0x01A2
	LB_CARETON             WinMsg = 0x01A3
	LB_CARETOFF            WinMsg = 0x01A4
	LB_SETLOCALE           WinMsg = 0x01A5
	LB_GETLOCALE           WinMsg = 0x01A6
	LB_SETCOUNT            WinMsg = 0x01A7
	LB_INITSTORAGE         WinMsg = 0x01A8
	LB_ITEMFROMPOINT       WinMsg = 0x01A9
	LB_MULTIPLEADDSTRING   WinMsg = 0x01B1
	LB_GETLISTBOXINFO      WinMsg = 0x01B2
	// Scroll Bar
	SBM_SETPOS           WinMsg = 0x00E0
	SBM_GETPOS           WinMsg = 0x00E1
	SBM_SETRANGE         WinMsg = 0x00E2
	SBM_GETRANGE         WinMsg = 0x00E3
	SBM_ENABLE_ARROWS    WinMsg = 0x00E4
	SBM_SETRANGEREDRAW   WinMsg = 0x00E6
	SBM_SETSCROLLINFO    WinMsg = 0x00E9
	SBM_GETSCROLLINFO    WinMsg = 0x00EA
	SBM_GETSCROLLBARINFO WinMsg = 0x00EB
	// Static
	STM_SETICON  WinMsg = 0x0170
	STM_GETICON  WinMsg = 0x0171
	STM_SETIMAGE WinMsg = 0x0172
	STM_GETIMAGE WinMsg = 0x0173
	// User
	WM_USER_1   WinMsg = 0x0401
	WM_USER_2   WinMsg = 0x0402
	WM_USER_3   WinMsg = 0x0403
	WM_USER_4   WinMsg = 0x0404
	WM_USER_5   WinMsg = 0x0405
	WM_USER_6   WinMsg = 0x0406
	WM_USER_7   WinMsg = 0x0407
	WM_USER_8   WinMsg = 0x0408
	WM_USER_9   WinMsg = 0x0409
	WM_USER_10  WinMsg = 0x040a
	WM_USER_11  WinMsg = 0x040b
	WM_USER_12  WinMsg = 0x040c
	WM_USER_13  WinMsg = 0x040d
	WM_USER_14  WinMsg = 0x040e
	WM_USER_15  WinMsg = 0x040f
	WM_USER_16  WinMsg = 0x0410
	WM_USER_17  WinMsg = 0x0411
	WM_USER_18  WinMsg = 0x0412
	WM_USER_19  WinMsg = 0x0413
	WM_USER_20  WinMsg = 0x0414
	WM_USER_21  WinMsg = 0x0415
	WM_USER_22  WinMsg = 0x0416
	WM_USER_23  WinMsg = 0x0417
	WM_USER_24  WinMsg = 0x0418
	WM_USER_25  WinMsg = 0x0419
	WM_USER_26  WinMsg = 0x041a
	WM_USER_27  WinMsg = 0x041b
	WM_USER_28  WinMsg = 0x041c
	WM_USER_29  WinMsg = 0x041d
	WM_USER_30  WinMsg = 0x041e
	WM_USER_31  WinMsg = 0x041f
	WM_USER_32  WinMsg = 0x0420
	WM_USER_33  WinMsg = 0x0421
	WM_USER_34  WinMsg = 0x0422
	WM_USER_35  WinMsg = 0x0423
	WM_USER_36  WinMsg = 0x0424
	WM_USER_37  WinMsg = 0x0425
	WM_USER_38  WinMsg = 0x0426
	WM_USER_39  WinMsg = 0x0427
	WM_USER_40  WinMsg = 0x0428
	WM_USER_41  WinMsg = 0x0429
	WM_USER_42  WinMsg = 0x042a
	WM_USER_43  WinMsg = 0x042b
	WM_USER_44  WinMsg = 0x042c
	WM_USER_45  WinMsg = 0x042d
	WM_USER_46  WinMsg = 0x042e
	WM_USER_47  WinMsg = 0x042f
	WM_USER_48  WinMsg = 0x0430
	WM_USER_49  WinMsg = 0x0431
	WM_USER_50  WinMsg = 0x0432
	WM_USER_51  WinMsg = 0x0433
	WM_USER_52  WinMsg = 0x0434
	WM_USER_53  WinMsg = 0x0435
	WM_USER_54  WinMsg = 0x0436
	WM_USER_55  WinMsg = 0x0437
	WM_USER_56  WinMsg = 0x0438
	WM_USER_57  WinMsg = 0x0439
	WM_USER_58  WinMsg = 0x043a
	WM_USER_59  WinMsg = 0x043b
	WM_USER_60  WinMsg = 0x043c
	WM_USER_61  WinMsg = 0x043d
	WM_USER_62  WinMsg = 0x043e
	WM_USER_63  WinMsg = 0x043f
	WM_USER_64  WinMsg = 0x0440
	WM_USER_65  WinMsg = 0x0441
	WM_USER_66  WinMsg = 0x0442
	WM_USER_67  WinMsg = 0x0443
	WM_USER_68  WinMsg = 0x0444
	WM_USER_69  WinMsg = 0x0445
	WM_USER_70  WinMsg = 0x0446
	WM_USER_71  WinMsg = 0x0447
	WM_USER_72  WinMsg = 0x0448
	WM_USER_73  WinMsg = 0x0449
	WM_USER_74  WinMsg = 0x044a
	WM_USER_75  WinMsg = 0x044b
	WM_USER_76  WinMsg = 0x044c
	WM_USER_77  WinMsg = 0x044d
	WM_USER_78  WinMsg = 0x044e
	WM_USER_79  WinMsg = 0x044f
	WM_USER_80  WinMsg = 0x0450
	WM_USER_81  WinMsg = 0x0451
	WM_USER_82  WinMsg = 0x0452
	WM_USER_83  WinMsg = 0x0453
	WM_USER_84  WinMsg = 0x0454
	WM_USER_85  WinMsg = 0x0455
	WM_USER_86  WinMsg = 0x0456
	WM_USER_87  WinMsg = 0x0457
	WM_USER_88  WinMsg = 0x0458
	WM_USER_89  WinMsg = 0x0459
	WM_USER_90  WinMsg = 0x045a
	WM_USER_91  WinMsg = 0x045b
	WM_USER_92  WinMsg = 0x045c
	WM_USER_93  WinMsg = 0x045d
	WM_USER_94  WinMsg = 0x045e
	WM_USER_95  WinMsg = 0x045f
	WM_USER_96  WinMsg = 0x0460
	WM_USER_97  WinMsg = 0x0461
	WM_USER_98  WinMsg = 0x0462
	WM_USER_99  WinMsg = 0x0463
	WM_USER_100 WinMsg = 0x0464
	// List View
	LVM_GETBKCOLOR               WinMsg = 0x00001000
	LVM_SETBKCOLOR               WinMsg = 0x00001001
	LVM_GETIMAGELIST             WinMsg = 0x00001002
	LVM_SETIMAGELIST             WinMsg = 0x00001003
	LVM_GETITEMCOUNT             WinMsg = 0x00001004
	LVM_GETITEMA                 WinMsg = 0x00001005
	LVM_GETITEMW                 WinMsg = 0x0000104b
	LVM_SETITEMA                 WinMsg = 0x00001006
	LVM_SETITEMW                 WinMsg = 0x0000104c
	LVM_INSERTITEMA              WinMsg = 0x00001007
	LVM_INSERTITEMW              WinMsg = 0x0000104d
	LVM_DELETEITEM               WinMsg = 0x00001008
	LVM_DELETEALLITEMS           WinMsg = 0x00001009
	LVM_GETCALLBACKMASK          WinMsg = 0x0000100a
	LVM_SETCALLBACKMASK          WinMsg = 0x0000100b
	LVM_GETNEXTITEM              WinMsg = 0x0000100c
	LVM_FINDITEMA                WinMsg = 0x0000100d
	LVM_FINDITEMW                WinMsg = 0x00001053
	LVM_GETITEMRECT              WinMsg = 0x0000100e
	LVM_SETITEMPOSITION          WinMsg = 0x0000100f
	LVM_GETITEMPOSITION          WinMsg = 0x00001010
	LVM_GETSTRINGWIDTHA          WinMsg = 0x00001011
	LVM_GETSTRINGWIDTHW          WinMsg = 0x00001057
	LVM_HITTEST                  WinMsg = 0x00001012
	LVM_ENSUREVISIBLE            WinMsg = 0x00001013
	LVM_SCROLL                   WinMsg = 0x00001014
	LVM_REDRAWITEMS              WinMsg = 0x00001015
	LVM_ARRANGE                  WinMsg = 0x00001016
	LVM_EDITLABELA               WinMsg = 0x00001017
	LVM_EDITLABELW               WinMsg = 0x00001076
	LVM_GETEDITCONTROL           WinMsg = 0x00001018
	LVM_GETCOLUMNA               WinMsg = 0x00001019
	LVM_GETCOLUMNW               WinMsg = 0x0000105f
	LVM_SETCOLUMNA               WinMsg = 0x0000101a
	LVM_SETCOLUMNW               WinMsg = 0x00001060
	LVM_INSERTCOLUMNA            WinMsg = 0x0000101b
	LVM_INSERTCOLUMNW            WinMsg = 0x00001061
	LVM_DELETECOLUMN             WinMsg = 0x0000101c
	LVM_GETCOLUMNWIDTH           WinMsg = 0x0000101d
	LVM_SETCOLUMNWIDTH           WinMsg = 0x0000101e
	LVM_GETHEADER                WinMsg = 0x0000101f
	LVM_CREATEDRAGIMAGE          WinMsg = 0x00001021
	LVM_GETVIEWRECT              WinMsg = 0x00001022
	LVM_GETTEXTCOLOR             WinMsg = 0x00001023
	LVM_SETTEXTCOLOR             WinMsg = 0x00001024
	LVM_GETTEXTBKCOLOR           WinMsg = 0x00001025
	LVM_SETTEXTBKCOLOR           WinMsg = 0x00001026
	LVM_GETTOPINDEX              WinMsg = 0x00001027
	LVM_GETCOUNTPERPAGE          WinMsg = 0x00001028
	LVM_GETORIGIN                WinMsg = 0x00001029
	LVM_UPDATE                   WinMsg = 0x0000102a
	LVM_SETITEMSTATE             WinMsg = 0x0000102b
	LVM_GETITEMSTATE             WinMsg = 0x0000102c
	LVM_GETITEMTEXTA             WinMsg = 0x0000102d
	LVM_GETITEMTEXTW             WinMsg = 0x00001073
	LVM_SETITEMTEXTA             WinMsg = 0x0000102e
	LVM_SETITEMTEXTW             WinMsg = 0x00001074
	LVM_SETITEMCOUNT             WinMsg = 0x0000102f
	LVM_SORTITEMS                WinMsg = 0x00001030
	LVM_SETITEMPOSITION32        WinMsg = 0x00001031
	LVM_GETSELECTEDCOUNT         WinMsg = 0x00001032
	LVM_GETITEMSPACING           WinMsg = 0x00001033
	LVM_GETISEARCHSTRINGA        WinMsg = 0x00001034
	LVM_GETISEARCHSTRINGW        WinMsg = 0x00001075
	LVM_SETICONSPACING           WinMsg = 0x00001035
	LVM_SETEXTENDEDLISTVIEWSTYLE WinMsg = 0x00001036
	LVM_GETEXTENDEDLISTVIEWSTYLE WinMsg = 0x00001037
	LVM_GETSUBITEMRECT           WinMsg = 0x00001038
	LVM_SUBITEMHITTEST           WinMsg = 0x00001039
	LVM_SETCOLUMNORDERARRAY      WinMsg = 0x0000103a
	LVM_GETCOLUMNORDERARRAY      WinMsg = 0x0000103b
	LVM_SETHOTITEM               WinMsg = 0x0000103c
	LVM_GETHOTITEM               WinMsg = 0x0000103d
	LVM_SETHOTCURSOR             WinMsg = 0x0000103e
	LVM_GETHOTCURSOR             WinMsg = 0x0000103f
	LVM_APPROXIMATEVIEWRECT      WinMsg = 0x00001040
	LVM_SETWORKAREAS             WinMsg = 0x00001041
	LVM_GETWORKAREAS             WinMsg = 0x00001046
	LVM_GETNUMBEROFWORKAREAS     WinMsg = 0x00001049
	LVM_GETSELECTIONMARK         WinMsg = 0x00001042
	LVM_SETSELECTIONMARK         WinMsg = 0x00001043
	LVM_SETHOVERTIME             WinMsg = 0x00001047
	LVM_GETHOVERTIME             WinMsg = 0x00001048
	LVM_SETTOOLTIPS              WinMsg = 0x0000104a
	LVM_GETTOOLTIPS              WinMsg = 0x0000104e
	LVM_SORTITEMSEX              WinMsg = 0x00001051
	LVM_SETBKIMAGEA              WinMsg = 0x00001044
	LVM_SETBKIMAGEW              WinMsg = 0x0000108a
	LVM_GETBKIMAGEA              WinMsg = 0x00001045
	LVM_GETBKIMAGEW              WinMsg = 0x0000108b
	LVM_SETSELECTEDCOLUMN        WinMsg = 0x0000108c
	LVM_SETVIEW                  WinMsg = 0x0000108e
	LVM_GETVIEW                  WinMsg = 0x0000108f
	LVM_INSERTGROUP              WinMsg = 0x00001091
	LVM_SETGROUPINFO             WinMsg = 0x00001093
	LVM_GETGROUPINFO             WinMsg = 0x00001095
	LVM_REMOVEGROUP              WinMsg = 0x00001096
	LVM_MOVEGROUP                WinMsg = 0x00001097
	LVM_GETGROUPCOUNT            WinMsg = 0x00001098
	LVM_GETGROUPINFOBYINDEX      WinMsg = 0x00001099
	LVM_MOVEITEMTOGROUP          WinMsg = 0x0000109a
	LVM_GETGROUPRECT             WinMsg = 0x00001062
	LVM_SETGROUPMETRICS          WinMsg = 0x0000109b
	LVM_GETGROUPMETRICS          WinMsg = 0x0000109c
	LVM_ENABLEGROUPVIEW          WinMsg = 0x0000109d
	LVM_SORTGROUPS               WinMsg = 0x0000109e
	LVM_INSERTGROUPSORTED        WinMsg = 0x0000109f
	LVM_REMOVEALLGROUPS          WinMsg = 0x000010a0
	LVM_HASGROUP                 WinMsg = 0x000010a1
	LVM_GETGROUPSTATE            WinMsg = 0x0000105c
	LVM_GETFOCUSEDGROUP          WinMsg = 0x0000105d
	LVM_SETTILEVIEWINFO          WinMsg = 0x000010a2
	LVM_GETTILEVIEWINFO          WinMsg = 0x000010a3
	LVM_SETTILEINFO              WinMsg = 0x000010a4
	LVM_GETTILEINFO              WinMsg = 0x000010a5
	LVM_SETINSERTMARK            WinMsg = 0x000010a6
	LVM_GETINSERTMARK            WinMsg = 0x000010a7
	LVM_INSERTMARKHITTEST        WinMsg = 0x000010a8
	LVM_GETINSERTMARKRECT        WinMsg = 0x000010a9
	LVM_SETINSERTMARKCOLOR       WinMsg = 0x000010aa
	LVM_GETINSERTMARKCOLOR       WinMsg = 0x000010ab
	LVM_SETINFOTIP               WinMsg = 0x000010ad
	LVM_GETSELECTEDCOLUMN        WinMsg = 0x000010ae
	LVM_ISGROUPVIEWENABLED       WinMsg = 0x000010af
	LVM_GETOUTLINECOLOR          WinMsg = 0x000010b0
	LVM_SETOUTLINECOLOR          WinMsg = 0x000010b1
	LVM_CANCELEDITLABEL          WinMsg = 0x000010b3
	LVM_MAPINDEXTOID             WinMsg = 0x000010b4
	LVM_MAPIDTOINDEX             WinMsg = 0x000010b5
	LVM_ISITEMVISIBLE            WinMsg = 0x000010b6
	LVM_GETEMPTYTEXT             WinMsg = 0x000010cc
	LVM_GETFOOTERRECT            WinMsg = 0x000010cd
	LVM_GETFOOTERINFO            WinMsg = 0x000010ce
	LVM_GETFOOTERITEMRECT        WinMsg = 0x000010cf
	LVM_GETFOOTERITEM            WinMsg = 0x000010d0
	LVM_GETITEMINDEXRECT         WinMsg = 0x000010d1
	LVM_SETITEMINDEXSTATE        WinMsg = 0x000010d2
	LVM_GETNEXTITEMINDEX         WinMsg = 0x000010d3
	// Tree View
	TVM_INSERTITEMA         WinMsg = 0x00001100
	TVM_INSERTITEMW         WinMsg = 0x00001132
	TVM_DELETEITEM          WinMsg = 0x00001101
	TVM_EXPAND              WinMsg = 0x00001102
	TVM_GETITEMRECT         WinMsg = 0x00001104
	TVM_GETCOUNT            WinMsg = 0x00001105
	TVM_GETINDENT           WinMsg = 0x00001106
	TVM_SETINDENT           WinMsg = 0x00001107
	TVM_GETIMAGELIST        WinMsg = 0x00001108
	TVM_SETIMAGELIST        WinMsg = 0x00001109
	TVM_GETNEXTITEM         WinMsg = 0x0000110a
	TVM_SELECTITEM          WinMsg = 0x0000110b
	TVM_GETITEMA            WinMsg = 0x0000110c
	TVM_GETITEMW            WinMsg = 0x0000113e
	TVM_SETITEMA            WinMsg = 0x0000110d
	TVM_SETITEMW            WinMsg = 0x0000113f
	TVM_EDITLABELA          WinMsg = 0x0000110e
	TVM_EDITLABELW          WinMsg = 0x00001141
	TVM_GETEDITCONTROL      WinMsg = 0x0000110f
	TVM_GETVISIBLECOUNT     WinMsg = 0x00001110
	TVM_HITTEST             WinMsg = 0x00001111
	TVM_CREATEDRAGIMAGE     WinMsg = 0x00001112
	TVM_SORTCHILDREN        WinMsg = 0x00001113
	TVM_ENSUREVISIBLE       WinMsg = 0x00001114
	TVM_SORTCHILDRENCB      WinMsg = 0x00001115
	TVM_ENDEDITLABELNOW     WinMsg = 0x00001116
	TVM_GETISEARCHSTRINGA   WinMsg = 0x00001117
	TVM_GETISEARCHSTRINGW   WinMsg = 0x00001140
	TVM_SETTOOLTIPS         WinMsg = 0x00001118
	TVM_GETTOOLTIPS         WinMsg = 0x00001119
	TVM_SETINSERTMARK       WinMsg = 0x0000111a
	TVM_SETITEMHEIGHT       WinMsg = 0x0000111b
	TVM_GETITEMHEIGHT       WinMsg = 0x0000111c
	TVM_SETBKCOLOR          WinMsg = 0x0000111d
	TVM_SETTEXTCOLOR        WinMsg = 0x0000111e
	TVM_GETBKCOLOR          WinMsg = 0x0000111f
	TVM_GETTEXTCOLOR        WinMsg = 0x00001120
	TVM_SETSCROLLTIME       WinMsg = 0x00001121
	TVM_GETSCROLLTIME       WinMsg = 0x00001122
	TVM_SETINSERTMARKCOLOR  WinMsg = 0x00001125
	TVM_GETINSERTMARKCOLOR  WinMsg = 0x00001126
	TVM_GETITEMSTATE        WinMsg = 0x00001127
	TVM_SETLINECOLOR        WinMsg = 0x00001128
	TVM_GETLINECOLOR        WinMsg = 0x00001129
	TVM_MAPACCIDTOHTREEITEM WinMsg = 0x0000112a
	TVM_MAPHTREEITEMTOACCID WinMsg = 0x0000112b
	TVM_SETEXTENDEDSTYLE    WinMsg = 0x0000112c
	TVM_GETEXTENDEDSTYLE    WinMsg = 0x0000112d
	TVM_SETAUTOSCROLLINFO   WinMsg = 0x0000113b
	TVM_GETSELECTEDCOUNT    WinMsg = 0x00001146
	TVM_SHOWINFOTIP         WinMsg = 0x00001147
	TVM_GETITEMPARTRECT     WinMsg = 0x00001148
	// Header Control
	HDM_GETITEMCOUNT           WinMsg = 0x00001200
	HDM_INSERTITEMA            WinMsg = 0x00001201
	HDM_INSERTITEMW            WinMsg = 0x0000120a
	HDM_DELETEITEM             WinMsg = 0x00001202
	HDM_GETITEMA               WinMsg = 0x00001203
	HDM_GETITEMW               WinMsg = 0x0000120b
	HDM_SETITEMA               WinMsg = 0x00001204
	HDM_SETITEMW               WinMsg = 0x0000120c
	HDM_LAYOUT                 WinMsg = 0x00001205
	HDM_HITTEST                WinMsg = 0x00001206
	HDM_GETITEMRECT            WinMsg = 0x00001207
	HDM_SETIMAGELIST           WinMsg = 0x00001208
	HDM_GETIMAGELIST           WinMsg = 0x00001209
	HDM_ORDERTOINDEX           WinMsg = 0x0000120f
	HDM_CREATEDRAGIMAGE        WinMsg = 0x00001210
	HDM_GETORDERARRAY          WinMsg = 0x00001211
	HDM_SETORDERARRAY          WinMsg = 0x00001212
	HDM_SETHOTDIVIDER          WinMsg = 0x00001213
	HDM_SETBITMAPMARGIN        WinMsg = 0x00001214
	HDM_GETBITMAPMARGIN        WinMsg = 0x00001215
	HDM_SETFILTERCHANGETIMEOUT WinMsg = 0x00001216
	HDM_EDITFILTER             WinMsg = 0x00001217
	HDM_CLEARFILTER            WinMsg = 0x00001218
	HDM_GETITEMDROPDOWNRECT    WinMsg = 0x00001219
	HDM_GETOVERFLOWRECT        WinMsg = 0x0000121a
	HDM_GETFOCUSEDITEM         WinMsg = 0x0000121b
	HDM_SETFOCUSEDITEM         WinMsg = 0x0000121c
	// Tab Control
	TCM_GETIMAGELIST     WinMsg = 0x00001302
	TCM_SETIMAGELIST     WinMsg = 0x00001303
	TCM_GETITEMCOUNT     WinMsg = 0x00001304
	TCM_GETITEMA         WinMsg = 0x00001305
	TCM_GETITEMW         WinMsg = 0x0000133c
	TCM_SETITEMA         WinMsg = 0x00001306
	TCM_SETITEMW         WinMsg = 0x0000133d
	TCM_INSERTITEMA      WinMsg = 0x00001307
	TCM_INSERTITEMW      WinMsg = 0x0000133e
	TCM_DELETEITEM       WinMsg = 0x00001308
	TCM_DELETEALLITEMS   WinMsg = 0x00001309
	TCM_GETITEMRECT      WinMsg = 0x0000130a
	TCM_GETCURSEL        WinMsg = 0x0000130b
	TCM_SETCURSEL        WinMsg = 0x0000130c
	TCM_HITTEST          WinMsg = 0x0000130d
	TCM_SETITEMEXTRA     WinMsg = 0x0000130e
	TCM_ADJUSTRECT       WinMsg = 0x00001328
	TCM_SETITEMSIZE      WinMsg = 0x00001329
	TCM_REMOVEIMAGE      WinMsg = 0x0000132a
	TCM_SETPADDING       WinMsg = 0x0000132b
	TCM_GETROWCOUNT      WinMsg = 0x0000132c
	TCM_GETTOOLTIPS      WinMsg = 0x0000132d
	TCM_SETTOOLTIPS      WinMsg = 0x0000132e
	TCM_GETCURFOCUS      WinMsg = 0x0000132f
	TCM_SETCURFOCUS      WinMsg = 0x00001330
	TCM_SETMINTABWIDTH   WinMsg = 0x00001331
	TCM_DESELECTALL      WinMsg = 0x00001332
	TCM_HIGHLIGHTITEM    WinMsg = 0x00001333
	TCM_SETEXTENDEDSTYLE WinMsg = 0x00001334
	TCM_GETEXTENDEDSTYLE WinMsg = 0x00001335
	// Pager Control
	PGM_SETCHILD       WinMsg = 0x00001401
	PGM_RECALCSIZE     WinMsg = 0x00001402
	PGM_FORWARDMOUSE   WinMsg = 0x00001403
	PGM_SETBKCOLOR     WinMsg = 0x00001404
	PGM_GETBKCOLOR     WinMsg = 0x00001405
	PGM_SETBORDER      WinMsg = 0x00001406
	PGM_GETBORDER      WinMsg = 0x00001407
	PGM_SETPOS         WinMsg = 0x00001408
	PGM_GETPOS         WinMsg = 0x00001409
	PGM_SETBUTTONSIZE  WinMsg = 0x0000140a
	PGM_GETBUTTONSIZE  WinMsg = 0x0000140b
	PGM_GETBUTTONSTATE WinMsg = 0x0000140c
	// Edit Control
	EM_SETCUEBANNER   WinMsg = 0x00001501
	EM_GETCUEBANNER   WinMsg = 0x00001502
	EM_SHOWBALLOONTIP WinMsg = 0x00001503
	EM_HIDEBALLOONTIP WinMsg = 0x00001504
	EM_SETHILITE      WinMsg = 0x00001505
	EM_GETHILITE      WinMsg = 0x00001506
	// Button Control
	BCM_GETIDEALSIZE     WinMsg = 0x00001601
	BCM_SETIMAGELIST     WinMsg = 0x00001602
	BCM_GETIMAGELIST     WinMsg = 0x00001603
	BCM_SETTEXTMARGIN    WinMsg = 0x00001604
	BCM_GETTEXTMARGIN    WinMsg = 0x00001605
	BCM_SETDROPDOWNSTATE WinMsg = 0x00001606
	BCM_SETSPLITINFO     WinMsg = 0x00001607
	BCM_GETSPLITINFO     WinMsg = 0x00001608
	BCM_SETNOTE          WinMsg = 0x00001609
	BCM_GETNOTE          WinMsg = 0x0000160a
	BCM_GETNOTELENGTH    WinMsg = 0x0000160b
	BCM_SETSHIELD        WinMsg = 0x0000160c
	// Combobox Control
	CB_SETMINVISIBLE WinMsg = 0x00001701
	CB_GETMINVISIBLE WinMsg = 0x00001702
	CB_SETCUEBANNER  WinMsg = 0x00001703
	CB_GETCUEBANNER  WinMsg = 0x00001704
	// Common Control Shared
	CCM_SETBKCOLOR       WinMsg = 0x00002001
	CCM_SETCOLORSCHEME   WinMsg = 0x00002002
	CCM_GETCOLORSCHEME   WinMsg = 0x00002003
	CCM_GETDROPTARGET    WinMsg = 0x00002004
	CCM_SETUNICODEFORMAT WinMsg = 0x00002005
	CCM_GETUNICODEFORMAT WinMsg = 0x00002006
	CCM_SETVERSION       WinMsg = 0x00002007
	CCM_GETVERSION       WinMsg = 0x00002008
	CCM_SETNOTIFYWINDOW  WinMsg = 0x00002009
	CCM_SETWINDOWTHEME   WinMsg = 0x0000200b
	CCM_DPISCALE         WinMsg = 0x0000200c
	WM_RASDIALEVENT      WinMsg = 0x0000cccd
)

// MSG
type MSG struct {
	Hwnd    HWND
	Message WinMsg
	WParam  WPARAM
	LParam  LPARAM
	Time    DWORD
	Pt      POINT
}

type LPMSG *MSG

// ImageType
type ImageType UINT

const (
	IMAGE_BITMAP      ImageType = 0
	IMAGE_ICON        ImageType = 1
	IMAGE_CURSOR      ImageType = 2
	IMAGE_ENHMETAFILE ImageType = 3
)

// LRFlags
type LRFlags UINT

const (
	LR_DEFAULTCOLOR     LRFlags = 0x00000000
	LR_MONOCHROME       LRFlags = 0x00000001
	LR_COLOR            LRFlags = 0x00000002
	LR_COPYRETURNORG    LRFlags = 0x00000004
	LR_COPYDELETEORG    LRFlags = 0x00000008
	LR_LOADFROMFILE     LRFlags = 0x00000010
	LR_LOADTRANSPARENT  LRFlags = 0x00000020
	LR_DEFAULTSIZE      LRFlags = 0x00000040
	LR_VGACOLOR         LRFlags = 0x00000080
	LR_LOADMAP3DCOLORS  LRFlags = 0x00001000
	LR_CREATEDIBSECTION LRFlags = 0x00002000
	LR_COPYFROMRESOURCE LRFlags = 0x00004000
	LR_SHARED           LRFlags = 0x00008000
)

// SBType
type SBType INT

const (
	SB_HORZ SBType = 0
	SB_VERT SBType = 1
	SB_CTL  SBType = 2
	SB_BOTH SBType = 3
)

// IsTextUnicodeFlags
type IsTextUnicodeFlags UINT

const (
	IS_TEXT_UNICODE_ASCII16            IsTextUnicodeFlags = 0x0001
	IS_TEXT_UNICODE_REVERSE_ASCII16    IsTextUnicodeFlags = 0x0010
	IS_TEXT_UNICODE_STATISTICS         IsTextUnicodeFlags = 0x0002
	IS_TEXT_UNICODE_REVERSE_STATISTICS IsTextUnicodeFlags = 0x0020
	IS_TEXT_UNICODE_CONTROLS           IsTextUnicodeFlags = 0x0004
	IS_TEXT_UNICODE_REVERSE_CONTROLS   IsTextUnicodeFlags = 0x0040
	IS_TEXT_UNICODE_SIGNATURE          IsTextUnicodeFlags = 0x0008
	IS_TEXT_UNICODE_REVERSE_SIGNATURE  IsTextUnicodeFlags = 0x0080
	IS_TEXT_UNICODE_ILLEGAL_CHARS      IsTextUnicodeFlags = 0x0100
	IS_TEXT_UNICODE_ODD_LENGTH         IsTextUnicodeFlags = 0x0200
	IS_TEXT_UNICODE_DBCS_LEADBYTE      IsTextUnicodeFlags = 0x0400
	IS_TEXT_UNICODE_NULL_BYTES         IsTextUnicodeFlags = 0x1000
	IS_TEXT_UNICODE_UNICODE_MASK       IsTextUnicodeFlags = 0x000F
	IS_TEXT_UNICODE_REVERSE_MASK       IsTextUnicodeFlags = 0x00F0
	IS_TEXT_UNICODE_NOT_UNICODE_MASK   IsTextUnicodeFlags = 0x0F00
	IS_TEXT_UNICODE_NOT_ASCII_MASK     IsTextUnicodeFlags = 0xF000
)

// WaitTimeout
type WaitTimeout DWORD

const (
	INFINITE WaitTimeout = 0xFFFFFFFF
)

// BorderFlag
type BorderFlag UINT

const (
	BF_LEFT                    BorderFlag = 0x0001
	BF_TOP                     BorderFlag = 0x0002
	BF_RIGHT                   BorderFlag = 0x0004
	BF_BOTTOM                  BorderFlag = 0x0008
	BF_TOPLEFT                 BorderFlag = 0x0003
	BF_TOPRIGHT                BorderFlag = 0x0006
	BF_BOTTOMLEFT              BorderFlag = 0x0009
	BF_BOTTOMRIGHT             BorderFlag = 0x000c
	BF_RECT                    BorderFlag = 0x000f
	BF_DIAGONAL                BorderFlag = 0x0010
	BF_DIAGONAL_ENDTOPRIGHT    BorderFlag = 0x0016
	BF_DIAGONAL_ENDTOPLEFT     BorderFlag = 0x0013
	BF_DIAGONAL_ENDBOTTOMLEFT  BorderFlag = 0x0019
	BF_DIAGONAL_ENDBOTTOMRIGHT BorderFlag = 0x001c
	BF_MIDDLE                  BorderFlag = 0x0800
	BF_SOFT                    BorderFlag = 0x1000
	BF_ADJUST                  BorderFlag = 0x2000
	BF_FLAT                    BorderFlag = 0x4000
	BF_MONO                    BorderFlag = 0x8000
)

// SysColorIndex
type SysColorIndex INT

const (
	COLOR_SCROLLBAR               SysColorIndex = 0
	COLOR_BACKGROUND              SysColorIndex = 1
	COLOR_ACTIVECAPTION           SysColorIndex = 2
	COLOR_INACTIVECAPTION         SysColorIndex = 3
	COLOR_MENU                    SysColorIndex = 4
	COLOR_WINDOW                  SysColorIndex = 5
	COLOR_WINDOWFRAME             SysColorIndex = 6
	COLOR_MENUTEXT                SysColorIndex = 7
	COLOR_WINDOWTEXT              SysColorIndex = 8
	COLOR_CAPTIONTEXT             SysColorIndex = 9
	COLOR_ACTIVEBORDER            SysColorIndex = 10
	COLOR_INACTIVEBORDER          SysColorIndex = 11
	COLOR_APPWORKSPACE            SysColorIndex = 12
	COLOR_HIGHLIGHT               SysColorIndex = 13
	COLOR_HIGHLIGHTTEXT           SysColorIndex = 14
	COLOR_BTNFACE                 SysColorIndex = 15
	COLOR_BTNSHADOW               SysColorIndex = 16
	COLOR_GRAYTEXT                SysColorIndex = 17
	COLOR_BTNTEXT                 SysColorIndex = 18
	COLOR_INACTIVECAPTIONTEXT     SysColorIndex = 19
	COLOR_BTNHIGHLIGHT            SysColorIndex = 20
	COLOR_3DDKSHADOW              SysColorIndex = 21
	COLOR_3DLIGHT                 SysColorIndex = 22
	COLOR_INFOTEXT                SysColorIndex = 23
	COLOR_INFOBK                  SysColorIndex = 24
	COLOR_HOTLIGHT                SysColorIndex = 26
	COLOR_GRADIENTACTIVECAPTION   SysColorIndex = 27
	COLOR_GRADIENTINACTIVECAPTION SysColorIndex = 28
	COLOR_MENUHILIGHT             SysColorIndex = 29
	COLOR_MENUBAR                 SysColorIndex = 30
)

// BorderEdge
type BorderEdge UINT

const (
	BDR_RAISEDOUTER BorderEdge = 0x0001
	BDR_SUNKENOUTER BorderEdge = 0x0002
	BDR_RAISEDINNER BorderEdge = 0x0004
	BDR_SUNKENINNER BorderEdge = 0x0008
	EDGE_RAISED     BorderEdge = 0x0005
	EDGE_SUNKEN     BorderEdge = 0x000a
	EDGE_ETCHED     BorderEdge = 0x0006
	EDGE_BUMP       BorderEdge = 0x0009
)

// DrawTextFlags
type DrawTextFlags DWORD

const (
	DT_TOP                  DrawTextFlags = 0x00000000
	DT_LEFT                 DrawTextFlags = 0x00000000
	DT_CENTER               DrawTextFlags = 0x00000001
	DT_RIGHT                DrawTextFlags = 0x00000002
	DT_VCENTER              DrawTextFlags = 0x00000004
	DT_BOTTOM               DrawTextFlags = 0x00000008
	DT_WORDBREAK            DrawTextFlags = 0x00000010
	DT_SINGLELINE           DrawTextFlags = 0x00000020
	DT_EXPANDTABS           DrawTextFlags = 0x00000040
	DT_TABSTOP              DrawTextFlags = 0x00000080
	DT_NOCLIP               DrawTextFlags = 0x00000100
	DT_EXTERNALLEADING      DrawTextFlags = 0x00000200
	DT_CALCRECT             DrawTextFlags = 0x00000400
	DT_NOPREFIX             DrawTextFlags = 0x00000800
	DT_INTERNAL             DrawTextFlags = 0x00001000
	DT_EDITCONTROL          DrawTextFlags = 0x00002000
	DT_PATH_ELLIPSIS        DrawTextFlags = 0x00004000
	DT_END_ELLIPSIS         DrawTextFlags = 0x00008000
	DT_MODIFYSTRING         DrawTextFlags = 0x00010000
	DT_RTLREADING           DrawTextFlags = 0x00020000
	DT_WORD_ELLIPSIS        DrawTextFlags = 0x00040000
	DT_NOFULLWIDTHCHARBREAK DrawTextFlags = 0x00080000
	DT_HIDEPREFIX           DrawTextFlags = 0x00100000
	DT_PREFIXONLY           DrawTextFlags = 0x00200000
)

// RECT
type RECT struct {
	Left   LONG
	Top    LONG
	Right  LONG
	Bottom LONG
}

type (
	LPRECT  *RECT
	LPCRECT LPRECT

	// RECTL
	RECTL    RECT
	LPCRECTL *RECTL
)

// CodePageEnum
type CodePageEnum UINT

const (
	CP_ACP              CodePageEnum = 0
	CP_OEMCP            CodePageEnum = 1
	CP_MACCP            CodePageEnum = 2
	CP_THREAD_ACP       CodePageEnum = 3
	CP_SYMBOL           CodePageEnum = 42
	MS_DOS_Latin_US     CodePageEnum = 437
	Thai                CodePageEnum = 874
	Japanese            CodePageEnum = 932
	Chinese_Simplified  CodePageEnum = 936
	Korean              CodePageEnum = 949
	Chinese_Traditional CodePageEnum = 950
	Unicode_UTF_16_LE   CodePageEnum = 1200
	Unicode_UTF_16_BE   CodePageEnum = 1201
	Central_European    CodePageEnum = 1250
	Cyrillic            CodePageEnum = 1251
	Western_European    CodePageEnum = 1252
	Greek               CodePageEnum = 1253
	Turkish             CodePageEnum = 1254
	Hebrew              CodePageEnum = 1255
	Arabic              CodePageEnum = 1256
	Baltic              CodePageEnum = 1257
	Vietnamese          CodePageEnum = 1258
	CP_UTF7             CodePageEnum = 65000
	CP_UTF8             CodePageEnum = 65001
)

// EXCEPTION_FLAGS
type EXCEPTION_FLAGS DWORD

// EXCEPTION_RECORD
type EXCEPTION_RECORD struct {
	ExceptionCode        ExceptionCode
	ExceptionFlags       DWORD
	ExceptionRecord      LPVOID
	ExceptionAddress     PVOID
	NumberParameters     DWORD
	ExceptionInformation [15]ULONG_PTR
}

type PEXCEPTION_RECORD *EXCEPTION_RECORD

// EXCEPTION_POINTERS
type EXCEPTION_POINTERS struct {
	ExceptionRecord PEXCEPTION_RECORD
	ContextRecord   PCONTEXT
}

type (
	PEXCEPTION_POINTERS  *EXCEPTION_POINTERS
	LPEXCEPTION_POINTERS *EXCEPTION_POINTERS
)

// SP_DEVINFO_DATA
type SP_DEVINFO_DATA struct {
	CbSize    DWORD
	ClassGuid GUID
	DevInst   DWORD
	Reserved  ULONG_PTR
}

type PSP_DEVINFO_DATA *SP_DEVINFO_DATA

// PROCESS_INFORMATION
type PROCESS_INFORMATION struct {
	HProcess    HANDLE
	HThread     HANDLE
	DwProcessId DWORD
	DwThreadId  DWORD
}

type LPPROCESS_INFORMATION *PROCESS_INFORMATION

// ShowWindowCmd
type ShowWindowCmd INT

const (
	SW_HIDE            ShowWindowCmd = 0
	SW_SHOWNORMAL      ShowWindowCmd = 1
	SW_NORMAL          ShowWindowCmd = 1
	SW_SHOWMINIMIZED   ShowWindowCmd = 2
	SW_SHOWMAXIMIZED   ShowWindowCmd = 3
	SW_MAXIMIZE        ShowWindowCmd = 3
	SW_SHOWNOACTIVATE  ShowWindowCmd = 4
	SW_SHOW            ShowWindowCmd = 5
	SW_MINIMIZE        ShowWindowCmd = 6
	SW_SHOWMINNOACTIVE ShowWindowCmd = 7
	SW_SHOWNA          ShowWindowCmd = 8
	SW_RESTORE         ShowWindowCmd = 9
	SW_SHOWDEFAULT     ShowWindowCmd = 10
	SW_FORCEMINIMIZE   ShowWindowCmd = 11
)

// STARTUPINFO_ShowWindow
type STARTUPINFO_ShowWindow WORD

// STARTUPINFO_Flags
type STARTUPINFO_Flags DWORD

const (
	STARTF_USESHOWWINDOW    STARTUPINFO_Flags = 0x00000001
	STARTF_USESIZE          STARTUPINFO_Flags = 0x00000002
	STARTF_USEPOSITION      STARTUPINFO_Flags = 0x00000004
	STARTF_USECOUNTCHARS    STARTUPINFO_Flags = 0x00000008
	STARTF_USEFILLATTRIBUTE STARTUPINFO_Flags = 0x00000010
	STARTF_RUNFULLSCREEN    STARTUPINFO_Flags = 0x00000020
	STARTF_FORCEONFEEDBACK  STARTUPINFO_Flags = 0x00000040
	STARTF_FORCEOFFFEEDBACK STARTUPINFO_Flags = 0x00000080
	STARTF_USESTDHANDLES    STARTUPINFO_Flags = 0x00000100
	STARTF_USEHOTKEY        STARTUPINFO_Flags = 0x00000200
	STARTF_TITLEISLINKNAME  STARTUPINFO_Flags = 0x00000800
	STARTF_TITLEISAPPID     STARTUPINFO_Flags = 0x00001000
	STARTF_PREVENTPINNING   STARTUPINFO_Flags = 0x00002000
)

// STARTUPINFO
type STARTUPINFO struct {
	Cb              DWORD
	LpReserved      LPWSTR
	LpDesktop       LPWSTR
	LpTitle         LPWSTR
	DwX             DWORD
	DwY             DWORD
	DwXSize         DWORD
	DwYSize         DWORD
	DwXCountChars   DWORD
	DwYCountChars   DWORD
	DwFillAttribute DWORD
	DwFlags         STARTUPINFO_Flags
	WShowWindow     STARTUPINFO_ShowWindow
	CbReserved2     WORD
	LpReserved2     LPBYTE
	HStdInput       HANDLE
	HStdOutput      HANDLE
	HStdError       HANDLE
}

type LPSTARTUPINFO *STARTUPINFO

// FILE_SEGMENT_ELEMENT
type FILE_SEGMENT_ELEMENT struct {
	Buffer    PVOID64
	Alignment ULONGLONG
}

// OVERLAPPED_u_s
type OVERLAPPED_u_s struct {
	Offset     DWORD
	OffsetHigh DWORD
}

// OVERLAPPED_u
type OVERLAPPED_u struct {
	OVERLAPPED_u_s
}

// OVERLAPPED
type OVERLAPPED struct {
	Internal     ULONG_PTR
	InternalHigh ULONG_PTR
	Offset       OVERLAPPED_u
	HEvent       HANDLE
}
type LPOVERLAPPED *OVERLAPPED

// ReparsePoint
type ReparsePoint DWORD

const (
	IO_REPARSE_TAG_MOUNT_POINT ReparsePoint = 0xA0000003
	IO_REPARSE_TAG_HSM         ReparsePoint = 0xC0000004
	IO_REPARSE_TAG_HSM2        ReparsePoint = 0x80000006
	IO_REPARSE_TAG_SIS         ReparsePoint = 0x80000007
	IO_REPARSE_TAG_WIM         ReparsePoint = 0x80000008
	IO_REPARSE_TAG_CSV         ReparsePoint = 0x80000009
	IO_REPARSE_TAG_DFS         ReparsePoint = 0x8000000A
	IO_REPARSE_TAG_SYMLINK     ReparsePoint = 0xA000000C
	IO_REPARSE_TAG_DFSR        ReparsePoint = 0x80000012
	IO_REPARSE_TAG_DEDUP       ReparsePoint = 0x80000013
	IO_REPARSE_TAG_NFS         ReparsePoint = 0x80000014
)

// FileAttributes
type FileAttributes DWORD

const (
	FILE_ATTRIBUTE_READONLY            FileAttributes = 0x00000001
	FILE_ATTRIBUTE_HIDDEN              FileAttributes = 0x00000002
	FILE_ATTRIBUTE_SYSTEM              FileAttributes = 0x00000004
	FILE_ATTRIBUTE_DIRECTORY           FileAttributes = 0x00000010
	FILE_ATTRIBUTE_ARCHIVE             FileAttributes = 0x00000020
	FILE_ATTRIBUTE_DEVICE              FileAttributes = 0x00000040
	FILE_ATTRIBUTE_NORMAL              FileAttributes = 0x00000080
	FILE_ATTRIBUTE_TEMPORARY           FileAttributes = 0x00000100
	FILE_ATTRIBUTE_SPARSE_FILE         FileAttributes = 0x00000200
	FILE_ATTRIBUTE_REPARSE_POINT       FileAttributes = 0x00000400
	FILE_ATTRIBUTE_COMPRESSED          FileAttributes = 0x00000800
	FILE_ATTRIBUTE_OFFLINE             FileAttributes = 0x00001000
	FILE_ATTRIBUTE_NOT_CONTENT_INDEXED FileAttributes = 0x00002000
	FILE_ATTRIBUTE_ENCRYPTED           FileAttributes = 0x00004000
	FILE_ATTRIBUTE_VIRTUAL             FileAttributes = 0x00010000
	FILE_ATTRIBUTE_NO_SCRUB_DATA       FileAttributes = 0x00020000
)

// WIN32_FIND_DATA
type WIN32_FIND_DATA struct {
	DwFileAttributes   FileAttributes
	FtCreationTime     FILETIME
	FtLastAccessTime   FILETIME
	FtLastWriteTime    FILETIME
	NFileSizeHigh      DWORD
	NFileSizeLow       DWORD
	DwReserved0        ReparsePoint
	DwReserved1        DWORD
	CFileName          [260]WCHAR
	CAlternateFileName [14]WCHAR
}

type LPWIN32_FIND_DATA *WIN32_FIND_DATA

// WIN32_FIND_DATAA
type WIN32_FIND_DATAA struct {
	DwFileAttributes   FileAttributes
	FtCreationTime     FILETIME
	FtLastAccessTime   FILETIME
	FtLastWriteTime    FILETIME
	NFileSizeHigh      DWORD
	NFileSizeLow       DWORD
	DwReserved0        ReparsePoint
	DwReserved1        DWORD
	CFileName          [260]CHAR
	CAlternateFileName [14]CHAR
}

// WIN32_FIND_DATAW
type WIN32_FIND_DATAW struct {
	DwFileAttributes   FileAttributes
	FtCreationTime     FILETIME
	FtLastAccessTime   FILETIME
	FtLastWriteTime    FILETIME
	NFileSizeHigh      DWORD
	NFileSizeLow       DWORD
	DwReserved0        ReparsePoint
	DwReserved1        DWORD
	CFileName          [260]WCHAR
	CAlternateFileName [14]WCHAR
}

// TIME_ZONE_INFORMATION
type TIME_ZONE_INFORMATION struct {
	Bias         LONG
	StandardName [32]WCHAR
	StandardDate SYSTEMTIME
	StandardBias LONG
	DaylightName [32]WCHAR
	DaylightDate SYSTEMTIME
	DaylightBias LONG
}

type LPTIME_ZONE_INFORMATION *TIME_ZONE_INFORMATION

// OFSTRUCT
type OFSTRUCT struct {
	CBytes     BYTE
	FFixedDisk BYTE
	NErrCode   WORD
	Reserved1  WORD
	Reserved2  WORD
	SzPathName [128]CHAR
}

type LPOFSTRUCT *OFSTRUCT

// PROCESSOR_ARCHITECTURE
type PROCESSOR_ARCHITECTURE WORD

const (
	PROCESSOR_ARCHITECTURE_INTEL   PROCESSOR_ARCHITECTURE = 0
	PROCESSOR_ARCHITECTURE_IA64    PROCESSOR_ARCHITECTURE = 6
	PROCESSOR_ARCHITECTURE_AMD64   PROCESSOR_ARCHITECTURE = 9
	PROCESSOR_ARCHITECTURE_UNKNOWN PROCESSOR_ARCHITECTURE = 0xFFFF
)

// MINIDUMP_TYPE
type MINIDUMP_TYPE DWORD

const (
	MiniDumpNormal                         MINIDUMP_TYPE = 0x00000000
	MiniDumpWithDataSegs                   MINIDUMP_TYPE = 0x00000001
	MiniDumpWithFullMemory                 MINIDUMP_TYPE = 0x00000002
	MiniDumpWithHandleData                 MINIDUMP_TYPE = 0x00000004
	MiniDumpFilterMemory                   MINIDUMP_TYPE = 0x00000008
	MiniDumpScanMemory                     MINIDUMP_TYPE = 0x00000010
	MiniDumpWithUnloadedModules            MINIDUMP_TYPE = 0x00000020
	MiniDumpWithIndirectlyReferencedMemory MINIDUMP_TYPE = 0x00000040
	MiniDumpFilterModulePaths              MINIDUMP_TYPE = 0x00000080
	MiniDumpWithProcessThreadData          MINIDUMP_TYPE = 0x00000100
	MiniDumpWithPrivateReadWriteMemory     MINIDUMP_TYPE = 0x00000200
	MiniDumpWithoutOptionalData            MINIDUMP_TYPE = 0x00000400
	MiniDumpWithFullMemoryInfo             MINIDUMP_TYPE = 0x00000800
	MiniDumpWithThreadInfo                 MINIDUMP_TYPE = 0x00001000
	MiniDumpWithCodeSegs                   MINIDUMP_TYPE = 0x00002000
	MiniDumpWithoutAuxiliaryState          MINIDUMP_TYPE = 0x00004000
	MiniDumpWithFullAuxiliaryState         MINIDUMP_TYPE = 0x00008000
	MiniDumpWithPrivateWriteCopyMemory     MINIDUMP_TYPE = 0x00010000
	MiniDumpIgnoreInaccessibleMemory       MINIDUMP_TYPE = 0x00020000
	MiniDumpWithTokenInformation           MINIDUMP_TYPE = 0x00040000
	MiniDumpWithModuleHeaders              MINIDUMP_TYPE = 0x00080000
	MiniDumpFilterTriage                   MINIDUMP_TYPE = 0x00100000
)

// THREAD_WRITE_FLAGS
type THREAD_WRITE_FLAGS DWORD

const (
	ThreadWriteThread            THREAD_WRITE_FLAGS = 0x0001
	ThreadWriteStack             THREAD_WRITE_FLAGS = 0x0002
	ThreadWriteContext           THREAD_WRITE_FLAGS = 0x0004
	ThreadWriteBackingStore      THREAD_WRITE_FLAGS = 0x0008
	ThreadWriteInstructionWindow THREAD_WRITE_FLAGS = 0x0010
	ThreadWriteThreadData        THREAD_WRITE_FLAGS = 0x0020
	ThreadWriteThreadInfo        THREAD_WRITE_FLAGS = 0x0040
)

// MODULE_WRITE_FLAGS
type MODULE_WRITE_FLAGS DWORD

const (
	ModuleWriteModule        MODULE_WRITE_FLAGS = 0x0001
	ModuleWriteDataSeg       MODULE_WRITE_FLAGS = 0x0002
	ModuleWriteMiscRecord    MODULE_WRITE_FLAGS = 0x0004
	ModuleWriteCvRecord      MODULE_WRITE_FLAGS = 0x0008
	ModuleReferencedByMemory MODULE_WRITE_FLAGS = 0x0010
	ModuleWriteTlsData       MODULE_WRITE_FLAGS = 0x0020
	ModuleWriteCodeSegs      MODULE_WRITE_FLAGS = 0x0040
)

// VirtKeyCode
type VirtKeyCode WORD

const (
	VK_LBUTTON             VirtKeyCode = 0x01
	VK_RBUTTON             VirtKeyCode = 0x02
	VK_CANCEL              VirtKeyCode = 0x03
	VK_MBUTTON             VirtKeyCode = 0x04
	VK_XBUTTON1            VirtKeyCode = 0x05
	VK_XBUTTON2            VirtKeyCode = 0x06
	VK_BACK                VirtKeyCode = 0x08
	VK_TAB                 VirtKeyCode = 0x09
	VK_CLEAR               VirtKeyCode = 0x0C
	VK_RETURN              VirtKeyCode = 0x0D
	VK_SHIFT               VirtKeyCode = 0x10
	VK_CONTROL             VirtKeyCode = 0x11
	VK_MENU                VirtKeyCode = 0x12
	VK_PAUSE               VirtKeyCode = 0x13
	VK_CAPITAL             VirtKeyCode = 0x14
	VK_KANA                VirtKeyCode = 0x15
	VK_JUNJA               VirtKeyCode = 0x17
	VK_FINAL               VirtKeyCode = 0x18
	VK_KANJI               VirtKeyCode = 0x19
	VK_ESCAPE              VirtKeyCode = 0x1B
	VK_CONVERT             VirtKeyCode = 0x1C
	VK_NONCONVERT          VirtKeyCode = 0x1D
	VK_ACCEPT              VirtKeyCode = 0x1E
	VK_MODECHANGE          VirtKeyCode = 0x1F
	VK_SPACE               VirtKeyCode = 0x20
	VK_PRIOR               VirtKeyCode = 0x21
	VK_NEXT                VirtKeyCode = 0x22
	VK_END                 VirtKeyCode = 0x23
	VK_HOME                VirtKeyCode = 0x24
	VK_LEFT                VirtKeyCode = 0x25
	VK_UP                  VirtKeyCode = 0x26
	VK_RIGHT               VirtKeyCode = 0x27
	VK_DOWN                VirtKeyCode = 0x28
	VK_SELECT              VirtKeyCode = 0x29
	VK_PRINT               VirtKeyCode = 0x2A
	VK_EXECUTE             VirtKeyCode = 0x2B
	VK_SNAPSHOT            VirtKeyCode = 0x2C
	VK_INSERT              VirtKeyCode = 0x2D
	VK_DELETE              VirtKeyCode = 0x2E
	VK_HELP                VirtKeyCode = 0x2F
	VK_LWIN                VirtKeyCode = 0x5B
	VK_RWIN                VirtKeyCode = 0x5C
	VK_APPS                VirtKeyCode = 0x5D
	VK_SLEEP               VirtKeyCode = 0x5F
	VK_NUMPAD0             VirtKeyCode = 0x60
	VK_NUMPAD1             VirtKeyCode = 0x61
	VK_NUMPAD2             VirtKeyCode = 0x62
	VK_NUMPAD3             VirtKeyCode = 0x63
	VK_NUMPAD4             VirtKeyCode = 0x64
	VK_NUMPAD5             VirtKeyCode = 0x65
	VK_NUMPAD6             VirtKeyCode = 0x66
	VK_NUMPAD7             VirtKeyCode = 0x67
	VK_NUMPAD8             VirtKeyCode = 0x68
	VK_NUMPAD9             VirtKeyCode = 0x69
	VK_MULTIPLY            VirtKeyCode = 0x6A
	VK_ADD                 VirtKeyCode = 0x6B
	VK_SEPARATOR           VirtKeyCode = 0x6C
	VK_SUBTRACT            VirtKeyCode = 0x6D
	VK_DECIMAL             VirtKeyCode = 0x6E
	VK_DIVIDE              VirtKeyCode = 0x6F
	VK_F1                  VirtKeyCode = 0x70
	VK_F2                  VirtKeyCode = 0x71
	VK_F3                  VirtKeyCode = 0x72
	VK_F4                  VirtKeyCode = 0x73
	VK_F5                  VirtKeyCode = 0x74
	VK_F6                  VirtKeyCode = 0x75
	VK_F7                  VirtKeyCode = 0x76
	VK_F8                  VirtKeyCode = 0x77
	VK_F9                  VirtKeyCode = 0x78
	VK_F10                 VirtKeyCode = 0x79
	VK_F11                 VirtKeyCode = 0x7A
	VK_F12                 VirtKeyCode = 0x7B
	VK_F13                 VirtKeyCode = 0x7C
	VK_F14                 VirtKeyCode = 0x7D
	VK_F15                 VirtKeyCode = 0x7E
	VK_F16                 VirtKeyCode = 0x7F
	VK_F17                 VirtKeyCode = 0x80
	VK_F18                 VirtKeyCode = 0x81
	VK_F19                 VirtKeyCode = 0x82
	VK_F20                 VirtKeyCode = 0x83
	VK_F21                 VirtKeyCode = 0x84
	VK_F22                 VirtKeyCode = 0x85
	VK_F23                 VirtKeyCode = 0x86
	VK_F24                 VirtKeyCode = 0x87
	VK_NUMLOCK             VirtKeyCode = 0x90
	VK_SCROLL              VirtKeyCode = 0x91
	VK_OEM_NEC_EQUAL       VirtKeyCode = 0x92
	VK_LSHIFT              VirtKeyCode = 0xA0
	VK_RSHIFT              VirtKeyCode = 0xA1
	VK_LCONTROL            VirtKeyCode = 0xA2
	VK_RCONTROL            VirtKeyCode = 0xA3
	VK_LMENU               VirtKeyCode = 0xA4
	VK_RMENU               VirtKeyCode = 0xA5
	VK_BROWSER_BACK        VirtKeyCode = 0xA6
	VK_BROWSER_FORWARD     VirtKeyCode = 0xA7
	VK_BROWSER_REFRESH     VirtKeyCode = 0xA8
	VK_BROWSER_STOP        VirtKeyCode = 0xA9
	VK_BROWSER_SEARCH      VirtKeyCode = 0xAA
	VK_BROWSER_FAVORITES   VirtKeyCode = 0xAB
	VK_BROWSER_HOME        VirtKeyCode = 0xAC
	VK_VOLUME_MUTE         VirtKeyCode = 0xAD
	VK_VOLUME_DOWN         VirtKeyCode = 0xAE
	VK_VOLUME_UP           VirtKeyCode = 0xAF
	VK_MEDIA_NEXT_TRACK    VirtKeyCode = 0xB0
	VK_MEDIA_PREV_TRACK    VirtKeyCode = 0xB1
	VK_MEDIA_STOP          VirtKeyCode = 0xB2
	VK_MEDIA_PLAY_PAUSE    VirtKeyCode = 0xB3
	VK_LAUNCH_MAIL         VirtKeyCode = 0xB4
	VK_LAUNCH_MEDIA_SELECT VirtKeyCode = 0xB5
	VK_LAUNCH_APP1         VirtKeyCode = 0xB6
	VK_LAUNCH_APP2         VirtKeyCode = 0xB7
	VK_OEM_1               VirtKeyCode = 0xBA
	VK_OEM_PLUS            VirtKeyCode = 0xBB
	VK_OEM_COMMA           VirtKeyCode = 0xBC
	VK_OEM_MINUS           VirtKeyCode = 0xBD
	VK_OEM_PERIOD          VirtKeyCode = 0xBE
	VK_OEM_2               VirtKeyCode = 0xBF
	VK_OEM_3               VirtKeyCode = 0xC0
	VK_OEM_4               VirtKeyCode = 0xDB
	VK_OEM_5               VirtKeyCode = 0xDC
	VK_OEM_6               VirtKeyCode = 0xDD
	VK_OEM_7               VirtKeyCode = 0xDE
	VK_OEM_8               VirtKeyCode = 0xDF
	VK_OEM_AX              VirtKeyCode = 0xE1
	VK_OEM_102             VirtKeyCode = 0xE2
	VK_ICO_HELP            VirtKeyCode = 0xE3
	VK_ICO_00              VirtKeyCode = 0xE4
	VK_PROCESSKEY          VirtKeyCode = 0xE5
	VK_ICO_CLEAR           VirtKeyCode = 0xE6
	VK_PACKET              VirtKeyCode = 0xE7
	VK_OEM_RESET           VirtKeyCode = 0xE9
	VK_OEM_JUMP            VirtKeyCode = 0xEA
	VK_OEM_PA1             VirtKeyCode = 0xEB
	VK_OEM_PA2             VirtKeyCode = 0xEC
	VK_OEM_PA3             VirtKeyCode = 0xED
	VK_OEM_WSCTRL          VirtKeyCode = 0xEE
	VK_OEM_CUSEL           VirtKeyCode = 0xEF
	VK_OEM_ATTN            VirtKeyCode = 0xF0
	VK_OEM_FINISH          VirtKeyCode = 0xF1
	VK_OEM_COPY            VirtKeyCode = 0xF2
	VK_OEM_AUTO            VirtKeyCode = 0xF3
	VK_OEM_ENLW            VirtKeyCode = 0xF4
	VK_OEM_BACKTAB         VirtKeyCode = 0xF5
	VK_ATTN                VirtKeyCode = 0xF6
	VK_CRSEL               VirtKeyCode = 0xF7
	VK_EXSEL               VirtKeyCode = 0xF8
	VK_EREOF               VirtKeyCode = 0xF9
	VK_PLAY                VirtKeyCode = 0xFA
	VK_ZOOM                VirtKeyCode = 0xFB
	VK_NONAME              VirtKeyCode = 0xFC
	VK_PA1                 VirtKeyCode = 0xFD
	VK_OEM_CLEAR           VirtKeyCode = 0xFE
)

// VER_PLATFORM
type VER_PLATFORM DWORD

const (
	VER_PLATFORM_WIN32s        VER_PLATFORM = 0
	VER_PLATFORM_WIN32_WINDOWS VER_PLATFORM = 1
	VER_PLATFORM_WIN32_NT      VER_PLATFORM = 2
)

// OSVERSIONINFO
type OSVERSIONINFO struct {
	DwOSVersionInfoSize DWORD
	DwMajorVersion      DWORD
	DwMinorVersion      DWORD
	DwBuildNumber       DWORD
	DwPlatformId        VER_PLATFORM
	SzCSDVersion        [128]WCHAR
}
type LPOSVERSIONINFO *OSVERSIONINFO

// RTL_OSVERSIONINFOW
type RTL_OSVERSIONINFOW struct {
	DwOSVersionInfoSize DWORD
	DwMajorVersion      DWORD
	DwMinorVersion      DWORD
	DwBuildNumber       DWORD
	DwPlatformId        VER_PLATFORM
	SzCSDVersion        [128]WCHAR
}
type PRTL_OSVERSIONINFOW *RTL_OSVERSIONINFOW

// VER_SUITE
type VER_SUITE WORD

const (
	VER_SUITE_SMALLBUSINESS            VER_SUITE = 0x00000001
	VER_SUITE_ENTERPRISE               VER_SUITE = 0x00000002
	VER_SUITE_BACKOFFICE               VER_SUITE = 0x00000004
	VER_SUITE_COMMUNICATIONS           VER_SUITE = 0x00000008
	VER_SUITE_TERMINAL                 VER_SUITE = 0x00000010
	VER_SUITE_SMALLBUSINESS_RESTRICTED VER_SUITE = 0x00000020
	VER_SUITE_EMBEDDEDNT               VER_SUITE = 0x00000040
	VER_SUITE_DATACENTER               VER_SUITE = 0x00000080
	VER_SUITE_SINGLEUSERTS             VER_SUITE = 0x00000100
	VER_SUITE_PERSONAL                 VER_SUITE = 0x00000200
	VER_SUITE_BLADE                    VER_SUITE = 0x00000400
	VER_SUITE_EMBEDDED_RESTRICTED      VER_SUITE = 0x00000800
	VER_SUITE_SECURITY_APPLIANCE       VER_SUITE = 0x00001000
	VER_SUITE_STORAGE_SERVER           VER_SUITE = 0x00002000
	VER_SUITE_COMPUTE_SERVER           VER_SUITE = 0x00004000
	VER_SUITE_WH_SERVER                VER_SUITE = 0x00008000
)

// VER_PRODUCT
type VER_PRODUCT BYTE

const (
	VER_NT_WORKSTATION       VER_PRODUCT = 0x0000001
	VER_NT_DOMAIN_CONTROLLER VER_PRODUCT = 0x0000002
	VER_NT_SERVER            VER_PRODUCT = 0x0000003
)

// OSVERSIONINFOEX
type OSVERSIONINFOEX struct {
	DwOSVersionInfoSize DWORD
	DwMajorVersion      DWORD
	DwMinorVersion      DWORD
	DwBuildNumber       DWORD
	DwPlatformId        DWORD
	SzCSDVersion        [128]WCHAR
	WServicePackMajor   WORD
	WServicePackMinor   WORD
	WSuiteMask          VER_SUITE
	WProductType        VER_PRODUCT
	WReserved           BYTE
}
type LPOSVERSIONINFOEX *OSVERSIONINFOEX

// RTL_OSVERSIONINFOEXW
type RTL_OSVERSIONINFOEXW struct {
	DwOSVersionInfoSize DWORD
	DwMajorVersion      DWORD
	DwMinorVersion      DWORD
	DwBuildNumber       DWORD
	DwPlatformId        DWORD
	SzCSDVersion        [128]WCHAR
	WServicePackMajor   WORD
	WServicePackMinor   WORD
	WSuiteMask          VER_SUITE
	WProductType        VER_PRODUCT
	WReserved           BYTE
}
type PRTL_OSVERSIONINFOEXW *RTL_OSVERSIONINFOEXW

// SymTagEnum
type SymTagEnum ULONG

const (
	SymTagNull             SymTagEnum = 0
	SymTagExe              SymTagEnum = 1
	SymTagCompiland        SymTagEnum = 2
	SymTagCompilandDetails SymTagEnum = 3
	SymTagCompilandEnv     SymTagEnum = 4
	SymTagFunction         SymTagEnum = 5
	SymTagBlock            SymTagEnum = 6
	SymTagData             SymTagEnum = 7
	SymTagAnnotation       SymTagEnum = 8
	SymTagLabel            SymTagEnum = 9
	SymTagPublicSymbol     SymTagEnum = 10
	SymTagUDT              SymTagEnum = 11
	SymTagEnum1            SymTagEnum = 12
	SymTagFunctionType     SymTagEnum = 13
	SymTagPointerType      SymTagEnum = 14
	SymTagArrayType        SymTagEnum = 15
	SymTagBaseType         SymTagEnum = 16
	SymTagTypedef          SymTagEnum = 17
	SymTagBaseClass        SymTagEnum = 18
	SymTagFriend           SymTagEnum = 19
	SymTagFunctionArgType  SymTagEnum = 20
	SymTagFuncDebugStart   SymTagEnum = 21
	SymTagFuncDebugEnd     SymTagEnum = 22
	SymTagUsingNamespace   SymTagEnum = 23
	SymTagVTableShape      SymTagEnum = 24
	SymTagVTable           SymTagEnum = 25
	SymTagCustom           SymTagEnum = 26
	SymTagThunk            SymTagEnum = 27
	SymTagCustomType       SymTagEnum = 28
	SymTagManagedType      SymTagEnum = 29
	SymTagDimension        SymTagEnum = 30
	SymTagCallSite         SymTagEnum = 31
)

type SymTagEnum_ULONG SymTagEnum

// IMAGE_DATA_DIRECTORY
type IMAGE_DATA_DIRECTORY struct {
	VirtualAddress DWORD
	Size           DWORD
}

// IMAGE_OPTIONAL_MAGIC
type IMAGE_OPTIONAL_MAGIC WORD

const (
	IMAGE_NT_OPTIONAL_HDR32_MAGIC IMAGE_OPTIONAL_MAGIC = 0x10b
	IMAGE_NT_OPTIONAL_HDR64_MAGIC IMAGE_OPTIONAL_MAGIC = 0x20b
	IMAGE_ROM_OPTIONAL_HDR_MAGIC  IMAGE_OPTIONAL_MAGIC = 0x107
)

// IMAGE_SUBSYSTEM
type IMAGE_SUBSYSTEM WORD

const (
	IMAGE_SUBSYSTEM_UNKNOWN                  IMAGE_SUBSYSTEM = 0
	IMAGE_SUBSYSTEM_NATIVE                   IMAGE_SUBSYSTEM = 1
	IMAGE_SUBSYSTEM_WINDOWS_GUI              IMAGE_SUBSYSTEM = 2
	IMAGE_SUBSYSTEM_WINDOWS_CUI              IMAGE_SUBSYSTEM = 3
	IMAGE_SUBSYSTEM_OS2_CUI                  IMAGE_SUBSYSTEM = 5
	IMAGE_SUBSYSTEM_POSIX_CUI                IMAGE_SUBSYSTEM = 7
	IMAGE_SUBSYSTEM_NATIVE_WINDOWS           IMAGE_SUBSYSTEM = 8
	IMAGE_SUBSYSTEM_WINDOWS_CE_GUI           IMAGE_SUBSYSTEM = 9
	IMAGE_SUBSYSTEM_EFI_APPLICATION          IMAGE_SUBSYSTEM = 10
	IMAGE_SUBSYSTEM_EFI_BOOT_SERVICE_DRIVER  IMAGE_SUBSYSTEM = 11
	IMAGE_SUBSYSTEM_EFI_RUNTIME_DRIVER       IMAGE_SUBSYSTEM = 12
	IMAGE_SUBSYSTEM_EFI_ROM                  IMAGE_SUBSYSTEM = 13
	IMAGE_SUBSYSTEM_XBOX                     IMAGE_SUBSYSTEM = 14
	IMAGE_SUBSYSTEM_WINDOWS_BOOT_APPLICATION IMAGE_SUBSYSTEM = 16
)

// IMAGE_SUBSYSTEM_ULONG
type IMAGE_SUBSYSTEM_ULONG ULONG

// IMAGE_DLLCHARACTERISTICS
type IMAGE_DLLCHARACTERISTICS WORD

const (
	IMAGE_LIBRARY_PROCESS_INIT                     IMAGE_DLLCHARACTERISTICS = 0x0001
	IMAGE_LIBRARY_PROCESS_TERM                     IMAGE_DLLCHARACTERISTICS = 0x0002
	IMAGE_LIBRARY_THREAD_INIT                      IMAGE_DLLCHARACTERISTICS = 0x0004
	IMAGE_LIBRARY_THREAD_TERM                      IMAGE_DLLCHARACTERISTICS = 0x0008
	IMAGE_DLLCHARACTERISTICS_HIGH_ENTROPY_VA       IMAGE_DLLCHARACTERISTICS = 0x0020
	IMAGE_DLLCHARACTERISTICS_DYNAMIC_BASE          IMAGE_DLLCHARACTERISTICS = 0x0040
	IMAGE_DLLCHARACTERISTICS_FORCE_INTEGRITY       IMAGE_DLLCHARACTERISTICS = 0x0080
	IMAGE_DLLCHARACTERISTICS_NX_COMPAT             IMAGE_DLLCHARACTERISTICS = 0x0100
	IMAGE_DLLCHARACTERISTICS_NO_ISOLATION          IMAGE_DLLCHARACTERISTICS = 0x0200
	IMAGE_DLLCHARACTERISTICS_NO_SEH                IMAGE_DLLCHARACTERISTICS = 0x0400
	IMAGE_DLLCHARACTERISTICS_NO_BIND               IMAGE_DLLCHARACTERISTICS = 0x0800
	IMAGE_DLLCHARACTERISTICS_APPCONTAINER          IMAGE_DLLCHARACTERISTICS = 0x1000
	IMAGE_DLLCHARACTERISTICS_WDM_DRIVER            IMAGE_DLLCHARACTERISTICS = 0x2000
	IMAGE_DLLCHARACTERISTICS_TERMINAL_SERVER_AWARE IMAGE_DLLCHARACTERISTICS = 0x8000
)

// IMAGE_FILE_MACHINE
type IMAGE_FILE_MACHINE WORD

const (
	IMAGE_FILE_MACHINE_UNKNOWN IMAGE_FILE_MACHINE = 0
	IMAGE_FILE_MACHINE_I386    IMAGE_FILE_MACHINE = 0x014c
	IMAGE_FILE_MACHINE_IA64    IMAGE_FILE_MACHINE = 0x0200
	IMAGE_FILE_MACHINE_AMD64   IMAGE_FILE_MACHINE = 0x8664
)

// IMAGE_FILE_CHARACTERISTICS
type IMAGE_FILE_CHARACTERISTICS WORD

const (
	IMAGE_FILE_RELOCS_STRIPPED         IMAGE_FILE_CHARACTERISTICS = 0x0001
	IMAGE_FILE_EXECUTABLE_IMAGE        IMAGE_FILE_CHARACTERISTICS = 0x0002
	IMAGE_FILE_LINE_NUMS_STRIPPED      IMAGE_FILE_CHARACTERISTICS = 0x0004
	IMAGE_FILE_LOCAL_SYMS_STRIPPED     IMAGE_FILE_CHARACTERISTICS = 0x0008
	IMAGE_FILE_AGGRESIVE_WS_TRIM       IMAGE_FILE_CHARACTERISTICS = 0x0010
	IMAGE_FILE_LARGE_ADDRESS_AWARE     IMAGE_FILE_CHARACTERISTICS = 0x0020
	IMAGE_FILE_BYTES_REVERSED_LO       IMAGE_FILE_CHARACTERISTICS = 0x0080
	IMAGE_FILE_32BIT_MACHINE           IMAGE_FILE_CHARACTERISTICS = 0x0100
	IMAGE_FILE_DEBUG_STRIPPED          IMAGE_FILE_CHARACTERISTICS = 0x0200
	IMAGE_FILE_REMOVABLE_RUN_FROM_SWAP IMAGE_FILE_CHARACTERISTICS = 0x0400
	IMAGE_FILE_NET_RUN_FROM_SWAP       IMAGE_FILE_CHARACTERISTICS = 0x0800
	IMAGE_FILE_SYSTEM                  IMAGE_FILE_CHARACTERISTICS = 0x1000
	IMAGE_FILE_DLL                     IMAGE_FILE_CHARACTERISTICS = 0x2000
	IMAGE_FILE_UP_SYSTEM_ONLY          IMAGE_FILE_CHARACTERISTICS = 0x4000
	IMAGE_FILE_BYTES_REVERSED_HI       IMAGE_FILE_CHARACTERISTICS = 0x8000
)

// IMAGE_FILE_HEADER
type IMAGE_FILE_HEADER struct {
	Machine              IMAGE_FILE_MACHINE
	NumberOfSections     WORD
	TimeDateStamp        DWORD
	PointerToSymbolTable DWORD
	NumberOfSymbols      DWORD
	SizeOfOptionalHeader WORD
	Characteristics      IMAGE_FILE_CHARACTERISTICS
}

// IMAGE_OPTIONAL_HEADER64
type IMAGE_OPTIONAL_HEADER64 struct {
	Magic                       IMAGE_OPTIONAL_MAGIC
	MajorLinkerVersion          BYTE
	MinorLinkerVersion          BYTE
	SizeOfCode                  DWORD
	SizeOfInitializedData       DWORD
	SizeOfUninitializedData     DWORD
	AddressOfEntryPoint         DWORD
	BaseOfCode                  DWORD
	ImageBase                   ULONGLONG
	SectionAlignment            DWORD
	FileAlignment               DWORD
	MajorOperatingSystemVersion WORD
	MinorOperatingSystemVersion WORD
	MajorImageVersion           WORD
	MinorImageVersion           WORD
	MajorSubsystemVersion       WORD
	MinorSubsystemVersion       WORD
	Win32VersionValue           DWORD
	SizeOfImage                 DWORD
	SizeOfHeaders               DWORD
	CheckSum                    DWORD
	Subsystem                   IMAGE_SUBSYSTEM
	DllCharacteristics          IMAGE_DLLCHARACTERISTICS
	SizeOfStackReserve          ULONGLONG
	SizeOfStackCommit           ULONGLONG
	SizeOfHeapReserve           ULONGLONG
	SizeOfHeapCommit            ULONGLONG
	LoaderFlags                 DWORD
	NumberOfRvaAndSizes         DWORD
	DataDirectory               [16]IMAGE_DATA_DIRECTORY
}

type IMAGE_OPTIONAL_HEADER IMAGE_OPTIONAL_HEADER64

// IMAGE_NT_HEADERS
type IMAGE_NT_HEADERS struct {
	Signature      DWORD
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER
}

type PIMAGE_NT_HEADERS *IMAGE_NT_HEADERS

// NUMBERFMT
type NUMBERFMT struct {
	NumDigits     UINT
	LeadingZero   UINT
	Grouping      UINT
	LpDecimalSep  LPWSTR
	LpThousandSep LPWSTR
	NegativeOrder UINT
}

// DwmWindowAttr
type DwmWindowAttr DWORD

const (
	DWMWA_NCRENDERING_ENABLED         DwmWindowAttr = 1
	DWMWA_NCRENDERING_POLICY          DwmWindowAttr = 2
	DWMWA_TRANSITIONS_FORCEDISABLED   DwmWindowAttr = 3
	DWMWA_ALLOW_NCPAINT               DwmWindowAttr = 4
	DWMWA_CAPTION_BUTTON_BOUNDS       DwmWindowAttr = 5
	DWMWA_NONCLIENT_RTL_LAYOUT        DwmWindowAttr = 6
	DWMWA_FORCE_ICONIC_REPRESENTATION DwmWindowAttr = 7
	DWMWA_FLIP3D_POLICY               DwmWindowAttr = 8
	DWMWA_EXTENDED_FRAME_BOUNDS       DwmWindowAttr = 9
	DWMWA_HAS_ICONIC_BITMAP           DwmWindowAttr = 10
	DWMWA_DISALLOW_PEEK               DwmWindowAttr = 11
	DWMWA_EXCLUDED_FROM_PEEK          DwmWindowAttr = 12
	DWMWA_CLOAK                       DwmWindowAttr = 13
	DWMWA_CLOAKED                     DwmWindowAttr = 14
	DWMWA_FREEZE_REPRESENTATION       DwmWindowAttr = 15
)

// FILE_NOTIFY_CHANGE_FLAGS
type FILE_NOTIFY_CHANGE_FLAGS DWORD

const (
	FILE_NOTIFY_CHANGE_FILE_NAME    FILE_NOTIFY_CHANGE_FLAGS = 0x00000001
	FILE_NOTIFY_CHANGE_DIR_NAME     FILE_NOTIFY_CHANGE_FLAGS = 0x00000002
	FILE_NOTIFY_CHANGE_NAME         FILE_NOTIFY_CHANGE_FLAGS = 0x00000003
	FILE_NOTIFY_CHANGE_ATTRIBUTES   FILE_NOTIFY_CHANGE_FLAGS = 0x00000004
	FILE_NOTIFY_CHANGE_SIZE         FILE_NOTIFY_CHANGE_FLAGS = 0x00000008
	FILE_NOTIFY_CHANGE_LAST_WRITE   FILE_NOTIFY_CHANGE_FLAGS = 0x00000010
	FILE_NOTIFY_CHANGE_LAST_ACCESS  FILE_NOTIFY_CHANGE_FLAGS = 0x00000020
	FILE_NOTIFY_CHANGE_CREATION     FILE_NOTIFY_CHANGE_FLAGS = 0x00000040
	FILE_NOTIFY_CHANGE_EA           FILE_NOTIFY_CHANGE_FLAGS = 0x00000080
	FILE_NOTIFY_CHANGE_SECURITY     FILE_NOTIFY_CHANGE_FLAGS = 0x00000100
	FILE_NOTIFY_CHANGE_STREAM_NAME  FILE_NOTIFY_CHANGE_FLAGS = 0x00000200
	FILE_NOTIFY_CHANGE_STREAM_SIZE  FILE_NOTIFY_CHANGE_FLAGS = 0x00000400
	FILE_NOTIFY_CHANGE_STREAM_WRITE FILE_NOTIFY_CHANGE_FLAGS = 0x00000800
)

// SEEK_TYPE
type SEEK_TYPE INT

const (
	SEEK_CUR SEEK_TYPE = 1
	SEEK_END SEEK_TYPE = 2
	SEEK_SET SEEK_TYPE = 0
)

// LocaleMappingFlags
type LocaleMappingFlags DWORD

const (
	LCMAP_LOWERCASE           LocaleMappingFlags = 0x00000100
	LCMAP_UPPERCASE           LocaleMappingFlags = 0x00000200
	LCMAP_SORTKEY             LocaleMappingFlags = 0x00000400
	LCMAP_BYTEREV             LocaleMappingFlags = 0x00000800
	LCMAP_HIRAGANA            LocaleMappingFlags = 0x00100000
	LCMAP_KATAKANA            LocaleMappingFlags = 0x00200000
	LCMAP_HALFWIDTH           LocaleMappingFlags = 0x00400000
	LCMAP_FULLWIDTH           LocaleMappingFlags = 0x00800000
	LCMAP_LINGUISTIC_CASING   LocaleMappingFlags = 0x01000000
	LCMAP_SIMPLIFIED_CHINESE  LocaleMappingFlags = 0x02000000
	LCMAP_TRADITIONAL_CHINESE LocaleMappingFlags = 0x04000000
	NORM_IGNORECASE           LocaleMappingFlags = 0x00000001
	NORM_IGNORENONSPACE       LocaleMappingFlags = 0x00000002
	NORM_IGNORESYMBOLS        LocaleMappingFlags = 0x00000004
	NORM_IGNOREKANATYPE       LocaleMappingFlags = 0x00010000
	NORM_IGNOREWIDTH          LocaleMappingFlags = 0x00020000
	NORM_LINGUISTIC_CASING    LocaleMappingFlags = 0x08000000
	SORT_STRINGSORT           LocaleMappingFlags = 0x00001000
)

// ACTCTX_FLAG
type ACTCTX_FLAG DWORD

const (
	ACTCTX_FLAG_PROCESSOR_ARCHITECTURE_VALID ACTCTX_FLAG = 0x00000001
	ACTCTX_FLAG_LANGID_VALID                 ACTCTX_FLAG = 0x00000002
	ACTCTX_FLAG_ASSEMBLY_DIRECTORY_VALID     ACTCTX_FLAG = 0x00000004
	ACTCTX_FLAG_RESOURCE_NAME_VALID          ACTCTX_FLAG = 0x00000008
	ACTCTX_FLAG_SET_PROCESS_DEFAULT          ACTCTX_FLAG = 0x00000010
	ACTCTX_FLAG_APPLICATION_NAME_VALID       ACTCTX_FLAG = 0x00000020
	ACTCTX_FLAG_SOURCE_IS_ASSEMBLYREF        ACTCTX_FLAG = 0x00000040
	ACTCTX_FLAG_HMODULE_VALID                ACTCTX_FLAG = 0x00000080
)

// ACTCTX
type ACTCTX struct {
	CbSize                 ULONG
	DwFlags                ACTCTX_FLAG
	LpSource               LPCWSTR
	WProcessorArchitecture PROCESSOR_ARCHITECTURE
	WLangId                LANGID
	LpAssemblyDirectory    LPCWSTR
	LpResourceName         LPCWSTR
	LpApplicationName      LPCWSTR
	HModule                HMODULE
}

type PACTCTX *ACTCTX

// HEAP_FLAGS
type HEAP_FLAGS DWORD

const (
	HEAP_NO_SERIALIZE             HEAP_FLAGS = 0x00000001
	HEAP_GROWABLE                 HEAP_FLAGS = 0x00000002
	HEAP_GENERATE_EXCEPTIONS      HEAP_FLAGS = 0x00000004
	HEAP_ZERO_MEMORY              HEAP_FLAGS = 0x00000008
	HEAP_REALLOC_IN_PLACE_ONLY    HEAP_FLAGS = 0x00000010
	HEAP_TAIL_CHECKING_ENABLED    HEAP_FLAGS = 0x00000020
	HEAP_FREE_CHECKING_ENABLED    HEAP_FLAGS = 0x00000040
	HEAP_DISABLE_COALESCE_ON_FREE HEAP_FLAGS = 0x00000080
	HEAP_CREATE_ALIGN_16          HEAP_FLAGS = 0x00010000
	HEAP_CREATE_ENABLE_TRACING    HEAP_FLAGS = 0x00020000
	HEAP_CREATE_ENABLE_EXECUTE    HEAP_FLAGS = 0x00040000
)

// HEAP_FLAGS_ULONG
type HEAP_FLAGS_ULONG HEAP_FLAGS

// IMAGE_FILE_CHARACTERISTICS_ULONG
type IMAGE_FILE_CHARACTERISTICS_ULONG ULONG

// FIND_ACTCTX_SECTION_FLAGS
type FIND_ACTCTX_SECTION_FLAGS DWORD

const (
	FIND_ACTCTX_SECTION_KEY_RETURN_HACTCTX           FIND_ACTCTX_SECTION_FLAGS = 0x00000001
	FIND_ACTCTX_SECTION_KEY_RETURN_FLAGS             FIND_ACTCTX_SECTION_FLAGS = 0x00000002
	FIND_ACTCTX_SECTION_KEY_RETURN_ASSEMBLY_METADATA FIND_ACTCTX_SECTION_FLAGS = 0x00000004
)

// ACTIVATION_CONTEXT_SECTION
type ACTIVATION_CONTEXT_SECTION ULONG

const (
	ACTIVATION_CONTEXT_SECTION_ASSEMBLY_INFORMATION         ACTIVATION_CONTEXT_SECTION = 1
	ACTIVATION_CONTEXT_SECTION_DLL_REDIRECTION              ACTIVATION_CONTEXT_SECTION = 2
	ACTIVATION_CONTEXT_SECTION_WINDOW_CLASS_REDIRECTION     ACTIVATION_CONTEXT_SECTION = 3
	ACTIVATION_CONTEXT_SECTION_COM_SERVER_REDIRECTION       ACTIVATION_CONTEXT_SECTION = 4
	ACTIVATION_CONTEXT_SECTION_COM_INTERFACE_REDIRECTION    ACTIVATION_CONTEXT_SECTION = 5
	ACTIVATION_CONTEXT_SECTION_COM_TYPE_LIBRARY_REDIRECTION ACTIVATION_CONTEXT_SECTION = 6
	ACTIVATION_CONTEXT_SECTION_COM_PROGID_REDIRECTION       ACTIVATION_CONTEXT_SECTION = 7
	ACTIVATION_CONTEXT_SECTION_GLOBAL_OBJECT_RENAME_TABLE   ACTIVATION_CONTEXT_SECTION = 8
	ACTIVATION_CONTEXT_SECTION_CLR_SURROGATES               ACTIVATION_CONTEXT_SECTION = 9
	ACTIVATION_CONTEXT_SECTION_APPLICATION_SETTINGS         ACTIVATION_CONTEXT_SECTION = 10
	ACTIVATION_CONTEXT_SECTION_COMPATIBILITY_INFO           ACTIVATION_CONTEXT_SECTION = 11
)

// WAIT_RESULT
type WAIT_RESULT DWORD

const (
	WAIT_OBJECT_0      WAIT_RESULT = 0x00000000
	WAIT_OBJECT_1      WAIT_RESULT = 0x00000001
	WAIT_OBJECT_2      WAIT_RESULT = 0x00000002
	WAIT_OBJECT_3      WAIT_RESULT = 0x00000003
	WAIT_OBJECT_4      WAIT_RESULT = 0x00000004
	WAIT_OBJECT_5      WAIT_RESULT = 0x00000005
	WAIT_ABANDONED_0   WAIT_RESULT = 0x00000080
	WAIT_ABANDONED_1   WAIT_RESULT = 0x00000081
	WAIT_ABANDONED_2   WAIT_RESULT = 0x00000082
	WAIT_ABANDONED_3   WAIT_RESULT = 0x00000083
	WAIT_ABANDONED_4   WAIT_RESULT = 0x00000084
	WAIT_ABANDONED_5   WAIT_RESULT = 0x00000085
	WAIT_IO_COMPLETION WAIT_RESULT = 0x000000C0
	WAIT_TIMEOUT1      WAIT_RESULT = 258
	WAIT_FAILED        WAIT_RESULT = 0xFFFFFFFF
)

// PROCESSOR_NUMBER
type PROCESSOR_NUMBER struct {
	Group    WORD
	Number   BYTE
	Reserved BYTE
}
type PPROCESSOR_NUMBER *PROCESSOR_NUMBER

// ISOLEVEL
type ISOLEVEL ULONG

const (
	ISOLATIONLEVEL_UNSPECIFIED     ISOLEVEL = 0xffffffff
	ISOLATIONLEVEL_CHAOS           ISOLEVEL = 0x10
	ISOLATIONLEVEL_BROWSE          ISOLEVEL = 0x100
	ISOLATIONLEVEL_CURSORSTABILITY ISOLEVEL = 0x1000
	ISOLATIONLEVEL_REPEATABLEREAD  ISOLEVEL = 0x10000
	ISOLATIONLEVEL_ISOLATED        ISOLEVEL = 0x100000
)

// ISOFLAG
type ISOFLAG ULONG

const (
	ISOFLAG_RETAIN_COMMIT_DC ISOFLAG = 1
	ISOFLAG_RETAIN_COMMIT    ISOFLAG = 2
	ISOFLAG_RETAIN_COMMIT_NO ISOFLAG = 3
	ISOFLAG_RETAIN_ABORT_DC  ISOFLAG = 4
	ISOFLAG_RETAIN_ABORT     ISOFLAG = 8
	ISOFLAG_RETAIN_ABORT_NO  ISOFLAG = 12
	ISOFLAG_RETAIN_DONTCARE  ISOFLAG = 5
	ISOFLAG_RETAIN_BOTH      ISOFLAG = 10
	ISOFLAG_RETAIN_NONE      ISOFLAG = 15
	ISOFLAG_OPTIMISTIC       ISOFLAG = 16
	ISOFLAG_READONLY         ISOFLAG = 32
)

// SHTDN_REASON
type SHTDN_REASON DWORD

const (
	SHTDN_REASON_FLAG_COMMENT_REQUIRED          SHTDN_REASON = 0x01000000
	SHTDN_REASON_FLAG_DIRTY_PROBLEM_ID_REQUIRED SHTDN_REASON = 0x02000000
	SHTDN_REASON_FLAG_CLEAN_UI                  SHTDN_REASON = 0x04000000
	SHTDN_REASON_FLAG_DIRTY_UI                  SHTDN_REASON = 0x08000000
	SHTDN_REASON_FLAG_USER_DEFINED              SHTDN_REASON = 0x40000000
	SHTDN_REASON_FLAG_PLANNED                   SHTDN_REASON = 0x80000000
	SHTDN_REASON_MAJOR_HARDWARE                 SHTDN_REASON = 0x00010000
	SHTDN_REASON_MAJOR_OPERATINGSYSTEM          SHTDN_REASON = 0x00020000
	SHTDN_REASON_MAJOR_SOFTWARE                 SHTDN_REASON = 0x00030000
	SHTDN_REASON_MAJOR_APPLICATION              SHTDN_REASON = 0x00040000
	SHTDN_REASON_MAJOR_SYSTEM                   SHTDN_REASON = 0x00050000
	SHTDN_REASON_MAJOR_POWER                    SHTDN_REASON = 0x00060000
	SHTDN_REASON_MAJOR_LEGACY_API               SHTDN_REASON = 0x00070000
	SHTDN_REASON_MINOR_NONE                     SHTDN_REASON = 0x000000ff
	SHTDN_REASON_MINOR_MAINTENANCE              SHTDN_REASON = 0x00000001
	SHTDN_REASON_MINOR_INSTALLATION             SHTDN_REASON = 0x00000002
	SHTDN_REASON_MINOR_UPGRADE                  SHTDN_REASON = 0x00000003
	SHTDN_REASON_MINOR_RECONFIG                 SHTDN_REASON = 0x00000004
	SHTDN_REASON_MINOR_HUNG                     SHTDN_REASON = 0x00000005
	SHTDN_REASON_MINOR_UNSTABLE                 SHTDN_REASON = 0x00000006
	SHTDN_REASON_MINOR_DISK                     SHTDN_REASON = 0x00000007
	SHTDN_REASON_MINOR_PROCESSOR                SHTDN_REASON = 0x00000008
	SHTDN_REASON_MINOR_NETWORKCARD              SHTDN_REASON = 0x00000009
	SHTDN_REASON_MINOR_POWER_SUPPLY             SHTDN_REASON = 0x0000000a
	SHTDN_REASON_MINOR_CORDUNPLUGGED            SHTDN_REASON = 0x0000000b
	SHTDN_REASON_MINOR_ENVIRONMENT              SHTDN_REASON = 0x0000000c
	SHTDN_REASON_MINOR_HARDWARE_DRIVER          SHTDN_REASON = 0x0000000d
	SHTDN_REASON_MINOR_OTHERDRIVER              SHTDN_REASON = 0x0000000e
	SHTDN_REASON_MINOR_BLUESCREEN               SHTDN_REASON = 0x0000000F
	SHTDN_REASON_MINOR_SERVICEPACK              SHTDN_REASON = 0x00000010
	SHTDN_REASON_MINOR_HOTFIX                   SHTDN_REASON = 0x00000011
	SHTDN_REASON_MINOR_SECURITYFIX              SHTDN_REASON = 0x00000012
	SHTDN_REASON_MINOR_SECURITY                 SHTDN_REASON = 0x00000013
	SHTDN_REASON_MINOR_NETWORK_CONNECTIVITY     SHTDN_REASON = 0x00000014
	SHTDN_REASON_MINOR_WMI                      SHTDN_REASON = 0x00000015
	SHTDN_REASON_MINOR_SERVICEPACK_UNINSTALL    SHTDN_REASON = 0x00000016
	SHTDN_REASON_MINOR_HOTFIX_UNINSTALL         SHTDN_REASON = 0x00000017
	SHTDN_REASON_MINOR_SECURITYFIX_UNINSTALL    SHTDN_REASON = 0x00000018
	SHTDN_REASON_MINOR_MMC                      SHTDN_REASON = 0x00000019
	SHTDN_REASON_MINOR_SYSTEMRESTORE            SHTDN_REASON = 0x0000001a
	SHTDN_REASON_MINOR_TERMSRV                  SHTDN_REASON = 0x00000020
	SHTDN_REASON_MINOR_DC_PROMOTION             SHTDN_REASON = 0x00000021
	SHTDN_REASON_MINOR_DC_DEMOTION              SHTDN_REASON = 0x00000022
	SHTDN_REASON_LEGACY_API                     SHTDN_REASON = 0x80070000
)

// EWX_FLAGS
type EWX_FLAGS UINT

const (
	EWX_LOGOFF          EWX_FLAGS = 0x00000000
	EWX_SHUTDOWN        EWX_FLAGS = 0x00000001
	EWX_REBOOT          EWX_FLAGS = 0x00000002
	EWX_FORCE           EWX_FLAGS = 0x00000004
	EWX_POWEROFF        EWX_FLAGS = 0x00000008
	EWX_FORCEIFHUNG     EWX_FLAGS = 0x00000010
	EWX_QUICKRESOLVE    EWX_FLAGS = 0x00000020
	EWX_RESTARTAPPS     EWX_FLAGS = 0x00000040
	EWX_HYBRID_SHUTDOWN EWX_FLAGS = 0x00400000
	EWX_BOOTOPTIONS     EWX_FLAGS = 0x01000000
)

// URLZONE
type URLZONE DWORD

const (
	URLZONE_LOCAL_MACHINE URLZONE = 0
	URLZONE_INTRANET      URLZONE = 1
	URLZONE_TRUSTED       URLZONE = 2
	URLZONE_INTERNET      URLZONE = 3
	URLZONE_UNTRUSTED     URLZONE = 4
)

// THREAD_PRIORITY
type THREAD_PRIORITY INT

const (
	THREAD_PRIORITY_LOWEST        THREAD_PRIORITY = -2
	THREAD_PRIORITY_BELOW_NORMAL  THREAD_PRIORITY = -1
	THREAD_PRIORITY_NORMAL        THREAD_PRIORITY = 0
	THREAD_PRIORITY_HIGHEST       THREAD_PRIORITY = 2
	THREAD_PRIORITY_ABOVE_NORMAL  THREAD_PRIORITY = 1
	THREAD_PRIORITY_TIME_CRITICAL THREAD_PRIORITY = 15
	THREAD_PRIORITY_IDLE          THREAD_PRIORITY = -15
	THREAD_PRIORITY_ERROR_RETURN  THREAD_PRIORITY = 0x7FFFFFFF
)

// MoveMethodEnum
type MoveMethodEnum DWORD

const (
	FILE_BEGIN   MoveMethodEnum = 0
	FILE_CURRENT MoveMethodEnum = 1
	FILE_END     MoveMethodEnum = 2
)

// XSTATE_FEATURE_64
type XSTATE_FEATURE_64 DWORD64

const (
	XSTATE_LEGACY_FLOATING_POINT XSTATE_FEATURE_64 = 0
	XSTATE_LEGACY_SSE            XSTATE_FEATURE_64 = 1
	XSTATE_AVX                   XSTATE_FEATURE_64 = 2
)

// XSTATE_MASK
type XSTATE_MASK DWORD64

const (
	XSTATE_MASK_LEGACY_FLOATING_POINT XSTATE_MASK = 0x1
	XSTATE_MASK_LEGACY_SSE            XSTATE_MASK = 0x2
	XSTATE_MASK_LEGACY                XSTATE_MASK = 0x3
	XSTATE_MASK_AVX                   XSTATE_MASK = 0x4
)
