// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type PostContactSummaryFailureCode string

// Enum values for PostContactSummaryFailureCode
const (
	PostContactSummaryFailureCodeQuotaExceeded                   PostContactSummaryFailureCode = "QUOTA_EXCEEDED"
	PostContactSummaryFailureCodeInsufficientConversationContent PostContactSummaryFailureCode = "INSUFFICIENT_CONVERSATION_CONTENT"
	PostContactSummaryFailureCodeFailedSafetyGuidelines          PostContactSummaryFailureCode = "FAILED_SAFETY_GUIDELINES"
	PostContactSummaryFailureCodeInvalidAnalysisConfiguration    PostContactSummaryFailureCode = "INVALID_ANALYSIS_CONFIGURATION"
	PostContactSummaryFailureCodeInternalError                   PostContactSummaryFailureCode = "INTERNAL_ERROR"
)

// Values returns all known values for PostContactSummaryFailureCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PostContactSummaryFailureCode) Values() []PostContactSummaryFailureCode {
	return []PostContactSummaryFailureCode{
		"QUOTA_EXCEEDED",
		"INSUFFICIENT_CONVERSATION_CONTENT",
		"FAILED_SAFETY_GUIDELINES",
		"INVALID_ANALYSIS_CONFIGURATION",
		"INTERNAL_ERROR",
	}
}

type PostContactSummaryStatus string

// Enum values for PostContactSummaryStatus
const (
	PostContactSummaryStatusFailed    PostContactSummaryStatus = "FAILED"
	PostContactSummaryStatusCompleted PostContactSummaryStatus = "COMPLETED"
)

// Values returns all known values for PostContactSummaryStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PostContactSummaryStatus) Values() []PostContactSummaryStatus {
	return []PostContactSummaryStatus{
		"FAILED",
		"COMPLETED",
	}
}

type SentimentValue string

// Enum values for SentimentValue
const (
	SentimentValuePositive SentimentValue = "POSITIVE"
	SentimentValueNeutral  SentimentValue = "NEUTRAL"
	SentimentValueNegative SentimentValue = "NEGATIVE"
)

// Values returns all known values for SentimentValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SentimentValue) Values() []SentimentValue {
	return []SentimentValue{
		"POSITIVE",
		"NEUTRAL",
		"NEGATIVE",
	}
}
