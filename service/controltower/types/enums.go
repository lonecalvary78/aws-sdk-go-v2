// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type BaselineOperationStatus string

// Enum values for BaselineOperationStatus
const (
	BaselineOperationStatusSucceeded  BaselineOperationStatus = "SUCCEEDED"
	BaselineOperationStatusFailed     BaselineOperationStatus = "FAILED"
	BaselineOperationStatusInProgress BaselineOperationStatus = "IN_PROGRESS"
)

// Values returns all known values for BaselineOperationStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (BaselineOperationStatus) Values() []BaselineOperationStatus {
	return []BaselineOperationStatus{
		"SUCCEEDED",
		"FAILED",
		"IN_PROGRESS",
	}
}

type BaselineOperationType string

// Enum values for BaselineOperationType
const (
	BaselineOperationTypeEnableBaseline        BaselineOperationType = "ENABLE_BASELINE"
	BaselineOperationTypeDisableBaseline       BaselineOperationType = "DISABLE_BASELINE"
	BaselineOperationTypeUpdateEnabledBaseline BaselineOperationType = "UPDATE_ENABLED_BASELINE"
	BaselineOperationTypeResetEnabledBaseline  BaselineOperationType = "RESET_ENABLED_BASELINE"
)

// Values returns all known values for BaselineOperationType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (BaselineOperationType) Values() []BaselineOperationType {
	return []BaselineOperationType{
		"ENABLE_BASELINE",
		"DISABLE_BASELINE",
		"UPDATE_ENABLED_BASELINE",
		"RESET_ENABLED_BASELINE",
	}
}

type ControlOperationStatus string

// Enum values for ControlOperationStatus
const (
	ControlOperationStatusSucceeded  ControlOperationStatus = "SUCCEEDED"
	ControlOperationStatusFailed     ControlOperationStatus = "FAILED"
	ControlOperationStatusInProgress ControlOperationStatus = "IN_PROGRESS"
)

// Values returns all known values for ControlOperationStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ControlOperationStatus) Values() []ControlOperationStatus {
	return []ControlOperationStatus{
		"SUCCEEDED",
		"FAILED",
		"IN_PROGRESS",
	}
}

type ControlOperationType string

// Enum values for ControlOperationType
const (
	ControlOperationTypeEnableControl        ControlOperationType = "ENABLE_CONTROL"
	ControlOperationTypeDisableControl       ControlOperationType = "DISABLE_CONTROL"
	ControlOperationTypeUpdateEnabledControl ControlOperationType = "UPDATE_ENABLED_CONTROL"
	ControlOperationTypeResetEnabledControl  ControlOperationType = "RESET_ENABLED_CONTROL"
)

// Values returns all known values for ControlOperationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ControlOperationType) Values() []ControlOperationType {
	return []ControlOperationType{
		"ENABLE_CONTROL",
		"DISABLE_CONTROL",
		"UPDATE_ENABLED_CONTROL",
		"RESET_ENABLED_CONTROL",
	}
}

type DriftStatus string

// Enum values for DriftStatus
const (
	DriftStatusDrifted     DriftStatus = "DRIFTED"
	DriftStatusInSync      DriftStatus = "IN_SYNC"
	DriftStatusNotChecking DriftStatus = "NOT_CHECKING"
	DriftStatusUnknown     DriftStatus = "UNKNOWN"
)

// Values returns all known values for DriftStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DriftStatus) Values() []DriftStatus {
	return []DriftStatus{
		"DRIFTED",
		"IN_SYNC",
		"NOT_CHECKING",
		"UNKNOWN",
	}
}

type EnabledBaselineDriftStatus string

// Enum values for EnabledBaselineDriftStatus
const (
	EnabledBaselineDriftStatusInSync  EnabledBaselineDriftStatus = "IN_SYNC"
	EnabledBaselineDriftStatusDrifted EnabledBaselineDriftStatus = "DRIFTED"
)

// Values returns all known values for EnabledBaselineDriftStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EnabledBaselineDriftStatus) Values() []EnabledBaselineDriftStatus {
	return []EnabledBaselineDriftStatus{
		"IN_SYNC",
		"DRIFTED",
	}
}

type EnablementStatus string

// Enum values for EnablementStatus
const (
	EnablementStatusSucceeded   EnablementStatus = "SUCCEEDED"
	EnablementStatusFailed      EnablementStatus = "FAILED"
	EnablementStatusUnderChange EnablementStatus = "UNDER_CHANGE"
)

// Values returns all known values for EnablementStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EnablementStatus) Values() []EnablementStatus {
	return []EnablementStatus{
		"SUCCEEDED",
		"FAILED",
		"UNDER_CHANGE",
	}
}

type LandingZoneDriftStatus string

// Enum values for LandingZoneDriftStatus
const (
	LandingZoneDriftStatusDrifted LandingZoneDriftStatus = "DRIFTED"
	LandingZoneDriftStatusInSync  LandingZoneDriftStatus = "IN_SYNC"
)

// Values returns all known values for LandingZoneDriftStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LandingZoneDriftStatus) Values() []LandingZoneDriftStatus {
	return []LandingZoneDriftStatus{
		"DRIFTED",
		"IN_SYNC",
	}
}

type LandingZoneOperationStatus string

// Enum values for LandingZoneOperationStatus
const (
	LandingZoneOperationStatusSucceeded  LandingZoneOperationStatus = "SUCCEEDED"
	LandingZoneOperationStatusFailed     LandingZoneOperationStatus = "FAILED"
	LandingZoneOperationStatusInProgress LandingZoneOperationStatus = "IN_PROGRESS"
)

// Values returns all known values for LandingZoneOperationStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LandingZoneOperationStatus) Values() []LandingZoneOperationStatus {
	return []LandingZoneOperationStatus{
		"SUCCEEDED",
		"FAILED",
		"IN_PROGRESS",
	}
}

type LandingZoneOperationType string

// Enum values for LandingZoneOperationType
const (
	LandingZoneOperationTypeDelete LandingZoneOperationType = "DELETE"
	LandingZoneOperationTypeCreate LandingZoneOperationType = "CREATE"
	LandingZoneOperationTypeUpdate LandingZoneOperationType = "UPDATE"
	LandingZoneOperationTypeReset  LandingZoneOperationType = "RESET"
)

// Values returns all known values for LandingZoneOperationType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LandingZoneOperationType) Values() []LandingZoneOperationType {
	return []LandingZoneOperationType{
		"DELETE",
		"CREATE",
		"UPDATE",
		"RESET",
	}
}

type LandingZoneStatus string

// Enum values for LandingZoneStatus
const (
	LandingZoneStatusActive     LandingZoneStatus = "ACTIVE"
	LandingZoneStatusProcessing LandingZoneStatus = "PROCESSING"
	LandingZoneStatusFailed     LandingZoneStatus = "FAILED"
)

// Values returns all known values for LandingZoneStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LandingZoneStatus) Values() []LandingZoneStatus {
	return []LandingZoneStatus{
		"ACTIVE",
		"PROCESSING",
		"FAILED",
	}
}
