package tdh

import "github.com/mel2oo/win32/types"

// Variables
type (
	TDH_HANDLE  types.HANDLE
	PTDH_HANDLE *TDH_HANDLE
)

// TRACE_PROVIDER_INFO
type TRACE_PROVIDER_INFO struct {
	ProviderGuid       types.GUID
	SchemaSource       types.ULONG
	ProviderNameOffset types.ULONG
}

// PROVIDER_ENUMERATION_INFO
type PROVIDER_ENUMERATION_INFO struct {
	NumberOfProviders      types.ULONG
	Reserved               types.ULONG
	TraceProviderInfoArray [1]TRACE_PROVIDER_INFO
}

type PPROVIDER_ENUMERATION_INFO *PROVIDER_ENUMERATION_INFO

// DecodingSource
type DecodingSource types.UINT

const (
	DecodingSourceXMLFile DecodingSource = 0
	DecodingSourceWbem    DecodingSource = 1
	DecodingSourceWPP     DecodingSource = 2
)

// TemplateFlags
type TemplateFlags types.UINT

const (
	TEMPLATE_EVENT_DATA TemplateFlags = 1
	TEMPLATE_USER_DATA  TemplateFlags = 2
)

// TDH_IN_TYPE
type TDH_IN_TYPE types.USHORT

const (
	TDH_INTYPE_NULL                        TDH_IN_TYPE = 0
	TDH_INTYPE_UNICODESTRING               TDH_IN_TYPE = 1
	TDH_INTYPE_ANSISTRING                  TDH_IN_TYPE = 2
	TDH_INTYPE_INT8                        TDH_IN_TYPE = 3
	TDH_INTYPE_UINT8                       TDH_IN_TYPE = 4
	TDH_INTYPE_INT16                       TDH_IN_TYPE = 5
	TDH_INTYPE_UINT16                      TDH_IN_TYPE = 6
	TDH_INTYPE_INT32                       TDH_IN_TYPE = 7
	TDH_INTYPE_UINT32                      TDH_IN_TYPE = 8
	TDH_INTYPE_INT64                       TDH_IN_TYPE = 9
	TDH_INTYPE_UINT64                      TDH_IN_TYPE = 10
	TDH_INTYPE_FLOAT                       TDH_IN_TYPE = 11
	TDH_INTYPE_DOUBLE                      TDH_IN_TYPE = 12
	TDH_INTYPE_BOOLEAN                     TDH_IN_TYPE = 13
	TDH_INTYPE_BINARY                      TDH_IN_TYPE = 14
	TDH_INTYPE_GUID                        TDH_IN_TYPE = 15
	TDH_INTYPE_POINTER                     TDH_IN_TYPE = 16
	TDH_INTYPE_FILETIME                    TDH_IN_TYPE = 17
	TDH_INTYPE_SYSTEMTIME                  TDH_IN_TYPE = 18
	TDH_INTYPE_SID                         TDH_IN_TYPE = 19
	TDH_INTYPE_HEXINT32                    TDH_IN_TYPE = 20
	TDH_INTYPE_HEXINT64                    TDH_IN_TYPE = 21
	TDH_INTYPE_COUNTEDSTRING               TDH_IN_TYPE = 300
	TDH_INTYPE_COUNTEDANSISTRING           TDH_IN_TYPE = 301
	TDH_INTYPE_REVERSEDCOUNTEDSTRING       TDH_IN_TYPE = 302
	TDH_INTYPE_REVERSEDCOUNTEDANSISTRING   TDH_IN_TYPE = 303
	TDH_INTYPE_NONNULLTERMINATEDSTRING     TDH_IN_TYPE = 304
	TDH_INTYPE_NONNULLTERMINATEDANSISTRING TDH_IN_TYPE = 305
	TDH_INTYPE_UNICODECHAR                 TDH_IN_TYPE = 306
	TDH_INTYPE_ANSICHAR                    TDH_IN_TYPE = 307
	TDH_INTYPE_SIZET                       TDH_IN_TYPE = 308
	TDH_INTYPE_HEXDUMP                     TDH_IN_TYPE = 309
	TDH_INTYPE_WBEMSID                     TDH_IN_TYPE = 310
)

// TDH_OUT_TYPE
type TDH_OUT_TYPE types.USHORT

const (
	TDH_OUTTYPE_NULL                         TDH_OUT_TYPE = 0
	TDH_OUTTYPE_STRING                       TDH_OUT_TYPE = 1
	TDH_OUTTYPE_DATETIME                     TDH_OUT_TYPE = 2
	TDH_OUTTYPE_BYTE                         TDH_OUT_TYPE = 3
	TDH_OUTTYPE_UNSIGNEDBYTE                 TDH_OUT_TYPE = 4
	TDH_OUTTYPE_SHORT                        TDH_OUT_TYPE = 5
	TDH_OUTTYPE_UNSIGNEDSHORT                TDH_OUT_TYPE = 6
	TDH_OUTTYPE_INT                          TDH_OUT_TYPE = 7
	TDH_OUTTYPE_UNSIGNEDINT                  TDH_OUT_TYPE = 8
	TDH_OUTTYPE_LONG                         TDH_OUT_TYPE = 9
	TDH_OUTTYPE_UNSIGNEDLONG                 TDH_OUT_TYPE = 10
	TDH_OUTTYPE_FLOAT                        TDH_OUT_TYPE = 11
	TDH_OUTTYPE_DOUBLE                       TDH_OUT_TYPE = 12
	TDH_OUTTYPE_BOOLEAN                      TDH_OUT_TYPE = 13
	TDH_OUTTYPE_GUID                         TDH_OUT_TYPE = 14
	TDH_OUTTYPE_HEXBINARY                    TDH_OUT_TYPE = 15
	TDH_OUTTYPE_HEXINT8                      TDH_OUT_TYPE = 16
	TDH_OUTTYPE_HEXINT16                     TDH_OUT_TYPE = 17
	TDH_OUTTYPE_HEXINT32                     TDH_OUT_TYPE = 18
	TDH_OUTTYPE_HEXINT64                     TDH_OUT_TYPE = 19
	TDH_OUTTYPE_PID                          TDH_OUT_TYPE = 20
	TDH_OUTTYPE_TID                          TDH_OUT_TYPE = 21
	TDH_OUTTYPE_PORT                         TDH_OUT_TYPE = 22
	TDH_OUTTYPE_IPV4                         TDH_OUT_TYPE = 23
	TDH_OUTTYPE_IPV6                         TDH_OUT_TYPE = 24
	TDH_OUTTYPE_SOCKETADDRESS                TDH_OUT_TYPE = 25
	TDH_OUTTYPE_CIMDATETIME                  TDH_OUT_TYPE = 26
	TDH_OUTTYPE_ETWTIME                      TDH_OUT_TYPE = 27
	TDH_OUTTYPE_XML                          TDH_OUT_TYPE = 28
	TDH_OUTTYPE_ERRORCODE                    TDH_OUT_TYPE = 29
	TDH_OUTTYPE_WIN32ERROR                   TDH_OUT_TYPE = 30
	TDH_OUTTYPE_NTSTATUS                     TDH_OUT_TYPE = 31
	TDH_OUTTYPE_HRESULT                      TDH_OUT_TYPE = 32
	TDH_OUTTYPE_CULTURE_INSENSITIVE_DATETIME TDH_OUT_TYPE = 33
	TDH_OUTTYPE_REDUCEDSTRING                TDH_OUT_TYPE = 300
	TDH_OUTTYPE_NOPRINT                      TDH_OUT_TYPE = 301
)

// PropertyFlags
type PropertyFlags types.UINT

const (
	PropertyStruct           PropertyFlags = 0x1
	PropertyParamLength      PropertyFlags = 0x2
	PropertyParamCount       PropertyFlags = 0x4
	PropertyWBEMXmlFragment  PropertyFlags = 0x8
	PropertyParamFixedLength PropertyFlags = 0x10
)

// EVENT_PROPERTY_INFO_u1
type StructType struct {
	StructStartIndex   types.USHORT
	NumOfStructMembers types.USHORT
	Padding            types.ULONG
}

// EventPropertyInfo
type EventPropertyInfo struct {
	Flags      PropertyFlags
	NameOffset types.ULONG
	StructType StructType
	Count      types.USHORT
	Length     types.USHORT
	Reserved   types.ULONG
}

// PROVIDER_FILTER_INFO
type PROVIDER_FILTER_INFO struct {
	Id                     types.UCHAR
	Version                types.UCHAR
	MessageOffset          types.ULONG
	Reserved               types.ULONG
	PropertyCount          types.ULONG
	EventPropertyInfoArray [1]EventPropertyInfo
}

type PPROVIDER_FILTER_INFO *PROVIDER_FILTER_INFO

// TraceEventInfo
type TraceEventInfo struct {
	ProviderGuid                types.GUID
	EventGuid                   types.GUID
	EventDescriptor             EventDescriptor
	DecodingSource              DecodingSource
	ProviderNameOffset          types.ULONG
	LevelNameOffset             types.ULONG
	ChannelNameOffset           types.ULONG
	KeywordsNameOffset          types.ULONG
	TaskNameOffset              types.ULONG
	OpcodeNameOffset            types.ULONG
	EventMessageOffset          types.ULONG
	ProviderMessageOffset       types.ULONG
	BinaryXMLOffset             types.ULONG
	BinaryXMLSize               types.ULONG
	ActivityIDNameOffset        types.ULONG
	RelatedActivityIDNameOffset types.ULONG
	PropertyCount               types.ULONG
	TopLevelPropertyCount       types.ULONG
	Flags                       TemplateFlags
	EventPropertyInfoArray      [1]EventPropertyInfo
}

type PTRACE_EVENT_INFO *TraceEventInfo

// MAP_VALUETYPE
type MAP_VALUETYPE types.UINT

const (
	EVENTMAP_ENTRY_VALUETYPE_ULONG  MAP_VALUETYPE = 0
	EVENTMAP_ENTRY_VALUETYPE_STRING MAP_VALUETYPE = 1
)

// MAP_FLAGS
type MAP_FLAGS types.UINT

const (
	EVENTMAP_INFO_FLAG_MANIFEST_VALUEMAP   MAP_FLAGS = 0x1
	EVENTMAP_INFO_FLAG_MANIFEST_BITMAP     MAP_FLAGS = 0x2
	EVENTMAP_INFO_FLAG_MANIFEST_PATTERNMAP MAP_FLAGS = 0x4
	EVENTMAP_INFO_FLAG_WBEM_VALUEMAP       MAP_FLAGS = 0x8
	EVENTMAP_INFO_FLAG_WBEM_BITMAP         MAP_FLAGS = 0x10
	EVENTMAP_INFO_FLAG_WBEM_FLAG           MAP_FLAGS = 0x20
	EVENTMAP_INFO_FLAG_WBEM_NO_MAP         MAP_FLAGS = 0x40
)

// EVENT_MAP_ENTRY_u
type EVENT_MAP_ENTRY_u struct {
	Value types.ULONG
}

// EVENT_MAP_ENTRY
type EVENT_MAP_ENTRY struct {
	OutputOffset types.ULONG
	U1           EVENT_MAP_ENTRY_u
}

// EVENT_MAP_INFO_u
type EVENT_MAP_INFO_u struct {
	MapEntryValueType MAP_VALUETYPE
}

// EVENT_MAP_INFO
type EVENT_MAP_INFO struct {
	NameOffset    types.ULONG
	Flag          MAP_FLAGS
	EntryCount    types.ULONG
	U1            EVENT_MAP_INFO_u
	MapEntryArray [1]EVENT_MAP_ENTRY
}

type PEVENT_MAP_INFO *EVENT_MAP_INFO

// EVENT_FIELD_TYPE
type EVENT_FIELD_TYPE types.UINT

const (
	EventKeywordInformation EVENT_FIELD_TYPE = 0
	EventLevelInformation   EVENT_FIELD_TYPE = 1
	EventChannelInformation EVENT_FIELD_TYPE = 2
	EventTaskInformation    EVENT_FIELD_TYPE = 3
	EventOpcodeInformation  EVENT_FIELD_TYPE = 4
)

// PROVIDER_FIELD_INFO
type PROVIDER_FIELD_INFO struct {
	NameOffset        types.ULONG
	DescriptionOffset types.ULONG
	Value             types.ULONGLONG
}

// PROVIDER_FIELD_INFOARRAY
type PROVIDER_FIELD_INFOARRAY struct {
	NumberOfElements types.ULONG
	FieldType        EVENT_FIELD_TYPE
	FieldInfoArray   [1]PROVIDER_FIELD_INFO
}

type PPROVIDER_FIELD_INFOARRAY *PROVIDER_FIELD_INFOARRAY

// TDH_CONTEXT_TYPE
type TDH_CONTEXT_TYPE types.UINT

const (
	TDH_CONTEXT_WPP_TMFFILE       TDH_CONTEXT_TYPE = 0
	TDH_CONTEXT_WPP_TMFSEARCHPATH TDH_CONTEXT_TYPE = 1
	TDH_CONTEXT_WPP_GMT           TDH_CONTEXT_TYPE = 2
	TDH_CONTEXT_POINTERSIZE       TDH_CONTEXT_TYPE = 3
)

// TDH_CONTEXT
type TDH_CONTEXT struct {
	ParameterValue types.ULONGLONG
	ParameterType  TDH_CONTEXT_TYPE
	ParameterSize  types.ULONG
}

type PTDH_CONTEXT *TDH_CONTEXT

// PropertyDataDescriptor
type PropertyDataDescriptor struct {
	PropertyName types.ULONGLONG
	ArrayIndex   types.ULONG
	Reserved     types.ULONG
}

// EVENT_HEADER_u_s
type EVENT_HEADER_u_s struct {
	KernelTime types.ULONG
	UserTime   types.ULONG
}

// EVENT_HEADER_u
type EVENT_HEADER_u struct {
	ProcessorTime types.ULONG64
}

// EventHeaderFlag
type EventHeaderFlag types.USHORT

const (
	EventHeaderFlagExtendedInfo   EventHeaderFlag = 0x0001
	EventHeaderFlagPrivateSession EventHeaderFlag = 0x0002
	EventHeaderFlagStringOnly     EventHeaderFlag = 0x0004
	EventHeaderFlagTraceMessage   EventHeaderFlag = 0x0008
	EventHeaderFlagNoCPUTime      EventHeaderFlag = 0x0010
	EventHeaderFlag32BitHeader    EventHeaderFlag = 0x0020
	EventHeaderFlag64BitHeader    EventHeaderFlag = 0x0040
	EventHeaderFlagClassicHeader  EventHeaderFlag = 0x0100
)

// EventHeaderProperty
type EventHeaderProperty types.USHORT

const (
	EventHeaderPropertyXML            EventHeaderProperty = 0x0001
	EventHeaderPropertyForwardedXML   EventHeaderProperty = 0x0002
	EventHeaderPropertyLegacyEventLog EventHeaderProperty = 0x0004
)

// EventDescriptor
type EventDescriptor struct {
	// ID represents event identifier.
	ID uint16
	// Version indicates a revision to the event definition.
	Version uint8
	// Channel is the audience for the event (e.g. administrator or developer).
	Channel uint8
	// Level is the severity or level of detail included in the event.
	Level uint8
	// Opcode is step in a sequence of operations being performed within the `Task` field. For MOF-defined events,
	// the `Opcode` member contains the event type value.
	Opcode uint8
	// Task represents a larger unit of work within an application or component.
	Task uint16
	// Keyword A bitmask that specifies a logical group of related events. Each bit corresponds to one group. An event may belong to one or more groups.
	// The keyword can contain one or more provider-defined keywords, standard keywords, or both.
	Keyword uint64
}

// EventHeader
type EventHeader struct {
	Size            types.USHORT
	HeaderType      types.USHORT
	Flags           EventHeaderFlag
	EventProperty   EventHeaderProperty
	ThreadId        types.ULONG
	ProcessId       types.ULONG
	TimeStamp       types.LARGE_INTEGER
	ProviderId      types.GUID
	EventDescriptor EventDescriptor
	ProcessorTime   types.ULONG64
	ActivityId      types.GUID
}

type BufferContext struct {
	ProcessorIndex [2]byte
	LoggerID       uint16
}

type Linkage struct {
	Linkage   uint16
	Resreved2 uint16
}

type EventHeaderExtendedDataItem struct {
	Reserved1 uint16
	ExtType   uint16
	Linkage
	DataSize uint16
	DataPtr  uint64
}

// EventRecord
type EventRecord struct {
	EventHeader       EventHeader
	BufferContext     BufferContext
	ExtendedDataCount types.USHORT
	UserDataLength    types.USHORT
	ExtendedData      *EventHeaderExtendedDataItem
	UserData          types.PVOID
	UserContext       types.PVOID
}
