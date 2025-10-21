package handler

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest/enums"
)

// IpHandler returns a middleware that inject client ip into context
func IpHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := RequestRealIP(r)
		ctx := context.WithValue(r.Context(), enums.IpKey, ip)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequestRealIP
func RequestRealIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip
	}
	ip = r.Header.Get("X-Forwarded-For")
	for _, v := range strings.Split(ip, ",") {
		if net.ParseIP(v) != nil {
			return v
		}
	}

	ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	if net.ParseIP(ip) != nil {
		return ip
	}

	return ""
}
