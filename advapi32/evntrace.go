//go:build windows
// +build windows

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/

package advapi32

import (
	"syscall"
	"unsafe"

	"github.com/mel2oo/win32/typedef"
)

type (
	TRACEHANDLE        typedef.ULONG64
	WNODE_HEADER_Flags typedef.ULONG
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
	BufferSize    typedef.ULONG
	ProviderId    typedef.ULONG
	U1            typedef.ULONG64
	U2            typedef.HANDLE
	Guid          typedef.GUID
	ClientContext typedef.ULONG
	Flags         WNODE_HEADER_Flags
}

type EventLogFileMode typedef.ULONG

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

type EventEnableFlags typedef.ULONG

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
	BufferSize          typedef.ULONG
	MinimumBuffers      typedef.ULONG
	MaximumBuffers      typedef.ULONG
	MaximumFileSize     typedef.ULONG
	LogFileMode         EventLogFileMode
	FlushTimer          typedef.ULONG
	EnableFlags         EventEnableFlags
	AgeLimit            typedef.LONG
	NumberOfBuffers     typedef.ULONG
	FreeBuffers         typedef.ULONG
	EventsLost          typedef.ULONG
	BuffersWritten      typedef.ULONG
	LogBuffersLost      typedef.ULONG
	RealTimeBuffersLost typedef.ULONG
	LoggerThreadId      typedef.HANDLE
	LogFileNameOffset   typedef.ULONG
	LoggerNameOffset    typedef.ULONG
}
type PEVENT_TRACE_PROPERTIES *EVENT_TRACE_PROPERTIES

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-closetrace
func CloseTrace(TraceHandle TRACEHANDLE) typedef.ULONG {
	ret, _, _ := procCloseTrace.Call(uintptr(TraceHandle))
	return typedef.ULONG(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-controltracew
func ControlTrace(TraceHandle TRACEHANDLE, InstanceName string,
	Properties PEVENT_TRACE_PROPERTIES, ControlCode typedef.ULONG) typedef.ULONG {
	a1, err := syscall.UTF16PtrFromString(InstanceName)
	if err != nil {
		return 0
	}

	ret, _, _ := procControlTrace.Call(
		uintptr(TraceHandle),
		uintptr(unsafe.Pointer(a1)),
		uintptr(unsafe.Pointer(Properties)),
		uintptr(ControlCode),
	)

	return typedef.ULONG(ret)
}
