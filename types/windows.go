package types

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
