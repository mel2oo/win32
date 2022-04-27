//go:build windows
// +build windows

package advapi32

import (
	"golang.org/x/sys/windows"
)

var (
	dll = windows.NewLazyDLL("advapi32.dll")

	// winreg.h
	ProcAbortSystemShutdown         = dll.NewProc("AbortSystemShutdownW")
	ProcInitiateShutdown            = dll.NewProc("InitiateShutdownW")
	ProcInitiateSystemShutdown      = dll.NewProc("InitiateSystemShutdownW")
	ProcInitiateSystemShutdownEx    = dll.NewProc("InitiateSystemShutdownExW")
	ProcRegCloseKey                 = dll.NewProc("RegCloseKey")
	ProcRegConnectRegistry          = dll.NewProc("RegConnectRegistryW")
	ProcRegCopyTree                 = dll.NewProc("RegCopyTreeW")
	ProcRegCreateKeyEx              = dll.NewProc("RegCreateKeyExW")
	ProcRegCreateKeyTransacted      = dll.NewProc("RegCreateKeyTransactedW")
	ProcRegCreateKey                = dll.NewProc("RegCreateKeyW")
	ProcRegDeleteKeyEx              = dll.NewProc("RegDeleteKeyExW")
	ProcRegDeleteKeyTransacted      = dll.NewProc("RegDeleteKeyTransactedW")
	ProcRegDeleteKeyValue           = dll.NewProc("RegDeleteKeyValueW")
	ProcRegDeleteKey                = dll.NewProc("RegDeleteKeyW")
	ProcRegDeleteTree               = dll.NewProc("RegDeleteTreeW")
	ProcRegDeleteValue              = dll.NewProc("RegDeleteValueW")
	ProcRegDisablePredefinedCache   = dll.NewProc("RegDisablePredefinedCache")
	ProcRegDisablePredefinedCacheEx = dll.NewProc("RegDisablePredefinedCacheEx")
	ProcRegDisableReflectionKey     = dll.NewProc("RegDisableReflectionKey")
	ProcRegEnableReflectionKey      = dll.NewProc("RegEnableReflectionKey")
	ProcRegEnumKeyEx                = dll.NewProc("RegEnumKeyExW")
	ProcRegEnumKey                  = dll.NewProc("RegEnumKeyW")
	ProcRegEnumValue                = dll.NewProc("RegEnumValueW")
	ProcRegFlushKey                 = dll.NewProc("RegFlushKey")
	ProcRegGetKeySecurity           = dll.NewProc("RegGetKeySecurity")
	ProcRegGetValue                 = dll.NewProc("RegGetValueW")
	ProcRegLoadAppKey               = dll.NewProc("RegLoadAppKeyW")
	ProcRegLoadKey                  = dll.NewProc("RegLoadKeyW")
	ProcRegLoadMUIString            = dll.NewProc("RegLoadMUIStringW")
	ProcRegNotifyChangeKeyValue     = dll.NewProc("RegNotifyChangeKeyValue")
	ProcRegOpenCurrentUser          = dll.NewProc("RegOpenCurrentUser")
	ProcRegOpenKeyEx                = dll.NewProc("RegOpenKeyExW")
	ProcRegOpenKeyTransacted        = dll.NewProc("RegOpenKeyTransactedW")
	ProcRegOpenKey                  = dll.NewProc("RegOpenKeyW")
	ProcRegOpenUserClassesRoot      = dll.NewProc("RegOpenUserClassesRoot")
	ProcRegOverridePredefKey        = dll.NewProc("RegOverridePredefKey")
	ProcRegQueryInfoKey             = dll.NewProc("RegQueryInfoKeyW")
	ProcRegQueryMultipleValues      = dll.NewProc("RegQueryMultipleValuesW")
	ProcRegQueryReflectionKey       = dll.NewProc("RegQueryReflectionKey")
	ProcRegQueryValueEx             = dll.NewProc("RegQueryValueExW")
	ProcRegQueryValue               = dll.NewProc("RegQueryValueW")
	ProcRegRenameKey                = dll.NewProc("RegRenameKey")
	ProcRegReplaceKey               = dll.NewProc("RegReplaceKeyW")
	ProcRegRestoreKey               = dll.NewProc("RegRestoreKeyW")
	ProcRegSaveKeyEx                = dll.NewProc("RegSaveKeyExW")
	ProcRegSaveKey                  = dll.NewProc("RegSaveKeyW")
	ProcRegSetKeySecurity           = dll.NewProc("RegSetKeySecurity")
	ProcRegSetKeyValue              = dll.NewProc("RegSetKeyValueW")
	ProcRegSetValueEx               = dll.NewProc("RegSetValueExW")
	ProcRegSetValue                 = dll.NewProc("RegSetValueW")
	ProcRegUnLoadKey                = dll.NewProc("RegUnLoadKeyW")

	// winsvc.h
	ProcChangeServiceConfig2             = dll.NewProc("ChangeServiceConfig2W")
	ProcChangeServiceConfig              = dll.NewProc("ChangeServiceConfigW")
	ProcCloseServiceHandle               = dll.NewProc("CloseServiceHandle")
	ProcControlService                   = dll.NewProc("ControlService")
	ProcControlServiceEx                 = dll.NewProc("ControlServiceExW")
	ProcCreateService                    = dll.NewProc("CreateServiceW")
	ProcDeleteService                    = dll.NewProc("DeleteService")
	ProcEnumDependentServices            = dll.NewProc("EnumDependentServicesW")
	ProcEnumServicesStatusEx             = dll.NewProc("EnumServicesStatusExW")
	ProcEnumServicesStatus               = dll.NewProc("EnumServicesStatusW")
	ProcGetServiceDirectory              = dll.NewProc("GetServiceDirectory")
	ProcGetServiceDisplayName            = dll.NewProc("GetServiceDisplayNameW")
	ProcGetServiceKeyName                = dll.NewProc("GetServiceKeyNameW")
	ProcGetServiceRegistryStateKey       = dll.NewProc("GetServiceRegistryStateKey")
	ProcGetSharedServiceDirectory        = dll.NewProc("GetSharedServiceDirectory")
	ProcGetSharedServiceRegistryStateKey = dll.NewProc("GetSharedServiceRegistryStateKey")
	ProcLockServiceDatabase              = dll.NewProc("LockServiceDatabase")
	ProcNotifyBootConfigStatus           = dll.NewProc("NotifyBootConfigStatus")
	ProcNotifyServiceStatusChange        = dll.NewProc("NotifyServiceStatusChangeW")
	ProcOpenSCManager                    = dll.NewProc("OpenSCManagerW")
	ProcOpenService                      = dll.NewProc("OpenServiceW")
	ProcQueryServiceConfig2              = dll.NewProc("QueryServiceConfig2W")
	ProcQueryServiceConfig               = dll.NewProc("QueryServiceConfigW")
	ProcQueryServiceDynamicInformation   = dll.NewProc("QueryServiceDynamicInformation")
	ProcQueryServiceLockStatus           = dll.NewProc("QueryServiceLockStatusW")
	ProcQueryServiceObjectSecurity       = dll.NewProc("QueryServiceObjectSecurity")
	ProcQueryServiceStatus               = dll.NewProc("QueryServiceStatus")
	ProcQueryServiceStatusEx             = dll.NewProc("QueryServiceStatusEx")
	ProcRegisterServiceCtrlHandler       = dll.NewProc("RegisterServiceCtrlHandlerW")
	ProcRegisterServiceCtrlHandlerEx     = dll.NewProc("RegisterServiceCtrlHandlerExW")
	ProcSetServiceObjectSecurity         = dll.NewProc("SetServiceObjectSecurity")
	ProcSetServiceStatus                 = dll.NewProc("SetServiceStatus")
	ProcStartService                     = dll.NewProc("StartServiceW")
	ProcStartServiceCtrlDispatcher       = dll.NewProc("StartServiceCtrlDispatcherW")
	ProcUnlockServiceDatabase            = dll.NewProc("UnlockServiceDatabase")

	// evntrace.h
	ProcCloseTrace                 = dll.NewProc("CloseTrace")
	ProcControlTrace               = dll.NewProc("ControlTraceW")
	ProcCreateTraceInstanceId      = dll.NewProc("CreateTraceInstanceId")
	ProcEnableTrace                = dll.NewProc("EnableTrace")
	ProcEnableTraceEx              = dll.NewProc("EnableTraceEx")
	ProcEnableTraceEx2             = dll.NewProc("EnableTraceEx2")
	ProcEnumerateTraceGuids        = dll.NewProc("EnumerateTraceGuids")
	ProcEnumerateTraceGuidsEx      = dll.NewProc("EnumerateTraceGuidsEx")
	ProcFlushTrace                 = dll.NewProc("FlushTraceW")
	ProcGetTraceEnableFlags        = dll.NewProc("GetTraceEnableFlags")
	ProcGetTraceEnableLevel        = dll.NewProc("GetTraceEnableLevel")
	ProcGetTraceLoggerHandle       = dll.NewProc("GetTraceLoggerHandle")
	ProcOpenTrace                  = dll.NewProc("OpenTraceW")
	ProcProcessTrace               = dll.NewProc("ProcessTrace")
	ProcQueryAllTraces             = dll.NewProc("QueryAllTracesW")
	ProcQueryTraceProcessingHandle = dll.NewProc("QueryTraceProcessingHandle")
	ProcQueryTrace                 = dll.NewProc("QueryTraceW")
	ProcRegisterTraceGuids         = dll.NewProc("RegisterTraceGuidsW")
	ProcRemoveTraceCallback        = dll.NewProc("RemoveTraceCallback")
	ProcSetTraceCallback           = dll.NewProc("SetTraceCallback")
	ProcStartTrace                 = dll.NewProc("StartTraceW")
	ProcStopTrace                  = dll.NewProc("StopTraceW")
	ProcTraceEvent                 = dll.NewProc("TraceEvent")
	ProcTraceEventInstance         = dll.NewProc("TraceEventInstance")
	ProcTraceMessage               = dll.NewProc("TraceMessage")
	ProcTraceMessageVa             = dll.NewProc("TraceMessageVa")
	ProcTraceQueryInformation      = dll.NewProc("TraceQueryInformation")
	ProcTraceSetInformation        = dll.NewProc("TraceSetInformation")
	ProcUnregisterTraceGuids       = dll.NewProc("UnregisterTraceGuids")
	ProcUpdateTrace                = dll.NewProc("UpdateTraceW")
)
