package types

// Structures (Different Versions)
type (
	LPCPROPSHEETPAGE   LPVOID
	LPCPROPSHEETHEADER LPVOID
	LPPROPSHEETHEADER  LPVOID
)

// interface IUnknown IUnknown IStream IStream IPropertySetStorage IPropertySetStorage AsyncIAdviseSink AsyncIAdviseSink AsyncIAdviseSink2 AsyncIAdviseSink2 AsyncIMultiQI AsyncIMultiQI AsyncIPipeByte AsyncIPipeByte AsyncIPipeDouble AsyncIPipeDouble AsyncIPipeLong AsyncIPipeLong IAddrExclusionControl IAddrExclusionControl IAddrTrackingControl IAddrTrackingControl IAdviseSink IAdviseSink IAdviseSink2 IAdviseSink2 IAsyncManager IAsyncManager IAsyncRpcChannelBuffer IAsyncRpcChannelBuffer IBindCtx IBindCtx IBlockingLock IBlockingLock ICallFactory ICallFactory ICancelMethodCalls ICancelMethodCalls IChannelHook IChannelHook IClassActivator IClassActivator IClientSecurity IClientSecurity IComThreadingInfo IComThreadingInfo IContext IContext IContinue IContinue IDataAdviseHolder IDataAdviseHolder IDataObject IDataObject
type (
	IDirectWriterLock          uintptr
	IDropSource                uintptr
	IDropSourceNotify          uintptr
	IDropTarget                uintptr
	IDummyHICONIncluder        uintptr
	IEnumContextProps          uintptr
	IEnumFORMATETC             uintptr
	IEnumMoniker               uintptr
	IEnumOLEVERB               uintptr
	IEnumSTATDATA              uintptr
	IEnumSTATSTG               uintptr
	IEnumString                uintptr
	IEnumUnknown               uintptr
	IExternalConnection        uintptr
	IFillLockBytes             uintptr
	IForegroundTransfer        uintptr
	IGlobalInterfaceTable      uintptr
	IGlobalOptions             uintptr
	IInitializeSpy             uintptr
	IInternalUnknown           uintptr
	ILayoutStorage             uintptr
	ILockBytes                 uintptr
	IMalloc                    uintptr
	IMallocSpy                 uintptr
	IMarshal                   uintptr
	IMarshal2                  uintptr
	IMessageFilter             uintptr
	IMoniker                   uintptr
	IMultiQI                   uintptr
	IObjContext                uintptr
	IOleAdviseHolder           uintptr
	IOleCache                  uintptr
	IOleCache2                 uintptr
	IOleCacheControl           uintptr
	IOleClientSite             uintptr
	IOleContainer              uintptr
	IOleInPlaceActiveObject    uintptr
	IOleInPlaceFrame           uintptr
	IOleInPlaceObject          uintptr
	IOleInPlaceSite            uintptr
	IOleInPlaceUIWindow        uintptr
	IOleItemContainer          uintptr
	IOleLink                   uintptr
	IOleObject                 uintptr
	IOleWindow                 uintptr
	IOplockStorage             uintptr
	IParseDisplayName          uintptr
	IPersist                   uintptr
	IPersistFile               uintptr
	IPersistStorage            uintptr
	IPersistStream             uintptr
	IPipeByte                  uintptr
	IPipeDouble                uintptr
	IPipeLong                  uintptr
	IProcessInitControl        uintptr
	IProcessLock               uintptr
	IProgressNotify            uintptr
	IPSFactoryBuffer           uintptr
	IReleaseMarshalBuffers     uintptr
	IRootStorage               uintptr
	IROTData                   uintptr
	IRpcChannelBuffer          uintptr
	IRpcChannelBuffer2         uintptr
	IRpcChannelBuffer3         uintptr
	IRpcHelper                 uintptr
	IRpcOptions                uintptr
	IRpcProxyBuffer            uintptr
	IRpcStubBuffer             uintptr
	IRpcSyntaxNegotiate        uintptr
	IRunnableObject            uintptr
	IRunningObjectTable        uintptr
	ISequentialStream          uintptr
	IServerSecurity            uintptr
	IStdMarshalInfo            uintptr
	IStorage                   uintptr
	IStream                    uintptr
	ISurrogate                 uintptr
	ISurrogateService          uintptr
	ISynchronize               uintptr
	ISynchronizeContainer      uintptr
	ISynchronizeEvent          uintptr
	ISynchronizeHandle         uintptr
	ISynchronizeMutex          uintptr
	IThumbnailExtractor        uintptr
	ITimeAndNoticeControl      uintptr
	IUrlMon                    uintptr
	IViewObject                uintptr
	IViewObject2               uintptr
	IWaitMultiple              uintptr
	IContinueCallback          uintptr
	IEnumOleDocumentViews      uintptr
	IOleCommandTarget          uintptr
	IOleDocument               uintptr
	IOleDocumentSite           uintptr
	IOleDocumentView           uintptr
	IPrint                     uintptr
	IProtectedModeMenuServices uintptr
	IProtectFocus              uintptr
	IZoomEvents                uintptr
	IPropertyStorage           uintptr
	IClassFactory              uintptr
	IRecordInfo                uintptr
	IEnumSTATPROPSTG           uintptr
	IEnumSTATPROPSETSTG        uintptr
)

// Forward declaration for IDispatch
type (
	IDispatch  uintptr
	LPDISPATCH *IDispatch
)

// Variables
type (
	WINOLEAPI       STDAPI
	HPROPSHEETPAGE  LPVOID
	LPCOLESTR       LPCWSTR
	LPOLESTR        LPCWSTR
	OLECHAR         WCHAR
	PROPID          ULONG
	FMTID           GUID
	REFFMTID        *FMTID
	CLIPFORMAT      WORD
	CPFLAGS         DWORD
	HTASK           HANDLE
	LPCBORDERWIDTHS LPCRECT
	HOLEMENU        HGLOBAL
	HMETAFILEPICT   HANDLE
)

// VARIANT_BOOL
type VARIANT_BOOL SHORT

const (
	VARIANT_TRUE  VARIANT_BOOL = -1
	VARIANT_FALSE VARIANT_BOOL = 0
)

// MSHLFLAGS
type MSHLFLAGS DWORD

const (
	MSHLFLAGS_NORMAL      MSHLFLAGS = 0
	MSHLFLAGS_TABLESTRONG MSHLFLAGS = 1
	MSHLFLAGS_TABLEWEAK   MSHLFLAGS = 2
	MSHLFLAGS_NOPING      MSHLFLAGS = 4
	MSHLFLAGS_RESERVED1   MSHLFLAGS = 8
	MSHLFLAGS_RESERVED2   MSHLFLAGS = 16
	MSHLFLAGS_RESERVED3   MSHLFLAGS = 32
	MSHLFLAGS_RESERVED4   MSHLFLAGS = 64
)

// DVTARGETDEVICE
type DVTARGETDEVICE struct {
	Size             DWORD
	DriverNameOffset WORD
	DeviceNameOffset WORD
	PortNameOffset   WORD
	ExtDevmodeOffset WORD
	Data             [1]BYTE
}

// SERIALIZEDPROPERTYVALUE
type SERIALIZEDPROPERTYVALUE struct {
	Type DWORD
	Rgb  [1]BYTE
}

// RPCOLEMESSAGE
type RPCOLEMESSAGE struct {
	Reserved1          LPVOID
	DataRepresentation ULONG
	Buffer             LPVOID
	CbBuffer           ULONG
	Method             ULONG
	Reserved2          [5]LPVOID
	RpcFlags           ULONG
}

// OLEMENUGROUPWIDTHS
type OLEMENUGROUPWIDTHS struct {
	Width [6]LONG
}

type LPOLEMENUGROUPWIDTHS *OLEMENUGROUPWIDTHS

// PALETTEENTRYFlags
type PALETTEENTRYFlags BYTE

const (
	PC_RESERVED   PALETTEENTRYFlags = 0x01
	PC_EXPLICIT   PALETTEENTRYFlags = 0x02
	PC_NOCOLLAPSE PALETTEENTRYFlags = 0x04
)

// PALETTEENTRY
type PALETTEENTRY struct {
	Red   BYTE
	Green BYTE
	Blue  BYTE
	Flags PALETTEENTRYFlags
}

type LPPALETTEENTRY *PALETTEENTRY
