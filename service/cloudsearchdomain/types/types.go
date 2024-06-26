// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// A container for facet information.
type Bucket struct {

	// The number of hits that contain the facet value in the specified facet field.
	Count int64

	// The facet value being counted.
	Value *string

	noSmithyDocumentSerde
}

// A container for the calculated facet values and counts.
type BucketInfo struct {

	// A list of the calculated facet values and counts.
	Buckets []Bucket

	noSmithyDocumentSerde
}

// A warning returned by the document service when an issue is discovered while
// processing an upload request.
type DocumentServiceWarning struct {

	// The description for a warning returned by the document service.
	Message *string

	noSmithyDocumentSerde
}

// The statistics for a field calculated in the request.
type FieldStats struct {

	// The number of documents that contain a value in the specified field in the
	// result set.
	Count int64

	// The maximum value found in the specified field in the result set.
	//
	// If the field is numeric ( int , int-array , double , or double-array ), max is
	// the string representation of a double-precision 64-bit floating point value. If
	// the field is date or date-array , max is the string representation of a date
	// with the format specified in [IETF RFC3339]: yyyy-mm-ddTHH:mm:ss.SSSZ.
	//
	// [IETF RFC3339]: http://tools.ietf.org/html/rfc3339
	Max *string

	// The average of the values found in the specified field in the result set.
	//
	// If the field is numeric ( int , int-array , double , or double-array ), mean is
	// the string representation of a double-precision 64-bit floating point value. If
	// the field is date or date-array , mean is the string representation of a date
	// with the format specified in [IETF RFC3339]: yyyy-mm-ddTHH:mm:ss.SSSZ.
	//
	// [IETF RFC3339]: http://tools.ietf.org/html/rfc3339
	Mean *string

	// The minimum value found in the specified field in the result set.
	//
	// If the field is numeric ( int , int-array , double , or double-array ), min is
	// the string representation of a double-precision 64-bit floating point value. If
	// the field is date or date-array , min is the string representation of a date
	// with the format specified in [IETF RFC3339]: yyyy-mm-ddTHH:mm:ss.SSSZ.
	//
	// [IETF RFC3339]: http://tools.ietf.org/html/rfc3339
	Min *string

	// The number of documents that do not contain a value in the specified field in
	// the result set.
	Missing int64

	// The standard deviation of the values in the specified field in the result set.
	Stddev float64

	// The sum of the field values across the documents in the result set. null for
	// date fields.
	Sum float64

	// The sum of all field values in the result set squared.
	SumOfSquares float64

	noSmithyDocumentSerde
}

// Information about a document that matches the search request.
type Hit struct {

	// The expressions returned from a document that matches the search request.
	Exprs map[string]string

	// The fields returned from a document that matches the search request.
	Fields map[string][]string

	// The highlights returned from a document that matches the search request.
	Highlights map[string]string

	// The document ID of a document that matches the search request.
	Id *string

	noSmithyDocumentSerde
}

// The collection of documents that match the search request.
type Hits struct {

	// A cursor that can be used to retrieve the next set of matching documents when
	// you want to page through a large result set.
	Cursor *string

	// The total number of documents that match the search request.
	Found int64

	// A document that matches the search request.
	Hit []Hit

	// The index of the first matching document.
	Start int64

	noSmithyDocumentSerde
}

// Contains the resource id ( rid ) and the time it took to process the request (
// timems ).
type SearchStatus struct {

	// The encrypted resource ID for the request.
	Rid *string

	// How long it took to process the request, in milliseconds.
	Timems int64

	noSmithyDocumentSerde
}

// An autocomplete suggestion that matches the query string specified in a
// SuggestRequest .
type SuggestionMatch struct {

	// The document ID of the suggested document.
	Id *string

	// The relevance score of a suggested match.
	Score int64

	// The string that matches the query string specified in the SuggestRequest .
	Suggestion *string

	noSmithyDocumentSerde
}

// Container for the suggestion information returned in a SuggestResponse .
type SuggestModel struct {

	// The number of documents that were found to match the query string.
	Found int64

	// The query string specified in the suggest request.
	Query *string

	// The documents that match the query string.
	Suggestions []SuggestionMatch

	noSmithyDocumentSerde
}

// Contains the resource id ( rid ) and the time it took to process the request (
// timems ).
type SuggestStatus struct {

	// The encrypted resource ID for the request.
	Rid *string

	// How long it took to process the request, in milliseconds.
	Timems int64

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
