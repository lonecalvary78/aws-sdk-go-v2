// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagentcore

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentcore/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateEvent struct {
}

func (*validateOpCreateEvent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateEvent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateEventInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateEventInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteEvent struct {
}

func (*validateOpDeleteEvent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteEvent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteEventInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteEventInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMemoryRecord struct {
}

func (*validateOpDeleteMemoryRecord) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMemoryRecord) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMemoryRecordInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMemoryRecordInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetBrowserSession struct {
}

func (*validateOpGetBrowserSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetBrowserSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetBrowserSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetBrowserSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCodeInterpreterSession struct {
}

func (*validateOpGetCodeInterpreterSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCodeInterpreterSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCodeInterpreterSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCodeInterpreterSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetEvent struct {
}

func (*validateOpGetEvent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetEvent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetEventInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetEventInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetMemoryRecord struct {
}

func (*validateOpGetMemoryRecord) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMemoryRecord) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMemoryRecordInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMemoryRecordInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResourceApiKey struct {
}

func (*validateOpGetResourceApiKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResourceApiKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResourceApiKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResourceApiKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResourceOauth2Token struct {
}

func (*validateOpGetResourceOauth2Token) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResourceOauth2Token) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResourceOauth2TokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResourceOauth2TokenInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetWorkloadAccessTokenForJWT struct {
}

func (*validateOpGetWorkloadAccessTokenForJWT) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetWorkloadAccessTokenForJWT) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetWorkloadAccessTokenForJWTInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetWorkloadAccessTokenForJWTInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetWorkloadAccessTokenForUserId struct {
}

func (*validateOpGetWorkloadAccessTokenForUserId) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetWorkloadAccessTokenForUserId) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetWorkloadAccessTokenForUserIdInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetWorkloadAccessTokenForUserIdInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetWorkloadAccessToken struct {
}

func (*validateOpGetWorkloadAccessToken) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetWorkloadAccessToken) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetWorkloadAccessTokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetWorkloadAccessTokenInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpInvokeAgentRuntime struct {
}

func (*validateOpInvokeAgentRuntime) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpInvokeAgentRuntime) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*InvokeAgentRuntimeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpInvokeAgentRuntimeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpInvokeCodeInterpreter struct {
}

func (*validateOpInvokeCodeInterpreter) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpInvokeCodeInterpreter) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*InvokeCodeInterpreterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpInvokeCodeInterpreterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListActors struct {
}

func (*validateOpListActors) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListActors) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListActorsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListActorsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListBrowserSessions struct {
}

func (*validateOpListBrowserSessions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListBrowserSessions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListBrowserSessionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListBrowserSessionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListCodeInterpreterSessions struct {
}

func (*validateOpListCodeInterpreterSessions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListCodeInterpreterSessions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListCodeInterpreterSessionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListCodeInterpreterSessionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListEvents struct {
}

func (*validateOpListEvents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListEvents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListEventsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListEventsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListMemoryRecords struct {
}

func (*validateOpListMemoryRecords) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListMemoryRecords) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListMemoryRecordsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListMemoryRecordsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListSessions struct {
}

func (*validateOpListSessions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListSessions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListSessionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListSessionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRetrieveMemoryRecords struct {
}

func (*validateOpRetrieveMemoryRecords) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRetrieveMemoryRecords) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RetrieveMemoryRecordsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRetrieveMemoryRecordsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartBrowserSession struct {
}

func (*validateOpStartBrowserSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartBrowserSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartBrowserSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartBrowserSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartCodeInterpreterSession struct {
}

func (*validateOpStartCodeInterpreterSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartCodeInterpreterSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartCodeInterpreterSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartCodeInterpreterSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopBrowserSession struct {
}

func (*validateOpStopBrowserSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopBrowserSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopBrowserSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopBrowserSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopCodeInterpreterSession struct {
}

func (*validateOpStopCodeInterpreterSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopCodeInterpreterSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopCodeInterpreterSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopCodeInterpreterSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateBrowserStream struct {
}

func (*validateOpUpdateBrowserStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateBrowserStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateBrowserStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateBrowserStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateEventValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateEvent{}, middleware.After)
}

func addOpDeleteEventValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteEvent{}, middleware.After)
}

func addOpDeleteMemoryRecordValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMemoryRecord{}, middleware.After)
}

func addOpGetBrowserSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetBrowserSession{}, middleware.After)
}

func addOpGetCodeInterpreterSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCodeInterpreterSession{}, middleware.After)
}

func addOpGetEventValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetEvent{}, middleware.After)
}

func addOpGetMemoryRecordValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMemoryRecord{}, middleware.After)
}

func addOpGetResourceApiKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetResourceApiKey{}, middleware.After)
}

func addOpGetResourceOauth2TokenValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetResourceOauth2Token{}, middleware.After)
}

func addOpGetWorkloadAccessTokenForJWTValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetWorkloadAccessTokenForJWT{}, middleware.After)
}

func addOpGetWorkloadAccessTokenForUserIdValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetWorkloadAccessTokenForUserId{}, middleware.After)
}

func addOpGetWorkloadAccessTokenValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetWorkloadAccessToken{}, middleware.After)
}

func addOpInvokeAgentRuntimeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpInvokeAgentRuntime{}, middleware.After)
}

func addOpInvokeCodeInterpreterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpInvokeCodeInterpreter{}, middleware.After)
}

func addOpListActorsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListActors{}, middleware.After)
}

func addOpListBrowserSessionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListBrowserSessions{}, middleware.After)
}

func addOpListCodeInterpreterSessionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListCodeInterpreterSessions{}, middleware.After)
}

func addOpListEventsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListEvents{}, middleware.After)
}

func addOpListMemoryRecordsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListMemoryRecords{}, middleware.After)
}

func addOpListSessionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListSessions{}, middleware.After)
}

func addOpRetrieveMemoryRecordsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRetrieveMemoryRecords{}, middleware.After)
}

func addOpStartBrowserSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartBrowserSession{}, middleware.After)
}

func addOpStartCodeInterpreterSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartCodeInterpreterSession{}, middleware.After)
}

func addOpStopBrowserSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopBrowserSession{}, middleware.After)
}

func addOpStopCodeInterpreterSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopCodeInterpreterSession{}, middleware.After)
}

func addOpUpdateBrowserStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateBrowserStream{}, middleware.After)
}

func validateBranch(v *types.Branch) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Branch"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateBranchFilter(v *types.BranchFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BranchFilter"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateConversational(v *types.Conversational) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Conversational"}
	if v.Content == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Content"))
	}
	if len(v.Role) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Role"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFilterInput(v *types.FilterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FilterInput"}
	if v.Branch != nil {
		if err := validateBranchFilter(v.Branch); err != nil {
			invalidParams.AddNested("Branch", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateInputContentBlock(v *types.InputContentBlock) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InputContentBlock"}
	if v.Path == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Path"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateInputContentBlockList(v []types.InputContentBlock) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InputContentBlockList"}
	for i := range v {
		if err := validateInputContentBlock(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePayloadType(v types.PayloadType) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PayloadType"}
	switch uv := v.(type) {
	case *types.PayloadTypeMemberConversational:
		if err := validateConversational(&uv.Value); err != nil {
			invalidParams.AddNested("[conversational]", err.(smithy.InvalidParamsError))
		}

	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePayloadTypeList(v []types.PayloadType) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PayloadTypeList"}
	for i := range v {
		if err := validatePayloadType(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSearchCriteria(v *types.SearchCriteria) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchCriteria"}
	if v.SearchQuery == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SearchQuery"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateToolArguments(v *types.ToolArguments) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ToolArguments"}
	if v.Content != nil {
		if err := validateInputContentBlockList(v.Content); err != nil {
			invalidParams.AddNested("Content", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateViewPort(v *types.ViewPort) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ViewPort"}
	if v.Width == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Width"))
	}
	if v.Height == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Height"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateEventInput(v *CreateEventInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateEventInput"}
	if v.MemoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemoryId"))
	}
	if v.ActorId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActorId"))
	}
	if v.EventTimestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EventTimestamp"))
	}
	if v.Payload == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Payload"))
	} else if v.Payload != nil {
		if err := validatePayloadTypeList(v.Payload); err != nil {
			invalidParams.AddNested("Payload", err.(smithy.InvalidParamsError))
		}
	}
	if v.Branch != nil {
		if err := validateBranch(v.Branch); err != nil {
			invalidParams.AddNested("Branch", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteEventInput(v *DeleteEventInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteEventInput"}
	if v.MemoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemoryId"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if v.EventId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EventId"))
	}
	if v.ActorId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActorId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMemoryRecordInput(v *DeleteMemoryRecordInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMemoryRecordInput"}
	if v.MemoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemoryId"))
	}
	if v.MemoryRecordId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemoryRecordId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetBrowserSessionInput(v *GetBrowserSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetBrowserSessionInput"}
	if v.BrowserIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BrowserIdentifier"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCodeInterpreterSessionInput(v *GetCodeInterpreterSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCodeInterpreterSessionInput"}
	if v.CodeInterpreterIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CodeInterpreterIdentifier"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetEventInput(v *GetEventInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetEventInput"}
	if v.MemoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemoryId"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if v.ActorId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActorId"))
	}
	if v.EventId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EventId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMemoryRecordInput(v *GetMemoryRecordInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMemoryRecordInput"}
	if v.MemoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemoryId"))
	}
	if v.MemoryRecordId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemoryRecordId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResourceApiKeyInput(v *GetResourceApiKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResourceApiKeyInput"}
	if v.WorkloadIdentityToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkloadIdentityToken"))
	}
	if v.ResourceCredentialProviderName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceCredentialProviderName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResourceOauth2TokenInput(v *GetResourceOauth2TokenInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResourceOauth2TokenInput"}
	if v.WorkloadIdentityToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkloadIdentityToken"))
	}
	if v.ResourceCredentialProviderName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceCredentialProviderName"))
	}
	if v.Scopes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Scopes"))
	}
	if len(v.Oauth2Flow) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Oauth2Flow"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetWorkloadAccessTokenForJWTInput(v *GetWorkloadAccessTokenForJWTInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetWorkloadAccessTokenForJWTInput"}
	if v.WorkloadName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkloadName"))
	}
	if v.UserToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetWorkloadAccessTokenForUserIdInput(v *GetWorkloadAccessTokenForUserIdInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetWorkloadAccessTokenForUserIdInput"}
	if v.WorkloadName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkloadName"))
	}
	if v.UserId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetWorkloadAccessTokenInput(v *GetWorkloadAccessTokenInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetWorkloadAccessTokenInput"}
	if v.WorkloadName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkloadName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpInvokeAgentRuntimeInput(v *InvokeAgentRuntimeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InvokeAgentRuntimeInput"}
	if v.AgentRuntimeArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AgentRuntimeArn"))
	}
	if v.Payload == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Payload"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpInvokeCodeInterpreterInput(v *InvokeCodeInterpreterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InvokeCodeInterpreterInput"}
	if v.CodeInterpreterIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CodeInterpreterIdentifier"))
	}
	if len(v.Name) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Arguments != nil {
		if err := validateToolArguments(v.Arguments); err != nil {
			invalidParams.AddNested("Arguments", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListActorsInput(v *ListActorsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListActorsInput"}
	if v.MemoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemoryId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListBrowserSessionsInput(v *ListBrowserSessionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListBrowserSessionsInput"}
	if v.BrowserIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BrowserIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListCodeInterpreterSessionsInput(v *ListCodeInterpreterSessionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListCodeInterpreterSessionsInput"}
	if v.CodeInterpreterIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CodeInterpreterIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListEventsInput(v *ListEventsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListEventsInput"}
	if v.MemoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemoryId"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if v.ActorId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActorId"))
	}
	if v.Filter != nil {
		if err := validateFilterInput(v.Filter); err != nil {
			invalidParams.AddNested("Filter", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListMemoryRecordsInput(v *ListMemoryRecordsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListMemoryRecordsInput"}
	if v.MemoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemoryId"))
	}
	if v.Namespace == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Namespace"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListSessionsInput(v *ListSessionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListSessionsInput"}
	if v.MemoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemoryId"))
	}
	if v.ActorId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActorId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRetrieveMemoryRecordsInput(v *RetrieveMemoryRecordsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RetrieveMemoryRecordsInput"}
	if v.MemoryId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemoryId"))
	}
	if v.Namespace == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Namespace"))
	}
	if v.SearchCriteria == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SearchCriteria"))
	} else if v.SearchCriteria != nil {
		if err := validateSearchCriteria(v.SearchCriteria); err != nil {
			invalidParams.AddNested("SearchCriteria", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartBrowserSessionInput(v *StartBrowserSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartBrowserSessionInput"}
	if v.BrowserIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BrowserIdentifier"))
	}
	if v.ViewPort != nil {
		if err := validateViewPort(v.ViewPort); err != nil {
			invalidParams.AddNested("ViewPort", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartCodeInterpreterSessionInput(v *StartCodeInterpreterSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartCodeInterpreterSessionInput"}
	if v.CodeInterpreterIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CodeInterpreterIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopBrowserSessionInput(v *StopBrowserSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopBrowserSessionInput"}
	if v.BrowserIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BrowserIdentifier"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopCodeInterpreterSessionInput(v *StopCodeInterpreterSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopCodeInterpreterSessionInput"}
	if v.CodeInterpreterIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CodeInterpreterIdentifier"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateBrowserStreamInput(v *UpdateBrowserStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateBrowserStreamInput"}
	if v.BrowserIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BrowserIdentifier"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if v.StreamUpdate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StreamUpdate"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
