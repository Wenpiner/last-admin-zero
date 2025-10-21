package errorx

import "net/http"

// ApiError represents an API error with both HTTP status code and business error code
type ApiError struct {
	Code    int    `json:"code"`    // Business error code
	Message string `json:"message"` // Error message
	Status  int    `json:"-"`  // HTTP status code
}

// Error implements the error interface
func (e *ApiError) Error() string {
	return e.Message
}

// GetStatus returns the HTTP status code
func (e *ApiError) GetStatus() int {
	return e.Status
}

// GetCode returns the business error code
func (e *ApiError) GetCode() int {
	return e.Code
}

// Business Error Codes
// 定义业务错误码常量，采用分层设计便于维护和扩展
const (
	// 系统级错误码 (10000-19999)
	CodeSuccess            = 0     // 成功
	CodeSystemError        = 10000 // 系统错误
	CodeInternalError      = 10001 // 内部服务器错误
	CodeServiceUnavailable = 10002 // 服务不可用
	CodeTimeout            = 10003 // 请求超时
	CodeTooManyRequests    = 10004 // 请求过于频繁
	CodeMaintenance        = 10005 // 系统维护中

	// 请求参数错误码 (20000-29999)
	CodeBadRequest      = 20000 // 请求参数错误
	CodeInvalidParams   = 20001 // 参数格式错误
	CodeMissingParams   = 20002 // 缺少必要参数
	CodeParamOutOfRange = 20003 // 参数超出范围
	CodeInvalidFormat   = 20004 // 格式错误
	CodeInvalidJSON     = 20005 // JSON格式错误
	CodeInvalidFileType = 20006 // 文件类型错误
	CodeFileTooLarge    = 20007 // 文件过大

	// 认证授权错误码 (30000-39999)
	CodeUnauthorized       = 30000 // 未认证
	CodeInvalidToken       = 30001 // 无效的令牌
	CodeTokenExpired       = 30002 // 令牌已过期
	CodeTokenMalformed     = 30003 // 令牌格式错误
	CodeForbidden          = 30004 // 权限不足
	CodeAccessDenied       = 30005 // 访问被拒绝
	CodeInvalidCredentials = 30006 // 凭据无效
	CodeAccountLocked      = 30007 // 账户被锁定
	CodeAccountDisabled    = 30008 // 账户被禁用
	CodeTOTPRequired       = 30009 // 需要TOTP验证
	CodeTOTPVerifyFailed   = 30010 // TOTP验证失败

	// 资源错误码 (40000-49999)
	CodeNotFound         = 40000 // 资源不存在
	CodeResourceNotFound = 40001 // 指定资源不存在
	CodeUserNotFound     = 40002 // 用户不存在
	CodeRecordNotFound   = 40003 // 记录不存在
	CodePageNotFound     = 40004 // 页面不存在

	// 业务逻辑错误码 (50000-59999)
	CodeBusinessError    = 50000 // 业务逻辑错误
	CodeDataConflict     = 50001 // 数据冲突
	CodeResourceExists   = 50002 // 资源已存在
	CodeOperationFailed  = 50003 // 操作失败
	CodeInvalidOperation = 50004 // 无效操作
	CodeStatusError      = 50005 // 状态错误
	CodeQuotaExceeded    = 50006 // 配额超限
	CodeDependencyError  = 50007 // 依赖错误

	// 数据库错误码 (60000-69999)
	CodeDatabaseError       = 60000 // 数据库错误
	CodeConnectionFailed    = 60001 // 数据库连接失败
	CodeQueryFailed         = 60002 // 查询失败
	CodeTransactionFailed   = 60003 // 事务失败
	CodeConstraintViolation = 60004 // 约束违反
	CodeDuplicateEntry      = 60005 // 重复条目

	// 外部服务错误码 (70000-79999)
	CodeExternalError   = 70000 // 外部服务错误
	CodeThirdPartyError = 70001 // 第三方服务错误
	CodeNetworkError    = 70002 // 网络错误
	CodeGatewayError    = 70003 // 网关错误
	CodeUpstreamError   = 70004 // 上游服务错误
)

// ============================================================================
// 基础构造函数
// ============================================================================

// NewApiError creates a new ApiError with custom business code, message and HTTP status
func NewApiError(code int, message string, status int) *ApiError {
	return &ApiError{
		Code:    code,
		Message: message,
		Status:  status,
	}
}

// NewApiErrorWithoutStatus creates a new ApiError with business code and message,
// HTTP status defaults to 400 Bad Request
func NewApiErrorWithoutStatus(code int, message string) *ApiError {
	return &ApiError{
		Code:    code,
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

// NewApiErrorWithDefaultCode creates a new ApiError with message and HTTP status,
// business code defaults to CodeSystemError
func NewApiErrorWithDefaultCode(message string, status int) *ApiError {
	return &ApiError{
		Code:    CodeSystemError,
		Message: message,
		Status:  status,
	}
}


// ============================================================================
// HTTP 状态码相关的便捷构造函数
// ============================================================================

// NewApiInternalError creates an ApiError for internal server errors (500)
func NewApiInternalError(message string) *ApiError {
	if message == "" {
		message = "Internal server error"
	}
	return &ApiError{
		Code:    CodeInternalError,
		Message: message,
		Status:  http.StatusInternalServerError,
	}
}

// NewApiBadRequestError creates an ApiError for bad request errors (400)
func NewApiBadRequestError(message string) *ApiError {
	if message == "" {
		message = "Bad request"
	}
	return &ApiError{
		Code:    CodeBadRequest,
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

// NewApiUnauthorizedError creates an ApiError for unauthorized errors (401)
func NewApiUnauthorizedError(message string) *ApiError {
	if message == "" {
		message = "Unauthorized"
	}
	return &ApiError{
		Code:    CodeUnauthorized,
		Message: message,
		Status:  http.StatusUnauthorized,
	}
}

// NewApiForbiddenError creates an ApiError for forbidden errors (403)
func NewApiForbiddenError(message string) *ApiError {
	if message == "" {
		message = "Forbidden"
	}
	return &ApiError{
		Code:    CodeForbidden,
		Message: message,
		Status:  http.StatusForbidden,
	}
}

// NewApiNotFoundError creates an ApiError for not found errors (404)
func NewApiNotFoundError(message string) *ApiError {
	if message == "" {
		message = "Resource not found"
	}
	return &ApiError{
		Code:    CodeNotFound,
		Message: message,
		Status:  http.StatusNotFound,
	}
}

// NewApiConflictError creates an ApiError for conflict errors (409)
func NewApiConflictError(message string) *ApiError {
	if message == "" {
		message = "Resource conflict"
	}
	return &ApiError{
		Code:    CodeDataConflict,
		Message: message,
		Status:  http.StatusConflict,
	}
}

// NewApiUnprocessableEntityError creates an ApiError for unprocessable entity errors (422)
func NewApiUnprocessableEntityError(message string) *ApiError {
	if message == "" {
		message = "Unprocessable entity"
	}
	return &ApiError{
		Code:    CodeInvalidParams,
		Message: message,
		Status:  http.StatusUnprocessableEntity,
	}
}

// NewApiTooManyRequestsError creates an ApiError for too many requests errors (429)
func NewApiTooManyRequestsError(message string) *ApiError {
	if message == "" {
		message = "Too many requests"
	}
	return &ApiError{
		Code:    CodeTooManyRequests,
		Message: message,
		Status:  http.StatusTooManyRequests,
	}
}

// NewApiBadGatewayError creates an ApiError for bad gateway errors (502)
func NewApiBadGatewayError(message string) *ApiError {
	if message == "" {
		message = "Bad gateway"
	}
	return &ApiError{
		Code:    CodeGatewayError,
		Message: message,
		Status:  http.StatusBadGateway,
	}
}

// NewApiServiceUnavailableError creates an ApiError for service unavailable errors (503)
func NewApiServiceUnavailableError(message string) *ApiError {
	if message == "" {
		message = "Service unavailable"
	}
	return &ApiError{
		Code:    CodeServiceUnavailable,
		Message: message,
		Status:  http.StatusServiceUnavailable,
	}
}

// NewApiGatewayTimeoutError creates an ApiError for gateway timeout errors (504)
func NewApiGatewayTimeoutError(message string) *ApiError {
	if message == "" {
		message = "Gateway timeout"
	}
	return &ApiError{
		Code:    CodeTimeout,
		Message: message,
		Status:  http.StatusGatewayTimeout,
	}
}

// ============================================================================
// 业务场景相关的便捷构造函数
// ============================================================================

// NewApiInvalidParamsError creates an ApiError for invalid parameters
func NewApiInvalidParamsError(message string) *ApiError {
	if message == "" {
		message = "Invalid parameters"
	}
	return &ApiError{
		Code:    CodeInvalidParams,
		Message: message,
		Status:  http.StatusOK,
	}
}

// NewApiMissingParamsError creates an ApiError for missing required parameters
func NewApiMissingParamsError(message string) *ApiError {
	if message == "" {
		message = "Missing required parameters"
	}
	return &ApiError{
		Code:    CodeMissingParams,
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

// NewApiInvalidTokenError creates an ApiError for invalid token
func NewApiInvalidTokenError(message string) *ApiError {
	if message == "" {
		message = "Invalid token"
	}
	return &ApiError{
		Code:    CodeInvalidToken,
		Message: message,
		Status:  http.StatusUnauthorized,
	}
}

// NewApiTokenExpiredError creates an ApiError for expired token
func NewApiTokenExpiredError(message string) *ApiError {
	if message == "" {
		message = "Token expired"
	}
	return &ApiError{
		Code:    CodeTokenExpired,
		Message: message,
		Status:  http.StatusUnauthorized,
	}
}

// NewApiResourceExistsError creates an ApiError for resource already exists
func NewApiResourceExistsError(message string) *ApiError {
	if message == "" {
		message = "Resource already exists"
	}
	return &ApiError{
		Code:    CodeResourceExists,
		Message: message,
		Status:  http.StatusConflict,
	}
}

// NewApiOperationFailedError creates an ApiError for operation failed
func NewApiOperationFailedError(message string) *ApiError {
	if message == "" {
		message = "Operation failed"
	}
	return &ApiError{
		Code:    CodeOperationFailed,
		Message: message,
		Status:  http.StatusInternalServerError,
	}
}

// NewApiDatabaseError creates an ApiError for database errors
func NewApiDatabaseError(message string) *ApiError {
	if message == "" {
		message = "Database error"
	}
	return &ApiError{
		Code:    CodeDatabaseError,
		Message: message,
		Status:  http.StatusInternalServerError,
	}
}

// NewApiExternalServiceError creates an ApiError for external service errors
func NewApiExternalServiceError(message string) *ApiError {
	if message == "" {
		message = "External service error"
	}
	return &ApiError{
		Code:    CodeExternalError,
		Message: message,
		Status:  http.StatusBadGateway,
	}
}

// ============================================================================
// 辅助函数和工具方法
// ============================================================================

// IsApiError checks if the given error is an ApiError
func IsApiError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ApiError)
	return ok
}

// FromError converts a standard error to ApiError with default internal error code
func FromError(err error) *ApiError {
	if err == nil {
		return nil
	}

	if apiErr, ok := err.(*ApiError); ok {
		return apiErr
	}

	return NewApiInternalError(err.Error())
}

// FromErrorWithCode converts a standard error to ApiError with specified business code
func FromErrorWithCode(err error, code int) *ApiError {
	if err == nil {
		return nil
	}

	if apiErr, ok := err.(*ApiError); ok {
		return apiErr
	}

	return NewApiErrorWithoutStatus(code, err.Error())
}

// FromErrorWithStatus converts a standard error to ApiError with specified HTTP status
func FromErrorWithStatus(err error, status int) *ApiError {
	if err == nil {
		return nil
	}

	if apiErr, ok := err.(*ApiError); ok {
		return apiErr
	}

	return NewApiErrorWithDefaultCode(err.Error(), status)
}

// WrapError wraps an existing error with additional context message
func WrapError(err error, contextMsg string) *ApiError {
	if err == nil {
		return nil
	}

	if apiErr, ok := err.(*ApiError); ok {
		return &ApiError{
			Code:    apiErr.Code,
			Message: contextMsg + ": " + apiErr.Message,
			Status:  apiErr.Status,
		}
	}

	return NewApiInternalError(contextMsg + ": " + err.Error())
}

// WithMessage creates a new ApiError with the same code and status but different message
func (e *ApiError) WithMessage(message string) *ApiError {
	return &ApiError{
		Code:    e.Code,
		Message: message,
		Status:  e.Status,
	}
}

// WithCode creates a new ApiError with the same message and status but different business code
func (e *ApiError) WithCode(code int) *ApiError {
	return &ApiError{
		Code:    code,
		Message: e.Message,
		Status:  e.Status,
	}
}

// WithStatus creates a new ApiError with the same code and message but different HTTP status
func (e *ApiError) WithStatus(status int) *ApiError {
	return &ApiError{
		Code:    e.Code,
		Message: e.Message,
		Status:  status,
	}
}

// IsSystemError checks if the error is a system-level error (10000-19999)
func (e *ApiError) IsSystemError() bool {
	return e.Code >= 10000 && e.Code < 20000
}

// IsParamError checks if the error is a parameter error (20000-29999)
func (e *ApiError) IsParamError() bool {
	return e.Code >= 20000 && e.Code < 30000
}

// IsAuthError checks if the error is an authentication/authorization error (30000-39999)
func (e *ApiError) IsAuthError() bool {
	return e.Code >= 30000 && e.Code < 40000
}

// IsResourceError checks if the error is a resource error (40000-49999)
func (e *ApiError) IsResourceError() bool {
	return e.Code >= 40000 && e.Code < 50000
}

// IsBusinessError checks if the error is a business logic error (50000-59999)
func (e *ApiError) IsBusinessError() bool {
	return e.Code >= 50000 && e.Code < 60000
}

// IsDatabaseError checks if the error is a database error (60000-69999)
func (e *ApiError) IsDatabaseError() bool {
	return e.Code >= 60000 && e.Code < 70000
}

// IsExternalError checks if the error is an external service error (70000-79999)
func (e *ApiError) IsExternalError() bool {
	return e.Code >= 70000 && e.Code < 80000
}

// ============================================================================
// 错误码描述映射 (用于调试和日志记录)
// ============================================================================

// GetCodeDescription returns a human-readable description of the error code
func GetCodeDescription(code int) string {
	descriptions := map[int]string{
		// 系统级错误码
		CodeSuccess:            "Success",
		CodeSystemError:        "System Error",
		CodeInternalError:      "Internal Server Error",
		CodeServiceUnavailable: "Service Unavailable",
		CodeTimeout:            "Request Timeout",
		CodeTooManyRequests:    "Too Many Requests",
		CodeMaintenance:        "System Under Maintenance",

		// 请求参数错误码
		CodeBadRequest:      "Bad Request",
		CodeInvalidParams:   "Invalid Parameters",
		CodeMissingParams:   "Missing Required Parameters",
		CodeParamOutOfRange: "Parameter Out of Range",
		CodeInvalidFormat:   "Invalid Format",
		CodeInvalidJSON:     "Invalid JSON Format",
		CodeInvalidFileType: "Invalid File Type",
		CodeFileTooLarge:    "File Too Large",

		// 认证授权错误码
		CodeUnauthorized:       "Unauthorized",
		CodeInvalidToken:       "Invalid Token",
		CodeTokenExpired:       "Token Expired",
		CodeTokenMalformed:     "Token Malformed",
		CodeForbidden:          "Forbidden",
		CodeAccessDenied:       "Access Denied",
		CodeInvalidCredentials: "Invalid Credentials",
		CodeAccountLocked:      "Account Locked",
		CodeAccountDisabled:    "Account Disabled",

		// 资源错误码
		CodeNotFound:         "Resource Not Found",
		CodeResourceNotFound: "Specified Resource Not Found",
		CodeUserNotFound:     "User Not Found",
		CodeRecordNotFound:   "Record Not Found",
		CodePageNotFound:     "Page Not Found",

		// 业务逻辑错误码
		CodeBusinessError:    "Business Logic Error",
		CodeDataConflict:     "Data Conflict",
		CodeResourceExists:   "Resource Already Exists",
		CodeOperationFailed:  "Operation Failed",
		CodeInvalidOperation: "Invalid Operation",
		CodeStatusError:      "Status Error",
		CodeQuotaExceeded:    "Quota Exceeded",
		CodeDependencyError:  "Dependency Error",

		// 数据库错误码
		CodeDatabaseError:       "Database Error",
		CodeConnectionFailed:    "Database Connection Failed",
		CodeQueryFailed:         "Database Query Failed",
		CodeTransactionFailed:   "Database Transaction Failed",
		CodeConstraintViolation: "Database Constraint Violation",
		CodeDuplicateEntry:      "Duplicate Entry",

		// 外部服务错误码
		CodeExternalError:   "External Service Error",
		CodeThirdPartyError: "Third Party Service Error",
		CodeNetworkError:    "Network Error",
		CodeGatewayError:    "Gateway Error",
		CodeUpstreamError:   "Upstream Service Error",
	}

	if desc, exists := descriptions[code]; exists {
		return desc
	}
	return "Unknown Error Code"
}

// GetCodeDescription returns the description of the current error code
func (e *ApiError) GetCodeDescription() string {
	return GetCodeDescription(e.Code)
}
