## nginx 配置

在请求(obtain)期间9980端口会处理 challenge

```nginx
## http/ 80 server
location /.well-known/acme-challenge {
    proxy_set_header Host $host;
    proxy_set_header X-Real_IP $remote_addr;
    proxy_set_header X-Forwarded-For $remote_addr:$remote_port;
    proxy_pass http://127.0.0.1:9980;
}
```