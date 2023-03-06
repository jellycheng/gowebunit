# gowebunit
```

```

## 启动服务
```
go run main.go -config=./.env

```

## nginx
```
upstream testservice {
    server 127.0.0.1:9909;
}

location ^~ /testcallback/ {
   proxy_pass http://testservice/;
   proxy_set_header X-Real-IP $remote_addr;
   proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
   proxy_set_header Host $host;
}

curl -X POST 'https://apicloud.qianguopai.com/testcallback/ximachannel/callback' \
-H 'Content-Type: application/json' \
-d '{
"result_code": 0,
"result_msg": "success",
"cp_order_id":"cp_order_id值",
"out_order_id":"out_order_id值",
"fee":990,
"pay_type":2,
"out_ext":"",
"sign":"sign值"
}'

```
