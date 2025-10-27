package handler

import (
	"net/http"

	"github.com/wenpiner/last-admin-common/ctx/langctx"
)

// I18nHandler returns a middleware that inject accept-language into context
func I18nHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("accept-language")
		if lang == "" {
			lang = "zh-CN"
		}
		ctx := langctx.WithLangToContext(r.Context(), lang)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
