package types

// ALG_ID
type ALG_ID UINT32

const (
	CALG_MD2                  ALG_ID = 0x00008001
	CALG_MD4                  ALG_ID = 0x00008002
	CALG_MD5                  ALG_ID = 0x00008003
	CALG_SHA                  ALG_ID = 0x00008004
	CALG_SHA1                 ALG_ID = 0x00008004
	CALG_MAC                  ALG_ID = 0x00008005
	CALG_RSA_SIGN             ALG_ID = 0x00002400
	CALG_DSS_SIGN             ALG_ID = 0x00002200
	CALG_NO_SIGN              ALG_ID = 0x00002000
	CALG_RSA_KEYX             ALG_ID = 0x0000a400
	CALG_DES                  ALG_ID = 0x00006601
	CALG_3DES_112             ALG_ID = 0x00006609
	CALG_3DES                 ALG_ID = 0x00006603
	CALG_DESX                 ALG_ID = 0x00006604
	CALG_RC2                  ALG_ID = 0x00006602
	CALG_RC4                  ALG_ID = 0x00006801
	CALG_SEAL                 ALG_ID = 0x00006802
	CALG_DH_SF                ALG_ID = 0x0000aa01
	CALG_DH_EPHEM             ALG_ID = 0x0000aa02
	CALG_AGREEDKEY_ANY        ALG_ID = 0x0000aa03
	CALG_KEA_KEYX             ALG_ID = 0x0000aa04
	CALG_HUGHES_MD5           ALG_ID = 0x0000a003
	CALG_SKIPJACK             ALG_ID = 0x0000660a
	CALG_TEK                  ALG_ID = 0x0000660b
	CALG_CYLINK_MEK           ALG_ID = 0x0000660c
	CALG_SSL3_SHAMD5          ALG_ID = 0x00008008
	CALG_SSL3_MASTER          ALG_ID = 0x00004c01
	CALG_SCHANNEL_MASTER_HASH ALG_ID = 0x00004c02
	CALG_SCHANNEL_MAC_KEY     ALG_ID = 0x00004c03
	CALG_SCHANNEL_ENC_KEY     ALG_ID = 0x00004c07
	CALG_PCT1_MASTER          ALG_ID = 0x00004c04
	CALG_SSL2_MASTER          ALG_ID = 0x00004c05
	CALG_TLS1_MASTER          ALG_ID = 0x00004c06
	CALG_RC5                  ALG_ID = 0x0000660d
	CALG_HMAC                 ALG_ID = 0x00008009
	CALG_TLS1PRF              ALG_ID = 0x0000800a
	CALG_HASH_REPLACE_OWF     ALG_ID = 0x0000800b
	CALG_AES_128              ALG_ID = 0x0000660e
	CALG_AES_192              ALG_ID = 0x0000660f
	CALG_AES_256              ALG_ID = 0x00006610
	CALG_AES                  ALG_ID = 0x00006611
	CALG_SHA_256              ALG_ID = 0x0000800c
	CALG_SHA_384              ALG_ID = 0x0000800d
	CALG_SHA_512              ALG_ID = 0x0000800e
	CALG_ECDH                 ALG_ID = 0x0000aa05
	CALG_ECMQV                ALG_ID = 0x0000a001
	CALG_ECDSA                ALG_ID = 0x00002203
)

// WTD_UI
type WTD_UI DWORD

const (
	WTD_UI_ALL    WTD_UI = 1
	WTD_UI_NONE   WTD_UI = 2
	WTD_UI_NOBAD  WTD_UI = 3
	WTD_UI_NOGOOD WTD_UI = 4
)