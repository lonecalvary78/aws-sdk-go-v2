// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AcceptanceType string

// Enum values for AcceptanceType
const (
	// Do not require explicit click-through acceptance of the Term associated with
	// this Report
	AcceptanceTypePassthrough AcceptanceType = "PASSTHROUGH"
	// Require explicit click-through acceptance of the Term associated with this
	// Report.
	AcceptanceTypeExplicit AcceptanceType = "EXPLICIT"
)

// Values returns all known values for AcceptanceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AcceptanceType) Values() []AcceptanceType {
	return []AcceptanceType{
		"PASSTHROUGH",
		"EXPLICIT",
	}
}

type AgreementType string

// Enum values for AgreementType
const (
	AgreementTypeCustom   AgreementType = "CUSTOM"
	AgreementTypeDefault  AgreementType = "DEFAULT"
	AgreementTypeModified AgreementType = "MODIFIED"
)

// Values returns all known values for AgreementType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AgreementType) Values() []AgreementType {
	return []AgreementType{
		"CUSTOM",
		"DEFAULT",
		"MODIFIED",
	}
}

type CustomerAgreementState string

// Enum values for CustomerAgreementState
const (
	CustomerAgreementStateActive             CustomerAgreementState = "ACTIVE"
	CustomerAgreementStateCustomerTerminated CustomerAgreementState = "CUSTOMER_TERMINATED"
	CustomerAgreementStateAwsTerminated      CustomerAgreementState = "AWS_TERMINATED"
)

// Values returns all known values for CustomerAgreementState. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CustomerAgreementState) Values() []CustomerAgreementState {
	return []CustomerAgreementState{
		"ACTIVE",
		"CUSTOMER_TERMINATED",
		"AWS_TERMINATED",
	}
}

type NotificationSubscriptionStatus string

// Enum values for NotificationSubscriptionStatus
const (
	NotificationSubscriptionStatusSubscribed    NotificationSubscriptionStatus = "SUBSCRIBED"
	NotificationSubscriptionStatusNotSubscribed NotificationSubscriptionStatus = "NOT_SUBSCRIBED"
)

// Values returns all known values for NotificationSubscriptionStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NotificationSubscriptionStatus) Values() []NotificationSubscriptionStatus {
	return []NotificationSubscriptionStatus{
		"SUBSCRIBED",
		"NOT_SUBSCRIBED",
	}
}

type PublishedState string

// Enum values for PublishedState
const (
	PublishedStatePublished   PublishedState = "PUBLISHED"
	PublishedStateUnpublished PublishedState = "UNPUBLISHED"
)

// Values returns all known values for PublishedState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PublishedState) Values() []PublishedState {
	return []PublishedState{
		"PUBLISHED",
		"UNPUBLISHED",
	}
}

type UploadState string

// Enum values for UploadState
const (
	UploadStateProcessing UploadState = "PROCESSING"
	UploadStateComplete   UploadState = "COMPLETE"
	UploadStateFailed     UploadState = "FAILED"
	UploadStateFault      UploadState = "FAULT"
)

// Values returns all known values for UploadState. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (UploadState) Values() []UploadState {
	return []UploadState{
		"PROCESSING",
		"COMPLETE",
		"FAILED",
		"FAULT",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "unknownOperation"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "cannotParse"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "fieldValidationFailed"
	ValidationExceptionReasonInvalidToken          ValidationExceptionReason = "invalidToken"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "other"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"unknownOperation",
		"cannotParse",
		"fieldValidationFailed",
		"invalidToken",
		"other",
	}
}
