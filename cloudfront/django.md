# Get the country from cloudfront using django

```
def my_view(request):
    country = request.META.get('HTTP_CLOUDFRONT_VIEWER_COUNTRY', None)
```
