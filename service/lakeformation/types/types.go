// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A new object to add to the governed table.
type AddObjectInput struct {

	// The Amazon S3 ETag of the object. Returned by GetTableObjects for validation
	// and used to identify changes to the underlying data.
	//
	// This member is required.
	ETag *string

	// The size of the Amazon S3 object in bytes.
	//
	// This member is required.
	Size int64

	// The Amazon S3 location of the object.
	//
	// This member is required.
	Uri *string

	// A list of partition values for the object. A value must be specified for each
	// partition key associated with the table.
	//
	// The supported data types are integer, long, date(yyyy-MM-dd),
	// timestamp(yyyy-MM-dd HH:mm:ssXXX or yyyy-MM-dd HH:mm:ss"), string and decimal.
	PartitionValues []string

	noSmithyDocumentSerde
}

// A structure that you pass to indicate you want all rows in a filter.
type AllRowsWildcard struct {
	noSmithyDocumentSerde
}

// A structure used to include auditing information on the privileged API.
type AuditContext struct {

	// The filter engine can populate the 'AdditionalAuditContext' information with
	// the request ID for you to track. This information will be displayed in
	// CloudTrail log in your account.
	AdditionalAuditContext *string

	noSmithyDocumentSerde
}

// A list of failures when performing a batch grant or batch revoke operation.
type BatchPermissionsFailureEntry struct {

	// An error message that applies to the failure of the entry.
	Error *ErrorDetail

	// An identifier for an entry of the batch request.
	RequestEntry *BatchPermissionsRequestEntry

	noSmithyDocumentSerde
}

// A permission to a resource granted by batch operation to the principal.
type BatchPermissionsRequestEntry struct {

	// A unique identifier for the batch permissions request entry.
	//
	// This member is required.
	Id *string

	// A Lake Formation condition, which applies to permissions and opt-ins that
	// contain an expression.
	Condition *Condition

	// The permissions to be granted.
	Permissions []Permission

	// Indicates if the option to pass permissions is granted.
	PermissionsWithGrantOption []Permission

	// The principal to be granted a permission.
	Principal *DataLakePrincipal

	// The resource to which the principal is to be granted a permission.
	Resource *Resource

	noSmithyDocumentSerde
}

// A structure for the catalog object.
type CatalogResource struct {

	// An identifier for the catalog resource.
	Id *string

	noSmithyDocumentSerde
}

// A structure containing the name of a column resource and the LF-tags attached
// to it.
type ColumnLFTag struct {

	// The LF-tags attached to a column resource.
	LFTags []LFTagPair

	// The name of a column resource.
	Name *string

	noSmithyDocumentSerde
}

// A wildcard object, consisting of an optional list of excluded column names or
// indexes.
type ColumnWildcard struct {

	// Excludes column names. Any column with this name will be excluded.
	ExcludedColumnNames []string

	noSmithyDocumentSerde
}

// A Lake Formation condition, which applies to permissions and opt-ins that
// contain an expression.
type Condition struct {

	// An expression written based on the Cedar Policy Language used to match the
	// principal attributes.
	Expression *string

	noSmithyDocumentSerde
}

// A structure for the database object.
type DatabaseResource struct {

	// The name of the database resource. Unique to the Data Catalog.
	//
	// This member is required.
	Name *string

	// The identifier for the Data Catalog. By default, it is the account ID of the
	// caller.
	CatalogId *string

	noSmithyDocumentSerde
}

// A structure that describes certain columns on certain rows.
type DataCellsFilter struct {

	// A database in the Glue Data Catalog.
	//
	// This member is required.
	DatabaseName *string

	// The name given by the user to the data filter cell.
	//
	// This member is required.
	Name *string

	// The ID of the catalog to which the table belongs.
	//
	// This member is required.
	TableCatalogId *string

	// A table in the database.
	//
	// This member is required.
	TableName *string

	// A list of column names and/or nested column attributes. When specifying nested
	// attributes, use a qualified dot (.) delimited format such as "address"."zip".
	// Nested attributes within this list may not exceed a depth of 5.
	ColumnNames []string

	// A wildcard with exclusions.
	//
	// You must specify either a ColumnNames list or the ColumnWildCard .
	ColumnWildcard *ColumnWildcard

	// A PartiQL predicate.
	RowFilter *RowFilter

	// The ID of the data cells filter version.
	VersionId *string

	noSmithyDocumentSerde
}

// A structure for a data cells filter resource.
type DataCellsFilterResource struct {

	// A database in the Glue Data Catalog.
	DatabaseName *string

	// The name of the data cells filter.
	Name *string

	// The ID of the catalog to which the table belongs.
	TableCatalogId *string

	// The name of the table.
	TableName *string

	noSmithyDocumentSerde
}

// The Lake Formation principal. Supported principals are IAM users or IAM roles.
type DataLakePrincipal struct {

	// An identifier for the Lake Formation principal.
	DataLakePrincipalIdentifier *string

	noSmithyDocumentSerde
}

// A structure representing a list of Lake Formation principals designated as data
// lake administrators and lists of principal permission entries for default create
// database and default create table permissions.
type DataLakeSettings struct {

	// Whether to allow Amazon EMR clusters to access data managed by Lake Formation.
	//
	// If true, you allow Amazon EMR clusters to access data in Amazon S3 locations
	// that are registered with Lake Formation.
	//
	// If false or null, no Amazon EMR clusters will be able to access data in Amazon
	// S3 locations that are registered with Lake Formation.
	//
	// For more information, see [(Optional) Allow external data filtering].
	//
	// [(Optional) Allow external data filtering]: https://docs.aws.amazon.com/lake-formation/latest/dg/initial-LF-setup.html#external-data-filter
	AllowExternalDataFiltering *bool

	// Whether to allow a third-party query engine to get data access credentials
	// without session tags when a caller has full data access permissions.
	AllowFullTableExternalDataAccess *bool

	// Lake Formation relies on a privileged process secured by Amazon EMR or the
	// third party integrator to tag the user's role while assuming it. Lake Formation
	// will publish the acceptable key-value pair, for example key =
	// "LakeFormationTrustedCaller" and value = "TRUE" and the third party integrator
	// must properly tag the temporary security credentials that will be used to call
	// Lake Formation's administrative APIs.
	AuthorizedSessionTagValueList []string

	// Specifies whether access control on newly created database is managed by Lake
	// Formation permissions or exclusively by IAM permissions.
	//
	// A null value indicates access control by Lake Formation permissions. A value
	// that assigns ALL to IAM_ALLOWED_PRINCIPALS indicates access control by IAM
	// permissions. This is referred to as the setting "Use only IAM access control,"
	// and is for backward compatibility with the Glue permission model implemented by
	// IAM permissions.
	//
	// The only permitted values are an empty array or an array that contains a single
	// JSON object that grants ALL to IAM_ALLOWED_PRINCIPALS.
	//
	// For more information, see [Changing the Default Security Settings for Your Data Lake].
	//
	// [Changing the Default Security Settings for Your Data Lake]: https://docs.aws.amazon.com/lake-formation/latest/dg/change-settings.html
	CreateDatabaseDefaultPermissions []PrincipalPermissions

	// Specifies whether access control on newly created table is managed by Lake
	// Formation permissions or exclusively by IAM permissions.
	//
	// A null value indicates access control by Lake Formation permissions. A value
	// that assigns ALL to IAM_ALLOWED_PRINCIPALS indicates access control by IAM
	// permissions. This is referred to as the setting "Use only IAM access control,"
	// and is for backward compatibility with the Glue permission model implemented by
	// IAM permissions.
	//
	// The only permitted values are an empty array or an array that contains a single
	// JSON object that grants ALL to IAM_ALLOWED_PRINCIPALS.
	//
	// For more information, see [Changing the Default Security Settings for Your Data Lake].
	//
	// [Changing the Default Security Settings for Your Data Lake]: https://docs.aws.amazon.com/lake-formation/latest/dg/change-settings.html
	CreateTableDefaultPermissions []PrincipalPermissions

	// A list of Lake Formation principals. Supported principals are IAM users or IAM
	// roles.
	DataLakeAdmins []DataLakePrincipal

	// A list of the account IDs of Amazon Web Services accounts with Amazon EMR
	// clusters that are to perform data filtering.>
	ExternalDataFilteringAllowList []DataLakePrincipal

	// A key-value map that provides an additional configuration on your data lake.
	// CROSS_ACCOUNT_VERSION is the key you can configure in the Parameters field.
	// Accepted values for the CrossAccountVersion key are 1, 2, 3, and 4.
	Parameters map[string]string

	// A list of Lake Formation principals with only view access to the resources,
	// without the ability to make changes. Supported principals are IAM users or IAM
	// roles.
	ReadOnlyAdmins []DataLakePrincipal

	// A list of the resource-owning account IDs that the caller's account can use to
	// share their user access details (user ARNs). The user ARNs can be logged in the
	// resource owner's CloudTrail log.
	//
	// You may want to specify this property when you are in a high-trust boundary,
	// such as the same team or company.
	TrustedResourceOwners []string

	noSmithyDocumentSerde
}

// A structure for a data location object where permissions are granted or
// revoked.
type DataLocationResource struct {

	// The Amazon Resource Name (ARN) that uniquely identifies the data location
	// resource.
	//
	// This member is required.
	ResourceArn *string

	// The identifier for the Data Catalog where the location is registered with Lake
	// Formation. By default, it is the account ID of the caller.
	CatalogId *string

	noSmithyDocumentSerde
}

// An object to delete from the governed table.
type DeleteObjectInput struct {

	// The Amazon S3 location of the object to delete.
	//
	// This member is required.
	Uri *string

	// The Amazon S3 ETag of the object. Returned by GetTableObjects for validation
	// and used to identify changes to the underlying data.
	ETag *string

	// A list of partition values for the object. A value must be specified for each
	// partition key associated with the governed table.
	PartitionValues []string

	noSmithyDocumentSerde
}

// A structure containing the additional details to be returned in the
// AdditionalDetails attribute of PrincipalResourcePermissions .
//
// If a catalog resource is shared through Resource Access Manager (RAM), then
// there will exist a corresponding RAM resource share ARN.
type DetailsMap struct {

	// A resource share ARN for a catalog resource shared through RAM.
	ResourceShare []string

	noSmithyDocumentSerde
}

// Contains details about an error.
type ErrorDetail struct {

	// The code associated with this error.
	ErrorCode *string

	// A message describing the error.
	ErrorMessage *string

	noSmithyDocumentSerde
}

// Statistics related to the processing of a query statement.
type ExecutionStatistics struct {

	// The average time the request took to be executed.
	AverageExecutionTimeMillis int64

	// The amount of data that was scanned in bytes.
	DataScannedBytes int64

	// The number of work units executed.
	WorkUnitsExecutedCount int64

	noSmithyDocumentSerde
}

// Configuration for enabling external data filtering for third-party applications
// to access data managed by Lake Formation .
type ExternalFilteringConfiguration struct {

	// List of third-party application ARNs integrated with Lake Formation.
	//
	// This member is required.
	AuthorizedTargets []string

	// Allows to enable or disable the third-party applications that are allowed to
	// access data managed by Lake Formation.
	//
	// This member is required.
	Status EnableStatus

	noSmithyDocumentSerde
}

// This structure describes the filtering of columns in a table based on a filter
// condition.
type FilterCondition struct {

	// The comparison operator used in the filter condition.
	ComparisonOperator ComparisonOperator

	// The field to filter in the filter condition.
	Field FieldNameString

	// A string with values used in evaluating the filter condition.
	StringValueList []string

	noSmithyDocumentSerde
}

// A single principal-resource pair that has Lake Formation permissins enforced.
type LakeFormationOptInsInfo struct {

	// A Lake Formation condition, which applies to permissions and opt-ins that
	// contain an expression.
	Condition *Condition

	// The last modified date and time of the record.
	LastModified *time.Time

	// The user who updated the record.
	LastUpdatedBy *string

	// The Lake Formation principal. Supported principals are IAM users or IAM roles.
	Principal *DataLakePrincipal

	// A structure for the resource.
	Resource *Resource

	noSmithyDocumentSerde
}

// A structure that allows an admin to grant user permissions on certain
// conditions. For example, granting a role access to all columns that do not have
// the LF-tag 'PII' in tables that have the LF-tag 'Prod'.
type LFTag struct {

	// The key-name for the LF-tag.
	//
	// This member is required.
	TagKey *string

	// A list of possible values an attribute can take.
	//
	// The maximum number of values that can be defined for a LF-Tag is 1000. A single
	// API call supports 50 values. You can use multiple API calls to add more values.
	//
	// This member is required.
	TagValues []string

	noSmithyDocumentSerde
}

// A structure containing an error related to a TagResource or UnTagResource
// operation.
type LFTagError struct {

	// An error that occurred with the attachment or detachment of the LF-tag.
	Error *ErrorDetail

	// The key-name of the LF-tag.
	LFTag *LFTagPair

	noSmithyDocumentSerde
}

// A structure consists LF-Tag expression name and catalog ID.
type LFTagExpression struct {

	// The identifier for the Data Catalog. By default, the account ID.
	CatalogId *string

	// A structure that contains information about the LF-Tag expression.
	Description *string

	// A logical expression composed of one or more LF-Tags.
	Expression []LFTag

	// The name for saved the LF-Tag expression.
	Name *string

	noSmithyDocumentSerde
}

// A structure containing a LF-Tag expression (keys and values).
type LFTagExpressionResource struct {

	// The name of the LF-Tag expression to grant permissions on.
	//
	// This member is required.
	Name *string

	// The identifier for the Data Catalog. By default, the account ID.
	CatalogId *string

	noSmithyDocumentSerde
}

// A structure containing an LF-tag key and values for a resource.
type LFTagKeyResource struct {

	// The key-name for the LF-tag.
	//
	// This member is required.
	TagKey *string

	// A list of possible values an attribute can take.
	//
	// This member is required.
	TagValues []string

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your Lake Formation
	// environment.
	CatalogId *string

	noSmithyDocumentSerde
}

// A structure containing an LF-tag key-value pair.
type LFTagPair struct {

	// The key-name for the LF-tag.
	//
	// This member is required.
	TagKey *string

	// A list of possible values an attribute can take.
	//
	// This member is required.
	TagValues []string

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your Lake Formation
	// environment.
	CatalogId *string

	noSmithyDocumentSerde
}

// A structure containing a list of LF-tag conditions or saved LF-Tag expressions
// that apply to a resource's LF-tag policy.
type LFTagPolicyResource struct {

	// The resource type for which the LF-tag policy applies.
	//
	// This member is required.
	ResourceType ResourceType

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your Lake Formation
	// environment.
	CatalogId *string

	// A list of LF-tag conditions or a saved expression that apply to the resource's
	// LF-tag policy.
	Expression []LFTag

	// If provided, permissions are granted to the Data Catalog resources whose
	// assigned LF-Tags match the expression body of the saved expression under the
	// provided ExpressionName .
	ExpressionName *string

	noSmithyDocumentSerde
}

// A structure containing a list of partition values and table objects.
type PartitionObjects struct {

	// A list of table objects
	Objects []TableObject

	// A list of partition values.
	PartitionValues []string

	noSmithyDocumentSerde
}

// Contains a list of values defining partitions.
type PartitionValueList struct {

	// The list of partition values.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// Statistics related to the processing of a query statement.
type PlanningStatistics struct {

	// An estimate of the data that was scanned in bytes.
	EstimatedDataToScanBytes int64

	// The time that it took to process the request.
	PlanningTimeMillis int64

	// The time the request was in queue to be processed.
	QueueTimeMillis int64

	// The number of work units generated.
	WorkUnitsGeneratedCount int64

	noSmithyDocumentSerde
}

// Permissions granted to a principal.
type PrincipalPermissions struct {

	// The permissions that are granted to the principal.
	Permissions []Permission

	// The principal who is granted permissions.
	Principal *DataLakePrincipal

	noSmithyDocumentSerde
}

// The permissions granted or revoked on a resource.
type PrincipalResourcePermissions struct {

	// This attribute can be used to return any additional details of
	// PrincipalResourcePermissions . Currently returns only as a RAM resource share
	// ARN.
	AdditionalDetails *DetailsMap

	// A Lake Formation condition, which applies to permissions and opt-ins that
	// contain an expression.
	Condition *Condition

	// The date and time when the resource was last updated.
	LastUpdated *time.Time

	// The user who updated the record.
	LastUpdatedBy *string

	// The permissions to be granted or revoked on the resource.
	Permissions []Permission

	// Indicates whether to grant the ability to grant permissions (as a subset of
	// permissions granted).
	PermissionsWithGrantOption []Permission

	// The Data Lake principal to be granted or revoked permissions.
	Principal *DataLakePrincipal

	// The resource where permissions are to be granted or revoked.
	Resource *Resource

	noSmithyDocumentSerde
}

// A structure containing information about the query plan.
type QueryPlanningContext struct {

	// The database containing the table.
	//
	// This member is required.
	DatabaseName *string

	// The ID of the Data Catalog where the partition in question resides. If none is
	// provided, the Amazon Web Services account ID is used by default.
	CatalogId *string

	// The time as of when to read the table contents. If not set, the most recent
	// transaction commit time will be used. Cannot be specified along with
	// TransactionId .
	QueryAsOfTime *time.Time

	// A map consisting of key-value pairs.
	QueryParameters map[string]string

	// The transaction ID at which to read the table contents. If this transaction is
	// not committed, the read will be treated as part of that transaction and will see
	// its writes. If this transaction has aborted, an error will be returned. If not
	// set, defaults to the most recent committed transaction. Cannot be specified
	// along with QueryAsOfTime .
	TransactionId *string

	noSmithyDocumentSerde
}

// A structure used as a protocol between query engines and Lake Formation or
// Glue. Contains both a Lake Formation generated authorization identifier and
// information from the request's authorization context.
type QuerySessionContext struct {

	// An opaque string-string map passed by the query engine.
	AdditionalContext map[string]string

	// An identifier string for the consumer cluster.
	ClusterId *string

	// A cryptographically generated query identifier generated by Glue or Lake
	// Formation.
	QueryAuthorizationId *string

	// A unique identifier generated by the query engine for the query.
	QueryId *string

	// A timestamp provided by the query engine for when the query started.
	QueryStartTime *time.Time

	noSmithyDocumentSerde
}

// A structure for the resource.
type Resource struct {

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your Lake Formation
	// environment.
	Catalog *CatalogResource

	// A data cell filter.
	DataCellsFilter *DataCellsFilterResource

	// The location of an Amazon S3 path where permissions are granted or revoked.
	DataLocation *DataLocationResource

	// The database for the resource. Unique to the Data Catalog. A database is a set
	// of associated table definitions organized into a logical group. You can Grant
	// and Revoke database permissions to a principal.
	Database *DatabaseResource

	// The LF-tag key and values attached to a resource.
	LFTag *LFTagKeyResource

	// LF-Tag expression resource. A logical expression composed of one or more LF-Tag
	// key:value pairs.
	LFTagExpression *LFTagExpressionResource

	// A list of LF-tag conditions or saved LF-Tag expressions that define a
	// resource's LF-tag policy.
	LFTagPolicy *LFTagPolicyResource

	// The table for the resource. A table is a metadata definition that represents
	// your data. You can Grant and Revoke table privileges to a principal.
	Table *TableResource

	// The table with columns for the resource. A principal with permissions to this
	// resource can select metadata from the columns of a table in the Data Catalog and
	// the underlying data in Amazon S3.
	TableWithColumns *TableWithColumnsResource

	noSmithyDocumentSerde
}

// A structure containing information about an Lake Formation resource.
type ResourceInfo struct {

	//  Indicates whether the data access of tables pointing to the location can be
	// managed by both Lake Formation permissions as well as Amazon S3 bucket policies.
	HybridAccessEnabled *bool

	// The date and time the resource was last modified.
	LastModified *time.Time

	// The Amazon Resource Name (ARN) of the resource.
	ResourceArn *string

	// The IAM role that registered a resource.
	RoleArn *string

	// Whether or not the resource is a federated resource.
	WithFederation *bool

	// Grants the calling principal the permissions to perform all supported Lake
	// Formation operations on the registered data location.
	WithPrivilegedAccess *bool

	noSmithyDocumentSerde
}

// A PartiQL predicate.
type RowFilter struct {

	// A wildcard for all rows.
	AllRowsWildcard *AllRowsWildcard

	// A filter expression.
	FilterExpression *string

	noSmithyDocumentSerde
}

// A structure describing the configuration and details of a storage optimizer.
type StorageOptimizer struct {

	// A map of the storage optimizer configuration. Currently contains only one
	// key-value pair: is_enabled indicates true or false for acceleration.
	Config map[string]string

	// A message that contains information about any error (if present).
	//
	// When an acceleration result has an enabled status, the error message is empty.
	//
	// When an acceleration result has a disabled status, the message describes an
	// error or simply indicates "disabled by the user".
	ErrorMessage *string

	// When an acceleration result has an enabled status, contains the details of the
	// last job run.
	LastRunDetails *string

	// The specific type of storage optimizer. The supported value is compaction .
	StorageOptimizerType OptimizerType

	// A message that contains information about any warnings (if present).
	Warnings *string

	noSmithyDocumentSerde
}

// Specifies the details of a governed table.
type TableObject struct {

	// The Amazon S3 ETag of the object. Returned by GetTableObjects for validation
	// and used to identify changes to the underlying data.
	ETag *string

	// The size of the Amazon S3 object in bytes.
	Size int64

	// The Amazon S3 location of the object.
	Uri *string

	noSmithyDocumentSerde
}

// A structure for the table object. A table is a metadata definition that
// represents your data. You can Grant and Revoke table privileges to a principal.
type TableResource struct {

	// The name of the database for the table. Unique to a Data Catalog. A database is
	// a set of associated table definitions organized into a logical group. You can
	// Grant and Revoke database privileges to a principal.
	//
	// This member is required.
	DatabaseName *string

	// The identifier for the Data Catalog. By default, it is the account ID of the
	// caller.
	CatalogId *string

	// The name of the table.
	Name *string

	// A wildcard object representing every table under a database.
	//
	// At least one of TableResource$Name or TableResource$TableWildcard is required.
	TableWildcard *TableWildcard

	noSmithyDocumentSerde
}

// A wildcard object representing every table under a database.
type TableWildcard struct {
	noSmithyDocumentSerde
}

// A structure for a table with columns object. This object is only used when
// granting a SELECT permission.
//
// This object must take a value for at least one of ColumnsNames , ColumnsIndexes
// , or ColumnsWildcard .
type TableWithColumnsResource struct {

	// The name of the database for the table with columns resource. Unique to the
	// Data Catalog. A database is a set of associated table definitions organized into
	// a logical group. You can Grant and Revoke database privileges to a principal.
	//
	// This member is required.
	DatabaseName *string

	// The name of the table resource. A table is a metadata definition that
	// represents your data. You can Grant and Revoke table privileges to a principal.
	//
	// This member is required.
	Name *string

	// The identifier for the Data Catalog. By default, it is the account ID of the
	// caller.
	CatalogId *string

	// The list of column names for the table. At least one of ColumnNames or
	// ColumnWildcard is required.
	ColumnNames []string

	// A wildcard specified by a ColumnWildcard object. At least one of ColumnNames or
	// ColumnWildcard is required.
	ColumnWildcard *ColumnWildcard

	noSmithyDocumentSerde
}

// A structure describing a database resource with LF-tags.
type TaggedDatabase struct {

	// A database that has LF-tags attached to it.
	Database *DatabaseResource

	// A list of LF-tags attached to the database.
	LFTags []LFTagPair

	noSmithyDocumentSerde
}

// A structure describing a table resource with LF-tags.
type TaggedTable struct {

	// A list of LF-tags attached to the database where the table resides.
	LFTagOnDatabase []LFTagPair

	// A list of LF-tags attached to columns in the table.
	LFTagsOnColumns []ColumnLFTag

	// A list of LF-tags attached to the table.
	LFTagsOnTable []LFTagPair

	// A table that has LF-tags attached to it.
	Table *TableResource

	noSmithyDocumentSerde
}

// A structure that contains information about a transaction.
type TransactionDescription struct {

	// The time when the transaction committed or aborted, if it is not currently
	// active.
	TransactionEndTime *time.Time

	// The ID of the transaction.
	TransactionId *string

	// The time when the transaction started.
	TransactionStartTime *time.Time

	// A status of ACTIVE, COMMITTED, or ABORTED.
	TransactionStatus TransactionStatus

	noSmithyDocumentSerde
}

// An object that defines an Amazon S3 object to be deleted if a transaction
// cancels, provided that VirtualPut was called before writing the object.
type VirtualObject struct {

	// The path to the Amazon S3 object. Must start with s3://
	//
	// This member is required.
	Uri *string

	// The ETag of the Amazon S3 object.
	ETag *string

	noSmithyDocumentSerde
}

// Defines the valid range of work unit IDs for querying the execution service.
type WorkUnitRange struct {

	// Defines the maximum work unit ID in the range. The maximum value is inclusive.
	//
	// This member is required.
	WorkUnitIdMax int64

	// Defines the minimum work unit ID in the range.
	//
	// This member is required.
	WorkUnitIdMin int64

	// A work token used to query the execution service.
	//
	// This member is required.
	WorkUnitToken *string

	noSmithyDocumentSerde
}

// Defines an object to add to or delete from a governed table.
type WriteOperation struct {

	// A new object to add to the governed table.
	AddObject *AddObjectInput

	// An object to delete from the governed table.
	DeleteObject *DeleteObjectInput

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
