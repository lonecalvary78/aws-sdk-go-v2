// Code generated by smithy-go-codegen DO NOT EDIT.

package acm

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAddTagsToCertificate struct {
}

func (*validateOpAddTagsToCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAddTagsToCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AddTagsToCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAddTagsToCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteCertificate struct {
}

func (*validateOpDeleteCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeCertificate struct {
}

func (*validateOpDescribeCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpExportCertificate struct {
}

func (*validateOpExportCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpExportCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ExportCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpExportCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCertificate struct {
}

func (*validateOpGetCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpImportCertificate struct {
}

func (*validateOpImportCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpImportCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ImportCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpImportCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForCertificate struct {
}

func (*validateOpListTagsForCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutAccountConfiguration struct {
}

func (*validateOpPutAccountConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutAccountConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutAccountConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutAccountConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemoveTagsFromCertificate struct {
}

func (*validateOpRemoveTagsFromCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemoveTagsFromCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemoveTagsFromCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemoveTagsFromCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRenewCertificate struct {
}

func (*validateOpRenewCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRenewCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RenewCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRenewCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRequestCertificate struct {
}

func (*validateOpRequestCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRequestCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RequestCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRequestCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpResendValidationEmail struct {
}

func (*validateOpResendValidationEmail) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpResendValidationEmail) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ResendValidationEmailInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpResendValidationEmailInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRevokeCertificate struct {
}

func (*validateOpRevokeCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRevokeCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RevokeCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRevokeCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateCertificateOptions struct {
}

func (*validateOpUpdateCertificateOptions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateCertificateOptions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateCertificateOptionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateCertificateOptionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAddTagsToCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAddTagsToCertificate{}, middleware.After)
}

func addOpDeleteCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteCertificate{}, middleware.After)
}

func addOpDescribeCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeCertificate{}, middleware.After)
}

func addOpExportCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpExportCertificate{}, middleware.After)
}

func addOpGetCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCertificate{}, middleware.After)
}

func addOpImportCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpImportCertificate{}, middleware.After)
}

func addOpListTagsForCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForCertificate{}, middleware.After)
}

func addOpPutAccountConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutAccountConfiguration{}, middleware.After)
}

func addOpRemoveTagsFromCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemoveTagsFromCertificate{}, middleware.After)
}

func addOpRenewCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRenewCertificate{}, middleware.After)
}

func addOpRequestCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRequestCertificate{}, middleware.After)
}

func addOpResendValidationEmailValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpResendValidationEmail{}, middleware.After)
}

func addOpRevokeCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRevokeCertificate{}, middleware.After)
}

func addOpUpdateCertificateOptionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateCertificateOptions{}, middleware.After)
}

func validateDomainValidationOption(v *types.DomainValidationOption) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DomainValidationOption"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if v.ValidationDomain == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ValidationDomain"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDomainValidationOptionList(v []types.DomainValidationOption) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DomainValidationOptionList"}
	for i := range v {
		if err := validateDomainValidationOption(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAddTagsToCertificateInput(v *AddTagsToCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddTagsToCertificateInput"}
	if v.CertificateArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteCertificateInput(v *DeleteCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteCertificateInput"}
	if v.CertificateArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeCertificateInput(v *DescribeCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeCertificateInput"}
	if v.CertificateArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpExportCertificateInput(v *ExportCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExportCertificateInput"}
	if v.CertificateArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateArn"))
	}
	if v.Passphrase == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Passphrase"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCertificateInput(v *GetCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCertificateInput"}
	if v.CertificateArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpImportCertificateInput(v *ImportCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImportCertificateInput"}
	if v.Certificate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Certificate"))
	}
	if v.PrivateKey == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PrivateKey"))
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForCertificateInput(v *ListTagsForCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForCertificateInput"}
	if v.CertificateArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutAccountConfigurationInput(v *PutAccountConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutAccountConfigurationInput"}
	if v.IdempotencyToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdempotencyToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemoveTagsFromCertificateInput(v *RemoveTagsFromCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemoveTagsFromCertificateInput"}
	if v.CertificateArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRenewCertificateInput(v *RenewCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RenewCertificateInput"}
	if v.CertificateArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRequestCertificateInput(v *RequestCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RequestCertificateInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if v.DomainValidationOptions != nil {
		if err := validateDomainValidationOptionList(v.DomainValidationOptions); err != nil {
			invalidParams.AddNested("DomainValidationOptions", err.(smithy.InvalidParamsError))
		}
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpResendValidationEmailInput(v *ResendValidationEmailInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResendValidationEmailInput"}
	if v.CertificateArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateArn"))
	}
	if v.Domain == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Domain"))
	}
	if v.ValidationDomain == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ValidationDomain"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRevokeCertificateInput(v *RevokeCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RevokeCertificateInput"}
	if v.CertificateArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateArn"))
	}
	if len(v.RevocationReason) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("RevocationReason"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateCertificateOptionsInput(v *UpdateCertificateOptionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateCertificateOptionsInput"}
	if v.CertificateArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateArn"))
	}
	if v.Options == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Options"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
