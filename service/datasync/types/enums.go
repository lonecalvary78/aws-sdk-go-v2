// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AgentStatus string

// Enum values for AgentStatus
const (
	AgentStatusOnline  AgentStatus = "ONLINE"
	AgentStatusOffline AgentStatus = "OFFLINE"
)

// Values returns all known values for AgentStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AgentStatus) Values() []AgentStatus {
	return []AgentStatus{
		"ONLINE",
		"OFFLINE",
	}
}

type Atime string

// Enum values for Atime
const (
	AtimeNone       Atime = "NONE"
	AtimeBestEffort Atime = "BEST_EFFORT"
)

// Values returns all known values for Atime. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Atime) Values() []Atime {
	return []Atime{
		"NONE",
		"BEST_EFFORT",
	}
}

type AzureAccessTier string

// Enum values for AzureAccessTier
const (
	AzureAccessTierHot     AzureAccessTier = "HOT"
	AzureAccessTierCool    AzureAccessTier = "COOL"
	AzureAccessTierArchive AzureAccessTier = "ARCHIVE"
)

// Values returns all known values for AzureAccessTier. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AzureAccessTier) Values() []AzureAccessTier {
	return []AzureAccessTier{
		"HOT",
		"COOL",
		"ARCHIVE",
	}
}

type AzureBlobAuthenticationType string

// Enum values for AzureBlobAuthenticationType
const (
	AzureBlobAuthenticationTypeSas  AzureBlobAuthenticationType = "SAS"
	AzureBlobAuthenticationTypeNone AzureBlobAuthenticationType = "NONE"
)

// Values returns all known values for AzureBlobAuthenticationType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AzureBlobAuthenticationType) Values() []AzureBlobAuthenticationType {
	return []AzureBlobAuthenticationType{
		"SAS",
		"NONE",
	}
}

type AzureBlobType string

// Enum values for AzureBlobType
const (
	AzureBlobTypeBlock AzureBlobType = "BLOCK"
)

// Values returns all known values for AzureBlobType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AzureBlobType) Values() []AzureBlobType {
	return []AzureBlobType{
		"BLOCK",
	}
}

type EfsInTransitEncryption string

// Enum values for EfsInTransitEncryption
const (
	EfsInTransitEncryptionNone  EfsInTransitEncryption = "NONE"
	EfsInTransitEncryptionTls12 EfsInTransitEncryption = "TLS1_2"
)

// Values returns all known values for EfsInTransitEncryption. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EfsInTransitEncryption) Values() []EfsInTransitEncryption {
	return []EfsInTransitEncryption{
		"NONE",
		"TLS1_2",
	}
}

type EndpointType string

// Enum values for EndpointType
const (
	EndpointTypePublic      EndpointType = "PUBLIC"
	EndpointTypePrivateLink EndpointType = "PRIVATE_LINK"
	EndpointTypeFips        EndpointType = "FIPS"
)

// Values returns all known values for EndpointType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EndpointType) Values() []EndpointType {
	return []EndpointType{
		"PUBLIC",
		"PRIVATE_LINK",
		"FIPS",
	}
}

type FilterType string

// Enum values for FilterType
const (
	FilterTypeSimplePattern FilterType = "SIMPLE_PATTERN"
)

// Values returns all known values for FilterType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FilterType) Values() []FilterType {
	return []FilterType{
		"SIMPLE_PATTERN",
	}
}

type Gid string

// Enum values for Gid
const (
	GidNone     Gid = "NONE"
	GidIntValue Gid = "INT_VALUE"
	GidName     Gid = "NAME"
	GidBoth     Gid = "BOTH"
)

// Values returns all known values for Gid. Note that this can be expanded in the
// future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Gid) Values() []Gid {
	return []Gid{
		"NONE",
		"INT_VALUE",
		"NAME",
		"BOTH",
	}
}

type HdfsAuthenticationType string

// Enum values for HdfsAuthenticationType
const (
	HdfsAuthenticationTypeSimple   HdfsAuthenticationType = "SIMPLE"
	HdfsAuthenticationTypeKerberos HdfsAuthenticationType = "KERBEROS"
)

// Values returns all known values for HdfsAuthenticationType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HdfsAuthenticationType) Values() []HdfsAuthenticationType {
	return []HdfsAuthenticationType{
		"SIMPLE",
		"KERBEROS",
	}
}

type HdfsDataTransferProtection string

// Enum values for HdfsDataTransferProtection
const (
	HdfsDataTransferProtectionDisabled       HdfsDataTransferProtection = "DISABLED"
	HdfsDataTransferProtectionAuthentication HdfsDataTransferProtection = "AUTHENTICATION"
	HdfsDataTransferProtectionIntegrity      HdfsDataTransferProtection = "INTEGRITY"
	HdfsDataTransferProtectionPrivacy        HdfsDataTransferProtection = "PRIVACY"
)

// Values returns all known values for HdfsDataTransferProtection. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HdfsDataTransferProtection) Values() []HdfsDataTransferProtection {
	return []HdfsDataTransferProtection{
		"DISABLED",
		"AUTHENTICATION",
		"INTEGRITY",
		"PRIVACY",
	}
}

type HdfsRpcProtection string

// Enum values for HdfsRpcProtection
const (
	HdfsRpcProtectionDisabled       HdfsRpcProtection = "DISABLED"
	HdfsRpcProtectionAuthentication HdfsRpcProtection = "AUTHENTICATION"
	HdfsRpcProtectionIntegrity      HdfsRpcProtection = "INTEGRITY"
	HdfsRpcProtectionPrivacy        HdfsRpcProtection = "PRIVACY"
)

// Values returns all known values for HdfsRpcProtection. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HdfsRpcProtection) Values() []HdfsRpcProtection {
	return []HdfsRpcProtection{
		"DISABLED",
		"AUTHENTICATION",
		"INTEGRITY",
		"PRIVACY",
	}
}

type LocationFilterName string

// Enum values for LocationFilterName
const (
	LocationFilterNameLocationUri  LocationFilterName = "LocationUri"
	LocationFilterNameLocationType LocationFilterName = "LocationType"
	LocationFilterNameCreationTime LocationFilterName = "CreationTime"
)

// Values returns all known values for LocationFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LocationFilterName) Values() []LocationFilterName {
	return []LocationFilterName{
		"LocationUri",
		"LocationType",
		"CreationTime",
	}
}

type LogLevel string

// Enum values for LogLevel
const (
	LogLevelOff      LogLevel = "OFF"
	LogLevelBasic    LogLevel = "BASIC"
	LogLevelTransfer LogLevel = "TRANSFER"
)

// Values returns all known values for LogLevel. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LogLevel) Values() []LogLevel {
	return []LogLevel{
		"OFF",
		"BASIC",
		"TRANSFER",
	}
}

type ManifestAction string

// Enum values for ManifestAction
const (
	ManifestActionTransfer ManifestAction = "TRANSFER"
)

// Values returns all known values for ManifestAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ManifestAction) Values() []ManifestAction {
	return []ManifestAction{
		"TRANSFER",
	}
}

type ManifestFormat string

// Enum values for ManifestFormat
const (
	ManifestFormatCsv ManifestFormat = "CSV"
)

// Values returns all known values for ManifestFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ManifestFormat) Values() []ManifestFormat {
	return []ManifestFormat{
		"CSV",
	}
}

type Mtime string

// Enum values for Mtime
const (
	MtimeNone     Mtime = "NONE"
	MtimePreserve Mtime = "PRESERVE"
)

// Values returns all known values for Mtime. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Mtime) Values() []Mtime {
	return []Mtime{
		"NONE",
		"PRESERVE",
	}
}

type NfsVersion string

// Enum values for NfsVersion
const (
	NfsVersionAutomatic NfsVersion = "AUTOMATIC"
	NfsVersionNfs3      NfsVersion = "NFS3"
	NfsVersionNfs40     NfsVersion = "NFS4_0"
	NfsVersionNfs41     NfsVersion = "NFS4_1"
)

// Values returns all known values for NfsVersion. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NfsVersion) Values() []NfsVersion {
	return []NfsVersion{
		"AUTOMATIC",
		"NFS3",
		"NFS4_0",
		"NFS4_1",
	}
}

type ObjectStorageServerProtocol string

// Enum values for ObjectStorageServerProtocol
const (
	ObjectStorageServerProtocolHttps ObjectStorageServerProtocol = "HTTPS"
	ObjectStorageServerProtocolHttp  ObjectStorageServerProtocol = "HTTP"
)

// Values returns all known values for ObjectStorageServerProtocol. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ObjectStorageServerProtocol) Values() []ObjectStorageServerProtocol {
	return []ObjectStorageServerProtocol{
		"HTTPS",
		"HTTP",
	}
}

type ObjectTags string

// Enum values for ObjectTags
const (
	ObjectTagsPreserve ObjectTags = "PRESERVE"
	ObjectTagsNone     ObjectTags = "NONE"
)

// Values returns all known values for ObjectTags. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ObjectTags) Values() []ObjectTags {
	return []ObjectTags{
		"PRESERVE",
		"NONE",
	}
}

type ObjectVersionIds string

// Enum values for ObjectVersionIds
const (
	ObjectVersionIdsInclude ObjectVersionIds = "INCLUDE"
	ObjectVersionIdsNone    ObjectVersionIds = "NONE"
)

// Values returns all known values for ObjectVersionIds. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ObjectVersionIds) Values() []ObjectVersionIds {
	return []ObjectVersionIds{
		"INCLUDE",
		"NONE",
	}
}

type Operator string

// Enum values for Operator
const (
	OperatorEq          Operator = "Equals"
	OperatorNe          Operator = "NotEquals"
	OperatorIn          Operator = "In"
	OperatorLe          Operator = "LessThanOrEqual"
	OperatorLt          Operator = "LessThan"
	OperatorGe          Operator = "GreaterThanOrEqual"
	OperatorGt          Operator = "GreaterThan"
	OperatorContains    Operator = "Contains"
	OperatorNotContains Operator = "NotContains"
	OperatorBeginsWith  Operator = "BeginsWith"
)

// Values returns all known values for Operator. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Operator) Values() []Operator {
	return []Operator{
		"Equals",
		"NotEquals",
		"In",
		"LessThanOrEqual",
		"LessThan",
		"GreaterThanOrEqual",
		"GreaterThan",
		"Contains",
		"NotContains",
		"BeginsWith",
	}
}

type OverwriteMode string

// Enum values for OverwriteMode
const (
	OverwriteModeAlways OverwriteMode = "ALWAYS"
	OverwriteModeNever  OverwriteMode = "NEVER"
)

// Values returns all known values for OverwriteMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OverwriteMode) Values() []OverwriteMode {
	return []OverwriteMode{
		"ALWAYS",
		"NEVER",
	}
}

type PhaseStatus string

// Enum values for PhaseStatus
const (
	PhaseStatusPending PhaseStatus = "PENDING"
	PhaseStatusSuccess PhaseStatus = "SUCCESS"
	PhaseStatusError   PhaseStatus = "ERROR"
)

// Values returns all known values for PhaseStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PhaseStatus) Values() []PhaseStatus {
	return []PhaseStatus{
		"PENDING",
		"SUCCESS",
		"ERROR",
	}
}

type PosixPermissions string

// Enum values for PosixPermissions
const (
	PosixPermissionsNone     PosixPermissions = "NONE"
	PosixPermissionsPreserve PosixPermissions = "PRESERVE"
)

// Values returns all known values for PosixPermissions. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PosixPermissions) Values() []PosixPermissions {
	return []PosixPermissions{
		"NONE",
		"PRESERVE",
	}
}

type PreserveDeletedFiles string

// Enum values for PreserveDeletedFiles
const (
	PreserveDeletedFilesPreserve PreserveDeletedFiles = "PRESERVE"
	PreserveDeletedFilesRemove   PreserveDeletedFiles = "REMOVE"
)

// Values returns all known values for PreserveDeletedFiles. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PreserveDeletedFiles) Values() []PreserveDeletedFiles {
	return []PreserveDeletedFiles{
		"PRESERVE",
		"REMOVE",
	}
}

type PreserveDevices string

// Enum values for PreserveDevices
const (
	PreserveDevicesNone     PreserveDevices = "NONE"
	PreserveDevicesPreserve PreserveDevices = "PRESERVE"
)

// Values returns all known values for PreserveDevices. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PreserveDevices) Values() []PreserveDevices {
	return []PreserveDevices{
		"NONE",
		"PRESERVE",
	}
}

type ReportLevel string

// Enum values for ReportLevel
const (
	ReportLevelErrorsOnly         ReportLevel = "ERRORS_ONLY"
	ReportLevelSuccessesAndErrors ReportLevel = "SUCCESSES_AND_ERRORS"
)

// Values returns all known values for ReportLevel. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReportLevel) Values() []ReportLevel {
	return []ReportLevel{
		"ERRORS_ONLY",
		"SUCCESSES_AND_ERRORS",
	}
}

type ReportOutputType string

// Enum values for ReportOutputType
const (
	ReportOutputTypeSummaryOnly ReportOutputType = "SUMMARY_ONLY"
	ReportOutputTypeStandard    ReportOutputType = "STANDARD"
)

// Values returns all known values for ReportOutputType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReportOutputType) Values() []ReportOutputType {
	return []ReportOutputType{
		"SUMMARY_ONLY",
		"STANDARD",
	}
}

type S3StorageClass string

// Enum values for S3StorageClass
const (
	S3StorageClassStandard                S3StorageClass = "STANDARD"
	S3StorageClassStandardIa              S3StorageClass = "STANDARD_IA"
	S3StorageClassOnezoneIa               S3StorageClass = "ONEZONE_IA"
	S3StorageClassIntelligentTiering      S3StorageClass = "INTELLIGENT_TIERING"
	S3StorageClassGlacier                 S3StorageClass = "GLACIER"
	S3StorageClassDeepArchive             S3StorageClass = "DEEP_ARCHIVE"
	S3StorageClassOutposts                S3StorageClass = "OUTPOSTS"
	S3StorageClassGlacierInstantRetrieval S3StorageClass = "GLACIER_INSTANT_RETRIEVAL"
)

// Values returns all known values for S3StorageClass. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3StorageClass) Values() []S3StorageClass {
	return []S3StorageClass{
		"STANDARD",
		"STANDARD_IA",
		"ONEZONE_IA",
		"INTELLIGENT_TIERING",
		"GLACIER",
		"DEEP_ARCHIVE",
		"OUTPOSTS",
		"GLACIER_INSTANT_RETRIEVAL",
	}
}

type ScheduleDisabledBy string

// Enum values for ScheduleDisabledBy
const (
	ScheduleDisabledByUser    ScheduleDisabledBy = "USER"
	ScheduleDisabledByService ScheduleDisabledBy = "SERVICE"
)

// Values returns all known values for ScheduleDisabledBy. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScheduleDisabledBy) Values() []ScheduleDisabledBy {
	return []ScheduleDisabledBy{
		"USER",
		"SERVICE",
	}
}

type ScheduleStatus string

// Enum values for ScheduleStatus
const (
	ScheduleStatusEnabled  ScheduleStatus = "ENABLED"
	ScheduleStatusDisabled ScheduleStatus = "DISABLED"
)

// Values returns all known values for ScheduleStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScheduleStatus) Values() []ScheduleStatus {
	return []ScheduleStatus{
		"ENABLED",
		"DISABLED",
	}
}

type SmbAuthenticationType string

// Enum values for SmbAuthenticationType
const (
	SmbAuthenticationTypeNtlm     SmbAuthenticationType = "NTLM"
	SmbAuthenticationTypeKerberos SmbAuthenticationType = "KERBEROS"
)

// Values returns all known values for SmbAuthenticationType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SmbAuthenticationType) Values() []SmbAuthenticationType {
	return []SmbAuthenticationType{
		"NTLM",
		"KERBEROS",
	}
}

type SmbSecurityDescriptorCopyFlags string

// Enum values for SmbSecurityDescriptorCopyFlags
const (
	SmbSecurityDescriptorCopyFlagsNone          SmbSecurityDescriptorCopyFlags = "NONE"
	SmbSecurityDescriptorCopyFlagsOwnerDacl     SmbSecurityDescriptorCopyFlags = "OWNER_DACL"
	SmbSecurityDescriptorCopyFlagsOwnerDaclSacl SmbSecurityDescriptorCopyFlags = "OWNER_DACL_SACL"
)

// Values returns all known values for SmbSecurityDescriptorCopyFlags. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SmbSecurityDescriptorCopyFlags) Values() []SmbSecurityDescriptorCopyFlags {
	return []SmbSecurityDescriptorCopyFlags{
		"NONE",
		"OWNER_DACL",
		"OWNER_DACL_SACL",
	}
}

type SmbVersion string

// Enum values for SmbVersion
const (
	SmbVersionAutomatic SmbVersion = "AUTOMATIC"
	SmbVersionSmb2      SmbVersion = "SMB2"
	SmbVersionSmb3      SmbVersion = "SMB3"
	SmbVersionSmb1      SmbVersion = "SMB1"
	SmbVersionSmb20     SmbVersion = "SMB2_0"
)

// Values returns all known values for SmbVersion. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SmbVersion) Values() []SmbVersion {
	return []SmbVersion{
		"AUTOMATIC",
		"SMB2",
		"SMB3",
		"SMB1",
		"SMB2_0",
	}
}

type TaskExecutionStatus string

// Enum values for TaskExecutionStatus
const (
	TaskExecutionStatusQueued       TaskExecutionStatus = "QUEUED"
	TaskExecutionStatusCancelling   TaskExecutionStatus = "CANCELLING"
	TaskExecutionStatusLaunching    TaskExecutionStatus = "LAUNCHING"
	TaskExecutionStatusPreparing    TaskExecutionStatus = "PREPARING"
	TaskExecutionStatusTransferring TaskExecutionStatus = "TRANSFERRING"
	TaskExecutionStatusVerifying    TaskExecutionStatus = "VERIFYING"
	TaskExecutionStatusSuccess      TaskExecutionStatus = "SUCCESS"
	TaskExecutionStatusError        TaskExecutionStatus = "ERROR"
)

// Values returns all known values for TaskExecutionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TaskExecutionStatus) Values() []TaskExecutionStatus {
	return []TaskExecutionStatus{
		"QUEUED",
		"CANCELLING",
		"LAUNCHING",
		"PREPARING",
		"TRANSFERRING",
		"VERIFYING",
		"SUCCESS",
		"ERROR",
	}
}

type TaskFilterName string

// Enum values for TaskFilterName
const (
	TaskFilterNameLocationId   TaskFilterName = "LocationId"
	TaskFilterNameCreationTime TaskFilterName = "CreationTime"
)

// Values returns all known values for TaskFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TaskFilterName) Values() []TaskFilterName {
	return []TaskFilterName{
		"LocationId",
		"CreationTime",
	}
}

type TaskMode string

// Enum values for TaskMode
const (
	TaskModeBasic    TaskMode = "BASIC"
	TaskModeEnhanced TaskMode = "ENHANCED"
)

// Values returns all known values for TaskMode. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TaskMode) Values() []TaskMode {
	return []TaskMode{
		"BASIC",
		"ENHANCED",
	}
}

type TaskQueueing string

// Enum values for TaskQueueing
const (
	TaskQueueingEnabled  TaskQueueing = "ENABLED"
	TaskQueueingDisabled TaskQueueing = "DISABLED"
)

// Values returns all known values for TaskQueueing. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TaskQueueing) Values() []TaskQueueing {
	return []TaskQueueing{
		"ENABLED",
		"DISABLED",
	}
}

type TaskStatus string

// Enum values for TaskStatus
const (
	TaskStatusAvailable   TaskStatus = "AVAILABLE"
	TaskStatusCreating    TaskStatus = "CREATING"
	TaskStatusQueued      TaskStatus = "QUEUED"
	TaskStatusRunning     TaskStatus = "RUNNING"
	TaskStatusUnavailable TaskStatus = "UNAVAILABLE"
)

// Values returns all known values for TaskStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TaskStatus) Values() []TaskStatus {
	return []TaskStatus{
		"AVAILABLE",
		"CREATING",
		"QUEUED",
		"RUNNING",
		"UNAVAILABLE",
	}
}

type TransferMode string

// Enum values for TransferMode
const (
	TransferModeChanged TransferMode = "CHANGED"
	TransferModeAll     TransferMode = "ALL"
)

// Values returns all known values for TransferMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TransferMode) Values() []TransferMode {
	return []TransferMode{
		"CHANGED",
		"ALL",
	}
}

type Uid string

// Enum values for Uid
const (
	UidNone     Uid = "NONE"
	UidIntValue Uid = "INT_VALUE"
	UidName     Uid = "NAME"
	UidBoth     Uid = "BOTH"
)

// Values returns all known values for Uid. Note that this can be expanded in the
// future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Uid) Values() []Uid {
	return []Uid{
		"NONE",
		"INT_VALUE",
		"NAME",
		"BOTH",
	}
}

type VerifyMode string

// Enum values for VerifyMode
const (
	VerifyModePointInTimeConsistent VerifyMode = "POINT_IN_TIME_CONSISTENT"
	VerifyModeOnlyFilesTransferred  VerifyMode = "ONLY_FILES_TRANSFERRED"
	VerifyModeNone                  VerifyMode = "NONE"
)

// Values returns all known values for VerifyMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (VerifyMode) Values() []VerifyMode {
	return []VerifyMode{
		"POINT_IN_TIME_CONSISTENT",
		"ONLY_FILES_TRANSFERRED",
		"NONE",
	}
}
