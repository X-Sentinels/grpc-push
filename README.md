# grpc-push

### 编译
#### server
```
git clone https://github.com/X-Sentinels/grpc-push.git
cd grpc-push
go get ./...
cd server
build
```
#### client
```
cd grpc-push/client
build
```

#### 配置文件
```
{
	"grpc": {
		"listen": "0.0.0.0:50051" # grpc 服务的监听地址
	},
	"http": {
		"listen": "0.0.0.0:18080", # http 服务的监听地址
		"x-api-key": "26d66c3822bff031b2aacbcbfe3d9d14" # http 接口的密钥
	},
	"channel_cache": 10 # 消息队列的长度，如果拥塞则无法接受新的消息
}
```

#### 运行
##### server
```
# ./control start
server started..., pid=12213
# ./control stop
server stoped...
# ./control help
./control start|stop|restart|status|tail
```
##### client
```
# ./client -a {server_address:port}
2019/02/18 23:45:44 Calling Register RPC
```

#### 接口
grpc 注册过的 client 列表
```
# curl -v -H "X-API-KEY: 26d66c3822bff031b2aacbcbfe3d9d14" http://127.0.0.1:18080/api/v1/clients
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 18080 (#0)
> GET /api/v1/clients HTTP/1.1
> Host: 127.0.0.1:18080
> User-Agent: curl/7.47.0
> Accept: */*
> X-API-KEY: 26d66c3822bff031b2aacbcbfe3d9d14
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Mon, 18 Feb 2019 15:47:31 GMT
< Content-Length: 45
< 
* Connection #0 to host 127.0.0.1 left intact
{"data":["7","36"],"msg":"ok","success":true}
```
向 grpc client 推送消息
```
# curl -v -H "X-API-KEY: 26d66c3822bff031b2aacbcbfe3d9d14" -H "Content-Type: application/json" -H "clientName:37" -X POST -d '{"message":"hello world"}' http://127.0.0.1:18080/api/v1/push 
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 18080 (#0)
> POST /api/v1/push HTTP/1.1
> Host: 127.0.0.1:18080
> User-Agent: curl/7.47.0
> Accept: */*
> X-API-KEY: 26d66c3822bff031b2aacbcbfe3d9d14
> Content-Type: application/json
> clientName:37
> Content-Length: 25
> 
* upload completely sent off: 25 out of 25 bytes
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Mon, 18 Feb 2019 15:49:48 GMT
< Content-Length: 27
< 
* Connection #0 to host 127.0.0.1 left intact
{"msg":"ok","success":true}

```
client 侧响应 
```
#./client.exe -a 202.120.83.82:50051
2019/02/18 23:47:19 Calling Register RPC
2019/02/18 23:50:36 notice:"{\"message\":\"hello world\"}"  <nil>
```
