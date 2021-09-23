// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a log group with the specified name. You can create up to 20,000 log
// groups per account. You must use the following guidelines when naming a log
// group:
//
// * Log group names must be unique within a region for an Amazon Web
// Services account.
//
// * Log group names can be between 1 and 512 characters
// long.
//
// * Log group names consist of the following characters: a-z, A-Z, 0-9, '_'
// (underscore), '-' (hyphen), '/' (forward slash), '.' (period), and '#' (number
// sign)
//
// When you create a log group, by default the log events in the log group
// never expire. To set a retention policy so that events expire and are deleted
// after a specified time, use PutRetentionPolicy
// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutRetentionPolicy.html).
// If you associate a Key Management Service customer master key (CMK) with the log
// group, ingested data is encrypted using the CMK. This association is stored as
// long as the data encrypted with the CMK is still within CloudWatch Logs. This
// enables CloudWatch Logs to decrypt this data whenever it is requested. If you
// attempt to associate a CMK with the log group but the CMK does not exist or the
// CMK is disabled, you receive an InvalidParameterException error. CloudWatch Logs
// supports only symmetric CMKs. Do not associate an asymmetric CMK with your log
// group. For more information, see Using Symmetric and Asymmetric Keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html).
func (c *Client) CreateLogGroup(ctx context.Context, params *CreateLogGroupInput, optFns ...func(*Options)) (*CreateLogGroupOutput, error) {
	if params == nil {
		params = &CreateLogGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLogGroup", params, optFns, c.addOperationCreateLogGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLogGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLogGroupInput struct {

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data. For
	// more information, see Amazon Resource Names - Key Management Service
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kms).
	KmsKeyId *string

	// The key-value pairs to use for the tags. CloudWatch Logs doesn’t support IAM
	// policies that prevent users from assigning specified tags to log groups using
	// the aws:Resource/key-name  or aws:TagKeys condition keys. For more information
	// about using tags to control access, see Controlling access to Amazon Web
	// Services resources using tags
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html).
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateLogGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLogGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLogGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLogGroup{}, middleware.After)
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
	if err = addOpCreateLogGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLogGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLogGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "CreateLogGroup",
	}
}
