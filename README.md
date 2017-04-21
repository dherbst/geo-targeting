# Geo-Targeting using CDNs
Many content distribution networks give you tools to target your content based on the country the user is in when they make a request to your web application.   Here are some examples:

## AWS - Cloudfront
Use the header `CloudFront-Viewer-Country`

### Get the country from cloudfront using django

```
def my_view(request):
    country = request.META.get('HTTP_CLOUDFRONT_VIEWER_COUNTRY', None)
```


## Google - AppEngine
Use the header `X-AppEngine-Country`

## Akamai - ion
Use the edgescape header `X-Akamai-Edgescape`, look for `country_code` in the string.

## Edgecast - WURL

Add `%{geo_country}` to a header using the advanced rules engine.
