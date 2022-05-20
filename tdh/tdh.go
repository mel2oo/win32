package tdh

import (
	"unsafe"

	"github.com/mel2oo/win32/types"
	"golang.org/x/sys/windows"
)

var (
	dll = windows.NewLazyDLL("tdh.dll")

	procTdhAggregatePayloadFilters             = dll.NewProc("TdhAggregatePayloadFilters")
	procTdhCleanupPayloadEventFilterDescriptor = dll.NewProc("TdhCleanupPayloadEventFilterDescriptor")
	procTdhCloseDecodingHandle                 = dll.NewProc("TdhCloseDecodingHandle")
	procTdhCreatePayloadFilter                 = dll.NewProc("TdhCreatePayloadFilter")
	procTdhDeletePayloadFilter                 = dll.NewProc("TdhDeletePayloadFilter")
	procTdhEnumerateManifestProviderEvents     = dll.NewProc("TdhEnumerateManifestProviderEvents")
	procTdhEnumerateProviderFieldInformation   = dll.NewProc("TdhEnumerateProviderFieldInformation")
	procTdhEnumerateProviderFilters            = dll.NewProc("TdhEnumerateProviderFilters")
	procTdhEnumerateProviders                  = dll.NewProc("TdhEnumerateProviders")
	procTdhFormatProperty                      = dll.NewProc("TdhFormatProperty")
	procTdhGetDecodingParameter                = dll.NewProc("TdhGetDecodingParameter")
	procTdhGetEventInformation                 = dll.NewProc("TdhGetEventInformation")
	procTdhGetEventMapInformation              = dll.NewProc("TdhGetEventMapInformation")
	procTdhGetManifestEventInformation         = dll.NewProc("TdhGetManifestEventInformation")
	procTdhGetProperty                         = dll.NewProc("TdhGetProperty")
	procTdhGetPropertySize                     = dll.NewProc("TdhGetPropertySize")
	procTdhGetWppMessage                       = dll.NewProc("TdhGetWppMessage")
	procTdhGetWppProperty                      = dll.NewProc("TdhGetWppProperty")
	procTdhLoadManifest                        = dll.NewProc("TdhLoadManifest")
	procTdhLoadManifestFromBinary              = dll.NewProc("TdhLoadManifestFromBinary")
	procTdhOpenDecodingHandle                  = dll.NewProc("TdhOpenDecodingHandle")
	procTdhQueryProviderFieldInformation       = dll.NewProc("TdhQueryProviderFieldInformation")
	procTdhSetDecodingParameter                = dll.NewProc("TdhSetDecodingParameter")
	procTdhUnloadManifest                      = dll.NewProc("TdhUnloadManifest")
)

// https://docs.microsoft.com/en-us/windows/win32/api/tdh/nf-tdh-tdhgeteventinformation
func TdhGetEventInformation(Event *EventRecord, TdhContextCount types.ULONG,
	TdhContext *TdhContext, Buffer *TraceEventInfo, BufferSize *types.ULONG) types.ErrorCode {
	ret, _, _ := procTdhGetEventInformation.Call(
		uintptr(unsafe.Pointer(Event)),
		uintptr(TdhContextCount),
		uintptr(unsafe.Pointer(TdhContext)),
		uintptr(unsafe.Pointer(Buffer)),
		uintptr(unsafe.Pointer(BufferSize)),
	)

	return types.ErrorCode(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/tdh/nf-tdh-tdhgetproperty
func TdhGetProperty(pEvent *EventRecord, TdhContextCount types.ULONG, pTdhContext *TdhContext,
	PropertyDataCount types.ULONG, pPropertyData *PropertyDataDescriptor, BufferSize types.ULONG,
	pBuffer *types.BYTE) types.ErrorCode {
	ret, _, _ := procTdhGetProperty.Call(
		uintptr(unsafe.Pointer(pEvent)),
		uintptr(TdhContextCount),
		uintptr(unsafe.Pointer(pTdhContext)),
		uintptr(PropertyDataCount),
		uintptr(unsafe.Pointer(pPropertyData)),
		uintptr(BufferSize),
		uintptr(unsafe.Pointer(pBuffer)),
	)

	return types.ErrorCode(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/tdh/nf-tdh-tdhgetpropertysize
func TdhGetPropertySize(pEvent *EventRecord, TdhContextCount types.ULONG, pTdhContext *TdhContext,
	PropertyDataCount types.ULONG, pPropertyData *PropertyDataDescriptor, pPropertySize *types.ULONG) types.ErrorCode {
	ret, _, _ := procTdhGetPropertySize.Call(
		uintptr(unsafe.Pointer(pEvent)),
		uintptr(TdhContextCount),
		uintptr(unsafe.Pointer(pTdhContext)),
		uintptr(PropertyDataCount),
		uintptr(unsafe.Pointer(pPropertyData)),
		uintptr(unsafe.Pointer(pPropertySize)),
	)

	return types.ErrorCode(ret)
}
