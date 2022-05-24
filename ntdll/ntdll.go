package ntdll

import "golang.org/x/sys/windows"

var (
	dll = windows.NewLazyDLL("ntdll.dll")

	ProcRtlEthernetAddressToString = dll.NewProc("RtlEthernetAddressToStringW")
	ProcRtlEthernetStringToAddress = dll.NewProc("RtlEthernetStringToAddressW")
	ProcRtlIpv4AddressToStringEx   = dll.NewProc("RtlIpv4AddressToStringExW")
	ProcRtlIpv4AddressToString     = dll.NewProc("RtlIpv4AddressToStringW")
	ProcRtlIpv4StringToAddressEx   = dll.NewProc("RtlIpv4StringToAddressExW")
	ProcRtlIpv4StringToAddress     = dll.NewProc("RtlIpv4StringToAddressW")
	ProcRtlIpv6AddressToStringEx   = dll.NewProc("RtlIpv6AddressToStringExW")
	ProcRtlIpv6AddressToString     = dll.NewProc("RtlIpv6AddressToStringW")
	ProcRtlIpv6StringToAddressEx   = dll.NewProc("RtlIpv6StringToAddressExW")
	ProcRtlIpv6StringToAddress     = dll.NewProc("RtlIpv6StringToAddressW")
)
