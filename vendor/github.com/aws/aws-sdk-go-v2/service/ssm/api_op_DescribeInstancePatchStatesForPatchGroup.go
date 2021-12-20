// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the high-level patch state for the instances in the specified patch
// group.
func (c *Client) DescribeInstancePatchStatesForPatchGroup(ctx context.Context, params *DescribeInstancePatchStatesForPatchGroupInput, optFns ...func(*Options)) (*DescribeInstancePatchStatesForPatchGroupOutput, error) {
	if params == nil {
		params = &DescribeInstancePatchStatesForPatchGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstancePatchStatesForPatchGroup", params, optFns, c.addOperationDescribeInstancePatchStatesForPatchGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstancePatchStatesForPatchGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstancePatchStatesForPatchGroupInput struct {

	// The name of the patch group for which the patch state information should be
	// retrieved.
	//
	// This member is required.
	PatchGroup *string

	// Each entry in the array is a structure containing:
	//
	// * Key (string between 1 and
	// 200 characters)
	//
	// * Values (array containing a single string)
	//
	// * Type (string
	// "Equal", "NotEqual", "LessThan", "GreaterThan")
	Filters []types.InstancePatchStateFilter

	// The maximum number of patches to return (per page).
	MaxResults int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeInstancePatchStatesForPatchGroupOutput struct {

	// The high-level patch state for the requested instances.
	InstancePatchStates []types.InstancePatchState

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstancePatchStatesForPatchGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInstancePatchStatesForPatchGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInstancePatchStatesForPatchGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeInstancePatchStatesForPatchGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstancePatchStatesForPatchGroup(options.Region), middleware.Before); err != nil {
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
	return nil
}

// DescribeInstancePatchStatesForPatchGroupAPIClient is a client that implements
// the DescribeInstancePatchStatesForPatchGroup operation.
type DescribeInstancePatchStatesForPatchGroupAPIClient interface {
	DescribeInstancePatchStatesForPatchGroup(context.Context, *DescribeInstancePatchStatesForPatchGroupInput, ...func(*Options)) (*DescribeInstancePatchStatesForPatchGroupOutput, error)
}

var _ DescribeInstancePatchStatesForPatchGroupAPIClient = (*Client)(nil)

// DescribeInstancePatchStatesForPatchGroupPaginatorOptions is the paginator
// options for DescribeInstancePatchStatesForPatchGroup
type DescribeInstancePatchStatesForPatchGroupPaginatorOptions struct {
	// The maximum number of patches to return (per page).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInstancePatchStatesForPatchGroupPaginator is a paginator for
// DescribeInstancePatchStatesForPatchGroup
type DescribeInstancePatchStatesForPatchGroupPaginator struct {
	options   DescribeInstancePatchStatesForPatchGroupPaginatorOptions
	client    DescribeInstancePatchStatesForPatchGroupAPIClient
	params    *DescribeInstancePatchStatesForPatchGroupInput
	nextToken *string
	firstPage bool
}

// NewDescribeInstancePatchStatesForPatchGroupPaginator returns a new
// DescribeInstancePatchStatesForPatchGroupPaginator
func NewDescribeInstancePatchStatesForPatchGroupPaginator(client DescribeInstancePatchStatesForPatchGroupAPIClient, params *DescribeInstancePatchStatesForPatchGroupInput, optFns ...func(*DescribeInstancePatchStatesForPatchGroupPaginatorOptions)) *DescribeInstancePatchStatesForPatchGroupPaginator {
	if params == nil {
		params = &DescribeInstancePatchStatesForPatchGroupInput{}
	}

	options := DescribeInstancePatchStatesForPatchGroupPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInstancePatchStatesForPatchGroupPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInstancePatchStatesForPatchGroupPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeInstancePatchStatesForPatchGroup page.
func (p *DescribeInstancePatchStatesForPatchGroupPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInstancePatchStatesForPatchGroupOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeInstancePatchStatesForPatchGroup(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeInstancePatchStatesForPatchGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeInstancePatchStatesForPatchGroup",
	}
}
