// Code generated by smithy-go-codegen DO NOT EDIT.

package document

import (
	internaldocument "github.com/aws/aws-sdk-go-v2/service/bedrockagentcore/internal/document"
)

// Interface defines a document which is a protocol-agnostic type which supports a
// JSON-like data-model. You can use this type to send UTF-8 strings, arbitrary
// precision numbers, booleans, nulls, a list of these values, and a map of UTF-8
// strings to these values.
//
// You create a document type using the NewLazyDocument function and passing it
// the Go type to marshal. When receiving a document in an API response, you use
// the document's UnmarshalSmithyDocument function to decode the response to your
// desired Go type. Unless documented specifically generated structure types in
// client packages or client types packages are not supported at this time. Such
// types embed a noSmithyDocumentSerde and will cause an error to be returned when
// attempting to send an API request.
//
// For more information see the accompanying package documentation and linked
// references.
type Interface = internaldocument.Interface

// You create document type using the NewLazyDocument function and passing it the
// Go type to be marshaled and sent to the service. The document marshaler supports
// semantics similar to the encoding/json Go standard library.
//
// For more information see the accompanying package documentation and linked
// references.
func NewLazyDocument(v interface{}) Interface {
	return internaldocument.NewDocumentMarshaler(v)
}
