package enums

type CtxKey string

const (
	// LangKey is the key for language in context
	LangKey CtxKey = "lang"
	// IpKey is the key for ip in context
	IpKey CtxKey = "client-ip"
	// DeptKey is the key for department id in context
	DeptKey CtxKey = "dept-id"
	// RoleKey is the key for role id in context
	RoleKey CtxKey = "role-id"
	// UserKey is the key for user id in context
	UserKey CtxKey = "user-id"
)
