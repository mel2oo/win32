package shell32

import (
	"unsafe"

	"github.com/mel2oo/win32/types"
	"golang.org/x/sys/windows"
)

var (
	dll = windows.NewLazyDLL("shell32.dll")

	procAssocCreateForClasses                 = dll.NewProc("AssocCreateForClasses")
	procCommandLineToArgv                     = dll.NewProc("CommandLineToArgvW")
	procDoEnvironmentSubst                    = dll.NewProc("DoEnvironmentSubstW")
	procDragAcceptFiles                       = dll.NewProc("DragAcceptFiles")
	procDragFinish                            = dll.NewProc("DragFinish")
	procDragQueryFile                         = dll.NewProc("DragQueryFileW")
	procDragQueryPoint                        = dll.NewProc("DragQueryPoint")
	procDuplicateIcon                         = dll.NewProc("DuplicateIcon")
	procExtractAssociatedIconEx               = dll.NewProc("ExtractAssociatedIconExW")
	procExtractAssociatedIcon                 = dll.NewProc("ExtractAssociatedIconW")
	procExtractIconEx                         = dll.NewProc("ExtractIconExW")
	procExtractIcon                           = dll.NewProc("ExtractIconW")
	procFindExecutable                        = dll.NewProc("FindExecutableW")
	procInitNetworkAddressControl             = dll.NewProc("InitNetworkAddressControl")
	procNetAddrDisplayErrorTip                = dll.NewProc("NetAddr_DisplayErrorTip")
	procNetAddrGetAddress                     = dll.NewProc("NetAddr_GetAddress")
	procNetAddrGetAllowType                   = dll.NewProc("NetAddr_GetAllowType")
	procNetAddrSetAllowType                   = dll.NewProc("NetAddr_SetAllowType")
	procSHAppBarMessage                       = dll.NewProc("SHAppBarMessage")
	procSHCreateProcessAsUser                 = dll.NewProc("SHCreateProcessAsUserW")
	procShellNotifyIconGetRect                = dll.NewProc("Shell_NotifyIconGetRect")
	procShellNotifyIcon                       = dll.NewProc("Shell_NotifyIconW")
	procShellAbout                            = dll.NewProc("ShellAboutW")
	procShellExecuteEx                        = dll.NewProc("ShellExecuteExW")
	procShellExecute                          = dll.NewProc("ShellExecuteW")
	procShellMessageBox                       = dll.NewProc("ShellMessageBoxW")
	procSHEmptyRecycleBin                     = dll.NewProc("SHEmptyRecycleBinW")
	procSHEnumerateUnreadMailAccounts         = dll.NewProc("SHEnumerateUnreadMailAccountsW")
	procSHEvaluateSystemCommandTemplate       = dll.NewProc("SHEvaluateSystemCommandTemplate")
	procSHFileOperation                       = dll.NewProc("SHFileOperationW")
	procSHFreeNameMappings                    = dll.NewProc("SHFreeNameMappings")
	procSHGetDiskFreeSpaceEx                  = dll.NewProc("SHGetDiskFreeSpaceExW")
	procSHGetDriveMedia                       = dll.NewProc("SHGetDriveMedia")
	procSHGetFileInfo                         = dll.NewProc("SHGetFileInfoW")
	procSHGetImageList                        = dll.NewProc("SHGetImageList")
	procSHGetLocalizedName                    = dll.NewProc("SHGetLocalizedName")
	procSHGetNewLinkInfo                      = dll.NewProc("SHGetNewLinkInfoW")
	procSHGetPropertyStoreForWindow           = dll.NewProc("SHGetPropertyStoreForWindow")
	procSHGetStockIconInfo                    = dll.NewProc("SHGetStockIconInfo")
	procSHGetUnreadMailCount                  = dll.NewProc("SHGetUnreadMailCountW")
	procSHInvokePrinterCommand                = dll.NewProc("SHInvokePrinterCommandW")
	procSHIsFileAvailableOffline              = dll.NewProc("SHIsFileAvailableOffline")
	procSHLoadNonloadedIconOverlayIdentifiers = dll.NewProc("SHLoadNonloadedIconOverlayIdentifiers")
	procSHQueryRecycleBin                     = dll.NewProc("SHQueryRecycleBinW")
	procSHQueryUserNotificationState          = dll.NewProc("SHQueryUserNotificationState")
	procSHRemoveLocalizedName                 = dll.NewProc("SHRemoveLocalizedName")
	procSHSetLocalizedName                    = dll.NewProc("SHSetLocalizedName")
	procSHSetUnreadMailCount                  = dll.NewProc("SHSetUnreadMailCountW")
	procSHTestTokenMembership                 = dll.NewProc("SHTestTokenMembership")
)

// https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellexecuteexw
func ShellExecuteExW(lpExecInfo *types.SHELLEXECUTEINFO) types.BOOL {
	ret, _, _ := procShellExecuteEx.Call(uintptr(unsafe.Pointer(lpExecInfo)))
	return types.BOOL(ret)
}
