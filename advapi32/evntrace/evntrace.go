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
func CloseTrace(handle TraceHandle) types.ULONG {
	ret, _, _ := advapi32.ProcCloseTrace.Call(uintptr(handle))
	return types.ULONG(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-controltracew
func ControlTrace(handle TraceHandle, InstanceName string,
	Properties *EventTraceProperties, ControlCode types.ULONG) types.ULONG {
	a1, err := syscall.UTF16PtrFromString(InstanceName)
	if err != nil {
		return 0
	}

	ret, _, _ := advapi32.ProcControlTrace.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(a1)),
		uintptr(unsafe.Pointer(Properties)),
		uintptr(ControlCode),
	)

	return types.ULONG(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-opentracew
func OpenTrace(Logfile *EventTraceLogFile) TraceHandle {
	ret, _, _ := advapi32.ProcOpenTrace.Call(
		uintptr(unsafe.Pointer(Logfile)),
	)

	return TraceHandle(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-processtrace
func ProcessTrace(handle *TraceHandle, HandleCount types.ULONG,
	StartTime types.LPFILETIME, EndTime types.LPFILETIME) types.ULONG {
	ret, _, _ := advapi32.ProcProcessTrace.Call(
		uintptr(unsafe.Pointer(handle)),
		uintptr(HandleCount),
		uintptr(unsafe.Pointer(StartTime)),
		uintptr(unsafe.Pointer(EndTime)),
	)

	return types.ULONG(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-starttracew
func StartTrace(handle *TraceHandle, InstanceName string, Properties *EventTraceProperties) types.ULONG {
	a1, err := syscall.UTF16PtrFromString(InstanceName)
	if err != nil {
		return 0
	}

	ret, _, _ := advapi32.ProcStartTrace.Call(
		uintptr(unsafe.Pointer(handle)),
		uintptr(unsafe.Pointer(a1)),
		uintptr(unsafe.Pointer(Properties)),
	)

	return types.ULONG(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-tracesetinformation
func TraceSetInformation(handle TraceHandle, InformationClass TraceInfoClass,
	TraceInformation []uint32, InformationLength types.ULONG) types.ULONG {
	ret, _, _ := advapi32.ProcTraceSetInformation.Call(
		uintptr(handle),
		uintptr(InformationClass),
		uintptr(unsafe.Pointer(&TraceInformation[0])),
		uintptr(InformationClass),
	)

	return types.ULONG(ret)
}
