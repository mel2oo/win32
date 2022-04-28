package evntrace

import (
	"syscall"

	"github.com/mel2oo/win32/types"
)

type (
	TRACEHANDLE        types.ULONG64
	PTRACEHANDLE       *TRACEHANDLE
	WNODE_HEADER_Flags types.ULONG
)

const (
	WNODE_FLAG_ALL_DATA              WNODE_HEADER_Flags = 0x00000001
	WNODE_FLAG_SINGLE_INSTANCE       WNODE_HEADER_Flags = 0x00000002
	WNODE_FLAG_SINGLE_ITEM           WNODE_HEADER_Flags = 0x00000004
	WNODE_FLAG_EVENT_ITEM            WNODE_HEADER_Flags = 0x00000008
	WNODE_FLAG_FIXED_INSTANCE_SIZE   WNODE_HEADER_Flags = 0x00000010
	WNODE_FLAG_TOO_SMALL             WNODE_HEADER_Flags = 0x00000020
	WNODE_FLAG_INSTANCES_SAME        WNODE_HEADER_Flags = 0x00000040
	WNODE_FLAG_STATIC_INSTANCE_NAMES WNODE_HEADER_Flags = 0x00000080
	WNODE_FLAG_INTERNAL              WNODE_HEADER_Flags = 0x00000100
	WNODE_FLAG_USE_TIMESTAMP         WNODE_HEADER_Flags = 0x00000200
	WNODE_FLAG_PERSIST_EVENT         WNODE_HEADER_Flags = 0x00000400
	WNODE_FLAG_EVENT_REFERENCE       WNODE_HEADER_Flags = 0x00002000
	WNODE_FLAG_ANSI_INSTANCENAMES    WNODE_HEADER_Flags = 0x00004000
	WNODE_FLAG_METHOD_ITEM           WNODE_HEADER_Flags = 0x00008000
	WNODE_FLAG_PDO_INSTANCE_NAMES    WNODE_HEADER_Flags = 0x00010000
	WNODE_FLAG_TRACED_GUID           WNODE_HEADER_Flags = 0x00020000
	WNODE_FLAG_LOG_WNODE             WNODE_HEADER_Flags = 0x00040000
	WNODE_FLAG_USE_GUID_PTR          WNODE_HEADER_Flags = 0x00080000
	WNODE_FLAG_USE_MOF_PTR           WNODE_HEADER_Flags = 0x00100000
	WNODE_FLAG_NO_HEADER             WNODE_HEADER_Flags = 0x00200000
	WNODE_FLAG_SEND_DATA_BLOCK       WNODE_HEADER_Flags = 0x00400000
	WNODE_FLAG_SEVERITY_MASK         WNODE_HEADER_Flags = 0xff000000
)

type WNODE_HEADER struct {
	BufferSize    types.ULONG
	ProviderId    types.ULONG
	U1            types.ULONG64
	U2            types.HANDLE
	Guid          types.GUID
	ClientContext types.ULONG
	Flags         WNODE_HEADER_Flags
}

type EventLogFileMode types.ULONG

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

type EventEnableFlags types.ULONG

const (
	EVENT_TRACE_FLAG_PROCESS            EventEnableFlags = 0x00000001
	EVENT_TRACE_FLAG_THREAD             EventEnableFlags = 0x00000002
	EVENT_TRACE_FLAG_IMAGE_LOAD         EventEnableFlags = 0x00000004
	EVENT_TRACE_FLAG_DISK_IO            EventEnableFlags = 0x00000100
	EVENT_TRACE_FLAG_DISK_FILE_IO       EventEnableFlags = 0x00000200
	EVENT_TRACE_FLAG_MEMORY_PAGE_FAULTS EventEnableFlags = 0x00001000
	EVENT_TRACE_FLAG_MEMORY_HARD_FAULTS EventEnableFlags = 0x00002000
	EVENT_TRACE_FLAG_NETWORK_TCPIP      EventEnableFlags = 0x00010000
	EVENT_TRACE_FLAG_REGISTRY           EventEnableFlags = 0x00020000
	EVENT_TRACE_FLAG_DBGPRINT           EventEnableFlags = 0x00040000
	EVENT_TRACE_FLAG_PROCESS_COUNTERS   EventEnableFlags = 0x00000008
	EVENT_TRACE_FLAG_CSWITCH            EventEnableFlags = 0x00000010
	EVENT_TRACE_FLAG_DPC                EventEnableFlags = 0x00000020
	EVENT_TRACE_FLAG_INTERRUPT          EventEnableFlags = 0x00000040
	EVENT_TRACE_FLAG_SYSTEMCALL         EventEnableFlags = 0x00000080
	EVENT_TRACE_FLAG_DISK_IO_INIT       EventEnableFlags = 0x00000400
	EVENT_TRACE_FLAG_ALPC               EventEnableFlags = 0x00100000
	EVENT_TRACE_FLAG_SPLIT_IO           EventEnableFlags = 0x00200000
	EVENT_TRACE_FLAG_DRIVER             EventEnableFlags = 0x00800000
	EVENT_TRACE_FLAG_PROFILE            EventEnableFlags = 0x01000000
	EVENT_TRACE_FLAG_FILE_IO            EventEnableFlags = 0x02000000
	EVENT_TRACE_FLAG_FILE_IO_INIT       EventEnableFlags = 0x04000000
	EVENT_TRACE_FLAG_DISPATCHER         EventEnableFlags = 0x00000800
	EVENT_TRACE_FLAG_VIRTUAL_ALLOC      EventEnableFlags = 0x00004000
	EVENT_TRACE_FLAG_EXTENSION          EventEnableFlags = 0x80000000
	EVENT_TRACE_FLAG_FORWARD_WMI        EventEnableFlags = 0x40000000
	EVENT_TRACE_FLAG_ENABLE_RESERVE     EventEnableFlags = 0x20000000
)

type EVENT_TRACE_PROPERTIES struct {
	Wnode               WNODE_HEADER
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
type PEVENT_TRACE_PROPERTIES *EVENT_TRACE_PROPERTIES

type EVENT_TRACE_HEADER struct {
	Size           types.USHORT
	FieldTypeFlags [2]byte
	Version        [4]byte
	ThreadId       types.ULONG
	ProcessId      types.ULONG
	TimeStamp      types.LARGE_INTEGER
	GUID           [16]byte
	ProcessorTime  [8]byte
}

// <Variable PEVENT_TRACE_HEADER" Type="Pointer" Base="EVENT_TRACE_HEADER

type EVENT_TRACE struct {
	Header           EVENT_TRACE_HEADER
	InstanceId       types.ULONG
	ParentInstanceId types.ULONG
	ParentGuid       types.GUID
	MofData          types.PVOID
	MofLength        types.ULONG
	Context          types.ULONG
}

type TRACE_LOGFILE_HEADER struct {
	BufferSize         types.ULONG
	Version            [4]byte
	ProviderVersion    types.ULONG
	NumberOfProcessors types.ULONG
	EndTime            types.LARGE_INTEGER
	TimerResolution    types.ULONG
	MaximumFileSize    types.ULONG
	LogFileMode        EventLogFileMode
	BuffersWritten     types.ULONG
	GUID               [16]byte
	LoggerName         string
	LogFileName        string
	TimeZone           syscall.Timezoneinformation
	BootTime           types.LARGE_INTEGER
	PerfFreq           types.LARGE_INTEGER
	StartTime          types.LARGE_INTEGER
	ReservedFlags      types.ULONG
	BuffersLost        types.ULONG
}

type PEVENT_TRACE_BUFFER_CALLBACK types.LPVOID

type EVENT_TRACE_LOGFILE struct {
	LogFileName    string
	LoggerName     string
	CurrentTime    types.LONGLONG
	BuffersRead    types.ULONG
	LogFileMode    types.ULONG
	CurrentEvent   EVENT_TRACE
	LogfileHeader  TRACE_LOGFILE_HEADER
	BufferCallback PEVENT_TRACE_BUFFER_CALLBACK
	BufferSize     types.ULONG
	Filled         types.ULONG
	EventsLost     types.ULONG
	EventCallback  [8]byte
	IsKernelTrace  types.ULONG
	Context        types.PVOID
}

type PEVENT_TRACE_LOGFILE *EVENT_TRACE_LOGFILE
