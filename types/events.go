package types

// Variables
type (
	PEVENT_CALLBACK              LPVOID
	PEVENT_TRACE_BUFFER_CALLBACK LPVOID
	PEVENT_RECORD_CALLBACK       LPVOID
)

// WNODE_HEADER_u1_s
type WNODE_HEADER_u1_s struct {
	Version ULONG
	Linkage ULONG
}

// WNODE_HEADER_u1
type WNODE_HEADER_u1 struct {
	HistoricalContext ULONG64
}

// WNODE_HEADER_u2
type WNODE_HEADER_u2 struct {
	CountLost ULONG
}

// WnodeHeaderFlags
type WnodeHeaderFlags ULONG

const (
	WnodeFlagAllData             WnodeHeaderFlags = 0x00000001
	WnodeFlagSingleInstance      WnodeHeaderFlags = 0x00000002
	WnodeFlagSingleItem          WnodeHeaderFlags = 0x00000004
	WnodeFlagEventItem           WnodeHeaderFlags = 0x00000008
	WnodeFlagFixedInstanceSize   WnodeHeaderFlags = 0x00000010
	WnodeFlagTooSmall            WnodeHeaderFlags = 0x00000020
	WnodeFlagInstancesSame       WnodeHeaderFlags = 0x00000040
	WnodeFlagStaticInstanceNames WnodeHeaderFlags = 0x00000080
	WnodeFlagInternal            WnodeHeaderFlags = 0x00000100
	WnodeFlagUseTimestamp        WnodeHeaderFlags = 0x00000200
	WnodeFlagPersistEvent        WnodeHeaderFlags = 0x00000400
	WnodeFlagEventReference      WnodeHeaderFlags = 0x00002000
	WnodeFlagAnsiInstanceNames   WnodeHeaderFlags = 0x00004000
	WnodeFlagMethodItem          WnodeHeaderFlags = 0x00008000
	WnodeFlagPDOInstanceNames    WnodeHeaderFlags = 0x00010000
	WnodeFlagTracedGUID          WnodeHeaderFlags = 0x00020000
	WnodeFlagLogWnode            WnodeHeaderFlags = 0x00040000
	WnodeFlagUseGuidPtr          WnodeHeaderFlags = 0x00080000
	WnodeFlagUseMofPtr           WnodeHeaderFlags = 0x00100000
	WnodeFlagNoHeader            WnodeHeaderFlags = 0x00200000
	WnodeFlagSendDataBlock       WnodeHeaderFlags = 0x00400000
	WnodeFlagSeverityMask        WnodeHeaderFlags = 0xff000000
)

// WnodeHeader
type WnodeHeader struct {
	BufferSize        ULONG
	ProviderId        ULONG
	HistoricalContext ULONG64
	KernelHandle      HANDLE
	Guid              GUID
	ClientContext     ULONG
	Flags             WnodeHeaderFlags
}

// EventLogFileMode
type EventLogFileMode ULONG

const (
	EVENT_TRACE_FILE_MODE_NONE             EventLogFileMode = 0x00000000
	EVENT_TRACE_FILE_MODE_SEQUENTIAL       EventLogFileMode = 0x00000001
	EVENT_TRACE_FILE_MODE_CIRCULAR         EventLogFileMode = 0x00000002
	EVENT_TRACE_FILE_MODE_APPEND           EventLogFileMode = 0x00000004
	EVENT_TRACE_REAL_TIME_MODE             EventLogFileMode = 0x00000100
	EVENT_TRACE_DELAY_OPEN_FILE_MODE       EventLogFileMode = 0x00000200
	EVENT_TRACE_BUFFERING_MODE             EventLogFileMode = 0x00000400
	EVENT_TRACE_PRIVATE_LOGGER_MODE        EventLogFileMode = 0x00000800
	EVENT_TRACE_ADD_HEADER_MODE            EventLogFileMode = 0x00001000
	EVENT_TRACE_USE_GLOBAL_SEQUENCE        EventLogFileMode = 0x00004000
	EVENT_TRACE_USE_LOCAL_SEQUENCE         EventLogFileMode = 0x00008000
	EVENT_TRACE_RELOG_MODE                 EventLogFileMode = 0x00010000
	EVENT_TRACE_USE_PAGED_MEMORY           EventLogFileMode = 0x01000000
	EVENT_TRACE_FILE_MODE_NEWFILE          EventLogFileMode = 0x00000008
	EVENT_TRACE_FILE_MODE_PREALLOCATE      EventLogFileMode = 0x00000020
	EVENT_TRACE_NONSTOPPABLE_MODE          EventLogFileMode = 0x00000040
	EVENT_TRACE_SECURE_MODE                EventLogFileMode = 0x00000080
	EVENT_TRACE_USE_KBYTES_FOR_SIZE        EventLogFileMode = 0x00002000
	EVENT_TRACE_PRIVATE_IN_PROC            EventLogFileMode = 0x00020000
	EVENT_TRACE_MODE_RESERVED              EventLogFileMode = 0x00100000
	EVENT_TRACE_NO_PER_PROCESSOR_BUFFERING EventLogFileMode = 0x10000000
)

// EventEnableFlags
type EventEnableFlags ULONG

const (
	EventTraceFlagProcess          EventEnableFlags = 0x00000001
	EventTraceFlagThread           EventEnableFlags = 0x00000002
	EventTraceFlagImageLoad        EventEnableFlags = 0x00000004
	EventTraceFlagProcessCounters  EventEnableFlags = 0x00000008
	EventTraceFlagCSwitch          EventEnableFlags = 0x00000010
	EventTraceFlagDPC              EventEnableFlags = 0x00000020
	EventTraceFlagInterrupt        EventEnableFlags = 0x00000040
	EventTraceFlagSystemcall       EventEnableFlags = 0x00000080
	EventTraceFlagDiskIO           EventEnableFlags = 0x00000100
	EventTraceFlagDiskFileIO       EventEnableFlags = 0x00000200
	EventTraceFlagDiskIOInit       EventEnableFlags = 0x00000400
	EventTraceFlagDispatcher       EventEnableFlags = 0x00000800
	EventTraceFlagMemoryPageFaults EventEnableFlags = 0x00001000
	EventTraceFlagMemoryHardFaults EventEnableFlags = 0x00002000
	EventTraceFlagVirtualAlloc     EventEnableFlags = 0x00004000
	EventTraceFlagVaMap            EventEnableFlags = 0x00008000
	EventTraceFlagNetworkTCPIP     EventEnableFlags = 0x00010000
	EventTraceFlagRegistry         EventEnableFlags = 0x00020000
	EventTraceFlagDbgPrint         EventEnableFlags = 0x00040000
	EventTraceFlagJOB              EventEnableFlags = 0x00080000
	EventTraceFlagALPC             EventEnableFlags = 0x00100000
	EventTraceFlagSplitIO          EventEnableFlags = 0x00200000
	EventTraceFlagDriver           EventEnableFlags = 0x00800000
	EventTraceFlagProfile          EventEnableFlags = 0x01000000
	EventTraceFlagFileIO           EventEnableFlags = 0x02000000
	EventTraceFlagFileIOInit       EventEnableFlags = 0x04000000
	EventTraceFlagEnableReserve    EventEnableFlags = 0x20000000
	EventTraceFlagForwardWMI       EventEnableFlags = 0x40000000
	EventTraceFlagExtension        EventEnableFlags = 0x80000000
)

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/ns-evntrace-event_trace_properties
type EventTraceProperties struct {
	Wnode               WnodeHeader
	BufferSize          ULONG
	MinimumBuffers      ULONG
	MaximumBuffers      ULONG
	MaximumFileSize     ULONG
	LogFileMode         EventLogFileMode
	FlushTimer          ULONG
	EnableFlags         EventEnableFlags
	AgeLimit            LONG
	NumberOfBuffers     ULONG
	FreeBuffers         ULONG
	EventsLost          ULONG
	BuffersWritten      ULONG
	LogBuffersLost      ULONG
	RealTimeBuffersLost ULONG
	LoggerThreadId      HANDLE
	LogFileNameOffset   ULONG
	LoggerNameOffset    ULONG
}

// EVENT_TRACE_HEADER_u1_s
type EVENT_TRACE_HEADER_u1_s struct {
	HeaderType  UCHAR
	MarkerFlags UCHAR
}

// EVENT_TRACE_HEADER_u1
type EVENT_TRACE_HEADER_u1 struct {
	FieldTypeFlags USHORT
}

// EVENT_TRACE_TYPE
type EVENT_TRACE_TYPE UCHAR

const (
	EVENT_TRACE_TYPE_INFO           EVENT_TRACE_TYPE = 0x00
	EVENT_TRACE_TYPE_START          EVENT_TRACE_TYPE = 0x01
	EVENT_TRACE_TYPE_END            EVENT_TRACE_TYPE = 0x02
	EVENT_TRACE_TYPE_STOP           EVENT_TRACE_TYPE = 0x02
	EVENT_TRACE_TYPE_DC_START       EVENT_TRACE_TYPE = 0x03
	EVENT_TRACE_TYPE_DC_END         EVENT_TRACE_TYPE = 0x04
	EVENT_TRACE_TYPE_EXTENSION      EVENT_TRACE_TYPE = 0x05
	EVENT_TRACE_TYPE_REPLY          EVENT_TRACE_TYPE = 0x06
	EVENT_TRACE_TYPE_DEQUEUE        EVENT_TRACE_TYPE = 0x07
	EVENT_TRACE_TYPE_RESUME         EVENT_TRACE_TYPE = 0x07
	EVENT_TRACE_TYPE_CHECKPOINT     EVENT_TRACE_TYPE = 0x08
	EVENT_TRACE_TYPE_SUSPEND        EVENT_TRACE_TYPE = 0x08
	EVENT_TRACE_TYPE_WINEVT_SEND    EVENT_TRACE_TYPE = 0x09
	EVENT_TRACE_TYPE_WINEVT_RECEIVE EVENT_TRACE_TYPE = 0xF0
)

// TRACE_LEVEL
type TRACE_LEVEL UCHAR

const (
	TRACE_LEVEL_NONE        TRACE_LEVEL = 0
	TRACE_LEVEL_FATAL       TRACE_LEVEL = 1
	TRACE_LEVEL_ERROR       TRACE_LEVEL = 2
	TRACE_LEVEL_WARNING     TRACE_LEVEL = 3
	TRACE_LEVEL_INFORMATION TRACE_LEVEL = 4
	TRACE_LEVEL_VERBOSE     TRACE_LEVEL = 5
	TRACE_LEVEL_RESERVED6   TRACE_LEVEL = 6
	TRACE_LEVEL_RESERVED7   TRACE_LEVEL = 7
	TRACE_LEVEL_RESERVED8   TRACE_LEVEL = 8
	TRACE_LEVEL_RESERVED9   TRACE_LEVEL = 9
)

// EVENT_TRACE_HEADER_u2_s
type EVENT_TRACE_HEADER_u2_s struct {
	Type    EVENT_TRACE_TYPE
	Level   TRACE_LEVEL
	Version USHORT
}

// EVENT_TRACE_HEADER_u2
type EVENT_TRACE_HEADER_u2 struct {
	Version ULONG
}

// EVENT_TRACE_HEADER_u3
type EVENT_TRACE_HEADER_u3 struct {
	Guid GUID
}

// EVENT_TRACE_HEADER_u4_s1
type EVENT_TRACE_HEADER_u4_s1 struct {
	KernelTime ULONG
	UserTime   ULONG
}

// EVENT_TRACE_HEADER_u4_s2
type EVENT_TRACE_HEADER_u4_s2 struct {
	ClientContext ULONG
	Flags         WnodeHeaderFlags
}

// EVENT_TRACE_HEADER_u4
type EVENT_TRACE_HEADER_u4 struct {
	ProcessorTime ULONG64
}

// EVENT_TRACE_HEADER
type EVENT_TRACE_HEADER struct {
	Size      USHORT
	U1        EVENT_TRACE_HEADER_u1
	U2        EVENT_TRACE_HEADER_u2
	ThreadId  ULONG
	ProcessId ULONG
	TimeStamp LARGE_INTEGER
	U3        EVENT_TRACE_HEADER_u3
	U4        EVENT_TRACE_HEADER_u4
}
type PEVENT_TRACE_HEADER *EVENT_TRACE_HEADER

// EVENT_INSTANCE_HEADER_u1_s
type EVENT_INSTANCE_HEADER_u1_s struct {
	EventId ULONG
	Flags   WnodeHeaderFlags
}

// EVENT_INSTANCE_HEADER_u
type EVENT_INSTANCE_HEADER_u struct {
	ProcessorTime ULONG64
}

// EVENT_INSTANCE_HEADER
type EVENT_INSTANCE_HEADER struct {
	Size             USHORT
	U1               EVENT_TRACE_HEADER_u1
	U2               EVENT_TRACE_HEADER_u2
	ThreadId         ULONG
	ProcessId        ULONG
	TimeStamp        LARGE_INTEGER
	RegHandle        ULONGLONG
	InstanceId       ULONG
	ParentInstanceId ULONG
	U                EVENT_INSTANCE_HEADER_u
	ParentRegHandle  ULONGLONG
}
type PEVENT_INSTANCE_HEADER *EVENT_INSTANCE_HEADER

// ProcessTraceMode
type ProcessTraceMode ULONG

const (
	PROCESS_TRACE_MODE_REAL_TIME     ProcessTraceMode = 0x00000100
	PROCESS_TRACE_MODE_RAW_TIMESTAMP ProcessTraceMode = 0x00001000
	PROCESS_TRACE_MODE_EVENT_RECORD  ProcessTraceMode = 0x10000000
)

// EVENT_TRACE_LOGFILE_u1
type EVENT_TRACE_LOGFILE_u1 struct {
	LogFileMode EventLogFileMode
}

// EVENT_TRACE_LOGFILE_u2
type EVENT_TRACE_LOGFILE_u2 struct {
	EventCallback PEVENT_CALLBACK
	// EventRecordCallback PEVENT_RECORD_CALLBACK
}

// ETW_BUFFER_CONTEXT
type ETW_BUFFER_CONTEXT struct {
	ProcessorNumber UCHAR
	Alignment       UCHAR
	LoggerId        USHORT
}

// EVENT_TRACE_u
type EVENT_TRACE_u struct {
	ClientContext ULONG
	// BufferContext ETW_BUFFER_CONTEXT
}

// EVENT_TRACE
type EVENT_TRACE struct {
	Header           EVENT_TRACE_HEADER
	InstanceId       ULONG
	ParentInstanceId ULONG
	ParentGuid       GUID
	MofData          PVOID
	MofLength        ULONG
	U                EVENT_TRACE_u
}

// TRACE_LOGFILE_HEADER_u1_s
type TRACE_LOGFILE_HEADER_u1_s struct {
	MajorVersion    UCHAR
	MinorVersion    UCHAR
	SubVersion      UCHAR
	SubMinorVersion UCHAR
}

// TRACE_LOGFILE_HEADER_u1
type TRACE_LOGFILE_HEADER_u1 struct {
	Version ULONG
	// VersionDetail TRACE_LOGFILE_HEADER_u1_s
}

// TRACE_LOGFILE_HEADER_u2_s
type TRACE_LOGFILE_HEADER_u2_s struct {
	StartBuffers  ULONG
	PointerSize   ULONG
	EventsLost    ULONG
	CpuSpeedInMHz ULONG
}

// TRACE_LOGFILE_HEADER_u2
type TRACE_LOGFILE_HEADER_u2 struct {
	LogInstanceGuid GUID
	//  [TRACE_LOGFILE_HEADER_u2_s]
}

// TRACE_LOGFILE_HEADER
type TRACE_LOGFILE_HEADER struct {
	BufferSize         ULONG
	U1                 TRACE_LOGFILE_HEADER_u1
	ProviderVersion    ULONG
	NumberOfProcessors ULONG
	EndTime            LARGE_INTEGER
	TimerResolution    ULONG
	MaximumFileSize    ULONG
	LogFileMode        EventLogFileMode
	BuffersWritten     ULONG
	U2                 TRACE_LOGFILE_HEADER_u2
	LoggerName         LPWSTR
	LogFileName        LPWSTR
	TimeZone           TIME_ZONE_INFORMATION
	BootTime           LARGE_INTEGER
	PerfFreq           LARGE_INTEGER
	StartTime          LARGE_INTEGER
	ReservedFlags      ULONG
	BuffersLost        ULONG
}

// EVENT_TRACE_LOGFILE
type EVENT_TRACE_LOGFILE struct {
	LogFileName    LPWSTR
	LoggerName     LPWSTR
	CurrentTime    LONGLONG
	BuffersRead    ULONG
	LogFileMode    EventLogFileMode
	CurrentEvent   EVENT_TRACE
	LogfileHeader  TRACE_LOGFILE_HEADER
	BufferCallback PEVENT_TRACE_BUFFER_CALLBACK
	BufferSize     ULONG
	Filled         ULONG
	EventsLost     ULONG
	EventCallback  uintptr
	IsKernelTrace  ULONG
	Context        PVOID
}
type PEVENT_TRACE_LOGFILE *EVENT_TRACE_LOGFILE

// EVENT_INSTANCE_INFO
type EVENT_INSTANCE_INFO struct {
	RegHandle  HANDLE
	InstanceId ULONG
}

type PEVENT_INSTANCE_INFO *EVENT_INSTANCE_INFO

// EVENT_DESCRIPTOR
type EVENT_DESCRIPTOR struct {
	Id      USHORT
	Version UCHAR
	Channel UCHAR
	Level   UCHAR
	Opcode  UCHAR
	Task    USHORT
	Keyword ULONGLONG
}

type PCEVENT_DESCRIPTOR *EVENT_DESCRIPTOR

// EVENT_DATA_DESCRIPTOR
type EVENT_DATA_DESCRIPTOR struct {
	Ptr      ULONGLONG
	Size     ULONG
	Reserved ULONG
}

type PEVENT_DATA_DESCRIPTOR *EVENT_DATA_DESCRIPTOR

// EventLogType
type EventLogType WORD

const (
	EVENTLOG_SUCCESS          EventLogType = 0x0000
	EVENTLOG_ERROR_TYPE       EventLogType = 0x0001
	EVENTLOG_WARNING_TYPE     EventLogType = 0x0002
	EVENTLOG_INFORMATION_TYPE EventLogType = 0x0004
	EVENTLOG_AUDIT_SUCCESS    EventLogType = 0x0008
	EVENTLOG_AUDIT_FAILURE    EventLogType = 0x0010
)

// EventActivity
type EventActivity ULONG

const (
	EVENT_ACTIVITY_CTRL_GET_ID        EventActivity = 1
	EVENT_ACTIVITY_CTRL_SET_ID        EventActivity = 2
	EVENT_ACTIVITY_CTRL_CREATE_ID     EventActivity = 3
	EVENT_ACTIVITY_CTRL_GET_SET_ID    EventActivity = 4
	EVENT_ACTIVITY_CTRL_CREATE_SET_ID EventActivity = 5
)
