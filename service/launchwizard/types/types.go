// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A field that details a condition of the specifications for a deployment.
type DeploymentConditionalField struct {

	// The comparator of the condition.
	//
	// Valid values: Equal | NotEqual
	Comparator *string

	// The name of the deployment condition.
	Name *string

	// The value of the condition.
	Value *string

	noSmithyDocumentSerde
}

// The data associated with a deployment.
type DeploymentData struct {

	// The time the deployment was created.
	CreatedAt *time.Time

	// The time the deployment was deleted.
	DeletedAt *time.Time

	// The Amazon Resource Name (ARN) of the deployment.
	DeploymentArn *string

	// The ID of the deployment.
	Id *string

	// The name of the deployment.
	Name *string

	// The pattern name of the deployment.
	PatternName *string

	// The resource group of the deployment.
	ResourceGroup *string

	// The settings specified for the deployment. These settings define how to deploy
	// and configure your resources created by the deployment. For more information
	// about the specifications required for creating a deployment for a SAP workload,
	// see [SAP deployment specifications]. To retrieve the specifications required to create a deployment for other
	// workloads, use the [GetWorkloadDeploymentPattern]GetWorkloadDeploymentPattern operation.
	//
	// [GetWorkloadDeploymentPattern]: https://docs.aws.amazon.com/launchwizard/latest/APIReference/API_GetWorkloadDeploymentPattern.html
	// [SAP deployment specifications]: https://docs.aws.amazon.com/launchwizard/latest/APIReference/launch-wizard-specifications-sap.html
	Specifications map[string]string

	// The status of the deployment.
	Status DeploymentStatus

	// Information about the tags attached to a deployment.
	Tags map[string]string

	// The name of the workload.
	WorkloadName *string

	noSmithyDocumentSerde
}

// A summary of the deployment data.
type DeploymentDataSummary struct {

	// The time the deployment was created.
	CreatedAt *time.Time

	// The ID of the deployment.
	Id *string

	// The name of the deployment
	Name *string

	// The name of the workload deployment pattern.
	PatternName *string

	// The status of the deployment.
	Status DeploymentStatus

	// The name of the workload.
	WorkloadName *string

	noSmithyDocumentSerde
}

// A summary of the deployment event data.
type DeploymentEventDataSummary struct {

	// The description of the deployment event.
	Description *string

	// The name of the deployment event.
	Name *string

	// The status of the deployment event.
	Status EventStatus

	// The reason of the deployment event status.
	StatusReason *string

	// The timestamp of the deployment event.
	Timestamp *time.Time

	noSmithyDocumentSerde
}

// A filter name and value pair that is used to return more specific results from
// a describe operation. Filters can be used to match a set of resources by
// specific criteria.
type DeploymentFilter struct {

	// The name of the filter. Filter names are case-sensitive.
	Name DeploymentFilterKey

	// The filter values. Filter values are case-sensitive. If you specify multiple
	// values for a filter, the values are joined with an OR , and the request returns
	// all results that match any of the specified values.
	Values []string

	noSmithyDocumentSerde
}

// A field that details a specification of a deployment pattern.
type DeploymentSpecificationsField struct {

	// The allowed values of the deployment specification.
	AllowedValues []string

	// The conditionals used for the deployment specification.
	Conditionals []DeploymentConditionalField

	// The description of the deployment specification.
	Description *string

	// The name of the deployment specification.
	Name *string

	// Indicates if the deployment specification is required.
	Required *string

	noSmithyDocumentSerde
}

// Describes a workload.
type WorkloadData struct {

	// The description of a workload.
	Description *string

	// The display name of a workload.
	DisplayName *string

	// The URL of a workload document.
	DocumentationUrl *string

	// The URL of a workload icon.
	IconUrl *string

	// The status of a workload.
	Status WorkloadStatus

	// The message about a workload's status.
	StatusMessage *string

	// The name of the workload.
	WorkloadName *string

	noSmithyDocumentSerde
}

// Describes workload data.
type WorkloadDataSummary struct {

	// The display name of the workload data.
	DisplayName *string

	// The name of the workload.
	WorkloadName *string

	noSmithyDocumentSerde
}

// The data that details a workload deployment pattern.
type WorkloadDeploymentPatternData struct {

	// The name of the deployment pattern.
	DeploymentPatternName *string

	// The description of the deployment pattern.
	Description *string

	// The display name of the deployment pattern.
	DisplayName *string

	// The settings specified for the deployment. These settings define how to deploy
	// and configure your resources created by the deployment. For more information
	// about the specifications required for creating a deployment for a SAP workload,
	// see [SAP deployment specifications]. To retrieve the specifications required to create a deployment for other
	// workloads, use the [GetWorkloadDeploymentPattern]GetWorkloadDeploymentPattern operation.
	//
	// [GetWorkloadDeploymentPattern]: https://docs.aws.amazon.com/launchwizard/latest/APIReference/API_GetWorkloadDeploymentPattern.html
	// [SAP deployment specifications]: https://docs.aws.amazon.com/launchwizard/latest/APIReference/launch-wizard-specifications-sap.html
	Specifications []DeploymentSpecificationsField

	// The status of the deployment pattern.
	Status WorkloadDeploymentPatternStatus

	// The status message of the deployment pattern.
	StatusMessage *string

	// The workload name of the deployment pattern.
	WorkloadName *string

	// The workload version name of the deployment pattern.
	WorkloadVersionName *string

	noSmithyDocumentSerde
}

// Describes a workload deployment pattern.
type WorkloadDeploymentPatternDataSummary struct {

	// The name of a workload deployment pattern.
	DeploymentPatternName *string

	// The description of a workload deployment pattern.
	Description *string

	// The display name of a workload deployment pattern.
	DisplayName *string

	// The status of a workload deployment pattern.
	Status WorkloadDeploymentPatternStatus

	// A message about a workload deployment pattern's status.
	StatusMessage *string

	// The name of the workload.
	WorkloadName *string

	// The name of the workload deployment pattern version.
	WorkloadVersionName *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
