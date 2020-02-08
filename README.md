# Carrier

Carrier is an HTTP proxy that redirects your request to the host and scheme specified in the `X-Carrier-Target` header.

That's all.

## Run It

```bash
make image
docker run -it --rm --net host --env HTTP_PORT=80 carrier
```

## Examples

Make a request to Google:

```bash
curl -H "X-Carrier-Target: http://google.com" http://localhost:8080
...
< HTTP/1.1 301 Moved Permanently
< Location: http://www.google.com/
...
```

Make a request to Google (HTTPS):

```bash
curl -H "X-Carrier-Target: https://google.com" http://localhost:8080
...
< HTTP/1.1 301 Moved Permanently
< Location: https://www.google.com/
...
```

### HTTP/2

Carrier can also proxy HTTP/2 traffic, but requires TLS. Provide the certificate and key via:

```bash
docker run -it --rm --net host --env HTTP_PORT=443 TLS_CERT_FILE=cert.pem TLS_KEY_FILE=key.pem carrier
```

Then make an HTTP/2 request:

```bash
curl --http2-prior-knowledge -H "X-Carrier-Target: https://http2.golang.org" https://localhost/
```
