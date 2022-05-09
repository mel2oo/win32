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

// LOGPALETTE
type LOGPALETTE struct {
	Version    WORD
	NumEntries WORD
	PalEntry   [1]PALETTEENTRY
}

// VARTYPE
type VARTYPE USHORT

const (
	VT_EMPTY            VARTYPE = 0
	VT_NULL             VARTYPE = 1
	VT_I2               VARTYPE = 2
	VT_I4               VARTYPE = 3
	VT_R4               VARTYPE = 4
	VT_R8               VARTYPE = 5
	VT_CY               VARTYPE = 6
	VT_DATE             VARTYPE = 7
	VT_BSTR             VARTYPE = 8
	VT_DISPATCH         VARTYPE = 9
	VT_ERROR            VARTYPE = 10
	VT_BOOL             VARTYPE = 11
	VT_VARIANT          VARTYPE = 12
	VT_UNKNOWN          VARTYPE = 13
	VT_DECIMAL          VARTYPE = 14
	VT_I1               VARTYPE = 16
	VT_UI1              VARTYPE = 17
	VT_UI2              VARTYPE = 18
	VT_UI4              VARTYPE = 19
	VT_I8               VARTYPE = 20
	VT_UI8              VARTYPE = 21
	VT_INT              VARTYPE = 22
	VT_UINT             VARTYPE = 23
	VT_VOID             VARTYPE = 24
	VT_HRESULT          VARTYPE = 25
	VT_PTR              VARTYPE = 26
	VT_SAFEARRAY        VARTYPE = 27
	VT_CARRAY           VARTYPE = 28
	VT_USERDEFINED      VARTYPE = 29
	VT_LPSTR            VARTYPE = 30
	VT_LPWSTR           VARTYPE = 31
	VT_RECORD           VARTYPE = 36
	VT_INT_PTR          VARTYPE = 37
	VT_UINT_PTR         VARTYPE = 38
	VT_FILETIME         VARTYPE = 64
	VT_BLOB             VARTYPE = 65
	VT_STREAM           VARTYPE = 66
	VT_STORAGE          VARTYPE = 67
	VT_STREAMED_OBJECT  VARTYPE = 68
	VT_STORED_OBJECT    VARTYPE = 69
	VT_BLOB_OBJECT      VARTYPE = 70
	VT_CF               VARTYPE = 71
	VT_CLSID            VARTYPE = 72
	VT_VERSIONED_STREAM VARTYPE = 73
	VT_BSTR_BLOB        VARTYPE = 0xfff
	VT_VECTOR           VARTYPE = 0x1000
	VT_ARRAY            VARTYPE = 0x2000
	VT_BYREF            VARTYPE = 0x4000
	VT_RESERVED         VARTYPE = 0x8000
)

// BRECORD
type BRECORD struct {
	Record  *PVOID
	RecInfo uintptr
}

// VARIANT
type VARIANT struct {
	VT        VARTYPE
	Reserved1 WORD
	Reserved2 WORD
	Reserved3 WORD
	Val       LONG
}
type LPVARIANT *VARIANT

// VARIANTARG
type VARIANTARG VARIANT

// PROPVARIANT
type (
	PROPVARIANT    VARIANT
	REFPROPVARIANT *PROPVARIANT
)

// TYMED
type TYMED DWORD

const (
	TYMED_HGLOBAL  TYMED = 1
	TYMED_FILE     TYMED = 2
	TYMED_ISTREAM  TYMED = 4
	TYMED_ISTORAGE TYMED = 8
	TYMED_GDI      TYMED = 16
	TYMED_MFPICT   TYMED = 32
	TYMED_ENHMF    TYMED = 64
	TYMED_NULL     TYMED = 0
)

// DVASPECT
type DVASPECT DWORD

const (
	DVASPECT_CONTENT   DVASPECT = 1
	DVASPECT_THUMBNAIL DVASPECT = 2
	DVASPECT_ICON      DVASPECT = 4
	DVASPECT_DOCPRINT  DVASPECT = 8
)

// FORMATETC
type FORMATETC struct {
	Format CLIPFORMAT
	Ptd    *DVTARGETDEVICE
	Aspect DVASPECT
	Lindex LONG
	Tymed  TYMED
}

type LPFORMATETC *FORMATETC

// GETPROPERTYSTOREFLAGS
type GETPROPERTYSTOREFLAGS UINT

const (
	GPS_DEFAULT               GETPROPERTYSTOREFLAGS = 0
	GPS_HANDLERPROPERTIESONLY GETPROPERTYSTOREFLAGS = 0x1
	GPS_READWRITE             GETPROPERTYSTOREFLAGS = 0x2
	GPS_TEMPORARY             GETPROPERTYSTOREFLAGS = 0x4
	GPS_FASTPROPERTIESONLY    GETPROPERTYSTOREFLAGS = 0x8
	GPS_OPENSLOWITEM          GETPROPERTYSTOREFLAGS = 0x10
	GPS_DELAYCREATION         GETPROPERTYSTOREFLAGS = 0x20
	GPS_BESTEFFORT            GETPROPERTYSTOREFLAGS = 0x40
	GPS_NO_OPLOCK             GETPROPERTYSTOREFLAGS = 0x80
	GPS_MASK_VALID            GETPROPERTYSTOREFLAGS = 0xff
)

// PROPERTYKEY
type PROPERTYKEY struct {
	Fmtid GUID
	Pid   DWORD
}

type REFPROPERTYKEY *PROPERTYKEY

// CSPLATFORM
type CSPLATFORM struct {
	PlatformId    DWORD
	VersionHi     DWORD
	VersionLo     DWORD
	ProcessorArch DWORD
}

// QUERYCONTEXT
type QUERYCONTEXT struct {
	Context   DWORD
	Platform  CSPLATFORM
	Locale    LCID
	VersionHi DWORD
	VersionLo DWORD
}

// MULTI_QI
type MULTI_QI struct {
	IID *IID
	Itf uintptr
	Hr  HRESULT
}

// STGM_FLAGS
type STGM_FLAGS DWORD

const (
	STGM_DIRECT           STGM_FLAGS = 0x00000000
	STGM_TRANSACTED       STGM_FLAGS = 0x00010000
	STGM_SIMPLE           STGM_FLAGS = 0x08000000
	STGM_READ             STGM_FLAGS = 0x00000000
	STGM_WRITE            STGM_FLAGS = 0x00000001
	STGM_READWRITE        STGM_FLAGS = 0x00000002
	STGM_SHARE_DENY_NONE  STGM_FLAGS = 0x00000040
	STGM_SHARE_DENY_READ  STGM_FLAGS = 0x00000030
	STGM_SHARE_DENY_WRITE STGM_FLAGS = 0x00000020
	STGM_SHARE_EXCLUSIVE  STGM_FLAGS = 0x00000010
	STGM_PRIORITY         STGM_FLAGS = 0x00040000
	STGM_DELETEONRELEASE  STGM_FLAGS = 0x04000000
	STGM_NOSCRATCH        STGM_FLAGS = 0x00100000
	STGM_CREATE           STGM_FLAGS = 0x00001000
	STGM_CONVERT          STGM_FLAGS = 0x00020000
	STGM_FAILIFTHERE      STGM_FLAGS = 0x00000000
	STGM_NOSNAPSHOT       STGM_FLAGS = 0x00200000
	STGM_DIRECT_SWMR      STGM_FLAGS = 0x00400000
)

// THDTYPE
type THDTYPE UINT

const (
	THDTYPE_BLOCKMESSAGES   THDTYPE = 0
	THDTYPE_PROCESSMESSAGES THDTYPE = 1
)

// GLOBALOPT_PROPERTIES
type GLOBALOPT_PROPERTIES UINT

const (
	COMGLB_EXCEPTION_HANDLING     GLOBALOPT_PROPERTIES = 1
	COMGLB_APPID                  GLOBALOPT_PROPERTIES = 2
	COMGLB_RPC_THREADPOOL_SETTING GLOBALOPT_PROPERTIES = 3
)

// RPCOPT_PROPERTIES
type RPCOPT_PROPERTIES UINT

const (
	COMBND_RPCTIMEOUT      RPCOPT_PROPERTIES = 0x1
	COMBND_SERVER_LOCALITY RPCOPT_PROPERTIES = 0x2
)

// ApplicationType
type ApplicationType UINT

const (
	ServerApplication  ApplicationType = 0
	LibraryApplication ApplicationType = 1
)

// ShutdownType
type ShutdownType UINT

const (
	IdleShutdown   ShutdownType = 0
	ForcedShutdown ShutdownType = 1
)

// LOCKTYPE
type LOCKTYPE DWORD

const (
	LOCK_WRITE     LOCKTYPE = 1
	LOCK_EXCLUSIVE LOCKTYPE = 2
	LOCK_ONLYONCE  LOCKTYPE = 4
)

// ADVF
type ADVF DWORD

const (
	ADVF_NODATA            ADVF = 1
	ADVF_PRIMEFIRST        ADVF = 2
	ADVF_ONLYONCE          ADVF = 4
	ADVF_DATAONSTOP        ADVF = 64
	ADVFCACHE_NOHANDLER    ADVF = 8
	ADVFCACHE_FORCEBUILTIN ADVF = 16
	ADVFCACHE_ONSAVE       ADVF = 32
)

// APTTYPE
type APTTYPE int32

const (
	APTTYPE_CURRENT APTTYPE = -1
	APTTYPE_STA     APTTYPE = 0
	APTTYPE_MTA     APTTYPE = 1
	APTTYPE_NA      APTTYPE = 2
	APTTYPE_MAINSTA APTTYPE = 3
)

// APTTYPEQUALIFIER
type APTTYPEQUALIFIER UINT

const (
	APTTYPEQUALIFIER_NONE               APTTYPEQUALIFIER = 0
	APTTYPEQUALIFIER_IMPLICIT_MTA       APTTYPEQUALIFIER = 1
	APTTYPEQUALIFIER_NA_ON_MTA          APTTYPEQUALIFIER = 2
	APTTYPEQUALIFIER_NA_ON_STA          APTTYPEQUALIFIER = 3
	APTTYPEQUALIFIER_NA_ON_IMPLICIT_MTA APTTYPEQUALIFIER = 4
	APTTYPEQUALIFIER_NA_ON_MAINSTA      APTTYPEQUALIFIER = 5
)

// DISPID
type DISPID LONG

const (
	DISPID_UNKNOWN                   DISPID = -1
	DISPID_VALUE                     DISPID = 0
	DISPID_PROPERTYPUT               DISPID = -3
	DISPID_NEWENUM                   DISPID = -4
	DISPID_EVALUATE                  DISPID = -5
	DISPID_CONSTRUCTOR               DISPID = -6
	DISPID_DESTRUCTOR                DISPID = -7
	DISPID_COLLECT                   DISPID = -8
	DISPID_AUTOSIZE                  DISPID = -500
	DISPID_BACKCOLOR                 DISPID = -501
	DISPID_BACKSTYLE                 DISPID = -502
	DISPID_BORDERCOLOR               DISPID = -503
	DISPID_BORDERSTYLE               DISPID = -504
	DISPID_BORDERWIDTH               DISPID = -505
	DISPID_DRAWMODE                  DISPID = -507
	DISPID_DRAWSTYLE                 DISPID = -508
	DISPID_DRAWWIDTH                 DISPID = -509
	DISPID_FILLCOLOR                 DISPID = -510
	DISPID_FILLSTYLE                 DISPID = -511
	DISPID_FONT                      DISPID = -512
	DISPID_FORECOLOR                 DISPID = -513
	DISPID_ENABLED                   DISPID = -514
	DISPID_HWND                      DISPID = -515
	DISPID_TABSTOP                   DISPID = -516
	DISPID_TEXT                      DISPID = -517
	DISPID_CAPTION                   DISPID = -518
	DISPID_BORDERVISIBLE             DISPID = -519
	DISPID_APPEARANCE                DISPID = -520
	DISPID_MOUSEPOINTER              DISPID = -521
	DISPID_MOUSEICON                 DISPID = -522
	DISPID_PICTURE                   DISPID = -523
	DISPID_VALID                     DISPID = -524
	DISPID_READYSTATE                DISPID = -525
	DISPID_LISTINDEX                 DISPID = -526
	DISPID_SELECTED                  DISPID = -527
	DISPID_LIST                      DISPID = -528
	DISPID_COLUMN                    DISPID = -529
	DISPID_LISTCOUNT                 DISPID = -531
	DISPID_MULTISELECT               DISPID = -532
	DISPID_MAXLENGTH                 DISPID = -533
	DISPID_PASSWORDCHAR              DISPID = -534
	DISPID_SCROLLBARS                DISPID = -535
	DISPID_WORDWRAP                  DISPID = -536
	DISPID_MULTILINE                 DISPID = -537
	DISPID_NUMBEROFROWS              DISPID = -538
	DISPID_NUMBEROFCOLUMNS           DISPID = -539
	DISPID_DISPLAYSTYLE              DISPID = -540
	DISPID_GROUPNAME                 DISPID = -541
	DISPID_IMEMODE                   DISPID = -542
	DISPID_ACCELERATOR               DISPID = -543
	DISPID_ENTERKEYBEHAVIOR          DISPID = -544
	DISPID_TABKEYBEHAVIOR            DISPID = -545
	DISPID_SELTEXT                   DISPID = -546
	DISPID_SELSTART                  DISPID = -547
	DISPID_SELLENGTH                 DISPID = -548
	DISPID_REFRESH                   DISPID = -550
	DISPID_DOCLICK                   DISPID = -551
	DISPID_ABOUTBOX                  DISPID = -552
	DISPID_ADDITEM                   DISPID = -553
	DISPID_CLEAR                     DISPID = -554
	DISPID_REMOVEITEM                DISPID = -555
	DISPID_CLICK                     DISPID = -600
	DISPID_DBLCLICK                  DISPID = -601
	DISPID_KEYDOWN                   DISPID = -602
	DISPID_KEYPRESS                  DISPID = -603
	DISPID_KEYUP                     DISPID = -604
	DISPID_MOUSEDOWN                 DISPID = -605
	DISPID_MOUSEMOVE                 DISPID = -606
	DISPID_MOUSEUP                   DISPID = -607
	DISPID_ERROREVENT                DISPID = -608
	DISPID_READYSTATECHANGE          DISPID = -609
	DISPID_CLICK_VALUE               DISPID = -610
	DISPID_RIGHTTOLEFT               DISPID = -611
	DISPID_TOPTOBOTTOM               DISPID = -612
	DISPID_THIS                      DISPID = -613
	DISPID_AMBIENT_BACKCOLOR         DISPID = -701
	DISPID_AMBIENT_DISPLAYNAME       DISPID = -702
	DISPID_AMBIENT_FONT              DISPID = -703
	DISPID_AMBIENT_FORECOLOR         DISPID = -704
	DISPID_AMBIENT_LOCALEID          DISPID = -705
	DISPID_AMBIENT_MESSAGEREFLECT    DISPID = -706
	DISPID_AMBIENT_SCALEUNITS        DISPID = -707
	DISPID_AMBIENT_TEXTALIGN         DISPID = -708
	DISPID_AMBIENT_USERMODE          DISPID = -709
	DISPID_AMBIENT_UIDEAD            DISPID = -710
	DISPID_AMBIENT_SHOWGRABHANDLES   DISPID = -711
	DISPID_AMBIENT_SHOWHATCHING      DISPID = -712
	DISPID_AMBIENT_DISPLAYASDEFAULT  DISPID = -713
	DISPID_AMBIENT_SUPPORTSMNEMONICS DISPID = -714
	DISPID_AMBIENT_AUTOCLIP          DISPID = -715
	DISPID_AMBIENT_APPEARANCE        DISPID = -716
	DISPID_AMBIENT_CODEPAGE          DISPID = -725
	DISPID_AMBIENT_PALETTE           DISPID = -726
	DISPID_AMBIENT_CHARSET           DISPID = -727
	DISPID_AMBIENT_TRANSFERPRIORITY  DISPID = -728
	DISPID_AMBIENT_RIGHTTOLEFT       DISPID = -732
	DISPID_AMBIENT_TOPTOBOTTOM       DISPID = -733
	DISPID_Name                      DISPID = -800
	DISPID_Delete                    DISPID = -801
	DISPID_Object                    DISPID = -802
	DISPID_Parent                    DISPID = -803
)

// BIND_FLAGS
type BIND_FLAGS DWORD

const (
	BIND_MAYBOTHERUSER     BIND_FLAGS = 1
	BIND_JUSTTESTEXISTENCE BIND_FLAGS = 2
)

// BIND_OPTS
type BIND_OPTS struct {
	Struct            DWORD
	Flags             BIND_FLAGS
	Mode              STGM_FLAGS
	TickCountDeadline DWORD
}

// STATSTG
type STATSTG struct {
	Name           LPOLESTR
	Type           DWORD
	Size           ULARGE_INTEGER
	Mtime          FILETIME
	Ctime          FILETIME
	Atime          FILETIME
	Mode           STGM_FLAGS
	LocksSupported LOCKTYPE
	Clsid          CLSID
	StateBits      DWORD
	Reserved       DWORD
}

// INTERFACEINFO
type INTERFACEINFO struct {
	Unk    uintptr
	Iid    IID
	Method WORD
}

type LPINTERFACEINFO *INTERFACEINFO

// OLEINPLACEFRAMEINFO
type OLEINPLACEFRAMEINFO struct {
	Cb           UINT
	MDIApp       BOOL
	Frame        HWND
	Haccel       HACCEL
	AccelEntries UINT
}

type LPOLEINPLACEFRAMEINFO *OLEINPLACEFRAMEINFO

// OLEVERB
type OLEVERB struct {
	Verb     LONG
	VerbName LPOLESTR
	Flags    DWORD
	Attribs  DWORD
}

type LPOLEVERB *OLEVERB

// STATDATA
type STATDATA struct {
	Formatetc  FORMATETC
	Advf       DWORD
	AdvSink    uintptr
	Connection DWORD
}

// ContextProperty
type ContextProperty struct {
	PolicyId GUID
	Flags    CPFLAGS
	Unk      uintptr
}

// StorageLayout
type StorageLayout struct {
	LayoutType  DWORD
	ElementName *OLECHAR
	Offset      LARGE_INTEGER
	Bytes       LARGE_INTEGER
}

// OLECMDF
type OLECMDF DWORD

const (
	OLECMDF_SUPPORTED         OLECMDF = 0x1
	OLECMDF_ENABLED           OLECMDF = 0x2
	OLECMDF_LATCHED           OLECMDF = 0x4
	OLECMDF_NINCHED           OLECMDF = 0x8
	OLECMDF_INVISIBLE         OLECMDF = 0x10
	OLECMDF_DEFHIDEONCTXTMENU OLECMDF = 0x20
)

// PROPBAG2_TYPE
type PROPBAG2_TYPE UINT

const (
	PROPBAG2_TYPE_UNDEFINED PROPBAG2_TYPE = 0
	PROPBAG2_TYPE_DATA      PROPBAG2_TYPE = 1
	PROPBAG2_TYPE_URL       PROPBAG2_TYPE = 2
	PROPBAG2_TYPE_OBJECT    PROPBAG2_TYPE = 3
	PROPBAG2_TYPE_STREAM    PROPBAG2_TYPE = 4
	PROPBAG2_TYPE_STORAGE   PROPBAG2_TYPE = 5
	PROPBAG2_TYPE_MONIKER   PROPBAG2_TYPE = 6
)

// PROPBAG2
type PROPBAG2 struct {
	Type   PROPBAG2_TYPE
	VT     VARTYPE
	CfType CLIPFORMAT
	Hint   DWORD
	Name   LPOLESTR
	Clsid  CLSID
}

// CLSCTX
type CLSCTX DWORD

const (
	CLSCTX_INPROC_SERVER          CLSCTX = 0x1
	CLSCTX_INPROC_HANDLER         CLSCTX = 0x2
	CLSCTX_LOCAL_SERVER           CLSCTX = 0x4
	CLSCTX_INPROC_SERVER16        CLSCTX = 0x8
	CLSCTX_REMOTE_SERVER          CLSCTX = 0x10
	CLSCTX_INPROC_HANDLER16       CLSCTX = 0x20
	CLSCTX_RESERVED1              CLSCTX = 0x40
	CLSCTX_RESERVED2              CLSCTX = 0x80
	CLSCTX_RESERVED3              CLSCTX = 0x100
	CLSCTX_RESERVED4              CLSCTX = 0x200
	CLSCTX_NO_CODE_DOWNLOAD       CLSCTX = 0x400
	CLSCTX_RESERVED5              CLSCTX = 0x800
	CLSCTX_NO_CUSTOM_MARSHAL      CLSCTX = 0x1000
	CLSCTX_ENABLE_CODE_DOWNLOAD   CLSCTX = 0x2000
	CLSCTX_NO_FAILURE_LOG         CLSCTX = 0x4000
	CLSCTX_DISABLE_AAA            CLSCTX = 0x8000
	CLSCTX_ENABLE_AAA             CLSCTX = 0x10000
	CLSCTX_FROM_DEFAULT_CONTEXT   CLSCTX = 0x20000
	CLSCTX_ACTIVATE_32_BIT_SERVER CLSCTX = 0x40000
	CLSCTX_ACTIVATE_64_BIT_SERVER CLSCTX = 0x80000
	CLSCTX_ENABLE_CLOAKING        CLSCTX = 0x100000
	CLSCTX_PS_DLL                 CLSCTX = 0x80000000
)

// CLSSPEC_u_s1
type CLSSPEC_u_s1 struct {
	PackageName LPOLESTR
	PolicyId    GUID
}

// CLSSPEC_u_s2
type CLSSPEC_u_s2 struct {
	ObjectId GUID
	PolicyId GUID
}

// TYSPEC
type TYSPEC DWORD

const (
	TYSPEC_CLSID       TYSPEC = 0
	TYSPEC_FILEEXT     TYSPEC = 1
	TYSPEC_MIMETYPE    TYSPEC = 2
	TYSPEC_FILENAME    TYSPEC = 3
	TYSPEC_PROGID      TYSPEC = 4
	TYSPEC_PACKAGENAME TYSPEC = 5
	TYSPEC_OBJECTID    TYSPEC = 6
)

// CLSSPEC
type CLSSPEC struct {
	Tyspec TYSPEC
	Clsid  CLSID
}

// STGMEDIUM
type STGMEDIUM struct {
	Tymed         TYMED
	Bitmap        HBITMAP
	UnkForRelease uintptr
}
type LPSTGMEDIUM *STGMEDIUM

// COINIT_FLAG
type COINIT_FLAG DWORD

const (
	COINIT_APARTMENTTHREADED COINIT_FLAG = 0x2
	COINIT_MULTITHREADED     COINIT_FLAG = 0x0
	COINIT_DISABLE_OLE1DDE   COINIT_FLAG = 0x4
	COINIT_SPEED_OVER_MEMORY COINIT_FLAG = 0x8
)

// FADF_FLAGS
type FADF_FLAGS USHORT

const (
	FADF_AUTO        FADF_FLAGS = 0x1
	FADF_STATIC      FADF_FLAGS = 0x2
	FADF_EMBEDDED    FADF_FLAGS = 0x4
	FADF_FIXEDSIZE   FADF_FLAGS = 0x10
	FADF_RECORD      FADF_FLAGS = 0x20
	FADF_HAVEIID     FADF_FLAGS = 0x40
	FADF_HAVEVARTYPE FADF_FLAGS = 0x80
	FADF_BSTR        FADF_FLAGS = 0x100
	FADF_UNKNOWN     FADF_FLAGS = 0x200
	FADF_DISPATCH    FADF_FLAGS = 0x400
	FADF_VARIANT     FADF_FLAGS = 0x800
	FADF_RESERVED    FADF_FLAGS = 0xf008
)

// SAFEARRAYBOUND
type SAFEARRAYBOUND struct {
	Elements ULONG
	Lbound   LONG
}
type LPSAFEARRAY *SAFEARRAYBOUND

// SAFEARRAY
type SAFEARRAY struct {
	Dims      USHORT
	Features  FADF_FLAGS
	Elements  ULONG
	Locks     ULONG
	Data      PVOID
	Rgsabound [1]SAFEARRAYBOUND
}

// CUSTDATAITEM
type CUSTDATAITEM struct {
	Guid  GUID
	Value VARIANTARG
}

type LPCUSTDATAITEM *CUSTDATAITEM

// CUSTDATA
type CUSTDATA struct {
	CustData    DWORD
	PrgCustData LPCUSTDATAITEM
}
type LPCUSTDATA *CUSTDATA

// InvokeFlags
type InvokeFlags WORD

const (
	DISPATCH_METHOD         InvokeFlags = 0x1
	DISPATCH_PROPERTYGET    InvokeFlags = 0x2
	DISPATCH_PROPERTYPUT    InvokeFlags = 0x4
	DISPATCH_PROPERTYPUTREF InvokeFlags = 0x8
)

// CALLCONV
type CALLCONV UINT

const (
	CC_FASTCALL   CALLCONV = 0
	CC_CDECL      CALLCONV = 1
	CC_PASCAL     CALLCONV = 2
	CC_MACPASCAL  CALLCONV = 3
	CC_STDCALL    CALLCONV = 4
	CC_FPFASTCALL CALLCONV = 5
	CC_SYSCALL    CALLCONV = 6
	CC_MPWCDECL   CALLCONV = 7
	CC_MPWPASCAL  CALLCONV = 8
)

// STATFLAG
type STATFLAG DWORD

const (
	STATFLAG_DEFAULT STATFLAG = 0
	STATFLAG_NONAME  STATFLAG = 1
	STATFLAG_NOOPEN  STATFLAG = 2
)

// STGC
type STGC DWORD

const (
	STGC_DEFAULT                            STGC = 0
	STGC_OVERWRITE                          STGC = 1
	STGC_ONLYIFCURRENT                      STGC = 2
	STGC_DANGEROUSLYCOMMITMERELYTODISKCACHE STGC = 4
	STGC_CONSOLIDATE                        STGC = 8
)

// DBKIND
type DBKIND DWORD

const (
	DBKIND_GUID_NAME    DBKIND = 0
	DBKIND_GUID_PROPID  DBKIND = 1
	DBKIND_NAME         DBKIND = 2
	DBKIND_PGUID_NAME   DBKIND = 3
	DBKIND_PGUID_PROPID DBKIND = 4
	DBKIND_PROPID       DBKIND = 5
	DBKIND_GUID         DBKIND = 6
)

// DBID
type DBID struct {
	Guid GUID
	Kind DBKIND
	Name LPOLESTR
}

// PROPSETFLAG
type PROPSETFLAG DWORD

const (
	PROPSETFLAG_DEFAULT        PROPSETFLAG = 0
	PROPSETFLAG_NONSIMPLE      PROPSETFLAG = 1
	PROPSETFLAG_ANSI           PROPSETFLAG = 2
	PROPSETFLAG_UNBUFFERED     PROPSETFLAG = 4
	PROPSETFLAG_CASE_SENSITIVE PROPSETFLAG = 8
)

// STATPROPSETSTG
type STATPROPSETSTG struct {
	Fmtid       FMTID
	Clsid       CLSID
	GrfFlags    PROPSETFLAG
	Mtime       FILETIME
	Ctime       FILETIME
	Atime       FILETIME
	DwOSVersion DWORD
}

// STATPROPSTG
type STATPROPSTG struct {
	Name   LPOLESTR
	Propid PROPID
	Vt     VARTYPE
}

// IMPLTYPEFLAGS
type IMPLTYPEFLAGS UINT

const (
	IMPLTYPEFLAG_FDEFAULT       IMPLTYPEFLAGS = 0x1
	IMPLTYPEFLAG_FSOURCE        IMPLTYPEFLAGS = 0x2
	IMPLTYPEFLAG_FRESTRICTED    IMPLTYPEFLAGS = 0x4
	IMPLTYPEFLAG_FDEFAULTVTABLE IMPLTYPEFLAGS = 0x8
)
