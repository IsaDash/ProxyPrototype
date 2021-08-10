FROM ubuntu:latest

COPY sources.list /etc/apt/sources.list

RUN apt-get update; \
    apt-get install -y libfontconfig1; \
    apt-get install -y libpcre3; \
    apt-get install -y libpcre3-dev; \
    apt-get install -y git; \
    apt-get install -y dpkg-dev; \
    apt-get install -y libpng-dev; \
    apt-get install -y libssl-dev; \
    apt-get autoclean && apt-get autoremove;

RUN apt-get source nginx; \
    git clone https://github.com/chobits/ngx_http_proxy_connect_module.git; \
    cd nginx-* && patch -p1 < ../ngx_http_proxy_connect_module/patch/proxy_connect_rewrite_1018.patch; \
    cd nginx-* && ./configure --add-module=/ngx_http_proxy_connect_module && make && make install;

COPY /certs etc/nginx/certs
COPY proxy_nginx.conf /etc/nginx/nginx.conf

EXPOSE 7000

CMD ["usr/local/nginx/sbin/nginx"]


