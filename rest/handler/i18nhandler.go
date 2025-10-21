package handler

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/enums"
)

// I18nHandler returns a middleware that inject accept-language into context
func I18nHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("accept-language")
		if lang == "" {
			lang = "zh-CN"
		}
		ctx := context.WithValue(r.Context(), enums.IpKey, lang)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
