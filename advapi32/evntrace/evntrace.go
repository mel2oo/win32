//go:build windows
// +build windows

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/

package evntrace

import (
	"syscall"
	"unsafe"

	"github.com/mel2oo/win32/advapi32"
	"github.com/mel2oo/win32/typedef"
)

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-closetrace
func CloseTrace(TraceHandle TRACEHANDLE) typedef.ULONG {
	ret, _, _ := advapi32.ProcCloseTrace.Call(uintptr(TraceHandle))
	return typedef.ULONG(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-controltracew
func ControlTrace(TraceHandle TRACEHANDLE, InstanceName string,
	Properties PEVENT_TRACE_PROPERTIES, ControlCode typedef.ULONG) typedef.ULONG {
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

	return typedef.ULONG(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/evntrace/nf-evntrace-starttracew
func StartTrace(TraceHandle PTRACEHANDLE, InstanceName string, Properties PEVENT_TRACE_PROPERTIES) typedef.ULONG {
	a1, err := syscall.UTF16PtrFromString(InstanceName)
	if err != nil {
		return 0
	}

	ret, _, _ := advapi32.ProcStartTrace.Call(
		uintptr(unsafe.Pointer(TraceHandle)),
		uintptr(unsafe.Pointer(a1)),
		uintptr(unsafe.Pointer(Properties)),
	)

	return typedef.ULONG(ret)
}
