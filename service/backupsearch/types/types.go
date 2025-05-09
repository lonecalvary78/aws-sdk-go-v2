// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// This filters by recovery points within the CreatedAfter and CreatedBefore
// timestamps.
type BackupCreationTimeFilter struct {

	// This timestamp includes recovery points only created after the specified time.
	CreatedAfter *time.Time

	// This timestamp includes recovery points only created before the specified time.
	CreatedBefore *time.Time

	noSmithyDocumentSerde
}

// This contains information results retrieved from a search job that may not have
// completed.
type CurrentSearchProgress struct {

	// This number is the sum of all items that match the item filters in a search job
	// in progress.
	ItemsMatchedCount *int64

	// This number is the sum of all items that have been scanned so far during a
	// search job.
	ItemsScannedCount *int64

	// This number is the sum of all backups that have been scanned so far during a
	// search job.
	RecoveryPointsScannedCount *int32

	noSmithyDocumentSerde
}

// This contains arrays of objects, which may include CreationTimes time condition
// objects, FilePaths string objects, LastModificationTimes time condition objects,
type EBSItemFilter struct {

	// You can include 1 to 10 values.
	//
	// If one is included, the results will return only items that match.
	//
	// If more than one is included, the results will return all items that match any
	// of the included values.
	CreationTimes []TimeCondition

	// You can include 1 to 10 values.
	//
	// If one file path is included, the results will return only items that match the
	// file path.
	//
	// If more than one file path is included, the results will return all items that
	// match any of the file paths.
	FilePaths []StringCondition

	// You can include 1 to 10 values.
	//
	// If one is included, the results will return only items that match.
	//
	// If more than one is included, the results will return all items that match any
	// of the included values.
	LastModificationTimes []TimeCondition

	// You can include 1 to 10 values.
	//
	// If one is included, the results will return only items that match.
	//
	// If more than one is included, the results will return all items that match any
	// of the included values.
	Sizes []LongCondition

	noSmithyDocumentSerde
}

// These are the items returned in the results of a search of Amazon EBS backup
// metadata.
type EBSResultItem struct {

	// These are one or more items in the results that match values for the Amazon
	// Resource Name (ARN) of recovery points returned in a search of Amazon EBS backup
	// metadata.
	BackupResourceArn *string

	// The name of the backup vault.
	BackupVaultName *string

	// These are one or more items in the results that match values for creation times
	// returned in a search of Amazon EBS backup metadata.
	CreationTime *time.Time

	// These are one or more items in the results that match values for file paths
	// returned in a search of Amazon EBS backup metadata.
	FilePath *string

	// These are one or more items in the results that match values for file sizes
	// returned in a search of Amazon EBS backup metadata.
	FileSize *int64

	// These are one or more items in the results that match values for file systems
	// returned in a search of Amazon EBS backup metadata.
	FileSystemIdentifier *string

	// These are one or more items in the results that match values for Last Modified
	// Time returned in a search of Amazon EBS backup metadata.
	LastModifiedTime *time.Time

	// These are one or more items in the results that match values for the Amazon
	// Resource Name (ARN) of source resources returned in a search of Amazon EBS
	// backup metadata.
	SourceResourceArn *string

	noSmithyDocumentSerde
}

// This is the summary of an export job.
type ExportJobSummary struct {

	// This is the unique string that identifies a specific export job.
	//
	// This member is required.
	ExportJobIdentifier *string

	// This is a timestamp of the time the export job compeleted.
	CompletionTime *time.Time

	// This is a timestamp of the time the export job was created.
	CreationTime *time.Time

	// This is the unique ARN (Amazon Resource Name) that belongs to the new export
	// job.
	ExportJobArn *string

	// The unique string that identifies the Amazon Resource Name (ARN) of the
	// specified search job.
	SearchJobArn *string

	// The status of the export job is one of the following:
	//
	// CREATED ; RUNNING ; FAILED ; or COMPLETED .
	Status ExportJobStatus

	// A status message is a string that is returned for an export job.
	//
	// A status message is included for any status other than COMPLETED without issues.
	StatusMessage *string

	noSmithyDocumentSerde
}

// This contains the export specification object.
//
// The following types satisfy this interface:
//
//	ExportSpecificationMemberS3ExportSpecification
type ExportSpecification interface {
	isExportSpecification()
}

// This specifies the destination Amazon S3 bucket for the export job. And, if
// included, it also specifies the destination prefix.
type ExportSpecificationMemberS3ExportSpecification struct {
	Value S3ExportSpecification

	noSmithyDocumentSerde
}

func (*ExportSpecificationMemberS3ExportSpecification) isExportSpecification() {}

// Item Filters represent all input item properties specified when the search was
// created.
//
// Contains either EBSItemFilters or S3ItemFilters
type ItemFilters struct {

	// This array can contain CreationTimes, FilePaths, LastModificationTimes, or
	// Sizes objects.
	EBSItemFilters []EBSItemFilter

	// This array can contain CreationTimes, ETags, ObjectKeys, Sizes, or VersionIds
	// objects.
	S3ItemFilters []S3ItemFilter

	noSmithyDocumentSerde
}

// The long condition contains a Value and can optionally contain an Operator .
type LongCondition struct {

	// The value of an item included in one of the search item filters.
	//
	// This member is required.
	Value *int64

	// A string that defines what values will be returned.
	//
	// If this is included, avoid combinations of operators that will return all
	// possible values. For example, including both EQUALS_TO and NOT_EQUALS_TO with a
	// value of 4 will return all values.
	Operator LongConditionOperator

	noSmithyDocumentSerde
}

// This is an object representing the item returned in the results of a search for
// a specific resource type.
//
// The following types satisfy this interface:
//
//	ResultItemMemberEBSResultItem
//	ResultItemMemberS3ResultItem
type ResultItem interface {
	isResultItem()
}

// These are items returned in the search results of an Amazon EBS search.
type ResultItemMemberEBSResultItem struct {
	Value EBSResultItem

	noSmithyDocumentSerde
}

func (*ResultItemMemberEBSResultItem) isResultItem() {}

// These are items returned in the search results of an Amazon S3 search.
type ResultItemMemberS3ResultItem struct {
	Value S3ResultItem

	noSmithyDocumentSerde
}

func (*ResultItemMemberS3ResultItem) isResultItem() {}

// This specification contains a required string of the destination bucket;
// optionally, you can include the destination prefix.
type S3ExportSpecification struct {

	// This specifies the destination Amazon S3 bucket for the export job.
	//
	// This member is required.
	DestinationBucket *string

	// This specifies the prefix for the destination Amazon S3 bucket for the export
	// job.
	DestinationPrefix *string

	noSmithyDocumentSerde
}

// This contains arrays of objects, which may include ObjectKeys, Sizes,
// CreationTimes, VersionIds, and/or Etags.
type S3ItemFilter struct {

	// You can include 1 to 10 values.
	//
	// If one value is included, the results will return only items that match the
	// value.
	//
	// If more than one value is included, the results will return all items that
	// match any of the values.
	CreationTimes []TimeCondition

	// You can include 1 to 10 values.
	//
	// If one value is included, the results will return only items that match the
	// value.
	//
	// If more than one value is included, the results will return all items that
	// match any of the values.
	ETags []StringCondition

	// You can include 1 to 10 values.
	//
	// If one value is included, the results will return only items that match the
	// value.
	//
	// If more than one value is included, the results will return all items that
	// match any of the values.
	ObjectKeys []StringCondition

	// You can include 1 to 10 values.
	//
	// If one value is included, the results will return only items that match the
	// value.
	//
	// If more than one value is included, the results will return all items that
	// match any of the values.
	Sizes []LongCondition

	// You can include 1 to 10 values.
	//
	// If one value is included, the results will return only items that match the
	// value.
	//
	// If more than one value is included, the results will return all items that
	// match any of the values.
	VersionIds []StringCondition

	noSmithyDocumentSerde
}

// These are the items returned in the results of a search of Amazon S3 backup
// metadata.
type S3ResultItem struct {

	// These are items in the returned results that match recovery point Amazon
	// Resource Names (ARN) input during a search of Amazon S3 backup metadata.
	BackupResourceArn *string

	// The name of the backup vault.
	BackupVaultName *string

	// These are one or more items in the returned results that match values for item
	// creation time input during a search of Amazon S3 backup metadata.
	CreationTime *time.Time

	// These are one or more items in the returned results that match values for ETags
	// input during a search of Amazon S3 backup metadata.
	ETag *string

	// This is one or more items returned in the results of a search of Amazon S3
	// backup metadata that match the values input for object key.
	ObjectKey *string

	// These are items in the returned results that match values for object size(s)
	// input during a search of Amazon S3 backup metadata.
	ObjectSize *int64

	// These are items in the returned results that match source Amazon Resource Names
	// (ARN) input during a search of Amazon S3 backup metadata.
	SourceResourceArn *string

	// These are one or more items in the returned results that match values for
	// version IDs input during a search of Amazon S3 backup metadata.
	VersionId *string

	noSmithyDocumentSerde
}

// This contains the information about recovery points returned in results of a
// search job.
type SearchJobBackupsResult struct {

	// This is the creation time of the backup (recovery point).
	BackupCreationTime *time.Time

	// The Amazon Resource Name (ARN) that uniquely identifies the backup resources.
	BackupResourceArn *string

	// This is the creation time of the backup index.
	IndexCreationTime *time.Time

	// This is the resource type of the search.
	ResourceType ResourceType

	// The Amazon Resource Name (ARN) that uniquely identifies the source resources.
	SourceResourceArn *string

	// This is the status of the search job backup result.
	Status SearchJobState

	// This is the status message included with the results.
	StatusMessage *string

	noSmithyDocumentSerde
}

// This is information pertaining to a search job.
type SearchJobSummary struct {

	// This is the completion time of the search job.
	CompletionTime *time.Time

	// This is the creation time of the search job.
	CreationTime *time.Time

	// This is the name of the search job.
	Name *string

	// The unique string that identifies the Amazon Resource Name (ARN) of the
	// specified search job.
	SearchJobArn *string

	// The unique string that specifies the search job.
	SearchJobIdentifier *string

	// Returned summary of the specified search job scope, including:
	//
	//   - TotalBackupsToScanCount, the number of recovery points returned by the
	//   search.
	//
	//   - TotalItemsToScanCount, the number of items returned by the search.
	SearchScopeSummary *SearchScopeSummary

	// This is the status of the search job.
	Status SearchJobState

	// A status message will be returned for either a earch job with a status of
	// ERRORED or a status of COMPLETED jobs with issues.
	//
	// For example, a message may say that a search contained recovery points unable
	// to be scanned because of a permissions issue.
	StatusMessage *string

	noSmithyDocumentSerde
}

// The search scope is all backup properties input into a search.
type SearchScope struct {

	// The resource types included in a search.
	//
	// Eligible resource types include S3 and EBS.
	//
	// This member is required.
	BackupResourceTypes []ResourceType

	// The Amazon Resource Name (ARN) that uniquely identifies the backup resources.
	BackupResourceArns []string

	// This is the time a backup resource was created.
	BackupResourceCreationTime *BackupCreationTimeFilter

	// These are one or more tags on the backup (recovery point).
	BackupResourceTags map[string]*string

	// The Amazon Resource Name (ARN) that uniquely identifies the source resources.
	SourceResourceArns []string

	noSmithyDocumentSerde
}

// The summary of the specified search job scope, including:
//
//   - TotalBackupsToScanCount, the number of recovery points returned by the
//     search.
//
//   - TotalItemsToScanCount, the number of items returned by the search.
type SearchScopeSummary struct {

	// This is the count of the total number of items that will be scanned in a search.
	TotalItemsToScanCount *int64

	// This is the count of the total number of backups that will be scanned in a
	// search.
	TotalRecoveryPointsToScanCount *int32

	noSmithyDocumentSerde
}

// This contains the value of the string and can contain one or more operators.
type StringCondition struct {

	// The value of the string.
	//
	// This member is required.
	Value *string

	// A string that defines what values will be returned.
	//
	// If this is included, avoid combinations of operators that will return all
	// possible values. For example, including both EQUALS_TO and NOT_EQUALS_TO with a
	// value of 4 will return all values.
	Operator StringConditionOperator

	noSmithyDocumentSerde
}

// A time condition denotes a creation time, last modification time, or other time.
type TimeCondition struct {

	// This is the timestamp value of the time condition.
	//
	// This member is required.
	Value *time.Time

	// A string that defines what values will be returned.
	//
	// If this is included, avoid combinations of operators that will return all
	// possible values. For example, including both EQUALS_TO and NOT_EQUALS_TO with a
	// value of 4 will return all values.
	Operator TimeConditionOperator

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isExportSpecification() {}
func (*UnknownUnionMember) isResultItem()          {}
