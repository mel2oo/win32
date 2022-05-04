package types

// Variables
type (
	PFN_SC_NOTIFY_CALLBACK LPVOID
	SERVICE_STATUS_HANDLE  HANDLE
	SC_HANDLE              HANDLE
	SC_LOCK                LPVOID
)

// SC_STATUS_TYPE
type SC_STATUS_TYPE UINT

const (
	SC_STATUS_PROCESS_INFO SC_STATUS_TYPE = 0
)

// ServiceType
type ServiceType DWORD

const (
	SERVICE_KERNEL_DRIVER       ServiceType = 0x00000001
	SERVICE_FILE_SYSTEM_DRIVER  ServiceType = 0x00000002
	SERVICE_ADAPTER             ServiceType = 0x00000004
	SERVICE_RECOGNIZER_DRIVER   ServiceType = 0x00000008
	SERVICE_DRIVER              ServiceType = 0x0000000B
	SERVICE_WIN32_OWN_PROCESS   ServiceType = 0x00000010
	SERVICE_WIN32_SHARE_PROCESS ServiceType = 0x00000020
	SERVICE_INTERACTIVE_PROCESS ServiceType = 0x00000100
	SERVICE_WIN32               ServiceType = 0x00000030
	SERVICE_NO_CHANGE           ServiceType = 0xffffffff
)

// ServiceState
type ServiceState DWORD

const (
	SERVICE_ACTIVE    ServiceState = 0x00000001
	SERVICE_INACTIVE  ServiceState = 0x00000002
	SERVICE_STATE_ALL ServiceState = 0x00000003
)

// ServiceCurrentState
type ServiceCurrentState DWORD

const (
	SERVICE_STOPPED          ServiceCurrentState = 0x00000001
	SERVICE_START_PENDING    ServiceCurrentState = 0x00000002
	SERVICE_STOP_PENDING     ServiceCurrentState = 0x00000003
	SERVICE_RUNNING          ServiceCurrentState = 0x00000004
	SERVICE_CONTINUE_PENDING ServiceCurrentState = 0x00000005
	SERVICE_PAUSE_PENDING    ServiceCurrentState = 0x00000006
	SERVICE_PAUSED           ServiceCurrentState = 0x00000007
)

// ServiceStartType
type ServiceStartType DWORD

const (
	SERVICE_BOOT_START   ServiceStartType = 0x00000000
	SERVICE_SYSTEM_START ServiceStartType = 0x00000001
	SERVICE_AUTO_START   ServiceStartType = 0x00000002
	SERVICE_DEMAND_START ServiceStartType = 0x00000003
	SERVICE_DISABLED     ServiceStartType = 0x00000004
)

// ServiceErrorControl
type ServiceErrorControl DWORD

const (
	SERVICE_ERROR_IGNORE   ServiceErrorControl = 0x00000000
	SERVICE_ERROR_NORMAL   ServiceErrorControl = 0x00000001
	SERVICE_ERROR_SEVERE   ServiceErrorControl = 0x00000002
	SERVICE_ERROR_CRITICAL ServiceErrorControl = 0x00000003
)

// ServiceInfoLevel
type ServiceInfoLevel DWORD

const (
	SERVICE_CONFIG_DESCRIPTION              ServiceInfoLevel = 1
	SERVICE_CONFIG_FAILURE_ACTIONS          ServiceInfoLevel = 2
	SERVICE_CONFIG_DELAYED_AUTO_START_INFO  ServiceInfoLevel = 3
	SERVICE_CONFIG_FAILURE_ACTIONS_FLAG     ServiceInfoLevel = 4
	SERVICE_CONFIG_SERVICE_SID_INFO         ServiceInfoLevel = 5
	SERVICE_CONFIG_REQUIRED_PRIVILEGES_INFO ServiceInfoLevel = 6
	SERVICE_CONFIG_PRESHUTDOWN_INFO         ServiceInfoLevel = 7
	SERVICE_CONFIG_TRIGGER_INFO             ServiceInfoLevel = 8
	SERVICE_CONFIG_PREFERRED_NODE           ServiceInfoLevel = 9
)

// SCManagerAccess
type SCManagerAccess DWORD

const (
	SC_MANAGER_GENERIC_READ       SCManagerAccess = 0x80000000
	SC_MANAGER_GENERIC_WRITE      SCManagerAccess = 0x40000000
	SC_MANAGER_GENERIC_EXECUTE    SCManagerAccess = 0x20000000
	SC_MANAGER_CONNECT            SCManagerAccess = 0x0001
	SC_MANAGER_CREATE_SERVICE     SCManagerAccess = 0x0002
	SC_MANAGER_ENUMERATE_SERVICE  SCManagerAccess = 0x0004
	SC_MANAGER_LOCK               SCManagerAccess = 0x0008
	SC_MANAGER_QUERY_LOCK_STATUS  SCManagerAccess = 0x0010
	SC_MANAGER_MODIFY_BOOT_CONFIG SCManagerAccess = 0x0020
	SC_MANAGER_ALL_ACCESS         SCManagerAccess = 0xF003F
)

// ServiceAccess
type ServiceAccess DWORD

const (
	SERVICE_QUERY_CONFIG           ServiceAccess = 0x0001
	SERVICE_CHANGE_CONFIG          ServiceAccess = 0x0002
	SERVICE_QUERY_STATUS           ServiceAccess = 0x0004
	SERVICE_ENUMERATE_DEPENDENTS   ServiceAccess = 0x0008
	SERVICE_START                  ServiceAccess = 0x0010
	SERVICE_STOP                   ServiceAccess = 0x0020
	SERVICE_PAUSE_CONTINUE         ServiceAccess = 0x0040
	SERVICE_INTERROGATE            ServiceAccess = 0x0080
	SERVICE_USER_DEFINED_CONTROL   ServiceAccess = 0x0100
	SERVICE_ALL_ACCESS             ServiceAccess = 0xF01FF
	SERVICE_ACCESS_SYSTEM_SECURITY ServiceAccess = 0x01000000
	SERVICE_DELETE                 ServiceAccess = 0x00010000
	SERVICE_READ_CONTROL           ServiceAccess = 0x00020000
	SERVICE_WRITE_DAC              ServiceAccess = 0x00040000
	SERVICE_WRITE_OWNER            ServiceAccess = 0x00080000
	SERVICE_GENERIC_READ           ServiceAccess = 0x80000000
	SERVICE_GENERIC_WRITE          ServiceAccess = 0x40000000
	SERVICE_GENERIC_EXECUTE        ServiceAccess = 0x20000000
	SERVICE_MAXIMUM_ALLOWED        ServiceAccess = 0x02000000
)

// ServiceControl
type ServiceControl DWORD

const (
	SERVICE_CONTROL_STOP                  ServiceControl = 0x00000001
	SERVICE_CONTROL_PAUSE                 ServiceControl = 0x00000002
	SERVICE_CONTROL_CONTINUE              ServiceControl = 0x00000003
	SERVICE_CONTROL_INTERROGATE           ServiceControl = 0x00000004
	SERVICE_CONTROL_SHUTDOWN              ServiceControl = 0x00000005
	SERVICE_CONTROL_PARAMCHANGE           ServiceControl = 0x00000006
	SERVICE_CONTROL_NETBINDADD            ServiceControl = 0x00000007
	SERVICE_CONTROL_NETBINDREMOVE         ServiceControl = 0x00000008
	SERVICE_CONTROL_NETBINDENABLE         ServiceControl = 0x00000009
	SERVICE_CONTROL_NETBINDDISABLE        ServiceControl = 0x0000000A
	SERVICE_CONTROL_DEVICEEVENT           ServiceControl = 0x0000000B
	SERVICE_CONTROL_HARDWAREPROFILECHANGE ServiceControl = 0x0000000C
	SERVICE_CONTROL_POWEREVENT            ServiceControl = 0x0000000D
	SERVICE_CONTROL_SESSIONCHANGE         ServiceControl = 0x0000000E
	SERVICE_CONTROL_PRESHUTDOWN           ServiceControl = 0x0000000F
)

// ServiceAcceptControls
type ServiceAcceptControls DWORD

const (
	SERVICE_ACCEPT_STOP                  ServiceAcceptControls = 0x00000001
	SERVICE_ACCEPT_PAUSE_CONTINUE        ServiceAcceptControls = 0x00000002
	SERVICE_ACCEPT_SHUTDOWN              ServiceAcceptControls = 0x00000004
	SERVICE_ACCEPT_PARAMCHANGE           ServiceAcceptControls = 0x00000008
	SERVICE_ACCEPT_NETBINDCHANGE         ServiceAcceptControls = 0x00000010
	SERVICE_ACCEPT_HARDWAREPROFILECHANGE ServiceAcceptControls = 0x00000020
	SERVICE_ACCEPT_POWEREVENT            ServiceAcceptControls = 0x00000040
	SERVICE_ACCEPT_SESSIONCHANGE         ServiceAcceptControls = 0x00000080
	SERVICE_ACCEPT_PRESHUTDOWN           ServiceAcceptControls = 0x00000100
	SERVICE_ACCEPT_TIMECHANGE            ServiceAcceptControls = 0x00000200
	SERVICE_ACCEPT_TRIGGEREVENT          ServiceAcceptControls = 0x00000400
)

// ServiceNotifyMask
type ServiceNotifyMask DWORD

const (
	SERVICE_NOTIFY_STOPPED          ServiceNotifyMask = 0x00000001
	SERVICE_NOTIFY_START_PENDING    ServiceNotifyMask = 0x00000002
	SERVICE_NOTIFY_STOP_PENDING     ServiceNotifyMask = 0x00000004
	SERVICE_NOTIFY_RUNNING          ServiceNotifyMask = 0x00000008
	SERVICE_NOTIFY_CONTINUE_PENDING ServiceNotifyMask = 0x00000010
	SERVICE_NOTIFY_PAUSE_PENDING    ServiceNotifyMask = 0x00000020
	SERVICE_NOTIFY_PAUSED           ServiceNotifyMask = 0x00000040
	SERVICE_NOTIFY_CREATED          ServiceNotifyMask = 0x00000080
	SERVICE_NOTIFY_DELETED          ServiceNotifyMask = 0x00000100
	SERVICE_NOTIFY_DELETE_PENDING   ServiceNotifyMask = 0x00000200
)

// ServiceFlags
type ServiceFlags DWORD

const (
	SERVICE_RUNS_IN_SYSTEM_PROCESS ServiceFlags = 0x00000001
)

// SERVICE_STATUS_PROCESS
type SERVICE_STATUS_PROCESS struct {
	ServiceType             ServiceType
	CurrentState            ServiceCurrentState
	ControlsAccepted        ServiceAcceptControls
	Win32ExitCode           DWORD
	ServiceSpecificExitCode DWORD
	CheckPoint              DWORD
	WaitHint                DWORD
	ProcessId               DWORD
	ServiceFlags            ServiceFlags
}

// SERVICE_NOTIFY
type SERVICE_NOTIFY struct {
	Version               DWORD
	NotifyCallback        PFN_SC_NOTIFY_CALLBACK
	Context               PVOID
	NotificationStatus    DWORD
	ServiceStatus         SERVICE_STATUS_PROCESS
	NotificationTriggered DWORD
	ServiceNames          LPWSTR
}

type PSERVICE_NOTIFY *SERVICE_NOTIFY

// SERVICE_NOTIFYA
type SERVICE_NOTIFYA struct {
	Version               DWORD
	NotifyCallback        PFN_SC_NOTIFY_CALLBACK
	Context               PVOID
	NotificationStatus    DWORD
	ServiceStatus         SERVICE_STATUS_PROCESS
	NotificationTriggered DWORD
	ServiceNames          LPSTR
}

type PSERVICE_NOTIFYA *SERVICE_NOTIFYA

// SERVICE_STATUS
type SERVICE_STATUS struct {
	ServiceType             ServiceType
	CurrentState            ServiceCurrentState
	ControlsAccepted        ServiceAcceptControls
	Win32ExitCode           DWORD
	ServiceSpecificExitCode DWORD
	CheckPoint              DWORD
	WaitHint                DWORD
}

type LPSERVICE_STATUS *SERVICE_STATUS

// SERVICE_TABLE_ENTRY
type LPSERVICE_MAIN_FUNCTION LPVOID

type SERVICE_TABLE_ENTRY struct {
	ServiceName LPWSTR
	ServiceProc LPSERVICE_MAIN_FUNCTION
}

// ENUM_SERVICE_STATUS
type ENUM_SERVICE_STATUS struct {
	ServiceName   LPWSTR
	DisplayName   LPWSTR
	ServiceStatus SERVICE_STATUS
}

type LPENUM_SERVICE_STATUS *ENUM_SERVICE_STATUS

// QUERY_SERVICE_CONFIG
type QUERY_SERVICE_CONFIG struct {
	ServiceType      ServiceType
	StartType        ServiceStartType
	ErrorControl     ServiceErrorControl
	BinaryPathName   LPWSTR
	LoadOrderGroup   LPWSTR
	TagId            DWORD
	Dependencies     LPWSTR
	ServiceStartName LPWSTR
	DisplayName      LPWSTR
}

type LPQUERY_SERVICE_CONFIG *QUERY_SERVICE_CONFIG

// QUERY_SERVICE_LOCK_STATUS
type QUERY_SERVICE_LOCK_STATUS struct {
	IsLocked     DWORD
	LockOwner    LPWSTR
	LockDuration DWORD
}

type LPQUERY_SERVICE_LOCK_STATUS *QUERY_SERVICE_LOCK_STATUS
