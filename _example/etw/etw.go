package main

import (
	"fmt"
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"github.com/mel2oo/win32/advapi32"
	"github.com/mel2oo/win32/advapi32/evntrace"
	"github.com/mel2oo/win32/tdh"
	"github.com/mel2oo/win32/types"
)

const (
	KernelLoggerSession = "NT Kernel Logger"
	maxStringLen        = 1024
	maxBufferSize       = 1024
)

var (
	callbackNext = uintptr(1)
	minBuffers   = runtime.NumCPU() * 2
	maxBuffers   = minBuffers + 20
)

var KernelTraceGUID = types.GUID{
	Data1: 0x9E814AAD,
	Data2: 0x3204,
	Data3: 0x11D2,
	Data4: [8]byte{0x9A, 0x82, 0x00, 0x60, 0x08, 0xA8, 0x69, 0x39},
}

func main() {
	props := types.EVENT_TRACE_PROPERTIES{
		Wnode: types.WNODE_HEADER{
			BufferSize: types.ULONG(unsafe.Sizeof(types.EVENT_TRACE_PROPERTIES{}) + 2*maxStringLen),
			Guid:       KernelTraceGUID,
			Flags:      types.WNODE_FLAG_TRACED_GUID,
		},
		BufferSize:        maxBufferSize,
		MinimumBuffers:    types.ULONG(minBuffers),
		MaximumBuffers:    types.ULONG(maxBuffers),
		LogFileMode:       types.EVENT_TRACE_REAL_TIME_MODE,
		FlushTimer:        types.ULONG(time.Second.Seconds()),
		EnableFlags:       types.EVENT_TRACE_FLAG_PROCESS,
		LogFileNameOffset: 0,
		LoggerNameOffset:  types.ULONG(unsafe.Sizeof(types.EVENT_TRACE_PROPERTIES{})),
	}

	var (
		errno   types.ULONG
		handle1 advapi32.TRACEHANDLE
		handle2 advapi32.TRACEHANDLE
	)
	errno = evntrace.StartTrace(&handle1, KernelLoggerSession, &props)
	if errno != 0 {
		if errno == types.ULONG(types.ERROR_ALREADY_EXISTS) {
			errno = evntrace.ControlTrace(0, KernelLoggerSession, &props, 0)
			errno = evntrace.ControlTrace(0, KernelLoggerSession, &props, 1)
		}
		errno = evntrace.StartTrace(&handle1, KernelLoggerSession, &props)
		if errno != 0 {
			return
		}
	}

	a1, _ := syscall.UTF16PtrFromString(KernelLoggerSession)

	trace := types.EVENT_TRACE_LOGFILE{
		LoggerName:     a1,
		LogFileMode:    types.EVENT_TRACE_REAL_TIME_MODE | types.EVENT_TRACE_NO_PER_PROCESSOR_BUFFERING,
		BufferCallback: types.PEVENT_TRACE_BUFFER_CALLBACK(syscall.NewCallback(bufferStatsCallback)),
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

func bufferStatsCallback(logfile *types.EVENT_TRACE_LOGFILE) uintptr {
	fmt.Println("bufferStatsCallback")
	return callbackNext
}

func processKeventCallback(evt *tdh.EVENT_RECORD) uintptr {
	fmt.Println("processKeventCallback")
	return callbackNext
}
