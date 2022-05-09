package shell32

import "github.com/mel2oo/win32/types"

// RESTRICTIONS
type RESTRICTIONS types.UINT

// Bitfields
type (
	LPSHELLSTATE     types.LPVOID
	LPSHELLFLAGSTATE types.LPVOID
)

// Interfaces
type (
	IShellFolder               uintptr
	IContextMenu               uintptr
	IDataObject                uintptr
	IEnumAssocHandlers         uintptr
	IShellItem                 uintptr
	IFileOperation             uintptr
	IShellView                 uintptr
	IShellItemArray            uintptr
	IShellBrowser              uintptr
	IFileOperationProgressSink uintptr
	IShellTaskScheduler        uintptr
	IDropSource                uintptr
	IContextMenuCB             uintptr
	IShellFolderViewCB         uintptr
)

// Variables
var (
	FARPROC16                types.LPVOID
	LPFNDFMCALLBACK          types.LPVOID
	HDROP                    types.HANDLE
	HPSXA                    types.HANDLE
	LPFNADDPROPSHEETPAGE     types.LPVOID
	LPFNVIEWCALLBACK         types.LPVOID
	PFNASYNCICONTASKBALLBACK types.LPVOID
)

// AUTO_SCROLL_DATA
type AUTO_SCROLL_DATA struct {
	NextSample int
	LastScroll types.DWORD
	Full       types.BOOL
	Pts        [3]types.POINT
	Times      [3]types.DWORD
}

// BROWSEINFO_FLAG
type BROWSEINFO_FLAG types.UINT

const (
	BIF_RETURNONLYFSDIRS    BROWSEINFO_FLAG = 0x00000001
	BIF_DONTGOBELOWDOMAIN   BROWSEINFO_FLAG = 0x00000002
	BIF_STATUSTEXT          BROWSEINFO_FLAG = 0x00000004
	BIF_RETURNFSANCESTORS   BROWSEINFO_FLAG = 0x00000008
	BIF_EDITBOX             BROWSEINFO_FLAG = 0x00000010
	BIF_VALIDATE            BROWSEINFO_FLAG = 0x00000020
	BIF_NEWDIALOGSTYLE      BROWSEINFO_FLAG = 0x00000040
	BIF_USENEWUI            BROWSEINFO_FLAG = 0x00000050
	BIF_BROWSEINCLUDEURLS   BROWSEINFO_FLAG = 0x00000080
	BIF_UAHINT              BROWSEINFO_FLAG = 0x00000100
	BIF_NONEWFOLDERBUTTON   BROWSEINFO_FLAG = 0x00000200
	BIF_NOTRANSLATETARGETS  BROWSEINFO_FLAG = 0x00000400
	BIF_BROWSEFORCOMPUTER   BROWSEINFO_FLAG = 0x00001000
	BIF_BROWSEFORPRINTER    BROWSEINFO_FLAG = 0x00002000
	BIF_BROWSEINCLUDEFILES  BROWSEINFO_FLAG = 0x00004000
	BIF_SHAREABLE           BROWSEINFO_FLAG = 0x00008000
	BIF_BROWSEFILEJUNCTIONS BROWSEINFO_FLAG = 0x00010000
)

// BROWSEINFO
// type BROWSEINFO struct {
// 	hwndOwner      types.HWND
// 	pidlRoot       types.PCIDLIST_ABSOLUTE
// 	pszDisplayName types.LPWSTR
// 	lpszTitle      types.LPCWSTR
// 	ulFlags        BROWSEINFO_FLAG
// 	lpfn           types.BFFCALLBACK
// 	lParam         types.LPARAM
// 	iImage         int
// }
