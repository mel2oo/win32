package main

import (
	"fmt"
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"github.com/mel2oo/win32/advapi32/evntrace"
	"github.com/mel2oo/win32/tdh"
	"github.com/mel2oo/win32/types"
)

const (
	maxStringLen              = 1024
	maxBufferSize             = 1024
	evtBufferSize types.ULONG = 4096
)

var (
	callbackNext = uintptr(1)
	minBuffers   = runtime.NumCPU() * 2
	maxBuffers   = minBuffers + 20
)

func main() {
	props := evntrace.EventTraceProperties{
		Wnode: evntrace.WnodeHeader{
			BufferSize: types.ULONG(unsafe.Sizeof(evntrace.EventTraceProperties{}) + 2*maxStringLen),
			Guid:       types.GUID(evntrace.SystemTraceControlGuid),
			Flags:      evntrace.WnodeFlagTracedGUID,
		},
		BufferSize:        maxBufferSize,
		MinimumBuffers:    types.ULONG(minBuffers),
		MaximumBuffers:    types.ULONG(maxBuffers),
		LogFileMode:       evntrace.EventTraceRealTimeMode,
		FlushTimer:        types.ULONG(time.Second.Seconds()),
		EnableFlags:       evntrace.EventTraceFlagProcess,
		LogFileNameOffset: 0,
		LoggerNameOffset:  types.ULONG(unsafe.Sizeof(evntrace.EventTraceProperties{})),
	}

	var (
		errno   types.ULONG
		handle1 evntrace.TraceHandle
		handle2 evntrace.TraceHandle
	)
	errno = evntrace.StartTrace(&handle1, evntrace.KernelLoggerName, &props)
	if errno != 0 {
		if errno == types.ULONG(types.ERROR_ALREADY_EXISTS) {
			errno = evntrace.ControlTrace(0, evntrace.KernelLoggerName, &props, types.ULONG(evntrace.EventTraceControlQuery))
			errno = evntrace.ControlTrace(0, evntrace.KernelLoggerName, &props, types.ULONG(evntrace.EventTraceControlStop))
			errno = evntrace.StartTrace(&handle1, evntrace.KernelLoggerName, &props)
			if errno != 0 {
				return
			}
		}
		return
	}

	sysTraceFlags := make([]uint32, 8)
	errno = evntrace.TraceSetInformation(
		handle1,
		evntrace.TraceSystemTraceEnableFlagsInfo,
		sysTraceFlags, 8)
	if errno != 0 {
		return
	}
	sysTraceFlags[0] = uint32(evntrace.EventTraceFlagProcess)
	sysTraceFlags[4] = uint32(handle1)

	errno = evntrace.TraceSetInformation(
		handle1,
		evntrace.TraceSystemTraceEnableFlagsInfo,
		sysTraceFlags, 8)
	if errno != 0 {
		return
	}

	a1, _ := syscall.UTF16PtrFromString(evntrace.KernelLoggerName)

	trace := evntrace.EventTraceLogFile{
		LoggerName:     a1,
		LogFileMode:    evntrace.EventTraceRealTimeMode | evntrace.EventTraceNoPerProcessorBuffering,
		BufferCallback: syscall.NewCallback(bufferStatsCallback),
		EventCallback:  syscall.NewCallback(processKeventCallback),
	}

	handle2 = evntrace.OpenTrace(&trace)

	go func() {
		errno = evntrace.ProcessTrace(&handle2, 1, nil, nil)
		if errno == types.ULONG(types.ERROR_SUCCESS) {
			fmt.Println("process trace close success")
			return
		}

		fmt.Println(errno)
	}()

	select {}
}

func bufferStatsCallback(logfile *evntrace.EventTraceLogFile) uintptr {
	fmt.Println("bufferStatsCallback")
	return callbackNext
}

func processKeventCallback(evt *tdh.EventRecord) uintptr {
	fmt.Println("processKeventCallback")

	bufferSize := evtBufferSize
	buffer := tdh.TRACE_EVENT_INFO{}

	errno := tdh.TdhGetEventInformation(evt, 0, nil, &buffer, &bufferSize)
	if errno != types.ERROR_SUCCESS {
		return callbackNext
	}

	return callbackNext
}
