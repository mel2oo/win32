package tdh

import "github.com/mel2oo/win32/types"

// Variables
type (
	TDH_HANDLE  types.HANDLE
	PTDH_HANDLE *TDH_HANDLE
)

type PEVENT_HEADER_EXTENDED_DATA_ITEM types.LPVOID

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

// DECODING_SOURCE
type DECODING_SOURCE types.UINT

const (
	DecodingSourceXMLFile DECODING_SOURCE = 0
	DecodingSourceWbem    DECODING_SOURCE = 1
	DecodingSourceWPP     DECODING_SOURCE = 2
)

// TEMPLATE_FLAGS
type TEMPLATE_FLAGS types.UINT

const (
	TEMPLATE_EVENT_DATA TEMPLATE_FLAGS = 1
	TEMPLATE_USER_DATA  TEMPLATE_FLAGS = 2
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

// PROPERTY_FLAGS
type PROPERTY_FLAGS types.UINT

const (
	PropertyStruct           PROPERTY_FLAGS = 0x1
	PropertyParamLength      PROPERTY_FLAGS = 0x2
	PropertyParamCount       PROPERTY_FLAGS = 0x4
	PropertyWBEMXmlFragment  PROPERTY_FLAGS = 0x8
	PropertyParamFixedLength PROPERTY_FLAGS = 0x10
)

// EVENT_PROPERTY_INFO_u1_s1
type EVENT_PROPERTY_INFO_u1_s1 struct {
	InType        TDH_IN_TYPE
	OutType       TDH_OUT_TYPE
	MapNameOffset types.ULONG
}

// EVENT_PROPERTY_INFO_u1_s2
type EVENT_PROPERTY_INFO_u1_s2 struct {
	StructStartIndex   types.USHORT
	NumOfStructMembers types.USHORT
	Padding            types.ULONG
}

// EVENT_PROPERTY_INFO_u1
type EVENT_PROPERTY_INFO_u1 struct {
	StructType EVENT_PROPERTY_INFO_u1_s2
}

// EVENT_PROPERTY_INFO_u2
type EVENT_PROPERTY_INFO_u2 struct {
	Count types.USHORT
}

// EVENT_PROPERTY_INFO_u3
type EVENT_PROPERTY_INFO_u3 struct {
	Length types.USHORT
}

// EVENT_PROPERTY_INFO
type EVENT_PROPERTY_INFO struct {
	Flags      PROPERTY_FLAGS
	NameOffset types.ULONG
	U1         EVENT_PROPERTY_INFO_u1
	U2         EVENT_PROPERTY_INFO_u2
	U3         EVENT_PROPERTY_INFO_u3
	Reserved   types.ULONG
}

// PROVIDER_FILTER_INFO
type PROVIDER_FILTER_INFO struct {
	Id                     types.UCHAR
	Version                types.UCHAR
	MessageOffset          types.ULONG
	Reserved               types.ULONG
	PropertyCount          types.ULONG
	EventPropertyInfoArray [1]EVENT_PROPERTY_INFO
}

type PPROVIDER_FILTER_INFO *PROVIDER_FILTER_INFO

// TRACE_EVENT_INFO
type TRACE_EVENT_INFO struct {
	ProviderGuid                types.GUID
	EventGuid                   types.GUID
	EventDescriptor             types.LPVOID
	DecodingSource              DECODING_SOURCE
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
	Flags                       TEMPLATE_FLAGS
	EventPropertyInfoArray      [1]EVENT_PROPERTY_INFO
}

type PTRACE_EVENT_INFO *TRACE_EVENT_INFO

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

// PROPERTY_DATA_DESCRIPTOR
type PROPERTY_DATA_DESCRIPTOR struct {
	PropertyName types.ULONGLONG
	ArrayIndex   types.ULONG
	Reserved     types.ULONG
}

type PPROPERTY_DATA_DESCRIPTOR *PROPERTY_DATA_DESCRIPTOR

// EVENT_HEADER_u_s
type EVENT_HEADER_u_s struct {
	KernelTime types.ULONG
	UserTime   types.ULONG
}

// EVENT_HEADER_u
type EVENT_HEADER_u struct {
	ProcessorTime types.ULONG64
}

// EVENT_HEADER_FLAG
type EVENT_HEADER_FLAG types.USHORT

const (
	EVENT_HEADER_FLAG_EXTENDED_INFO   EVENT_HEADER_FLAG = 0x0001
	EVENT_HEADER_FLAG_PRIVATE_SESSION EVENT_HEADER_FLAG = 0x0002
	EVENT_HEADER_FLAG_STRING_ONLY     EVENT_HEADER_FLAG = 0x0004
	EVENT_HEADER_FLAG_TRACE_MESSAGE   EVENT_HEADER_FLAG = 0x0008
	EVENT_HEADER_FLAG_NO_CPUTIME      EVENT_HEADER_FLAG = 0x0010
	EVENT_HEADER_FLAG_32_BIT_HEADER   EVENT_HEADER_FLAG = 0x0020
	EVENT_HEADER_FLAG_64_BIT_HEADER   EVENT_HEADER_FLAG = 0x0040
	EVENT_HEADER_FLAG_CLASSIC_HEADER  EVENT_HEADER_FLAG = 0x0100
)

// EVENT_HEADER_PROPERTY
type EVENT_HEADER_PROPERTY types.USHORT

const (
	EVENT_HEADER_PROPERTY_XML             EVENT_HEADER_PROPERTY = 0x0001
	EVENT_HEADER_PROPERTY_FORWARDED_XML   EVENT_HEADER_PROPERTY = 0x0002
	EVENT_HEADER_PROPERTY_LEGACY_EVENTLOG EVENT_HEADER_PROPERTY = 0x0004
)

// EVENT_HEADER
type EVENT_HEADER struct {
	Size            types.USHORT
	HeaderType      types.USHORT
	Flags           EVENT_HEADER_FLAG
	EventProperty   EVENT_HEADER_PROPERTY
	ThreadId        types.ULONG
	ProcessId       types.ULONG
	TimeStamp       types.LARGE_INTEGER
	ProviderId      types.GUID
	EventDescriptor types.LPVOID
	U1              EVENT_HEADER_u
	ActivityId      types.GUID
}

// EVENT_RECORD
type EVENT_RECORD struct {
	EventHeader       EVENT_HEADER
	BufferContext     types.PVOID
	ExtendedDataCount types.USHORT
	UserDataLength    types.USHORT
	ExtendedData      PEVENT_HEADER_EXTENDED_DATA_ITEM
	UserData          types.PVOID
	UserContext       types.PVOID
}

type PEVENT_RECORD *EVENT_RECORD
