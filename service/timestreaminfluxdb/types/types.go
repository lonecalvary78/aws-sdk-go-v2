// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// Contains a summary of a DB instance.
type DbInstanceSummary struct {

	// The Amazon Resource Name (ARN) of the DB instance.
	//
	// This member is required.
	Arn *string

	// The service-generated unique identifier of the DB instance.
	//
	// This member is required.
	Id *string

	// This customer-supplied name uniquely identifies the DB instance when
	// interacting with the Amazon Timestream for InfluxDB API and AWS CLI commands.
	//
	// This member is required.
	Name *string

	// The amount of storage to allocate for your DbStorageType in GiB (gibibytes).
	AllocatedStorage *int32

	// The Timestream for InfluxDB instance type to run InfluxDB on.
	DbInstanceType DbInstanceType

	// The storage type for your DB instance.
	DbStorageType DbStorageType

	// Single-Instance or with a MultiAZ Standby for High availability.
	DeploymentType DeploymentType

	// The endpoint used to connect to InfluxDB. The default InfluxDB port is 8086.
	Endpoint *string

	// The status of the DB instance.
	Status Status

	noSmithyDocumentSerde
}

// Contains a summary of a DB parameter group.
type DbParameterGroupSummary struct {

	// The Amazon Resource Name (ARN) of the DB parameter group.
	//
	// This member is required.
	Arn *string

	// A service-generated unique identifier.
	//
	// This member is required.
	Id *string

	// This customer-supplied name uniquely identifies the parameter group.
	//
	// This member is required.
	Name *string

	// A description of the DB parameter group.
	Description *string

	noSmithyDocumentSerde
}

// All the customer-modifiable InfluxDB v2 parameters in Timestream for InfluxDB.
type InfluxDBv2Parameters struct {

	// Include option to show detailed logs for Flux queries.
	//
	// Default: false
	FluxLogEnabled *bool

	// Log output level. InfluxDB outputs log entries with severity levels greater
	// than or equal to the level specified.
	//
	// Default: info
	LogLevel LogLevel

	// Disable the HTTP /metrics endpoint which exposes [internal InfluxDB metrics].
	//
	// Default: false
	//
	// [internal InfluxDB metrics]: https://docs.influxdata.com/influxdb/v2/reference/internals/metrics/
	MetricsDisabled *bool

	// Disable the task scheduler. If problematic tasks prevent InfluxDB from
	// starting, use this option to start InfluxDB without scheduling or executing
	// tasks.
	//
	// Default: false
	NoTasks *bool

	// Number of queries allowed to execute concurrently. Setting to 0 allows an
	// unlimited number of concurrent queries.
	//
	// Default: 0
	QueryConcurrency *int32

	// Maximum number of queries allowed in execution queue. When queue limit is
	// reached, new queries are rejected. Setting to 0 allows an unlimited number of
	// queries in the queue.
	//
	// Default: 0
	QueryQueueSize *int32

	// Enable tracing in InfluxDB and specifies the tracing type. Tracing is disabled
	// by default.
	TracingType TracingType

	noSmithyDocumentSerde
}

// Configuration for sending InfluxDB engine logs to send to specified S3 bucket.
type LogDeliveryConfiguration struct {

	// Configuration for S3 bucket log delivery.
	//
	// This member is required.
	S3Configuration *S3Configuration

	noSmithyDocumentSerde
}

// The parameters that comprise the parameter group.
//
// The following types satisfy this interface:
//
//	ParametersMemberInfluxDBv2
type Parameters interface {
	isParameters()
}

// All the customer-modifiable InfluxDB v2 parameters in Timestream for InfluxDB.
type ParametersMemberInfluxDBv2 struct {
	Value InfluxDBv2Parameters

	noSmithyDocumentSerde
}

func (*ParametersMemberInfluxDBv2) isParameters() {}

// Configuration for S3 bucket log delivery.
type S3Configuration struct {

	// The name of the S3 bucket to deliver logs to.
	//
	// This member is required.
	BucketName *string

	// Indicates whether log delivery to the S3 bucket is enabled.
	//
	// This member is required.
	Enabled *bool

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

func (*UnknownUnionMember) isParameters() {}
