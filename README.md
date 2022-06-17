## lebron
Highly concurrent mall system built on go-zero

## Architecture 
A busy sum, when there is time to make up for it

## Series of courses
[go-zero 微服务实战系列（一、开篇）](https://mp.weixin.qq.com/s?__biz=Mzg2ODU1MTI0OA==&mid=2247485597&idx=1&sn=7e85894b7847cc50df51d66092792453&scene=21#wechat_redirect)

[go-zero 微服务实战系列（二、服务拆分）](https://mp.weixin.qq.com/s?__biz=Mzg2ODU1MTI0OA==&mid=2247485645&idx=1&sn=d329f56741dbe1f3e09713a6e4d1f7f0&scene=21#wechat_redirect)

[go-zero微服务实战系列（三、API定义和表结构设计）](https://mp.weixin.qq.com/s/ZWfzuJuJKeyJM3PMJ-SysQ)

[go-zero微服务实战系列（四、CRUD热身）](https://mp.weixin.qq.com/s/AIcJkMKTL0odqy1NzeJkxg)


## Business Function
![img.png](doc/imgs/img.png)

## Server Port 

### API

| server name    | port  |
| -------------- | ------|
| api bff        | 8001  |


### RPC
| server name    | port  |
| -------------- | ------|
| user rpc       | 9001  |
| product rpc    | 9002  |
| order rpc      | 9003  |