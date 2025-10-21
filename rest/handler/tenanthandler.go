package handler

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/enums"
	"google.golang.org/grpc/metadata"
)

// TenantHandler returns a middleware that inject tenant id into context
func TenantHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenantId := r.Header.Get(string(enums.TenantIdHeaderKey))
		ctx := context.WithValue(r.Context(), enums.TenantKey, tenantId)
		ctx = metadata.AppendToOutgoingContext(ctx, string(enums.TenantKey), tenantId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
