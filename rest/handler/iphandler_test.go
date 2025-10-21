package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/rest/enums"
)

func TestIpHandlerWithXRealIP(t *testing.T) {
	ip := "192.168.1.100"
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("X-Real-IP", ip)

	handler := IpHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify IP is in context
		ctx := r.Context()
		ctxIP := ctx.Value(enums.IpKey)
		assert.Equal(t, ip, ctxIP)

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("ok"))
		assert.Nil(t, err)
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "ok", resp.Body.String())
}

func TestIpHandlerWithXForwardedFor(t *testing.T) {
	ip := "10.0.0.1"
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("X-Forwarded-For", ip+",192.168.1.1")

	handler := IpHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctxIP := ctx.Value(enums.IpKey)
		assert.Equal(t, ip, ctxIP)

		w.WriteHeader(http.StatusOK)
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestIpHandlerWithRemoteAddr(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.RemoteAddr = "172.16.0.1:8080"

	handler := IpHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctxIP := ctx.Value(enums.IpKey)
		assert.Equal(t, "172.16.0.1", ctxIP)

		w.WriteHeader(http.StatusOK)
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestIpHandlerPriority(t *testing.T) {
	// X-Real-IP should have highest priority
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("X-Real-IP", "192.168.1.100")
	req.Header.Set("X-Forwarded-For", "10.0.0.1")
	req.RemoteAddr = "172.16.0.1:8080"

	handler := IpHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctxIP := ctx.Value(enums.IpKey)
		assert.Equal(t, "192.168.1.100", ctxIP)

		w.WriteHeader(http.StatusOK)
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestIpHandlerXForwardedForPriority(t *testing.T) {
	// X-Forwarded-For should have second priority
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("X-Forwarded-For", "10.0.0.1,192.168.1.1")
	req.RemoteAddr = "172.16.0.1:8080"

	handler := IpHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctxIP := ctx.Value(enums.IpKey)
		assert.Equal(t, "10.0.0.1", ctxIP)

		w.WriteHeader(http.StatusOK)
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestRequestRealIPWithXRealIP(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("X-Real-IP", "192.168.1.100")

	ip := RequestRealIP(req)
	assert.Equal(t, "192.168.1.100", ip)
}

func TestRequestRealIPWithXForwardedFor(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("X-Forwarded-For", "10.0.0.1,192.168.1.1")

	ip := RequestRealIP(req)
	assert.Equal(t, "10.0.0.1", ip)
}

func TestRequestRealIPWithRemoteAddr(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.RemoteAddr = "172.16.0.1:8080"

	ip := RequestRealIP(req)
	assert.Equal(t, "172.16.0.1", ip)
}

func TestRequestRealIPWithInvalidIPs(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("X-Real-IP", "invalid-ip")
	req.Header.Set("X-Forwarded-For", "invalid-ip-1,invalid-ip-2")
	req.RemoteAddr = "invalid-remote-addr"

	ip := RequestRealIP(req)
	assert.Equal(t, "", ip)
}

func TestRequestRealIPWithMixedValidInvalidIPs(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("X-Forwarded-For", "invalid-ip,10.0.0.1,192.168.1.1")

	ip := RequestRealIP(req)
	assert.Equal(t, "10.0.0.1", ip)
}

func TestRequestRealIPWithIPv6(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("X-Real-IP", "2001:db8::1")

	ip := RequestRealIP(req)
	assert.Equal(t, "2001:db8::1", ip)
}

func TestIpHandlerCallsNextHandler(t *testing.T) {
	handlerCalled := false
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("X-Real-IP", "192.168.1.100")

	handler := IpHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerCalled = true
		w.WriteHeader(http.StatusOK)
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.True(t, handlerCalled)
	assert.Equal(t, http.StatusOK, resp.Code)
}

