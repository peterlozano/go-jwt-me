# go-jwt-me
Proof of concept, jwt decoder endpoint.

Build:
`go build`

Start Server:
`./jwt-me`

Single request with curl:
`curl -H "Authorization: Bearer <JWT_TOKEN>" http://127.0.0.1:1323`

Benchmark with apache benchmark:
`ab -n 10000 -c 5 -H "Authorization: Bearer <JWT_TOKEN>" http://127.0.0.1:1323/`

