// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearchdomain

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/tracing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpSearch struct {
}

func (*awsRestjson1_serializeOpSearch) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpSearch) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*SearchInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/2013-01-01/search?format=sdk&pretty=true")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsSearchInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsSearchInput(v *SearchInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Cursor != nil {
		encoder.SetQuery("cursor").String(*v.Cursor)
	}

	if v.Expr != nil {
		encoder.SetQuery("expr").String(*v.Expr)
	}

	if v.Facet != nil {
		encoder.SetQuery("facet").String(*v.Facet)
	}

	if v.FilterQuery != nil {
		encoder.SetQuery("fq").String(*v.FilterQuery)
	}

	if v.Highlight != nil {
		encoder.SetQuery("highlight").String(*v.Highlight)
	}

	if v.Partial {
		encoder.SetQuery("partial").Boolean(v.Partial)
	}

	if v.Query != nil {
		encoder.SetQuery("q").String(*v.Query)
	}

	if v.QueryOptions != nil {
		encoder.SetQuery("q.options").String(*v.QueryOptions)
	}

	if len(v.QueryParser) > 0 {
		encoder.SetQuery("q.parser").String(string(v.QueryParser))
	}

	if v.Return != nil {
		encoder.SetQuery("return").String(*v.Return)
	}

	if v.Size != 0 {
		encoder.SetQuery("size").Long(v.Size)
	}

	if v.Sort != nil {
		encoder.SetQuery("sort").String(*v.Sort)
	}

	if v.Start != 0 {
		encoder.SetQuery("start").Long(v.Start)
	}

	if v.Stats != nil {
		encoder.SetQuery("stats").String(*v.Stats)
	}

	return nil
}

type awsRestjson1_serializeOpSuggest struct {
}

func (*awsRestjson1_serializeOpSuggest) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpSuggest) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*SuggestInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/2013-01-01/suggest?format=sdk&pretty=true")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsSuggestInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsSuggestInput(v *SuggestInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Query != nil {
		encoder.SetQuery("q").String(*v.Query)
	}

	if v.Size != 0 {
		encoder.SetQuery("size").Long(v.Size)
	}

	if v.Suggester != nil {
		encoder.SetQuery("suggester").String(*v.Suggester)
	}

	return nil
}

type awsRestjson1_serializeOpUploadDocuments struct {
}

func (*awsRestjson1_serializeOpUploadDocuments) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUploadDocuments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UploadDocumentsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/2013-01-01/documents/batch?format=sdk")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsUploadDocumentsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if !restEncoder.HasHeader("Content-Type") {
		ctx = smithyhttp.SetIsContentTypeDefaultValue(ctx, true)
		restEncoder.SetHeader("Content-Type").String("application/octet-stream")
	}

	if input.Documents != nil {
		payload := input.Documents
		if request, err = request.SetStream(payload); err != nil {
			return out, metadata, &smithy.SerializationError{Err: err}
		}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsUploadDocumentsInput(v *UploadDocumentsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if len(v.ContentType) > 0 {
		locationName := "Content-Type"
		encoder.SetHeader(locationName).String(string(v.ContentType))
	}

	return nil
}
