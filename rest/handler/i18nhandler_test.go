package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/rest/enums"
)

func TestI18nHandlerWithAcceptLanguage(t *testing.T) {
	lang := "en-US"
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("accept-language", lang)

	handler := I18nHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify language is in context
		ctx := r.Context()
		ctxLang := ctx.Value(enums.IpKey)
		assert.Equal(t, lang, ctxLang)

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("ok"))
		assert.Nil(t, err)
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "ok", resp.Body.String())
}

func TestI18nHandlerWithoutAcceptLanguage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	// Don't set accept-language header

	handler := I18nHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify default language is set to zh-CN
		ctx := r.Context()
		ctxLang := ctx.Value(enums.IpKey)
		assert.Equal(t, "zh-CN", ctxLang)

		w.WriteHeader(http.StatusOK)
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestI18nHandlerWithEmptyAcceptLanguage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("accept-language", "")

	handler := I18nHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify default language is set to zh-CN when header is empty
		ctx := r.Context()
		ctxLang := ctx.Value(enums.IpKey)
		assert.Equal(t, "zh-CN", ctxLang)

		w.WriteHeader(http.StatusOK)
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestI18nHandlerMultipleLanguages(t *testing.T) {
	testCases := []struct {
		name     string
		header   string
		expected string
	}{
		{"English", "en-US", "en-US"},
		{"Chinese", "zh-CN", "zh-CN"},
		{"French", "fr-FR", "fr-FR"},
		{"German", "de-DE", "de-DE"},
		{"Japanese", "ja-JP", "ja-JP"},
		{"Empty", "", "zh-CN"},
		{"Multiple", "en-US,zh-CN;q=0.9", "en-US,zh-CN;q=0.9"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
			if tc.header != "" {
				req.Header.Set("accept-language", tc.header)
			}

			handler := I18nHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				ctxLang := ctx.Value(enums.IpKey)
				assert.Equal(t, tc.expected, ctxLang)
				w.WriteHeader(http.StatusOK)
			}))

			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, req)
			assert.Equal(t, http.StatusOK, resp.Code)
		})
	}
}

func TestI18nHandlerPreservesExistingContext(t *testing.T) {
	lang := "es-ES"
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("accept-language", lang)

	// Add some existing context value
	ctx := req.Context()
	ctx = context.WithValue(ctx, "user-id", "user-123")
	req = req.WithContext(ctx)

	handler := I18nHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Verify existing value is preserved
		userID := ctx.Value("user-id")
		assert.Equal(t, "user-123", userID)

		// Verify language is added
		ctxLang := ctx.Value(enums.IpKey)
		assert.Equal(t, lang, ctxLang)

		w.WriteHeader(http.StatusOK)
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestI18nHandlerCallsNextHandler(t *testing.T) {
	handlerCalled := false
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("accept-language", "en-US")

	handler := I18nHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerCalled = true
		w.WriteHeader(http.StatusOK)
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.True(t, handlerCalled)
	assert.Equal(t, http.StatusOK, resp.Code)
}

