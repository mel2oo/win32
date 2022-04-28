package types

type ExceptionCode DWORD

const (
	EXCEPTION_ACCESS_VIOLATION         ExceptionCode = 0xC0000005
	EXCEPTION_ARRAY_BOUNDS_EXCEEDED    ExceptionCode = 0xC000008C
	EXCEPTION_BREAKPOINT               ExceptionCode = 0x80000003
	EXCEPTION_DATATYPE_MISALIGNMENT    ExceptionCode = 0xC00002C5
	EXCEPTION_FLT_DENORMAL_OPERAND     ExceptionCode = 0xC000008D
	EXCEPTION_FLT_DIVIDE_BY_ZERO       ExceptionCode = 0xC000008E
	EXCEPTION_FLT_INEXACT_RESULT       ExceptionCode = 0xC000008F
	EXCEPTION_FLT_INVALID_OPERATION    ExceptionCode = 0xC0000090
	EXCEPTION_FLT_OVERFLOW             ExceptionCode = 0xC0000091
	EXCEPTION_FLT_STACK_CHECK          ExceptionCode = 0xC0000092
	EXCEPTION_FLT_UNDERFLOW            ExceptionCode = 0xC0000093
	EXCEPTION_GUARD_PAGE               ExceptionCode = 0x80000001
	EXCEPTION_ILLEGAL_INSTRUCTION      ExceptionCode = 0xC000001D
	EXCEPTION_IN_PAGE_ERROR            ExceptionCode = 0xC0000006
	EXCEPTION_INT_DIVIDE_BY_ZERO       ExceptionCode = 0xC0000094
	EXCEPTION_INT_OVERFLOW             ExceptionCode = 0xC0000095
	EXCEPTION_INVALID_DISPOSITION      ExceptionCode = 0xC0000026
	EXCEPTION_INVALID_HANDLE           ExceptionCode = 0xC0000008
	EXCEPTION_NONCONTINUABLE_EXCEPTION ExceptionCode = 0xC0000025
	EXCEPTION_PRIV_INSTRUCTION         ExceptionCode = 0xC0000096
	EXCEPTION_SINGLE_STEP              ExceptionCode = 0x80000004
	EXCEPTION_STACK_OVERFLOW           ExceptionCode = 0xC00000FD
)
