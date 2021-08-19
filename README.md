# ProxyPrototype

To get the proxy running locally:

1. Download nginx source code version 1.20.1:
	http://nginx.org/en/download.html

2. Follow the instructions for building the patch for ngx_http_proxy_connect_module:
https://github.com/chobits/ngx_http_proxy_connect_module#proxy_connect

*should use proxy_connect_rewrite_1018.patch
*when at the configure step, include the following modules:

```
./configure 
--add-module=../ngx_http_proxy_connect_module (modify to own path)
--with-http_ssl_module 
--with-http_stub_status_module 
--with-http_realip_module 
--with-threads 
--with-stream 
--with-stream_ssl_preread_module 
--with-stream_ssl_module 
--with-openssl=../openssl (modify to own path)
```

3. Copy the contents of proxy_nginx.conf into the local nginx.conf file
- Can find location of nginx.conf file by running `nginx -t`

4. Start the proxy by running `sudo nginx -s reload`

Now, the proxy is running on localhost:7070, which can be passed into the `HTTP_PROXY` env variable on crdb
