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
func CloseTrace(TraceHandle TRACEHANDLE) types.ULONG {
	ret, _, _ := advapi32.ProcCloseTrace.Call(uintptr(TraceHandle))
	return types.ULONG(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-controltracew
func ControlTrace(TraceHandle TRACEHANDLE, InstanceName string,
	Properties PEVENT_TRACE_PROPERTIES, ControlCode types.ULONG) types.ULONG {
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
func OpenTrace(Logfile PEVENT_TRACE_LOGFILE) TRACEHANDLE {
	ret, _, _ := advapi32.ProcOpenTrace.Call(
		uintptr(unsafe.Pointer(Logfile)),
	)

	return TRACEHANDLE(ret)
}

// func ProcessTrace(
// 	HandleArray PTRACEHANDLE,
// 	HandleCount types.ULONG,
// 	StartTime LPFILETIME,
// 	EndTime LPFILETIME,
// ) types.ULONG

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-starttracew
func StartTrace(TraceHandle PTRACEHANDLE, InstanceName string, Properties PEVENT_TRACE_PROPERTIES) types.ULONG {
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

func TraceSetInformation()
