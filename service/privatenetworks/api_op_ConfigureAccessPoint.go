// Code generated by smithy-go-codegen DO NOT EDIT.

package privatenetworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/privatenetworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Configures the specified network resource.
//
// Use this action to specify the geographic position of the hardware. You must
// provide Certified Professional Installer (CPI) credentials in the request so
// that we can obtain spectrum grants. For more information, see [Radio units]in the Amazon Web
// Services Private 5G User Guide.
//
// [Radio units]: https://docs.aws.amazon.com/private-networks/latest/userguide/radio-units.html
func (c *Client) ConfigureAccessPoint(ctx context.Context, params *ConfigureAccessPointInput, optFns ...func(*Options)) (*ConfigureAccessPointOutput, error) {
	if params == nil {
		params = &ConfigureAccessPointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConfigureAccessPoint", params, optFns, c.addOperationConfigureAccessPointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConfigureAccessPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ConfigureAccessPointInput struct {

	// The Amazon Resource Name (ARN) of the network resource.
	//
	// This member is required.
	AccessPointArn *string

	// A Base64 encoded string of the CPI certificate associated with the CPI user who
	// is certifying the coordinates of the network resource.
	CpiSecretKey *string

	// The CPI user ID of the CPI user who is certifying the coordinates of the
	// network resource.
	CpiUserId *string

	// The CPI password associated with the CPI certificate in cpiSecretKey .
	CpiUserPassword *string

	// The CPI user name of the CPI user who is certifying the coordinates of the
	// radio unit.
	CpiUsername *string

	// The position of the network resource.
	Position *types.Position

	noSmithyDocumentSerde
}

type ConfigureAccessPointOutput struct {

	// Information about the network resource.
	//
	// This member is required.
	AccessPoint *types.NetworkResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationConfigureAccessPointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpConfigureAccessPoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpConfigureAccessPoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ConfigureAccessPoint"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpConfigureAccessPointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opConfigureAccessPoint(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opConfigureAccessPoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ConfigureAccessPoint",
	}
}
