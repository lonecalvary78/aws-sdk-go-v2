// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

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
func TestOpCancelBatchJobExecutionSRAOperationOrder(t *testing.T) {
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
	_, err := svc.CancelBatchJobExecution(context.Background(), nil)
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
func TestOpCreateApplicationSRAOperationOrder(t *testing.T) {
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
	_, err := svc.CreateApplication(context.Background(), nil)
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
func TestOpCreateDataSetExportTaskSRAOperationOrder(t *testing.T) {
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
	_, err := svc.CreateDataSetExportTask(context.Background(), nil)
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
func TestOpCreateDataSetImportTaskSRAOperationOrder(t *testing.T) {
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
	_, err := svc.CreateDataSetImportTask(context.Background(), nil)
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
func TestOpCreateDeploymentSRAOperationOrder(t *testing.T) {
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
	_, err := svc.CreateDeployment(context.Background(), nil)
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
func TestOpCreateEnvironmentSRAOperationOrder(t *testing.T) {
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
	_, err := svc.CreateEnvironment(context.Background(), nil)
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
func TestOpDeleteApplicationSRAOperationOrder(t *testing.T) {
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
	_, err := svc.DeleteApplication(context.Background(), nil)
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
func TestOpDeleteApplicationFromEnvironmentSRAOperationOrder(t *testing.T) {
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
	_, err := svc.DeleteApplicationFromEnvironment(context.Background(), nil)
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
func TestOpDeleteEnvironmentSRAOperationOrder(t *testing.T) {
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
	_, err := svc.DeleteEnvironment(context.Background(), nil)
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
func TestOpGetApplicationSRAOperationOrder(t *testing.T) {
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
	_, err := svc.GetApplication(context.Background(), nil)
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
func TestOpGetApplicationVersionSRAOperationOrder(t *testing.T) {
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
	_, err := svc.GetApplicationVersion(context.Background(), nil)
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
func TestOpGetBatchJobExecutionSRAOperationOrder(t *testing.T) {
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
	_, err := svc.GetBatchJobExecution(context.Background(), nil)
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
func TestOpGetDataSetDetailsSRAOperationOrder(t *testing.T) {
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
	_, err := svc.GetDataSetDetails(context.Background(), nil)
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
func TestOpGetDataSetExportTaskSRAOperationOrder(t *testing.T) {
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
	_, err := svc.GetDataSetExportTask(context.Background(), nil)
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
func TestOpGetDataSetImportTaskSRAOperationOrder(t *testing.T) {
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
	_, err := svc.GetDataSetImportTask(context.Background(), nil)
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
func TestOpGetDeploymentSRAOperationOrder(t *testing.T) {
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
	_, err := svc.GetDeployment(context.Background(), nil)
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
func TestOpGetEnvironmentSRAOperationOrder(t *testing.T) {
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
	_, err := svc.GetEnvironment(context.Background(), nil)
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
func TestOpGetSignedBluinsightsUrlSRAOperationOrder(t *testing.T) {
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
	_, err := svc.GetSignedBluinsightsUrl(context.Background(), nil)
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
func TestOpListApplicationsSRAOperationOrder(t *testing.T) {
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
	_, err := svc.ListApplications(context.Background(), nil)
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
func TestOpListApplicationVersionsSRAOperationOrder(t *testing.T) {
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
	_, err := svc.ListApplicationVersions(context.Background(), nil)
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
func TestOpListBatchJobDefinitionsSRAOperationOrder(t *testing.T) {
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
	_, err := svc.ListBatchJobDefinitions(context.Background(), nil)
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
func TestOpListBatchJobExecutionsSRAOperationOrder(t *testing.T) {
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
	_, err := svc.ListBatchJobExecutions(context.Background(), nil)
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
func TestOpListBatchJobRestartPointsSRAOperationOrder(t *testing.T) {
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
	_, err := svc.ListBatchJobRestartPoints(context.Background(), nil)
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
func TestOpListDataSetExportHistorySRAOperationOrder(t *testing.T) {
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
	_, err := svc.ListDataSetExportHistory(context.Background(), nil)
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
func TestOpListDataSetImportHistorySRAOperationOrder(t *testing.T) {
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
	_, err := svc.ListDataSetImportHistory(context.Background(), nil)
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
func TestOpListDataSetsSRAOperationOrder(t *testing.T) {
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
	_, err := svc.ListDataSets(context.Background(), nil)
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
func TestOpListDeploymentsSRAOperationOrder(t *testing.T) {
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
	_, err := svc.ListDeployments(context.Background(), nil)
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
func TestOpListEngineVersionsSRAOperationOrder(t *testing.T) {
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
	_, err := svc.ListEngineVersions(context.Background(), nil)
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
func TestOpListEnvironmentsSRAOperationOrder(t *testing.T) {
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
	_, err := svc.ListEnvironments(context.Background(), nil)
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
func TestOpListTagsForResourceSRAOperationOrder(t *testing.T) {
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
	_, err := svc.ListTagsForResource(context.Background(), nil)
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
func TestOpStartApplicationSRAOperationOrder(t *testing.T) {
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
	_, err := svc.StartApplication(context.Background(), nil)
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
func TestOpStartBatchJobSRAOperationOrder(t *testing.T) {
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
	_, err := svc.StartBatchJob(context.Background(), nil)
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
func TestOpStopApplicationSRAOperationOrder(t *testing.T) {
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
	_, err := svc.StopApplication(context.Background(), nil)
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
func TestOpTagResourceSRAOperationOrder(t *testing.T) {
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
	_, err := svc.TagResource(context.Background(), nil)
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
func TestOpUntagResourceSRAOperationOrder(t *testing.T) {
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
	_, err := svc.UntagResource(context.Background(), nil)
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
func TestOpUpdateApplicationSRAOperationOrder(t *testing.T) {
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
	_, err := svc.UpdateApplication(context.Background(), nil)
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
func TestOpUpdateEnvironmentSRAOperationOrder(t *testing.T) {
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
	_, err := svc.UpdateEnvironment(context.Background(), nil)
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
