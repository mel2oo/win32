//go:build windows
// +build windows

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/

package evntrace

import (
	"syscall"
	"unsafe"

	"github.com/mel2oo/win32/advapi32"
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

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-closetrace
func CloseTrace(TraceHandle advapi32.TRACEHANDLE) types.ULONG {
	ret, _, _ := advapi32.ProcCloseTrace.Call(uintptr(TraceHandle))
	return types.ULONG(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-controltracew
func ControlTrace(TraceHandle advapi32.TRACEHANDLE, InstanceName string,
	Properties *types.EventTraceProperties, ControlCode types.ULONG) types.ULONG {
	a1, err := syscall.UTF16PtrFromString(InstanceName)
	if err != nil {
		return 0
	}

	ret, _, _ := advapi32.ProcControlTrace.Call(
		uintptr(TraceHandle),
		uintptr(unsafe.Pointer(a1)),
		uintptr(unsafe.Pointer(Properties)),
		uintptr(ControlCode),
	)

	return types.ULONG(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-opentracew
func OpenTrace(Logfile types.PEVENT_TRACE_LOGFILE) advapi32.TRACEHANDLE {
	ret, _, _ := advapi32.ProcOpenTrace.Call(
		uintptr(unsafe.Pointer(Logfile)),
	)

	return advapi32.TRACEHANDLE(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-processtrace
func ProcessTrace(HandleArray advapi32.PTRACEHANDLE, HandleCount types.ULONG,
	StartTime types.LPFILETIME, EndTime types.LPFILETIME) types.ULONG {
	ret, _, _ := advapi32.ProcProcessTrace.Call(
		uintptr(unsafe.Pointer(HandleArray)),
		uintptr(HandleCount),
		uintptr(unsafe.Pointer(StartTime)),
		uintptr(unsafe.Pointer(EndTime)),
	)

	return types.ULONG(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-starttracew
func StartTrace(TraceHandle advapi32.PTRACEHANDLE, InstanceName string, Properties *types.EventTraceProperties) types.ULONG {
	a1, err := syscall.UTF16PtrFromString(InstanceName)
	if err != nil {
		return 0
	}

	ret, _, _ := advapi32.ProcStartTrace.Call(
		uintptr(unsafe.Pointer(TraceHandle)),
		uintptr(unsafe.Pointer(a1)),
		uintptr(unsafe.Pointer(Properties)),
	)

	return types.ULONG(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-tracesetinformation
func TraceSetInformation(SessionHandle advapi32.TRACEHANDLE, InformationClass advapi32.TRACE_INFO_CLASS,
	TraceInformation uint32, InformationLength types.ULONG) types.ULONG {
	ret, _, _ := advapi32.ProcTraceSetInformation.Call(
		uintptr(SessionHandle),
		uintptr(InformationClass),
		uintptr(unsafe.Pointer(&TraceInformation)),
		uintptr(InformationClass),
	)

	return types.ULONG(ret)
}
