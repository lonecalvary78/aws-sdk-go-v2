// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/synthetics/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAssociateResource struct {
}

func (*validateOpAssociateResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateCanary struct {
}

func (*validateOpCreateCanary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCanary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateCanaryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateCanaryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateGroup struct {
}

func (*validateOpCreateGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteCanary struct {
}

func (*validateOpDeleteCanary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteCanary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteCanaryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteCanaryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteGroup struct {
}

func (*validateOpDeleteGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateResource struct {
}

func (*validateOpDisassociateResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCanary struct {
}

func (*validateOpGetCanary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCanary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCanaryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCanaryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCanaryRuns struct {
}

func (*validateOpGetCanaryRuns) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCanaryRuns) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCanaryRunsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCanaryRunsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetGroup struct {
}

func (*validateOpGetGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListAssociatedGroups struct {
}

func (*validateOpListAssociatedGroups) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListAssociatedGroups) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListAssociatedGroupsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListAssociatedGroupsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListGroupResources struct {
}

func (*validateOpListGroupResources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListGroupResources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListGroupResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListGroupResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartCanaryDryRun struct {
}

func (*validateOpStartCanaryDryRun) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartCanaryDryRun) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartCanaryDryRunInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartCanaryDryRunInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartCanary struct {
}

func (*validateOpStartCanary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartCanary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartCanaryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartCanaryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopCanary struct {
}

func (*validateOpStopCanary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopCanary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopCanaryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopCanaryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateCanary struct {
}

func (*validateOpUpdateCanary) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateCanary) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateCanaryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateCanaryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAssociateResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateResource{}, middleware.After)
}

func addOpCreateCanaryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCanary{}, middleware.After)
}

func addOpCreateGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateGroup{}, middleware.After)
}

func addOpDeleteCanaryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteCanary{}, middleware.After)
}

func addOpDeleteGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteGroup{}, middleware.After)
}

func addOpDisassociateResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateResource{}, middleware.After)
}

func addOpGetCanaryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCanary{}, middleware.After)
}

func addOpGetCanaryRunsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCanaryRuns{}, middleware.After)
}

func addOpGetGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetGroup{}, middleware.After)
}

func addOpListAssociatedGroupsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListAssociatedGroups{}, middleware.After)
}

func addOpListGroupResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListGroupResources{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpStartCanaryDryRunValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartCanaryDryRun{}, middleware.After)
}

func addOpStartCanaryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartCanary{}, middleware.After)
}

func addOpStopCanaryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopCanary{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateCanaryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateCanary{}, middleware.After)
}

func validateBaseScreenshot(v *types.BaseScreenshot) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BaseScreenshot"}
	if v.ScreenshotName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScreenshotName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateBaseScreenshots(v []types.BaseScreenshot) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BaseScreenshots"}
	for i := range v {
		if err := validateBaseScreenshot(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCanaryCodeInput(v *types.CanaryCodeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CanaryCodeInput"}
	if v.Handler == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Handler"))
	}
	if v.Dependencies != nil {
		if err := validateDependencies(v.Dependencies); err != nil {
			invalidParams.AddNested("Dependencies", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCanaryScheduleInput(v *types.CanaryScheduleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CanaryScheduleInput"}
	if v.Expression == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Expression"))
	}
	if v.RetryConfig != nil {
		if err := validateRetryConfigInput(v.RetryConfig); err != nil {
			invalidParams.AddNested("RetryConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDependencies(v []types.Dependency) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Dependencies"}
	for i := range v {
		if err := validateDependency(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDependency(v *types.Dependency) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Dependency"}
	if v.Reference == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Reference"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRetryConfigInput(v *types.RetryConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RetryConfigInput"}
	if v.MaxRetries == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MaxRetries"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateVisualReferenceInput(v *types.VisualReferenceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VisualReferenceInput"}
	if v.BaseScreenshots != nil {
		if err := validateBaseScreenshots(v.BaseScreenshots); err != nil {
			invalidParams.AddNested("BaseScreenshots", err.(smithy.InvalidParamsError))
		}
	}
	if v.BaseCanaryRunId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BaseCanaryRunId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateResourceInput(v *AssociateResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateResourceInput"}
	if v.GroupIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GroupIdentifier"))
	}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateCanaryInput(v *CreateCanaryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateCanaryInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Code == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Code"))
	} else if v.Code != nil {
		if err := validateCanaryCodeInput(v.Code); err != nil {
			invalidParams.AddNested("Code", err.(smithy.InvalidParamsError))
		}
	}
	if v.ArtifactS3Location == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ArtifactS3Location"))
	}
	if v.ExecutionRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExecutionRoleArn"))
	}
	if v.Schedule == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Schedule"))
	} else if v.Schedule != nil {
		if err := validateCanaryScheduleInput(v.Schedule); err != nil {
			invalidParams.AddNested("Schedule", err.(smithy.InvalidParamsError))
		}
	}
	if v.RuntimeVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RuntimeVersion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateGroupInput(v *CreateGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateGroupInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteCanaryInput(v *DeleteCanaryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteCanaryInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteGroupInput(v *DeleteGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteGroupInput"}
	if v.GroupIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GroupIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateResourceInput(v *DisassociateResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateResourceInput"}
	if v.GroupIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GroupIdentifier"))
	}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCanaryInput(v *GetCanaryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCanaryInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCanaryRunsInput(v *GetCanaryRunsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCanaryRunsInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetGroupInput(v *GetGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetGroupInput"}
	if v.GroupIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GroupIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListAssociatedGroupsInput(v *ListAssociatedGroupsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListAssociatedGroupsInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListGroupResourcesInput(v *ListGroupResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListGroupResourcesInput"}
	if v.GroupIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GroupIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartCanaryDryRunInput(v *StartCanaryDryRunInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartCanaryDryRunInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Code != nil {
		if err := validateCanaryCodeInput(v.Code); err != nil {
			invalidParams.AddNested("Code", err.(smithy.InvalidParamsError))
		}
	}
	if v.VisualReference != nil {
		if err := validateVisualReferenceInput(v.VisualReference); err != nil {
			invalidParams.AddNested("VisualReference", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartCanaryInput(v *StartCanaryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartCanaryInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopCanaryInput(v *StopCanaryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopCanaryInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateCanaryInput(v *UpdateCanaryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateCanaryInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Code != nil {
		if err := validateCanaryCodeInput(v.Code); err != nil {
			invalidParams.AddNested("Code", err.(smithy.InvalidParamsError))
		}
	}
	if v.Schedule != nil {
		if err := validateCanaryScheduleInput(v.Schedule); err != nil {
			invalidParams.AddNested("Schedule", err.(smithy.InvalidParamsError))
		}
	}
	if v.VisualReference != nil {
		if err := validateVisualReferenceInput(v.VisualReference); err != nil {
			invalidParams.AddNested("VisualReference", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
