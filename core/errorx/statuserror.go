package errorx

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ============================================================================
// gRPC Status Error 便捷构造函数
// ============================================================================

// NewOKError returns status error with OK code (rarely used)
func NewOKError(msg string) error {
	return status.Error(codes.OK, msg)
}

// NewCanceledError returns status error with Canceled code
func NewCanceledError(msg string) error {
	return status.Error(codes.Canceled, msg)
}

// NewUnknownError returns status error with Unknown code
func NewUnknownError(msg string) error {
	return status.Error(codes.Unknown, msg)
}

// NewInvalidArgumentError returns status error with InvalidArgument code
func NewInvalidArgumentError(msg string) error {
	return status.Error(codes.InvalidArgument, msg)
}

// NewDeadlineExceededError returns status error with DeadlineExceeded code
func NewDeadlineExceededError(msg string) error {
	return status.Error(codes.DeadlineExceeded, msg)
}

// NewNotFoundError returns status error with NotFound code
func NewNotFoundError(msg string) error {
	return status.Error(codes.NotFound, msg)
}

// NewAlreadyExistsError returns status error with AlreadyExists code
func NewAlreadyExistsError(msg string) error {
	return status.Error(codes.AlreadyExists, msg)
}

// NewPermissionDeniedError returns status error with PermissionDenied code
func NewPermissionDeniedError(msg string) error {
	return status.Error(codes.PermissionDenied, msg)
}

// NewResourceExhaustedError returns status error with ResourceExhausted code
func NewResourceExhaustedError(msg string) error {
	return status.Error(codes.ResourceExhausted, msg)
}

// NewFailedPreconditionError returns status error with FailedPrecondition code
func NewFailedPreconditionError(msg string) error {
	return status.Error(codes.FailedPrecondition, msg)
}

// NewAbortedError returns status error with Aborted code
func NewAbortedError(msg string) error {
	return status.Error(codes.Aborted, msg)
}

// NewOutOfRangeError returns status error with OutOfRange code
func NewOutOfRangeError(msg string) error {
	return status.Error(codes.OutOfRange, msg)
}

// NewUnimplementedError returns status error with Unimplemented code
func NewUnimplementedError(msg string) error {
	return status.Error(codes.Unimplemented, msg)
}

// NewInternalError returns status error with Internal code
func NewInternalError(msg string) error {
	return status.Error(codes.Internal, msg)
}

// NewUnavailableError returns status error with Unavailable code
func NewUnavailableError(msg string) error {
	return status.Error(codes.Unavailable, msg)
}

// NewDataLossError returns status error with DataLoss code
func NewDataLossError(msg string) error {
	return status.Error(codes.DataLoss, msg)
}

// NewUnauthenticatedError returns status error with Unauthenticated code
func NewUnauthenticatedError(msg string) error {
	return status.Error(codes.Unauthenticated, msg)
}

// ============================================================================
// 带格式化参数的便捷构造函数
// ============================================================================

// NewInvalidArgumentErrorf returns status error with InvalidArgument code and formatted message
func NewInvalidArgumentErrorf(format string, args ...any) error {
	return status.Errorf(codes.InvalidArgument, format, args...)
}

// NewNotFoundErrorf returns status error with NotFound code and formatted message
func NewNotFoundErrorf(format string, args ...any) error {
	return status.Errorf(codes.NotFound, format, args...)
}

// NewAlreadyExistsErrorf returns status error with AlreadyExists code and formatted message
func NewAlreadyExistsErrorf(format string, args ...any) error {
	return status.Errorf(codes.AlreadyExists, format, args...)
}

// NewPermissionDeniedErrorf returns status error with PermissionDenied code and formatted message
func NewPermissionDeniedErrorf(format string, args ...any) error {
	return status.Errorf(codes.PermissionDenied, format, args...)
}

// NewInternalErrorf returns status error with Internal code and formatted message
func NewInternalErrorf(format string, args ...any) error {
	return status.Errorf(codes.Internal, format, args...)
}

// NewUnavailableErrorf returns status error with Unavailable code and formatted message
func NewUnavailableErrorf(format string, args ...any) error {
	return status.Errorf(codes.Unavailable, format, args...)
}

// NewUnauthenticatedErrorf returns status error with Unauthenticated code and formatted message
func NewUnauthenticatedErrorf(format string, args ...any) error {
	return status.Errorf(codes.Unauthenticated, format, args...)
}

// NewDeadlineExceededErrorf returns status error with DeadlineExceeded code and formatted message
func NewDeadlineExceededErrorf(format string, args ...any) error {
	return status.Errorf(codes.DeadlineExceeded, format, args...)
}

// NewFailedPreconditionErrorf returns status error with FailedPrecondition code and formatted message
func NewFailedPreconditionErrorf(format string, args ...any) error {
	return status.Errorf(codes.FailedPrecondition, format, args...)
}

// NewResourceExhaustedErrorf returns status error with ResourceExhausted code and formatted message
func NewResourceExhaustedErrorf(format string, args ...any) error {
	return status.Errorf(codes.ResourceExhausted, format, args...)
}

// ============================================================================
// 常用业务场景的便捷函数
// ============================================================================

// NewValidationError returns InvalidArgument error for validation failures
func NewValidationError(field, reason string) error {
	return status.Errorf(codes.InvalidArgument, "validation failed for field '%s': %s", field, reason)
}

// NewRequiredFieldError returns InvalidArgument error for missing required fields
func NewRequiredFieldError(field string) error {
	return status.Errorf(codes.InvalidArgument, "required field '%s' is missing", field)
}

// NewResourceNotFoundError returns NotFound error for specific resource
func NewResourceNotFoundError(resourceType, resourceId string) error {
	return status.Errorf(codes.NotFound, "%s with id '%s' not found", resourceType, resourceId)
}

// NewResourceAlreadyExistsError returns AlreadyExists error for specific resource
func NewResourceAlreadyExistsError(resourceType, resourceId string) error {
	return status.Errorf(codes.AlreadyExists, "%s with id '%s' already exists", resourceType, resourceId)
}

// NewAccessDeniedError returns PermissionDenied error for access control
func NewAccessDeniedError(resource, action string) error {
	return status.Errorf(codes.PermissionDenied, "access denied: insufficient permissions to %s %s", action, resource)
}

// NewTokenExpiredError returns Unauthenticated error for expired tokens
func NewTokenExpiredError() error {
	return status.Error(codes.Unauthenticated, "token has expired")
}

// NewInvalidTokenError returns Unauthenticated error for invalid tokens
func NewInvalidTokenError() error {
	return status.Error(codes.Unauthenticated, "invalid token")
}

// NewRateLimitExceededError returns ResourceExhausted error for rate limiting
func NewRateLimitExceededError() error {
	return status.Error(codes.ResourceExhausted, "rate limit exceeded")
}

// NewServiceUnavailableError returns Unavailable error for service maintenance
func NewServiceUnavailableError(reason string) error {
	if reason == "" {
		reason = "service temporarily unavailable"
	}
	return status.Error(codes.Unavailable, reason)
}

// NewTimeoutError returns DeadlineExceeded error for operation timeouts
func NewTimeoutError(operation string) error {
	return status.Errorf(codes.DeadlineExceeded, "operation '%s' timed out", operation)
}

// NewDatabaseError returns Internal error for database issues
func NewDatabaseError(operation string, err error) error {
	return status.Errorf(codes.Internal, "database error during %s: %v", operation, err)
}

// NewExternalServiceError returns Unavailable error for external service failures
func NewExternalServiceError(service string, err error) error {
	return status.Errorf(codes.Unavailable, "external service '%s' error: %v", service, err)
}

// ============================================================================
// 错误检查和转换工具函数
// ============================================================================

// IsGrpcError checks if the error is a gRPC status error
func IsGrpcError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := status.FromError(err)
	return ok
}

// GetGrpcCode extracts the gRPC status code from error
func GetGrpcCode(err error) codes.Code {
	if err == nil {
		return codes.OK
	}
	if st, ok := status.FromError(err); ok {
		return st.Code()
	}
	return codes.Unknown
}

// GetGrpcMessage extracts the gRPC status message from error
func GetGrpcMessage(err error) string {
	if err == nil {
		return ""
	}
	if st, ok := status.FromError(err); ok {
		return st.Message()
	}
	return err.Error()
}

// IsCode checks if the error has the specified gRPC status code
func IsCode(err error, code codes.Code) bool {
	return GetGrpcCode(err) == code
}

// IsNotFound checks if the error is a NotFound error
func IsNotFound(err error) bool {
	return IsCode(err, codes.NotFound)
}

// IsInvalidArgument checks if the error is an InvalidArgument error
func IsInvalidArgument(err error) bool {
	return IsCode(err, codes.InvalidArgument)
}

// IsPermissionDenied checks if the error is a PermissionDenied error
func IsPermissionDenied(err error) bool {
	return IsCode(err, codes.PermissionDenied)
}

// IsUnauthenticated checks if the error is an Unauthenticated error
func IsUnauthenticated(err error) bool {
	return IsCode(err, codes.Unauthenticated)
}

// IsAlreadyExists checks if the error is an AlreadyExists error
func IsAlreadyExists(err error) bool {
	return IsCode(err, codes.AlreadyExists)
}

// IsInternal checks if the error is an Internal error
func IsInternal(err error) bool {
	return IsCode(err, codes.Internal)
}

// IsUnavailable checks if the error is an Unavailable error
func IsUnavailable(err error) bool {
	return IsCode(err, codes.Unavailable)
}

// IsDeadlineExceeded checks if the error is a DeadlineExceeded error
func IsDeadlineExceeded(err error) bool {
	return IsCode(err, codes.DeadlineExceeded)
}

// ConvertToGrpcError converts a regular error to gRPC status error
func ConvertToGrpcError(err error) error {
	if err == nil {
		return nil
	}
	if IsGrpcError(err) {
		return err
	}
	return status.Error(codes.Internal, err.Error())
}
