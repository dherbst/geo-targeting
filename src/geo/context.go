package geo

import (
	"context"
	"net/http"
)

type geoKey string

const (
	countryKey = geoKey("countryKey")
)

// WithCountry adds the country code to the context
func WithCountry(parent context.Context, r *http.Request) context.Context {
	country := r.Header.Get("CloudFront-Viewer-Country")
	c := context.WithValue(parent, countryKey, country)
	return c
}

// CountryFromContext returns the country code string from the context, or false if it is
// not a string
func CountryFromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(countryKey).(string)
	return v, ok
}
