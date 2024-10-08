// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type GroupConfigurationStatus string

// Enum values for GroupConfigurationStatus
const (
	GroupConfigurationStatusUpdating       GroupConfigurationStatus = "UPDATING"
	GroupConfigurationStatusUpdateComplete GroupConfigurationStatus = "UPDATE_COMPLETE"
	GroupConfigurationStatusUpdateFailed   GroupConfigurationStatus = "UPDATE_FAILED"
)

// Values returns all known values for GroupConfigurationStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GroupConfigurationStatus) Values() []GroupConfigurationStatus {
	return []GroupConfigurationStatus{
		"UPDATING",
		"UPDATE_COMPLETE",
		"UPDATE_FAILED",
	}
}

type GroupFilterName string

// Enum values for GroupFilterName
const (
	GroupFilterNameResourceType      GroupFilterName = "resource-type"
	GroupFilterNameConfigurationType GroupFilterName = "configuration-type"
	GroupFilterNameOwner             GroupFilterName = "owner"
	GroupFilterNameDisplayName       GroupFilterName = "display-name"
	GroupFilterNameCriticality       GroupFilterName = "criticality"
)

// Values returns all known values for GroupFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GroupFilterName) Values() []GroupFilterName {
	return []GroupFilterName{
		"resource-type",
		"configuration-type",
		"owner",
		"display-name",
		"criticality",
	}
}

type GroupingStatus string

// Enum values for GroupingStatus
const (
	GroupingStatusSuccess    GroupingStatus = "SUCCESS"
	GroupingStatusFailed     GroupingStatus = "FAILED"
	GroupingStatusInProgress GroupingStatus = "IN_PROGRESS"
	GroupingStatusSkipped    GroupingStatus = "SKIPPED"
)

// Values returns all known values for GroupingStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GroupingStatus) Values() []GroupingStatus {
	return []GroupingStatus{
		"SUCCESS",
		"FAILED",
		"IN_PROGRESS",
		"SKIPPED",
	}
}

type GroupingType string

// Enum values for GroupingType
const (
	GroupingTypeGroup   GroupingType = "GROUP"
	GroupingTypeUngroup GroupingType = "UNGROUP"
)

// Values returns all known values for GroupingType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GroupingType) Values() []GroupingType {
	return []GroupingType{
		"GROUP",
		"UNGROUP",
	}
}

type GroupLifecycleEventsDesiredStatus string

// Enum values for GroupLifecycleEventsDesiredStatus
const (
	GroupLifecycleEventsDesiredStatusActive   GroupLifecycleEventsDesiredStatus = "ACTIVE"
	GroupLifecycleEventsDesiredStatusInactive GroupLifecycleEventsDesiredStatus = "INACTIVE"
)

// Values returns all known values for GroupLifecycleEventsDesiredStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GroupLifecycleEventsDesiredStatus) Values() []GroupLifecycleEventsDesiredStatus {
	return []GroupLifecycleEventsDesiredStatus{
		"ACTIVE",
		"INACTIVE",
	}
}

type GroupLifecycleEventsStatus string

// Enum values for GroupLifecycleEventsStatus
const (
	GroupLifecycleEventsStatusActive     GroupLifecycleEventsStatus = "ACTIVE"
	GroupLifecycleEventsStatusInactive   GroupLifecycleEventsStatus = "INACTIVE"
	GroupLifecycleEventsStatusInProgress GroupLifecycleEventsStatus = "IN_PROGRESS"
	GroupLifecycleEventsStatusError      GroupLifecycleEventsStatus = "ERROR"
)

// Values returns all known values for GroupLifecycleEventsStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GroupLifecycleEventsStatus) Values() []GroupLifecycleEventsStatus {
	return []GroupLifecycleEventsStatus{
		"ACTIVE",
		"INACTIVE",
		"IN_PROGRESS",
		"ERROR",
	}
}

type ListGroupingStatusesFilterName string

// Enum values for ListGroupingStatusesFilterName
const (
	ListGroupingStatusesFilterNameStatus      ListGroupingStatusesFilterName = "status"
	ListGroupingStatusesFilterNameResourceArn ListGroupingStatusesFilterName = "resource-arn"
)

// Values returns all known values for ListGroupingStatusesFilterName. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ListGroupingStatusesFilterName) Values() []ListGroupingStatusesFilterName {
	return []ListGroupingStatusesFilterName{
		"status",
		"resource-arn",
	}
}

type QueryErrorCode string

// Enum values for QueryErrorCode
const (
	QueryErrorCodeCloudformationStackInactive        QueryErrorCode = "CLOUDFORMATION_STACK_INACTIVE"
	QueryErrorCodeCloudformationStackNotExisting     QueryErrorCode = "CLOUDFORMATION_STACK_NOT_EXISTING"
	QueryErrorCodeCloudformationStackUnassumableRole QueryErrorCode = "CLOUDFORMATION_STACK_UNASSUMABLE_ROLE"
	QueryErrorCodeResourceTypeNotSupported           QueryErrorCode = "RESOURCE_TYPE_NOT_SUPPORTED"
)

// Values returns all known values for QueryErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QueryErrorCode) Values() []QueryErrorCode {
	return []QueryErrorCode{
		"CLOUDFORMATION_STACK_INACTIVE",
		"CLOUDFORMATION_STACK_NOT_EXISTING",
		"CLOUDFORMATION_STACK_UNASSUMABLE_ROLE",
		"RESOURCE_TYPE_NOT_SUPPORTED",
	}
}

type QueryType string

// Enum values for QueryType
const (
	QueryTypeTagFilters10          QueryType = "TAG_FILTERS_1_0"
	QueryTypeCloudformationStack10 QueryType = "CLOUDFORMATION_STACK_1_0"
)

// Values returns all known values for QueryType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QueryType) Values() []QueryType {
	return []QueryType{
		"TAG_FILTERS_1_0",
		"CLOUDFORMATION_STACK_1_0",
	}
}

type ResourceFilterName string

// Enum values for ResourceFilterName
const (
	ResourceFilterNameResourceType ResourceFilterName = "resource-type"
)

// Values returns all known values for ResourceFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceFilterName) Values() []ResourceFilterName {
	return []ResourceFilterName{
		"resource-type",
	}
}

type ResourceStatusValue string

// Enum values for ResourceStatusValue
const (
	ResourceStatusValuePending ResourceStatusValue = "PENDING"
)

// Values returns all known values for ResourceStatusValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceStatusValue) Values() []ResourceStatusValue {
	return []ResourceStatusValue{
		"PENDING",
	}
}

type TagSyncTaskStatus string

// Enum values for TagSyncTaskStatus
const (
	TagSyncTaskStatusActive TagSyncTaskStatus = "ACTIVE"
	TagSyncTaskStatusError  TagSyncTaskStatus = "ERROR"
)

// Values returns all known values for TagSyncTaskStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TagSyncTaskStatus) Values() []TagSyncTaskStatus {
	return []TagSyncTaskStatus{
		"ACTIVE",
		"ERROR",
	}
}
