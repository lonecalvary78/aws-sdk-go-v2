// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// You don't have sufficient access to perform this action.
type AccessDeniedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AccessDeniedException"
	}
	return *e.ErrorCodeOverride
}
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There is an error in the call or in a SQL statement. (This error only appears
// in calls from Aurora Serverless v1 databases.)
type BadRequestException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "BadRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There was an error in processing the SQL statement.
type DatabaseErrorException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DatabaseErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DatabaseErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DatabaseErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DatabaseErrorException"
	}
	return *e.ErrorCodeOverride
}
func (e *DatabaseErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The DB cluster doesn't have a DB instance.
type DatabaseNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DatabaseNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DatabaseNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DatabaseNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DatabaseNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *DatabaseNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A request was cancelled because the Aurora Serverless v2 DB instance was
// paused. The Data API request automatically resumes the DB instance. Wait a few
// seconds and try again.
type DatabaseResumingException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DatabaseResumingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DatabaseResumingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DatabaseResumingException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DatabaseResumingException"
	}
	return *e.ErrorCodeOverride
}
func (e *DatabaseResumingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The writer instance in the DB cluster isn't available.
type DatabaseUnavailableException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DatabaseUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DatabaseUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DatabaseUnavailableException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DatabaseUnavailableException"
	}
	return *e.ErrorCodeOverride
}
func (e *DatabaseUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// There are insufficient privileges to make the call.
type ForbiddenException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ForbiddenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ForbiddenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ForbiddenException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ForbiddenException"
	}
	return *e.ErrorCodeOverride
}
func (e *ForbiddenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The HTTP endpoint for using RDS Data API isn't enabled for the DB cluster.
type HttpEndpointNotEnabledException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *HttpEndpointNotEnabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *HttpEndpointNotEnabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *HttpEndpointNotEnabledException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "HttpEndpointNotEnabledException"
	}
	return *e.ErrorCodeOverride
}
func (e *HttpEndpointNotEnabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An internal error occurred.
type InternalServerErrorException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServerErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServerErrorException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServerErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The resource is in an invalid state.
type InvalidResourceStateException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidResourceStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidResourceStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidResourceStateException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidResourceStateException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidResourceStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Secrets Manager secret used with the request isn't valid.
type InvalidSecretException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidSecretException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSecretException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSecretException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSecretException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSecretException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resourceArn , secretArn , or transactionId value can't be found.
type NotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There was a problem with the Secrets Manager secret used with the request,
// caused by one of the following conditions:
//
//   - RDS Data API timed out retrieving the secret.
//
//   - The secret provided wasn't found.
//
//   - The secret couldn't be decrypted.
type SecretsErrorException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SecretsErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SecretsErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SecretsErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SecretsErrorException"
	}
	return *e.ErrorCodeOverride
}
func (e *SecretsErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service specified by the resourceArn parameter isn't available.
type ServiceUnavailableError struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServiceUnavailableError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableError) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceUnavailableError"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceUnavailableError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The execution of the SQL statement timed out.
type StatementTimeoutException struct {
	Message *string

	ErrorCodeOverride *string

	DbConnectionId int64

	noSmithyDocumentSerde
}

func (e *StatementTimeoutException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StatementTimeoutException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StatementTimeoutException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "StatementTimeoutException"
	}
	return *e.ErrorCodeOverride
}
func (e *StatementTimeoutException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The transaction ID wasn't found.
type TransactionNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TransactionNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TransactionNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TransactionNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TransactionNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *TransactionNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There was a problem with the result because of one of the following conditions:
//
//   - It contained an unsupported data type.
//
//   - It contained a multidimensional array.
//
//   - The size was too large.
type UnsupportedResultException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UnsupportedResultException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedResultException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedResultException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnsupportedResultException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnsupportedResultException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
