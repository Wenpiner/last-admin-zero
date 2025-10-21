package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/rest/enums"
	"google.golang.org/grpc/metadata"
)

func TestTenantHandler(t *testing.T) {
	tenantID := "test-tenant-123"
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set(string(enums.TenantIdHeaderKey), tenantID)

	handler := TenantHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify tenant ID is in context
		ctx := r.Context()
		ctxTenantID := ctx.Value(enums.TenantKey)
		assert.Equal(t, tenantID, ctxTenantID)

		// Verify tenant ID is in gRPC metadata
		md, ok := metadata.FromOutgoingContext(ctx)
		assert.True(t, ok)
		tenantValues := md.Get(string(enums.TenantKey))
		assert.Equal(t, 1, len(tenantValues))
		assert.Equal(t, tenantID, tenantValues[0])

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("ok"))
		assert.Nil(t, err)
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "ok", resp.Body.String())
}

func TestTenantHandlerWithEmptyTenantID(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	// Don't set tenant ID header

	handler := TenantHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify empty tenant ID is in context
		ctx := r.Context()
		ctxTenantID := ctx.Value(enums.TenantKey)
		assert.Equal(t, "", ctxTenantID)

		// Verify empty tenant ID is in gRPC metadata
		md, ok := metadata.FromOutgoingContext(ctx)
		assert.True(t, ok)
		tenantValues := md.Get(string(enums.TenantKey))
		assert.Equal(t, 1, len(tenantValues))
		assert.Equal(t, "", tenantValues[0])

		w.WriteHeader(http.StatusOK)
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestTenantHandlerMultipleTenants(t *testing.T) {
	testCases := []string{
		"tenant-1",
		"tenant-2",
		"tenant-abc-123",
		"",
	}

	for _, tenantID := range testCases {
		t.Run(tenantID, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
			if tenantID != "" {
				req.Header.Set(string(enums.TenantIdHeaderKey), tenantID)
			}

			handler := TenantHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				ctxTenantID := ctx.Value(enums.TenantKey)
				assert.Equal(t, tenantID, ctxTenantID)
				w.WriteHeader(http.StatusOK)
			}))

			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, req)
			assert.Equal(t, http.StatusOK, resp.Code)
		})
	}
}

func TestTenantHandlerPreservesExistingContext(t *testing.T) {
	tenantID := "test-tenant"
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set(string(enums.TenantIdHeaderKey), tenantID)

	// Add some existing context value
	ctx := context.WithValue(req.Context(), "existing-key", "existing-value")
	req = req.WithContext(ctx)

	handler := TenantHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Verify existing value is preserved
		existingValue := ctx.Value("existing-key")
		assert.Equal(t, "existing-value", existingValue)

		// Verify tenant ID is added
		ctxTenantID := ctx.Value(enums.TenantKey)
		assert.Equal(t, tenantID, ctxTenantID)

		w.WriteHeader(http.StatusOK)
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

