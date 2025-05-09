// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccountTargeting string

// Enum values for AccountTargeting
const (
	AccountTargetingSingleAccount AccountTargeting = "single-account"
	AccountTargetingMultiAccount  AccountTargeting = "multi-account"
)

// Values returns all known values for AccountTargeting. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AccountTargeting) Values() []AccountTargeting {
	return []AccountTargeting{
		"single-account",
		"multi-account",
	}
}

type ActionsMode string

// Enum values for ActionsMode
const (
	ActionsModeSkipAll ActionsMode = "skip-all"
	ActionsModeRunAll  ActionsMode = "run-all"
)

// Values returns all known values for ActionsMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ActionsMode) Values() []ActionsMode {
	return []ActionsMode{
		"skip-all",
		"run-all",
	}
}

type EmptyTargetResolutionMode string

// Enum values for EmptyTargetResolutionMode
const (
	EmptyTargetResolutionModeFail EmptyTargetResolutionMode = "fail"
	EmptyTargetResolutionModeSkip EmptyTargetResolutionMode = "skip"
)

// Values returns all known values for EmptyTargetResolutionMode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EmptyTargetResolutionMode) Values() []EmptyTargetResolutionMode {
	return []EmptyTargetResolutionMode{
		"fail",
		"skip",
	}
}

type ExperimentActionStatus string

// Enum values for ExperimentActionStatus
const (
	ExperimentActionStatusPending    ExperimentActionStatus = "pending"
	ExperimentActionStatusInitiating ExperimentActionStatus = "initiating"
	ExperimentActionStatusRunning    ExperimentActionStatus = "running"
	ExperimentActionStatusCompleted  ExperimentActionStatus = "completed"
	ExperimentActionStatusCancelled  ExperimentActionStatus = "cancelled"
	ExperimentActionStatusStopping   ExperimentActionStatus = "stopping"
	ExperimentActionStatusStopped    ExperimentActionStatus = "stopped"
	ExperimentActionStatusFailed     ExperimentActionStatus = "failed"
	ExperimentActionStatusSkipped    ExperimentActionStatus = "skipped"
)

// Values returns all known values for ExperimentActionStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExperimentActionStatus) Values() []ExperimentActionStatus {
	return []ExperimentActionStatus{
		"pending",
		"initiating",
		"running",
		"completed",
		"cancelled",
		"stopping",
		"stopped",
		"failed",
		"skipped",
	}
}

type ExperimentReportStatus string

// Enum values for ExperimentReportStatus
const (
	ExperimentReportStatusPending   ExperimentReportStatus = "pending"
	ExperimentReportStatusRunning   ExperimentReportStatus = "running"
	ExperimentReportStatusCompleted ExperimentReportStatus = "completed"
	ExperimentReportStatusCancelled ExperimentReportStatus = "cancelled"
	ExperimentReportStatusFailed    ExperimentReportStatus = "failed"
)

// Values returns all known values for ExperimentReportStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExperimentReportStatus) Values() []ExperimentReportStatus {
	return []ExperimentReportStatus{
		"pending",
		"running",
		"completed",
		"cancelled",
		"failed",
	}
}

type ExperimentStatus string

// Enum values for ExperimentStatus
const (
	ExperimentStatusPending    ExperimentStatus = "pending"
	ExperimentStatusInitiating ExperimentStatus = "initiating"
	ExperimentStatusRunning    ExperimentStatus = "running"
	ExperimentStatusCompleted  ExperimentStatus = "completed"
	ExperimentStatusStopping   ExperimentStatus = "stopping"
	ExperimentStatusStopped    ExperimentStatus = "stopped"
	ExperimentStatusFailed     ExperimentStatus = "failed"
	ExperimentStatusCancelled  ExperimentStatus = "cancelled"
)

// Values returns all known values for ExperimentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExperimentStatus) Values() []ExperimentStatus {
	return []ExperimentStatus{
		"pending",
		"initiating",
		"running",
		"completed",
		"stopping",
		"stopped",
		"failed",
		"cancelled",
	}
}

type SafetyLeverStatus string

// Enum values for SafetyLeverStatus
const (
	SafetyLeverStatusDisengaged SafetyLeverStatus = "disengaged"
	SafetyLeverStatusEngaged    SafetyLeverStatus = "engaged"
	SafetyLeverStatusEngaging   SafetyLeverStatus = "engaging"
)

// Values returns all known values for SafetyLeverStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SafetyLeverStatus) Values() []SafetyLeverStatus {
	return []SafetyLeverStatus{
		"disengaged",
		"engaged",
		"engaging",
	}
}

type SafetyLeverStatusInput string

// Enum values for SafetyLeverStatusInput
const (
	SafetyLeverStatusInputDisengaged SafetyLeverStatusInput = "disengaged"
	SafetyLeverStatusInputEngaged    SafetyLeverStatusInput = "engaged"
)

// Values returns all known values for SafetyLeverStatusInput. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SafetyLeverStatusInput) Values() []SafetyLeverStatusInput {
	return []SafetyLeverStatusInput{
		"disengaged",
		"engaged",
	}
}
