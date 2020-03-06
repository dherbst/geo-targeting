package geo

import (
	"context"
	"net/http"
	"testing"
)

func TestWithNoCountry(t *testing.T) {
	ctx := context.TODO()
	req, err := http.NewRequest("GET", "/example", nil)
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	ctx = WithCountry(ctx, req)

	country, ok := CountryFromContext(ctx)
	if !ok {
		t.Fatal("Did not map empty string to string")
	}
	if country != "" {
		t.Fatalf("Expected empty country, got %v\n", country)
	}
}

func TestWithCountry(t *testing.T) {
	ctx := context.TODO()
	req, err := http.NewRequest("GET", "/example", nil)
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	req.Header.Add("CloudFront-Viewer-Country", "US")
	ctx = WithCountry(ctx, req)

	country, ok := CountryFromContext(ctx)
	if !ok {
		t.Fatal("Did not get ok")
	}
	if country != "US" {
		t.Fatalf("Did not get US, got %v\n", country)
	}
}
