package main

import (
	"runtime"
	"time"
	"unsafe"

	"github.com/mel2oo/win32/advapi32"
	"github.com/mel2oo/win32/advapi32/evntrace"
	"github.com/mel2oo/win32/types"
)

const (
	KernelLoggerSession = "NT Kernel Logger"
	maxStringLen        = 1024
	maxBufferSize       = 1024
)

var (
	minBuffers = runtime.NumCPU() * 2
	maxBuffers = minBuffers + 20
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
		errno  types.ULONG
		handle advapi32.TRACEHANDLE
	)
	errno = evntrace.StartTrace(&handle, KernelLoggerSession, &props)
	if errno != 0 {
		return
	}
}
