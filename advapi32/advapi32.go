//go:build windows
// +build windows

package advapi32

import (
	"github.com/mel2oo/win32/types"
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

// REG_ERROR
type REG_ERROR types.ERROR_CODE

// HW_PROFILE_INFO
type HW_PROFILE_INFO struct {
	DwDockInfo      types.DWORD
	SzHwProfileGuid [39]types.WCHAR
	SzHwProfileName [80]types.WCHAR
}
type LPHW_PROFILE_INFO *HW_PROFILE_INFO

// WCT_OBJECT_TYPE
type WCT_OBJECT_TYPE types.UINT

const (
	WctCriticalSectionType WCT_OBJECT_TYPE = 1
	WctSendMessageType     WCT_OBJECT_TYPE = 2
	WctMutexType           WCT_OBJECT_TYPE = 3
	WctAlpcType            WCT_OBJECT_TYPE = 4
	WctComType             WCT_OBJECT_TYPE = 5
	WctThreadWaitType      WCT_OBJECT_TYPE = 6
	WctProcessWaitType     WCT_OBJECT_TYPE = 7
	WctThreadType          WCT_OBJECT_TYPE = 8
	WctComActivationType   WCT_OBJECT_TYPE = 9
	WctUnknownType         WCT_OBJECT_TYPE = 10
	WctSocketIoType        WCT_OBJECT_TYPE = 11
	WctSmbIoType           WCT_OBJECT_TYPE = 12
)

// WCT_OBJECT_STATUS
type WCT_OBJECT_STATUS types.UINT

const (
	WctStatusNoAccess     WCT_OBJECT_STATUS = 1
	WctStatusRunning      WCT_OBJECT_STATUS = 2
	WctStatusBlocked      WCT_OBJECT_STATUS = 3
	WctStatusPidOnly      WCT_OBJECT_STATUS = 4
	WctStatusPidOnlyRpcss WCT_OBJECT_STATUS = 5
	WctStatusOwned        WCT_OBJECT_STATUS = 6
	WctStatusNotOwned     WCT_OBJECT_STATUS = 7
	WctStatusAbandoned    WCT_OBJECT_STATUS = 8
	WctStatusUnknown      WCT_OBJECT_STATUS = 9
	WctStatusError        WCT_OBJECT_STATUS = 10
)

// WAITCHAIN_NODE_INFO_u_s1
type WAITCHAIN_NODE_INFO_u_s1 struct {
	ObjectName [128]types.WCHAR
	Timeout    types.LARGE_INTEGER
	Alertable  types.BOOL
}

// WAITCHAIN_NODE_INFO_u_s2
type WAITCHAIN_NODE_INFO_u_s2 struct {
	ProcessId       types.DWORD
	ThreadId        types.DWORD
	WaitTime        types.DWORD
	ContextSwitches types.DWORD
}

// <Variable Name="[WAITCHAIN_NODE_INFO_u]" Type="Union">
// <Display Name="union" />
// <Field Type="[WAITCHAIN_NODE_INFO_u_s1]"    Name="LockObject" />
// <Field Type="[WAITCHAIN_NODE_INFO_u_s2]"    Name="ThreadObject" />
// </Variable>
type WAITCHAIN_NODE_INFO_u [16]byte

// WAITCHAIN_NODE_INFO
type WAITCHAIN_NODE_INFO struct {
	ObjectType    WCT_OBJECT_TYPE
	ObjectStatus  WCT_OBJECT_STATUS
	ObjectStatusU WAITCHAIN_NODE_INFO_u
}
type PWAITCHAIN_NODE_INFO *WAITCHAIN_NODE_INFO

// SC_ENUM_TYPE
type SC_ENUM_TYPE types.UINT

const (
	SC_ENUM_PROCESS_INFO SC_ENUM_TYPE = 0
)

// TRACE_INFO_CLASS
type TRACE_INFO_CLASS types.UINT

const (
	TraceGuidQueryList              TRACE_INFO_CLASS = 0
	TraceGuidQueryInfo              TRACE_INFO_CLASS = 1
	TraceGuidQueryProcess           TRACE_INFO_CLASS = 2
	TraceStackTracingInfo           TRACE_INFO_CLASS = 3
	TraceSystemTraceEnableFlagsInfo TRACE_INFO_CLASS = 4
	TraceSampledProfileIntervalInfo TRACE_INFO_CLASS = 5
	TraceProfileSourceConfigInfo    TRACE_INFO_CLASS = 6
	TraceProfileSourceListInfo      TRACE_INFO_CLASS = 7
	TracePmcEventListInfo           TRACE_INFO_CLASS = 8
	TracePmcCounterListInfo         TRACE_INFO_CLASS = 9
)

// TRACE_QUERY_INFO_CLASS
type TRACE_QUERY_INFO_CLASS TRACE_INFO_CLASS

// PROG_INVOKE_SETTING
type PROG_INVOKE_SETTING types.UINT

const (
	ProgressInvokeNever        PROG_INVOKE_SETTING = 1
	ProgressInvokeEveryObject  PROG_INVOKE_SETTING = 2
	ProgressInvokeOnError      PROG_INVOKE_SETTING = 3
	ProgressCancelOperation    PROG_INVOKE_SETTING = 4
	ProgressRetryOperation     PROG_INVOKE_SETTING = 5
	ProgressInvokePrePostError PROG_INVOKE_SETTING = 6
)

// TRUSTED_INFORMATION_CLASS
type TRUSTED_INFORMATION_CLASS types.UINT

const (
	TrustedDomainNameInformation          TRUSTED_INFORMATION_CLASS = 1
	TrustedControllersInformation         TRUSTED_INFORMATION_CLASS = 2
	TrustedPosixOffsetInformation         TRUSTED_INFORMATION_CLASS = 3
	TrustedPasswordInformation            TRUSTED_INFORMATION_CLASS = 4
	TrustedDomainInformationBasic         TRUSTED_INFORMATION_CLASS = 5
	TrustedDomainInformationEx            TRUSTED_INFORMATION_CLASS = 6
	TrustedDomainAuthInformation          TRUSTED_INFORMATION_CLASS = 7
	TrustedDomainFullInformation          TRUSTED_INFORMATION_CLASS = 8
	TrustedDomainAuthInformationInternal  TRUSTED_INFORMATION_CLASS = 9
	TrustedDomainFullInformationInternal  TRUSTED_INFORMATION_CLASS = 10
	TrustedDomainInformationEx2Internal   TRUSTED_INFORMATION_CLASS = 11
	TrustedDomainFullInformation2Internal TRUSTED_INFORMATION_CLASS = 12
	TrustedDomainSupportedEncryptionTypes TRUSTED_INFORMATION_CLASS = 13
)

// SEF_Flags
type SEF_Flags types.ULONG

const (
	SEF_DACL_AUTO_INHERIT             SEF_Flags = 0x01
	SEF_SACL_AUTO_INHERIT             SEF_Flags = 0x02
	SEF_DEFAULT_DESCRIPTOR_FOR_OBJECT SEF_Flags = 0x04
	SEF_AVOID_PRIVILEGE_CHECK         SEF_Flags = 0x08
	SEF_AVOID_OWNER_CHECK             SEF_Flags = 0x10
	SEF_DEFAULT_OWNER_FROM_PARENT     SEF_Flags = 0x20
	SEF_DEFAULT_GROUP_FROM_PARENT     SEF_Flags = 0x40
	SEF_MACL_NO_WRITE_UP              SEF_Flags = 0x100
	SEF_MACL_NO_READ_UP               SEF_Flags = 0x200
	SEF_MACL_NO_EXECUTE_UP            SEF_Flags = 0x400
	SEF_AVOID_OWNER_RESTRICTION       SEF_Flags = 0x1000
)

// SaferScopeId
type SaferScopeId types.DWORD

const (
	SAFER_SCOPEID_MACHINE SaferScopeId = 1
	SAFER_SCOPEID_USER    SaferScopeId = 2
)

// MandatoryPolicyFlags
type MandatoryPolicyFlags types.DWORD

const (
	SYSTEM_MANDATORY_LABEL_NO_WRITE_UP   MandatoryPolicyFlags = 0x1
	SYSTEM_MANDATORY_LABEL_NO_READ_UP    MandatoryPolicyFlags = 0x2
	SYSTEM_MANDATORY_LABEL_NO_EXECUTE_UP MandatoryPolicyFlags = 0x4
)

// LogonFlags
type LogonFlags types.DWORD

const (
	LOGON_WITH_PROFILE         LogonFlags = 0x1
	LOGON_NETCREDENTIALS_ONLY  LogonFlags = 0x2
	LOGON_ZERO_PASSWORD_BUFFER LogonFlags = 0x4
)

// LogonType
type LogonType types.DWORD

const (
	LOGON32_LOGON_INTERACTIVE       LogonType = 2
	LOGON32_LOGON_NETWORK           LogonType = 3
	LOGON32_LOGON_BATCH             LogonType = 4
	LOGON32_LOGON_SERVICE           LogonType = 5
	LOGON32_LOGON_UNLOCK            LogonType = 7
	LOGON32_LOGON_NETWORK_CLEARTEXT LogonType = 8
	LOGON32_LOGON_NEW_CREDENTIALS   LogonType = 9
)

// LogonProvider
type LogonProvider types.DWORD

const (
	LOGON32_PROVIDER_DEFAULT LogonProvider = 0
	LOGON32_PROVIDER_WINNT35 LogonProvider = 1
	LOGON32_PROVIDER_WINNT40 LogonProvider = 2
	LOGON32_PROVIDER_WINNT50 LogonProvider = 3
	LOGON32_PROVIDER_VIRTUAL LogonProvider = 4
)

// TreeSecAction
type TreeSecAction types.DWORD

const (
	TREE_SEC_INFO_SET                 TreeSecAction = 0x00000001
	TREE_SEC_INFO_RESET               TreeSecAction = 0x00000002
	TREE_SEC_INFO_RESET_KEEP_EXPLICIT TreeSecAction = 0x00000003
)

// CredGetTargetInfoFlags
type CredGetTargetInfoFlags types.DWORD

const (
	CRED_ALLOW_NAME_RESOLUTION CredGetTargetInfoFlags = 0x1
)

// CredType
type CredType types.DWORD

const (
	CRED_TYPE_GENERIC                 CredType = 1
	CRED_TYPE_DOMAIN_PASSWORD         CredType = 2
	CRED_TYPE_DOMAIN_CERTIFICATE      CredType = 3
	CRED_TYPE_DOMAIN_VISIBLE_PASSWORD CredType = 4
	CRED_TYPE_GENERIC_CERTIFICATE     CredType = 5
	CRED_TYPE_DOMAIN_EXTENDED         CredType = 6
)

// CredEnumerateFlags
type CredEnumerateFlags types.DWORD

const (
	CRED_ENUMERATE_ALL_CREDENTIALS CredEnumerateFlags = 0x1
)

// CreateRestrictedTokenFlags
type CreateRestrictedTokenFlags types.DWORD

const (
	DISABLE_MAX_PRIVILEGE CreateRestrictedTokenFlags = 0x1
	SANDBOX_INERT         CreateRestrictedTokenFlags = 0x2
	LUA_TOKEN             CreateRestrictedTokenFlags = 0x4
	WRITE_RESTRICTED      CreateRestrictedTokenFlags = 0x8
)

// POLICY_INFORMATION_CLASS
type POLICY_INFORMATION_CLASS types.UINT

const (
	PolicyAuditLogInformation           POLICY_INFORMATION_CLASS = 1
	PolicyAuditEventsInformation        POLICY_INFORMATION_CLASS = 2
	PolicyPrimaryDomainInformation      POLICY_INFORMATION_CLASS = 3
	PolicyPdAccountInformation          POLICY_INFORMATION_CLASS = 4
	PolicyAccountDomainInformation      POLICY_INFORMATION_CLASS = 5
	PolicyLsaServerRoleInformation      POLICY_INFORMATION_CLASS = 6
	PolicyReplicaSourceInformation      POLICY_INFORMATION_CLASS = 7
	PolicyDefaultQuotaInformation       POLICY_INFORMATION_CLASS = 8
	PolicyModificationInformation       POLICY_INFORMATION_CLASS = 9
	PolicyAuditFullSetInformation       POLICY_INFORMATION_CLASS = 10
	PolicyAuditFullQueryInformation     POLICY_INFORMATION_CLASS = 11
	PolicyDnsDomainInformation          POLICY_INFORMATION_CLASS = 12
	PolicyDnsDomainInformationInt       POLICY_INFORMATION_CLASS = 13
	PolicyLocalAccountDomainInformation POLICY_INFORMATION_CLASS = 14
)

// POLICY_DOMAIN_INFORMATION_CLASS
type POLICY_DOMAIN_INFORMATION_CLASS types.UINT

const (
	PolicyDomainQualityOfServiceInformation POLICY_DOMAIN_INFORMATION_CLASS = 1
	PolicyDomainEfsInformation              POLICY_DOMAIN_INFORMATION_CLASS = 2
	PolicyDomainKerberosTicketInformation   POLICY_DOMAIN_INFORMATION_CLASS = 3
)

// EFS_HASH_BLOB
type EFS_HASH_BLOB struct {
	CbData types.DWORD
	PbData types.PBYTE
}
type PEFS_HASH_BLOB *EFS_HASH_BLOB

//ENCRYPTION_CERTIFICATE_HASH
type ENCRYPTION_CERTIFICATE_HASH struct {
	TotalLength        types.DWORD
	UserSid            types.PSID
	Hash               PEFS_HASH_BLOB
	DisplayInformation types.LPWSTR
}
type PENCRYPTION_CERTIFICATE_HASH *ENCRYPTION_CERTIFICATE_HASH

// ENCRYPTION_CERTIFICATE_HASH_LIST
type ENCRYPTION_CERTIFICATE_HASH_LIST struct {
	CertHash types.DWORD
	Users    *PENCRYPTION_CERTIFICATE_HASH
}
type PENCRYPTION_CERTIFICATE_HASH_LIST *ENCRYPTION_CERTIFICATE_HASH_LIST

// INSTALLSPECTYPE
type INSTALLSPECTYPE types.UINT

const (
	APPNAME  INSTALLSPECTYPE = 1
	FILEEXT  INSTALLSPECTYPE = 2
	PROGID   INSTALLSPECTYPE = 3
	COMCLASS INSTALLSPECTYPE = 4
)

// INSTALLSPEC_s1
type INSTALLSPEC_s1 struct {
	Name  *types.WCHAR
	GPOId types.GUID
}

// INSTALLSPEC_s2
type INSTALLSPEC_s2 struct {
	Clsid  types.GUID
	ClsCtx types.CLSCTX
}

// INSTALLSPEC
type INSTALLSPEC INSTALLSPEC_s2

// INSTALLDATA
type INSTALLDATA struct {
	Type INSTALLSPECTYPE
	Spec INSTALLSPEC
}
type PINSTALLDATA *INSTALLDATA

// VALENT
type VALENT struct {
	Valuename types.LPWSTR
	Valuelen  types.DWORD
	Valueptr  types.DWORD_PTR
	Vtype     types.DWORD
}
type PVALENT *VALENT

// EFS_CERTIFICATE_BLOB
type EFS_CERTIFICATE_BLOB struct {
	CertEncodingType types.DWORD
	CbData           types.DWORD
	PbData           types.PBYTE
}
type PEFS_CERTIFICATE_BLOB *EFS_CERTIFICATE_BLOB

// ENCRYPTION_CERTIFICATE
type ENCRYPTION_CERTIFICATE struct {
	CbTotalLength types.DWORD
	PUserSid      types.PSID
	PCertBlob     PEFS_CERTIFICATE_BLOB
}
type PENCRYPTION_CERTIFICATE *ENCRYPTION_CERTIFICATE

// ENCRYPTION_CERTIFICATE_LIST
type ENCRYPTION_CERTIFICATE_LIST struct {
	NUsers types.DWORD
	PUsers *PENCRYPTION_CERTIFICATE
}
type PENCRYPTION_CERTIFICATE_LIST *ENCRYPTION_CERTIFICATE_LIST

// LocalState
type LocalState types.DWORD

const (
	LOCALSTATE_ASSIGNED               LocalState = 0x1
	LOCALSTATE_PUBLISHED              LocalState = 0x2
	LOCALSTATE_UNINSTALL_UNMANAGED    LocalState = 0x4
	LOCALSTATE_POLICYREMOVE_ORPHAN    LocalState = 0x8
	LOCALSTATE_POLICYREMOVE_UNINSTALL LocalState = 0x10
	LOCALSTATE_ORPHANED               LocalState = 0x20
	LOCALSTATE_UNINSTALLED            LocalState = 0x40
)

// LOCALMANAGEDAPPLICATION
type LOCALMANAGEDAPPLICATION struct {
	PszDeploymentName types.LPWSTR
	PszPolicyName     types.LPWSTR
	PszProductId      types.LPWSTR
	DwState           LocalState
}

type PLOCALMANAGEDAPPLICATION *LOCALMANAGEDAPPLICATION

// MANAGEDAPPLICATION
type MANAGEDAPPLICATION struct {
	PszPackageName types.LPWSTR
	PszPublisher   types.LPWSTR
	DwVersionHi    types.DWORD
	DwVersionLo    types.DWORD
	DwRevision     types.DWORD
	GpoId          types.GUID
	PszPolicyName  types.LPWSTR
	ProductId      types.GUID
	Language       types.LANGID
	PszOwner       types.LPWSTR
	PszCompany     types.LPWSTR
	PszComments    types.LPWSTR
	PszContact     types.LPWSTR
	PszSupportUrl  types.LPWSTR
	DwPathType     types.DWORD
	BInstalled     types.BOOL
}
type PMANAGEDAPPLICATION *MANAGEDAPPLICATION

// PERF_COUNTERSET_INSTANCE
type PERF_COUNTERSET_INSTANCE struct {
	CounterSetGuid     types.GUID
	dwSize             types.ULONG
	InstanceId         types.ULONG
	InstanceNameOffset types.ULONG
	InstanceNameSize   types.ULONG
}

type PPERF_COUNTERSET_INSTANCE *PERF_COUNTERSET_INSTANCE

// PERF_COUNTERSET_INFO
type PERF_COUNTERSET_INFO struct {
	CounterSetGuid types.GUID
	ProviderGuid   types.GUID
	NumCounters    types.ULONG
	InstanceType   types.ULONG
}

type PPERF_COUNTERSET_INFO PERF_COUNTERSET_INFO

// PERF_PROVIDER_CONTEXT
type (
	PERFLIBREQUEST types.LPVOID
	PERF_MEM_ALLOC types.LPVOID
	PERF_MEM_FREE  types.LPVOID
)

type PERF_PROVIDER_CONTEXT struct {
	ContextSize     types.DWORD
	Reserved        types.DWORD
	ControlCallback PERFLIBREQUEST
	MemAllocRoutine PERF_MEM_ALLOC
	MemFreeRoutine  PERF_MEM_FREE
	MemContext      types.LPVOID
}
type PPERF_PROVIDER_CONTEXT *PERF_PROVIDER_CONTEXT

// ENABLE_TRACE_PARAMETERS
type PEVENT_FILTER_DESCRIPTOR types.LPVOID
type ENABLE_TRACE_PARAMETERS struct {
	Version          types.ULONG
	EnableProperty   types.ULONG
	ControlFlags     types.ULONG
	SourceId         types.GUID
	EnableFilterDesc PEVENT_FILTER_DESCRIPTOR
}
type PENABLE_TRACE_PARAMETERS *ENABLE_TRACE_PARAMETERS

// TRACE_GUID_PROPERTIES
type TRACE_GUID_PROPERTIES struct {
	Guid        types.GUID
	GuidType    types.ULONG
	LoggerId    types.ULONG
	EnableLevel types.ULONG
	EnableFlags types.ULONG
	IsEnable    types.BOOLEAN
}
type PTRACE_GUID_PROPERTIES *TRACE_GUID_PROPERTIES

// TRACE_GUID_REGISTRATION
type TRACE_GUID_REGISTRATION struct {
	Guid      types.LPCGUID
	RegHandle types.HANDLE
}
type PTRACE_GUID_REGISTRATION *TRACE_GUID_REGISTRATION

// INHERITED_FROM
type INHERITED_FROM struct {
	GenerationGap types.LONG
	AncestorName  types.LPWSTR
}
type PINHERITED_FROM *INHERITED_FROM

// FN_OBJECT_MGR_FUNCTS
type FN_OBJECT_MGR_FUNCTS struct {
	Placeholder types.ULONG
}
type PFN_OBJECT_MGR_FUNCTS *FN_OBJECT_MGR_FUNCTS

// LSA_TRANSLATED_SID
type LSA_TRANSLATED_SID struct {
	Use         types.SID_NAME_USE
	RelativeId  types.ULONG
	DomainIndex types.LONG
}
type PLSA_TRANSLATED_SID *LSA_TRANSLATED_SID

// LSA_TRANSLATED_SID2
type LSA_TRANSLATED_SID2 struct {
	Use         types.SID_NAME_USE
	Sid         types.PSID
	DomainIndex types.LONG
	Flags       types.ULONG
}
type PLSA_TRANSLATED_SID2 *LSA_TRANSLATED_SID2

// LSA_TRANSLATED_NAME
type LSA_TRANSLATED_NAME struct {
	Use         types.SID_NAME_USE
	Name        types.LSA_UNICODE_STRING
	DomainIndex types.LONG
}
type PLSA_TRANSLATED_NAME *LSA_TRANSLATED_NAME

// TRUSTED_DOMAIN_INFORMATION_EX
type TRUSTED_DOMAIN_INFORMATION_EX struct {
	Name            types.LSA_UNICODE_STRING
	FlatName        types.LSA_UNICODE_STRING
	Sid             types.PSID
	TrustDirection  types.ULONG
	TrustType       types.ULONG
	TrustAttributes types.ULONG
}
type PTRUSTED_DOMAIN_INFORMATION_EX *TRUSTED_DOMAIN_INFORMATION_EX

// LsaAuthType
type LsaAuthType types.ULONG

const (
	TRUST_AUTH_TYPE_NONE    LsaAuthType = 0
	TRUST_AUTH_TYPE_NT4OWF  LsaAuthType = 1
	TRUST_AUTH_TYPE_CLEAR   LsaAuthType = 2
	TRUST_AUTH_TYPE_VERSION LsaAuthType = 3
)

// LSA_AUTH_INFORMATION
type LSA_AUTH_INFORMATION struct {
	LastUpdateTime types.LARGE_INTEGER
	AuthType       LsaAuthType
	AuthInfoLength types.ULONG
	AuthInfo       types.PUCHAR
}
type PLSA_AUTH_INFORMATION *LSA_AUTH_INFORMATION

// TRUSTED_DOMAIN_AUTH_INFORMATION
type TRUSTED_DOMAIN_AUTH_INFORMATION struct {
	IncomingAuthInfos                         types.ULONG
	IncomingAuthenticationInformation         PLSA_AUTH_INFORMATION
	IncomingPreviousAuthenticationInformation PLSA_AUTH_INFORMATION
	OutgoingAuthInfos                         types.ULONG
	OutgoingAuthenticationInformation         PLSA_AUTH_INFORMATION
	OutgoingPreviousAuthenticationInformation PLSA_AUTH_INFORMATION
}
type PTRUSTED_DOMAIN_AUTH_INFORMATION *TRUSTED_DOMAIN_AUTH_INFORMATION

// TRACE_CONTROL_CODE
type TRACE_CONTROL_CODE types.ULONG

const (
	EVENT_TRACE_CONTROL_QUERY          TRACE_CONTROL_CODE = 0
	EVENT_TRACE_CONTROL_STOP           TRACE_CONTROL_CODE = 1
	EVENT_TRACE_CONTROL_UPDATE         TRACE_CONTROL_CODE = 2
	EVENT_TRACE_CONTROL_FLUSH          TRACE_CONTROL_CODE = 3
	EVENT_TRACE_CONTROL_INCREMENT_FILE TRACE_CONTROL_CODE = 4
)

// TRACE_MESSAGE_FLAGS
type TRACE_MESSAGE_FLAGS types.ULONG

const (
	TRACE_MESSAGE_SEQUENCE              TRACE_MESSAGE_FLAGS = 1
	TRACE_MESSAGE_GUID                  TRACE_MESSAGE_FLAGS = 2
	TRACE_MESSAGE_COMPONENTID           TRACE_MESSAGE_FLAGS = 4
	TRACE_MESSAGE_TIMESTAMP             TRACE_MESSAGE_FLAGS = 8
	TRACE_MESSAGE_PERFORMANCE_TIMESTAMP TRACE_MESSAGE_FLAGS = 16
	TRACE_MESSAGE_SYSTEMINFO            TRACE_MESSAGE_FLAGS = 32
	TRACE_MESSAGE_POINTER32             TRACE_MESSAGE_FLAGS = 0x0040
	TRACE_MESSAGE_POINTER64             TRACE_MESSAGE_FLAGS = 0x0080
)

// LSA_LOOKUP_FLAGS
type LSA_LOOKUP_FLAGS types.ULONG

const (
	LSA_LOOKUP_ISOLATED_AS_LOCAL LSA_LOOKUP_FLAGS = 0x80000000
)

// Variables
type (
	PFE_EXPORT_FUNC         types.LPVOID
	PFE_IMPORT_FUNC         types.LPVOID
	LPHANDLER_FUNCTION      types.LPVOID
	LPHANDLER_FUNCTION_EX   types.LPVOID
	TRACEHANDLE             types.ULONG64
	PTRACEHANDLE            *TRACEHANDLE
	WMIDPREQUEST            types.LPVOID
	ENABLECALLBACK          types.LPVOID
	FN_PROGRESS             types.LPVOID
	LSA_HANDLE              types.LPVOID
	PLSA_HANDLE             *LSA_HANDLE
	LSA_ENUMERATION_HANDLE  types.ULONG
	PLSA_ENUMERATION_HANDLE *LSA_ENUMERATION_HANDLE
	HWCT                    types.LPVOID
	PWAITCHAINCALLBACK      types.LPVOID
	PCOGETCALLSTATE         types.LPVOID
	PCOGETACTIVATIONSTATE   types.LPVOID
	REGHANDLE               types.ULONGULONG
	PREGHANDLE              *REGHANDLE
	PLSA_OBJECT_ATTRIBUTES  *types.OBJECT_ATTRIBUTES
)

// SHUTDOWN_FLAGS
type SHUTDOWN_FLAGS types.DWORD

const (
	SHUTDOWN_FORCE_OTHERS         SHUTDOWN_FLAGS = 0x00000001
	SHUTDOWN_FORCE_SELF           SHUTDOWN_FLAGS = 0x00000002
	SHUTDOWN_RESTART              SHUTDOWN_FLAGS = 0x00000004
	SHUTDOWN_POWEROFF             SHUTDOWN_FLAGS = 0x00000008
	SHUTDOWN_NOREBOOT             SHUTDOWN_FLAGS = 0x00000010
	SHUTDOWN_GRACE_OVERRIDE       SHUTDOWN_FLAGS = 0x00000020
	SHUTDOWN_INSTALL_UPDATES      SHUTDOWN_FLAGS = 0x00000040
	SHUTDOWN_RESTARTAPPS          SHUTDOWN_FLAGS = 0x00000080
	SHUTDOWN_SKIP_SVC_PRESHUTDOWN SHUTDOWN_FLAGS = 0x00000100
	SHUTDOWN_HYBRID               SHUTDOWN_FLAGS = 0x00000200
	SHUTDOWN_RESTART_BOOTOPTIONS  SHUTDOWN_FLAGS = 0x00000400
)

// TOKEN_ACCESS_MASK_DWORD
type TOKEN_ACCESS_MASK_DWORD types.DWORD

// EVENT_INFO_CLASS
type EVENT_INFO_CLASS types.UINT

const (
	EventProviderBinaryTrackInfo EVENT_INFO_CLASS = 0
)

// CENTRAL_ACCESS_POLICY_ENTRY
type CENTRAL_ACCESS_POLICY_ENTRY struct {
	Name            types.LSA_UNICODE_STRING
	Description     types.LSA_UNICODE_STRING
	ChangeId        types.LSA_UNICODE_STRING
	LengthAppliesTo types.ULONG
	AppliesTo       types.PUCHAR
	LengthSD        types.ULONG
	SD              types.PSECURITY_DESCRIPTOR
	LengthStagedSD  types.ULONG
	StagedSD        types.PSECURITY_DESCRIPTOR
	Flags           types.ULONG
}
type PCENTRAL_ACCESS_POLICY_ENTRY *CENTRAL_ACCESS_POLICY_ENTRY

// CENTRAL_ACCESS_POLICY
type CENTRAL_ACCESS_POLICY struct {
	CAPID       types.PSID
	Name        types.LSA_UNICODE_STRING
	Description types.LSA_UNICODE_STRING
	ChangeId    types.LSA_UNICODE_STRING
	Flags       types.ULONG
	CAPECount   types.ULONG
	CAPEs       *PCENTRAL_ACCESS_POLICY_ENTRY
}
type PCENTRAL_ACCESS_POLICY *CENTRAL_ACCESS_POLICY

type RegLoadAppKey_Options types.DWORD

const (
	REG_PROCESS_APPKEY RegLoadAppKey_Options = 0x00000001
)

type RTL_ENCRYPT_OPTION_FLAGS types.ULONG

const (
	RTL_ENCRYPT_OPTION_CROSS_PROCESS RTL_ENCRYPT_OPTION_FLAGS = 0x01
	RTL_ENCRYPT_OPTION_SAME_LOGON    RTL_ENCRYPT_OPTION_FLAGS = 0x02
)
