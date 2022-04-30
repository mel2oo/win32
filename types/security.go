package types

// PSEC_WINNT_AUTH_IDENTITY_OPAQUE
type PSEC_WINNT_AUTH_IDENTITY_OPAQUE LPVOID

// SECURITY_STATUS
type (
	SECURITY_STATUS    HRESULT
	SAFER_LEVEL_HANDLE HANDLE
)

// ISecurityInformation
type (
	ISecurityInformation LPVOID
	LPSECURITYINFO       ISecurityInformation
)

// AceFlags
type AceFlags DWORD

const (
	OBJECT_INHERIT_ACE                 AceFlags = 0x1
	CONTAINER_INHERIT_ACE              AceFlags = 0x2
	SUB_CONTAINERS_AND_OBJECTS_INHERIT AceFlags = 0x3
	NO_PROPAGATE_INHERIT_ACE           AceFlags = 0x4
	INHERIT_ONLY_ACE                   AceFlags = 0x8
	INHERITED_ACE                      AceFlags = 0x10
)

//ACCESS_MODE
type ACCESS_MODE UINT

const (
	NOT_USED_ACCESS   ACCESS_MODE = 0
	GRANT_ACCESS      ACCESS_MODE = 1
	SET_ACCESS        ACCESS_MODE = 2
	DENY_ACCESS       ACCESS_MODE = 3
	REVOKE_ACCESS     ACCESS_MODE = 4
	SET_AUDIT_SUCCESS ACCESS_MODE = 5
	SET_AUDIT_FAILURE ACCESS_MODE = 6
)

// MULTIPLE_TRUSTEE_OPERATION
type MULTIPLE_TRUSTEE_OPERATION UINT

const (
	NO_MULTIPLE_TRUSTEE    MULTIPLE_TRUSTEE_OPERATION = 0
	TRUSTEE_IS_IMPERSONATE MULTIPLE_TRUSTEE_OPERATION = 1
)

// TRUSTEE_FORM
type TRUSTEE_FORM UINT

const (
	TRUSTEE_IS_SID              TRUSTEE_FORM = 0
	TRUSTEE_IS_NAME             TRUSTEE_FORM = 1
	TRUSTEE_BAD_FORM            TRUSTEE_FORM = 2
	TRUSTEE_IS_OBJECTS_AND_SID  TRUSTEE_FORM = 3
	TRUSTEE_IS_OBJECTS_AND_NAME TRUSTEE_FORM = 4
)

// TRUSTEE_TYPE
type TRUSTEE_TYPE UINT

const (
	TRUSTEE_IS_UNKNOWN          TRUSTEE_TYPE = 0
	TRUSTEE_IS_USER             TRUSTEE_TYPE = 1
	TRUSTEE_IS_GROUP            TRUSTEE_TYPE = 2
	TRUSTEE_IS_DOMAIN           TRUSTEE_TYPE = 3
	TRUSTEE_IS_ALIAS            TRUSTEE_TYPE = 4
	TRUSTEE_IS_WELL_KNOWN_GROUP TRUSTEE_TYPE = 5
	TRUSTEE_IS_DELETED          TRUSTEE_TYPE = 6
	TRUSTEE_IS_INVALID          TRUSTEE_TYPE = 7
	TRUSTEE_IS_COMPUTER         TRUSTEE_TYPE = 8
)

// OBJECTS_AND_SID
type OBJECTS_AND_SID struct {
	ObjectsPresent          DWORD
	ObjectTypeGuid          GUID
	InheritedObjectTypeGuid GUID
	Sid                     PSID
}
type POBJECTS_AND_SID *OBJECTS_AND_SID

// SE_OBJECT_TYPE
type SE_OBJECT_TYPE UINT

const (
	SE_UNKNOWN_OBJECT_TYPE     SE_OBJECT_TYPE = 0
	SE_FILE_OBJECT             SE_OBJECT_TYPE = 1
	SE_SERVICE                 SE_OBJECT_TYPE = 2
	SE_PRINTER                 SE_OBJECT_TYPE = 3
	SE_REGISTRY_KEY            SE_OBJECT_TYPE = 4
	SE_LMSHARE                 SE_OBJECT_TYPE = 5
	SE_KERNEL_OBJECT           SE_OBJECT_TYPE = 6
	SE_WINDOW_OBJECT           SE_OBJECT_TYPE = 7
	SE_DS_OBJECT               SE_OBJECT_TYPE = 8
	SE_DS_OBJECT_ALL           SE_OBJECT_TYPE = 9
	SE_PROVIDER_DEFINED_OBJECT SE_OBJECT_TYPE = 10
	SE_WMIGUID_OBJECT          SE_OBJECT_TYPE = 11
	SE_REGISTRY_WOW64_32KEY    SE_OBJECT_TYPE = 12
)

// OBJECTS_AND_NAME
type OBJECTS_AND_NAME struct {
	ObjectsPresent          DWORD
	ObjectType              SE_OBJECT_TYPE
	ObjectTypeName          LPWSTR
	InheritedObjectTypeName LPWSTR
	Name                    LPWSTR
}
type POBJECTS_AND_NAME *OBJECTS_AND_NAME

// TRUSTEE_u
type TRUSTEE_u struct {
	Name           LPWSTR
	Sid            PSID
	ObjectsAndSid  *OBJECTS_AND_SID
	ObjectsAndName *OBJECTS_AND_NAME
}

// TRUSTEE
type TRUSTEE struct {
	MultipleTrustee          LPVOID
	MultipleTrusteeOperation MULTIPLE_TRUSTEE_OPERATION
	TrusteeForm              TRUSTEE_FORM
	TrusteeType              TRUSTEE_TYPE
	TRUSTEE                  TRUSTEE_u
}
type PTRUSTEE *TRUSTEE

// EXPLICIT_ACCESS
type EXPLICIT_ACCESS struct {
	AccessPermissions ACCESS_MASK_DWORD
	AccessMode        ACCESS_MODE
	Inheritance       AceFlags
	Trustee           TRUSTEE
}
type PEXPLICIT_ACCESS *EXPLICIT_ACCESS

// TOKEN_SOURCE
type TOKEN_SOURCE struct {
	SourceName       [8]CHAR
	SourceIdentifier LUID
}
type PTOKEN_SOURCE *TOKEN_SOURCE

// SAFER_CRITERIA
type SAFER_CRITERIA DWORD

const (
	SAFER_CRITERIA_IMAGEPATH    SAFER_CRITERIA = 0x00001
	SAFER_CRITERIA_NOSIGNEDHASH SAFER_CRITERIA = 0x00002
	SAFER_CRITERIA_IMAGEHASH    SAFER_CRITERIA = 0x00004
	SAFER_CRITERIA_AUTHENTICODE SAFER_CRITERIA = 0x00008
	SAFER_CRITERIA_URLZONE      SAFER_CRITERIA = 0x00010
	SAFER_CRITERIA_IMAGEPATH_NT SAFER_CRITERIA = 0x01000
)

// SAFER_CODE_PROPERTIES
type SAFER_CODE_PROPERTIES struct {
	Size               DWORD
	CheckFlags         SAFER_CRITERIA
	ImagePath          LPCWSTR
	ImageFileHandle    HANDLE
	UrlZoneId          URLZONE
	ImageHash          [64]BYTE
	ImageHashSize      DWORD
	ImageSize          LARGE_INTEGER
	HashAlgorithm      ALG_ID
	ByteBlock          LPBYTE
	WndParent          HWND
	WVTUIChoice        WTD_UI
	PackageMoniker     LPCWSTR
	PackagePublisher   LPCWSTR
	PackageName        LPCWSTR
	PackageVersion     ULONG64
	PackageIsFramework BOOL
}
type PSAFER_CODE_PROPERTIES *SAFER_CODE_PROPERTIES

// SID_IDENTIFIER_AUTHORITY
type SID_IDENTIFIER_AUTHORITY struct {
	Value [6]BYTE
}
type PSID_IDENTIFIER_AUTHORITY *SID_IDENTIFIER_AUTHORITY

// AUDIT_POLICY_INFORMATION
type AUDIT_POLICY_INFORMATION struct {
	AuditSubCategoryGuid GUID
	AuditingInformation  ULONG
	AuditCategoryGuid    GUID
}
type PCAUDIT_POLICY_INFORMATION *AUDIT_POLICY_INFORMATION

// AceType
type AceType BYTE

const (
	ACCESS_ALLOWED_ACE_TYPE                 AceType = 0x0
	ACCESS_DENIED_ACE_TYPE                  AceType = 0x1
	SYSTEM_AUDIT_ACE_TYPE                   AceType = 0x2
	SYSTEM_ALARM_ACE_TYPE                   AceType = 0x3
	ACCESS_ALLOWED_COMPOUND_ACE_TYPE        AceType = 0x4
	ACCESS_ALLOWED_OBJECT_ACE_TYPE          AceType = 0x5
	ACCESS_DENIED_OBJECT_ACE_TYPE           AceType = 0x6
	SYSTEM_AUDIT_OBJECT_ACE_TYPE            AceType = 0x7
	SYSTEM_ALARM_OBJECT_ACE_TYPE            AceType = 0x8
	ACCESS_ALLOWED_CALLBACK_ACE_TYPE        AceType = 0x9
	ACCESS_DENIED_CALLBACK_ACE_TYPE         AceType = 0xA
	ACCESS_ALLOWED_CALLBACK_OBJECT_ACE_TYPE AceType = 0xB
	ACCESS_DENIED_CALLBACK_OBJECT_ACE_TYPE  AceType = 0xC
	SYSTEM_AUDIT_CALLBACK_ACE_TYPE          AceType = 0xD
	SYSTEM_ALARM_CALLBACK_ACE_TYPE          AceType = 0xE
	SYSTEM_AUDIT_CALLBACK_OBJECT_ACE_TYPE   AceType = 0xF
	SYSTEM_ALARM_CALLBACK_OBJECT_ACE_TYPE   AceType = 0x10
	SYSTEM_MANDATORY_LABEL_ACE_TYPE         AceType = 0x11
	SYSTEM_RESOURCE_ATTRIBUTE_ACE_TYPE      AceType = 0x12
	SYSTEM_SCOPED_POLICY_ID_ACE_TYPE        AceType = 0x13
)

// AceFlagsByte
type AceFlagsByte BYTE

// ACE_HEADER
type ACE_HEADER struct {
	AceType  AceType
	AceFlags AceFlagsByte
	AceSize  WORD
}

// ACE
type ACE struct {
	Header     ACE_HEADER
	AccessMask ACCESS_MASK
}
type PACE *ACE

// GENERIC_MAPPING
type GENERIC_MAPPING struct {
	GenericRead    ACCESS_MASK
	GenericWrite   ACCESS_MASK
	GenericExecute ACCESS_MASK
	GenericAll     ACCESS_MASK
}
type PGENERIC_MAPPING *GENERIC_MAPPING

// LUID_AND_ATTRIBUTES
type LUID_AND_ATTRIBUTES struct {
	Luid       LUID
	Attributes DWORD
}
type PLUID_AND_ATTRIBUTES *LUID_AND_ATTRIBUTES

// PRIVILEGE_SET_ALL_NECESSARY
type PRIVILEGE_SET DWORD

const (
	PRIVILEGE_SET_ALL_NECESSARY PRIVILEGE_SET = 1
)

type PPRIVILEGE_SET *PRIVILEGE_SET

// TOKEN_PRIVILEGES
type TOKEN_PRIVILEGES struct {
	PrivilegeCount DWORD
	Privileges     [1]LUID_AND_ATTRIBUTES
}
type PTOKEN_PRIVILEGES *TOKEN_PRIVILEGES

// ACL
type ACL struct {
	AclRevision BYTE
	Sbz1        BYTE
	AclSize     WORD
	AceCount    WORD
	Sbz2        WORD
}
type PACL *ACL

// CREDENTIAL_ATTRIBUTE
type CREDENTIAL_ATTRIBUTE struct {
	Keyword   LPWSTR
	Flags     DWORD
	ValueSize DWORD
	Value     LPBYTE
}
type PCREDENTIAL_ATTRIBUTE *CREDENTIAL_ATTRIBUTE

// CREDENTIAL
type CREDENTIAL struct {
	Flags              DWORD
	Type               DWORD
	TargetName         LPWSTR
	Comment            LPWSTR
	LastWritten        FILETIME
	CredentialBlobSize DWORD
	CredentialBlob     LPBYTE
	Persist            DWORD
	AttributeCount     DWORD
	Attributes         PCREDENTIAL_ATTRIBUTE
	TargetAlias        LPWSTR
	UserName           LPWSTR
}
type PCREDENTIAL *CREDENTIAL

// CREDENTIAL_TARGET_INFORMATION
type CREDENTIAL_TARGET_INFORMATION struct {
	TargetName        LPWSTR
	NetbiosServerName LPWSTR
	DnsServerName     LPWSTR
	NetbiosDomainName LPWSTR
	DnsDomainName     LPWSTR
	DnsTreeName       LPWSTR
	PackageName       LPWSTR
	Flags             ULONG
	CredTypeCount     DWORD
	CredTypes         LPDWORD
}
type PCREDENTIAL_TARGET_INFORMATION *CREDENTIAL_TARGET_INFORMATION

// POLICY_AUDIT_SID_ARRAY
type POLICY_AUDIT_SID_ARRAY struct {
	UsersCount   ULONG
	UserSidArray *PSID
}
type PPOLICY_AUDIT_SID_ARRAY *POLICY_AUDIT_SID_ARRAY

// QUOTA_LIMITS
type QUOTA_LIMITS struct {
	PagedPoolLimit        SIZE_T
	NonPagedPoolLimit     SIZE_T
	MinimumWorkingSetSize SIZE_T
	MaximumWorkingSetSize SIZE_T
	PagefileLimit         SIZE_T
	TimeLimit             LARGE_INTEGER
}
type PQUOTA_LIMITS *QUOTA_LIMITS

// SecHandle
type SecHandle struct {
	DwLower ULONG_PTR
	DwUpper ULONG_PTR
}
type PCtxtHandle *SecHandle

// SECURITY_DESCRIPTOR_CONTROL
type SECURITY_DESCRIPTOR_CONTROL WORD

const (
	SE_OWNER_DEFAULTED       SECURITY_DESCRIPTOR_CONTROL = 0x0001
	SE_GROUP_DEFAULTED       SECURITY_DESCRIPTOR_CONTROL = 0x0002
	SE_DACL_PRESENT          SECURITY_DESCRIPTOR_CONTROL = 0x0004
	SE_DACL_DEFAULTED        SECURITY_DESCRIPTOR_CONTROL = 0x0008
	SE_SACL_PRESENT          SECURITY_DESCRIPTOR_CONTROL = 0x0010
	SE_SACL_DEFAULTED        SECURITY_DESCRIPTOR_CONTROL = 0x0020
	SE_DACL_AUTO_INHERIT_REQ SECURITY_DESCRIPTOR_CONTROL = 0x0100
	SE_SACL_AUTO_INHERIT_REQ SECURITY_DESCRIPTOR_CONTROL = 0x0200
	SE_DACL_AUTO_INHERITED   SECURITY_DESCRIPTOR_CONTROL = 0x0400
	SE_SACL_AUTO_INHERITED   SECURITY_DESCRIPTOR_CONTROL = 0x0800
	SE_DACL_PROTECTED        SECURITY_DESCRIPTOR_CONTROL = 0x1000
	SE_SACL_PROTECTED        SECURITY_DESCRIPTOR_CONTROL = 0x2000
	SE_RM_CONTROL_VALID      SECURITY_DESCRIPTOR_CONTROL = 0x4000
	SE_SELF_RELATIVE         SECURITY_DESCRIPTOR_CONTROL = 0x8000
)

type PSECURITY_DESCRIPTOR_CONTROL *SECURITY_DESCRIPTOR_CONTROL

// TOKEN_TYPE
type TOKEN_TYPE UINT

const (
	TokenPrimary       TOKEN_TYPE = 1
	TokenImpersonation TOKEN_TYPE = 2
)

// TOKEN_INFORMATION_CLASS
type TOKEN_INFORMATION_CLASS UINT

const (
	TokenUser                            TOKEN_INFORMATION_CLASS = 1
	TokenGroups                          TOKEN_INFORMATION_CLASS = 2
	TokenPrivileges                      TOKEN_INFORMATION_CLASS = 3
	TokenOwner                           TOKEN_INFORMATION_CLASS = 4
	TokenPrimaryGroup                    TOKEN_INFORMATION_CLASS = 5
	TokenDefaultDacl                     TOKEN_INFORMATION_CLASS = 6
	TokenSource                          TOKEN_INFORMATION_CLASS = 7
	TokenType                            TOKEN_INFORMATION_CLASS = 8
	TokenImpersonationLevel              TOKEN_INFORMATION_CLASS = 9
	TokenStatistics                      TOKEN_INFORMATION_CLASS = 10
	TokenRestrictedSids                  TOKEN_INFORMATION_CLASS = 11
	TokenSessionId                       TOKEN_INFORMATION_CLASS = 12
	TokenGroupsAndPrivileges             TOKEN_INFORMATION_CLASS = 13
	TokenSessionReference                TOKEN_INFORMATION_CLASS = 14
	TokenSandBoxInert                    TOKEN_INFORMATION_CLASS = 15
	TokenAuditPolicy                     TOKEN_INFORMATION_CLASS = 16
	TokenOrigin                          TOKEN_INFORMATION_CLASS = 17
	TokenElevationType                   TOKEN_INFORMATION_CLASS = 18
	TokenLinkedToken                     TOKEN_INFORMATION_CLASS = 19
	TokenElevation                       TOKEN_INFORMATION_CLASS = 20
	TokenHasRestrictions                 TOKEN_INFORMATION_CLASS = 21
	TokenAccessInformation               TOKEN_INFORMATION_CLASS = 22
	TokenVirtualizationAllowed           TOKEN_INFORMATION_CLASS = 23
	TokenVirtualizationEnabled           TOKEN_INFORMATION_CLASS = 24
	TokenIntegrityLevel                  TOKEN_INFORMATION_CLASS = 25
	TokenUIAccess                        TOKEN_INFORMATION_CLASS = 26
	TokenMandatoryPolicy                 TOKEN_INFORMATION_CLASS = 27
	TokenLogonSid                        TOKEN_INFORMATION_CLASS = 28
	TokenIsAppContainer                  TOKEN_INFORMATION_CLASS = 29
	TokenCapabilities                    TOKEN_INFORMATION_CLASS = 30
	TokenAppContainerSid                 TOKEN_INFORMATION_CLASS = 31
	TokenAppContainerNumber              TOKEN_INFORMATION_CLASS = 32
	TokenUserClaimAttributes             TOKEN_INFORMATION_CLASS = 33
	TokenDeviceClaimAttributes           TOKEN_INFORMATION_CLASS = 34
	TokenRestrictedUserClaimAttributes   TOKEN_INFORMATION_CLASS = 35
	TokenRestrictedDeviceClaimAttributes TOKEN_INFORMATION_CLASS = 36
	TokenDeviceGroups                    TOKEN_INFORMATION_CLASS = 37
	TokenRestrictedDeviceGroups          TOKEN_INFORMATION_CLASS = 38
	TokenSecurityAttributes              TOKEN_INFORMATION_CLASS = 39
	TokenIsRestricted                    TOKEN_INFORMATION_CLASS = 40
)

// SECURITY_IMPERSONATION_LEVEL
type SECURITY_IMPERSONATION_LEVEL UINT

const (
	SecurityAnonymous      SECURITY_IMPERSONATION_LEVEL = 0
	SecurityIdentification SECURITY_IMPERSONATION_LEVEL = 1
	SecurityImpersonation  SECURITY_IMPERSONATION_LEVEL = 2
	SecurityDelegation     SECURITY_IMPERSONATION_LEVEL = 3
)

// SECURITY_CONTEXT_TRACKING_MODE
type SECURITY_CONTEXT_TRACKING_MODE BYTE

const (
	SECURITY_DYNAMIC_TRACKING SECURITY_CONTEXT_TRACKING_MODE = 1
	SECURITY_STATIC_TRACKING  SECURITY_CONTEXT_TRACKING_MODE = 0
)

// SECURITY_QUALITY_OF_SERVICE
type SECURITY_QUALITY_OF_SERVICE struct {
	Length              DWORD
	ImpersonationLevel  SECURITY_IMPERSONATION_LEVEL
	ContextTrackingMode SECURITY_CONTEXT_TRACKING_MODE
	EffectiveOnly       BOOLEAN
}
type PSECURITY_QUALITY_OF_SERVICE *SECURITY_QUALITY_OF_SERVICE

// ACL_INFORMATION_CLASS
type ACL_INFORMATION_CLASS UINT

const (
	AclRevisionInformation ACL_INFORMATION_CLASS = 1
	AclSizeInformation     ACL_INFORMATION_CLASS = 2
)

// SAFER_POLICY_INFO_CLASS
type SAFER_POLICY_INFO_CLASS UINT

const (
	SaferPolicyLevelList                    SAFER_POLICY_INFO_CLASS = 1
	SaferPolicyEnableTransparentEnforcement SAFER_POLICY_INFO_CLASS = 2
	SaferPolicyDefaultLevel                 SAFER_POLICY_INFO_CLASS = 3
	SaferPolicyEvaluateUserScope            SAFER_POLICY_INFO_CLASS = 4
	SaferPolicyScopeFlags                   SAFER_POLICY_INFO_CLASS = 5
	SaferPolicyDefaultLevelFlags            SAFER_POLICY_INFO_CLASS = 6
	SaferPolicyAuthenticodeEnabled          SAFER_POLICY_INFO_CLASS = 7
)

// WELL_KNOWN_SID_TYPE
type WELL_KNOWN_SID_TYPE UINT

const (
	WinNullSid                                    WELL_KNOWN_SID_TYPE = 0
	WinWorldSid                                   WELL_KNOWN_SID_TYPE = 1
	WinLocalSid                                   WELL_KNOWN_SID_TYPE = 2
	WinCreatorOwnerSid                            WELL_KNOWN_SID_TYPE = 3
	WinCreatorGroupSid                            WELL_KNOWN_SID_TYPE = 4
	WinCreatorOwnerServerSid                      WELL_KNOWN_SID_TYPE = 5
	WinCreatorGroupServerSid                      WELL_KNOWN_SID_TYPE = 6
	WinNtAuthoritySid                             WELL_KNOWN_SID_TYPE = 7
	WinDialupSid                                  WELL_KNOWN_SID_TYPE = 8
	WinNetworkSid                                 WELL_KNOWN_SID_TYPE = 9
	WinBatchSid                                   WELL_KNOWN_SID_TYPE = 10
	WinInteractiveSid                             WELL_KNOWN_SID_TYPE = 11
	WinServiceSid                                 WELL_KNOWN_SID_TYPE = 12
	WinAnonymousSid                               WELL_KNOWN_SID_TYPE = 13
	WinProxySid                                   WELL_KNOWN_SID_TYPE = 14
	WinEnterpriseControllersSid                   WELL_KNOWN_SID_TYPE = 15
	WinSelfSid                                    WELL_KNOWN_SID_TYPE = 16
	WinAuthenticatedUserSid                       WELL_KNOWN_SID_TYPE = 17
	WinRestrictedCodeSid                          WELL_KNOWN_SID_TYPE = 18
	WinTerminalServerSid                          WELL_KNOWN_SID_TYPE = 19
	WinRemoteLogonIdSid                           WELL_KNOWN_SID_TYPE = 20
	WinLogonIdsSid                                WELL_KNOWN_SID_TYPE = 21
	WinLocalSystemSid                             WELL_KNOWN_SID_TYPE = 22
	WinLocalServiceSid                            WELL_KNOWN_SID_TYPE = 23
	WinNetworkServiceSid                          WELL_KNOWN_SID_TYPE = 24
	WinBuiltinDomainSid                           WELL_KNOWN_SID_TYPE = 25
	WinBuiltinAdministratorsSid                   WELL_KNOWN_SID_TYPE = 26
	WinBuiltinUsersSid                            WELL_KNOWN_SID_TYPE = 27
	WinBuiltinGuestsSid                           WELL_KNOWN_SID_TYPE = 28
	WinBuiltinPowerUsersSid                       WELL_KNOWN_SID_TYPE = 29
	WinBuiltinAccountOperatorsSid                 WELL_KNOWN_SID_TYPE = 30
	WinBuiltinSystemOperatorsSid                  WELL_KNOWN_SID_TYPE = 31
	WinBuiltinPrintOperatorsSid                   WELL_KNOWN_SID_TYPE = 32
	WinBuiltinBackupOperatorsSid                  WELL_KNOWN_SID_TYPE = 33
	WinBuiltinReplicatorSid                       WELL_KNOWN_SID_TYPE = 34
	WinBuiltinPreWindows2000CompatibleAccessSid   WELL_KNOWN_SID_TYPE = 35
	WinBuiltinRemoteDesktopUsersSid               WELL_KNOWN_SID_TYPE = 36
	WinBuiltinNetworkConfigurationOperatorsSid    WELL_KNOWN_SID_TYPE = 37
	WinAccountAdministratorSid                    WELL_KNOWN_SID_TYPE = 38
	WinAccountGuestSid                            WELL_KNOWN_SID_TYPE = 39
	WinAccountKrbtgtSid                           WELL_KNOWN_SID_TYPE = 40
	WinAccountDomainAdminsSid                     WELL_KNOWN_SID_TYPE = 41
	WinAccountDomainUsersSid                      WELL_KNOWN_SID_TYPE = 42
	WinAccountDomainGuestsSid                     WELL_KNOWN_SID_TYPE = 43
	WinAccountComputersSid                        WELL_KNOWN_SID_TYPE = 44
	WinAccountControllersSid                      WELL_KNOWN_SID_TYPE = 45
	WinAccountCertAdminsSid                       WELL_KNOWN_SID_TYPE = 46
	WinAccountSchemaAdminsSid                     WELL_KNOWN_SID_TYPE = 47
	WinAccountEnterpriseAdminsSid                 WELL_KNOWN_SID_TYPE = 48
	WinAccountPolicyAdminsSid                     WELL_KNOWN_SID_TYPE = 49
	WinAccountRasAndIasServersSid                 WELL_KNOWN_SID_TYPE = 50
	WinNTLMAuthenticationSid                      WELL_KNOWN_SID_TYPE = 51
	WinDigestAuthenticationSid                    WELL_KNOWN_SID_TYPE = 52
	WinSChannelAuthenticationSid                  WELL_KNOWN_SID_TYPE = 53
	WinThisOrganizationSid                        WELL_KNOWN_SID_TYPE = 54
	WinOtherOrganizationSid                       WELL_KNOWN_SID_TYPE = 55
	WinBuiltinIncomingForestTrustBuildersSid      WELL_KNOWN_SID_TYPE = 56
	WinBuiltinPerfMonitoringUsersSid              WELL_KNOWN_SID_TYPE = 57
	WinBuiltinPerfLoggingUsersSid                 WELL_KNOWN_SID_TYPE = 58
	WinBuiltinAuthorizationAccessSid              WELL_KNOWN_SID_TYPE = 59
	WinBuiltinTerminalServerLicenseServersSid     WELL_KNOWN_SID_TYPE = 60
	WinBuiltinDCOMUsersSid                        WELL_KNOWN_SID_TYPE = 61
	WinBuiltinIUsersSid                           WELL_KNOWN_SID_TYPE = 62
	WinIUserSid                                   WELL_KNOWN_SID_TYPE = 63
	WinBuiltinCryptoOperatorsSid                  WELL_KNOWN_SID_TYPE = 64
	WinUntrustedLabelSid                          WELL_KNOWN_SID_TYPE = 65
	WinLowLabelSid                                WELL_KNOWN_SID_TYPE = 66
	WinMediumLabelSid                             WELL_KNOWN_SID_TYPE = 67
	WinHighLabelSid                               WELL_KNOWN_SID_TYPE = 68
	WinSystemLabelSid                             WELL_KNOWN_SID_TYPE = 69
	WinWriteRestrictedCodeSid                     WELL_KNOWN_SID_TYPE = 70
	WinCreatorOwnerRightsSid                      WELL_KNOWN_SID_TYPE = 71
	WinCacheablePrincipalsGroupSid                WELL_KNOWN_SID_TYPE = 72
	WinNonCacheablePrincipalsGroupSid             WELL_KNOWN_SID_TYPE = 73
	WinEnterpriseReadonlyControllersSid           WELL_KNOWN_SID_TYPE = 74
	WinAccountReadonlyControllersSid              WELL_KNOWN_SID_TYPE = 75
	WinBuiltinEventLogReadersGroup                WELL_KNOWN_SID_TYPE = 76
	WinNewEnterpriseReadonlyControllersSid        WELL_KNOWN_SID_TYPE = 77
	WinBuiltinCertSvcDComAccessGroup              WELL_KNOWN_SID_TYPE = 78
	WinMediumPlusLabelSid                         WELL_KNOWN_SID_TYPE = 79
	WinLocalLogonSid                              WELL_KNOWN_SID_TYPE = 80
	WinConsoleLogonSid                            WELL_KNOWN_SID_TYPE = 81
	WinThisOrganizationCertificateSid             WELL_KNOWN_SID_TYPE = 82
	WinApplicationPackageAuthoritySid             WELL_KNOWN_SID_TYPE = 83
	WinBuiltinAnyPackageSid                       WELL_KNOWN_SID_TYPE = 84
	WinCapabilityInternetClientSid                WELL_KNOWN_SID_TYPE = 85
	WinCapabilityInternetClientServerSid          WELL_KNOWN_SID_TYPE = 86
	WinCapabilityPrivateNetworkClientServerSid    WELL_KNOWN_SID_TYPE = 87
	WinCapabilityPicturesLibrarySid               WELL_KNOWN_SID_TYPE = 88
	WinCapabilityVideosLibrarySid                 WELL_KNOWN_SID_TYPE = 89
	WinCapabilityMusicLibrarySid                  WELL_KNOWN_SID_TYPE = 90
	WinCapabilityDocumentsLibrarySid              WELL_KNOWN_SID_TYPE = 91
	WinCapabilitySharedUserCertificatesSid        WELL_KNOWN_SID_TYPE = 92
	WinCapabilityDefaultWindowsCredentialsSid     WELL_KNOWN_SID_TYPE = 93
	WinCapabilityRemovableStorageSid              WELL_KNOWN_SID_TYPE = 94
	WinBuiltinRDSRemoteAccessServersSid           WELL_KNOWN_SID_TYPE = 95
	WinBuiltinRDSEndpointServersSid               WELL_KNOWN_SID_TYPE = 96
	WinBuiltinRDSManagementServersSid             WELL_KNOWN_SID_TYPE = 97
	WinUserModeDriversSid                         WELL_KNOWN_SID_TYPE = 98
	WinBuiltinHyperVAdminsSid                     WELL_KNOWN_SID_TYPE = 99
	WinAccountCloneableControllersSid             WELL_KNOWN_SID_TYPE = 100
	WinBuiltinAccessControlAssistanceOperatorsSid WELL_KNOWN_SID_TYPE = 101
	WinBuiltinRemoteManagementUsersSid            WELL_KNOWN_SID_TYPE = 102
	WinAuthenticationAuthorityAssertedSid         WELL_KNOWN_SID_TYPE = 103
	WinAuthenticationServiceAssertedSid           WELL_KNOWN_SID_TYPE = 104
)

// SID_NAME_USE
type SID_NAME_USE UINT

const (
	SidTypeUser           SID_NAME_USE = 1
	SidTypeGroup          SID_NAME_USE = 2
	SidTypeDomain         SID_NAME_USE = 3
	SidTypeAlias          SID_NAME_USE = 4
	SidTypeWellKnownGroup SID_NAME_USE = 5
	SidTypeDeletedAccount SID_NAME_USE = 6
	SidTypeInvalid        SID_NAME_USE = 7
	SidTypeUnknown        SID_NAME_USE = 8
	SidTypeComputer       SID_NAME_USE = 9
	SidTypeLabel          SID_NAME_USE = 10
)

type PSID_NAME_USE *SID_NAME_USE

// SAFER_OBJECT_INFO_CLASS
type SAFER_OBJECT_INFO_CLASS UINT

const (
	SaferObjectLevelId                 SAFER_OBJECT_INFO_CLASS = 1
	SaferObjectScopeId                 SAFER_OBJECT_INFO_CLASS = 2
	SaferObjectFriendlyName            SAFER_OBJECT_INFO_CLASS = 3
	SaferObjectDescription             SAFER_OBJECT_INFO_CLASS = 4
	SaferObjectBuiltin                 SAFER_OBJECT_INFO_CLASS = 5
	SaferObjectDisallowed              SAFER_OBJECT_INFO_CLASS = 6
	SaferObjectDisableMaxPrivilege     SAFER_OBJECT_INFO_CLASS = 7
	SaferObjectInvertDeletedPrivileges SAFER_OBJECT_INFO_CLASS = 8
	SaferObjectDeletedPrivileges       SAFER_OBJECT_INFO_CLASS = 9
	SaferObjectDefaultOwner            SAFER_OBJECT_INFO_CLASS = 10
	SaferObjectSidsToDisable           SAFER_OBJECT_INFO_CLASS = 11
	SaferObjectRestrictedSidsInverted  SAFER_OBJECT_INFO_CLASS = 12
	SaferObjectRestrictedSidsAdded     SAFER_OBJECT_INFO_CLASS = 13
	SaferObjectAllIdentificationGuids  SAFER_OBJECT_INFO_CLASS = 14
	SaferObjectSingleIdentification    SAFER_OBJECT_INFO_CLASS = 15
	SaferObjectExtendedError           SAFER_OBJECT_INFO_CLASS = 16
)

// CRED_PROTECTION_TYPE
type CRED_PROTECTION_TYPE UINT

const (
	CredUnprotected       CRED_PROTECTION_TYPE = 0
	CredUserProtection    CRED_PROTECTION_TYPE = 1
	CredTrustedProtection CRED_PROTECTION_TYPE = 2
)

// CRED_MARSHAL_TYPE
type CRED_MARSHAL_TYPE UINT

const (
	CertCredential               CRED_MARSHAL_TYPE = 1
	UsernameTargetCredential     CRED_MARSHAL_TYPE = 2
	BinaryBlobCredential         CRED_MARSHAL_TYPE = 3
	UsernameForPackedCredentials CRED_MARSHAL_TYPE = 4
)

type PCRED_MARSHAL_TYPE *CRED_MARSHAL_TYPE

// AUDIT_EVENT_TYPE
type AUDIT_EVENT_TYPE UINT

const (
	AuditEventObjectAccess           AUDIT_EVENT_TYPE = 0
	AuditEventDirectoryServiceAccess AUDIT_EVENT_TYPE = 1
)

// POLICY_AUDIT_EVENT_TYPE
type POLICY_AUDIT_EVENT_TYPE UINT

const (
	AuditCategorySystem                 POLICY_AUDIT_EVENT_TYPE = 0
	AuditCategoryLogon                  POLICY_AUDIT_EVENT_TYPE = 1
	AuditCategoryObjectAccess           POLICY_AUDIT_EVENT_TYPE = 2
	AuditCategoryPrivilegeUse           POLICY_AUDIT_EVENT_TYPE = 3
	AuditCategoryDetailedTracking       POLICY_AUDIT_EVENT_TYPE = 4
	AuditCategoryPolicyChange           POLICY_AUDIT_EVENT_TYPE = 5
	AuditCategoryAccountManagement      POLICY_AUDIT_EVENT_TYPE = 6
	AuditCategoryDirectoryServiceAccess POLICY_AUDIT_EVENT_TYPE = 7
	AuditCategoryAccountLogon           POLICY_AUDIT_EVENT_TYPE = 8
)

type PPOLICY_AUDIT_EVENT_TYPE *POLICY_AUDIT_EVENT_TYPE

// OBJECT_TYPE_LIST
type OBJECT_TYPE_LIST struct {
	Level      WORD
	Sbz        WORD
	ObjectType *GUID
}
type POBJECT_TYPE_LIST *OBJECT_TYPE_LIST

// SID_AND_ATTRIBUTES
type SID_AND_ATTRIBUTES struct {
	Sid        PSID
	Attributes DWORD
}
type PSID_AND_ATTRIBUTES *SID_AND_ATTRIBUTES

// TOKEN_GROUPS
type TOKEN_GROUPS struct {
	GroupCount DWORD
	Groups     [1]SID_AND_ATTRIBUTES
}
type PTOKEN_GROUPS *TOKEN_GROUPS

// SECURITY_DESCRIPTOR
type SECURITY_DESCRIPTOR struct {
	Revision BYTE
	Sbz1     BYTE
	Control  SECURITY_DESCRIPTOR_CONTROL
	Owner    PSID
	Group    PSID
	Sacl     PACL
	Dacl     PACL
}
type PSECURITY_DESCRIPTOR *SECURITY_DESCRIPTOR

// SECURITY_ATTRIBUTES
type SECURITY_ATTRIBUTES struct {
	Length             DWORD
	SecurityDescriptor PSECURITY_DESCRIPTOR
	InheritHandle      BOOL
}
type LPSECURITY_ATTRIBUTES *SECURITY_ATTRIBUTES
type PSECURITY_ATTRIBUTES *SECURITY_ATTRIBUTES

// SECURITY_DESCRIPTOR_RELATIVE
type SECURITY_DESCRIPTOR_RELATIVE struct {
	Revision BYTE
	Sbz1     BYTE
	Control  SECURITY_DESCRIPTOR_CONTROL
	Owner    DWORD
	Group    DWORD
	Sacl     DWORD
	Dacl     DWORD
}

type PSECURITY_DESCRIPTOR_RELATIVE *SECURITY_DESCRIPTOR_RELATIVE
type PISECURITY_DESCRIPTOR_RELATIVE *SECURITY_DESCRIPTOR_RELATIVE
