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

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-closetrace
func CloseTrace(TraceHandle advapi32.TRACEHANDLE) types.ULONG {
	ret, _, _ := advapi32.ProcCloseTrace.Call(uintptr(TraceHandle))
	return types.ULONG(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-controltracew
func ControlTrace(TraceHandle advapi32.TRACEHANDLE, InstanceName string,
	Properties types.PEVENT_TRACE_PROPERTIES, ControlCode types.ULONG) types.ULONG {
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
func StartTrace(TraceHandle advapi32.PTRACEHANDLE, InstanceName string, Properties types.PEVENT_TRACE_PROPERTIES) types.ULONG {
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
