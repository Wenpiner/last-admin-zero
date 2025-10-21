package enums

type CtxKey string

type HeaderKey string

const (
	// TenantIdHeaderKey is the key for tenant id in header
	TenantIdHeaderKey HeaderKey = "tenant-id"
)

const (
	// LangKey is the key for language in context
	LangKey CtxKey = "lang"
	// IpKey is the key for ip in context
	IpKey CtxKey = "client-ip"
	// TenantKey is the key for tenant id in context
	TenantKey CtxKey = "tenant-id"
)
