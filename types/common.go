// From MSDN: Windows Data Types
// http://msdn.microsoft.com/en-us/library/s3f49ktz.aspx
// http://msdn.microsoft.com/en-us/library/windows/desktop/aa383751.aspx

package types

import (
	"golang.org/x/sys/windows"
)

type (
	// UINT_PTR and INT_PTR
	INT_PTR  uintptr
	UINT_PTR uintptr

	// HMODULE todo
	HMODULE windows.Handle

	// void*
	LPVOID  uintptr
	LPCVOID LPVOID
	PVOID   LPVOID

	// PVOID64
	PVOID64 UINT64

	// byte
	BYTE   byte
	LPBYTE *BYTE
	PBYTE  *BYTE

	// 8-bit unsigned integer
	UCHAR  uint8
	PUCHAR *UCHAR
	UINT8  uint8
	PUINT8 *UINT8

	// 8-bit signed integer
	INT8  int8
	INT16 int16

	// 16-bit unsigned integer
	UINT16  uint16
	WORD    UINT16
	PWORD   *WORD
	LPWORD  PWORD
	USHORT  UINT16
	PUSHORT *UINT16

	// 16-bit signed integer
	SHORT int16

	// 32/64 bit unsigned integer
	PUINT_PTR  *UINT_PTR
	ULONG_PTR  UINT_PTR
	PULONG_PTR *ULONG_PTR
	DWORD_PTR  ULONG_PTR
	PDWORD_PTR *DWORD_PTR

	// 32/64 bit signed integer
	LONG_PTR INT_PTR

	// BOOL
	BOOL   int32
	PBOOL  *BOOL
	LPBOOL PBOOL

	// BOOLEAN
	BOOLEAN  byte
	PBOOLEAN *BOOLEAN

	// 32-bit unsigned integer
	UINT32  uint32
	PUINT32 *UINT32
	ULONG   UINT32
	PULONG  *ULONG
	UINT    UINT32
	PUINT   *UINT
	LPUINT  PUINT
	ULONG32 ULONG

	// 32-bit signed integer
	INT32  int32
	LONG   INT32
	PLONG  *LONG
	LPLONG *LONG
	INT    INT32
	PINT   *INT
	LPINT  PINT

	// 64-bit integers
	INT64     int64
	LONGLONG  INT64
	PLONGLONG *LONGLONG
	LONG64    INT64
	PLONG64   *LONG64

	// 64-bit unsigned integers
	UINT64     uint64
	PUINT64    *UINT64
	ULONGLONG  UINT64
	PULONGLONG *ULONGLONG
	ULONGULONG UINT64
	ULONG64    UINT64
	PULONG64   *ULONG64
	DWORD64    UINT64
	PDWORD64   *DWORD64
	DWORDLONG  ULONGLONG
	PDWORDLONG *DWORDLONG

	// DWORD
	DWORD   uint32
	PDWORD  *DWORD
	LPDWORD PDWORD

	// LPSTR
	CHAR   int8
	LPSTR  *CHAR
	LPCSTR LPSTR
	PCSTR  LPSTR
	PSTR   LPSTR
	PCHAR  LPSTR
	LPCH   LPSTR

	// LPWSTR
	WCHAR   uint16
	LPWSTR  *uint16
	PWSTR   LPWSTR
	PCWSTR  PWSTR
	LPCWSTR PWSTR
	PWCHAR  PWSTR

	// size_t
	SIZE_T  INT_PTR
	PSIZE_T *SIZE_T

	// time_t
	TIME_T INT64

	// float
	FLOAT   float32
	PFLOAT  *FLOAT
	FLOAT32 FLOAT

	// double
	DOUBLE float64

	// va_list
	VA_LIST LPVOID
)
