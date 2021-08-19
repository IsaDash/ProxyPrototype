# ProxyPrototype

This branch demonstrates TLS authentication between client and proxy.

Use OpenSSL to generate client and server certificates and keys, and modify the paths in client.go, server.go, and proxy_nginx.conf

The proxy and server are containerized and can be run by `docker compose up -d`

Run the client by running `go run client.go` in the client folder
