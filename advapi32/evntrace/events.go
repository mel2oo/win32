package evntrace

import (
	"syscall"

	"github.com/mel2oo/win32/types"
)

var (
	// EventTraceGuid is used to identify a event tracing session
	EventTraceGuid = syscall.GUID{
		Data1: 0x68fdd900,
		Data2: 0x4a3e,
		Data3: 0x11d1,
		Data4: [8]byte{0x84, 0xf4, 0x00, 0x00, 0xf8, 0x04, 0x64, 0xe3},
	}

	// SystemTraceControlGuid. Used to specify event tracing for kernel
	SystemTraceControlGuid = syscall.GUID{
		Data1: 0x9e814aad,
		Data2: 0x3204,
		Data3: 0x11d2,
		Data4: [8]byte{0x9a, 0x82, 0x00, 0x60, 0x08, 0xa8, 0x69, 0x39},
	}

	// EventTraceConfigGuid. Used to report system configuration records
	EventTraceConfigGuid = syscall.GUID{
		Data1: 0x01853a65,
		Data2: 0x418f,
		Data3: 0x4f36,
		Data4: [8]byte{0xae, 0xfc, 0xdc, 0x0f, 0x1d, 0x2f, 0xd2, 0x35},
	}

	// DefaultTraceSecurityGuid. Specifies the default event tracing security
	DefaultTraceSecurityGuid = syscall.GUID{
		Data1: 0x0811c1af,
		Data2: 0x7a07,
		Data3: 0x4a06,
		Data4: [8]byte{0x82, 0xed, 0x86, 0x94, 0x55, 0xcd, 0xf7, 0x13},
	}

	// PrivateLoggerNotificationGuid. Used for private cross-process logger notifications.
	PrivateLoggerNotificationGuid = syscall.GUID{
		Data1: 0x3595ab5c,
		Data2: 0x042a,
		Data3: 0x4c8e,
		Data4: [8]byte{0xb9, 0x42, 0x2d, 0x05, 0x9b, 0xfe, 0xb1, 0xb1},
	}
)

const (
	KernelLoggerName = "NT Kernel Logger"
	GlobalLoggerName = "GlobalLogger"
	EventLoggerName  = "EventLog"
	DiagLoggerName   = "DiagLog"
)

// Variables
type (
	TraceHandle              types.ULONG64
	PEVENT_CALLBACK          types.LPVOID
	EventTraceBufferCallback types.LPVOID
	PEVENT_RECORD_CALLBACK   types.LPVOID
)

// WNODE_HEADER_u1_s
type WNODE_HEADER_u1_s struct {
	Version types.ULONG
	Linkage types.ULONG
}

// WNODE_HEADER_u1
type WNODE_HEADER_u1 struct {
	HistoricalContext types.ULONG64
}

// WNODE_HEADER_u2
type WNODE_HEADER_u2 struct {
	CountLost types.ULONG
}

// WnodeHeaderFlags
type WnodeHeaderFlags types.ULONG

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
	BufferSize        types.ULONG
	ProviderId        types.ULONG
	HistoricalContext types.ULONG64
	KernelHandle      types.HANDLE
	Guid              types.GUID
	ClientContext     types.ULONG
	Flags             WnodeHeaderFlags
}

// EventLogFileMode
type EventLogFileMode types.ULONG

const (
	EventTraceFileModeNone            EventLogFileMode = 0x00000000
	EventTraceFileModeSequential      EventLogFileMode = 0x00000001
	EventTraceFileModeCircular        EventLogFileMode = 0x00000002
	EventTraceFileModeAppend          EventLogFileMode = 0x00000004
	EventTraceRealTimeMode            EventLogFileMode = 0x00000100
	EventTraceDelayOpenFileMode       EventLogFileMode = 0x00000200
	EventTraceBufferingMode           EventLogFileMode = 0x00000400
	EventTracePrivateLoggerMode       EventLogFileMode = 0x00000800
	EventTraceAddHeaderMode           EventLogFileMode = 0x00001000
	EventTraceUseGlobalSequence       EventLogFileMode = 0x00004000
	EventTraceUseLocalSequence        EventLogFileMode = 0x00008000
	EventTraceRelogMode               EventLogFileMode = 0x00010000
	EventTraceUsePagedMemory          EventLogFileMode = 0x01000000
	EventTraceFileModeNewFile         EventLogFileMode = 0x00000008
	EventTraceFileModePreAllocate     EventLogFileMode = 0x00000020
	EventTraceNonStoppableMode        EventLogFileMode = 0x00000040
	EventTraceSecureMode              EventLogFileMode = 0x00000080
	EventTraceUseKbytesForSize        EventLogFileMode = 0x00002000
	EventTracePrivateInProc           EventLogFileMode = 0x00020000
	EventTraceModeReserved            EventLogFileMode = 0x00100000
	EventTraceNoPerProcessorBuffering EventLogFileMode = 0x10000000
)

// EventEnableFlags
type EventEnableFlags types.ULONG

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
	EventTraceFlagHandle           EventEnableFlags = 0x80000040
)

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/ns-evntrace-event_trace_properties
type EventTraceProperties struct {
	Wnode               WnodeHeader
	BufferSize          types.ULONG
	MinimumBuffers      types.ULONG
	MaximumBuffers      types.ULONG
	MaximumFileSize     types.ULONG
	LogFileMode         EventLogFileMode
	FlushTimer          types.ULONG
	EnableFlags         EventEnableFlags
	AgeLimit            types.LONG
	NumberOfBuffers     types.ULONG
	FreeBuffers         types.ULONG
	EventsLost          types.ULONG
	BuffersWritten      types.ULONG
	LogBuffersLost      types.ULONG
	RealTimeBuffersLost types.ULONG
	LoggerThreadId      types.HANDLE
	LogFileNameOffset   types.ULONG
	LoggerNameOffset    types.ULONG
}

// TraceControlCode
type TraceControlCode types.ULONG

const (
	EventTraceControlQuery         TraceControlCode = 0
	EventTraceControlStop          TraceControlCode = 1
	EventTraceControlUpdate        TraceControlCode = 2
	EventTraceControlFlush         TraceControlCode = 3
	EventTraceControlIncrementFile TraceControlCode = 4
)

// TraceMessageFlags
type TraceMessageFlags types.ULONG

const (
	TraceMessageSequence             TraceMessageFlags = 1
	TraceMessageGuid                 TraceMessageFlags = 2
	TraceMessageComponentid          TraceMessageFlags = 4
	TraceMessageTimestamp            TraceMessageFlags = 8
	TraceMessagePerformanceTimestamp TraceMessageFlags = 16
	TraceMessageSysteminfo           TraceMessageFlags = 32
	TraceMessagePointer32            TraceMessageFlags = 0x0040
	TraceMessagePointer64            TraceMessageFlags = 0x0080
)

// EVENT_TRACE_HEADER_u1_s
type EVENT_TRACE_HEADER_u1_s struct {
	HeaderType  types.UCHAR
	MarkerFlags types.UCHAR
}

// EventTraceType
type EventTraceType types.UCHAR

const (
	EventTraceTypeInfo          EventTraceType = 0x00
	EventTraceTypeStart         EventTraceType = 0x01
	EventTraceTypeEnd           EventTraceType = 0x02
	EventTraceTypeStop          EventTraceType = 0x02
	EventTraceTypeDCStart       EventTraceType = 0x03
	EventTraceTypeDCEnd         EventTraceType = 0x04
	EventTraceTypeExtension     EventTraceType = 0x05
	EventTraceTypeReply         EventTraceType = 0x06
	EventTraceTypeDequeue       EventTraceType = 0x07
	EventTraceTypeResume        EventTraceType = 0x07
	EventTraceTypeCheckPoint    EventTraceType = 0x08
	EventTraceTypeSuspend       EventTraceType = 0x08
	EventTraceTypeWinevtSend    EventTraceType = 0x09
	EventTraceTypeWinevtReceive EventTraceType = 0xF0
)

// TraceLevel
type TraceLevel types.UCHAR

const (
	TraceLevelNone        TraceLevel = 0
	TraceLevelFatal       TraceLevel = 1
	TraceLevelErrot       TraceLevel = 2
	TraceLevelWarning     TraceLevel = 3
	TraceLevelInformation TraceLevel = 4
	TraceLevelVerbose     TraceLevel = 5
	TraceLevelReserved6   TraceLevel = 6
	TraceLevelReserved7   TraceLevel = 7
	TraceLevelReserved8   TraceLevel = 8
	TraceLevelReserved9   TraceLevel = 9
)

// EVENT_TRACE_HEADER_u2_s
type EVENT_TRACE_HEADER_u2_s struct {
	Type    EventTraceType
	Level   TraceLevel
	Version types.USHORT
}

// EVENT_TRACE_HEADER_u4_s1
type EVENT_TRACE_HEADER_u4_s1 struct {
	KernelTime types.ULONG
	UserTime   types.ULONG
}

// EVENT_TRACE_HEADER_u4_s2
type EVENT_TRACE_HEADER_u4_s2 struct {
	ClientContext types.ULONG
	Flags         WnodeHeaderFlags
}

// EventTraceHeader
type EventTraceHeader struct {
	Size           types.USHORT
	FieldTypeFlags types.USHORT
	Version        types.ULONG
	ThreadId       types.ULONG
	ProcessId      types.ULONG
	TimeStamp      types.LARGE_INTEGER
	Guid           types.GUID
	ProcessorTime  types.ULONG64
}

// EVENT_INSTANCE_HEADER_u1_s
type EVENT_INSTANCE_HEADER_u1_s struct {
	EventId types.ULONG
	Flags   WnodeHeaderFlags
}

// EVENT_INSTANCE_HEADER_u
type EVENT_INSTANCE_HEADER_u struct {
	ProcessorTime types.ULONG64
}

// EVENT_INSTANCE_HEADER
type EVENT_INSTANCE_HEADER struct {
	Size             types.USHORT
	FieldTypeFlags   types.USHORT
	Version          types.ULONG
	ThreadId         types.ULONG
	ProcessId        types.ULONG
	TimeStamp        types.LARGE_INTEGER
	RegHandle        types.ULONGLONG
	InstanceId       types.ULONG
	ParentInstanceId types.ULONG
	U                EVENT_INSTANCE_HEADER_u
	ParentRegHandle  types.ULONGLONG
}
type PEVENT_INSTANCE_HEADER *EVENT_INSTANCE_HEADER

// ProcessTraceMode
type ProcessTraceMode types.ULONG

const (
	ProcessTraceModeRealTime     ProcessTraceMode = 0x00000100
	ProcessTraceModeRawTimestamp ProcessTraceMode = 0x00001000
	ProcessTraceModeEventRecord  ProcessTraceMode = 0x10000000
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

// EtwBufferContext
type EtwBufferContext struct {
	ProcessorNumber types.UCHAR
	Alignment       types.UCHAR
	LoggerId        types.USHORT
}

// EventTrace
type EventTrace struct {
	Header           EventTraceHeader
	InstanceId       types.ULONG
	ParentInstanceId types.ULONG
	ParentGuid       types.GUID
	MofData          types.PVOID
	MofLength        types.ULONG
	ClientContext    types.ULONG
}

// TRACE_LOGFILE_HEADER_u1_s
type TRACE_LOGFILE_HEADER_u1_s struct {
	MajorVersion    types.UCHAR
	MinorVersion    types.UCHAR
	SubVersion      types.UCHAR
	SubMinorVersion types.UCHAR
}

// TraceLogFileHeader
type TraceLogFileHeader struct {
	BufferSize         types.ULONG
	Version            types.ULONG
	ProviderVersion    types.ULONG
	NumberOfProcessors types.ULONG
	EndTime            types.LARGE_INTEGER
	TimerResolution    types.ULONG
	MaximumFileSize    types.ULONG
	LogFileMode        EventLogFileMode
	BuffersWritten     types.ULONG
	LogInstanceGuid    types.GUID
	LoggerName         types.LPWSTR
	LogFileName        types.LPWSTR
	TimeZone           types.TIME_ZONE_INFORMATION
	BootTime           types.LARGE_INTEGER
	PerfFreq           types.LARGE_INTEGER
	StartTime          types.LARGE_INTEGER
	ReservedFlags      types.ULONG
	BuffersLost        types.ULONG
}

// EventTraceLogFile
type EventTraceLogFile struct {
	LogFileName    types.LPWSTR
	LoggerName     types.LPWSTR
	CurrentTime    types.LONGLONG
	BuffersRead    types.ULONG
	LogFileMode    EventLogFileMode
	CurrentEvent   EventTrace
	LogfileHeader  TraceLogFileHeader
	BufferCallback uintptr
	BufferSize     types.ULONG
	Filled         types.ULONG
	EventsLost     types.ULONG
	EventCallback  uintptr
	IsKernelTrace  types.ULONG
	Context        types.PVOID
}

// EVENT_INSTANCE_INFO
type EVENT_INSTANCE_INFO struct {
	RegHandle  types.HANDLE
	InstanceId types.ULONG
}

type PEVENT_INSTANCE_INFO *EVENT_INSTANCE_INFO

// EventDescriptor
type EventDescriptor struct {
	Id      types.USHORT
	Version types.UCHAR
	Channel types.UCHAR
	Level   types.UCHAR
	Opcode  types.UCHAR
	Task    types.USHORT
	Keyword types.ULONGLONG
}

// EventDataDescriptor
type EventDataDescriptor struct {
	Ptr      types.ULONGLONG
	Size     types.ULONG
	Reserved types.ULONG
}

// EventLogType
type EventLogType types.WORD

const (
	EVENTLOG_SUCCESS          EventLogType = 0x0000
	EVENTLOG_ERROR_TYPE       EventLogType = 0x0001
	EVENTLOG_WARNING_TYPE     EventLogType = 0x0002
	EVENTLOG_INFORMATION_TYPE EventLogType = 0x0004
	EVENTLOG_AUDIT_SUCCESS    EventLogType = 0x0008
	EVENTLOG_AUDIT_FAILURE    EventLogType = 0x0010
)

// EventActivity
type EventActivity types.ULONG

const (
	EVENT_ACTIVITY_CTRL_GET_ID        EventActivity = 1
	EVENT_ACTIVITY_CTRL_SET_ID        EventActivity = 2
	EVENT_ACTIVITY_CTRL_CREATE_ID     EventActivity = 3
	EVENT_ACTIVITY_CTRL_GET_SET_ID    EventActivity = 4
	EVENT_ACTIVITY_CTRL_CREATE_SET_ID EventActivity = 5
)

// TraceInfoClass
type TraceInfoClass types.UINT

const (
	TraceGuidQueryList              TraceInfoClass = 0
	TraceGuidQueryInfo              TraceInfoClass = 1
	TraceGuidQueryProcess           TraceInfoClass = 2
	TraceStackTracingInfo           TraceInfoClass = 3
	TraceSystemTraceEnableFlagsInfo TraceInfoClass = 4
	TraceSampledProfileIntervalInfo TraceInfoClass = 5
	TraceProfileSourceConfigInfo    TraceInfoClass = 6
	TraceProfileSourceListInfo      TraceInfoClass = 7
	TracePmcEventListInfo           TraceInfoClass = 8
	TracePmcCounterListInfo         TraceInfoClass = 9
)

// TraceQueryInfoClass
type TraceQueryInfoClass TraceInfoClass
