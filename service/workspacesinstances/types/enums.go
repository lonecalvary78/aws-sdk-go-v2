// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AmdSevSnpEnum string

// Enum values for AmdSevSnpEnum
const (
	AmdSevSnpEnumEnabled  AmdSevSnpEnum = "enabled"
	AmdSevSnpEnumDisabled AmdSevSnpEnum = "disabled"
)

// Values returns all known values for AmdSevSnpEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AmdSevSnpEnum) Values() []AmdSevSnpEnum {
	return []AmdSevSnpEnum{
		"enabled",
		"disabled",
	}
}

type AutoRecoveryEnum string

// Enum values for AutoRecoveryEnum
const (
	AutoRecoveryEnumDisabled AutoRecoveryEnum = "disabled"
	AutoRecoveryEnumDefault  AutoRecoveryEnum = "default"
)

// Values returns all known values for AutoRecoveryEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AutoRecoveryEnum) Values() []AutoRecoveryEnum {
	return []AutoRecoveryEnum{
		"disabled",
		"default",
	}
}

type BandwidthWeightingEnum string

// Enum values for BandwidthWeightingEnum
const (
	BandwidthWeightingEnumDefault BandwidthWeightingEnum = "default"
	BandwidthWeightingEnumVpc1    BandwidthWeightingEnum = "vpc-1"
	BandwidthWeightingEnumEbs1    BandwidthWeightingEnum = "ebs-1"
)

// Values returns all known values for BandwidthWeightingEnum. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (BandwidthWeightingEnum) Values() []BandwidthWeightingEnum {
	return []BandwidthWeightingEnum{
		"default",
		"vpc-1",
		"ebs-1",
	}
}

type CapacityReservationPreferenceEnum string

// Enum values for CapacityReservationPreferenceEnum
const (
	CapacityReservationPreferenceEnumCapacityReservationsOnly CapacityReservationPreferenceEnum = "capacity-reservations-only"
	CapacityReservationPreferenceEnumOpen                     CapacityReservationPreferenceEnum = "open"
	CapacityReservationPreferenceEnumNone                     CapacityReservationPreferenceEnum = "none"
)

// Values returns all known values for CapacityReservationPreferenceEnum. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CapacityReservationPreferenceEnum) Values() []CapacityReservationPreferenceEnum {
	return []CapacityReservationPreferenceEnum{
		"capacity-reservations-only",
		"open",
		"none",
	}
}

type CpuCreditsEnum string

// Enum values for CpuCreditsEnum
const (
	CpuCreditsEnumStandard  CpuCreditsEnum = "standard"
	CpuCreditsEnumUnlimited CpuCreditsEnum = "unlimited"
)

// Values returns all known values for CpuCreditsEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CpuCreditsEnum) Values() []CpuCreditsEnum {
	return []CpuCreditsEnum{
		"standard",
		"unlimited",
	}
}

type DisassociateModeEnum string

// Enum values for DisassociateModeEnum
const (
	DisassociateModeEnumForce   DisassociateModeEnum = "FORCE"
	DisassociateModeEnumNoForce DisassociateModeEnum = "NO_FORCE"
)

// Values returns all known values for DisassociateModeEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DisassociateModeEnum) Values() []DisassociateModeEnum {
	return []DisassociateModeEnum{
		"FORCE",
		"NO_FORCE",
	}
}

type HostnameTypeEnum string

// Enum values for HostnameTypeEnum
const (
	HostnameTypeEnumIpName       HostnameTypeEnum = "ip-name"
	HostnameTypeEnumResourceName HostnameTypeEnum = "resource-name"
)

// Values returns all known values for HostnameTypeEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HostnameTypeEnum) Values() []HostnameTypeEnum {
	return []HostnameTypeEnum{
		"ip-name",
		"resource-name",
	}
}

type HttpEndpointEnum string

// Enum values for HttpEndpointEnum
const (
	HttpEndpointEnumEnabled  HttpEndpointEnum = "enabled"
	HttpEndpointEnumDisabled HttpEndpointEnum = "disabled"
)

// Values returns all known values for HttpEndpointEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HttpEndpointEnum) Values() []HttpEndpointEnum {
	return []HttpEndpointEnum{
		"enabled",
		"disabled",
	}
}

type HttpProtocolIpv6Enum string

// Enum values for HttpProtocolIpv6Enum
const (
	HttpProtocolIpv6EnumEnabled  HttpProtocolIpv6Enum = "enabled"
	HttpProtocolIpv6EnumDisabled HttpProtocolIpv6Enum = "disabled"
)

// Values returns all known values for HttpProtocolIpv6Enum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HttpProtocolIpv6Enum) Values() []HttpProtocolIpv6Enum {
	return []HttpProtocolIpv6Enum{
		"enabled",
		"disabled",
	}
}

type HttpTokensEnum string

// Enum values for HttpTokensEnum
const (
	HttpTokensEnumOptional HttpTokensEnum = "optional"
	HttpTokensEnumRequired HttpTokensEnum = "required"
)

// Values returns all known values for HttpTokensEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HttpTokensEnum) Values() []HttpTokensEnum {
	return []HttpTokensEnum{
		"optional",
		"required",
	}
}

type InstanceInterruptionBehaviorEnum string

// Enum values for InstanceInterruptionBehaviorEnum
const (
	InstanceInterruptionBehaviorEnumHibernate InstanceInterruptionBehaviorEnum = "hibernate"
	InstanceInterruptionBehaviorEnumStop      InstanceInterruptionBehaviorEnum = "stop"
)

// Values returns all known values for InstanceInterruptionBehaviorEnum. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InstanceInterruptionBehaviorEnum) Values() []InstanceInterruptionBehaviorEnum {
	return []InstanceInterruptionBehaviorEnum{
		"hibernate",
		"stop",
	}
}

type InstanceMetadataTagsEnum string

// Enum values for InstanceMetadataTagsEnum
const (
	InstanceMetadataTagsEnumEnabled  InstanceMetadataTagsEnum = "enabled"
	InstanceMetadataTagsEnumDisabled InstanceMetadataTagsEnum = "disabled"
)

// Values returns all known values for InstanceMetadataTagsEnum. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InstanceMetadataTagsEnum) Values() []InstanceMetadataTagsEnum {
	return []InstanceMetadataTagsEnum{
		"enabled",
		"disabled",
	}
}

type InterfaceTypeEnum string

// Enum values for InterfaceTypeEnum
const (
	InterfaceTypeEnumInterface InterfaceTypeEnum = "interface"
	InterfaceTypeEnumEfa       InterfaceTypeEnum = "efa"
	InterfaceTypeEnumEfaOnly   InterfaceTypeEnum = "efa-only"
)

// Values returns all known values for InterfaceTypeEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InterfaceTypeEnum) Values() []InterfaceTypeEnum {
	return []InterfaceTypeEnum{
		"interface",
		"efa",
		"efa-only",
	}
}

type MarketTypeEnum string

// Enum values for MarketTypeEnum
const (
	MarketTypeEnumSpot          MarketTypeEnum = "spot"
	MarketTypeEnumCapacityBlock MarketTypeEnum = "capacity-block"
)

// Values returns all known values for MarketTypeEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MarketTypeEnum) Values() []MarketTypeEnum {
	return []MarketTypeEnum{
		"spot",
		"capacity-block",
	}
}

type ProvisionStateEnum string

// Enum values for ProvisionStateEnum
const (
	ProvisionStateEnumAllocating        ProvisionStateEnum = "ALLOCATING"
	ProvisionStateEnumAllocated         ProvisionStateEnum = "ALLOCATED"
	ProvisionStateEnumDeallocating      ProvisionStateEnum = "DEALLOCATING"
	ProvisionStateEnumDeallocated       ProvisionStateEnum = "DEALLOCATED"
	ProvisionStateEnumErrorAllocating   ProvisionStateEnum = "ERROR_ALLOCATING"
	ProvisionStateEnumErrorDeallocating ProvisionStateEnum = "ERROR_DEALLOCATING"
)

// Values returns all known values for ProvisionStateEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ProvisionStateEnum) Values() []ProvisionStateEnum {
	return []ProvisionStateEnum{
		"ALLOCATING",
		"ALLOCATED",
		"DEALLOCATING",
		"DEALLOCATED",
		"ERROR_ALLOCATING",
		"ERROR_DEALLOCATING",
	}
}

type ResourceTypeEnum string

// Enum values for ResourceTypeEnum
const (
	ResourceTypeEnumInstance             ResourceTypeEnum = "instance"
	ResourceTypeEnumVolume               ResourceTypeEnum = "volume"
	ResourceTypeEnumSpotInstancesRequest ResourceTypeEnum = "spot-instances-request"
	ResourceTypeEnumNetworkInterface     ResourceTypeEnum = "network-interface"
)

// Values returns all known values for ResourceTypeEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceTypeEnum) Values() []ResourceTypeEnum {
	return []ResourceTypeEnum{
		"instance",
		"volume",
		"spot-instances-request",
		"network-interface",
	}
}

type SpotInstanceTypeEnum string

// Enum values for SpotInstanceTypeEnum
const (
	SpotInstanceTypeEnumOneTime    SpotInstanceTypeEnum = "one-time"
	SpotInstanceTypeEnumPersistent SpotInstanceTypeEnum = "persistent"
)

// Values returns all known values for SpotInstanceTypeEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SpotInstanceTypeEnum) Values() []SpotInstanceTypeEnum {
	return []SpotInstanceTypeEnum{
		"one-time",
		"persistent",
	}
}

type TenancyEnum string

// Enum values for TenancyEnum
const (
	TenancyEnumDefault   TenancyEnum = "default"
	TenancyEnumDedicated TenancyEnum = "dedicated"
	TenancyEnumHost      TenancyEnum = "host"
)

// Values returns all known values for TenancyEnum. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TenancyEnum) Values() []TenancyEnum {
	return []TenancyEnum{
		"default",
		"dedicated",
		"host",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "UNKNOWN_OPERATION"
	ValidationExceptionReasonUnsupportedOperation  ValidationExceptionReason = "UNSUPPORTED_OPERATION"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "CANNOT_PARSE"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReasonDependencyFailure     ValidationExceptionReason = "DEPENDENCY_FAILURE"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "OTHER"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"UNKNOWN_OPERATION",
		"UNSUPPORTED_OPERATION",
		"CANNOT_PARSE",
		"FIELD_VALIDATION_FAILED",
		"DEPENDENCY_FAILURE",
		"OTHER",
	}
}

type VolumeTypeEnum string

// Enum values for VolumeTypeEnum
const (
	VolumeTypeEnumStandard VolumeTypeEnum = "standard"
	VolumeTypeEnumIo1      VolumeTypeEnum = "io1"
	VolumeTypeEnumIo2      VolumeTypeEnum = "io2"
	VolumeTypeEnumGp2      VolumeTypeEnum = "gp2"
	VolumeTypeEnumSc1      VolumeTypeEnum = "sc1"
	VolumeTypeEnumSt1      VolumeTypeEnum = "st1"
	VolumeTypeEnumGp3      VolumeTypeEnum = "gp3"
)

// Values returns all known values for VolumeTypeEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (VolumeTypeEnum) Values() []VolumeTypeEnum {
	return []VolumeTypeEnum{
		"standard",
		"io1",
		"io2",
		"gp2",
		"sc1",
		"st1",
		"gp3",
	}
}
