# gowebunit
```
这个项目用于模拟喜马项目回调，打印接收到的回调参数

```

## 启动服务
```
go run main.go -config=./.env

```

## nginx配置示例
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

```

## 请求示例

### 通用通知回调-成功 `/commonnotify/callback01`
```
curl -X POST 'http://127.0.0.1:9909/commonnotify/callback01' \
-H 'Content-Type: application/json' \
-d '{
    "order_id": "order_001",
    "status": "success"
}'
```

### 通用通知回调-失败 `/commonnotify/callback02`
```
curl -X POST 'http://127.0.0.1:9909/commonnotify/callback02' \
-H 'Content-Type: application/json' \
-d '{
    "order_id": "order_001",
    "status": "fail"
}'
```

### 喜马拉雅电信通知 `/ximalaya/callback`

点播通知（POST form）：
```
curl -X POST 'http://127.0.0.1:9909/ximalaya/callback' \
-H 'Content-Type: application/x-www-form-urlencoded' \
-d 'cp_order_id=cp_order_001&fee=990&pay_type=2'
```

包月通知（GET）：
```
curl -X GET 'http://127.0.0.1:9909/ximalaya/callback?cp_order_id=cp_order_001&op_type=1'
```

### 喜马渠道方回调 `/ximachannel/callback`
```
curl -X POST 'http://127.0.0.1:9909/ximachannel/callback' \
-H 'Content-Type: application/json' \
-d '{
    "result_code": 0,
    "result_msg": "success",
    "cp_order_id": "cp_order_id值",
    "out_order_id": "out_order_id值",
    "fee": 990,
    "pay_type": 2,
    "out_ext": "",
    "sign": "sign值"
}'
```

经 nginx 代理访问示例：
```
curl -X POST 'https://api.example.com/testcallback/ximachannel/callback' \
-H 'Content-Type: application/json' \
-d '{
    "result_code": 0,
    "result_msg": "success",
    "cp_order_id": "cp_order_id值",
    "out_order_id": "out_order_id值",
    "fee": 990,
    "pay_type": 2,
    "out_ext": "",
    "sign": "sign值"
}'
```
