package types

// Variables
type (
	HICON           HANDLE
	HMENU           HANDLE
	HDC             HANDLE
	HPALETTE        HANDLE
	HCURSOR         HICON
	HMONITOR        HANDLE
	HIMAGELIST      LPVOID
	HRGN            HANDLE
	HMETAFILE       HANDLE
	HBRUSH          HANDLE
	HFONT           HANDLE
	LPCCHOOKPROC    LPVOID
	LPCFHOOKPROC    LPVOID
	LPPAGESETUPHOOK LPVOID
	LPPAGEPAINTHOOK LPVOID
	LPFRHOOKPROC    LPVOID
	LPPRINTHOOKPROC LPVOID
	LPSETUPHOOKPROC LPVOID
	REFERENCE_TIME  LONGLONG
	HPROFILE        HANDLE
	PHPROFILE       *HPROFILE
	FXPT2DOT30      LONG
	PBMCALLBACKFN   LPVOID
	LPBMCALLBACKFN  PBMCALLBACKFN
	DLGPROC         LPVOID
)

// RGBQUAD
type RGBQUAD struct {
	Blue     BYTE
	Green    BYTE
	Red      BYTE
	Reserved BYTE
}

// RegionType
type RegionType DWORD

const (
	RDH_RECTANGLES RegionType = 1
)

// RGNDATAHEADER
type RGNDATAHEADER struct {
	Size    DWORD
	Type    RegionType
	Count   DWORD
	RgnSize DWORD
	Bound   RECT
}

// RGNDATA
type RGNDATA struct {
	Rdh    RGNDATAHEADER
	Buffer CHAR
}
type LPRGNDATA *RGNDATA

// COLORREF
type COLORREF DWORD

const (
	CLR_INVALID COLORREF = 0xFFFFFFFF
)

type LPCOLORREF *COLORREF

// SIZE
type SIZE struct {
	X LONG
	Y LONG
}

type (
	LPSIZE *SIZE
	// SIZEL
	SIZEL   SIZE
	LPSIZEL *SIZEL
)

// ChooseColorFlags
type ChooseColorFlags DWORD

const (
	CC_RGBINIT              ChooseColorFlags = 0x00000001
	CC_FULLOPEN             ChooseColorFlags = 0x00000002
	CC_PREVENTFULLOPEN      ChooseColorFlags = 0x00000004
	CC_SHOWHELP             ChooseColorFlags = 0x00000008
	CC_ENABLEHOOK           ChooseColorFlags = 0x00000010
	CC_ENABLETEMPLATE       ChooseColorFlags = 0x00000020
	CC_ENABLETEMPLATEHANDLE ChooseColorFlags = 0x00000040
	CC_SOLIDCOLOR           ChooseColorFlags = 0x00000080
	CC_ANYCOLOR             ChooseColorFlags = 0x00000100
)

// CHOOSECOLOR
type CHOOSECOLOR struct {
	StructSize   DWORD
	Owner        HWND
	Instance     HWND
	Result       COLORREF
	CustColors   *COLORREF
	Flags        ChooseColorFlags
	CustData     LPARAM
	Hook         LPCCHOOKPROC
	TemplateName LPCWSTR
}

// ChoooseFontFlags
type ChoooseFontFlags DWORD

const (
	CF_SCREENFONTS          ChoooseFontFlags = 0x00000001
	CF_PRINTERFONTS         ChoooseFontFlags = 0x00000002
	CF_BOTH                 ChoooseFontFlags = 0x00000003
	CF_SHOWHELP             ChoooseFontFlags = 0x00000004
	CF_ENABLEHOOK           ChoooseFontFlags = 0x00000008
	CF_ENABLETEMPLATE       ChoooseFontFlags = 0x00000010
	CF_ENABLETEMPLATEHANDLE ChoooseFontFlags = 0x00000020
	CF_INITTOLOGFONTSTRUCT  ChoooseFontFlags = 0x00000040
	CF_USESTYLE             ChoooseFontFlags = 0x00000080
	CF_EFFECTS              ChoooseFontFlags = 0x00000100
	CF_APPLY                ChoooseFontFlags = 0x00000200
	CF_ANSIONLY             ChoooseFontFlags = 0x00000400
	CF_NOVECTORFONTS        ChoooseFontFlags = 0x00000800
	CF_NOSIMULATIONS        ChoooseFontFlags = 0x00001000
	CF_LIMITSIZE            ChoooseFontFlags = 0x00002000
	CF_FIXEDPITCHONLY       ChoooseFontFlags = 0x00004000
	CF_WYSIWYG              ChoooseFontFlags = 0x00008000
	CF_FORCEFONTEXIST       ChoooseFontFlags = 0x00010000
	CF_SCALABLEONLY         ChoooseFontFlags = 0x00020000
	CF_TTONLY               ChoooseFontFlags = 0x00040000
	CF_NOFACESEL            ChoooseFontFlags = 0x00080000
	CF_NOSTYLESEL           ChoooseFontFlags = 0x00100000
	CF_NOSIZESEL            ChoooseFontFlags = 0x00200000
	CF_SELECTSCRIPT         ChoooseFontFlags = 0x00400000
	CF_NOSCRIPTSEL          ChoooseFontFlags = 0x00800000
	CF_NOVERTFONTS          ChoooseFontFlags = 0x01000000
	CF_INACTIVEFONTS        ChoooseFontFlags = 0x02000000
)

// FontType
type FontType WORD

const (
	SIMULATED_FONTTYPE FontType = 0x8000
	PRINTER_FONTTYPE   FontType = 0x4000
	SCREEN_FONTTYPE    FontType = 0x2000
	BOLD_FONTTYPE      FontType = 0x0100
	ITALIC_FONTTYPE    FontType = 0x0200
	REGULAR_FONTTYPE   FontType = 0x0400
)

// FontWeight
type FontWeight LONG

const (
	FW_DONTCARE   FontWeight = 0
	FW_THIN       FontWeight = 100
	FW_EXTRALIGHT FontWeight = 200
	FW_LIGHT      FontWeight = 300
	FW_NORMAL     FontWeight = 400
	FW_MEDIUM     FontWeight = 500
	FW_SEMIBOLD   FontWeight = 600
	FW_BOLD       FontWeight = 700
	FW_EXTRABOLD  FontWeight = 800
	FW_HEAVY      FontWeight = 900
)

// FontCharset
type FontCharset BYTE

const (
	ANSI_CHARSET        FontCharset = 0
	DEFAULT_CHARSET     FontCharset = 1
	SYMBOL_CHARSET      FontCharset = 2
	SHIFTJIS_CHARSET    FontCharset = 128
	HANGEUL_CHARSET     FontCharset = 129
	HANGUL_CHARSET      FontCharset = 129
	GB2312_CHARSET      FontCharset = 134
	CHINESEBIG5_CHARSET FontCharset = 136
	OEM_CHARSET         FontCharset = 255
	JOHAB_CHARSET       FontCharset = 130
	HEBREW_CHARSET      FontCharset = 177
	ARABIC_CHARSET      FontCharset = 178
	GREEK_CHARSET       FontCharset = 161
	TURKISH_CHARSET     FontCharset = 162
	VIETNAMESE_CHARSET  FontCharset = 163
	THAI_CHARSET        FontCharset = 222
	EASTEUROPE_CHARSET  FontCharset = 238
	RUSSIAN_CHARSET     FontCharset = 204
	MAC_CHARSET         FontCharset = 77
	BALTIC_CHARSET      FontCharset = 186
)

// FontOutputPrecision
type FontOutputPrecision BYTE

const (
	OUT_DEFAULT_PRECIS        FontOutputPrecision = 0
	OUT_STRING_PRECIS         FontOutputPrecision = 1
	OUT_CHARACTER_PRECIS      FontOutputPrecision = 2
	OUT_STROKE_PRECIS         FontOutputPrecision = 3
	OUT_TT_PRECIS             FontOutputPrecision = 4
	OUT_DEVICE_PRECIS         FontOutputPrecision = 5
	OUT_RASTER_PRECIS         FontOutputPrecision = 6
	OUT_TT_ONLY_PRECIS        FontOutputPrecision = 7
	OUT_OUTLINE_PRECIS        FontOutputPrecision = 8
	OUT_SCREEN_OUTLINE_PRECIS FontOutputPrecision = 9
	OUT_PS_ONLY_PRECIS        FontOutputPrecision = 10
)

// FontClipPrecision
type FontClipPrecision BYTE

const (
	CLIP_DEFAULT_PRECIS   FontClipPrecision = 0
	CLIP_CHARACTER_PRECIS FontClipPrecision = 1
	CLIP_STROKE_PRECIS    FontClipPrecision = 2
	CLIP_MASK             FontClipPrecision = 0xf
	CLIP_LH_ANGLES        FontClipPrecision = 0x10
	CLIP_TT_ALWAYS        FontClipPrecision = 0x20
	CLIP_DFA_DISABLE      FontClipPrecision = 0x40
	CLIP_EMBEDDED         FontClipPrecision = 0x80
)

// FontQuality
type FontQuality BYTE

const (
	DEFAULT_QUALITY           FontQuality = 0
	DRAFT_QUALITY             FontQuality = 1
	PROOF_QUALITY             FontQuality = 2
	NONANTIALIASED_QUALITY    FontQuality = 3
	ANTIALIASED_QUALITY       FontQuality = 4
	CLEARTYPE_QUALITY         FontQuality = 5
	CLEARTYPE_NATURAL_QUALITY FontQuality = 6
)

// FontPitchAndFamily
type FontPitchAndFamily BYTE

const (
	DEFAULT_PITCH  FontPitchAndFamily = 0
	FIXED_PITCH    FontPitchAndFamily = 1
	VARIABLE_PITCH FontPitchAndFamily = 2
	FF_DONTCARE    FontPitchAndFamily = 0
	FF_ROMAN       FontPitchAndFamily = 0x10
	FF_SWISS       FontPitchAndFamily = 0x20
	FF_MODERN      FontPitchAndFamily = 0x30
	FF_SCRIPT      FontPitchAndFamily = 0x40
	FF_DECORATIVE  FontPitchAndFamily = 0x50
)

// LOGFONT
type LOGFONT struct {
	Height         LONG
	Width          LONG
	Escapement     LONG
	Orientation    LONG
	Weight         FontWeight
	Italic         BYTE
	Underline      BYTE
	StrikeOut      BYTE
	CharSet        FontCharset
	OutPrecision   FontOutputPrecision
	ClipPrecision  FontClipPrecision
	Quality        FontQuality
	PitchAndFamily FontPitchAndFamily
	FaceName       [32]WCHAR
}
type LPLOGFONT *LOGFONT

// CHOOSEFONT
type CHOOSEFONT struct {
	StructSize       DWORD
	Owner            HWND
	DC               HDC
	LogFont          LPLOGFONT
	PointSize        INT
	Flags            ChoooseFontFlags
	Colors           COLORREF
	CustData         LPARAM
	Hook             LPCFHOOKPROC
	TemplateName     LPCWSTR
	Instance         HINSTANCE
	Style            LPWSTR
	FontType         FontType
	MISSINGALIGNMENT WORD
	SizeMin          INT
	SizeMax          INT
}

type LPCHOOSEFONT *CHOOSEFONT

// FindReplaceFlags
type FindReplaceFlags DWORD

const (
	FR_DOWN                 FindReplaceFlags = 0x00000001
	FR_WHOLEWORD            FindReplaceFlags = 0x00000002
	FR_MATCHCASE            FindReplaceFlags = 0x00000004
	FR_FINDNEXT             FindReplaceFlags = 0x00000008
	FR_REPLACE              FindReplaceFlags = 0x00000010
	FR_REPLACEALL           FindReplaceFlags = 0x00000020
	FR_DIALOGTERM           FindReplaceFlags = 0x00000040
	FR_SHOWHELP             FindReplaceFlags = 0x00000080
	FR_ENABLEHOOK           FindReplaceFlags = 0x00000100
	FR_ENABLETEMPLATE       FindReplaceFlags = 0x00000200
	FR_NOUPDOWN             FindReplaceFlags = 0x00000400
	FR_NOMATCHCASE          FindReplaceFlags = 0x00000800
	FR_NOWHOLEWORD          FindReplaceFlags = 0x00001000
	FR_ENABLETEMPLATEHANDLE FindReplaceFlags = 0x00002000
	FR_HIDEUPDOWN           FindReplaceFlags = 0x00004000
	FR_HIDEMATCHCASE        FindReplaceFlags = 0x00008000
	FR_HIDEWHOLEWORD        FindReplaceFlags = 0x00010000
	FR_RAW                  FindReplaceFlags = 0x00020000
	FR_MATCHDIAC            FindReplaceFlags = 0x20000000
	FR_MATCHKASHIDA         FindReplaceFlags = 0x40000000
	FR_MATCHALEFHAMZA       FindReplaceFlags = 0x80000000
)

// FINDREPLACE
type FINDREPLACE struct {
	StructSize     DWORD
	Owner          HWND
	Instance       HINSTANCE
	Flags          FindReplaceFlags
	FindWhat       LPWSTR
	ReplaceWith    LPWSTR
	FindWhatLen    WORD
	ReplaceWithLen WORD
	CustData       LPARAM
	Hook           LPFRHOOKPROC
	TemplateName   LPCWSTR
}

type LPFINDREPLACE *FINDREPLACE

// OfnFlags
type OfnFlags DWORD

const (
	OFN_READONLY             OfnFlags = 0x00000001
	OFN_OVERWRITEPROMPT      OfnFlags = 0x00000002
	OFN_HIDEREADONLY         OfnFlags = 0x00000004
	OFN_NOCHANGEDIR          OfnFlags = 0x00000008
	OFN_SHOWHELP             OfnFlags = 0x00000010
	OFN_ENABLEHOOK           OfnFlags = 0x00000020
	OFN_ENABLETEMPLATE       OfnFlags = 0x00000040
	OFN_ENABLETEMPLATEHANDLE OfnFlags = 0x00000080
	OFN_NOVALIDATE           OfnFlags = 0x00000100
	OFN_ALLOWMULTISELECT     OfnFlags = 0x00000200
	OFN_EXTENSIONDIFFERENT   OfnFlags = 0x00000400
	OFN_PATHMUSTEXIST        OfnFlags = 0x00000800
	OFN_FILEMUSTEXIST        OfnFlags = 0x00001000
	OFN_CREATEPROMPT         OfnFlags = 0x00002000
	OFN_SHAREAWARE           OfnFlags = 0x00004000
	OFN_NOREADONLYRETURN     OfnFlags = 0x00008000
	OFN_NOTESTFILECREATE     OfnFlags = 0x00010000
	OFN_NONETWORKBUTTON      OfnFlags = 0x00020000
	OFN_NOLONGNAMES          OfnFlags = 0x00040000
	OFN_EXPLORER             OfnFlags = 0x00080000
	OFN_NODEREFERENCELINKS   OfnFlags = 0x00100000
	OFN_LONGNAMES            OfnFlags = 0x00200000
	OFN_ENABLEINCLUDENOTIFY  OfnFlags = 0x00400000
	OFN_ENABLESIZING         OfnFlags = 0x00800000
	OFN_DONTADDTORECENT      OfnFlags = 0x02000000
	OFN_FORCESHOWHIDDEN      OfnFlags = 0x10000000
)

// OfnFlagsEx
type OfnFlagsEx DWORD

const (
	OFN_EX_NOPLACESBAR OfnFlagsEx = 0x00000001
)

type OPENFILENAME struct {
	StructSize    DWORD
	Owner         HWND
	Instance      HINSTANCE
	Filter        LPCWSTR
	CustomFilter  LPWSTR
	MaxCustFilter DWORD
	FilterIndex   DWORD
	File          LPWSTR
	MaxFile       DWORD
	FileTitle     LPWSTR
	MaxFileTitle  DWORD
	InitialDir    LPCWSTR
	Title         LPCWSTR
	Flags         OfnFlags
	FileOffset    WORD
	FileExtension WORD
	DefExt        LPCWSTR
	CustData      LPARAM
	Hook          LPVOID
	TemplateName  LPCWSTR
	PvReserved    LPVOID
	DwReserved    DWORD
	FlagsEx       OfnFlagsEx
}
type LPOPENFILENAME *OPENFILENAME

// PageSetupDialogFlags
type PageSetupDialogFlags DWORD

const (
	PSD_DEFAULTMINMARGINS             PageSetupDialogFlags = 0x00000000
	PSD_INWININIINTLMEASURE           PageSetupDialogFlags = 0x00000000
	PSD_MINMARGINS                    PageSetupDialogFlags = 0x00000001
	PSD_MARGINS                       PageSetupDialogFlags = 0x00000002
	PSD_INTHOUSANDTHSOFINCHES         PageSetupDialogFlags = 0x00000004
	PSD_INHUNDREDTHSOFMILLIMETERS     PageSetupDialogFlags = 0x00000008
	PSD_DISABLEMARGINS                PageSetupDialogFlags = 0x00000010
	PSD_DISABLEPRINTER                PageSetupDialogFlags = 0x00000020
	PSD_NOWARNING                     PageSetupDialogFlags = 0x00000080
	PSD_DISABLEORIENTATION            PageSetupDialogFlags = 0x00000100
	PSD_RETURNDEFAULT                 PageSetupDialogFlags = 0x00000400
	PSD_DISABLEPAPER                  PageSetupDialogFlags = 0x00000200
	PSD_SHOWHELP                      PageSetupDialogFlags = 0x00000800
	PSD_ENABLEPAGESETUPHOOK           PageSetupDialogFlags = 0x00002000
	PSD_ENABLEPAGESETUPTEMPLATE       PageSetupDialogFlags = 0x00008000
	PSD_ENABLEPAGESETUPTEMPLATEHANDLE PageSetupDialogFlags = 0x00020000
	PSD_ENABLEPAGEPAINTHOOK           PageSetupDialogFlags = 0x00040000
	PSD_DISABLEPAGEPAINTING           PageSetupDialogFlags = 0x00080000
	PSD_NONETWORKBUTTON               PageSetupDialogFlags = 0x00200000
)

// PAGESETUPDLG
type PAGESETUPDLG struct {
	StructSize            DWORD
	Owner                 HWND
	DevMode               HGLOBAL
	DevNames              HGLOBAL
	Flags                 PageSetupDialogFlags
	PaperSize             POINT
	MinMargin             RECT
	Margin                RECT
	Instance              HINSTANCE
	CustData              LPARAM
	PageSetupHook         LPPAGESETUPHOOK
	PagePaintHook         LPPAGEPAINTHOOK
	PageSetupTemplateName LPCWSTR
	PageSetupTemplate     HGLOBAL
}

type LPPAGESETUPDLG *PAGESETUPDLG

// PrintDlgFlags
type PrintDlgFlags DWORD

const (
	PD_ALLPAGES                   PrintDlgFlags = 0x00000000
	PD_SELECTION                  PrintDlgFlags = 0x00000001
	PD_PAGENUMS                   PrintDlgFlags = 0x00000002
	PD_NOSELECTION                PrintDlgFlags = 0x00000004
	PD_NOPAGENUMS                 PrintDlgFlags = 0x00000008
	PD_COLLATE                    PrintDlgFlags = 0x00000010
	PD_PRINTTOFILE                PrintDlgFlags = 0x00000020
	PD_PRINTSETUP                 PrintDlgFlags = 0x00000040
	PD_NOWARNING                  PrintDlgFlags = 0x00000080
	PD_RETURNDC                   PrintDlgFlags = 0x00000100
	PD_RETURNIC                   PrintDlgFlags = 0x00000200
	PD_RETURNDEFAULT              PrintDlgFlags = 0x00000400
	PD_SHOWHELP                   PrintDlgFlags = 0x00000800
	PD_ENABLEPRINTHOOK            PrintDlgFlags = 0x00001000
	PD_ENABLESETUPHOOK            PrintDlgFlags = 0x00002000
	PD_ENABLEPRINTTEMPLATE        PrintDlgFlags = 0x00004000
	PD_ENABLESETUPTEMPLATE        PrintDlgFlags = 0x00008000
	PD_ENABLEPRINTTEMPLATEHANDLE  PrintDlgFlags = 0x00010000
	PD_ENABLESETUPTEMPLATEHANDLE  PrintDlgFlags = 0x00020000
	PD_USEDEVMODECOPIES           PrintDlgFlags = 0x00040000
	PD_USEDEVMODECOPIESANDCOLLATE PrintDlgFlags = 0x00040000
	PD_DISABLEPRINTTOFILE         PrintDlgFlags = 0x00080000
	PD_HIDEPRINTTOFILE            PrintDlgFlags = 0x00100000
	PD_NONETWORKBUTTON            PrintDlgFlags = 0x00200000
	PD_CURRENTPAGE                PrintDlgFlags = 0x00400000
	PD_NOCURRENTPAGE              PrintDlgFlags = 0x00800000
	PD_EXCLUSIONFLAGS             PrintDlgFlags = 0x01000000
	PD_USELARGETEMPLATE           PrintDlgFlags = 0x10000000
)

// PRINTDLG
type PRINTDLG struct {
	StructSize        DWORD
	Owner             HWND
	DevMode           HGLOBAL
	DevNames          HGLOBAL
	DC                HDC
	Flags             PrintDlgFlags
	FromPage          WORD
	ToPage            WORD
	MinPage           WORD
	MaxPage           WORD
	Copies            WORD
	Instance          HINSTANCE
	CustData          LPARAM
	PrintHook         LPPRINTHOOKPROC
	SetupHook         LPSETUPHOOKPROC
	PrintTemplateName LPCWSTR
	SetupTemplateName LPCWSTR
	PrintTemplate     HGLOBAL
	SetupTemplate     HGLOBAL
}

type LPPRINTDLG *PRINTDLG

// PRINTPAGERANGE
type PRINTPAGERANGE struct {
	FromPage DWORD
	ToPage   DWORD
}

type LPPRINTPAGERANGE *PRINTPAGERANGE

// PRINTDLGEX
type PRINTDLGEX struct {
	StructSize        DWORD
	Owner             HWND
	DevMode           HGLOBAL
	DevNames          HGLOBAL
	DC                HDC
	Flags             PrintDlgFlags
	Flags2            DWORD
	ExclusionFlags    DWORD
	PageRanges        DWORD
	MaxPageRanges     DWORD
	LpPageRanges      LPPRINTPAGERANGE
	MinPage           DWORD
	MaxPage           DWORD
	Copies            DWORD
	Instance          HINSTANCE
	PrintTemplateName LPCWSTR
	Callback          uintptr
	PropertyPages     DWORD
	LphPropertyPages  *HPROPSHEETPAGE
	StartPage         DWORD
	ResultAction      DWORD
}

type LPPRINTDLGEX *PRINTDLGEX

// DRAWTEXTPARAMS
type DRAWTEXTPARAMS struct {
	Size        UINT
	TabLength   int
	LeftMargin  int
	RightMargin int
	LengthDrawn UINT
}

type LPDRAWTEXTPARAMS *DRAWTEXTPARAMS

// TRACKMOUSEEVENT_Flags
type TRACKMOUSEEVENT_Flags DWORD

const (
	TME_HOVER     TRACKMOUSEEVENT_Flags = 0x00000001
	TME_LEAVE     TRACKMOUSEEVENT_Flags = 0x00000002
	TME_NONCLIENT TRACKMOUSEEVENT_Flags = 0x00000010
	TME_QUERY     TRACKMOUSEEVENT_Flags = 0x40000000
	TME_CANCEL    TRACKMOUSEEVENT_Flags = 0x80000000
)

// TRACKMOUSEEVENT
type TRACKMOUSEEVENT struct {
	Size      DWORD
	Flags     TRACKMOUSEEVENT_Flags
	Track     HWND
	HoverTime DWORD
}

type LPTRACKMOUSEEVENT *TRACKMOUSEEVENT

// SCROLLINFO_Flags
type SCROLLINFO_Flags UINT

const (
	SIF_RANGE           SCROLLINFO_Flags = 0x0001
	SIF_PAGE            SCROLLINFO_Flags = 0x0002
	SIF_POS             SCROLLINFO_Flags = 0x0004
	SIF_DISABLENOSCROLL SCROLLINFO_Flags = 0x0008
	SIF_TRACKPOS        SCROLLINFO_Flags = 0x0010
	SIF_ALL             SCROLLINFO_Flags = 0x0017
)

// SCROLLINFO
type SCROLLINFO struct {
	Size     UINT
	Mask     SCROLLINFO_Flags
	Min      int
	Max      int
	Page     UINT
	Pos      int
	TrackPos int
}

type (
	LPSCROLLINFO  *SCROLLINFO
	LPCSCROLLINFO *SCROLLINFO
)

// MARGINS
type MARGINS struct {
	LeftWidth    int
	RightWidth   int
	TopHeight    int
	BottomHeight int
}

// BLENDFUNCTION
type BLENDFUNCTION struct {
	BlendOp             BYTE
	BlendFlags          BYTE
	SourceConstantAlpha BYTE
	AlphaFormat         BYTE
}

// BITMAPINFOHEADER
type BITMAPINFOHEADER struct {
	Size          DWORD
	Width         LONG
	Height        LONG
	Planes        WORD
	BitCount      WORD
	Compression   DWORD
	SizeImage     DWORD
	XPelsPerMeter LONG
	YPelsPerMeter LONG
	ClrUsed       DWORD
	ClrImportant  DWORD
}

type LPBITMAPINFOHEADER *BITMAPINFOHEADER

// BITMAPINFO
type BITMAPINFO struct {
	Header BITMAPINFOHEADER
	Colors [1]RGBQUAD
}

type LPBITMAPINFO *BITMAPINFO

// TEXTMETRIC_Pitch
type TEXTMETRIC_Pitch BYTE

const (
	TMPF_FIXED_PITCH TEXTMETRIC_Pitch = 0x01
	TMPF_VECTOR      TEXTMETRIC_Pitch = 0x02
	TMPF_TRUETYPE    TEXTMETRIC_Pitch = 0x04
	TMPF_DEVICE      TEXTMETRIC_Pitch = 0x08
	TMPF_DONTCARE    TEXTMETRIC_Pitch = 0
	TMPF_ROMAN       TEXTMETRIC_Pitch = 0x10
	TMPF_SWISS       TEXTMETRIC_Pitch = 0x20
	TMPF_MODERN      TEXTMETRIC_Pitch = 0x30
	TMPF_SCRIPT      TEXTMETRIC_Pitch = 0x40
	TMPF_DECORATIVE  TEXTMETRIC_Pitch = 0x50
)

// TEXTMETRIC
type TEXTMETRIC struct {
	Height           LONG
	Ascent           LONG
	Descent          LONG
	InternalLeading  LONG
	ExternalLeading  LONG
	AveCharWidth     LONG
	MaxCharWidth     LONG
	Weight           LONG
	Overhang         LONG
	DigitizedAspectX LONG
	DigitizedAspectY LONG
	FirstChar        WCHAR
	LastChar         WCHAR
	DefaultChar      WCHAR
	BreakChar        WCHAR
	Italic           BYTE
	Underlined       BYTE
	StruckOut        BYTE
	PitchAndFamily   TEXTMETRIC_Pitch
	CharSet          FontCharset
}

type LPTEXTMETRIC *TEXTMETRIC

// MONITORINFO_Flags
type MONITORINFO_Flags DWORD

const (
	MONITORINFOF_PRIMARY MONITORINFO_Flags = 0x00000001
)

// VIDEOINFOHEADER
type VIDEOINFOHEADER struct {
	Source          RECT
	Target          RECT
	BitRate         DWORD
	BitErrorRate    DWORD
	AvgTimePerFrame REFERENCE_TIME
	Header          BITMAPINFOHEADER
}

// AMINTERLACE_FLAGS
type AMINTERLACE_FLAGS DWORD

const (
	AMINTERLACE_IsInterlaced          AMINTERLACE_FLAGS = 0x00000001
	AMINTERLACE_1FieldPerSample       AMINTERLACE_FLAGS = 0x00000002
	AMINTERLACE_Field1First           AMINTERLACE_FLAGS = 0x00000004
	AMINTERLACE_FieldPatField1Only    AMINTERLACE_FLAGS = 0x00000000
	AMINTERLACE_FieldPatField2Only    AMINTERLACE_FLAGS = 0x00000010
	AMINTERLACE_FieldPatBothRegular   AMINTERLACE_FLAGS = 0x00000020
	AMINTERLACE_FieldPatBothIrregular AMINTERLACE_FLAGS = 0x00000030
	AMINTERLACE_DisplayModeBobOnly    AMINTERLACE_FLAGS = 0x00000000
	AMINTERLACE_DisplayModeWeaveOnly  AMINTERLACE_FLAGS = 0x00000040
	AMINTERLACE_DisplayModeBobOrWeave AMINTERLACE_FLAGS = 0x00000080
)

// AMCOPYPROTECT_FLAGS
type AMCOPYPROTECT_FLAGS DWORD

const (
	AMCOPYPROTECT_RestrictDuplication AMCOPYPROTECT_FLAGS = 0x00000001
)

// AMCONTROL_FLAGS
type AMCONTROL_FLAGS DWORD

const (
	AMCONTROL_USED              AMCONTROL_FLAGS = 0x00000001
	AMCONTROL_PAD_TO_4x3        AMCONTROL_FLAGS = 0x00000002
	AMCONTROL_PAD_TO_16x9       AMCONTROL_FLAGS = 0x00000004
	AMCONTROL_COLORINFO_PRESENT AMCONTROL_FLAGS = 0x00000080
)

// VIDEOINFOHEADER2
type VIDEOINFOHEADER2 struct {
	Source           RECT
	Target           RECT
	BitRate          DWORD
	BitErrorRate     DWORD
	AvgTimePerFrame  REFERENCE_TIME
	InterlaceFlags   AMINTERLACE_FLAGS
	CopyProtectFlags AMCOPYPROTECT_FLAGS
	PictAspectRatioX DWORD
	PictAspectRatioY DWORD
	ControlFlags     AMCONTROL_FLAGS
	Reserved2        DWORD
	Header           BITMAPINFOHEADER
}

// PixelFormat
type PixelFormat UINT

const (
	PixelFormat4bppIndexed    PixelFormat = 0x30402
	PixelFormat8bppIndexed    PixelFormat = 0x30803
	PixelFormat16bppGrayScale PixelFormat = 0x101004
	PixelFormat16bppRGB555    PixelFormat = 0x21005
	PixelFormat16bppRGB565    PixelFormat = 0x21006
	PixelFormat16bppARGB1555  PixelFormat = 0x61007
	PixelFormat24bppRGB       PixelFormat = 0x21808
	PixelFormat32bppRGB       PixelFormat = 0x22009
	PixelFormat32bppARGB      PixelFormat = 0x26200a
	PixelFormat32bppPARGB     PixelFormat = 0xe200b
	PixelFormat48bppRGB       PixelFormat = 0x10300c
	PixelFormat64bppARGB      PixelFormat = 0x34400d
	PixelFormat64bppPARGB     PixelFormat = 0x1c400e
	PixelFormat32bppCMYK      PixelFormat = 0x200f
)

// InterpolationMode
type InterpolationMode INT

const (
	InterpolationModeInvalid             InterpolationMode = -1
	InterpolationModeDefault             InterpolationMode = 0
	InterpolationModeLowQuality          InterpolationMode = 1
	InterpolationModeHighQuality         InterpolationMode = 2
	InterpolationModeBilinear            InterpolationMode = 3
	InterpolationModeBicubic             InterpolationMode = 4
	InterpolationModeNearestNeighbor     InterpolationMode = 5
	InterpolationModeHighQualityBilinear InterpolationMode = 6
	InterpolationModeHighQualityBicubic  InterpolationMode = 7
)

// EncoderParameterValueType
type EncoderParameterValueType ULONG

const (
	EncoderParameterValueTypeByte          EncoderParameterValueType = 1
	EncoderParameterValueTypeASCII         EncoderParameterValueType = 2
	EncoderParameterValueTypeShort         EncoderParameterValueType = 3
	EncoderParameterValueTypeLong          EncoderParameterValueType = 4
	EncoderParameterValueTypeRational      EncoderParameterValueType = 5
	EncoderParameterValueTypeLongRange     EncoderParameterValueType = 6
	EncoderParameterValueTypeUndefined     EncoderParameterValueType = 7
	EncoderParameterValueTypeRationalRange EncoderParameterValueType = 8
	EncoderParameterValueTypePointer       EncoderParameterValueType = 9
)

// EncoderParameter
type EncoderParameter struct {
	Guid           GUID
	NumberOfValues ULONG
	Type           EncoderParameterValueType
	Value          LPVOID
}

// EncoderParameters
type EncoderParameters struct {
	Count     UINT
	Parameter [1]EncoderParameter
}

// ABC
type ABC struct {
	A int
	B UINT
	C int
}
type LPABC *ABC

// LCSCSTYPE
type LCSCSTYPE LONG

const (
	LCS_sRGB                LCSCSTYPE = 0x73524742
	LCS_WINDOWS_COLOR_SPACE LCSCSTYPE = 0x57696e20
	LCS_CALIBRATED_RGB      LCSCSTYPE = 0x00000000
)

type LCSCSTYPE_DWORD LCSCSTYPE

// LCSGAMUTMATCH
type LCSGAMUTMATCH LONG

const (
	LCS_GM_BUSINESS         LCSGAMUTMATCH = 0x00000001
	LCS_GM_GRAPHICS         LCSGAMUTMATCH = 0x00000002
	LCS_GM_IMAGES           LCSGAMUTMATCH = 0x00000004
	LCS_GM_ABS_COLORIMETRIC LCSGAMUTMATCH = 0x00000008
)

// CIEXYZ
type CIEXYZ struct {
	X FXPT2DOT30
	Y FXPT2DOT30
	Z FXPT2DOT30
}

// CIEXYZTRIPLE
type CIEXYZTRIPLE struct {
	Red   CIEXYZ
	Green CIEXYZ
	Blue  CIEXYZ
}

// LOGCOLORSPACE
type LOGCOLORSPACE struct {
	Signature  DWORD
	Version    DWORD
	Size       DWORD
	CSType     LCSCSTYPE
	Intent     LCSGAMUTMATCH
	Endpoints  CIEXYZTRIPLE
	GammaRed   DWORD
	GammaGreen DWORD
	GammaBlue  DWORD
	Filename   [260]WCHAR
}

type LPLOGCOLORSPACE *LOGCOLORSPACE

// BMFORMAT
type BMFORMAT UINT

const (
	BM_x555RGB             BMFORMAT = 0x0000
	BM_x555XYZ             BMFORMAT = 0x0101
	BM_x555Yxy             BMFORMAT = 0x0102
	BM_x555Lab             BMFORMAT = 0x0103
	BM_x555G3CH            BMFORMAT = 0x0104
	BM_RGBTRIPLETS         BMFORMAT = 0x0002
	BM_BGRTRIPLETS         BMFORMAT = 0x0004
	BM_XYZTRIPLETS         BMFORMAT = 0x0201
	BM_YxyTRIPLETS         BMFORMAT = 0x0202
	BM_LabTRIPLETS         BMFORMAT = 0x0203
	BM_G3CHTRIPLETS        BMFORMAT = 0x0204
	BM_5CHANNEL            BMFORMAT = 0x0205
	BM_6CHANNEL            BMFORMAT = 0x0206
	BM_7CHANNEL            BMFORMAT = 0x0207
	BM_8CHANNEL            BMFORMAT = 0x0208
	BM_GRAY                BMFORMAT = 0x0209
	BM_xRGBQUADS           BMFORMAT = 0x0008
	BM_xBGRQUADS           BMFORMAT = 0x0010
	BM_xG3CHQUADS          BMFORMAT = 0x0304
	BM_KYMCQUADS           BMFORMAT = 0x0305
	BM_CMYKQUADS           BMFORMAT = 0x0020
	BM_10b_RGB             BMFORMAT = 0x0009
	BM_10b_XYZ             BMFORMAT = 0x0401
	BM_10b_Yxy             BMFORMAT = 0x0402
	BM_10b_Lab             BMFORMAT = 0x0403
	BM_10b_G3CH            BMFORMAT = 0x0404
	BM_NAMED_INDEX         BMFORMAT = 0x0405
	BM_16b_RGB             BMFORMAT = 0x000A
	BM_16b_XYZ             BMFORMAT = 0x0501
	BM_16b_Yxy             BMFORMAT = 0x0502
	BM_16b_Lab             BMFORMAT = 0x0503
	BM_16b_G3CH            BMFORMAT = 0x0504
	BM_16b_GRAY            BMFORMAT = 0x0505
	BM_565RGB              BMFORMAT = 0x0001
	BM_32b_scRGB           BMFORMAT = 0x0601
	BM_32b_scARGB          BMFORMAT = 0x0602
	BM_S2DOT13FIXED_scRGB  BMFORMAT = 0x0603
	BM_S2DOT13FIXED_scARGB BMFORMAT = 0x0604
	BM_R10G10B10A2         BMFORMAT = 0x0701
	BM_R10G10B10A2_XR      BMFORMAT = 0x0702
	BM_R16G16B16A16_FLOAT  BMFORMAT = 0x0703
)

// NAMED_PROFILE_INFO
type NAMED_PROFILE_INFO struct {
	Flags               DWORD
	Count               DWORD
	CountDevCoordinates DWORD
	Prefix              [32]CHAR
	Suffix              [32]CHAR
}

type PNAMED_PROFILE_INFO *NAMED_PROFILE_INFO

// GRAYCOLOR
type GRAYCOLOR struct {
	Gray WORD
}

// RGBCOLOR
type RGBCOLOR struct {
	Red   WORD
	Green WORD
	Blue  WORD
}

// CMYKCOLOR
type CMYKCOLOR struct {
	Cyan    WORD
	Magenta WORD
	Yellow  WORD
	Black   WORD
}

// XYZCOLOR
type XYZCOLOR struct {
	X WORD
	Y WORD
	Z WORD
}

// YxyCOLOR
type YxyCOLOR struct {
	Y  WORD
	X  WORD
	Y2 WORD
}

// LabCOLOR
type LabCOLOR struct {
	L WORD
	A WORD
	B WORD
}

// GENERIC3CHANNEL
type GENERIC3CHANNEL struct {
	Ch1 WORD
	Ch2 WORD
	Ch3 WORD
}

// NAMEDCOLOR
type NAMEDCOLOR struct {
	Index DWORD
}

// HiFiCOLOR
type HiFiCOLOR struct {
	Channel [8]BYTE
}

// COLOR_s
type COLOR_s struct {
	Reserved1 DWORD
	Reserved2 LPVOID
}

// COLORDATATYPE
type COLORDATATYPE UINT

const (
	COLOR_BYTE               COLORDATATYPE = 1
	COLOR_WORD               COLORDATATYPE = 2
	COLOR_FLOAT              COLORDATATYPE = 3
	COLOR_S2DOT13FIXED       COLORDATATYPE = 4
	COLOR_10b_R10G10B10A2    COLORDATATYPE = 5
	COLOR_10b_R10G10B10A2_XR COLORDATATYPE = 6
	COLOR_FLOAT16            COLORDATATYPE = 7
)

// ENUMTYPE_FIELDS
type ENUMTYPE_FIELDS DWORD

const (
	ET_DEVICENAME      ENUMTYPE_FIELDS = 0x00000001
	ET_MEDIATYPE       ENUMTYPE_FIELDS = 0x00000002
	ET_DITHERMODE      ENUMTYPE_FIELDS = 0x00000004
	ET_RESOLUTION      ENUMTYPE_FIELDS = 0x00000008
	ET_CMMTYPE         ENUMTYPE_FIELDS = 0x00000010
	ET_CLASS           ENUMTYPE_FIELDS = 0x00000020
	ET_DATACOLORSPACE  ENUMTYPE_FIELDS = 0x00000040
	ET_CONNECTIONSPACE ENUMTYPE_FIELDS = 0x00000080
	ET_SIGNATURE       ENUMTYPE_FIELDS = 0x00000100
	ET_PLATFORM        ENUMTYPE_FIELDS = 0x00000200
	ET_PROFILEFLAGS    ENUMTYPE_FIELDS = 0x00000400
	ET_MANUFACTURER    ENUMTYPE_FIELDS = 0x00000800
	ET_MODEL           ENUMTYPE_FIELDS = 0x00001000
	ET_ATTRIBUTES      ENUMTYPE_FIELDS = 0x00002000
	ET_RENDERINGINTENT ENUMTYPE_FIELDS = 0x00004000
	ET_CREATOR         ENUMTYPE_FIELDS = 0x00008000
	ET_DEVICECLASS     ENUMTYPE_FIELDS = 0x00010000
)

// PROFILE_CLASS
type PROFILE_CLASS DWORD

const (
	CLASS_MONITOR    PROFILE_CLASS = 0x6d6e7472
	CLASS_PRINTER    PROFILE_CLASS = 0x70727472
	CLASS_SCANNER    PROFILE_CLASS = 0x73636e72
	CLASS_LINK       PROFILE_CLASS = 0x6c696e6b
	CLASS_ABSTRACT   PROFILE_CLASS = 0x61627374
	CLASS_COLORSPACE PROFILE_CLASS = 0x73706163
	CLASS_NAMED      PROFILE_CLASS = 0x6e6d636c
	CLASS_CAMP       PROFILE_CLASS = 0x63616d70
	CLASS_GMMP       PROFILE_CLASS = 0x676d6d70
)

// DEVICE_CLASS
type DEVICE_CLASS DWORD

const (
	DEVICE_MONITOR DEVICE_CLASS = 0x6d6e7472
	DEVICE_PRINTER DEVICE_CLASS = 0x70727472
	DEVICE_SCANNER DEVICE_CLASS = 0x73636e72
)

// COLOR_SPACE
type COLOR_SPACE DWORD

const (
	SPACE_XYZ       COLOR_SPACE = 0x58595A20
	SPACE_Lab       COLOR_SPACE = 0x4C616220
	SPACE_Luv       COLOR_SPACE = 0x4C757620
	SPACE_YCbCr     COLOR_SPACE = 0x59436272
	SPACE_Yxy       COLOR_SPACE = 0x59787920
	SPACE_RGB       COLOR_SPACE = 0x52474220
	SPACE_GRAY      COLOR_SPACE = 0x47524159
	SPACE_HSV       COLOR_SPACE = 0x48535620
	SPACE_HLS       COLOR_SPACE = 0x484C5320
	SPACE_CMYK      COLOR_SPACE = 0x434D594B
	SPACE_CMY       COLOR_SPACE = 0x434D5920
	SPACE_2_CHANNEL COLOR_SPACE = 0x32434C52
	SPACE_3_CHANNEL COLOR_SPACE = 0x33434C52
	SPACE_4_CHANNEL COLOR_SPACE = 0x34434C52
	SPACE_5_CHANNEL COLOR_SPACE = 0x35434C52
	SPACE_6_CHANNEL COLOR_SPACE = 0x36434C52
	SPACE_7_CHANNEL COLOR_SPACE = 0x37434C52
	SPACE_8_CHANNEL COLOR_SPACE = 0x38434C52
)

// CONNECTION_SPACE
type CONNECTION_SPACE DWORD

const (
	CONNECTION_XYZ CONNECTION_SPACE = 0x58595A20
	CONNECTION_Lab CONNECTION_SPACE = 0x4C616220
)

// PROFILE_FLAG
type PROFILE_FLAG DWORD

const (
	FLAG_EMBEDDEDPROFILE             PROFILE_FLAG = 0x00000001
	FLAG_DEPENDENTONDATA             PROFILE_FLAG = 0x00000002
	FLAG_ENABLE_CHROMATIC_ADAPTATION PROFILE_FLAG = 0x02000000
)

// PROFILE_ATTRIBUTES
type PROFILE_ATTRIBUTES DWORD

const (
	ATTRIB_TRANSPARENCY PROFILE_ATTRIBUTES = 0x00000001
	ATTRIB_MATTE        PROFILE_ATTRIBUTES = 0x00000002
)

// RENDERING_INTENT
type RENDERING_INTENT DWORD

const (
	INTENT_PERCEPTUAL            RENDERING_INTENT = 0
	INTENT_RELATIVE_COLORIMETRIC RENDERING_INTENT = 1
	INTENT_SATURATION            RENDERING_INTENT = 2
	INTENT_ABSOLUTE_COLORIMETRIC RENDERING_INTENT = 3
)

// ENUMTYPE
type ENUMTYPE struct {
	Size            DWORD
	Version         DWORD
	Fields          ENUMTYPE_FIELDS
	DeviceName      PCWSTR
	MediaType       DWORD
	DitheringMode   DWORD
	Resolution      [2]DWORD
	CMMType         DWORD
	Class           PROFILE_CLASS
	DataColorSpace  COLOR_SPACE
	ConnectionSpace CONNECTION_SPACE
	Signature       DWORD
	Platform        DWORD
	ProfileFlags    PROFILE_FLAG
	Manufacturer    DWORD
	Model           DWORD
	Attributes      [2]DWORD
	RenderingIntent RENDERING_INTENT
	Creator         DWORD
	DeviceClass     DEVICE_CLASS
}

type PENUMTYPE *ENUMTYPE

// CCT_FLAGS
type CCT_FLAGS DWORD

const (
	PROOF_MODE                CCT_FLAGS = 0x00000001
	NORMAL_MODE               CCT_FLAGS = 0x00000002
	BEST_MODE                 CCT_FLAGS = 0x00000003
	ENABLE_GAMUT_CHECKING     CCT_FLAGS = 0x00010000
	USE_RELATIVE_COLORIMETRIC CCT_FLAGS = 0x00020000
	FAST_TRANSLATE            CCT_FLAGS = 0x00040000
	PRESERVEBLACK             CCT_FLAGS = 0x00100000
	WCS_ALWAYS                CCT_FLAGS = 0x00200000
	SEQUENTIAL_TRANSFORM      CCT_FLAGS = 0x80800000
)

// CMM_INFO
type CMM_INFO DWORD

const (
	CMM_WIN_VERSION    CMM_INFO = 0
	CMM_IDENT          CMM_INFO = 1
	CMM_DRIVER_VERSION CMM_INFO = 2
	CMM_DLL_VERSION    CMM_INFO = 3
	CMM_VERSION        CMM_INFO = 4
	CMM_DESCRIPTION    CMM_INFO = 5
	CMM_LOGOICON       CMM_INFO = 6
)
