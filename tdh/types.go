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

// TdhInType
type TdhInType types.USHORT

const (
	TdhIntypeNull                        TdhInType = 0
	TdhIntypeUnicodestring               TdhInType = 1
	TdhIntypeAnsiString                  TdhInType = 2
	TdhIntypeInt8                        TdhInType = 3
	TdhIntypeUint8                       TdhInType = 4
	TdhIntypeInt16                       TdhInType = 5
	TdhIntypeUint16                      TdhInType = 6
	TdhIntypeInt32                       TdhInType = 7
	TdhIntypeUint32                      TdhInType = 8
	TdhIntypeInt64                       TdhInType = 9
	TdhIntypeUint64                      TdhInType = 10
	TdhIntypeFloat                       TdhInType = 11
	TdhIntypeDouble                      TdhInType = 12
	TdhIntypeBoolean                     TdhInType = 13
	TdhIntypeBinary                      TdhInType = 14
	TdhIntypeGuid                        TdhInType = 15
	TdhIntypePointer                     TdhInType = 16
	TdhIntypeFileTime                    TdhInType = 17
	TdhIntypeSystemTime                  TdhInType = 18
	TdhIntypeSid                         TdhInType = 19
	TdhIntypeHexInt32                    TdhInType = 20
	TdhIntypeHexInt64                    TdhInType = 21
	TdhIntypeCountedString               TdhInType = 300
	TdhIntypeCountedAnsiString           TdhInType = 301
	TdhIntypeReversedCountedString       TdhInType = 302
	TdhIntypeReversedCountedAnsiString   TdhInType = 303
	TdhIntypeNonnullTerminatedString     TdhInType = 304
	TdhIntypeNonnullTerminatedAnsiString TdhInType = 305
	TdhIntypeUnicodeChar                 TdhInType = 306
	TdhIntypeAnsiChar                    TdhInType = 307
	TdhIntypeSizet                       TdhInType = 308
	TdhIntypeHexDump                     TdhInType = 309
	TdhIntypeWbemSid                     TdhInType = 310
)

// TdhOutType
type TdhOutType types.USHORT

const (
	TdhOutypeNull                       TdhOutType = 0
	TdhOutypeString                     TdhOutType = 1
	TdhOutypeDateTime                   TdhOutType = 2
	TdhOutypeByte                       TdhOutType = 3
	TdhOutypeUnsignedByte               TdhOutType = 4
	TdhOutypeShort                      TdhOutType = 5
	TdhOutypeUnsignedShort              TdhOutType = 6
	TdhOutypeInt                        TdhOutType = 7
	TdhOutypeUnsignedInt                TdhOutType = 8
	TdhOutypeLong                       TdhOutType = 9
	TdhOutypeUnsignedLong               TdhOutType = 10
	TdhOutypeFloat                      TdhOutType = 11
	TdhOutypeDouble                     TdhOutType = 12
	TdhOutypeBoolean                    TdhOutType = 13
	TdhOutypeGuid                       TdhOutType = 14
	TdhOutypeHexBinary                  TdhOutType = 15
	TdhOutypeHexInt8                    TdhOutType = 16
	TdhOutypeHexInt16                   TdhOutType = 17
	TdhOutypeHexInt32                   TdhOutType = 18
	TdhOutypeHexInt64                   TdhOutType = 19
	TdhOutypePid                        TdhOutType = 20
	TdhOutypeTid                        TdhOutType = 21
	TdhOutypePort                       TdhOutType = 22
	TdhOutypeIpv4                       TdhOutType = 23
	TdhOutypeIpv6                       TdhOutType = 24
	TdhOutypeSocketAddress              TdhOutType = 25
	TdhOutypeCIMDateTime                TdhOutType = 26
	TdhOutypeETWTime                    TdhOutType = 27
	TdhOutypeXml                        TdhOutType = 28
	TdhOutypeErrorCode                  TdhOutType = 29
	TdhOutypeWin32Error                 TdhOutType = 30
	TdhOutypeNtstatus                   TdhOutType = 31
	TdhOutypeHresult                    TdhOutType = 32
	TdhOutypeCultureInsensitiveDateTime TdhOutType = 33
	TdhOutypeReducedString              TdhOutType = 300
	TdhOutypeNoPrint                    TdhOutType = 301
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
	InType        TdhInType
	OutType       TdhOutType
	MapNameOffset types.ULONG
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

// TdhContextType
type TdhContextType types.UINT

const (
	TdhContextWPPTmfFile       TdhContextType = 0
	TdhContextWPPTmfSearchpath TdhContextType = 1
	TdhContextWPPGMT           TdhContextType = 2
	TdhContextPointersize      TdhContextType = 3
)

// TdhContext
type TdhContext struct {
	ParameterValue types.ULONGLONG
	ParameterType  TdhContextType
	ParameterSize  types.ULONG
}

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
