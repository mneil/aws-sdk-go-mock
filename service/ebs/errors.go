// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ebs

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient access to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConcurrentLimitExceededException for service response error code
	// "ConcurrentLimitExceededException".
	//
	// You have reached the limit for concurrent API requests. For more information,
	// see Optimizing performance of the EBS direct APIs (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-accessing-snapshot.html#ebsapi-performance)
	// in the Amazon Elastic Compute Cloud User Guide.
	ErrCodeConcurrentLimitExceededException = "ConcurrentLimitExceededException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The request uses the same client token as a previous, but non-identical request.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// An internal error has occurred. For more information see Error retries (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html).
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeRequestThrottledException for service response error code
	// "RequestThrottledException".
	//
	// The number of API requests has exceeded the maximum allowed API request throttling
	// limit for the snapshot. For more information see Error retries (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html).
	ErrCodeRequestThrottledException = "RequestThrottledException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// Your current service quotas do not allow you to perform this action.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The input fails to satisfy the constraints of the EBS direct APIs.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":            newErrorAccessDeniedException,
	"ConcurrentLimitExceededException": newErrorConcurrentLimitExceededException,
	"ConflictException":                newErrorConflictException,
	"InternalServerException":          newErrorInternalServerException,
	"RequestThrottledException":        newErrorRequestThrottledException,
	"ResourceNotFoundException":        newErrorResourceNotFoundException,
	"ServiceQuotaExceededException":    newErrorServiceQuotaExceededException,
	"ValidationException":              newErrorValidationException,
}
