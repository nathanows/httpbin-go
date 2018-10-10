# httpbin-go: HTTP Request & Response Service
httpbin-go is a Go port of the popular [requests/httpbin](https://github.com/requests/httpbin), and its corresponding Docker image [kennethreitz/httpbin](https://hub.docker.com/r/kennethreitz/httpbin/).

The original httpbin Docker container tips the scales at ~500MB, httpbin-go is **100x smaller** (~5MB) and **10x faster** (see performance comparison below).

### Build and Run Locally

### Docker Container
```
docker pull ndwhtlssthr/httpbin-go
```
# Endpoints Implemented
### HTTP
- [x] `/delete` [DELETE]
- [x] `/get` [GET]
- [x] `/patch` [PATCH]
- [x] `/post` [POST]
- [x] `/put` [PUT]

### Auth
- [x] `/basic-auth/{user}/{passwd}` [GET]
- [x] `/bearer` [GET]
- [ ] `/digest-auth/{qop}/{user}/{passwd}` [GET]
- [ ] `/digest-auth/{qop}/{user}/{passwd}/{algorithm}` [GET]
- [ ] `/digest-auth/{qop}/{user}/{passwd}/{algorithm}/{stale_after}` [GET]
- [x] `/hidden-basic-auth/{user}/{passwd}` [GET]

### Status Codes
- [x] `/status/{codes}` [DELETE, GET, PATCH, POST, PUT]

### Request Inspection
- [x] `/headers` [GET]
- [x] `/ip` [GET]
- [x] `/user-agent` [GET]

### Response Inspection
- [x] `/cache` [GET]
- [x] `/cache/{value}` [GET]
- [x] `/etag/{etag}` [GET]
- [x] `/response-headers` [GET, POST]

### Response Formats
- [ ] `/brotli` [GET]
- [ ] `/deflate` [GET]
- [x] `/deny` [GET]
- [x] `/encoding/utf8` [GET]
- [ ] `/gzip` [GET]
- [x] `/html` [GET]
- [x] `/json` [GET]
- [x] `/robots.txt` [GET]
- [x] `/xml` [GET]

### Dynamic Data
- [x] `/base64/{value}` [GET]
- [x] `/bytes/{n}` [GET]
- [x] `/delay/{delay}` [DELETE, GET, PATCH, POST, PUT]
- [x] `/drip` [GET]
- [x] `/links/{n}/{offset}` [GET]
- [x] `/range/{numbytes}` [GET]
- [x] `/stream-bytes/{n}` [GET]
- [x] `/stream/{n}` [GET]
- [x] `/uuid` [GET]

### Cookies
- [x] `/cookies` [GET]
- [x] `/cookies/delete` [GET]
- [x] `/cookies/set` [GET]
- [x] `/cookies/set/{name}/{value}` [GET]

### Images
- [x] `/image` [GET]
- [x] `/image/jpeg` [GET]
- [x] `/image/png` [GET]
- [x] `/image/svg` [GET]
- [x] `/image/webp` [GET]

### Redirects
- [ ] `/absolute-redirect/{n}` [GET]
- [x] `/redirect-to` [DELETE, GET, PATCH, POST, PUT]
- [ ] `/redirect-to/{n}` [GET]
- [ ] `/relative-redirect/{n}` [GET]

### Anything
- [x] `/anything` [DELETE, GET, PATCH, POST, PUT]
- [x] `/anything/{anything}` [DELETE, GET, PATCH, POST, PUT]
