// Code generated by smithy-go-codegen DO NOT EDIT.

package s3vectors

import (
	"context"
	"errors"
	"github.com/aws/smithy-go/middleware"
	"slices"
	"strings"
	"testing"
)

var errTestReturnEarly = errors.New("errTestReturnEarly")

func captureMiddlewareStack(stack *middleware.Stack) func(*middleware.Stack) error {
	return func(inner *middleware.Stack) error {
		*stack = *inner
		return errTestReturnEarly
	}
}
func TestOpCreateIndexSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.CreateIndex(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpCreateVectorBucketSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.CreateVectorBucket(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpDeleteIndexSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.DeleteIndex(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpDeleteVectorBucketSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.DeleteVectorBucket(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpDeleteVectorBucketPolicySRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.DeleteVectorBucketPolicy(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpDeleteVectorsSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.DeleteVectors(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpGetIndexSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.GetIndex(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpGetVectorBucketSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.GetVectorBucket(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpGetVectorBucketPolicySRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.GetVectorBucketPolicy(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpGetVectorsSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.GetVectors(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpListIndexesSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.ListIndexes(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpListVectorBucketsSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.ListVectorBuckets(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpListVectorsSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.ListVectors(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpPutVectorBucketPolicySRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.PutVectorBucketPolicy(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpPutVectorsSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.PutVectors(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpQueryVectorsSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.QueryVectors(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
