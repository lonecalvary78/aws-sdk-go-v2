// Code generated by smithy-go-codegen DO NOT EDIT.

package smithyrpcv2cbor

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/smithyrpcv2cbor/types"
	smithy "github.com/aws/smithy-go"
	smithycbor "github.com/aws/smithy-go/encoding/cbor"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/tracing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"net/http"
	"time"
)

type smithyRpcv2cbor_serializeOpEmptyInputOutput struct {
}

func (*smithyRpcv2cbor_serializeOpEmptyInputOutput) ID() string {
	return "OperationSerializer"
}

func (m *smithyRpcv2cbor_serializeOpEmptyInputOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	input, ok := in.Parameters.(*EmptyInputOutputInput)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
	}
	_ = input

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", in.Request)
	}

	req.Method = http.MethodPost
	req.URL.Path = "/service/RpcV2Protocol/operation/EmptyInputOutput"
	req.Header.Set("smithy-protocol", "rpc-v2-cbor")

	req.Header.Set("Content-Type", "application/cbor")
	req.Header.Set("Accept", "application/cbor")

	cv, err := serializeCBOR_EmptyInputOutputInput(input)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	payload := bytes.NewReader(smithycbor.Encode(cv))
	if req, err = req.SetStream(payload); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	in.Request = req

	endTimer()
	span.End()

	return next.HandleSerialize(ctx, in)
}

type smithyRpcv2cbor_serializeOpFloat16 struct {
}

func (*smithyRpcv2cbor_serializeOpFloat16) ID() string {
	return "OperationSerializer"
}

func (m *smithyRpcv2cbor_serializeOpFloat16) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	input, ok := in.Parameters.(*Float16Input)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
	}
	_ = input

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", in.Request)
	}

	req.Method = http.MethodPost
	req.URL.Path = "/service/RpcV2Protocol/operation/Float16"
	req.Header.Set("smithy-protocol", "rpc-v2-cbor")

	req.Header.Set("Accept", "application/cbor")

	endTimer()
	span.End()

	return next.HandleSerialize(ctx, in)
}

type smithyRpcv2cbor_serializeOpFractionalSeconds struct {
}

func (*smithyRpcv2cbor_serializeOpFractionalSeconds) ID() string {
	return "OperationSerializer"
}

func (m *smithyRpcv2cbor_serializeOpFractionalSeconds) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	input, ok := in.Parameters.(*FractionalSecondsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
	}
	_ = input

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", in.Request)
	}

	req.Method = http.MethodPost
	req.URL.Path = "/service/RpcV2Protocol/operation/FractionalSeconds"
	req.Header.Set("smithy-protocol", "rpc-v2-cbor")

	req.Header.Set("Accept", "application/cbor")

	endTimer()
	span.End()

	return next.HandleSerialize(ctx, in)
}

type smithyRpcv2cbor_serializeOpGreetingWithErrors struct {
}

func (*smithyRpcv2cbor_serializeOpGreetingWithErrors) ID() string {
	return "OperationSerializer"
}

func (m *smithyRpcv2cbor_serializeOpGreetingWithErrors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	input, ok := in.Parameters.(*GreetingWithErrorsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
	}
	_ = input

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", in.Request)
	}

	req.Method = http.MethodPost
	req.URL.Path = "/service/RpcV2Protocol/operation/GreetingWithErrors"
	req.Header.Set("smithy-protocol", "rpc-v2-cbor")

	req.Header.Set("Accept", "application/cbor")

	endTimer()
	span.End()

	return next.HandleSerialize(ctx, in)
}

type smithyRpcv2cbor_serializeOpNoInputOutput struct {
}

func (*smithyRpcv2cbor_serializeOpNoInputOutput) ID() string {
	return "OperationSerializer"
}

func (m *smithyRpcv2cbor_serializeOpNoInputOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	input, ok := in.Parameters.(*NoInputOutputInput)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
	}
	_ = input

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", in.Request)
	}

	req.Method = http.MethodPost
	req.URL.Path = "/service/RpcV2Protocol/operation/NoInputOutput"
	req.Header.Set("smithy-protocol", "rpc-v2-cbor")

	req.Header.Set("Accept", "application/cbor")

	endTimer()
	span.End()

	return next.HandleSerialize(ctx, in)
}

type smithyRpcv2cbor_serializeOpOperationWithDefaults struct {
}

func (*smithyRpcv2cbor_serializeOpOperationWithDefaults) ID() string {
	return "OperationSerializer"
}

func (m *smithyRpcv2cbor_serializeOpOperationWithDefaults) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	input, ok := in.Parameters.(*OperationWithDefaultsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
	}
	_ = input

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", in.Request)
	}

	req.Method = http.MethodPost
	req.URL.Path = "/service/RpcV2Protocol/operation/OperationWithDefaults"
	req.Header.Set("smithy-protocol", "rpc-v2-cbor")

	req.Header.Set("Content-Type", "application/cbor")
	req.Header.Set("Accept", "application/cbor")

	cv, err := serializeCBOR_OperationWithDefaultsInput(input)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	payload := bytes.NewReader(smithycbor.Encode(cv))
	if req, err = req.SetStream(payload); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	in.Request = req

	endTimer()
	span.End()

	return next.HandleSerialize(ctx, in)
}

type smithyRpcv2cbor_serializeOpOptionalInputOutput struct {
}

func (*smithyRpcv2cbor_serializeOpOptionalInputOutput) ID() string {
	return "OperationSerializer"
}

func (m *smithyRpcv2cbor_serializeOpOptionalInputOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	input, ok := in.Parameters.(*OptionalInputOutputInput)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
	}
	_ = input

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", in.Request)
	}

	req.Method = http.MethodPost
	req.URL.Path = "/service/RpcV2Protocol/operation/OptionalInputOutput"
	req.Header.Set("smithy-protocol", "rpc-v2-cbor")

	req.Header.Set("Content-Type", "application/cbor")
	req.Header.Set("Accept", "application/cbor")

	cv, err := serializeCBOR_OptionalInputOutputInput(input)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	payload := bytes.NewReader(smithycbor.Encode(cv))
	if req, err = req.SetStream(payload); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	in.Request = req

	endTimer()
	span.End()

	return next.HandleSerialize(ctx, in)
}

type smithyRpcv2cbor_serializeOpRecursiveShapes struct {
}

func (*smithyRpcv2cbor_serializeOpRecursiveShapes) ID() string {
	return "OperationSerializer"
}

func (m *smithyRpcv2cbor_serializeOpRecursiveShapes) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	input, ok := in.Parameters.(*RecursiveShapesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
	}
	_ = input

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", in.Request)
	}

	req.Method = http.MethodPost
	req.URL.Path = "/service/RpcV2Protocol/operation/RecursiveShapes"
	req.Header.Set("smithy-protocol", "rpc-v2-cbor")

	req.Header.Set("Content-Type", "application/cbor")
	req.Header.Set("Accept", "application/cbor")

	cv, err := serializeCBOR_RecursiveShapesInput(input)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	payload := bytes.NewReader(smithycbor.Encode(cv))
	if req, err = req.SetStream(payload); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	in.Request = req

	endTimer()
	span.End()

	return next.HandleSerialize(ctx, in)
}

type smithyRpcv2cbor_serializeOpRpcV2CborDenseMaps struct {
}

func (*smithyRpcv2cbor_serializeOpRpcV2CborDenseMaps) ID() string {
	return "OperationSerializer"
}

func (m *smithyRpcv2cbor_serializeOpRpcV2CborDenseMaps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	input, ok := in.Parameters.(*RpcV2CborDenseMapsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
	}
	_ = input

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", in.Request)
	}

	req.Method = http.MethodPost
	req.URL.Path = "/service/RpcV2Protocol/operation/RpcV2CborDenseMaps"
	req.Header.Set("smithy-protocol", "rpc-v2-cbor")

	req.Header.Set("Content-Type", "application/cbor")
	req.Header.Set("Accept", "application/cbor")

	cv, err := serializeCBOR_RpcV2CborDenseMapsInput(input)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	payload := bytes.NewReader(smithycbor.Encode(cv))
	if req, err = req.SetStream(payload); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	in.Request = req

	endTimer()
	span.End()

	return next.HandleSerialize(ctx, in)
}

type smithyRpcv2cbor_serializeOpRpcV2CborLists struct {
}

func (*smithyRpcv2cbor_serializeOpRpcV2CborLists) ID() string {
	return "OperationSerializer"
}

func (m *smithyRpcv2cbor_serializeOpRpcV2CborLists) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	input, ok := in.Parameters.(*RpcV2CborListsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
	}
	_ = input

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", in.Request)
	}

	req.Method = http.MethodPost
	req.URL.Path = "/service/RpcV2Protocol/operation/RpcV2CborLists"
	req.Header.Set("smithy-protocol", "rpc-v2-cbor")

	req.Header.Set("Content-Type", "application/cbor")
	req.Header.Set("Accept", "application/cbor")

	cv, err := serializeCBOR_RpcV2CborListsInput(input)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	payload := bytes.NewReader(smithycbor.Encode(cv))
	if req, err = req.SetStream(payload); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	in.Request = req

	endTimer()
	span.End()

	return next.HandleSerialize(ctx, in)
}

type smithyRpcv2cbor_serializeOpRpcV2CborSparseMaps struct {
}

func (*smithyRpcv2cbor_serializeOpRpcV2CborSparseMaps) ID() string {
	return "OperationSerializer"
}

func (m *smithyRpcv2cbor_serializeOpRpcV2CborSparseMaps) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	input, ok := in.Parameters.(*RpcV2CborSparseMapsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
	}
	_ = input

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", in.Request)
	}

	req.Method = http.MethodPost
	req.URL.Path = "/service/RpcV2Protocol/operation/RpcV2CborSparseMaps"
	req.Header.Set("smithy-protocol", "rpc-v2-cbor")

	req.Header.Set("Content-Type", "application/cbor")
	req.Header.Set("Accept", "application/cbor")

	cv, err := serializeCBOR_RpcV2CborSparseMapsInput(input)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	payload := bytes.NewReader(smithycbor.Encode(cv))
	if req, err = req.SetStream(payload); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	in.Request = req

	endTimer()
	span.End()

	return next.HandleSerialize(ctx, in)
}

type smithyRpcv2cbor_serializeOpSimpleScalarProperties struct {
}

func (*smithyRpcv2cbor_serializeOpSimpleScalarProperties) ID() string {
	return "OperationSerializer"
}

func (m *smithyRpcv2cbor_serializeOpSimpleScalarProperties) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	input, ok := in.Parameters.(*SimpleScalarPropertiesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
	}
	_ = input

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", in.Request)
	}

	req.Method = http.MethodPost
	req.URL.Path = "/service/RpcV2Protocol/operation/SimpleScalarProperties"
	req.Header.Set("smithy-protocol", "rpc-v2-cbor")

	req.Header.Set("Content-Type", "application/cbor")
	req.Header.Set("Accept", "application/cbor")

	cv, err := serializeCBOR_SimpleScalarPropertiesInput(input)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	payload := bytes.NewReader(smithycbor.Encode(cv))
	if req, err = req.SetStream(payload); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	in.Request = req

	endTimer()
	span.End()

	return next.HandleSerialize(ctx, in)
}

type smithyRpcv2cbor_serializeOpSparseNullsOperation struct {
}

func (*smithyRpcv2cbor_serializeOpSparseNullsOperation) ID() string {
	return "OperationSerializer"
}

func (m *smithyRpcv2cbor_serializeOpSparseNullsOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	input, ok := in.Parameters.(*SparseNullsOperationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
	}
	_ = input

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", in.Request)
	}

	req.Method = http.MethodPost
	req.URL.Path = "/service/RpcV2Protocol/operation/SparseNullsOperation"
	req.Header.Set("smithy-protocol", "rpc-v2-cbor")

	req.Header.Set("Content-Type", "application/cbor")
	req.Header.Set("Accept", "application/cbor")

	cv, err := serializeCBOR_SparseNullsOperationInput(input)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	payload := bytes.NewReader(smithycbor.Encode(cv))
	if req, err = req.SetStream(payload); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	in.Request = req

	endTimer()
	span.End()

	return next.HandleSerialize(ctx, in)
}
func serializeCBOR_Blob(v []byte) (smithycbor.Value, error) {
	return smithycbor.Slice(v), nil
}

func serializeCBOR_Bool(v bool) (smithycbor.Value, error) {
	return smithycbor.Bool(v), nil
}

func serializeCBOR_Float32(v float32) (smithycbor.Value, error) {
	return smithycbor.Float32(v), nil
}

func serializeCBOR_Float64(v float64) (smithycbor.Value, error) {
	return smithycbor.Float64(v), nil
}

func serializeCBOR_Int16(v int16) (smithycbor.Value, error) {
	if v < 0 {
		return smithycbor.NegInt(uint64(-v)), nil
	}
	return smithycbor.Uint(uint64(v)), nil
}

func serializeCBOR_Int32(v int32) (smithycbor.Value, error) {
	if v < 0 {
		return smithycbor.NegInt(uint64(-v)), nil
	}
	return smithycbor.Uint(uint64(v)), nil
}

func serializeCBOR_Int64(v int64) (smithycbor.Value, error) {
	if v < 0 {
		return smithycbor.NegInt(uint64(-v)), nil
	}
	return smithycbor.Uint(uint64(v)), nil
}

func serializeCBOR_Int8(v int8) (smithycbor.Value, error) {
	if v < 0 {
		return smithycbor.NegInt(uint64(-v)), nil
	}
	return smithycbor.Uint(uint64(v)), nil
}

func serializeCBOR_String(v string) (smithycbor.Value, error) {
	return smithycbor.String(v), nil
}

func serializeCBOR_Time(v time.Time) (smithycbor.Value, error) {
	return &smithycbor.Tag{
		ID:    1,
		Value: smithycbor.Float64(float64(v.UnixMilli()) / 1000),
	}, nil
}

func serializeCBOR_EmptyInputOutputInput(v *EmptyInputOutputInput) (smithycbor.Value, error) {
	vm := smithycbor.Map{}

	return vm, nil
}

func serializeCBOR_Float16Input(v *Float16Input) (smithycbor.Value, error) {
	vm := smithycbor.Map{}

	return vm, nil
}

func serializeCBOR_FractionalSecondsInput(v *FractionalSecondsInput) (smithycbor.Value, error) {
	vm := smithycbor.Map{}

	return vm, nil
}

func serializeCBOR_GreetingWithErrorsInput(v *GreetingWithErrorsInput) (smithycbor.Value, error) {
	vm := smithycbor.Map{}

	return vm, nil
}

func serializeCBOR_NoInputOutputInput(v *NoInputOutputInput) (smithycbor.Value, error) {
	vm := smithycbor.Map{}

	return vm, nil
}

func serializeCBOR_OperationWithDefaultsInput(v *OperationWithDefaultsInput) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	if v.Defaults != nil {
		ser, err := serializeCBOR_Defaults(v.Defaults)
		if err != nil {
			return nil, err
		}
		vm["defaults"] = ser
	}
	if v.ClientOptionalDefaults != nil {
		ser, err := serializeCBOR_ClientOptionalDefaults(v.ClientOptionalDefaults)
		if err != nil {
			return nil, err
		}
		vm["clientOptionalDefaults"] = ser
	}
	if v.TopLevelDefault != nil {
		ser, err := serializeCBOR_String(*v.TopLevelDefault)
		if err != nil {
			return nil, err
		}
		vm["topLevelDefault"] = ser
	}
	serotherTopLevelDefault, err := serializeCBOR_Int32(v.OtherTopLevelDefault)
	if err != nil {
		return nil, err
	}
	vm["otherTopLevelDefault"] = serotherTopLevelDefault
	return vm, nil
}

func serializeCBOR_OptionalInputOutputInput(v *OptionalInputOutputInput) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	if v.Value != nil {
		ser, err := serializeCBOR_String(*v.Value)
		if err != nil {
			return nil, err
		}
		vm["value"] = ser
	}
	return vm, nil
}

func serializeCBOR_RecursiveShapesInput(v *RecursiveShapesInput) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	if v.Nested != nil {
		ser, err := serializeCBOR_RecursiveShapesInputOutputNested1(v.Nested)
		if err != nil {
			return nil, err
		}
		vm["nested"] = ser
	}
	return vm, nil
}

func serializeCBOR_RpcV2CborDenseMapsInput(v *RpcV2CborDenseMapsInput) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	if v.DenseStructMap != nil {
		ser, err := serializeCBOR_DenseStructMap(v.DenseStructMap)
		if err != nil {
			return nil, err
		}
		vm["denseStructMap"] = ser
	}
	if v.DenseNumberMap != nil {
		ser, err := serializeCBOR_DenseNumberMap(v.DenseNumberMap)
		if err != nil {
			return nil, err
		}
		vm["denseNumberMap"] = ser
	}
	if v.DenseBooleanMap != nil {
		ser, err := serializeCBOR_DenseBooleanMap(v.DenseBooleanMap)
		if err != nil {
			return nil, err
		}
		vm["denseBooleanMap"] = ser
	}
	if v.DenseStringMap != nil {
		ser, err := serializeCBOR_DenseStringMap(v.DenseStringMap)
		if err != nil {
			return nil, err
		}
		vm["denseStringMap"] = ser
	}
	if v.DenseSetMap != nil {
		ser, err := serializeCBOR_DenseSetMap(v.DenseSetMap)
		if err != nil {
			return nil, err
		}
		vm["denseSetMap"] = ser
	}
	return vm, nil
}

func serializeCBOR_RpcV2CborListsInput(v *RpcV2CborListsInput) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	if v.StringList != nil {
		ser, err := serializeCBOR_StringList(v.StringList)
		if err != nil {
			return nil, err
		}
		vm["stringList"] = ser
	}
	if v.StringSet != nil {
		ser, err := serializeCBOR_StringSet(v.StringSet)
		if err != nil {
			return nil, err
		}
		vm["stringSet"] = ser
	}
	if v.IntegerList != nil {
		ser, err := serializeCBOR_IntegerList(v.IntegerList)
		if err != nil {
			return nil, err
		}
		vm["integerList"] = ser
	}
	if v.BooleanList != nil {
		ser, err := serializeCBOR_BooleanList(v.BooleanList)
		if err != nil {
			return nil, err
		}
		vm["booleanList"] = ser
	}
	if v.TimestampList != nil {
		ser, err := serializeCBOR_TimestampList(v.TimestampList)
		if err != nil {
			return nil, err
		}
		vm["timestampList"] = ser
	}
	if v.EnumList != nil {
		ser, err := serializeCBOR_FooEnumList(v.EnumList)
		if err != nil {
			return nil, err
		}
		vm["enumList"] = ser
	}
	if v.IntEnumList != nil {
		ser, err := serializeCBOR_IntegerEnumList(v.IntEnumList)
		if err != nil {
			return nil, err
		}
		vm["intEnumList"] = ser
	}
	if v.NestedStringList != nil {
		ser, err := serializeCBOR_NestedStringList(v.NestedStringList)
		if err != nil {
			return nil, err
		}
		vm["nestedStringList"] = ser
	}
	if v.StructureList != nil {
		ser, err := serializeCBOR_StructureList(v.StructureList)
		if err != nil {
			return nil, err
		}
		vm["structureList"] = ser
	}
	if v.BlobList != nil {
		ser, err := serializeCBOR_BlobList(v.BlobList)
		if err != nil {
			return nil, err
		}
		vm["blobList"] = ser
	}
	return vm, nil
}

func serializeCBOR_RpcV2CborSparseMapsInput(v *RpcV2CborSparseMapsInput) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	if v.SparseStructMap != nil {
		ser, err := serializeCBOR_SparseStructMap(v.SparseStructMap)
		if err != nil {
			return nil, err
		}
		vm["sparseStructMap"] = ser
	}
	if v.SparseNumberMap != nil {
		ser, err := serializeCBOR_SparseNumberMap(v.SparseNumberMap)
		if err != nil {
			return nil, err
		}
		vm["sparseNumberMap"] = ser
	}
	if v.SparseBooleanMap != nil {
		ser, err := serializeCBOR_SparseBooleanMap(v.SparseBooleanMap)
		if err != nil {
			return nil, err
		}
		vm["sparseBooleanMap"] = ser
	}
	if v.SparseStringMap != nil {
		ser, err := serializeCBOR_SparseStringMap(v.SparseStringMap)
		if err != nil {
			return nil, err
		}
		vm["sparseStringMap"] = ser
	}
	if v.SparseSetMap != nil {
		ser, err := serializeCBOR_SparseSetMap(v.SparseSetMap)
		if err != nil {
			return nil, err
		}
		vm["sparseSetMap"] = ser
	}
	return vm, nil
}

func serializeCBOR_SimpleScalarPropertiesInput(v *SimpleScalarPropertiesInput) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	if v.TrueBooleanValue != nil {
		ser, err := serializeCBOR_Bool(*v.TrueBooleanValue)
		if err != nil {
			return nil, err
		}
		vm["trueBooleanValue"] = ser
	}
	if v.FalseBooleanValue != nil {
		ser, err := serializeCBOR_Bool(*v.FalseBooleanValue)
		if err != nil {
			return nil, err
		}
		vm["falseBooleanValue"] = ser
	}
	if v.ByteValue != nil {
		ser, err := serializeCBOR_Int8(*v.ByteValue)
		if err != nil {
			return nil, err
		}
		vm["byteValue"] = ser
	}
	if v.DoubleValue != nil {
		ser, err := serializeCBOR_Float64(*v.DoubleValue)
		if err != nil {
			return nil, err
		}
		vm["doubleValue"] = ser
	}
	if v.FloatValue != nil {
		ser, err := serializeCBOR_Float32(*v.FloatValue)
		if err != nil {
			return nil, err
		}
		vm["floatValue"] = ser
	}
	if v.IntegerValue != nil {
		ser, err := serializeCBOR_Int32(*v.IntegerValue)
		if err != nil {
			return nil, err
		}
		vm["integerValue"] = ser
	}
	if v.LongValue != nil {
		ser, err := serializeCBOR_Int64(*v.LongValue)
		if err != nil {
			return nil, err
		}
		vm["longValue"] = ser
	}
	if v.ShortValue != nil {
		ser, err := serializeCBOR_Int16(*v.ShortValue)
		if err != nil {
			return nil, err
		}
		vm["shortValue"] = ser
	}
	if v.StringValue != nil {
		ser, err := serializeCBOR_String(*v.StringValue)
		if err != nil {
			return nil, err
		}
		vm["stringValue"] = ser
	}
	if v.BlobValue != nil {
		ser, err := serializeCBOR_Blob(v.BlobValue)
		if err != nil {
			return nil, err
		}
		vm["blobValue"] = ser
	}
	return vm, nil
}

func serializeCBOR_SparseNullsOperationInput(v *SparseNullsOperationInput) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	if v.SparseStringList != nil {
		ser, err := serializeCBOR_SparseStringList(v.SparseStringList)
		if err != nil {
			return nil, err
		}
		vm["sparseStringList"] = ser
	}
	if v.SparseStringMap != nil {
		ser, err := serializeCBOR_SparseStringMap(v.SparseStringMap)
		if err != nil {
			return nil, err
		}
		vm["sparseStringMap"] = ser
	}
	return vm, nil
}

func serializeCBOR_ClientOptionalDefaults(v *types.ClientOptionalDefaults) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	if v.Member != nil {
		ser, err := serializeCBOR_Int32(*v.Member)
		if err != nil {
			return nil, err
		}
		vm["member"] = ser
	}
	return vm, nil
}

func serializeCBOR_Defaults(v *types.Defaults) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	if v.DefaultString != nil {
		ser, err := serializeCBOR_String(*v.DefaultString)
		if err != nil {
			return nil, err
		}
		vm["defaultString"] = ser
	}
	if v.DefaultBoolean != nil {
		ser, err := serializeCBOR_Bool(*v.DefaultBoolean)
		if err != nil {
			return nil, err
		}
		vm["defaultBoolean"] = ser
	}
	if v.DefaultList != nil {
		ser, err := serializeCBOR_TestStringList(v.DefaultList)
		if err != nil {
			return nil, err
		}
		vm["defaultList"] = ser
	}
	if v.DefaultTimestamp != nil {
		ser, err := serializeCBOR_Time(*v.DefaultTimestamp)
		if err != nil {
			return nil, err
		}
		vm["defaultTimestamp"] = ser
	}
	if v.DefaultBlob != nil {
		ser, err := serializeCBOR_Blob(v.DefaultBlob)
		if err != nil {
			return nil, err
		}
		vm["defaultBlob"] = ser
	}
	if v.DefaultByte != nil {
		ser, err := serializeCBOR_Int8(*v.DefaultByte)
		if err != nil {
			return nil, err
		}
		vm["defaultByte"] = ser
	}
	if v.DefaultShort != nil {
		ser, err := serializeCBOR_Int16(*v.DefaultShort)
		if err != nil {
			return nil, err
		}
		vm["defaultShort"] = ser
	}
	if v.DefaultInteger != nil {
		ser, err := serializeCBOR_Int32(*v.DefaultInteger)
		if err != nil {
			return nil, err
		}
		vm["defaultInteger"] = ser
	}
	if v.DefaultLong != nil {
		ser, err := serializeCBOR_Int64(*v.DefaultLong)
		if err != nil {
			return nil, err
		}
		vm["defaultLong"] = ser
	}
	if v.DefaultFloat != nil {
		ser, err := serializeCBOR_Float32(*v.DefaultFloat)
		if err != nil {
			return nil, err
		}
		vm["defaultFloat"] = ser
	}
	if v.DefaultDouble != nil {
		ser, err := serializeCBOR_Float64(*v.DefaultDouble)
		if err != nil {
			return nil, err
		}
		vm["defaultDouble"] = ser
	}
	if v.DefaultMap != nil {
		ser, err := serializeCBOR_TestStringMap(v.DefaultMap)
		if err != nil {
			return nil, err
		}
		vm["defaultMap"] = ser
	}
	if len(v.DefaultEnum) > 0 {
		ser, err := serializeCBOR_TestEnum(v.DefaultEnum)
		if err != nil {
			return nil, err
		}
		vm["defaultEnum"] = ser
	}
	serdefaultIntEnum, err := serializeCBOR_TestIntEnum(v.DefaultIntEnum)
	if err != nil {
		return nil, err
	}
	vm["defaultIntEnum"] = serdefaultIntEnum
	if v.EmptyString != nil {
		ser, err := serializeCBOR_String(*v.EmptyString)
		if err != nil {
			return nil, err
		}
		vm["emptyString"] = ser
	}
	serfalseBoolean, err := serializeCBOR_Bool(v.FalseBoolean)
	if err != nil {
		return nil, err
	}
	vm["falseBoolean"] = serfalseBoolean
	if v.EmptyBlob != nil {
		ser, err := serializeCBOR_Blob(v.EmptyBlob)
		if err != nil {
			return nil, err
		}
		vm["emptyBlob"] = ser
	}
	serzeroByte, err := serializeCBOR_Int8(v.ZeroByte)
	if err != nil {
		return nil, err
	}
	vm["zeroByte"] = serzeroByte
	serzeroShort, err := serializeCBOR_Int16(v.ZeroShort)
	if err != nil {
		return nil, err
	}
	vm["zeroShort"] = serzeroShort
	serzeroInteger, err := serializeCBOR_Int32(v.ZeroInteger)
	if err != nil {
		return nil, err
	}
	vm["zeroInteger"] = serzeroInteger
	serzeroLong, err := serializeCBOR_Int64(v.ZeroLong)
	if err != nil {
		return nil, err
	}
	vm["zeroLong"] = serzeroLong
	serzeroFloat, err := serializeCBOR_Float32(v.ZeroFloat)
	if err != nil {
		return nil, err
	}
	vm["zeroFloat"] = serzeroFloat
	serzeroDouble, err := serializeCBOR_Float64(v.ZeroDouble)
	if err != nil {
		return nil, err
	}
	vm["zeroDouble"] = serzeroDouble
	return vm, nil
}

func serializeCBOR_DenseBooleanMap(v map[string]bool) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	for k, vv := range v {

		ser, err := serializeCBOR_Bool(vv)
		if err != nil {
			return nil, err
		}
		vm[k] = ser
	}
	return vm, nil
}

func serializeCBOR_DenseNumberMap(v map[string]int32) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	for k, vv := range v {

		ser, err := serializeCBOR_Int32(vv)
		if err != nil {
			return nil, err
		}
		vm[k] = ser
	}
	return vm, nil
}

func serializeCBOR_DenseSetMap(v map[string][]string) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	for k, vv := range v {
		if vv == nil {
			vm[k] = &smithycbor.Nil{}
			continue
		}
		ser, err := serializeCBOR_StringSet(vv)
		if err != nil {
			return nil, err
		}
		vm[k] = ser
	}
	return vm, nil
}

func serializeCBOR_DenseStringMap(v map[string]string) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	for k, vv := range v {

		ser, err := serializeCBOR_String(vv)
		if err != nil {
			return nil, err
		}
		vm[k] = ser
	}
	return vm, nil
}

func serializeCBOR_DenseStructMap(v map[string]types.GreetingStruct) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	for k, vv := range v {

		ser, err := serializeCBOR_GreetingStruct(&vv)
		if err != nil {
			return nil, err
		}
		vm[k] = ser
	}
	return vm, nil
}

func serializeCBOR_RecursiveShapesInputOutputNested1(v *types.RecursiveShapesInputOutputNested1) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	if v.Foo != nil {
		ser, err := serializeCBOR_String(*v.Foo)
		if err != nil {
			return nil, err
		}
		vm["foo"] = ser
	}
	if v.Nested != nil {
		ser, err := serializeCBOR_RecursiveShapesInputOutputNested2(v.Nested)
		if err != nil {
			return nil, err
		}
		vm["nested"] = ser
	}
	return vm, nil
}

func serializeCBOR_RecursiveShapesInputOutputNested2(v *types.RecursiveShapesInputOutputNested2) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	if v.Bar != nil {
		ser, err := serializeCBOR_String(*v.Bar)
		if err != nil {
			return nil, err
		}
		vm["bar"] = ser
	}
	if v.RecursiveMember != nil {
		ser, err := serializeCBOR_RecursiveShapesInputOutputNested1(v.RecursiveMember)
		if err != nil {
			return nil, err
		}
		vm["recursiveMember"] = ser
	}
	return vm, nil
}

func serializeCBOR_SparseBooleanMap(v map[string]*bool) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	for k, vv := range v {
		if vv == nil {
			vm[k] = &smithycbor.Nil{}
			continue
		}
		ser, err := serializeCBOR_Bool(*vv)
		if err != nil {
			return nil, err
		}
		vm[k] = ser
	}
	return vm, nil
}

func serializeCBOR_SparseNumberMap(v map[string]*int32) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	for k, vv := range v {
		if vv == nil {
			vm[k] = &smithycbor.Nil{}
			continue
		}
		ser, err := serializeCBOR_Int32(*vv)
		if err != nil {
			return nil, err
		}
		vm[k] = ser
	}
	return vm, nil
}

func serializeCBOR_SparseSetMap(v map[string][]string) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	for k, vv := range v {
		if vv == nil {
			vm[k] = &smithycbor.Nil{}
			continue
		}
		ser, err := serializeCBOR_StringSet(vv)
		if err != nil {
			return nil, err
		}
		vm[k] = ser
	}
	return vm, nil
}

func serializeCBOR_SparseStructMap(v map[string]*types.GreetingStruct) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	for k, vv := range v {
		if vv == nil {
			vm[k] = &smithycbor.Nil{}
			continue
		}
		ser, err := serializeCBOR_GreetingStruct(vv)
		if err != nil {
			return nil, err
		}
		vm[k] = ser
	}
	return vm, nil
}

func serializeCBOR_StructureList(v []types.StructureListMember) (smithycbor.Value, error) {
	vl := smithycbor.List{}
	for i := range v {

		ser, err := serializeCBOR_StructureListMember(&v[i])
		if err != nil {
			return nil, err
		}
		vl = append(vl, ser)
	}
	return vl, nil
}

func serializeCBOR_StructureListMember(v *types.StructureListMember) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	if v.A != nil {
		ser, err := serializeCBOR_String(*v.A)
		if err != nil {
			return nil, err
		}
		vm["a"] = ser
	}
	if v.B != nil {
		ser, err := serializeCBOR_String(*v.B)
		if err != nil {
			return nil, err
		}
		vm["b"] = ser
	}
	return vm, nil
}

func serializeCBOR_TestEnum(v types.TestEnum) (smithycbor.Value, error) {
	return smithycbor.String(string(v)), nil
}

func serializeCBOR_TestIntEnum(v types.TestIntEnum) (smithycbor.Value, error) {
	if v < 0 {
		return smithycbor.NegInt(uint64(-v)), nil
	}
	return smithycbor.Uint(uint64(v)), nil
}

func serializeCBOR_TestStringList(v []string) (smithycbor.Value, error) {
	vl := smithycbor.List{}
	for i := range v {

		ser, err := serializeCBOR_String(v[i])
		if err != nil {
			return nil, err
		}
		vl = append(vl, ser)
	}
	return vl, nil
}

func serializeCBOR_TestStringMap(v map[string]string) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	for k, vv := range v {

		ser, err := serializeCBOR_String(vv)
		if err != nil {
			return nil, err
		}
		vm[k] = ser
	}
	return vm, nil
}

func serializeCBOR_BlobList(v [][]byte) (smithycbor.Value, error) {
	vl := smithycbor.List{}
	for i := range v {

		ser, err := serializeCBOR_Blob(v[i])
		if err != nil {
			return nil, err
		}
		vl = append(vl, ser)
	}
	return vl, nil
}

func serializeCBOR_BooleanList(v []bool) (smithycbor.Value, error) {
	vl := smithycbor.List{}
	for i := range v {

		ser, err := serializeCBOR_Bool(v[i])
		if err != nil {
			return nil, err
		}
		vl = append(vl, ser)
	}
	return vl, nil
}

func serializeCBOR_FooEnum(v types.FooEnum) (smithycbor.Value, error) {
	return smithycbor.String(string(v)), nil
}

func serializeCBOR_FooEnumList(v []types.FooEnum) (smithycbor.Value, error) {
	vl := smithycbor.List{}
	for i := range v {

		ser, err := serializeCBOR_FooEnum(v[i])
		if err != nil {
			return nil, err
		}
		vl = append(vl, ser)
	}
	return vl, nil
}

func serializeCBOR_GreetingStruct(v *types.GreetingStruct) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	if v.Hi != nil {
		ser, err := serializeCBOR_String(*v.Hi)
		if err != nil {
			return nil, err
		}
		vm["hi"] = ser
	}
	return vm, nil
}

func serializeCBOR_IntegerEnum(v types.IntegerEnum) (smithycbor.Value, error) {
	if v < 0 {
		return smithycbor.NegInt(uint64(-v)), nil
	}
	return smithycbor.Uint(uint64(v)), nil
}

func serializeCBOR_IntegerEnumList(v []types.IntegerEnum) (smithycbor.Value, error) {
	vl := smithycbor.List{}
	for i := range v {

		ser, err := serializeCBOR_IntegerEnum(v[i])
		if err != nil {
			return nil, err
		}
		vl = append(vl, ser)
	}
	return vl, nil
}

func serializeCBOR_IntegerList(v []int32) (smithycbor.Value, error) {
	vl := smithycbor.List{}
	for i := range v {

		ser, err := serializeCBOR_Int32(v[i])
		if err != nil {
			return nil, err
		}
		vl = append(vl, ser)
	}
	return vl, nil
}

func serializeCBOR_NestedStringList(v [][]string) (smithycbor.Value, error) {
	vl := smithycbor.List{}
	for i := range v {
		if v[i] == nil {
			vl = append(vl, &smithycbor.Nil{})
			continue
		}
		ser, err := serializeCBOR_StringList(v[i])
		if err != nil {
			return nil, err
		}
		vl = append(vl, ser)
	}
	return vl, nil
}

func serializeCBOR_SparseStringList(v []*string) (smithycbor.Value, error) {
	vl := smithycbor.List{}
	for i := range v {
		if v[i] == nil {
			vl = append(vl, &smithycbor.Nil{})
			continue
		}
		ser, err := serializeCBOR_String(*v[i])
		if err != nil {
			return nil, err
		}
		vl = append(vl, ser)
	}
	return vl, nil
}

func serializeCBOR_SparseStringMap(v map[string]*string) (smithycbor.Value, error) {
	vm := smithycbor.Map{}
	for k, vv := range v {
		if vv == nil {
			vm[k] = &smithycbor.Nil{}
			continue
		}
		ser, err := serializeCBOR_String(*vv)
		if err != nil {
			return nil, err
		}
		vm[k] = ser
	}
	return vm, nil
}

func serializeCBOR_StringList(v []string) (smithycbor.Value, error) {
	vl := smithycbor.List{}
	for i := range v {

		ser, err := serializeCBOR_String(v[i])
		if err != nil {
			return nil, err
		}
		vl = append(vl, ser)
	}
	return vl, nil
}

func serializeCBOR_StringSet(v []string) (smithycbor.Value, error) {
	vl := smithycbor.List{}
	for i := range v {

		ser, err := serializeCBOR_String(v[i])
		if err != nil {
			return nil, err
		}
		vl = append(vl, ser)
	}
	return vl, nil
}

func serializeCBOR_TimestampList(v []time.Time) (smithycbor.Value, error) {
	vl := smithycbor.List{}
	for i := range v {

		ser, err := serializeCBOR_Time(v[i])
		if err != nil {
			return nil, err
		}
		vl = append(vl, ser)
	}
	return vl, nil
}
