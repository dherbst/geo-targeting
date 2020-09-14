[![Go Report Card](https://goreportcard.com/badge/github.com/dherbst/geo-targeting?style=flat-square)](https://goreportcard.com/report/github.com/dherbst/geo-targeting)

# Geo-Targeting using CDNs
Many content distribution networks give you tools to target your content based on the country the user is in when they make a request to your web application.   Here are some examples:

## AWS - Cloudfront
Use the header `CloudFront-Viewer-Country`

### Get the country from cloudfront using django

```
def my_view(request):
    country = request.META.get('HTTP_CLOUDFRONT_VIEWER_COUNTRY', None)
```

### Country from cloudfront using golang and context

```
import (
       "context"
       "fmt"
       "github.com/dherbst/geo-targeting/src/geo"
       "net/http"
)

func init() {
	http.HandleFunc("/country", func(w http.ResponseWriter, r *http.Request) {
		ctx := geo.WithCountry(context.Background(), r)
		country, _ := geo.CountryFromContext(ctx)
		fmt.Fprintf(w, "Country=%v\n", country)
	})
}
```

## Google - AppEngine
Use the header `X-AppEngine-Country`

## Akamai - ion
Use the edgescape header `X-Akamai-Edgescape`, look for `country_code` in the string.

## Edgecast - WURL

Add `%{geo_country}` to a header using the advanced rules engine.

## Cloudflare
Use the header `CF-IPCountry`.

[Documentation](https://support.cloudflare.com/hc/en-us/articles/200168236-Configuring-Cloudflare-IP-Geolocation)
